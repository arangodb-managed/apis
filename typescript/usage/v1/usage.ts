//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { Empty as arangodb_cloud_common_v1_Empty } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'
import { Version as arangodb_cloud_common_v1_Version } from '../../common/v1/common'

// File: usage/v1/usage.proto
// Package: arangodb.cloud.usage.v1

// Request arguments for ListUsageItems
export interface ListUsageItemsRequest {
  // Request usage items for the organization with this id.
  // This is a required field.
  // string
  organization_id?: string;
  
  // Request usage items that overlaps in time with the time period that starts with this timestamp (inclusive).
  // This is a required field.
  // googleTypes.Timestamp
  from?: googleTypes.Timestamp;
  
  // Request usage items that overlaps in time with the time period that ends with this timestamp (exclusive).
  // This is a required field.
  // googleTypes.Timestamp
  to?: googleTypes.Timestamp;
  
  // Sort descending (new to old) on started_at field (default is ascending).
  // boolean
  sort_descending?: boolean;
  
  // Limit to usage items of this kind
  // string
  kind?: string;
  
  // Standard list options
  // This is an optional field.
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
  
  // Limit to usage items for the resource with this URL.
  // This is an optional field.
  // string
  resource_url?: string;
  
  // Limit to usage items for the resource with this kind.
  // This is an optional field.
  // string
  resource_kind?: string;
  
  // Limit to usage items for the project with this id.
  // This is an optional field.
  // string
  project_id?: string;
  
  // Limit to usage items for the deployment with this id.
  // This is an optional field.
  // string
  deployment_id?: string;
  
  // Limit to usage items for deployments with this node size.
  // This is an optional field.
  // string
  node_size_id?: string;
  
  // Limit to usage items for deployments in this region.
  // This is an optional field.
  // string
  region_id?: string;
  
  // If set, limit to usage items that have no invoice_id set.
  // boolean
  has_no_invoice_id?: boolean;
  
  // If set, limit to usage items that have an invoice_id set.
  // boolean
  has_invoice_id?: boolean;
  
  // If set, limit to usage items that have the invoice_id set to this specific value.
  // This is an optional field.
  // string
  invoice_id?: string;
  
  // Request usage items that start at or after given timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  not_start_before?: googleTypes.Timestamp;
  
  // Limit to usage items for the deployment with this model.
  // string
  included_tier_ids?: string[];
  
  // Limit to usage items for the deployment with this model.
  // string
  excluded_tier_ids?: string[];
}

// A UsageItem message contained usage tracking information for a tracked
// resource (usually deployment) in a specific time period.
export interface UsageItem {
  // System identifier of the usage item.
  // string
  id?: string;
  
  // URL of this resource
  // string
  url?: string;
  
  // Kind of usage item
  // string
  kind?: string;
  
  // Identification of the resource covered by this usage item
  // UsageItem_Resource
  resource?: UsageItem_Resource;
  
  // This usage item covers a time period that starts at this timestamp
  // googleTypes.Timestamp
  starts_at?: googleTypes.Timestamp;
  
  // This usage item covers a time period that ends at this timestamp.
  // If the usage item has not yet ended, this field is is set to the current time.
  // googleTypes.Timestamp
  ends_at?: googleTypes.Timestamp;
  
  // Set when this usage item has ended.
  // boolean
  has_ended?: boolean;
  
  // Identifier of the tier the organization was using at the start of this usage period.
  // string
  tier_id?: string;
  
  // Identifier of the invoice that includes this usage item.
  // The usage item must be ended when this field it set.
  // string
  invoice_id?: string;
  
  // Amount of (computer) resources used by the resource covered by this usage item.
  // This field is only set when the usage item is of kind DeploymentSize.
  // UsageItem_DeploymentSize
  deployment_size?: UsageItem_DeploymentSize;
  
  // Amount of network traffic used by the resource covered by this usage item.
  // This field is only set when the usage item is of kind NetworkTransferSize.
  // UsageItem_NetworkTransferSize
  network_transfer_size?: UsageItem_NetworkTransferSize;
  
  // Amount of backup related cloud storage used by the resource covered by this usage item.
  // This field is only set when the usage item is of kind BackupStorageSize.
  // UsageItem_BackupStorageSize
  backup_storage_size?: UsageItem_BackupStorageSize;
  
