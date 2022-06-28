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

// MlInferenceIngestProcessorCount type.
type MlInferenceIngestProcessorCount struct {
	Max int64 `json:"max"`
	Min int64 `json:"min"`
	Sum int64 `json:"sum"`
}

// MlInferenceIngestProcessorCountBuilder holds MlInferenceIngestProcessorCount struct and provides a builder API.
type MlInferenceIngestProcessorCountBuilder struct {
	v *MlInferenceIngestProcessorCount
}

// NewMlInferenceIngestProcessorCount provides a builder for the MlInferenceIngestProcessorCount struct.
func NewMlInferenceIngestProcessorCount() *MlInferenceIngestProcessorCountBuilder {
	r := MlInferenceIngestProcessorCountBuilder{
		&MlInferenceIngestProcessorCount{},
	}

	return &r
}

// Build finalize the chain and returns the MlInferenceIngestProcessorCount struct
func (rb *MlInferenceIngestProcessorCountBuilder) Build() MlInferenceIngestProcessorCount {
	return *rb.v
}

// Max set the Max property for MlInferenceIngestProcessorCountBuilder.
func (rb *MlInferenceIngestProcessorCountBuilder) Max(max int64) *MlInferenceIngestProcessorCountBuilder {
	rb.v.Max = max
	return rb
}

// Min set the Min property for MlInferenceIngestProcessorCountBuilder.
func (rb *MlInferenceIngestProcessorCountBuilder) Min(min int64) *MlInferenceIngestProcessorCountBuilder {
	rb.v.Min = min
	return rb
}

// Sum set the Sum property for MlInferenceIngestProcessorCountBuilder.
func (rb *MlInferenceIngestProcessorCountBuilder) Sum(sum int64) *MlInferenceIngestProcessorCountBuilder {
	rb.v.Sum = sum
	return rb
}
