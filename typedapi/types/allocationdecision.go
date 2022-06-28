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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/allocationexplaindecision"
)

// AllocationDecision type.
type AllocationDecision struct {
	Decider     string                                              `json:"decider"`
	Decision    allocationexplaindecision.AllocationExplainDecision `json:"decision"`
	Explanation string                                              `json:"explanation"`
}

// AllocationDecisionBuilder holds AllocationDecision struct and provides a builder API.
type AllocationDecisionBuilder struct {
	v *AllocationDecision
}

// NewAllocationDecision provides a builder for the AllocationDecision struct.
func NewAllocationDecision() *AllocationDecisionBuilder {
	r := AllocationDecisionBuilder{
		&AllocationDecision{},
	}

	return &r
}

// Build finalize the chain and returns the AllocationDecision struct
func (rb *AllocationDecisionBuilder) Build() AllocationDecision {
	return *rb.v
}

// Decider set the Decider property for AllocationDecisionBuilder.
func (rb *AllocationDecisionBuilder) Decider(decider string) *AllocationDecisionBuilder {
	rb.v.Decider = decider
	return rb
}

// Decision set the Decision property for AllocationDecisionBuilder.
func (rb *AllocationDecisionBuilder) Decision(decision allocationexplaindecision.AllocationExplainDecision) *AllocationDecisionBuilder {
	rb.v.Decision = decision
	return rb
}

// Explanation set the Explanation property for AllocationDecisionBuilder.
func (rb *AllocationDecisionBuilder) Explanation(explanation string) *AllocationDecisionBuilder {
	rb.v.Explanation = explanation
	return rb
}
