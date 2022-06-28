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

// NodeInfoSettingsNode type.
type NodeInfoSettingsNode struct {
	Attr                 map[string]interface{} `json:"attr"`
	MaxLocalStorageNodes *string                `json:"max_local_storage_nodes,omitempty"`
	Name                 Name                   `json:"name"`
}

// NodeInfoSettingsNodeBuilder holds NodeInfoSettingsNode struct and provides a builder API.
type NodeInfoSettingsNodeBuilder struct {
	v *NodeInfoSettingsNode
}

// NewNodeInfoSettingsNode provides a builder for the NodeInfoSettingsNode struct.
func NewNodeInfoSettingsNode() *NodeInfoSettingsNodeBuilder {
	r := NodeInfoSettingsNodeBuilder{
		&NodeInfoSettingsNode{
			Attr: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsNode struct
func (rb *NodeInfoSettingsNodeBuilder) Build() NodeInfoSettingsNode {
	return *rb.v
}

// Attr set the Attr property for NodeInfoSettingsNodeBuilder.
func (rb *NodeInfoSettingsNodeBuilder) Attr(value map[string]interface{}) *NodeInfoSettingsNodeBuilder {
	rb.v.Attr = value
	return rb
}

// MaxLocalStorageNodes set the MaxLocalStorageNodes property for NodeInfoSettingsNodeBuilder.
func (rb *NodeInfoSettingsNodeBuilder) MaxLocalStorageNodes(maxlocalstoragenodes string) *NodeInfoSettingsNodeBuilder {
	rb.v.MaxLocalStorageNodes = &maxlocalstoragenodes
	return rb
}

// Name set the Name property for NodeInfoSettingsNodeBuilder.
func (rb *NodeInfoSettingsNodeBuilder) Name(name Name) *NodeInfoSettingsNodeBuilder {
	rb.v.Name = name
	return rb
}
