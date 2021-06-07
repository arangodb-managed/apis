//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { Empty as arangodb_cloud_common_v1_Empty } from '../../common/v1/common'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'
import { Version as arangodb_cloud_common_v1_Version } from '../../common/v1/common'
import { Deployment as arangodb_cloud_data_v1_Deployment } from '../../data/v1/data'
import { Deployment_ModelSpec as arangodb_cloud_data_v1_Deployment_ModelSpec } from '../../data/v1/data'

// File: prepaid/v1/prepaid.proto
// Package: arangodb.cloud.prepaid.v1

// CloneFromBackupRequest is used to create a new deployment based on PrepaidDeployment
// with prepaid_deployment_id and restore data from Backup with given backup_id
export interface CloneFromBackupRequest {
  // Identifier of prepaid deployment
  // string
  prepaid_deployment_id?: string;
  
  // Identifier of the backup to restore from
  // Backup specification has to match prepaid_deployment specification in order to succeed
  // string
  backup_id?: string;
}

// CreateDeploymentRequest is used to create a new deployment based on PrepaidDeployment
// with given prepaid_deployment_id and attach it to this PrepaidDeployment
export interface CreateDeploymentRequest {
  // Identifier of prepaid deployment to use as a specification and attach the newly created Deployment to
  // string
  prepaid_deployment_id?: string;
  
  // Identifier of the project that owns the newly created deployment.
  // string
  project_id?: string;
  
  // Optional identifier of IP allowlist to use for this deployment.
  // string
  ipallowlist_id?: string;
  
  // ArangoDB version to use for this deployment.
  // string
  version?: string;
  
  // CreateDeploymentRequest_CertificateSpec
  certificates?: CreateDeploymentRequest_CertificateSpec;
}
export interface CreateDeploymentRequest_CertificateSpec {
  // Identifier of the CACertificate used to sign TLS certificates for the deployment.
  // If you change this value after the creation of the deployment a complete
  // rotation of the deployment is required, which will result in some downtime.
  // string
  ca_certificate_id?: string;
  
  // Zero or more DNS names to include in the TLS certificate of the deployment.
  // string
  alternate_dns_names?: string[];
}

// ListPrepaidDeploymentsRequest is used to request a list of PrepaidDeployments for
// organization with given organization_id
export interface ListPrepaidDeploymentsRequest {
  // identifier of the organization to get a list of prepaid deployments for
  // string
  organization_id?: string;
  
  // common listing options
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// A PrepaidDeployment contains all attributes of a future deployment that is already paid for.
export interface PrepaidDeployment {
  // System identifier of the prepaid deployment.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value.
  // string
  url?: string;
  
  // The name of prepaid deployment, not related to created deployment
  // string
  name?: string;
  
  // An optional description for prepaid deployment, not related to created deployment
  // string
  description?: string;
  
  // Identifier of an organization that owns this prepaid deployment
  // string
  organization_id?: string;
  
  // Identifier of the region in which a deployment is going to be created.
  // string
  region_id?: string;
  
  // The creation timestamp of the prepaid deployment
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the prepaid deployment
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Set when this deployment is deleted.
  // boolean
  is_deleted?: boolean;
  
  // Start of the period for which the PrepaidDeployment was purchased
  // googleTypes.Timestamp
  starts_at?: googleTypes.Timestamp;
  
  // End of the period for which the PrepaidDeployment was purchased
  // googleTypes.Timestamp
  ends_at?: googleTypes.Timestamp;
  
  // Identifier of the support plan selected for this prepaid deployment.
  // string
  support_plan_id?: string;
  
  // Model specification for created deployment
  // arangodb.cloud.data.v1.Deployment.ModelSpec
  model?: arangodb_cloud_data_v1_Deployment_ModelSpec;
  
  // Identifies the addons that will be used on the deployment
  // string
  addons?: string[];
  
  // PrepaidDeployment_Status
  status?: PrepaidDeployment_Status;
}

// Status of the prepaid deployment
// All members of this field are read-only.
export interface PrepaidDeployment_Status {
  // id of created deployment
  // if there is no deployment associated with this prepaid deployment it's empty
  // string
  deployment_id?: string;
  
  // timestamp when the deployment was created for or attached to PrepaidDeployment
  // googleTypes.Timestamp
  attached_at?: googleTypes.Timestamp;
  
