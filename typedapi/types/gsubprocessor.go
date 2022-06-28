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

// GsubProcessor type.
type GsubProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Pattern       string               `json:"pattern"`
	Replacement   string               `json:"replacement"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
}

// GsubProcessorBuilder holds GsubProcessor struct and provides a builder API.
type GsubProcessorBuilder struct {
	v *GsubProcessor
}

// NewGsubProcessor provides a builder for the GsubProcessor struct.
func NewGsubProcessor() *GsubProcessorBuilder {
	r := GsubProcessorBuilder{
		&GsubProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the GsubProcessor struct
func (rb *GsubProcessorBuilder) Build() GsubProcessor {
	return *rb.v
}

// Field set the Field property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) Field(field Field) *GsubProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) If_(if_ string) *GsubProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) IgnoreFailure(ignorefailure bool) *GsubProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) IgnoreMissing(ignoremissing bool) *GsubProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// OnFailure set the OnFailure property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *GsubProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Pattern set the Pattern property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) Pattern(pattern string) *GsubProcessorBuilder {
	rb.v.Pattern = pattern
	return rb
}

// Replacement set the Replacement property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) Replacement(replacement string) *GsubProcessorBuilder {
	rb.v.Replacement = replacement
	return rb
}

// Tag set the Tag property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) Tag(tag string) *GsubProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for GsubProcessorBuilder.
func (rb *GsubProcessorBuilder) TargetField(targetfield Field) *GsubProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
