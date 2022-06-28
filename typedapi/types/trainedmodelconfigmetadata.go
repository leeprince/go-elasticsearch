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

// TrainedModelConfigMetadata type.
type TrainedModelConfigMetadata struct {
	// FeatureImportanceBaseline An object that contains the baseline for feature importance values. For
	// regression analysis, it is a single value. For classification analysis, there
	// is a value for each class.
	FeatureImportanceBaseline map[string]string `json:"feature_importance_baseline,omitempty"`
	// Hyperparameters List of the available hyperparameters optimized during the
	// fine_parameter_tuning phase as well as specified by the user.
	Hyperparameters []Hyperparameter `json:"hyperparameters,omitempty"`
	ModelAliases    []string         `json:"model_aliases,omitempty"`
	// TotalFeatureImportance An array of the total feature importance for each feature used from the
	// training data set. This array of objects is returned if data frame analytics
	// trained the model and the request includes total_feature_importance in the
	// include request parameter.
	TotalFeatureImportance []TotalFeatureImportance `json:"total_feature_importance,omitempty"`
}

// TrainedModelConfigMetadataBuilder holds TrainedModelConfigMetadata struct and provides a builder API.
type TrainedModelConfigMetadataBuilder struct {
	v *TrainedModelConfigMetadata
}

// NewTrainedModelConfigMetadata provides a builder for the TrainedModelConfigMetadata struct.
func NewTrainedModelConfigMetadata() *TrainedModelConfigMetadataBuilder {
	r := TrainedModelConfigMetadataBuilder{
		&TrainedModelConfigMetadata{
			FeatureImportanceBaseline: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelConfigMetadata struct
func (rb *TrainedModelConfigMetadataBuilder) Build() TrainedModelConfigMetadata {
	return *rb.v
}

// FeatureImportanceBaseline An object that contains the baseline for feature importance values. For
// regression analysis, it is a single value. For classification analysis, there
// is a value for each class.
func (rb *TrainedModelConfigMetadataBuilder) FeatureImportanceBaseline(value map[string]string) *TrainedModelConfigMetadataBuilder {
	rb.v.FeatureImportanceBaseline = value
	return rb
}

// Hyperparameters List of the available hyperparameters optimized during the
// fine_parameter_tuning phase as well as specified by the user.
func (rb *TrainedModelConfigMetadataBuilder) Hyperparameters(hyperparameters ...Hyperparameter) *TrainedModelConfigMetadataBuilder {
	rb.v.Hyperparameters = append(rb.v.Hyperparameters, hyperparameters...)
	return rb
}

// ModelAliases set the ModelAliases property for TrainedModelConfigMetadataBuilder.
func (rb *TrainedModelConfigMetadataBuilder) ModelAliases(model_aliases ...string) *TrainedModelConfigMetadataBuilder {
	rb.v.ModelAliases = append(rb.v.ModelAliases, model_aliases...)
	return rb
}

// TotalFeatureImportance An array of the total feature importance for each feature used from the
// training data set. This array of objects is returned if data frame analytics
// trained the model and the request includes total_feature_importance in the
// include request parameter.
func (rb *TrainedModelConfigMetadataBuilder) TotalFeatureImportance(total_feature_importance ...TotalFeatureImportance) *TrainedModelConfigMetadataBuilder {
	rb.v.TotalFeatureImportance = append(rb.v.TotalFeatureImportance, total_feature_importance...)
	return rb
}