  // Amount of audit log related resources used by the resource covered by this usage item.
  // This field is only set when the usage item is of kind AuditLogSize.
  // UsageItem_AuditLogSize
  auditlog_size?: UsageItem_AuditLogSize;
  
  // Amount of audit log storage related resources used by the resource covered by this usage item.
  // This field is only set when the usage item is of kind AuditLogCloudSize.
  // UsageItem_AuditLogStorageSize
  auditlog_storage_size?: UsageItem_AuditLogStorageSize;
  
  // Amount of cloud resources used by a given Notebook.
  // This field is only set when the usage item is of kind NotebookSize.
  // UsageItem_NotebookSize
  notebook_size?: UsageItem_NotebookSize;
  
  // Amount of compute resources used by a given MLServices resource.
  // This field is only set when the usage item is of kind MLServicesSize.
  // UsageItem_MLServicesSize
  mlservices_size?: UsageItem_MLServicesSize;
  
  // Amount of compute resources used by a given ML job.
  // This field is only set when the usage item is of kind MLJobSize.
  // UsageItem_MLJobSize
  mljob_size?: UsageItem_MLJobSize;
  
  // Amount of compute resources used by a given ML job.
  // This field is only set when the usage item is of kind MLJobSize.
  // UsageItem_GraphAnalyticsJobSize
  graphanalyticsjob_size?: UsageItem_GraphAnalyticsJobSize;
}

// Amount of audit log related resources used by the resource covered by this usage item.
export interface UsageItem_AuditLogSize {
  // Type of destination.
  // Possible values are: "cloud", "https-post"
  // string
  destination_type?: string;
  
  // Number of events used by audit log.
  // This is the value of DestinationCounters.events (of the timespan covered by this usage item).
  // number
  event_count?: number;
  
  // Number of bytes used by audit log.
  // This is the total of DestinationCounters.bytes_succeeded and bytes_failed (of the timespan covered by this usage item)
  // and depending on the destination type send to the cloud or used as the body of a https post request.
  // number
  event_size?: number;
  
  // Number of https post invocations used by audit log.
  // This is the total of DestinationCounters.https_posts_succeeded and https_posts_failed (of the timespan covered by this usage item).
  // Set when destination_type is "https-post" only
  // number
  https_post_count?: number;
}

// Amount of audit log storage related resources used by the resource covered by this usage item.
// When this usage type is used, the audit log destination_type is "cloud".
export interface UsageItem_AuditLogStorageSize {
  // Amount of cloud storage (in bytes) used by audit log.
  // number
  cloud_storage_size?: number;
}

// Amount of backup related cloud storage used by the resource covered by this usage item.
export interface UsageItem_BackupStorageSize {
  // Amount of cloud storage (in bytes) used by backups of a deployment.
  // number
  cloud_storage_size?: number;
}

// Amount of (computer) resources used by the resource covered by this usage item.
export interface UsageItem_DeploymentSize {
  // Number of coordinators of the deployment
  // number
  coordinators?: number;
  
  // Amount of memory (in GB) allocated for each coordinator.
  // number
  coordinator_memory_size?: number;
  
  // Number of dbservers of the deployment
  // number
  dbservers?: number;
  
  // Amount of memory (in GB) allocated for each dbserver.
  // number
  dbserver_memory_size?: number;
  
  // Amount of disk space (in GB) allocated for each dbserver.
  // number
  dbserver_disk_size?: number;
  
  // Number of agents of the deployment
  // number
  agents?: number;
  
  // Amount of memory (in GB) allocated for each agent.
  // number
  agent_memory_size?: number;
  
  // Amount of disk space (in GB) allocated for each agent.
  // number
  agent_disk_size?: number;
  
  // Identifier of the node-size used for this deployment (empty for flexible)
  // string
  node_size_id?: string;
  
  // Identifier of disk performance used for this deployment (if any).
  // string
  disk_performance_id?: string;
  
  // List of identifiers of addons used by the deployment.
  // string
  addon_ids?: string[];
  
  // IF true the deployment was in paused state
  // boolean
  is_paused?: boolean;
  
  // Number of gateways of the deployment.
  // number
  gateways?: number;
  
  // Amount of memory (in GB) allocated for each gateway.
  // number
  gateway_memory_size?: number;
}

