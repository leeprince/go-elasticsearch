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

// RerouteExplanation type.
type RerouteExplanation struct {
	Command    string            `json:"command"`
	Decisions  []RerouteDecision `json:"decisions"`
	Parameters RerouteParameters `json:"parameters"`
}

// RerouteExplanationBuilder holds RerouteExplanation struct and provides a builder API.
type RerouteExplanationBuilder struct {
	v *RerouteExplanation
}

// NewRerouteExplanation provides a builder for the RerouteExplanation struct.
func NewRerouteExplanation() *RerouteExplanationBuilder {
	r := RerouteExplanationBuilder{
		&RerouteExplanation{},
	}

	return &r
}

// Build finalize the chain and returns the RerouteExplanation struct
func (rb *RerouteExplanationBuilder) Build() RerouteExplanation {
	return *rb.v
}

// Command set the Command property for RerouteExplanationBuilder.
func (rb *RerouteExplanationBuilder) Command(command string) *RerouteExplanationBuilder {
	rb.v.Command = command
	return rb
}

// Decisions set the Decisions property for RerouteExplanationBuilder.
func (rb *RerouteExplanationBuilder) Decisions(decisions ...RerouteDecision) *RerouteExplanationBuilder {
	rb.v.Decisions = append(rb.v.Decisions, decisions...)
	return rb
}

// Parameters set the Parameters property for RerouteExplanationBuilder.
func (rb *RerouteExplanationBuilder) Parameters(parameters RerouteParameters) *RerouteExplanationBuilder {
	rb.v.Parameters = parameters
	return rb
}
