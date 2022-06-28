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

// RerouteDecision type.
type RerouteDecision struct {
	Decider     string `json:"decider"`
	Decision    string `json:"decision"`
	Explanation string `json:"explanation"`
}

// RerouteDecisionBuilder holds RerouteDecision struct and provides a builder API.
type RerouteDecisionBuilder struct {
	v *RerouteDecision
}

// NewRerouteDecision provides a builder for the RerouteDecision struct.
func NewRerouteDecision() *RerouteDecisionBuilder {
	r := RerouteDecisionBuilder{
		&RerouteDecision{},
	}

	return &r
}

// Build finalize the chain and returns the RerouteDecision struct
func (rb *RerouteDecisionBuilder) Build() RerouteDecision {
	return *rb.v
}

// Decider set the Decider property for RerouteDecisionBuilder.
func (rb *RerouteDecisionBuilder) Decider(decider string) *RerouteDecisionBuilder {
	rb.v.Decider = decider
	return rb
}

// Decision set the Decision property for RerouteDecisionBuilder.
func (rb *RerouteDecisionBuilder) Decision(decision string) *RerouteDecisionBuilder {
	rb.v.Decision = decision
	return rb
}

// Explanation set the Explanation property for RerouteDecisionBuilder.
func (rb *RerouteDecisionBuilder) Explanation(explanation string) *RerouteDecisionBuilder {
	rb.v.Explanation = explanation
	return rb
}
