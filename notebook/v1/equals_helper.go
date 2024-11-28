//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
//

package v1

import (
	"google.golang.org/protobuf/proto"

	common "github.com/arangodb-managed/apis/common/v1"
)

// Equals returns true when source and other have the same values.
func (source *Status) Equals(other *Status) bool {
	return source.GetEndpoint() == other.GetEndpoint() &&
		source.GetPhase() == other.GetPhase() &&
		source.GetMessage() == other.GetMessage() &&
		common.TimestampsEqual(source.GetLastUpdatedAt(), other.GetLastUpdatedAt()) &&
		common.TimestampsEqual(source.GetLastActiveAt(), other.GetLastActiveAt()) &&
		proto.Equal(source.GetUsage(), other.GetUsage())
}
