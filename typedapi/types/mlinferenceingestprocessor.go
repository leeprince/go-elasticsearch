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

// MlInferenceIngestProcessor type.
type MlInferenceIngestProcessor struct {
	NumDocsProcessed MlInferenceIngestProcessorCount `json:"num_docs_processed"`
	NumFailures      MlInferenceIngestProcessorCount `json:"num_failures"`
	Pipelines        MlCounter                       `json:"pipelines"`
	TimeMs           MlInferenceIngestProcessorCount `json:"time_ms"`
}

// MlInferenceIngestProcessorBuilder holds MlInferenceIngestProcessor struct and provides a builder API.
type MlInferenceIngestProcessorBuilder struct {
	v *MlInferenceIngestProcessor
}

// NewMlInferenceIngestProcessor provides a builder for the MlInferenceIngestProcessor struct.
func NewMlInferenceIngestProcessor() *MlInferenceIngestProcessorBuilder {
	r := MlInferenceIngestProcessorBuilder{
		&MlInferenceIngestProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the MlInferenceIngestProcessor struct
func (rb *MlInferenceIngestProcessorBuilder) Build() MlInferenceIngestProcessor {
	return *rb.v
}

// NumDocsProcessed set the NumDocsProcessed property for MlInferenceIngestProcessorBuilder.
func (rb *MlInferenceIngestProcessorBuilder) NumDocsProcessed(numdocsprocessed MlInferenceIngestProcessorCount) *MlInferenceIngestProcessorBuilder {
	rb.v.NumDocsProcessed = numdocsprocessed
	return rb
}

// NumFailures set the NumFailures property for MlInferenceIngestProcessorBuilder.
func (rb *MlInferenceIngestProcessorBuilder) NumFailures(numfailures MlInferenceIngestProcessorCount) *MlInferenceIngestProcessorBuilder {
	rb.v.NumFailures = numfailures
	return rb
}

// Pipelines set the Pipelines property for MlInferenceIngestProcessorBuilder.
func (rb *MlInferenceIngestProcessorBuilder) Pipelines(pipelines MlCounter) *MlInferenceIngestProcessorBuilder {
	rb.v.Pipelines = pipelines
	return rb
}

// TimeMs set the TimeMs property for MlInferenceIngestProcessorBuilder.
func (rb *MlInferenceIngestProcessorBuilder) TimeMs(timems MlInferenceIngestProcessorCount) *MlInferenceIngestProcessorBuilder {
	rb.v.TimeMs = timems
	return rb
}
