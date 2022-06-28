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

// IndicesValidationExplanation type.
type IndicesValidationExplanation struct {
	Error       *string   `json:"error,omitempty"`
	Explanation *string   `json:"explanation,omitempty"`
	Index       IndexName `json:"index"`
	Valid       bool      `json:"valid"`
}

// IndicesValidationExplanationBuilder holds IndicesValidationExplanation struct and provides a builder API.
type IndicesValidationExplanationBuilder struct {
	v *IndicesValidationExplanation
}

// NewIndicesValidationExplanation provides a builder for the IndicesValidationExplanation struct.
func NewIndicesValidationExplanation() *IndicesValidationExplanationBuilder {
	r := IndicesValidationExplanationBuilder{
		&IndicesValidationExplanation{},
	}

	return &r
}

// Build finalize the chain and returns the IndicesValidationExplanation struct
func (rb *IndicesValidationExplanationBuilder) Build() IndicesValidationExplanation {
	return *rb.v
}

// Error set the Error property for IndicesValidationExplanationBuilder.
func (rb *IndicesValidationExplanationBuilder) Error(error string) *IndicesValidationExplanationBuilder {
	rb.v.Error = &error
	return rb
}

// Explanation set the Explanation property for IndicesValidationExplanationBuilder.
func (rb *IndicesValidationExplanationBuilder) Explanation(explanation string) *IndicesValidationExplanationBuilder {
	rb.v.Explanation = &explanation
	return rb
}

// Index set the Index property for IndicesValidationExplanationBuilder.
func (rb *IndicesValidationExplanationBuilder) Index(index IndexName) *IndicesValidationExplanationBuilder {
	rb.v.Index = index
	return rb
}

// Valid set the Valid property for IndicesValidationExplanationBuilder.
func (rb *IndicesValidationExplanationBuilder) Valid(valid bool) *IndicesValidationExplanationBuilder {
	rb.v.Valid = valid
	return rb
}
