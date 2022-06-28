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

// NodeInfoScript type.
type NodeInfoScript struct {
	AllowedTypes               string `json:"allowed_types"`
	DisableMaxCompilationsRate string `json:"disable_max_compilations_rate"`
}

// NodeInfoScriptBuilder holds NodeInfoScript struct and provides a builder API.
type NodeInfoScriptBuilder struct {
	v *NodeInfoScript
}

// NewNodeInfoScript provides a builder for the NodeInfoScript struct.
func NewNodeInfoScript() *NodeInfoScriptBuilder {
	r := NodeInfoScriptBuilder{
		&NodeInfoScript{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoScript struct
func (rb *NodeInfoScriptBuilder) Build() NodeInfoScript {
	return *rb.v
}

// AllowedTypes set the AllowedTypes property for NodeInfoScriptBuilder.
func (rb *NodeInfoScriptBuilder) AllowedTypes(allowedtypes string) *NodeInfoScriptBuilder {
	rb.v.AllowedTypes = allowedtypes
	return rb
}

// DisableMaxCompilationsRate set the DisableMaxCompilationsRate property for NodeInfoScriptBuilder.
func (rb *NodeInfoScriptBuilder) DisableMaxCompilationsRate(disablemaxcompilationsrate string) *NodeInfoScriptBuilder {
	rb.v.DisableMaxCompilationsRate = disablemaxcompilationsrate
	return rb
}
