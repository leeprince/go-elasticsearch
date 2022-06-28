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

// LowercaseProcessor type.
type LowercaseProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
}

// LowercaseProcessorBuilder holds LowercaseProcessor struct and provides a builder API.
type LowercaseProcessorBuilder struct {
	v *LowercaseProcessor
}

// NewLowercaseProcessor provides a builder for the LowercaseProcessor struct.
func NewLowercaseProcessor() *LowercaseProcessorBuilder {
	r := LowercaseProcessorBuilder{
		&LowercaseProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the LowercaseProcessor struct
func (rb *LowercaseProcessorBuilder) Build() LowercaseProcessor {
	return *rb.v
}

// Field set the Field property for LowercaseProcessorBuilder.
func (rb *LowercaseProcessorBuilder) Field(field Field) *LowercaseProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for LowercaseProcessorBuilder.
func (rb *LowercaseProcessorBuilder) If_(if_ string) *LowercaseProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for LowercaseProcessorBuilder.
func (rb *LowercaseProcessorBuilder) IgnoreFailure(ignorefailure bool) *LowercaseProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for LowercaseProcessorBuilder.
func (rb *LowercaseProcessorBuilder) IgnoreMissing(ignoremissing bool) *LowercaseProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// OnFailure set the OnFailure property for LowercaseProcessorBuilder.
func (rb *LowercaseProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *LowercaseProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for LowercaseProcessorBuilder.
func (rb *LowercaseProcessorBuilder) Tag(tag string) *LowercaseProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for LowercaseProcessorBuilder.
func (rb *LowercaseProcessorBuilder) TargetField(targetfield Field) *LowercaseProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
