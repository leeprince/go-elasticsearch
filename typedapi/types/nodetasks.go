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

// NodeTasks type.
type NodeTasks struct {
	Attributes       map[string]string   `json:"attributes,omitempty"`
	Host             *Host               `json:"host,omitempty"`
	Ip               *Ip                 `json:"ip,omitempty"`
	Name             *NodeId             `json:"name,omitempty"`
	Roles            []string            `json:"roles,omitempty"`
	Tasks            map[TaskId]TaskInfo `json:"tasks"`
	TransportAddress *TransportAddress   `json:"transport_address,omitempty"`
}

// NodeTasksBuilder holds NodeTasks struct and provides a builder API.
type NodeTasksBuilder struct {
	v *NodeTasks
}

// NewNodeTasks provides a builder for the NodeTasks struct.
func NewNodeTasks() *NodeTasksBuilder {
	r := NodeTasksBuilder{
		&NodeTasks{
			Attributes: make(map[string]string, 0),
			Tasks:      make(map[TaskId]TaskInfo, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeTasks struct
func (rb *NodeTasksBuilder) Build() NodeTasks {
	return *rb.v
}

// Attributes set the Attributes property for NodeTasksBuilder.
func (rb *NodeTasksBuilder) Attributes(value map[string]string) *NodeTasksBuilder {
	rb.v.Attributes = value
	return rb
}

// Host set the Host property for NodeTasksBuilder.
func (rb *NodeTasksBuilder) Host(host Host) *NodeTasksBuilder {
	rb.v.Host = &host
	return rb
}

// Ip set the Ip property for NodeTasksBuilder.
func (rb *NodeTasksBuilder) Ip(ip Ip) *NodeTasksBuilder {
	rb.v.Ip = &ip
	return rb
}

// Name set the Name property for NodeTasksBuilder.
func (rb *NodeTasksBuilder) Name(name NodeId) *NodeTasksBuilder {
	rb.v.Name = &name
	return rb
}

// Roles set the Roles property for NodeTasksBuilder.
func (rb *NodeTasksBuilder) Roles(roles ...string) *NodeTasksBuilder {
	rb.v.Roles = append(rb.v.Roles, roles...)
	return rb
}

// Tasks set the Tasks property for NodeTasksBuilder.
func (rb *NodeTasksBuilder) Tasks(value map[TaskId]TaskInfo) *NodeTasksBuilder {
	rb.v.Tasks = value
	return rb
}

// TransportAddress set the TransportAddress property for NodeTasksBuilder.
func (rb *NodeTasksBuilder) TransportAddress(transportaddress TransportAddress) *NodeTasksBuilder {
	rb.v.TransportAddress = &transportaddress
	return rb
}
