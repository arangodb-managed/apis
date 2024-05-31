//
// DISCLAIMER
//
// Copyright 2024 ArangoDB GmbH, Cologne, Germany
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

func (d *Deployment) Notification(id string) (*Notification, bool) {
	v, ok := d.GetNotifications()[id]
	return v, ok
}

func (d *Deployment) NotificationEquals(id string, notification string, severity NotificationSeverity) bool {
	v, ok := d.GetNotifications()[id]
	if !ok {
		return false
	}

	return v.Equals(notification, severity)
}

func (n *Notification) Equals(notification string, severity NotificationSeverity) bool {
	if n == nil {
		return false
	}

	return n.Severity == severity && n.Notification == notification
}
