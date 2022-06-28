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

// NodeInfoNetwork type.
type NodeInfoNetwork struct {
	PrimaryInterface NodeInfoNetworkInterface `json:"primary_interface"`
	RefreshInterval  int                      `json:"refresh_interval"`
}

// NodeInfoNetworkBuilder holds NodeInfoNetwork struct and provides a builder API.
type NodeInfoNetworkBuilder struct {
	v *NodeInfoNetwork
}

// NewNodeInfoNetwork provides a builder for the NodeInfoNetwork struct.
func NewNodeInfoNetwork() *NodeInfoNetworkBuilder {
	r := NodeInfoNetworkBuilder{
		&NodeInfoNetwork{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoNetwork struct
func (rb *NodeInfoNetworkBuilder) Build() NodeInfoNetwork {
	return *rb.v
}

// PrimaryInterface set the PrimaryInterface property for NodeInfoNetworkBuilder.
func (rb *NodeInfoNetworkBuilder) PrimaryInterface(primaryinterface NodeInfoNetworkInterface) *NodeInfoNetworkBuilder {
	rb.v.PrimaryInterface = primaryinterface
	return rb
}

// RefreshInterval set the RefreshInterval property for NodeInfoNetworkBuilder.
func (rb *NodeInfoNetworkBuilder) RefreshInterval(refreshinterval int) *NodeInfoNetworkBuilder {
	rb.v.RefreshInterval = refreshinterval
	return rb
}
