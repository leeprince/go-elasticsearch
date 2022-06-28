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

// JoinProcessor type.
type JoinProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Separator     string               `json:"separator"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   *Field               `json:"target_field,omitempty"`
}

// JoinProcessorBuilder holds JoinProcessor struct and provides a builder API.
type JoinProcessorBuilder struct {
	v *JoinProcessor
}

// NewJoinProcessor provides a builder for the JoinProcessor struct.
func NewJoinProcessor() *JoinProcessorBuilder {
	r := JoinProcessorBuilder{
		&JoinProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the JoinProcessor struct
func (rb *JoinProcessorBuilder) Build() JoinProcessor {
	return *rb.v
}

// Field set the Field property for JoinProcessorBuilder.
func (rb *JoinProcessorBuilder) Field(field Field) *JoinProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for JoinProcessorBuilder.
func (rb *JoinProcessorBuilder) If_(if_ string) *JoinProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for JoinProcessorBuilder.
func (rb *JoinProcessorBuilder) IgnoreFailure(ignorefailure bool) *JoinProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// OnFailure set the OnFailure property for JoinProcessorBuilder.
func (rb *JoinProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *JoinProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Separator set the Separator property for JoinProcessorBuilder.
func (rb *JoinProcessorBuilder) Separator(separator string) *JoinProcessorBuilder {
	rb.v.Separator = separator
	return rb
}

// Tag set the Tag property for JoinProcessorBuilder.
func (rb *JoinProcessorBuilder) Tag(tag string) *JoinProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for JoinProcessorBuilder.
func (rb *JoinProcessorBuilder) TargetField(targetfield Field) *JoinProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