  // timestamp when the deployment was detached from PrepaidDeployment
  // googleTypes.Timestamp
  detached_at?: googleTypes.Timestamp;
}

// PrepaidDeploymentList contains a list of PrepaidDeployment items
export interface PrepaidDeploymentList {
  // prepaid deployment items
  // PrepaidDeployment
  items?: PrepaidDeployment[];
}

// UpdateDeploymentRequest is used to update deployment attached to PrepaidDeployment
// with give prepaid_deployment_id
export interface UpdateDeploymentRequest {
  // Identifier of prepaid deployment
  // string
  prepaid_deployment_id?: string;
}

// PrepaidService is the API used to configure prepaid objects.
export interface IPrepaidService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Fetch all prepaid deployments for organization.
  // Required permissions:
  // - prepaid.prepaiddeployment.list on the organization
  ListPrepaidDeployments: (req: ListPrepaidDeploymentsRequest) => Promise<PrepaidDeploymentList>;
  
  // Fetch a deployment by its id.
  // Required permissions:
  // - prepaid.prepaiddeployment.get on the deployment identified by the given ID
  GetPrepaidDeployment: (req: arangodb_cloud_common_v1_IDOptions) => Promise<PrepaidDeployment>;
  
  // Creates a new deployment from a prepaid deployment and attached the newly created deployment to the prepaid deployment.
  // Required permissions:
  // - prepaid.prepaiddeployment.create on the project in which the deployment is going to be created
  // - prepaid.prepaiddeployment.get on the deployment identified by the given prepaid_deployment_id
  CreateDeployment: (req: CreateDeploymentRequest) => Promise<arangodb_cloud_data_v1_Deployment>;
  
  // Update the deployment by prepaid deployment's id
  // Required permissions:
  // - data.deployment.update on the deployment attached to the prepaid deployment
  UpdateDeployment: (req: UpdateDeploymentRequest) => Promise<arangodb_cloud_data_v1_Deployment>;
  
  // Creates a cloned deployment from a backup and attaches it to the prepaid deployment. This takes the deployment specification from the prepaid deployment, which must match the specification mentioned in the backup.
  // Required permissions:
  // - prepaid.prepaiddeployment.create on the project in which the deployment is going to be created
  // - prepaid.prepaiddeployment.get on the prepaid deployment identified by the given prepaid_deployment_id
  CloneDeploymentFromBackup: (req: CloneFromBackupRequest) => Promise<arangodb_cloud_data_v1_Deployment>;
}

// PrepaidService is the API used to configure prepaid objects.
export class PrepaidService implements IPrepaidService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/prepaid/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all prepaid deployments for organization.
  // Required permissions:
  // - prepaid.prepaiddeployment.list on the organization
  async ListPrepaidDeployments(req: ListPrepaidDeploymentsRequest): Promise<PrepaidDeploymentList> {
    const path = `/api/prepaid/v1/organizations/${encodeURIComponent(req.organization_id || '')}/prepaiddeployments`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.post(url, undefined);
  }
  
  // Fetch a deployment by its id.
  // Required permissions:
  // - prepaid.prepaiddeployment.get on the deployment identified by the given ID
  async GetPrepaidDeployment(req: arangodb_cloud_common_v1_IDOptions): Promise<PrepaidDeployment> {
    const path = `/api/prepaid/v1/prepaiddeployments/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Creates a new deployment from a prepaid deployment and attached the newly created deployment to the prepaid deployment.
  // Required permissions:
  // - prepaid.prepaiddeployment.create on the project in which the deployment is going to be created
  // - prepaid.prepaiddeployment.get on the deployment identified by the given prepaid_deployment_id
  async CreateDeployment(req: CreateDeploymentRequest): Promise<arangodb_cloud_data_v1_Deployment> {
    const url = `/api/prepaid/v1/prepaiddeployments/${encodeURIComponent(req.prepaid_deployment_id || '')}/create`;
    return api.post(url, req);
  }
  
  // Update the deployment by prepaid deployment's id
  // Required permissions:
  // - data.deployment.update on the deployment attached to the prepaid deployment
  async UpdateDeployment(req: UpdateDeploymentRequest): Promise<arangodb_cloud_data_v1_Deployment> {
    const url = `/api/prepaid/v1/prepaiddeployments/${encodeURIComponent(req.prepaid_deployment_id || '')}/update`;
    return api.post(url, req);
  }
  
  // Creates a cloned deployment from a backup and attaches it to the prepaid deployment. This takes the deployment specification from the prepaid deployment, which must match the specification mentioned in the backup.
  // Required permissions:
  // - prepaid.prepaiddeployment.create on the project in which the deployment is going to be created
  // - prepaid.prepaiddeployment.get on the prepaid deployment identified by the given prepaid_deployment_id
  async CloneDeploymentFromBackup(req: CloneFromBackupRequest): Promise<arangodb_cloud_data_v1_Deployment> {
    const url = `/api/prepaid/v1/prepaiddeployments/${encodeURIComponent(req.prepaid_deployment_id || '')}/clone`;
    return api.post(url, req);
  }
}
