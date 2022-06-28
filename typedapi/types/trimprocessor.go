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

// TrimProcessor type.
type TrimProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
}

// TrimProcessorBuilder holds TrimProcessor struct and provides a builder API.
type TrimProcessorBuilder struct {
	v *TrimProcessor
}

// NewTrimProcessor provides a builder for the TrimProcessor struct.
func NewTrimProcessor() *TrimProcessorBuilder {
	r := TrimProcessorBuilder{
		&TrimProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the TrimProcessor struct
func (rb *TrimProcessorBuilder) Build() TrimProcessor {
	return *rb.v
}

// Field set the Field property for TrimProcessorBuilder.
func (rb *TrimProcessorBuilder) Field(field Field) *TrimProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for TrimProcessorBuilder.
func (rb *TrimProcessorBuilder) If_(if_ string) *TrimProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for TrimProcessorBuilder.
func (rb *TrimProcessorBuilder) IgnoreFailure(ignorefailure bool) *TrimProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for TrimProcessorBuilder.
func (rb *TrimProcessorBuilder) IgnoreMissing(ignoremissing bool) *TrimProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// OnFailure set the OnFailure property for TrimProcessorBuilder.
func (rb *TrimProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *TrimProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for TrimProcessorBuilder.
func (rb *TrimProcessorBuilder) Tag(tag string) *TrimProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for TrimProcessorBuilder.
func (rb *TrimProcessorBuilder) TargetField(targetfield Field) *TrimProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
