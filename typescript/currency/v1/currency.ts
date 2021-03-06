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

// File: currency/v1/currency.proto
// Package: arangodb.cloud.currency.v1

// Currency represents a specific monetary currency.
export interface Currency {
  // System identifier of the currency.
  // E.g. "eur" or "usd"
  // string
  id?: string;
  
  // Human readable name of the currency
  // E.g. "US Dollar"
  // string
  name?: string;
  
  // Human readable sign for the currency.
  // E.g. "$"
  // string
  sign?: string;
  
  // ISO 4217 currency code.
  // E.g. "USD"
  // string
  iso4217_code?: string;
}

// List of currencies.
export interface CurrencyList {
  // Currency
  items?: Currency[];
}

// Request arguments for GetDefaultCurrency.
export interface GetDefaultCurrencyRequest {
  // Optional identifier for the organization to request the default
  // currency for.
  // string
  organization_id?: string;
}

// CurrencyService is the API used to query for supported currencies.
export interface ICurrencyService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  GetAPIVersion: (req?: arangodb_cloud_common_v1_Empty) => Promise<arangodb_cloud_common_v1_Version>;
  
  // Fetch all providers that are supported by the ArangoDB cloud.
  // Required permissions:
  // - None
  ListCurrencies: (req: arangodb_cloud_common_v1_ListOptions) => Promise<CurrencyList>;
  
  // Fetch a currency by its id.
  // Required permissions:
  // - None
  GetCurrency: (req: arangodb_cloud_common_v1_IDOptions) => Promise<Currency>;
  
  // Fetch the default currency for a given (optional) organization.
  // Required permissions:
  // - resourcemanager.organization.get On the organization identified by given id.
  // - None In case no organization identifier was given.
  GetDefaultCurrency: (req: GetDefaultCurrencyRequest) => Promise<Currency>;
}

// CurrencyService is the API used to query for supported currencies.
export class CurrencyService implements ICurrencyService {
  // Get the current API version of this service.
  // Required permissions:
  // - None
  async GetAPIVersion(req?: arangodb_cloud_common_v1_Empty): Promise<arangodb_cloud_common_v1_Version> {
    const path = `/api/currency/v1/api-version`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch all providers that are supported by the ArangoDB cloud.
  // Required permissions:
  // - None
  async ListCurrencies(req: arangodb_cloud_common_v1_ListOptions): Promise<CurrencyList> {
    const path = `/api/currency/v1/currencies`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
  
  // Fetch a currency by its id.
  // Required permissions:
  // - None
  async GetCurrency(req: arangodb_cloud_common_v1_IDOptions): Promise<Currency> {
    const path = `/api/currency/v1/currencies/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch the default currency for a given (optional) organization.
  // Required permissions:
  // - resourcemanager.organization.get On the organization identified by given id.
  // - None In case no organization identifier was given.
  async GetDefaultCurrency(req: GetDefaultCurrencyRequest): Promise<Currency> {
    const path = `/api/currency/v1/default-currency`;
    const url = path + api.queryString(req, []);
    return api.get(url, undefined);
  }
}
