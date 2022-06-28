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

// NodeInfoXpackSecurityAuthcRealmsStatus type.
type NodeInfoXpackSecurityAuthcRealmsStatus struct {
	Enabled *string `json:"enabled,omitempty"`
	Order   string  `json:"order"`
}

// NodeInfoXpackSecurityAuthcRealmsStatusBuilder holds NodeInfoXpackSecurityAuthcRealmsStatus struct and provides a builder API.
type NodeInfoXpackSecurityAuthcRealmsStatusBuilder struct {
	v *NodeInfoXpackSecurityAuthcRealmsStatus
}

// NewNodeInfoXpackSecurityAuthcRealmsStatus provides a builder for the NodeInfoXpackSecurityAuthcRealmsStatus struct.
func NewNodeInfoXpackSecurityAuthcRealmsStatus() *NodeInfoXpackSecurityAuthcRealmsStatusBuilder {
	r := NodeInfoXpackSecurityAuthcRealmsStatusBuilder{
		&NodeInfoXpackSecurityAuthcRealmsStatus{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoXpackSecurityAuthcRealmsStatus struct
func (rb *NodeInfoXpackSecurityAuthcRealmsStatusBuilder) Build() NodeInfoXpackSecurityAuthcRealmsStatus {
	return *rb.v
}

// Enabled set the Enabled property for NodeInfoXpackSecurityAuthcRealmsStatusBuilder.
func (rb *NodeInfoXpackSecurityAuthcRealmsStatusBuilder) Enabled(enabled string) *NodeInfoXpackSecurityAuthcRealmsStatusBuilder {
	rb.v.Enabled = &enabled
	return rb
}

// Order set the Order property for NodeInfoXpackSecurityAuthcRealmsStatusBuilder.
func (rb *NodeInfoXpackSecurityAuthcRealmsStatusBuilder) Order(order string) *NodeInfoXpackSecurityAuthcRealmsStatusBuilder {
	rb.v.Order = order
	return rb
}
