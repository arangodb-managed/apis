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

// File: notebook/v1/notebook.proto
// Package: arangodb.cloud.notebook.v1

// Request for creating a notebook.
export interface CreateNotebookRequest {
  // ID of the Deployment this notebook belongs to.
  // string
  deployment_id?: string;
  
  // Name of the notebook.
  // string
  name?: string;
  
  // Description of the notebook.
  // string
  description?: string;
  
  // Model specification for the notebook.
  // Note: This field is read-only after creating.
  // ModelSpec
  model?: ModelSpec;
}

// Request for listing notebooks.
export interface ListNotebookRequest {
  // List notebooks for this deployment ID.
  // string
  deployment_id?: string;
  
  // Optional common list options, the context_id is ignored
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Model specification for the notebook.
export interface ModelSpec {
  // Type of model being used.
  // Currently supported:
  // - basic [CPU: 1, Memory: 4GiB]
  // string
  model?: string;
  
  // Disk size required by the notebook instance.
  // Should be expressed as power-of-two: Mi, Gi, Ti, Pi, etc.
  // string
  disk_size?: string;
}

// Contains the specification and status of a given notebook instance.
export interface Notebook {
  // ID of the Notebook.
  // string
  id?: string;
  
  // ID of the Deployment this notebook belongs to.
  // string
  deployment_id?: string;
  
  // URL of the Notebook.
  // Note: This is a read-only property.
  // string
  url?: string;
  
  // Name of the notebook.
  // string
  name?: string;
  
  // Description of the notebook.
  // string
  description?: string;
  
  // If the notebook should be paused.
  // string
  paused?: string;
  
  // Identifier of the user that created this notebook.
  // string
  created_by_id?: string;
  
  // Time at which this notebook was created.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // Model specification for the notebook.
  // Note: This field is read-only after creating.
  // ModelSpec
  model?: ModelSpec;
  
  // If the notebook should be deleted.
  // boolean
  deleted?: boolean;
  
  // Time at which this notebook was deleted.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Status of the notebook. Represents the state of the notebook as observed by the controller.
  // Note: all fields in this block are read-only.
  // Status
  status?: Status;
}

// List of Notebooks.
export interface NotebookList {
  // Notebook
  items?: Notebook[];
}

// Status of the notebook. Represents the state of the notebook as observed by the controller.
// Note: all fields in this block are read-only.
export interface Status {
  // Where the notebook is in its lifecycle at any given time.
  // Should only contain only one of the following values:
  // "Initialising"   - Notebook is initialising.
  // "Running"        - Notebook is running.
  // "Hibernating"    - Notebook is hibernating.
  // "Error"          - Notebook is in an errored state. Additional information can be obtained from `message` field.
  // string
  phase?: string;
  
  // Supporting information about the notebook phase - such as error messages in case of failures.
  // string
  message?: string;
  
  // The last time this notebook was updated.
  // googleTypes.Timestamp
  last_updated_at?: googleTypes.Timestamp;
  
  // Resource usage of the notebook.
  // Status_Usage
  usage?: Status_Usage;
}

// Resource usage of the notebook.
export interface Status_Usage {
  // CPU usage in percentage.
  // number
  cpu?: number;
  
  // Amount of memory being consumed.
  // number
  memory?: number;
  
  // Amount of disk space used.
  // number
  disk?: number;
}

// Request to pause a notebook.
export interface UpdateNotebookRequest {
  // ID of the notebook.
  // string
  id?: string;
  
  // Set if notebook should be paused.
  // boolean
  paused?: boolean;
  
  // Name of the notebook.
  // string
  name?: string;
  
  // Description of the notebook.
  // string
  description?: string;
}

// Notebook service is used to manage notebooks.
export interface INotebookService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Get a Notebook using its ID.
  // Required permissions:
  // - notebook.notebook.get
  GetNotebook: (req: arangodb_cloud_common_v1_IDOptions) => Promise<Notebook>;
  
  // Create a new Notebook by specifying its configuration.
  // Required permissions:
  // - notebook.notebook.create
  CreateNotebook: (req: CreateNotebookRequest) => Promise<Notebook>;
  
  // Delete an existing notebook using its ID.
  // This initially marks the notebook for deletion. It is deleted from CP once all its child resources are deleted.
  // Required permissions:
  // - notebook.notebook.delete
  DeleteNotebook: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
  
  // Update an existing notebook. Returns updated Notebook.
  // Required permissions:
  // - notebook.notebook.update
  UpdateNotebook: (req: UpdateNotebookRequest) => Promise<Notebook>;
  
  // List all notebooks for a deployment.
  // Required permissions:
  // - notebook.notebook.list
  ListNotebooks: (req: ListNotebookRequest) => Promise<NotebookList>;
}

// Notebook service is used to manage notebooks.
export class NotebookService implements INotebookService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/notebook/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Get a Notebook using its ID.
  // Required permissions:
  // - notebook.notebook.get
  async GetNotebook(req: arangodb_cloud_common_v1_IDOptions): Promise<Notebook> {
    const path = `/api/notebook/v1/notebook/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new Notebook by specifying its configuration.
  // Required permissions:
  // - notebook.notebook.create
  async CreateNotebook(req: CreateNotebookRequest): Promise<Notebook> {
    const url = `/api/notebook/v1/notebook`;
    return api.put(url, req);
  }
  
  // Delete an existing notebook using its ID.
  // This initially marks the notebook for deletion. It is deleted from CP once all its child resources are deleted.
  // Required permissions:
  // - notebook.notebook.delete
  async DeleteNotebook(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/notebook/v1/notebook/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Update an existing notebook. Returns updated Notebook.
  // Required permissions:
  // - notebook.notebook.update
  async UpdateNotebook(req: UpdateNotebookRequest): Promise<Notebook> {
    const url = `/api/notebook/v1/notebook/${encodeURIComponent(req.id || '')}`;
    return api.post(url, req);
  }
  
  // List all notebooks for a deployment.
  // Required permissions:
  // - notebook.notebook.list
  async ListNotebooks(req: ListNotebookRequest): Promise<NotebookList> {
    const url = `/api/notebook/v1/notebooks`;
    return api.post(url, req);
  }
}
