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

// DocumentSimulation type.
type DocumentSimulation struct {
	Id_      Id                     `json:"_id"`
	Index_   IndexName              `json:"_index"`
	Ingest_  Ingest                 `json:"_ingest"`
	Parent_  *string                `json:"_parent,omitempty"`
	Routing_ *string                `json:"_routing,omitempty"`
	Source_  map[string]interface{} `json:"_source"`
}

// DocumentSimulationBuilder holds DocumentSimulation struct and provides a builder API.
type DocumentSimulationBuilder struct {
	v *DocumentSimulation
}

// NewDocumentSimulation provides a builder for the DocumentSimulation struct.
func NewDocumentSimulation() *DocumentSimulationBuilder {
	r := DocumentSimulationBuilder{
		&DocumentSimulation{
			Source_: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DocumentSimulation struct
func (rb *DocumentSimulationBuilder) Build() DocumentSimulation {
	return *rb.v
}

// Id_ set the Id_ property for DocumentSimulationBuilder.
func (rb *DocumentSimulationBuilder) Id_(id_ Id) *DocumentSimulationBuilder {
	rb.v.Id_ = id_
	return rb
}

// Index_ set the Index_ property for DocumentSimulationBuilder.
func (rb *DocumentSimulationBuilder) Index_(index_ IndexName) *DocumentSimulationBuilder {
	rb.v.Index_ = index_
	return rb
}

// Ingest_ set the Ingest_ property for DocumentSimulationBuilder.
func (rb *DocumentSimulationBuilder) Ingest_(ingest_ Ingest) *DocumentSimulationBuilder {
	rb.v.Ingest_ = ingest_
	return rb
}

// Parent_ set the Parent_ property for DocumentSimulationBuilder.
func (rb *DocumentSimulationBuilder) Parent_(parent_ string) *DocumentSimulationBuilder {
	rb.v.Parent_ = &parent_
	return rb
}

// Routing_ set the Routing_ property for DocumentSimulationBuilder.
func (rb *DocumentSimulationBuilder) Routing_(routing_ string) *DocumentSimulationBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

// Source_ set the Source_ property for DocumentSimulationBuilder.
func (rb *DocumentSimulationBuilder) Source_(value map[string]interface{}) *DocumentSimulationBuilder {
	rb.v.Source_ = value
	return rb
}
