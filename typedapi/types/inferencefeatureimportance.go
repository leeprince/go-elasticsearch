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

// InferenceFeatureImportance type.
type InferenceFeatureImportance struct {
	Classes     []InferenceClassImportance `json:"classes,omitempty"`
	FeatureName string                     `json:"feature_name"`
	Importance  *float64                   `json:"importance,omitempty"`
}

// InferenceFeatureImportanceBuilder holds InferenceFeatureImportance struct and provides a builder API.
type InferenceFeatureImportanceBuilder struct {
	v *InferenceFeatureImportance
}

// NewInferenceFeatureImportance provides a builder for the InferenceFeatureImportance struct.
func NewInferenceFeatureImportance() *InferenceFeatureImportanceBuilder {
	r := InferenceFeatureImportanceBuilder{
		&InferenceFeatureImportance{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceFeatureImportance struct
func (rb *InferenceFeatureImportanceBuilder) Build() InferenceFeatureImportance {
	return *rb.v
}

// Classes set the Classes property for InferenceFeatureImportanceBuilder.
func (rb *InferenceFeatureImportanceBuilder) Classes(classes ...InferenceClassImportance) *InferenceFeatureImportanceBuilder {
	rb.v.Classes = append(rb.v.Classes, classes...)
	return rb
}

// FeatureName set the FeatureName property for InferenceFeatureImportanceBuilder.
func (rb *InferenceFeatureImportanceBuilder) FeatureName(featurename string) *InferenceFeatureImportanceBuilder {
	rb.v.FeatureName = featurename
	return rb
}

// Importance set the Importance property for InferenceFeatureImportanceBuilder.
func (rb *InferenceFeatureImportanceBuilder) Importance(importance float64) *InferenceFeatureImportanceBuilder {
	rb.v.Importance = &importance
	return rb
}
