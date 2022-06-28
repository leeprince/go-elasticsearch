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

// SetSecurityUserProcessor type.
type SetSecurityUserProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Properties    []string             `json:"properties,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// SetSecurityUserProcessorBuilder holds SetSecurityUserProcessor struct and provides a builder API.
type SetSecurityUserProcessorBuilder struct {
	v *SetSecurityUserProcessor
}

// NewSetSecurityUserProcessor provides a builder for the SetSecurityUserProcessor struct.
func NewSetSecurityUserProcessor() *SetSecurityUserProcessorBuilder {
	r := SetSecurityUserProcessorBuilder{
		&SetSecurityUserProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the SetSecurityUserProcessor struct
func (rb *SetSecurityUserProcessorBuilder) Build() SetSecurityUserProcessor {
	return *rb.v
}

// Field set the Field property for SetSecurityUserProcessorBuilder.
func (rb *SetSecurityUserProcessorBuilder) Field(field Field) *SetSecurityUserProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for SetSecurityUserProcessorBuilder.
func (rb *SetSecurityUserProcessorBuilder) If_(if_ string) *SetSecurityUserProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for SetSecurityUserProcessorBuilder.
func (rb *SetSecurityUserProcessorBuilder) IgnoreFailure(ignorefailure bool) *SetSecurityUserProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// OnFailure set the OnFailure property for SetSecurityUserProcessorBuilder.
func (rb *SetSecurityUserProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *SetSecurityUserProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Properties set the Properties property for SetSecurityUserProcessorBuilder.
func (rb *SetSecurityUserProcessorBuilder) Properties(properties ...string) *SetSecurityUserProcessorBuilder {
	rb.v.Properties = append(rb.v.Properties, properties...)
	return rb
}

// Tag set the Tag property for SetSecurityUserProcessorBuilder.
func (rb *SetSecurityUserProcessorBuilder) Tag(tag string) *SetSecurityUserProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
