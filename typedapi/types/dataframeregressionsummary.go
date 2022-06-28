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

// DataframeRegressionSummary type.
type DataframeRegressionSummary struct {
	Huber    *DataframeEvaluationValue `json:"huber,omitempty"`
	Mse      *DataframeEvaluationValue `json:"mse,omitempty"`
	Msle     *DataframeEvaluationValue `json:"msle,omitempty"`
	RSquared *DataframeEvaluationValue `json:"r_squared,omitempty"`
}

// DataframeRegressionSummaryBuilder holds DataframeRegressionSummary struct and provides a builder API.
type DataframeRegressionSummaryBuilder struct {
	v *DataframeRegressionSummary
}

// NewDataframeRegressionSummary provides a builder for the DataframeRegressionSummary struct.
func NewDataframeRegressionSummary() *DataframeRegressionSummaryBuilder {
	r := DataframeRegressionSummaryBuilder{
		&DataframeRegressionSummary{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeRegressionSummary struct
func (rb *DataframeRegressionSummaryBuilder) Build() DataframeRegressionSummary {
	return *rb.v
}

// Huber set the Huber property for DataframeRegressionSummaryBuilder.
func (rb *DataframeRegressionSummaryBuilder) Huber(huber DataframeEvaluationValue) *DataframeRegressionSummaryBuilder {
	rb.v.Huber = &huber
	return rb
}

// Mse set the Mse property for DataframeRegressionSummaryBuilder.
func (rb *DataframeRegressionSummaryBuilder) Mse(mse DataframeEvaluationValue) *DataframeRegressionSummaryBuilder {
	rb.v.Mse = &mse
	return rb
}

// Msle set the Msle property for DataframeRegressionSummaryBuilder.
func (rb *DataframeRegressionSummaryBuilder) Msle(msle DataframeEvaluationValue) *DataframeRegressionSummaryBuilder {
	rb.v.Msle = &msle
	return rb
}

// RSquared set the RSquared property for DataframeRegressionSummaryBuilder.
func (rb *DataframeRegressionSummaryBuilder) RSquared(rsquared DataframeEvaluationValue) *DataframeRegressionSummaryBuilder {
	rb.v.RSquared = &rsquared
	return rb
}
