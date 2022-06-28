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

// UppercaseProcessor type.
type UppercaseProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
}

// UppercaseProcessorBuilder holds UppercaseProcessor struct and provides a builder API.
type UppercaseProcessorBuilder struct {
	v *UppercaseProcessor
}

// NewUppercaseProcessor provides a builder for the UppercaseProcessor struct.
func NewUppercaseProcessor() *UppercaseProcessorBuilder {
	r := UppercaseProcessorBuilder{
		&UppercaseProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the UppercaseProcessor struct
func (rb *UppercaseProcessorBuilder) Build() UppercaseProcessor {
	return *rb.v
}

// Field set the Field property for UppercaseProcessorBuilder.
func (rb *UppercaseProcessorBuilder) Field(field Field) *UppercaseProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for UppercaseProcessorBuilder.
func (rb *UppercaseProcessorBuilder) If_(if_ string) *UppercaseProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for UppercaseProcessorBuilder.
func (rb *UppercaseProcessorBuilder) IgnoreFailure(ignorefailure bool) *UppercaseProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for UppercaseProcessorBuilder.
func (rb *UppercaseProcessorBuilder) IgnoreMissing(ignoremissing bool) *UppercaseProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// OnFailure set the OnFailure property for UppercaseProcessorBuilder.
func (rb *UppercaseProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *UppercaseProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for UppercaseProcessorBuilder.
func (rb *UppercaseProcessorBuilder) Tag(tag string) *UppercaseProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for UppercaseProcessorBuilder.
func (rb *UppercaseProcessorBuilder) TargetField(targetfield Field) *UppercaseProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
