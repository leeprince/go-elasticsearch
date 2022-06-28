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

// Pipeline type.
type Pipeline struct {
	Description      string           `json:"description"`
	LastModified     Timestamp        `json:"last_modified"`
	Pipeline         string           `json:"pipeline"`
	PipelineMetadata PipelineMetadata `json:"pipeline_metadata"`
	PipelineSettings PipelineSettings `json:"pipeline_settings"`
	Username         string           `json:"username"`
}

// PipelineBuilder holds Pipeline struct and provides a builder API.
type PipelineBuilder struct {
	v *Pipeline
}

// NewPipeline provides a builder for the Pipeline struct.
func NewPipeline() *PipelineBuilder {
	r := PipelineBuilder{
		&Pipeline{},
	}

	return &r
}

// Build finalize the chain and returns the Pipeline struct
func (rb *PipelineBuilder) Build() Pipeline {
	return *rb.v
}

// Description set the Description property for PipelineBuilder.
func (rb *PipelineBuilder) Description(description string) *PipelineBuilder {
	rb.v.Description = description
	return rb
}

// LastModified set the LastModified property for PipelineBuilder.
func (rb *PipelineBuilder) LastModified(lastmodified Timestamp) *PipelineBuilder {
	rb.v.LastModified = lastmodified
	return rb
}

// Pipeline set the Pipeline property for PipelineBuilder.
func (rb *PipelineBuilder) Pipeline(pipeline string) *PipelineBuilder {
	rb.v.Pipeline = pipeline
	return rb
}

// PipelineMetadata set the PipelineMetadata property for PipelineBuilder.
func (rb *PipelineBuilder) PipelineMetadata(pipelinemetadata PipelineMetadata) *PipelineBuilder {
	rb.v.PipelineMetadata = pipelinemetadata
	return rb
}

// PipelineSettings set the PipelineSettings property for PipelineBuilder.
func (rb *PipelineBuilder) PipelineSettings(pipelinesettings PipelineSettings) *PipelineBuilder {
	rb.v.PipelineSettings = pipelinesettings
	return rb
}

// Username set the Username property for PipelineBuilder.
func (rb *PipelineBuilder) Username(username string) *PipelineBuilder {
	rb.v.Username = username
	return rb
}
