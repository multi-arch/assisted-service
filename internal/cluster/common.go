package cluster

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/openshift/assisted-service/internal/cluster/validations"
	"github.com/openshift/assisted-service/internal/common"
	eventgen "github.com/openshift/assisted-service/internal/common/events"
	eventsapi "github.com/openshift/assisted-service/internal/events/api"
	"github.com/openshift/assisted-service/internal/network"
	"github.com/openshift/assisted-service/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

const (
	MinMastersNeededForInstallation = 3
)

const (
	StatusInfoReady                           = "Cluster ready to be installed"
	StatusInfoInsufficient                    = "Cluster is not ready for install"
	statusInfoInstalling                      = "Installation in progress"
	statusInfoFinalizing                      = "Finalizing cluster installation"
	statusInfoInstalled                       = "Cluster is installed"
	StatusInfoDegraded                        = "Cluster is installed but degraded"
	StatusInfoNotAllWorkersInstalled          = "Cluster is installed but some workers did not join"
	statusInfoPreparingForInstallation        = "Preparing cluster for installation"
	statusInfoPreparingForInstallationTimeout = "Preparing cluster for installation timeout"
	statusInfoFinalizingTimeout               = "Cluster installation timeout while finalizing"
	statusInfoPendingForInput                 = "User input required"
	statusInfoError                           = "cluster has hosts in error"
	statusInfoTimeout                         = "cluster installation timed out while pending user action (a manual booting from installation disk)"
	statusInfoAddingHosts                     = "cluster is adding hosts to existing OCP cluster"
	statusInfoInstallingPendingUserAction     = "Cluster has hosts pending user action"
	statusInfoUnpreparingHostExists           = "At least one host has stopped preparing for installation"
	statusInfoClusterFailedToPrepare          = "Cluster failed to prepare for installation"
)

func updateClusterStatus(ctx context.Context, log logrus.FieldLogger, db *gorm.DB, clusterId strfmt.UUID, srcStatus string,
	newStatus string, statusInfo string, events eventsapi.Handler, extra ...interface{}) (*common.Cluster, error) {
	var cluster *common.Cluster
	var err error
	extra = append(append(make([]interface{}, 0), "status", newStatus, "status_info", statusInfo), extra...)

	if newStatus != srcStatus {
		now := strfmt.DateTime(time.Now())
		extra = append(extra, "status_updated_at", now)

		installationCompletedStatuses := []string{models.ClusterStatusInstalled, models.ClusterStatusError, models.ClusterStatusCancelled}
		if funk.ContainsString(installationCompletedStatuses, swag.StringValue(&newStatus)) {
			extra = append(extra, "install_completed_at", now)
		}
	}

	if newStatus != srcStatus {
		extra = append(extra, "trigger_monitor_timestamp", time.Now())
	}

	if cluster, err = UpdateCluster(log, db, clusterId, srcStatus, extra...); err != nil ||
		swag.StringValue(cluster.Status) != newStatus {
		return nil, errors.Wrapf(err, "failed to update cluster %s state from %s to %s",
			clusterId, srcStatus, newStatus)
	}

	if newStatus != srcStatus {
		eventgen.SendClusterStatusUpdatedEvent(ctx, events, clusterId, *cluster.Status, statusInfo)
		log.Infof("cluster %s has been updated with the following updates %+v", clusterId, extra)
	}

	return cluster, nil
}

func updateLogsProgress(log logrus.FieldLogger, db *gorm.DB, c *common.Cluster, progress string) error {
	var updates map[string]interface{}

	switch progress {
	case string(models.LogsStateRequested):
		updates = map[string]interface{}{
			"logs_info":                  progress,
			"controller_logs_started_at": strfmt.DateTime(time.Now()),
		}
	default:
		updates = map[string]interface{}{
			"logs_info": progress,
		}
	}

	result := db.Model(&common.Cluster{}).Where("id = ?", c.ID.String()).Updates(updates)

	if err := result.Error; err != nil {
		log.WithError(err).Errorf("could not update log progress %v on cluster %s", updates, *c.ID)
		return err
	}
	if result.RowsAffected == 1 {
		updatedCluster, err := common.GetClusterFromDB(db, *c.ID, common.UseEagerLoading)
		if err != nil {
			log.WithError(err).Errorf("could not update log progress %v on cluster %s", updates, *c.ID)
			return err
		}
		*c = *updatedCluster
	}

	log.Infof("cluster %s has been updated with the following log progress %+v", *c.ID, updates)
	return nil
}

