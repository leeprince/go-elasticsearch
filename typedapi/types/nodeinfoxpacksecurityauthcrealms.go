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

// NodeInfoXpackSecurityAuthcRealms type.
type NodeInfoXpackSecurityAuthcRealms struct {
	File   map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"file,omitempty"`
	Native map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"native,omitempty"`
	Pki    map[string]NodeInfoXpackSecurityAuthcRealmsStatus `json:"pki,omitempty"`
}

// NodeInfoXpackSecurityAuthcRealmsBuilder holds NodeInfoXpackSecurityAuthcRealms struct and provides a builder API.
type NodeInfoXpackSecurityAuthcRealmsBuilder struct {
	v *NodeInfoXpackSecurityAuthcRealms
}

// NewNodeInfoXpackSecurityAuthcRealms provides a builder for the NodeInfoXpackSecurityAuthcRealms struct.
func NewNodeInfoXpackSecurityAuthcRealms() *NodeInfoXpackSecurityAuthcRealmsBuilder {
	r := NodeInfoXpackSecurityAuthcRealmsBuilder{
		&NodeInfoXpackSecurityAuthcRealms{
			File:   make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, 0),
			Native: make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, 0),
			Pki:    make(map[string]NodeInfoXpackSecurityAuthcRealmsStatus, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoXpackSecurityAuthcRealms struct
func (rb *NodeInfoXpackSecurityAuthcRealmsBuilder) Build() NodeInfoXpackSecurityAuthcRealms {
	return *rb.v
}

// File set the File property for NodeInfoXpackSecurityAuthcRealmsBuilder.
func (rb *NodeInfoXpackSecurityAuthcRealmsBuilder) File(value map[string]NodeInfoXpackSecurityAuthcRealmsStatus) *NodeInfoXpackSecurityAuthcRealmsBuilder {
	rb.v.File = value
	return rb
}

// Native set the Native property for NodeInfoXpackSecurityAuthcRealmsBuilder.
func (rb *NodeInfoXpackSecurityAuthcRealmsBuilder) Native(value map[string]NodeInfoXpackSecurityAuthcRealmsStatus) *NodeInfoXpackSecurityAuthcRealmsBuilder {
	rb.v.Native = value
	return rb
}

// Pki set the Pki property for NodeInfoXpackSecurityAuthcRealmsBuilder.
func (rb *NodeInfoXpackSecurityAuthcRealmsBuilder) Pki(value map[string]NodeInfoXpackSecurityAuthcRealmsStatus) *NodeInfoXpackSecurityAuthcRealmsBuilder {
	rb.v.Pki = value
	return rb
}
