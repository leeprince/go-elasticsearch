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

// DissectProcessor type.
type DissectProcessor struct {
	AppendSeparator string               `json:"append_separator"`
	Field           Field                `json:"field"`
	If              *string              `json:"if,omitempty"`
	IgnoreFailure   *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing   bool                 `json:"ignore_missing"`
	OnFailure       []ProcessorContainer `json:"on_failure,omitempty"`
	Pattern         string               `json:"pattern"`
	Tag             *string              `json:"tag,omitempty"`
}

// DissectProcessorBuilder holds DissectProcessor struct and provides a builder API.
type DissectProcessorBuilder struct {
	v *DissectProcessor
}

// NewDissectProcessor provides a builder for the DissectProcessor struct.
func NewDissectProcessor() *DissectProcessorBuilder {
	r := DissectProcessorBuilder{
		&DissectProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the DissectProcessor struct
func (rb *DissectProcessorBuilder) Build() DissectProcessor {
	return *rb.v
}

// AppendSeparator set the AppendSeparator property for DissectProcessorBuilder.
func (rb *DissectProcessorBuilder) AppendSeparator(appendseparator string) *DissectProcessorBuilder {
	rb.v.AppendSeparator = appendseparator
	return rb
}

// Field set the Field property for DissectProcessorBuilder.
func (rb *DissectProcessorBuilder) Field(field Field) *DissectProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for DissectProcessorBuilder.
func (rb *DissectProcessorBuilder) If_(if_ string) *DissectProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for DissectProcessorBuilder.
func (rb *DissectProcessorBuilder) IgnoreFailure(ignorefailure bool) *DissectProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for DissectProcessorBuilder.
func (rb *DissectProcessorBuilder) IgnoreMissing(ignoremissing bool) *DissectProcessorBuilder {
	rb.v.IgnoreMissing = ignoremissing
	return rb
}

// OnFailure set the OnFailure property for DissectProcessorBuilder.
func (rb *DissectProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *DissectProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Pattern set the Pattern property for DissectProcessorBuilder.
func (rb *DissectProcessorBuilder) Pattern(pattern string) *DissectProcessorBuilder {
	rb.v.Pattern = pattern
	return rb
}

// Tag set the Tag property for DissectProcessorBuilder.
func (rb *DissectProcessorBuilder) Tag(tag string) *DissectProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
