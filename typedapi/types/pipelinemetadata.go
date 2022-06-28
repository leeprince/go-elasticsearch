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

// PipelineMetadata type.
type PipelineMetadata struct {
	Type    string `json:"type"`
	Version string `json:"version"`
}

// PipelineMetadataBuilder holds PipelineMetadata struct and provides a builder API.
type PipelineMetadataBuilder struct {
	v *PipelineMetadata
}

// NewPipelineMetadata provides a builder for the PipelineMetadata struct.
func NewPipelineMetadata() *PipelineMetadataBuilder {
	r := PipelineMetadataBuilder{
		&PipelineMetadata{},
	}

	return &r
}

// Build finalize the chain and returns the PipelineMetadata struct
func (rb *PipelineMetadataBuilder) Build() PipelineMetadata {
	return *rb.v
}

// Type set the Type property for PipelineMetadataBuilder.
func (rb *PipelineMetadataBuilder) Type_(type_ string) *PipelineMetadataBuilder {
	rb.v.Type = type_
	return rb
}

// Version set the Version property for PipelineMetadataBuilder.
func (rb *PipelineMetadataBuilder) Version(version string) *PipelineMetadataBuilder {
	rb.v.Version = version
	return rb
}
