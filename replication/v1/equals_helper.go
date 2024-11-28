//
// DISCLAIMER
//
// Copyright 2022-2024 ArangoDB GmbH, Cologne, Germany
//

package v1

import common "github.com/arangodb-managed/apis/common/v1"

// Equals returns true when source and other have the same values.
func (source *DeploymentReplication_Status) Equals(other *DeploymentReplication_Status) bool {
	return source.GetPhase() == other.GetPhase() &&
		source.GetMessage() == other.GetMessage() &&
		source.GetProgress() == other.GetProgress() &&
		source.GetSyncEndpoint() == other.GetSyncEndpoint() &&
		source.GetForwarderEndpoint() == other.GetForwarderEndpoint()
}

// Equals returns true when source and other have the same values.
func (source *DeploymentMigration_Status) Equals(other *DeploymentMigration_Status) bool {
	return source.GetPhase() == other.GetPhase() &&
		source.GetDescription() == other.GetDescription() &&
		source.GetBackupId() == other.GetBackupId() &&
		source.GetTargetDeploymentId() == other.GetTargetDeploymentId() &&
		common.TimestampsEqual(source.GetLastUpdatedAt(), other.GetLastUpdatedAt())
}
