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

// AggregateOutput type.
type AggregateOutput struct {
	Exponent           *Weights `json:"exponent,omitempty"`
	LogisticRegression *Weights `json:"logistic_regression,omitempty"`
	WeightedMode       *Weights `json:"weighted_mode,omitempty"`
	WeightedSum        *Weights `json:"weighted_sum,omitempty"`
}

// AggregateOutputBuilder holds AggregateOutput struct and provides a builder API.
type AggregateOutputBuilder struct {
	v *AggregateOutput
}

// NewAggregateOutput provides a builder for the AggregateOutput struct.
func NewAggregateOutput() *AggregateOutputBuilder {
	r := AggregateOutputBuilder{
		&AggregateOutput{},
	}

	return &r
}

// Build finalize the chain and returns the AggregateOutput struct
func (rb *AggregateOutputBuilder) Build() AggregateOutput {
	return *rb.v
}

// Exponent set the Exponent property for AggregateOutputBuilder.
func (rb *AggregateOutputBuilder) Exponent(exponent Weights) *AggregateOutputBuilder {
	rb.v.Exponent = &exponent
	return rb
}

// LogisticRegression set the LogisticRegression property for AggregateOutputBuilder.
func (rb *AggregateOutputBuilder) LogisticRegression(logisticregression Weights) *AggregateOutputBuilder {
	rb.v.LogisticRegression = &logisticregression
	return rb
}

// WeightedMode set the WeightedMode property for AggregateOutputBuilder.
func (rb *AggregateOutputBuilder) WeightedMode(weightedmode Weights) *AggregateOutputBuilder {
	rb.v.WeightedMode = &weightedmode
	return rb
}

// WeightedSum set the WeightedSum property for AggregateOutputBuilder.
func (rb *AggregateOutputBuilder) WeightedSum(weightedsum Weights) *AggregateOutputBuilder {
	rb.v.WeightedSum = &weightedsum
	return rb
}
