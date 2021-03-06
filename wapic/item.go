/*
   Onix Web API Client Library - Copyright (c) 2020 by www.gatblau.org

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
   Unless required by applicable law or agreed to in writing, software distributed under
   the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
   either express or implied.
   See the License for the specific language governing permissions and limitations under the License.

   Contributors to this project, hereby assign copyright in this code to the project,
   to be licensed under the same terms as the rest of the code.
*/
package wapic

import (
	"bytes"
)

type MAP map[string]interface{}

type Item struct {
	Key         string        `json:"key"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Status      int           `json:"status"`
	Type        string        `json:"type"`
	Tag         []interface{} `json:"tag"`
	Meta        MAP           `json:"meta"`
	Attribute   MAP           `json:"attribute"`
	Partition   string        `json:"partition"`
}

func (item *Item) ToJSON() (*bytes.Reader, error) {
	return getJSONBytesReader(item)
}

func (item *Item) KeyValue() string {
	return item.Key
}
