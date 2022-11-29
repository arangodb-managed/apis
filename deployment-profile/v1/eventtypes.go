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

const (
	// Deployment Profile event types

	// EventTypeDeploymentProfileCreated is the type of event fired after a Deployment Profile has been created.
	// SubjectID contains the ID of the Deployment Profile.
	EventTypeDeploymentProfileCreated = "deploymentprofile.deploymentprofile.created"
	// EventTypeDeploymentProfileUpdated is the type of event fired after a Deployment Profile has been updated.
	// SubjectID contains the ID of the Deployment Profile.
	EventTypeDeploymentProfileUpdated = "deploymentprofile.deploymentprofile.updated"
	// EventTypeDeploymentProfileDeleted is the type of event fired after a Deployment Profile has been deleted.
	// SubjectID contains the ID of the Deployment Profile.
	EventTypeDeploymentProfileDeleted = "deploymentprofile.deploymentprofile.deleted"
)
