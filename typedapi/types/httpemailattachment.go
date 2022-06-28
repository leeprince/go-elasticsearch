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

// HttpEmailAttachment type.
type HttpEmailAttachment struct {
	ContentType *string                     `json:"content_type,omitempty"`
	Inline      *bool                       `json:"inline,omitempty"`
	Request     *HttpInputRequestDefinition `json:"request,omitempty"`
}

// HttpEmailAttachmentBuilder holds HttpEmailAttachment struct and provides a builder API.
type HttpEmailAttachmentBuilder struct {
	v *HttpEmailAttachment
}

// NewHttpEmailAttachment provides a builder for the HttpEmailAttachment struct.
func NewHttpEmailAttachment() *HttpEmailAttachmentBuilder {
	r := HttpEmailAttachmentBuilder{
		&HttpEmailAttachment{},
	}

	return &r
}

// Build finalize the chain and returns the HttpEmailAttachment struct
func (rb *HttpEmailAttachmentBuilder) Build() HttpEmailAttachment {
	return *rb.v
}

// ContentType set the ContentType property for HttpEmailAttachmentBuilder.
func (rb *HttpEmailAttachmentBuilder) ContentType(contenttype string) *HttpEmailAttachmentBuilder {
	rb.v.ContentType = &contenttype
	return rb
}

// Inline set the Inline property for HttpEmailAttachmentBuilder.
func (rb *HttpEmailAttachmentBuilder) Inline(inline bool) *HttpEmailAttachmentBuilder {
	rb.v.Inline = &inline
	return rb
}

// Request set the Request property for HttpEmailAttachmentBuilder.
func (rb *HttpEmailAttachmentBuilder) Request(request HttpInputRequestDefinition) *HttpEmailAttachmentBuilder {
	rb.v.Request = &request
	return rb
}
