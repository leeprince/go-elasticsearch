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

// AppendProcessor type.
type AppendProcessor struct {
	AllowDuplicates *bool                `json:"allow_duplicates,omitempty"`
	Field           Field                `json:"field"`
	If              *string              `json:"if,omitempty"`
	IgnoreFailure   *bool                `json:"ignore_failure,omitempty"`
	OnFailure       []ProcessorContainer `json:"on_failure,omitempty"`
	Tag             *string              `json:"tag,omitempty"`
	Value           []interface{}        `json:"value"`
}

// AppendProcessorBuilder holds AppendProcessor struct and provides a builder API.
type AppendProcessorBuilder struct {
	v *AppendProcessor
}

// NewAppendProcessor provides a builder for the AppendProcessor struct.
func NewAppendProcessor() *AppendProcessorBuilder {
	r := AppendProcessorBuilder{
		&AppendProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the AppendProcessor struct
func (rb *AppendProcessorBuilder) Build() AppendProcessor {
	return *rb.v
}

// AllowDuplicates set the AllowDuplicates property for AppendProcessorBuilder.
func (rb *AppendProcessorBuilder) AllowDuplicates(allowduplicates bool) *AppendProcessorBuilder {
	rb.v.AllowDuplicates = &allowduplicates
	return rb
}

// Field set the Field property for AppendProcessorBuilder.
func (rb *AppendProcessorBuilder) Field(field Field) *AppendProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for AppendProcessorBuilder.
func (rb *AppendProcessorBuilder) If_(if_ string) *AppendProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for AppendProcessorBuilder.
func (rb *AppendProcessorBuilder) IgnoreFailure(ignorefailure bool) *AppendProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// OnFailure set the OnFailure property for AppendProcessorBuilder.
func (rb *AppendProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *AppendProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for AppendProcessorBuilder.
func (rb *AppendProcessorBuilder) Tag(tag string) *AppendProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// Value set the Value property for AppendProcessorBuilder.
func (rb *AppendProcessorBuilder) Value(value ...interface{}) *AppendProcessorBuilder {
	rb.v.Value = append(rb.v.Value, value...)
	return rb
}
