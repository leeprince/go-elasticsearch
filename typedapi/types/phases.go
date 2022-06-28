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

// Phases type.
type Phases struct {
	Cold   *Phase `json:"cold,omitempty"`
	Delete *Phase `json:"delete,omitempty"`
	Hot    *Phase `json:"hot,omitempty"`
	Warm   *Phase `json:"warm,omitempty"`
}

// PhasesBuilder holds Phases struct and provides a builder API.
type PhasesBuilder struct {
	v *Phases
}

// NewPhases provides a builder for the Phases struct.
func NewPhases() *PhasesBuilder {
	r := PhasesBuilder{
		&Phases{},
	}

	return &r
}

// Build finalize the chain and returns the Phases struct
func (rb *PhasesBuilder) Build() Phases {
	return *rb.v
}

// Cold set the Cold property for PhasesBuilder.
func (rb *PhasesBuilder) Cold(cold Phase) *PhasesBuilder {
	rb.v.Cold = &cold
	return rb
}

// Delete set the Delete property for PhasesBuilder.
func (rb *PhasesBuilder) Delete(delete Phase) *PhasesBuilder {
	rb.v.Delete = &delete
	return rb
}

// Hot set the Hot property for PhasesBuilder.
func (rb *PhasesBuilder) Hot(hot Phase) *PhasesBuilder {
	rb.v.Hot = &hot
	return rb
}

// Warm set the Warm property for PhasesBuilder.
func (rb *PhasesBuilder) Warm(warm Phase) *PhasesBuilder {
	rb.v.Warm = &warm
	return rb
}
