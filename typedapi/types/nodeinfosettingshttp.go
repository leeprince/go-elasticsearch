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

// NodeInfoSettingsHttp type.
type NodeInfoSettingsHttp struct {
	Compression string                   `json:"compression,omitempty"`
	Port        string                   `json:"port,omitempty"`
	Type        NodeInfoSettingsHttpType `json:"type"`
	TypeDefault *string                  `json:"type.default,omitempty"`
}

// NodeInfoSettingsHttpBuilder holds NodeInfoSettingsHttp struct and provides a builder API.
type NodeInfoSettingsHttpBuilder struct {
	v *NodeInfoSettingsHttp
}

// NewNodeInfoSettingsHttp provides a builder for the NodeInfoSettingsHttp struct.
func NewNodeInfoSettingsHttp() *NodeInfoSettingsHttpBuilder {
	r := NodeInfoSettingsHttpBuilder{
		&NodeInfoSettingsHttp{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsHttp struct
func (rb *NodeInfoSettingsHttpBuilder) Build() NodeInfoSettingsHttp {
	return *rb.v
}

// Compression set the Compression property for NodeInfoSettingsHttpBuilder.
func (rb *NodeInfoSettingsHttpBuilder) Compression(arg string) *NodeInfoSettingsHttpBuilder {
	rb.v.Compression = arg
	return rb
}

// Port set the Port property for NodeInfoSettingsHttpBuilder.
func (rb *NodeInfoSettingsHttpBuilder) Port(arg string) *NodeInfoSettingsHttpBuilder {
	rb.v.Port = arg
	return rb
}

// Type set the Type property for NodeInfoSettingsHttpBuilder.
func (rb *NodeInfoSettingsHttpBuilder) Type_(type_ NodeInfoSettingsHttpType) *NodeInfoSettingsHttpBuilder {
	rb.v.Type = type_
	return rb
}

// TypeDefault set the TypeDefault property for NodeInfoSettingsHttpBuilder.
func (rb *NodeInfoSettingsHttpBuilder) TypeDefault(typedefault string) *NodeInfoSettingsHttpBuilder {
	rb.v.TypeDefault = &typedefault
	return rb
}
