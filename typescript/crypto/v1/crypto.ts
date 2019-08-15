//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'

// File: crypto/v1/crypto.proto
// Package: arangodb.cloud.crypto.v1

// A CACertificate is represents a self-signed certificate authority used to sign
// TLS certificates for deployments & client authentication.
export interface CACertificate {
  // System identifier of the CA certificate.
  // This is a read-only value.
  // string
  id?: string;
  
  // URL of this resource
  // This is a read-only value.
  // string
  url?: string;
  
  // Name of the CA certificate
  // string
  name?: string;
  
  // Description of the CA certificate
  // string
  description?: string;
  
  // Identifier of the project that owns this CA certificate.
  // This value cannot be changed after creation.
  // string
  project_id?: string;
  
  // Time from creation of the CA certificate to expiration.
  // This value cannot be changed after creation.
  // googleTypes.Duration
  lifetime?: googleTypes.Duration;
  
  // The creation timestamp of the CA certificate
  // This is a read-only value.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // The deletion timestamp of the CA certificate
  // This is a read-only value.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // The expiration timestamp of the CA certificate
  // This is a read-only value.
  // googleTypes.Timestamp
  expires_at?: googleTypes.Timestamp;
  
  // A PEM encoded representation of the public key of the CA certificate.
  // This is a read-only value.
  // string
  certificate_pem?: string;
  
  // Set when this CA certificate is deleted.
  // This is a read-only value.
  // boolean
  is_deleted?: boolean;
  
  // Set when this CA certificate has expired.
  // This is a read-only value.
  // boolean
  is_expired?: boolean;
  
  // Set when this CA certificate will expire in the next month.
  // This is a read-only value.
  // boolean
  will_expire_soon?: boolean;
}

// Instructions for installing & uninstalling CA certificates
export interface CACertificateInstructions {
  // Per platform instructions for install/uninstall of the CA certificate
  // CACertificateInstructions_PlatformInstructions
  platforms?: CACertificateInstructions_PlatformInstructions[];
}

// Instructions for a specific platform
export interface CACertificateInstructions_PlatformInstructions {
  // Human readable description of platform.
  // E.g. "MacOS"
  // string
  platform?: string;
  
  // Steps needed to install
  // string
  install_steps?: string[];
  
  // Steps needed to uninstall
  // string
  uninstall_steps?: string[];
}

// List of CACertificates.
export interface CACertificateList {
  // CACertificate
  items?: CACertificate[];
}

// CryptoService is the API used to configure various crypto objects.
export class CryptoService {
  // Fetch all CA certificates in the project identified by the given context ID.
  // Required permissions:
  // - crypto.cacertificate.list on the project identified by the given context ID
  async ListCACertificates(req: arangodb_cloud_common_v1_ListOptions): Promise<CACertificateList> {
    const path = `/api/crypto/v1/projects/${encodeURIComponent(req.context_id || '')}/cacertificates`;
    const url = path + api.queryString(req, [`context_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a CA certificate by its id.
  // Required permissions:
  // - crypto.cacertificate.get on the CA certificate identified by the given ID
  async GetCACertificate(req: arangodb_cloud_common_v1_IDOptions): Promise<CACertificate> {
    const path = `/api/crypto/v1/cacertificates/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch instructions for installing & unistalling a CA certificate identified by its id
  // on various platforms.
  // Required permissions:
  // - crypto.cacertificate.get on the CA certificate identified by the given ID
  async GetCACertificateInstructions(req: arangodb_cloud_common_v1_IDOptions): Promise<CACertificateInstructions> {
    const path = `/api/crypto/v1/cacertificates/${encodeURIComponent(req.id || '')}/instructions`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Create a new CA certificate
  // Required permissions:
  // - crypto.cacertificate.create on the project that owns the CA certificate
  async CreateCACertificate(req: CACertificate): Promise<CACertificate> {
    const url = `/api/crypto/v1/projects/${encodeURIComponent(req.project_id || '')}/cacertificates`;
    return api.post(url, req);
  }
  
  // Update a CA certificate
  // Required permissions:
  // - crypto.cacertificate.update on the CA certificate
  async UpdateCACertificate(req: CACertificate): Promise<CACertificate> {
    const url = `/api/crypto/v1/cacertificates/${encodeURIComponent(req.id || '')}`;
    return api.patch(url, req);
  }
  
  // Delete a CA certificate
  // Note that CA certificate are initially only marked for deleted.
  // Once all the resources that depend on it are removed the CA certificate itself is deleted
  // and cannot be restored.
  // Required permissions:
  // - crypto.cacertificate.delete on the CA certificate
  async DeleteCACertificate(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/crypto/v1/cacertificates/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
}
