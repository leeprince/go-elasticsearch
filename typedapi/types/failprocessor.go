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

// FailProcessor type.
type FailProcessor struct {
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	Message       string               `json:"message"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// FailProcessorBuilder holds FailProcessor struct and provides a builder API.
type FailProcessorBuilder struct {
	v *FailProcessor
}

// NewFailProcessor provides a builder for the FailProcessor struct.
func NewFailProcessor() *FailProcessorBuilder {
	r := FailProcessorBuilder{
		&FailProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the FailProcessor struct
func (rb *FailProcessorBuilder) Build() FailProcessor {
	return *rb.v
}

// If set the If property for FailProcessorBuilder.
func (rb *FailProcessorBuilder) If_(if_ string) *FailProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for FailProcessorBuilder.
func (rb *FailProcessorBuilder) IgnoreFailure(ignorefailure bool) *FailProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// Message set the Message property for FailProcessorBuilder.
func (rb *FailProcessorBuilder) Message(message string) *FailProcessorBuilder {
	rb.v.Message = message
	return rb
}

// OnFailure set the OnFailure property for FailProcessorBuilder.
func (rb *FailProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *FailProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for FailProcessorBuilder.
func (rb *FailProcessorBuilder) Tag(tag string) *FailProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
