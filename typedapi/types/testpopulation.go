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

// TestPopulation type.
type TestPopulation struct {
	Field  Field           `json:"field"`
	Filter *QueryContainer `json:"filter,omitempty"`
	Script *Script         `json:"script,omitempty"`
}

// TestPopulationBuilder holds TestPopulation struct and provides a builder API.
type TestPopulationBuilder struct {
	v *TestPopulation
}

// NewTestPopulation provides a builder for the TestPopulation struct.
func NewTestPopulation() *TestPopulationBuilder {
	r := TestPopulationBuilder{
		&TestPopulation{},
	}

	return &r
}

// Build finalize the chain and returns the TestPopulation struct
func (rb *TestPopulationBuilder) Build() TestPopulation {
	return *rb.v
}

// Field set the Field property for TestPopulationBuilder.
func (rb *TestPopulationBuilder) Field(field Field) *TestPopulationBuilder {
	rb.v.Field = field
	return rb
}

// Filter set the Filter property for TestPopulationBuilder.
func (rb *TestPopulationBuilder) Filter(filter QueryContainer) *TestPopulationBuilder {
	rb.v.Filter = &filter
	return rb
}

// Script set the Script property for TestPopulationBuilder.
func (rb *TestPopulationBuilder) Script(script Script) *TestPopulationBuilder {
	rb.v.Script = &script
	return rb
}
