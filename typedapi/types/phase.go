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

// Phase type.
type Phase struct {
	Actions        *Actions        `json:"actions,omitempty"`
	Configurations *Configurations `json:"configurations,omitempty"`
	MinAge         *Time           `json:"min_age,omitempty"`
}

// PhaseBuilder holds Phase struct and provides a builder API.
type PhaseBuilder struct {
	v *Phase
}

// NewPhase provides a builder for the Phase struct.
func NewPhase() *PhaseBuilder {
	r := PhaseBuilder{
		&Phase{},
	}

	return &r
}

// Build finalize the chain and returns the Phase struct
func (rb *PhaseBuilder) Build() Phase {
	return *rb.v
}

// Actions set the Actions property for PhaseBuilder.
func (rb *PhaseBuilder) Actions(actions Actions) *PhaseBuilder {
	rb.v.Actions = &actions
	return rb
}

// Configurations set the Configurations property for PhaseBuilder.
func (rb *PhaseBuilder) Configurations(configurations Configurations) *PhaseBuilder {
	rb.v.Configurations = &configurations
	return rb
}

// MinAge set the MinAge property for PhaseBuilder.
func (rb *PhaseBuilder) MinAge(minage Time) *PhaseBuilder {
	rb.v.MinAge = &minage
	return rb
}
