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

// MlInference type.
type MlInference struct {
	Deployments      *MlInferenceDeployments               `json:"deployments,omitempty"`
	IngestProcessors map[string]MlInferenceIngestProcessor `json:"ingest_processors"`
	TrainedModels    MlInferenceTrainedModels              `json:"trained_models"`
}

// MlInferenceBuilder holds MlInference struct and provides a builder API.
type MlInferenceBuilder struct {
	v *MlInference
}

// NewMlInference provides a builder for the MlInference struct.
func NewMlInference() *MlInferenceBuilder {
	r := MlInferenceBuilder{
		&MlInference{
			IngestProcessors: make(map[string]MlInferenceIngestProcessor, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MlInference struct
func (rb *MlInferenceBuilder) Build() MlInference {
	return *rb.v
}

// Deployments set the Deployments property for MlInferenceBuilder.
func (rb *MlInferenceBuilder) Deployments(deployments MlInferenceDeployments) *MlInferenceBuilder {
	rb.v.Deployments = &deployments
	return rb
}

// IngestProcessors set the IngestProcessors property for MlInferenceBuilder.
func (rb *MlInferenceBuilder) IngestProcessors(value map[string]MlInferenceIngestProcessor) *MlInferenceBuilder {
	rb.v.IngestProcessors = value
	return rb
}

// TrainedModels set the TrainedModels property for MlInferenceBuilder.
func (rb *MlInferenceBuilder) TrainedModels(trainedmodels MlInferenceTrainedModels) *MlInferenceBuilder {
	rb.v.TrainedModels = trainedmodels
	return rb
}
