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

// NodeInfoTransport type.
type NodeInfoTransport struct {
	BoundAddress   []string          `json:"bound_address"`
	Profiles       map[string]string `json:"profiles"`
	PublishAddress string            `json:"publish_address"`
}

// NodeInfoTransportBuilder holds NodeInfoTransport struct and provides a builder API.
type NodeInfoTransportBuilder struct {
	v *NodeInfoTransport
}

// NewNodeInfoTransport provides a builder for the NodeInfoTransport struct.
func NewNodeInfoTransport() *NodeInfoTransportBuilder {
	r := NodeInfoTransportBuilder{
		&NodeInfoTransport{
			Profiles: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoTransport struct
func (rb *NodeInfoTransportBuilder) Build() NodeInfoTransport {
	return *rb.v
}

// BoundAddress set the BoundAddress property for NodeInfoTransportBuilder.
func (rb *NodeInfoTransportBuilder) BoundAddress(bound_address ...string) *NodeInfoTransportBuilder {
	rb.v.BoundAddress = append(rb.v.BoundAddress, bound_address...)
	return rb
}

// Profiles set the Profiles property for NodeInfoTransportBuilder.
func (rb *NodeInfoTransportBuilder) Profiles(value map[string]string) *NodeInfoTransportBuilder {
	rb.v.Profiles = value
	return rb
}

// PublishAddress set the PublishAddress property for NodeInfoTransportBuilder.
func (rb *NodeInfoTransportBuilder) PublishAddress(publishaddress string) *NodeInfoTransportBuilder {
	rb.v.PublishAddress = publishaddress
	return rb
}
