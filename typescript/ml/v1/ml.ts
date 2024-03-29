//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { Empty as arangodb_cloud_common_v1_Empty } from '../../common/v1/common'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { Version as arangodb_cloud_common_v1_Version } from '../../common/v1/common'

// File: ml/v1/ml.proto
// Package: arangodb.cloud.ml.v1
export interface ListMLServicesSizesRequest {
  // Optional ID of the Deployment for which sizes are being requested.
  // If set, the response will exclude any sizes that are unavailable for the specified deployment model.
  // string
  deployment_id?: string;
}

// MLServices is a single resource which represents the state and configuration
// of ML Services (ArangoGraphML) for a deployment specified by deployment_id.
export interface MLServices {
  // Identifier of the deployment for this MLServices resource.
  // This is a ready-only value.
  // string
  deployment_id?: string;
  
  // Set to true if ML services are enabled for this deployment.
  // boolean
  enabled?: boolean;
  
  // Size to use for the ML Jobs.
  // Use `ListMLServicesSizes` to get a list of available sizes.
  // If unspecified, the MLServiceSize marked as `is_default` is used.
  // This is an optional field.
  // string
  size?: string;
  
  // The creation timestamp of the MLServices.
  // This also serves as a timestamp of when MLServices were first enabled.
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // Status of the MLServices.
  // This is a read-only value.
  // Status
  status?: Status;
}

// MLServicesSize represents the resources allocated for MLServices.
// Note that the specified configuration is applied for the ML jobs.
export interface MLServicesSize {
  // Identifier of the size configuration.
  // string
  size_id?: string;
  
  // If set, this is the default size when unspecified in MLServices.
  // boolean
  is_default?: boolean;
  
  // Amount of CPU allocated (in vCPU units)
  // number
  cpu?: number;
  
  // Amount of Memory allocated (in GB)
  // number
  memory?: number;
  
  // Amount of GPUs allocated
  // number
  gpu?: number;
}

// List of MLServicesSize.
export interface MLServicesSizeList {
  // Items in this list.
  // MLServicesSize
  items?: MLServicesSize[];
}

// Status of a single ArangoGraphML component.
export interface ServiceStatus {
  // Type of service.
  // Should be one of: [training|prediction|projects]
  // string
  type?: string;
  
  // Set to true if the service is available.
  // Every service is always in ONLY ONE of the following states: (available|failed)
  // boolean
  available?: boolean;
  
  // Set to true if the service is in a failed state.
  // Every service is always in ONLY ONE of the following states: (available|failed)
  // boolean
  failed?: boolean;
  
  // Resource usage information for this service.
  // ServiceStatus_Usage
  usage?: ServiceStatus_Usage;
  
  // Number of replicas running for this service.
  // number
  replicas?: number;
}

// Resource usage for this service.
export interface ServiceStatus_Usage {
  // Last known memory usage in bytes
  // number
  last_memory_usage?: number;
  
  // Last known CPU usage in vCPU units
  // number
  last_cpu_usage?: number;
  
  // Last known memory limit in bytes
  // number
  last_memory_limit?: number;
  
  // Last known CPU limit in vCPU units
  // number
  last_cpu_limit?: number;
}

// Status of the MLServices.
// Note: All fields are read-only.
export interface Status {
  // Overall status of where the MLServices resource is in its lifecycle at a given time.
  // It will contain only one of the following values:
  // "Bootstrapping"  - ArangoDB Deployment is being bootstrapped with the required databases, schemas and data.
  // "Initialising"   - The services needed for ArangoGraphML are being installed.
  // "Running"        - ArangoGraphML is setup and running correctly.
  // "Error"          - Indicates that there was an error with setting up ArangoGraphML. Check `message` field for additional info.
  // "Hibernated"     - Indicates that ArangoGraphML and all its associated services are hibernated.
  // string
  phase?: string;
  
  // Supporting information about the phase of MLServices (such as error messages in case of failures).
  // string
  message?: string;
  
  // The timestamp of when this status was last updated.
  // googleTypes.Timestamp
  last_updated_at?: googleTypes.Timestamp;
  
  // Status of each ArangoGraphML components/services.
  // ServiceStatus
  services?: ServiceStatus[];
  
  // Total number of hours ML Jobs have run for this Deployment.
  // number
  hours_used?: number;
  
  // Total number of runtime hours allowed for ML Jobs for this Deployment.
  // Set to 0 if unlimited (i.e, no restriction).
  // number
  hours_allowed?: number;
  
  // Timestamp after which MLServices are no longer usable.
  // This is set during trial use.
  // If unset, no expiry.
  // googleTypes.Timestamp
  expires_at?: googleTypes.Timestamp;
  
  // Timestamp of when MLServices were last enabled for this deployment.
  // googleTypes.Timestamp
  last_enabled_at?: googleTypes.Timestamp;
}

// MLService is the API used to configure ArangoML on ArangoGraph Insights Platform.
export interface IMLService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Get an existing MLServices resource for a given deployment (specified by the id).
  // Required permissions:
  // - ml.mlservices.get
  GetMLServices: (req: arangodb_cloud_common_v1_IDOptions) => Promise<MLServices>;
  
  // Update an existing MLServices resource. If it does not exist, this will create a new one.
  // Pass the desired updated state of MLServices to this call.
  // Required permissions:
  // - ml.mlservices.update
  UpdateMLServices: (req: MLServices) => Promise<MLServices>;
  
  // List the available size configurations for MLServices.
  // Note that the returned size specifications are applied for ML Jobs.
  // Required permissions:
  // - ml.mlservicessize.list on the deployment (if deployment_id is provided)
  // - None, authenticated only
  ListMLServicesSizes: (req: ListMLServicesSizesRequest) => Promise<MLServicesSizeList>;
}

// MLService is the API used to configure ArangoML on ArangoGraph Insights Platform.
export class MLService implements IMLService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/ml/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Get an existing MLServices resource for a given deployment (specified by the id).
  // Required permissions:
  // - ml.mlservices.get
  async GetMLServices(req: arangodb_cloud_common_v1_IDOptions): Promise<MLServices> {
    const path = `/api/ml/v1/mlservices/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Update an existing MLServices resource. If it does not exist, this will create a new one.
  // Pass the desired updated state of MLServices to this call.
  // Required permissions:
  // - ml.mlservices.update
  async UpdateMLServices(req: MLServices): Promise<MLServices> {
    const url = `/api/ml/v1/mlservices/${encodeURIComponent(req.deployment_id || '')}`;
    return api.put(url, req);
  }
  
  // List the available size configurations for MLServices.
  // Note that the returned size specifications are applied for ML Jobs.
  // Required permissions:
  // - ml.mlservicessize.list on the deployment (if deployment_id is provided)
  // - None, authenticated only
  async ListMLServicesSizes(req: ListMLServicesSizesRequest): Promise<MLServicesSizeList> {
    const path = `/api/ml/v1/sizes`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
}
