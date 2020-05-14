//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { Empty as arangodb_cloud_common_v1_Empty } from '../../common/v1/common'
import { Version as arangodb_cloud_common_v1_Version } from '../../common/v1/common'
import { Deployment as arangodb_cloud_data_v1_Deployment } from '../../data/v1/data'

// File: replication/v1/replication.proto
// Package: arangodb.cloud.replication.v1

// CloneDeploymentFromBackupRequest defines a request object for clone deployment call.
export interface CloneDeploymentFromBackupRequest {
  // The ID of the backup to clone a deployment from.
  // string
  backup_id?: string;
  
  // Region ID
  // string
  region_id?: string;
}

// ReplicationService is the API used to replicate a deployment.
export interface IReplicationService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Takes a backup and creates a deployment from it. For all intents and purposes this new deployment
  // will be the same as the deployment at that exact moment when the backup was taken from it. This means that
  // the new deployment will be in the same region and in the same project and use the same provider as the old
  // deployment did. Furthermore, the new deployment will have the same server settings ( count, mode, replication
  // factor ) as the old deployment did at the time of taking the backup. After the new deployment successfully
  // started, the backup will be used to restore the data into the new deployment. The new deployment will have
  // a different endpoint, and the password will also be reset for it. All other user settings will remain the same.
  // The old deployment will not be touched.
  // Required permissions:
  // - replication.deployment.clone-from-backup
  CloneDeploymentFromBackup: (req: CloneDeploymentFromBackupRequest) => Promise<arangodb_cloud_data_v1_Deployment>;
}

// ReplicationService is the API used to replicate a deployment.
export class ReplicationService implements IReplicationService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/replication/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Takes a backup and creates a deployment from it. For all intents and purposes this new deployment
  // will be the same as the deployment at that exact moment when the backup was taken from it. This means that
  // the new deployment will be in the same region and in the same project and use the same provider as the old
  // deployment did. Furthermore, the new deployment will have the same server settings ( count, mode, replication
  // factor ) as the old deployment did at the time of taking the backup. After the new deployment successfully
  // started, the backup will be used to restore the data into the new deployment. The new deployment will have
  // a different endpoint, and the password will also be reset for it. All other user settings will remain the same.
  // The old deployment will not be touched.
  // Required permissions:
  // - replication.deployment.clone-from-backup
  async CloneDeploymentFromBackup(req: CloneDeploymentFromBackupRequest): Promise<arangodb_cloud_data_v1_Deployment> {
    const url = `/api/replication/v1/backup/${encodeURIComponent(req.backup_id || '')}/clone`;
    return api.post(url, req);
  }
}
