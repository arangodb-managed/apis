//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'

// File: billing/v1/billing.proto
// Package: arangodb.cloud.billing.v1

// An Invoice message describes a transaction for usage of ArangoDB Oasis.
export interface Invoice {
  // System identifier of the invoice.
  // string
  id?: string;
  
  // URL of this resource
  // string
  url?: string;
  
  // Identifier of the organization that is responsible for the payment of this invoice.
  // string
  organization_id?: string;
  
  // Name of the organization that is responsible for the payment of this invoice.
  // string
  organization_name?: string;
  
  // Identifier of the legal entity that is the sender of this invoice.
  // string
  entity_id?: string;
  
  // Name of the legal entity that is the sender of this invoice.
  // string
  entity_name?: string;
  
  // The creation date of the invoice
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // All items of the invoice
  // Invoice_Item
  items?: Invoice_Item[];
  
  // Currency for all amounts
  // string
  currency_id?: string;
  
  // Sum all amount for all items
  // number
  total_amount_ex_vat?: number;
  
  // VAT amount for all items
  // number
  total_vat?: number;
  
  // Sum of total_amount_ex_vat + total_vat.
  // This is the amount that the customer will be charged for.
  // number
  total_amount_including_vat?: number;
  
  // Invoice_Status
  status?: Invoice_Status;
}

// A single item of the invoice
export interface Invoice_Item {
  // Identifier of the UsageItem that this item covers.
  // string
  usageitem_id?: string;
  
  // Amount of money (ex VAT) for this item
  // number
  amount?: number;
  
  // Human readable description of this item
  // string
  description?: string;
}

// Status of the invoice
export interface Invoice_Status {
  // If set, this invoice is still being processed.
  // boolean
  is_pending?: boolean;
  
  // If set, this invoice has been payed for succesfully.
  // boolean
  is_completed?: boolean;
  
  // If set, payment for this invoice has been rejected.
  // boolean
  is_rejected?: boolean;
  
  // The timestamp of succesfull completion of the payment.
  // googleTypes.Timestamp
  completed_at?: googleTypes.Timestamp;
  
  // The timestamp of rejected completion of the payment.
  // googleTypes.Timestamp
  rejected_at?: googleTypes.Timestamp;
}

// List of Invoices.
export interface InvoiceList {
  // Invoice
  items?: Invoice[];
}

// Request arguments for ListInvoices
export interface ListInvoicesRequest {
  // Request invoices for the organization with this id.
  // This is a required field.
  // string
  organization_id?: string;
  
  // Request invoices that are created at or after this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  from?: googleTypes.Timestamp;
  
  // Request invoices that are created before this timestamp.
  // This is a required field.
  // googleTypes.Timestamp
  to?: googleTypes.Timestamp;
  
  // Standard list options
  // This is an optional field.
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// BillingService is the API used to fetch billing information.
export class BillingService {
  // Fetch all Invoice resources for the organization identified by the given
  // organization ID that match the given criteria.
  // Required permissions:
  // - billing.invoice.list on the organization identified by the given organization ID
  async ListInvoices(req: ListInvoicesRequest): Promise<InvoiceList> {
    const path = `/api/billing/v1/organization/${encodeURIComponent(req.organization_id || '')}/invoices`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a specific Invoice identified by the given ID.
  // Required permissions:
  // - billing.invoice.get on the organization that owns the invoice
  // with given ID.
  async GetInvoice(req: arangodb_cloud_common_v1_IDOptions): Promise<Invoice> {
    const path = `/api/billing/v1/invoices/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
}
