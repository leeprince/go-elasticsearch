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

// AttachmentProcessor type.
type AttachmentProcessor struct {
	Field             Field                `json:"field"`
	If                *string              `json:"if,omitempty"`
	IgnoreFailure     *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing     *bool                `json:"ignore_missing,omitempty"`
	IndexedChars      *int64               `json:"indexed_chars,omitempty"`
	IndexedCharsField *Field               `json:"indexed_chars_field,omitempty"`
	OnFailure         []ProcessorContainer `json:"on_failure,omitempty"`
	Properties        []string             `json:"properties,omitempty"`
	ResourceName      *string              `json:"resource_name,omitempty"`
	Tag               *string              `json:"tag,omitempty"`
	TargetField       *Field               `json:"target_field,omitempty"`
}

// AttachmentProcessorBuilder holds AttachmentProcessor struct and provides a builder API.
type AttachmentProcessorBuilder struct {
	v *AttachmentProcessor
}

// NewAttachmentProcessor provides a builder for the AttachmentProcessor struct.
func NewAttachmentProcessor() *AttachmentProcessorBuilder {
	r := AttachmentProcessorBuilder{
		&AttachmentProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the AttachmentProcessor struct
func (rb *AttachmentProcessorBuilder) Build() AttachmentProcessor {
	return *rb.v
}

// Field set the Field property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) Field(field Field) *AttachmentProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) If_(if_ string) *AttachmentProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) IgnoreFailure(ignorefailure bool) *AttachmentProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) IgnoreMissing(ignoremissing bool) *AttachmentProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// IndexedChars set the IndexedChars property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) IndexedChars(indexedchars int64) *AttachmentProcessorBuilder {
	rb.v.IndexedChars = &indexedchars
	return rb
}

// IndexedCharsField set the IndexedCharsField property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) IndexedCharsField(indexedcharsfield Field) *AttachmentProcessorBuilder {
	rb.v.IndexedCharsField = &indexedcharsfield
	return rb
}

// OnFailure set the OnFailure property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *AttachmentProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Properties set the Properties property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) Properties(properties ...string) *AttachmentProcessorBuilder {
	rb.v.Properties = append(rb.v.Properties, properties...)
	return rb
}

// ResourceName set the ResourceName property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) ResourceName(resourcename string) *AttachmentProcessorBuilder {
	rb.v.ResourceName = &resourcename
	return rb
}

// Tag set the Tag property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) Tag(tag string) *AttachmentProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for AttachmentProcessorBuilder.
func (rb *AttachmentProcessorBuilder) TargetField(targetfield Field) *AttachmentProcessorBuilder {
	rb.v.TargetField = &targetfield
	return rb
}
