//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'

// File: credits/v1/credits.proto
// Package: arangodb.cloud.credits.v1
export interface CreditBundle {
  // ID of the credit bundle.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this credit bundle.
  // This is a read-only value.
  // string
  url?: string;
  
  // The organization this credit bundle belongs to.
  // This is a read-only value.
  // string
  organization_id?: string;
  
  // The number of credits purchased in this bundle.
  // This is a read-only value.
  // number
  credits_purchased?: number;
  
  // The total price of these credits.
  // This is a read-only value.
  // number
  total_price?: number;
  
  // Currency used for total_price.
  // This is a read-only value.
  // string
  currency?: string;
  
  // The number of credits remaining in this bundle.
  // This is a read-only value.
  // number
  credits_remaining?: number;
  
  // The date at which this bundle was purchased.
  // This is a read-only value.
  // googleTypes.Timestamp
  purchased_at?: googleTypes.Timestamp;
  
  // The date from which this bundle is valid.
  // This is a read-only value.
  // googleTypes.Timestamp
  valid_from?: googleTypes.Timestamp;
  
  // The date until which this bundle is valid.
  // This is a read-only value.
  // googleTypes.Timestamp
  valid_until?: googleTypes.Timestamp;
  
  // Timestamp of when this credit bundle was last used.
  // googleTypes.Timestamp
  last_used_at?: googleTypes.Timestamp;
}

// Usage of credit bundle.
export interface CreditBundleUsage {
  // Unique identifier of this credit bundle usage.
  // string
  id?: string;
  
  // ID of the usage item this credit bundle usage corresponds to.
  // string
  usage_item_id?: string;
  
  // ID of the credit bundle from which credit was used.
  // string
  credit_bundle_id?: string;
  
  // ID of the organization this credit bundle (and usage) belongs to.
  // string
  organization_id?: string;
  
  // Amount of credits used from the specified credit_bundle_id.
  // number
  usage?: number;
  
  // Amount of credits remaining after this usage.
  // number
  remaining?: number;
  
  // Timestamp at which the credits were used.
  // googleTypes.Timestamp
  used_at?: googleTypes.Timestamp;
}
export interface CreditBundleUsageList {
  // CreditBundleUsage
  items?: CreditBundleUsage[];
}

// List of credit bundles
export interface CreditBundlesList {
  // CreditBundle
  items?: CreditBundle[];
}

// Request for listing credit usage.
export interface ListCreditBundleUsageRequest {
  // The organization this credit bundle belongs to.
  // This is a required field.
  // string
  organization_id?: string;
  
  // If set, list the usage for the specified credit bundle only.
  // By default, returns the usage for all credit bundles in this organization.
  // This is an optional field.
  // string
  credit_bundle_id?: string;
  
  // The date from which credit usage should be listed.
  // If unspecified, defaults to the date 7 days before `ends_at`.
  // This is an optional field.
  // googleTypes.Timestamp
  starts_at?: googleTypes.Timestamp;
  
  // The date until which credit usage should be listed.
  // If unspecified, defaults to the current date (at the time of calling the API).
  // This is an optional field.
  // googleTypes.Timestamp
  end_at?: googleTypes.Timestamp;
  
  // Common list options
  // context_id is ignored.
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Request for listing credit bundles
export interface ListCreditBundlesRequest {
  // ID of the organization for which credit bundles are listed.
  // This is a required field.
  // string
  organization_id?: string;
  
  // If set, exclude expired bundles.
  // boolean
  exclude_expired?: boolean;
}

// CreditsService is the API used for managing credits.
export interface ICreditsService {
  // List credit bundles for an organization.
  // Required permissions:
  // - credit.creditbundle.list on the organization identified by the given organization ID
  ListCreditBundles: (req: ListCreditBundlesRequest) => Promise<CreditBundlesList>;
  
  // List credit usage for the specified organization ID.
  // Required permisisons:
  // - credit.creditbundleusage.list on the organization identified by the given organization ID.
  ListCreditBundlesUsage: (req: ListCreditBundleUsageRequest) => Promise<CreditBundleUsageList>;
}

// CreditsService is the API used for managing credits.
export class CreditsService implements ICreditsService {
  // List credit bundles for an organization.
  // Required permissions:
  // - credit.creditbundle.list on the organization identified by the given organization ID
  async ListCreditBundles(req: ListCreditBundlesRequest): Promise<CreditBundlesList> {
    const path = `/api/credit/v1/${encodeURIComponent(req.organization_id || '')}/creditbundles`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
  
  // List credit usage for the specified organization ID.
  // Required permisisons:
  // - credit.creditbundleusage.list on the organization identified by the given organization ID.
  async ListCreditBundlesUsage(req: ListCreditBundleUsageRequest): Promise<CreditBundleUsageList> {
    const path = `/api/credit/v1/${encodeURIComponent(req.organization_id || '')}/creditbundleusages`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
}
