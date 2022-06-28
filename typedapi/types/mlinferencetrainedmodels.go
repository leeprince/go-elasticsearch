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

// MlInferenceTrainedModels type.
type MlInferenceTrainedModels struct {
	All_                          MlCounter                      `json:"_all"`
	Count                         *MlInferenceTrainedModelsCount `json:"count,omitempty"`
	EstimatedHeapMemoryUsageBytes *JobStatistics                 `json:"estimated_heap_memory_usage_bytes,omitempty"`
	EstimatedOperations           *JobStatistics                 `json:"estimated_operations,omitempty"`
	ModelSizeBytes                *JobStatistics                 `json:"model_size_bytes,omitempty"`
}

// MlInferenceTrainedModelsBuilder holds MlInferenceTrainedModels struct and provides a builder API.
type MlInferenceTrainedModelsBuilder struct {
	v *MlInferenceTrainedModels
}

// NewMlInferenceTrainedModels provides a builder for the MlInferenceTrainedModels struct.
func NewMlInferenceTrainedModels() *MlInferenceTrainedModelsBuilder {
	r := MlInferenceTrainedModelsBuilder{
		&MlInferenceTrainedModels{},
	}

	return &r
}

// Build finalize the chain and returns the MlInferenceTrainedModels struct
func (rb *MlInferenceTrainedModelsBuilder) Build() MlInferenceTrainedModels {
	return *rb.v
}

// All_ set the All_ property for MlInferenceTrainedModelsBuilder.
func (rb *MlInferenceTrainedModelsBuilder) All_(all_ MlCounter) *MlInferenceTrainedModelsBuilder {
	rb.v.All_ = all_
	return rb
}

// Count set the Count property for MlInferenceTrainedModelsBuilder.
func (rb *MlInferenceTrainedModelsBuilder) Count(count MlInferenceTrainedModelsCount) *MlInferenceTrainedModelsBuilder {
	rb.v.Count = &count
	return rb
}

// EstimatedHeapMemoryUsageBytes set the EstimatedHeapMemoryUsageBytes property for MlInferenceTrainedModelsBuilder.
func (rb *MlInferenceTrainedModelsBuilder) EstimatedHeapMemoryUsageBytes(estimatedheapmemoryusagebytes JobStatistics) *MlInferenceTrainedModelsBuilder {
	rb.v.EstimatedHeapMemoryUsageBytes = &estimatedheapmemoryusagebytes
	return rb
}

// EstimatedOperations set the EstimatedOperations property for MlInferenceTrainedModelsBuilder.
func (rb *MlInferenceTrainedModelsBuilder) EstimatedOperations(estimatedoperations JobStatistics) *MlInferenceTrainedModelsBuilder {
	rb.v.EstimatedOperations = &estimatedoperations
	return rb
}

// ModelSizeBytes set the ModelSizeBytes property for MlInferenceTrainedModelsBuilder.
func (rb *MlInferenceTrainedModelsBuilder) ModelSizeBytes(modelsizebytes JobStatistics) *MlInferenceTrainedModelsBuilder {
	rb.v.ModelSizeBytes = &modelsizebytes
	return rb
}
