//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { Empty as arangodb_cloud_common_v1_Empty } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'
import { Version as arangodb_cloud_common_v1_Version } from '../../common/v1/common'

// File: notifications/v1/notifications.proto
// Package: arangodb.cloud.prepaid.v1

// ListDeploymentNotificationsRequest is used to request a list of Notifications for given deployment
export interface ListDeploymentNotificationsRequest {
  // identifier of the deployment_id to get a list of notifications for
  // string
  deployment_id?: string;
  
  // common listing options
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Notification contains all attributes of a notification
export interface Notification {
  // System identifier of the notification
  // This is a read-only value.
  // string
  id?: string;
  
  // Type of notification.
  // This is a read-only value.
  // string
  type?: string;
  
  // The creation timestamp of the prepaid deployment
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // Title of notification.
  // This is a read-only value.
  // string
  title?: string;
  
  // Receipments of notification.
  // This is a read-only value.
  // string
  receipment?: string[];
  
  // Content of notification
  // NotificationContent
  content?: NotificationContent[];
}

// NotificationContent holds content and it's mime type
export interface NotificationContent {
  // MIME Type of notification.
  // This is a read-only value.
  // string
  content_type?: string;
  
  // Content of notification
  // This is a read-only value.
  // string
  content?: string;
}

// NotificationList contains a list of Notification items
export interface NotificationList {
  // notification items
  // Notification
  items?: Notification[];
}

// NotificationsServise is the API used to interact with deployment notifications
export interface INotificationsService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Fetch all notifications related to given deployment
  // Required permissions:
  // - notifications.deployment-notifications.list on the deployment identified by given deployment_id
  ListDeploymentNotifications: (req: ListDeploymentNotificationsRequest) => Promise<NotificationList>;
}

// NotificationsServise is the API used to interact with deployment notifications
export class NotificationsService implements INotificationsService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/notifications/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all notifications related to given deployment
  // Required permissions:
  // - notifications.deployment-notifications.list on the deployment identified by given deployment_id
  async ListDeploymentNotifications(req: ListDeploymentNotificationsRequest): Promise<NotificationList> {
    const path = `/api/notifications/v1/${encodeURIComponent(req.deployment_id || '')}`;
    const url = path + api.queryString(req, [`deployment_id`]);
    return api.post(url, undefined);
  }
}
