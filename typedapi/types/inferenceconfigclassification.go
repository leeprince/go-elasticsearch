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

// InferenceConfigClassification type.
type InferenceConfigClassification struct {
	NumTopClasses                 *int    `json:"num_top_classes,omitempty"`
	NumTopFeatureImportanceValues *int    `json:"num_top_feature_importance_values,omitempty"`
	PredictionFieldType           *string `json:"prediction_field_type,omitempty"`
	ResultsField                  *Field  `json:"results_field,omitempty"`
	TopClassesResultsField        *Field  `json:"top_classes_results_field,omitempty"`
}

// InferenceConfigClassificationBuilder holds InferenceConfigClassification struct and provides a builder API.
type InferenceConfigClassificationBuilder struct {
	v *InferenceConfigClassification
}

// NewInferenceConfigClassification provides a builder for the InferenceConfigClassification struct.
func NewInferenceConfigClassification() *InferenceConfigClassificationBuilder {
	r := InferenceConfigClassificationBuilder{
		&InferenceConfigClassification{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceConfigClassification struct
func (rb *InferenceConfigClassificationBuilder) Build() InferenceConfigClassification {
	return *rb.v
}

// NumTopClasses set the NumTopClasses property for InferenceConfigClassificationBuilder.
func (rb *InferenceConfigClassificationBuilder) NumTopClasses(numtopclasses int) *InferenceConfigClassificationBuilder {
	rb.v.NumTopClasses = &numtopclasses
	return rb
}

// NumTopFeatureImportanceValues set the NumTopFeatureImportanceValues property for InferenceConfigClassificationBuilder.
func (rb *InferenceConfigClassificationBuilder) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *InferenceConfigClassificationBuilder {
	rb.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues
	return rb
}

// PredictionFieldType set the PredictionFieldType property for InferenceConfigClassificationBuilder.
func (rb *InferenceConfigClassificationBuilder) PredictionFieldType(predictionfieldtype string) *InferenceConfigClassificationBuilder {
	rb.v.PredictionFieldType = &predictionfieldtype
	return rb
}

// ResultsField set the ResultsField property for InferenceConfigClassificationBuilder.
func (rb *InferenceConfigClassificationBuilder) ResultsField(resultsfield Field) *InferenceConfigClassificationBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// TopClassesResultsField set the TopClassesResultsField property for InferenceConfigClassificationBuilder.
func (rb *InferenceConfigClassificationBuilder) TopClassesResultsField(topclassesresultsfield Field) *InferenceConfigClassificationBuilder {
	rb.v.TopClassesResultsField = &topclassesresultsfield
	return rb
}
