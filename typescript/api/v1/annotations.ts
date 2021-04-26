//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'

// File: api/v1/annotations.proto
// Package: arangodb.cloud.api.v1

// MethodProfile specifies profile options for methods.
export interface MethodProfile {
  // Maximum timeout in seconds.
  // number
  timeout?: number;
  
  // If set, the method can be retried automatically.
  // boolean
  retryable?: boolean;
}

// Extension .google.protobuf.MethodOptions.profile: Not implemented
