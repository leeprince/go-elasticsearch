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

// NodePackagingType type.
type NodePackagingType struct {
	Count  int    `json:"count"`
	Flavor string `json:"flavor"`
	Type   string `json:"type"`
}

// NodePackagingTypeBuilder holds NodePackagingType struct and provides a builder API.
type NodePackagingTypeBuilder struct {
	v *NodePackagingType
}

// NewNodePackagingType provides a builder for the NodePackagingType struct.
func NewNodePackagingType() *NodePackagingTypeBuilder {
	r := NodePackagingTypeBuilder{
		&NodePackagingType{},
	}

	return &r
}

// Build finalize the chain and returns the NodePackagingType struct
func (rb *NodePackagingTypeBuilder) Build() NodePackagingType {
	return *rb.v
}

// Count set the Count property for NodePackagingTypeBuilder.
func (rb *NodePackagingTypeBuilder) Count(count int) *NodePackagingTypeBuilder {
	rb.v.Count = count
	return rb
}

// Flavor set the Flavor property for NodePackagingTypeBuilder.
func (rb *NodePackagingTypeBuilder) Flavor(flavor string) *NodePackagingTypeBuilder {
	rb.v.Flavor = flavor
	return rb
}

// Type set the Type property for NodePackagingTypeBuilder.
func (rb *NodePackagingTypeBuilder) Type_(type_ string) *NodePackagingTypeBuilder {
	rb.v.Type = type_
	return rb
}
