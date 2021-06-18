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
// Author Marcin Swiderski
//

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepaidDeploymentURL(t *testing.T) {
	assert.Equal(t, "/Organization/org_id/PrepaidDeployment/pd_id", PrepaidDeploymentURL("org_id", "pd_id"))
	assert.Equal(t, "/Organization/org%2Fid/PrepaidDeployment/pd%2Fid", PrepaidDeploymentURL("org/id", "pd/id"))
	assert.Equal(t, "/Organization/123/PrepaidDeployment/456", PrepaidDeploymentURL("123", "456"))
	assert.Equal(t, "/Organization/123/PrepaidDeployment/%3Fx=456", PrepaidDeploymentURL("123", "?x=456"))
}
