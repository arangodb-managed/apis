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
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrganizationURL(t *testing.T) {
	assert.Equal(t, "/Organization/123", OrganizationURL("123"))
	assert.Equal(t, "/Organization/123%2F567", OrganizationURL("123/567"))
}

func TestOrganizationURLFormat(t *testing.T) {
	u, err := url.Parse(OrganizationURL("123"))
	require.NoError(t, err)
	assert.Equal(t, "", u.Host)
	assert.Equal(t, "/Organization/123", u.Path)
}

func TestProjectURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Project/p1", ProjectURL("123", "p1"))
	assert.Equal(t, "/Organization/123%2F567/Project/ab", ProjectURL("123/567", "ab"))
	assert.Equal(t, "/Organization/123%2F567/Project/a%25b", ProjectURL("123/567", "a%b"))
}

func TestOrganizationInviteURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/OrganizationInvite/p1", OrganizationInviteURL("123", "p1"))
	assert.Equal(t, "/Organization/123%2F567/OrganizationInvite/ab", OrganizationInviteURL("123/567", "ab"))
	assert.Equal(t, "/Organization/123%2F567/OrganizationInvite/a%25b", OrganizationInviteURL("123/567", "a%b"))
}

func TestParseResourceURL(t *testing.T) {
	u, err := ParseResourceURL("/Organization/123/Project/p1")
	require.NoError(t, err)
	assert.Equal(t, "123", u.OrganizationID())
	assert.Equal(t, "p1", u.ProjectID())
	assert.Equal(t, "", u.ProjectChildID("Foo"))
	assert.Equal(t, "/Organization/123/Project/p1", u.String())

	u, err = ParseResourceURL("/Organization/123%2F567/Project/p1")
	require.NoError(t, err)
	assert.Equal(t, "123/567", u.OrganizationID())
	assert.Equal(t, "p1", u.ProjectID())
	assert.Equal(t, "", u.ProjectChildID("Foo"))
	assert.Equal(t, "/Organization/123%2F567/Project/p1", u.String())

	u, err = ParseResourceURL("/Organization/123%2F567/Project/p1/Backup/b52")
	require.NoError(t, err)
	assert.Equal(t, "123/567", u.OrganizationID())
	assert.Equal(t, "p1", u.ProjectID())
	assert.Equal(t, "b52", u.ProjectChildID("Backup"))
	assert.Equal(t, "/Organization/123%2F567/Project/p1/Backup/b52", u.String())

	u, err = ParseResourceURL("/Organization/123%2F567/Project/p1/Foo/b52/Bar/a17")
	require.NoError(t, err)
	assert.Equal(t, "123/567", u.OrganizationID())
	assert.Equal(t, "p1", u.ProjectID())
	assert.Equal(t, "b52", u.ProjectChildID("Foo"))
	assert.Equal(t, "a17", u.ProjectChildID("Foo", "Bar"))
	assert.Equal(t, "/Organization/123%2F567/Project/p1/Foo/b52/Bar/a17", u.String())
}
