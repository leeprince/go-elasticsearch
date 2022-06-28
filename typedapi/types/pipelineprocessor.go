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

// PipelineProcessor type.
type PipelineProcessor struct {
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	Name          Name                 `json:"name"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// PipelineProcessorBuilder holds PipelineProcessor struct and provides a builder API.
type PipelineProcessorBuilder struct {
	v *PipelineProcessor
}

// NewPipelineProcessor provides a builder for the PipelineProcessor struct.
func NewPipelineProcessor() *PipelineProcessorBuilder {
	r := PipelineProcessorBuilder{
		&PipelineProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the PipelineProcessor struct
func (rb *PipelineProcessorBuilder) Build() PipelineProcessor {
	return *rb.v
}

// If set the If property for PipelineProcessorBuilder.
func (rb *PipelineProcessorBuilder) If_(if_ string) *PipelineProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for PipelineProcessorBuilder.
func (rb *PipelineProcessorBuilder) IgnoreFailure(ignorefailure bool) *PipelineProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// Name set the Name property for PipelineProcessorBuilder.
func (rb *PipelineProcessorBuilder) Name(name Name) *PipelineProcessorBuilder {
	rb.v.Name = name
	return rb
}

// OnFailure set the OnFailure property for PipelineProcessorBuilder.
func (rb *PipelineProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *PipelineProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for PipelineProcessorBuilder.
func (rb *PipelineProcessorBuilder) Tag(tag string) *PipelineProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
