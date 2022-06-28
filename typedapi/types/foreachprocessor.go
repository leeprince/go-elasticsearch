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

// ForeachProcessor type.
type ForeachProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Processor     *ProcessorContainer  `json:"processor,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// ForeachProcessorBuilder holds ForeachProcessor struct and provides a builder API.
type ForeachProcessorBuilder struct {
	v *ForeachProcessor
}

// NewForeachProcessor provides a builder for the ForeachProcessor struct.
func NewForeachProcessor() *ForeachProcessorBuilder {
	r := ForeachProcessorBuilder{
		&ForeachProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the ForeachProcessor struct
func (rb *ForeachProcessorBuilder) Build() ForeachProcessor {
	return *rb.v
}

// Field set the Field property for ForeachProcessorBuilder.
func (rb *ForeachProcessorBuilder) Field(field Field) *ForeachProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for ForeachProcessorBuilder.
func (rb *ForeachProcessorBuilder) If_(if_ string) *ForeachProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for ForeachProcessorBuilder.
func (rb *ForeachProcessorBuilder) IgnoreFailure(ignorefailure bool) *ForeachProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for ForeachProcessorBuilder.
func (rb *ForeachProcessorBuilder) IgnoreMissing(ignoremissing bool) *ForeachProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// OnFailure set the OnFailure property for ForeachProcessorBuilder.
func (rb *ForeachProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *ForeachProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Processor set the Processor property for ForeachProcessorBuilder.
func (rb *ForeachProcessorBuilder) Processor(processor ProcessorContainer) *ForeachProcessorBuilder {
	rb.v.Processor = &processor
	return rb
}

// Tag set the Tag property for ForeachProcessorBuilder.
func (rb *ForeachProcessorBuilder) Tag(tag string) *ForeachProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
