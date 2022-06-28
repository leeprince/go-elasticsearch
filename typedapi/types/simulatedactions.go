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

// SimulatedActions type.
type SimulatedActions struct {
	Actions []string          `json:"actions"`
	All     *SimulatedActions `json:"all,omitempty"`
	UseAll  bool              `json:"use_all"`
}

// SimulatedActionsBuilder holds SimulatedActions struct and provides a builder API.
type SimulatedActionsBuilder struct {
	v *SimulatedActions
}

// NewSimulatedActions provides a builder for the SimulatedActions struct.
func NewSimulatedActions() *SimulatedActionsBuilder {
	r := SimulatedActionsBuilder{
		&SimulatedActions{},
	}

	return &r
}

// Build finalize the chain and returns the SimulatedActions struct
func (rb *SimulatedActionsBuilder) Build() SimulatedActions {
	return *rb.v
}

// Actions set the Actions property for SimulatedActionsBuilder.
func (rb *SimulatedActionsBuilder) Actions(actions ...string) *SimulatedActionsBuilder {
	rb.v.Actions = append(rb.v.Actions, actions...)
	return rb
}

// All set the All property for SimulatedActionsBuilder.
func (rb *SimulatedActionsBuilder) All(all SimulatedActions) *SimulatedActionsBuilder {
	rb.v.All = &all
	return rb
}

// UseAll set the UseAll property for SimulatedActionsBuilder.
func (rb *SimulatedActionsBuilder) UseAll(useall bool) *SimulatedActionsBuilder {
	rb.v.UseAll = useall
	return rb
}
