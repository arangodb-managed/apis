//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CauseFunc specifies the prototype of a function that must return the cause
// of the given error.
// If there is not underlying cause, the given error itself must be retured.
// If nil is passed, nil must be returned.
type CauseFunc = func(error) error

// Cause is the cause function used by the error helpers in this module.
func Cause(err error) error {
	for err != nil {
		if s, ok := status.FromError(err); ok {
			return s.Err()
		}
		err = errors.Unwrap(err)
	}
	return nil
}

// IsCanceled returns true if the given error signals a request that was canceled. Typically by the caller.
func IsCanceled(err error) bool {
	return status.Code(Cause(err)) == codes.Canceled
}

// Canceled creates a new error that signals a request that was canceled. Typically by the caller.
func Canceled(msg string) error {
	return status.Error(codes.Canceled, msg)
}

// Canceledf creates a new error that signals a request that was canceled. Typically by the caller.
func Canceledf(format string, args ...interface{}) error {
	return status.Errorf(codes.Canceled, format, args...)
}

// IsDeadlineExceeded returns true if the given error signals a request that timed out.
func IsDeadlineExceeded(err error) bool {
	return status.Code(Cause(err)) == codes.DeadlineExceeded
}

// DeadlineExceeded creates a new error that signals a request that timed out.
func DeadlineExceeded(msg string) error {
	return status.Error(codes.DeadlineExceeded, msg)
}

// DeadlineExceededf creates a new error that signals a request that timed out.
func DeadlineExceededf(format string, args ...interface{}) error {
	return status.Errorf(codes.DeadlineExceeded, format, args...)
}

// IsInvalidArgument returns true if the given error signals a request with invalid arguments.
func IsInvalidArgument(err error) bool {
	return status.Code(Cause(err)) == codes.InvalidArgument
}

// InvalidArgument creates a new error that signals a request with invalid arguments.
func InvalidArgument(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}

// InvalidArgumentf creates a new error that signals a request with invalid arguments.
func InvalidArgumentf(format string, args ...interface{}) error {
	return status.Errorf(codes.InvalidArgument, format, args...)
}

// IsNotFound returns true if the given error signals a request to an object that is not found.
func IsNotFound(err error) bool {
	return status.Code(Cause(err)) == codes.NotFound
}

// NotFound creates a new error that signals a request to an object that is not found.
func NotFound(msg string) error {
	return status.Error(codes.NotFound, msg)
}

// NotFoundf creates a new error that signals a request to an object that is not found.
func NotFoundf(format string, args ...interface{}) error {
	return status.Errorf(codes.NotFound, format, args...)
}

// IsAlreadyExists returns true if the given error signals a request to create an object that already exists.
func IsAlreadyExists(err error) bool {
	return status.Code(Cause(err)) == codes.AlreadyExists
}

// AlreadyExists creates a new error that signals a request to create an object that already exists.
func AlreadyExists(msg string) error {
	return status.Error(codes.AlreadyExists, msg)
}

// AlreadyExistsf creates a new error that signals a request to create an object that already exists.
func AlreadyExistsf(format string, args ...interface{}) error {
	return status.Errorf(codes.AlreadyExists, format, args...)
}

// IsPermissionDenied returns true if the given error signals a request that the caller has not enough permissions for.
func IsPermissionDenied(err error) bool {
	return status.Code(Cause(err)) == codes.PermissionDenied
}

// PermissionDenied creates a new error that signals a request that the caller has not enough permissions for.
func PermissionDenied(msg string) error {
	return status.Error(codes.PermissionDenied, msg)
}

// PermissionDeniedf creates a new error that signals a request that the caller has not enough permissions for.
func PermissionDeniedf(format string, args ...interface{}) error {
	return status.Errorf(codes.PermissionDenied, format, args...)
}

// IsPreconditionFailed returns true if the given error signals a precondition of the request has failed.
func IsPreconditionFailed(err error) bool {
	return status.Code(Cause(err)) == codes.FailedPrecondition
}

// PreconditionFailed creates a new error that signals a request that a precondition of the call has failed.
func PreconditionFailed(msg string) error {
	return status.Error(codes.FailedPrecondition, msg)
}

// PreconditionFailedf creates a new error that signals a request that a precondition of the call has failed.
func PreconditionFailedf(format string, args ...interface{}) error {
	return status.Errorf(codes.FailedPrecondition, format, args...)
}

// IsUnauthenticated returns true if the given error signals an unauthenticated request.
func IsUnauthenticated(err error) bool {
	return status.Code(Cause(err)) == codes.Unauthenticated
}

// Unauthenticated creates a new error that signals an unauthenticated request.
func Unauthenticated(msg string) error {
	return status.Error(codes.Unauthenticated, msg)
}

// Unauthenticatedf creates a new error that signals an unauthenticated request.
func Unauthenticatedf(format string, args ...interface{}) error {
	return status.Errorf(codes.Unauthenticated, format, args...)
}

// IsResourceExhausted returns true if the given error signals a request that failed because of lack
// of resources, e.g. user quotas.
func IsResourceExhausted(err error) bool {
	return status.Code(Cause(err)) == codes.ResourceExhausted
}

// ResourceExhausted creates a new error that signals  a request that failed because of lack
// of resources, e.g. user quotas.
func ResourceExhausted(msg string) error {
	return status.Error(codes.ResourceExhausted, msg)
}

// ResourceExhaustedf creates a new error that signals  a request that failed because of lack
// of resources, e.g. user quotas.
func ResourceExhaustedf(format string, args ...interface{}) error {
	return status.Errorf(codes.ResourceExhausted, format, args...)
}

// IsUnknown returns true if the given error signals an unknown error.
func IsUnknown(err error) bool {
	return status.Code(Cause(err)) == codes.Unknown
}

// Unknown creates a new error that signals an unknown error.
func Unknown(msg string) error {
	return status.Error(codes.Unknown, msg)
}

// Unknownf creates a new error that signals an unknown error.
func Unknownf(format string, args ...interface{}) error {
	return status.Errorf(codes.Unknown, format, args...)
}

// IsUnavailable returns true if the given error signals an unavailable error.
// This is a most likely a transient condition and may be corrected
// by retrying with a backoff. Note that it is not always safe to retry
// non-idempotent operations.
func IsUnavailable(err error) bool {
	return status.Code(Cause(err)) == codes.Unavailable
}

// Unavailable creates a new error that signals an unavailable service.
func Unavailable(msg string) error {
	return status.Error(codes.Unavailable, msg)
}

// Unavailablef creates a new error that signals an unavailable service.
func Unavailablef(format string, args ...interface{}) error {
	return status.Errorf(codes.Unavailable, format, args...)
}

// IsAborted returns true if the given error signals that the operation was aborted.
func IsAborted(err error) bool {
	return status.Code(Cause(err)) == codes.Aborted
}

// Aborted creates a new error that signals that an operation was aborted.
func Aborted(msg string) error {
	return status.Error(codes.Aborted, msg)
}

// Abortedf creates a new error that signals that an operation was aborted.
func Abortedf(format string, args ...interface{}) error {
	return status.Errorf(codes.Aborted, format, args...)
}

// CommonError extracts common error from given error and returns it
// If the given err is nil it returns nil also
// If the given err Cause is not any of known common errors it return Unknown error
func CommonError(err error) error {
	if err == nil {
		return nil
	}
	if common := Cause(err); common != nil {
		return common
	}
	return Unknown("Unknown error")
}
