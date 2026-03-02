//
// DISCLAIMER
//
// Copyright 2020-2026 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

package v1

const (
	// UsageItems kind

	// UsageItemKindV2CPUHour (Billing 2.0) indicates a UsageItem that contains
	// the number of CPU resources allocated for one hour by a deployment (or member of a deployment).
	UsageItemKindV2CPUHour = "CPUHour"

	// UsageItemKindV2MemoryHour (Billing 2.0) indicates a UsageItem that contains
	// the number of memory resources allocated for one hour by a deployment (or member of a deployment).
	UsageItemKindV2MemoryHour = "MemoryHour"

	// UsageItemKindV2StorageHour (Billing 2.0) indicates a UsageItem that contains
	// the number of storage resources allocated for one hour by a deployment (or member of a deployment).
	UsageItemKindV2StorageHour = "StorageHour"

	// UsageItemKindV2StoragePerformanceHour (Billing 2.0) indicates a UsageItem that contains
	// the amount of storage performance resources allocated for one hour by a deployment (or member of a deployment).
	UsageItemKindV2StoragePerformanceHour = "StoragePerformanceHour"

	// UsageItemKindV2GPUHour (Billing 2.0) indicates a UsageItem that contains
	// the number of GPU resources allocated for one hour by a deployment (or member of a deployment).
	UsageItemKindV2GPUHour = "GPUHour"

	// UsageItemKindV2NetworkSize (Billing 2.0) indicates a UsageItem that contains
	// the GiB of the Transfer Network allocated for a deployment (or member of a deployment).
	UsageItemKindV2NetworkSize = "NetworkUsageSize"

	// UsageItemKindV2CloudStorageHour (Billing 2.0) indicates a UsageItem that contains the amount of Cloud Storage usage,
	// e.g., for backup, audit, platform, etc., allocated/used for one hour.
	UsageItemKindV2CloudStorageHour = "CloudStorageHour"

	// UsageItemKindV2DeploymentAEU (Billing 2.0) indicates a UsageItem that contains deployment AEU usage,
	// where AEU = Deployment Size x Deployment Node Count x Deployment AEU Base x Deployment Type Ratio.
	UsageItemKindV2DeploymentAEU = "DeploymentAEU"

	/*
		###
		Deprecated section - to be removed in future releases:
		###
	*/

	// UsageItemKindNetworkTransferSize indicates a UsageItem that contains
	// the amount of network traffic caused by a deployment (or member of a deployment).
	// UsageItems of this kind will be closed automatically, so no open items can exist.
	// The timespan of this UsageItem is normally 24h (unless the deployment is deleted)
	UsageItemKindNetworkTransferSize = "NetworkTransferSize"

	// Deprecated: UsageItemKindDeploymentSize indicates a UsageItem that contains cloud
	// resources for Deployment.
	// UsageItems of this kind will not be closed automatically, so open items can exist.
	// This is replaced by CPUHour + MemoryHour + StorageHour usage kinds
	UsageItemKindDeploymentSize = "DeploymentSize"

	// Deprecated: UsageItemKindBackupStorageSize indicates a UsageItem that contains
	// the amount of cloud storage used by backups of a deployment.
	// UsageItems of this kind will not be closed automatically, so open items can exists.
	// This is replaced by CloudStorageHour usage kind
	UsageItemKindBackupStorageSize = "BackupStorageSize"

	// Deprecated: UsageItemKindAuditLogSize indicates a UsageItem that contains
	// the amount of resources used by audit log (of a deployment).
	// UsageItems of this kind will be closed automatically, so no open items can exists.
	// The timespan of this UsageItem is normally 24h (unless the auditlog or deployment is deleted)
	// This is replaced by StorageHour usage kind
	UsageItemKindAuditLogSize = "AuditLogSize"

	// Deprecated: UsageItemKindAuditLogStorageSize indicates a UsageItem that contains
	// the amount of cloud storage used by audit log (of a deployment).
	// UsageItems of this kind will not be closed automatically, so open items can exists.
	// This is replaced by CloudStorageHour usage kind
	UsageItemKindAuditLogStorageSize = "AuditLogStorageSize"

	// Deprecated: UsageItemKindNotebookSize indicates a UsageItem that contains
	// the number of resources used by the notebooks (of a deployment).
	// UsageItems of this kind will not be closed automatically, so open items can exists.
	UsageItemKindNotebookSize = "NotebookSize"
	// Deprecated: UsageItemKindMLServicesSize indicates a UsageItem that contains
	// the number of resources used by a MLServices resources (of a deployment).
	// UsageItems of this kind will not be closed automatically, so open items can exist.
	UsageItemKindMLServicesSize = "MLServicesSize"
	// Deprecated: UsageItemKindMLJobSize indicates a UsageItem that contains
	// the number of resources used by a ML job (of a MLServices).
	// UsageItems of this kind are closed on creation itself, so open items cannot exist.
	UsageItemKindMLJobSize = "MLJobSize"
	// Deprecated: UsageItemKindGraphAnalyticsJobSize indicates a UsageItem that contains
	// the number of resources used by a GraphAnalytics job.
	// UsageItems of this kind are closed on creation, so open items cannot exist.
	UsageItemKindGraphAnalyticsJobSize = "GraphAnalyticsJobSize"
)
