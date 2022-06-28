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

// Hop type.
type Hop struct {
	Connections *Hop               `json:"connections,omitempty"`
	Query       QueryContainer     `json:"query"`
	Vertices    []VertexDefinition `json:"vertices"`
}

// HopBuilder holds Hop struct and provides a builder API.
type HopBuilder struct {
	v *Hop
}

// NewHop provides a builder for the Hop struct.
func NewHop() *HopBuilder {
	r := HopBuilder{
		&Hop{},
	}

	return &r
}

// Build finalize the chain and returns the Hop struct
func (rb *HopBuilder) Build() Hop {
	return *rb.v
}

// Connections set the Connections property for HopBuilder.
func (rb *HopBuilder) Connections(connections Hop) *HopBuilder {
	rb.v.Connections = &connections
	return rb
}

// Query set the Query property for HopBuilder.
func (rb *HopBuilder) Query(query QueryContainer) *HopBuilder {
	rb.v.Query = query
	return rb
}

// Vertices set the Vertices property for HopBuilder.
func (rb *HopBuilder) Vertices(vertices ...VertexDefinition) *HopBuilder {
	rb.v.Vertices = append(rb.v.Vertices, vertices...)
	return rb
}
