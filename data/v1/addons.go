//
// DISCLAIMER
//
// Copyright 2021-2026 ArangoDB GmbH, Cologne, Germany
//
// Author Robert Stam
//

package v1

const (
	// Addon IDs

	// AddonIDIAMProvider is ID of the IAM provider addon.
	// That addon is used when a deployment is using a custom IAM provider.
	AddonIDIAMProvider = "iamprovider"

	// AddonIDAuditLog is ID of the Audit Log addon.
	// That addon is used when a deployment has an attached audit log.
	AddonIDAuditLog = "auditlog"

	// AddonIDPriviteEndpointService is ID of the Private endpoint service addon.
	// That addon is used when a deployment is using a private endpoint service.
	AddonIDPrivateEndpointService = "privateendpointservice"

	// AddonIDDedicatedDataCluster is ID of the Dedicated Data Cluster addon.
	// That addon is used when an organization runs its deployments on data
	// clusters dedicated (private) to the organization. Unlike the other addons
	// it belongs to a data cluster, not to a single deployment; in Billing 2.0
	// it is charged hourly via AddonHour usage items.
	AddonIDDedicatedDataCluster = "dedicateddatacluster"
)
