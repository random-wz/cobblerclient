/*
Copyright 2015 Container Solutions

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cobblerclient

type KickstartFile struct {
	Name string // The name the kickstart file will be saved in Cobbler
	Body string // The contents of the kickstart file
}

// Creates a kickstart file in Cobbler.
// Returns true/false and returns an optional error in case
// that anything goes wrong.
func (c *Client) CreateKickstartFile(f *KickstartFile) (bool, error) {
	body := tplCreateKickstartFile(f.Name, f.Body, c.token)
	res, err := c.post(body)

	if err != nil {
		return false, err
	}

	return boolFromResponse(res)
}