// Amount of compute resources used by a Graph Analytics Job.
export interface UsageItem_GraphAnalyticsJobSize {
  // Amount of memory (in GB) allocated for the job.
  // number
  memory_size?: number;
  
  // Amount of CPU units allocated to the job.
  // 1 CPU unit equals 1 physical / virtual CPU.
  // number
  cpu_size?: number;
  
  // Amount of GPU units allocated to the job.
  // 1 CPU unit equals 1 physical / virtual GPU.
  // number
  gpu_size?: number;
  
  // ID of the ML job.
  // string
  job_id?: string;
  
  // Type of Graph Analytics job.
  // string
  job_type?: string;
}

// Amount of compute resources used by a ML Job.
export interface UsageItem_MLJobSize {
  // Amount of memory (in GB) allocated for the job.
  // number
  memory_size?: number;
  
  // Amount of CPU units allocated to the job.
  // 1 CPU unit equals 1 physical / virtual CPU.
  // number
  cpu_size?: number;
  
  // Amount of GPU units allocated to the job.
  // 1 CPU unit equals 1 physical / virtual GPU.
  // number
  gpu_size?: number;
  
  // ID of the ML job.
  // string
  job_id?: string;
  
  // Type of ML Job (training, prediction, etc.)
  // string
  job_type?: string;
}

// Amount of compute resources used by a given MLServices resource.
export interface UsageItem_MLServicesSize {
  // Number of training API servers.
  // number
  training_apis?: number;
  
  // Amount of memory (in GB) allocated for each training API server.
  // number
  training_api_memory_size?: number;
  
  // Amount of CPU units allocated to the training API server.
  // 1 CPU unit equals 1 physical / virtual CPU.
  // number
  training_api_cpu_size?: number;
  
  // Number of prediction API servers.
  // number
  prediction_apis?: number;
  
  // Amount of memory (in GB) allocated for each prediction API server.
  // number
  prediction_api_memory_size?: number;
  
  // Amount of CPU units allocated to the prediction API server.
  // 1 CPU unit equals 1 physical / virtual CPU.
  // number
  prediction_api_cpu_size?: number;
  
  // Number of projects API servers.
  // number
  projects_apis?: number;
  
  // Amount of memory (in GB) allocated for each projects API server.
  // number
  projects_api_memory_size?: number;
  
  // Amount of CPU units allocated to the projects API server.
  // 1 CPU unit equals 1 physical / virtual CPU.
  // number
  projects_api_cpu_size?: number;
}

// Amount of network traffic used by the resource covered by this usage item.
export interface UsageItem_NetworkTransferSize {
  // The destination (or source) the network traffic going to (or coming from)
  // Can be 'Internet' or 'InCluster'
  // An empty string means 'Internet'
  // string
  destination?: string;
  
  // Total amount of network ingress traffic (in bytes) caused by the use of a deployment.
  // Destination 'Internet': This is traffic coming from the internet, so excluding inner cluster traffic
  // Destination 'InCluster': This is in-cluster traffic only
  // This is always excluding backup traffic (downloads).
  // number
  total_transfer_ingress_size?: number;
  
  // Total amount of network egress traffic (in bytes) caused by the use of a deployment.
  // Destination 'Internet': This is traffic going to the internet, so excluding inner cluster traffic
  // Destination 'InCluster': This is in-cluster traffic only
  // This is always excluding backup traffic (uploads).
  // Note: In the future we want to split between cross_region_transfer_x and inner_region_transfer_x,
  // the total_transfer_x is the sum of these 2. Inner region can be cross availability zone.
  // number
  total_transfer_egress_size?: number;
}

// Amount of cloud resources used by a given Notebook.
// This field is only set when the usage item is of kind NotebookSize.
export interface UsageItem_NotebookSize {
  // Amount of CPU units allocated to the notebook.
  // 1 CPU unit equals 1 physical / virtual CPU.
  // number
  cpu_size?: number;
  
  // Amount of memory allocated to the notebook (in GiB).
  // number
  memory_size?: number;
  
  // Size of the disk allocated to the notebook (in GiB).
  // number
  disk_size?: number;
  
  // If the Notebook is paused (Hibernated phase).
  // boolean
  is_paused?: boolean;
  
  // ID of the type of Notebook that is being used.
  // string
  notebook_model_id?: string;
  
