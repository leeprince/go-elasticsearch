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

// SetProcessor type.
type SetProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Override      *bool                `json:"override,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	Value         interface{}          `json:"value,omitempty"`
}

// SetProcessorBuilder holds SetProcessor struct and provides a builder API.
type SetProcessorBuilder struct {
	v *SetProcessor
}

// NewSetProcessor provides a builder for the SetProcessor struct.
func NewSetProcessor() *SetProcessorBuilder {
	r := SetProcessorBuilder{
		&SetProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the SetProcessor struct
func (rb *SetProcessorBuilder) Build() SetProcessor {
	return *rb.v
}

// Field set the Field property for SetProcessorBuilder.
func (rb *SetProcessorBuilder) Field(field Field) *SetProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for SetProcessorBuilder.
func (rb *SetProcessorBuilder) If_(if_ string) *SetProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for SetProcessorBuilder.
func (rb *SetProcessorBuilder) IgnoreFailure(ignorefailure bool) *SetProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// OnFailure set the OnFailure property for SetProcessorBuilder.
func (rb *SetProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *SetProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Override set the Override property for SetProcessorBuilder.
func (rb *SetProcessorBuilder) Override(override bool) *SetProcessorBuilder {
	rb.v.Override = &override
	return rb
}

// Tag set the Tag property for SetProcessorBuilder.
func (rb *SetProcessorBuilder) Tag(tag string) *SetProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// Value set the Value property for SetProcessorBuilder.
func (rb *SetProcessorBuilder) Value(value interface{}) *SetProcessorBuilder {
	rb.v.Value = value
	return rb
}
