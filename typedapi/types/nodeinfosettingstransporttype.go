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

// NodeInfoSettingsTransportType type.
type NodeInfoSettingsTransportType struct {
	Default string `json:"default"`
}

// NodeInfoSettingsTransportTypeBuilder holds NodeInfoSettingsTransportType struct and provides a builder API.
type NodeInfoSettingsTransportTypeBuilder struct {
	v *NodeInfoSettingsTransportType
}

// NewNodeInfoSettingsTransportType provides a builder for the NodeInfoSettingsTransportType struct.
func NewNodeInfoSettingsTransportType() *NodeInfoSettingsTransportTypeBuilder {
	r := NodeInfoSettingsTransportTypeBuilder{
		&NodeInfoSettingsTransportType{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsTransportType struct
func (rb *NodeInfoSettingsTransportTypeBuilder) Build() NodeInfoSettingsTransportType {
	return *rb.v
}

// Default set the Default property for NodeInfoSettingsTransportTypeBuilder.
func (rb *NodeInfoSettingsTransportTypeBuilder) Default_(default_ string) *NodeInfoSettingsTransportTypeBuilder {
	rb.v.Default = default_
	return rb
}