  // Amount of GPU units allocated to the notebook.
  // 1 GPU unit equals 1 physical / virtual GPU.
  // number
  gpu_size?: number;
}
export interface UsageItem_Resource {
  // System identifier of the resource that this usage item covers.
  // string
  id?: string;
  
  // URL of the resource that this usage item covers
  // string
  url?: string;
  
  // Kind of resource that this usage item covers.
  // string
  kind?: string;
  
  // Human readable description of the resource that this usage item covers.
  // string
  description?: string;
  
  // Identifier of the organization that owns the resource that this usage item covers.
  // string
  organization_id?: string;
  
  // Name of the organization that owns the resource that this usage item covers.
  // string
  organization_name?: string;
  
  // Identifier of the project that owns the resource that this usage item covers.
  // This field is optional when the kind is AuditLogSize, depending if the audit log is for a deployment (set) or ArangoGraph Insights Platform (empty)
  // string
  project_id?: string;
  
  // Name of the project that owns the resource that this usage item covers.
  // This field is optional when the kind is AuditLogSize, depending if the audit log is for a deployment (set) or ArangoGraph Insights Platform (empty)
  // string
  project_name?: string;
  
  // Identifier of the deployment that owns the resource that this usage item covers.
  // This field is optional when the kind is AuditLogSize, depending if the audit log is for a deployment (set) or ArangoGraph Insights Platform (empty)
  // string
  deployment_id?: string;
  
  // Name of the deployment that owns the resource that this usage item covers.
  // This field is optional when the kind is AuditLogSize, depending if the audit log is for a deployment (set) or ArangoGraph Insights Platform (empty)
  // string
  deployment_name?: string;
  
  // Name of the deployment member that owns the resource that this usage item covers.
  // This field is only set when the usage item is specific for a member of the deployment (e.g. network transfer)
  // string
  deployment_member_name?: string;
  
  // Identifier of the cloud provider that is used to run the deployment.
  // This field is optional when the kind is AuditLogSize, depending if the audit log is for a deployment (set) or ArangoGraph Insights Platform (empty)
  // string
  cloud_provider_id?: string;
  
  // Identifier of the cloud region that is used to run the deployment.
  // This field is optional when the kind is AuditLogSize, depending if the audit log is for a deployment (set) or ArangoGraph Insights Platform (empty)
  // string
  cloud_region_id?: string;
  
  // Identifier of the support plan that is attached to the deployment.
  // This field is not filled-out when the kind is AuditLogSize
  // string
  support_plan_id?: string;
  
  // Model of the deployment
  // This field is not filled-out when the kind is AuditLogSize
  // string
  deployment_model?: string;
  
  // Identifier of the PrepaidDeployment that this deployment is attached to (if any)
  // string
  prepaid_deployment_id?: string;
  
  // Timestamp when the prepaid_deployment starts (relevant when prepaid_deployment_id is set only)
  // googleTypes.Timestamp
  prepaid_deployment_starts_at?: googleTypes.Timestamp;
  
  // Timestamp when the prepaid_deployment ends (relevant when prepaid_deployment_id is set only)
  // googleTypes.Timestamp
  prepaid_deployment_ends_at?: googleTypes.Timestamp;
  
  // Identifiers of the credit bundles (if any) that were used to pay for this usage item.
  // string
  credit_bundle_ids?: string[];
}

// List of UsageItems.
export interface UsageItemList {
  // UsageItem
  items?: UsageItem[];
}

// UsageService is the API used to fetch usage tracking information.
export interface IUsageService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Fetch all UsageItem resources in the organization identified by the given
  // organization ID that match the given criteria.
  // Required permissions:
  // - usage.usageitem.list on the organization identified by the given organization ID
  ListUsageItems: (req: ListUsageItemsRequest) => Promise<UsageItemList>;
}

// UsageService is the API used to fetch usage tracking information.
export class UsageService implements IUsageService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/usage/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all UsageItem resources in the organization identified by the given
  // organization ID that match the given criteria.
  // Required permissions:
  // - usage.usageitem.list on the organization identified by the given organization ID
  async ListUsageItems(req: ListUsageItemsRequest): Promise<UsageItemList> {
    const path = `/api/usage/v1/organization/${encodeURIComponent(req.organization_id || '')}/usageitems`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
}