func ClusterExists(db *gorm.DB, clusterId strfmt.UUID) bool {
	where := make(map[string]interface{})
	return clusterExistsInDB(db, clusterId, where)
}

func clusterExistsInDB(db *gorm.DB, clusterId strfmt.UUID, where map[string]interface{}) bool {
	where["id"] = clusterId.String()
	var cluster common.Cluster
	return db.Select("id").Take(&cluster, where).Error == nil
}

func UpdateCluster(log logrus.FieldLogger, db *gorm.DB, clusterId strfmt.UUID, srcStatus string, extra ...interface{}) (*common.Cluster, error) {
	updates := make(map[string]interface{})
	updateChangesVips := false

	if len(extra)%2 != 0 {
		return nil, errors.Errorf("invalid update extra parameters %+v", extra)
	}
	for i := 0; i < len(extra); i += 2 {
		if extra[i].(string) == "api_vip" || extra[i].(string) == "ingress_vip" {
			updateChangesVips = true
		}
		updates[extra[i].(string)] = extra[i+1]
	}

	// Query by <cluster-id, status>
	// Status is required as well to avoid races between different components.
	dbReply := db.Model(&common.Cluster{}).Where("id = ? and status = ?", clusterId, srcStatus).Updates(updates)

	if dbReply.Error != nil {
		return nil, errors.Wrapf(dbReply.Error, "failed to update cluster %s", clusterId)
	}

	if dbReply.RowsAffected == 0 && !clusterExistsInDB(db, clusterId, updates) {
		return nil, errors.Errorf("failed to update cluster %s. nothing has changed", clusterId)
	}

	// (MGMT-9915)
	// The block below is responsible for updating plural VIP structure in any scenario when an
	// update of singular VIP has been requested via the `updates[]` mechanism. This is
	// a workaround for e.g. a scenario when SNO updates its VIP from the Host inventory and not
	// using the regular get&set for the Cluster.*Vip field.
	//
	// This part covers specifically the scenario as follows
	//   * updates[api_vip] or updates[ingress_vip] contains a value
	//   * Cluster.ApiVips and Cluster.IngressVips are not updated
	//
	// Without this block below, such a scenario would lead to unsynchronized state of the DB.
	if updateChangesVips {
		updatedCluster, err := common.GetClusterFromDB(db, clusterId, common.UseEagerLoading)
		if err != nil {
			return nil, err
		}
		apiVips, err := validations.HandleApiVipBackwardsCompatibility(*updatedCluster.ID, updatedCluster.APIVip, nil)
		if err != nil {
			return nil, err
		}
		ingressVips, err := validations.HandleIngressVipBackwardsCompatibility(*updatedCluster.ID, updatedCluster.IngressVip, nil)
		if err != nil {
			return nil, err
		}

		if err = network.UpdateVipsTables(db,
			&common.Cluster{Cluster: models.Cluster{
				ID:          updatedCluster.ID,
				APIVip:      updatedCluster.APIVip,
				APIVips:     apiVips,
				IngressVip:  updatedCluster.IngressVip,
				IngressVips: ingressVips,
			}},
			true,
			true,
		); err != nil {
			return nil, err
		}
	}

	return common.GetClusterFromDB(db, clusterId, common.UseEagerLoading)
}

func getKnownMastersNodesIds(c *common.Cluster, db *gorm.DB) ([]*strfmt.UUID, error) {
	var cluster *common.Cluster
	var err error
	var masterNodesIds []*strfmt.UUID
	if cluster, err = common.GetClusterFromDB(db, *c.ID, common.UseEagerLoading); err != nil {
		return nil, errors.Errorf("cluster %s not found", c.ID)
	}

	allowedStatuses := []string{models.HostStatusKnown, models.HostStatusPreparingForInstallation}
	for _, host := range cluster.Hosts {
		if common.GetEffectiveRole(host) == models.HostRoleMaster && funk.ContainsString(allowedStatuses, swag.StringValue(host.Status)) {
			masterNodesIds = append(masterNodesIds, host.ID)
		}
	}
	return masterNodesIds, nil
}

