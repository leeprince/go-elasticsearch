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

// ConfusionMatrixPrediction type.
type ConfusionMatrixPrediction struct {
	Count          int  `json:"count"`
	PredictedClass Name `json:"predicted_class"`
}

// ConfusionMatrixPredictionBuilder holds ConfusionMatrixPrediction struct and provides a builder API.
type ConfusionMatrixPredictionBuilder struct {
	v *ConfusionMatrixPrediction
}

// NewConfusionMatrixPrediction provides a builder for the ConfusionMatrixPrediction struct.
func NewConfusionMatrixPrediction() *ConfusionMatrixPredictionBuilder {
	r := ConfusionMatrixPredictionBuilder{
		&ConfusionMatrixPrediction{},
	}

	return &r
}

// Build finalize the chain and returns the ConfusionMatrixPrediction struct
func (rb *ConfusionMatrixPredictionBuilder) Build() ConfusionMatrixPrediction {
	return *rb.v
}

// Count set the Count property for ConfusionMatrixPredictionBuilder.
func (rb *ConfusionMatrixPredictionBuilder) Count(count int) *ConfusionMatrixPredictionBuilder {
	rb.v.Count = count
	return rb
}

// PredictedClass set the PredictedClass property for ConfusionMatrixPredictionBuilder.
func (rb *ConfusionMatrixPredictionBuilder) PredictedClass(predictedclass Name) *ConfusionMatrixPredictionBuilder {
	rb.v.PredictedClass = predictedclass
	return rb
}
