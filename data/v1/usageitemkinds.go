//
// DISCLAIMER
//
// Copyright 2021-2026 ArangoDB GmbH, Cologne, Germany
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
// Author Robert Stam
//

package v1

const (
	// Data usage item resource kinds

	// UsageItemResourceKindDeployment is the kind used inside the UsageItem.Resource to refer to a deployment.
	UsageItemResourceKindDeployment = "Deployment"

	// UsageItemResourceKindDataCluster is the kind used inside the UsageItem.Resource to refer to a data cluster.
	// Used for usage that belongs to a whole data cluster rather than a single deployment
	// (e.g. AddonHour items for dedicated data clusters).
	UsageItemResourceKindDataCluster = "DataCluster"
)
