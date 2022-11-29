//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
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
	"net/url"
	"path"

	rm "github.com/arangodb-managed/apis/resourcemanager/v1"
)

const (
	// KindDeploymentProfile is a constant for the kind of DeploymentProfile resources.
	KindDeploymentProfile = "DeploymentProfile"
)

// DeploymentProfileURL creates a resource URL for the DeploymentProfile with given ID
// in given context.
func DeploymentProfileURL(organizationID, deploymentProfileID string) string {
	return path.Join(rm.OrganizationURL(organizationID), KindDeploymentProfile, url.PathEscape(deploymentProfileID))
}
