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

// RemoveProcessor type.
type RemoveProcessor struct {
	Field         Fields               `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// RemoveProcessorBuilder holds RemoveProcessor struct and provides a builder API.
type RemoveProcessorBuilder struct {
	v *RemoveProcessor
}

// NewRemoveProcessor provides a builder for the RemoveProcessor struct.
func NewRemoveProcessor() *RemoveProcessorBuilder {
	r := RemoveProcessorBuilder{
		&RemoveProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the RemoveProcessor struct
func (rb *RemoveProcessorBuilder) Build() RemoveProcessor {
	return *rb.v
}

// Field set the Field property for RemoveProcessorBuilder.
func (rb *RemoveProcessorBuilder) Field(field Fields) *RemoveProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for RemoveProcessorBuilder.
func (rb *RemoveProcessorBuilder) If_(if_ string) *RemoveProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for RemoveProcessorBuilder.
func (rb *RemoveProcessorBuilder) IgnoreFailure(ignorefailure bool) *RemoveProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for RemoveProcessorBuilder.
func (rb *RemoveProcessorBuilder) IgnoreMissing(ignoremissing bool) *RemoveProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// OnFailure set the OnFailure property for RemoveProcessorBuilder.
func (rb *RemoveProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *RemoveProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for RemoveProcessorBuilder.
func (rb *RemoveProcessorBuilder) Tag(tag string) *RemoveProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
