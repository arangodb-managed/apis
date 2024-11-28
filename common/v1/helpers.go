//
// DISCLAIMER
//
// Copyright 2020-2024 ArangoDB GmbH, Cologne, Germany
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

package v1

import (
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// DefaultPageSize is the default number of items per List request.
	DefaultPageSize = 50
)

// CloneTimestamp creates a deep clone of the given timestamp
func CloneTimestamp(s *timestamppb.Timestamp) *timestamppb.Timestamp {
	if s == nil {
		return nil
	}
	clone := *s
	return &clone
}

// CloneDuration creates a deep copy of the given duration
func CloneDuration(s *durationpb.Duration) *durationpb.Duration {
	if s == nil {
		return nil
	}
	clone := *s
	return &clone
}

// CloneOrDefault creates a clone of the given options if not nil,
// or creates an empty ListOptions.
// In either case, if current pageSize is zero, it is set to the given
// default page size or DefaultPageSize is not default page size is given.
func (opts *ListOptions) CloneOrDefault(defaultPageSize ...int32) *ListOptions {
	if opts == nil {
		opts = &ListOptions{}
	} else {
		clone := *opts
		opts = &clone
	}
	if opts.PageSize == 0 {
		if len(defaultPageSize) > 0 {
			opts.PageSize = defaultPageSize[0]
		} else {
			opts.PageSize = DefaultPageSize
		}
	}
	return opts
}

// TimestampsEqual compares two *timestamppb.Timestamp instances.
func TimestampsEqual(ts1, ts2 *timestamppb.Timestamp) bool {
	if ts1 == nil && ts2 == nil {
		return true
	}
	if ts1 == nil || ts2 == nil {
		return false
	}
	return ts1.Seconds == ts2.Seconds && ts1.Nanos == ts2.Nanos
}

// CompareTimestamps compares two *timestamppb.Timestamp instances.
// Returns -1 if ts1 < ts2, 1 if ts1 > ts2, and 0 if they are equal.
func CompareTimestamps(ts1, ts2 *timestamppb.Timestamp) int {
	if ts1 == nil && ts2 == nil {
		return 0
	}
	if ts1 == nil {
		return -1
	}
	if ts2 == nil {
		return 1
	}

	if ts1.Seconds < ts2.Seconds {
		return -1
	}
	if ts1.Seconds > ts2.Seconds {
		return 1
	}
	// Seconds are equal; compare nanoseconds.
	if ts1.Nanos < ts2.Nanos {
		return -1
	}
	if ts1.Nanos > ts2.Nanos {
		return 1
	}
	return 0
}

// DurationsEqual compares two *durationpb.Duration instances for equality.
func DurationsEqual(d1, d2 *durationpb.Duration) bool {
	if d1 == nil && d2 == nil {
		return true
	}
	if d1 == nil || d2 == nil {
		return false
	}
	return d1.Seconds == d2.Seconds && d1.Nanos == d2.Nanos
}
