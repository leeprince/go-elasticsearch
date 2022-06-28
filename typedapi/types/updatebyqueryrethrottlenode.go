// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.


// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification


package types

// UpdateByQueryRethrottleNode type.
type UpdateByQueryRethrottleNode struct {
	Attributes       map[string]string   `json:"attributes"`
	Host             Host                `json:"host"`
	Ip               Ip                  `json:"ip"`
	Name             Name                `json:"name"`
	Roles            *NodeRoles          `json:"roles,omitempty"`
	Tasks            map[TaskId]TaskInfo `json:"tasks"`
	TransportAddress TransportAddress    `json:"transport_address"`
}

// UpdateByQueryRethrottleNodeBuilder holds UpdateByQueryRethrottleNode struct and provides a builder API.
type UpdateByQueryRethrottleNodeBuilder struct {
	v *UpdateByQueryRethrottleNode
}

// NewUpdateByQueryRethrottleNode provides a builder for the UpdateByQueryRethrottleNode struct.
func NewUpdateByQueryRethrottleNode() *UpdateByQueryRethrottleNodeBuilder {
	r := UpdateByQueryRethrottleNodeBuilder{
		&UpdateByQueryRethrottleNode{
			Attributes: make(map[string]string, 0),
			Tasks:      make(map[TaskId]TaskInfo, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the UpdateByQueryRethrottleNode struct
func (rb *UpdateByQueryRethrottleNodeBuilder) Build() UpdateByQueryRethrottleNode {
	return *rb.v
}

// Attributes set the Attributes property for UpdateByQueryRethrottleNodeBuilder.
func (rb *UpdateByQueryRethrottleNodeBuilder) Attributes(value map[string]string) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Attributes = value
	return rb
}

// Host set the Host property for UpdateByQueryRethrottleNodeBuilder.
func (rb *UpdateByQueryRethrottleNodeBuilder) Host(host Host) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Host = host
	return rb
}

// Ip set the Ip property for UpdateByQueryRethrottleNodeBuilder.
func (rb *UpdateByQueryRethrottleNodeBuilder) Ip(ip Ip) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Ip = ip
	return rb
}

// Name set the Name property for UpdateByQueryRethrottleNodeBuilder.
func (rb *UpdateByQueryRethrottleNodeBuilder) Name(name Name) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Name = name
	return rb
}

// Roles set the Roles property for UpdateByQueryRethrottleNodeBuilder.
func (rb *UpdateByQueryRethrottleNodeBuilder) Roles(roles NodeRoles) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Roles = &roles
	return rb
}

// Tasks set the Tasks property for UpdateByQueryRethrottleNodeBuilder.
func (rb *UpdateByQueryRethrottleNodeBuilder) Tasks(value map[TaskId]TaskInfo) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.Tasks = value
	return rb
}

// TransportAddress set the TransportAddress property for UpdateByQueryRethrottleNodeBuilder.
func (rb *UpdateByQueryRethrottleNodeBuilder) TransportAddress(transportaddress TransportAddress) *UpdateByQueryRethrottleNodeBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