func NumberOfWorkers(c *common.Cluster) int {
	num := 0
	for _, host := range c.Hosts {
		if common.GetEffectiveRole(host) != models.HostRoleWorker {
			continue
		}
		num += 1
	}
	return num
}

func HostsInStatus(c *common.Cluster, statuses []string) (int, int) {
	mappedMastersByRole := MapMasterHostsByStatus(c)
	mappedWorkersByRole := MapWorkersHostsByStatus(c)
	mastersInSomeInstallingStatus := 0
	workersInSomeInstallingStatus := 0

	for _, status := range statuses {
		mastersInSomeInstallingStatus += len(mappedMastersByRole[status])
		workersInSomeInstallingStatus += len(mappedWorkersByRole[status])
	}
	return mastersInSomeInstallingStatus, workersInSomeInstallingStatus
}

func MapMasterHostsByStatus(c *common.Cluster) map[string][]*models.Host {
	return mapHostsByStatus(c, models.HostRoleMaster)
}

func MapWorkersHostsByStatus(c *common.Cluster) map[string][]*models.Host {
	return mapHostsByStatus(c, models.HostRoleWorker)
}

func mapHostsByStatus(c *common.Cluster, role models.HostRole) map[string][]*models.Host {
	hostMap := make(map[string][]*models.Host)
	for _, host := range c.Hosts {
		if role != "" && common.GetEffectiveRole(host) != role {
			continue
		}
		if _, ok := hostMap[swag.StringValue(host.Status)]; ok {
			hostMap[swag.StringValue(host.Status)] = append(hostMap[swag.StringValue(host.Status)], host)
		} else {
			hostMap[swag.StringValue(host.Status)] = []*models.Host{host}
		}
	}
	return hostMap
}

func MapHostsByStatus(c *common.Cluster) map[string][]*models.Host {
	return mapHostsByStatus(c, "")
}

func UpdateMachineNetwork(db *gorm.DB, cluster *common.Cluster, machineNetwork []string) error {
	if len(machineNetwork) > 2 {
		return common.NewApiError(http.StatusInternalServerError,
			errors.Errorf("for cluster %s received request to update %d machine networks", cluster.ID, len(machineNetwork)))
	}

	previousPrimaryMachineNetwork := ""
	previousSecondaryMachineNetwork := ""
	if network.IsMachineCidrAvailable(cluster) {
		previousPrimaryMachineNetwork = network.GetMachineCidrById(cluster, 0)
		previousSecondaryMachineNetwork = network.GetMachineCidrById(cluster, 1)
	}

	if (len(machineNetwork) > 0 && machineNetwork[0] != previousPrimaryMachineNetwork) || (len(machineNetwork) > 1 && machineNetwork[1] != previousSecondaryMachineNetwork) {
		err := db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("cluster_id = ?", *cluster.ID).Delete(&models.MachineNetwork{}).Error; err != nil {
				err = errors.Wrapf(err, "failed to delete machine networks of cluster %s", *cluster.ID)
				return common.NewApiError(http.StatusInternalServerError, err)
			}
			for _, cidr := range machineNetwork {
				if cidr != "" {
					machineNetwork := &models.MachineNetwork{
						ClusterID: *cluster.ID,
						Cidr:      models.Subnet(cidr),
					}
					if err := tx.Save(machineNetwork).Error; err != nil {
						err = errors.Wrapf(err, "failed to update cluster machineNetwork %s of cluster %s", cidr, *cluster.ID)
						return common.NewApiError(http.StatusInternalServerError, err)
					}
				}
			}
			if err := tx.Model(&common.Cluster{}).Where("id = ?", cluster.ID.String()).Updates(map[string]interface{}{
				"machine_network_cidr_updated_at": time.Now()}).Error; err != nil {
				err = errors.Wrapf(err, "failed to update machine networks timestamp in cluster %s", *cluster.ID)
				return common.NewApiError(http.StatusInternalServerError, err)
			}
			return nil
		})
		return err
	}
	return nil
}
