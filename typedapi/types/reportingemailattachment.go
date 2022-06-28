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

// ReportingEmailAttachment type.
type ReportingEmailAttachment struct {
	Inline   *bool                       `json:"inline,omitempty"`
	Interval *Time                       `json:"interval,omitempty"`
	Request  *HttpInputRequestDefinition `json:"request,omitempty"`
	Retries  *int                        `json:"retries,omitempty"`
	Url      string                      `json:"url"`
}

// ReportingEmailAttachmentBuilder holds ReportingEmailAttachment struct and provides a builder API.
type ReportingEmailAttachmentBuilder struct {
	v *ReportingEmailAttachment
}

// NewReportingEmailAttachment provides a builder for the ReportingEmailAttachment struct.
func NewReportingEmailAttachment() *ReportingEmailAttachmentBuilder {
	r := ReportingEmailAttachmentBuilder{
		&ReportingEmailAttachment{},
	}

	return &r
}

// Build finalize the chain and returns the ReportingEmailAttachment struct
func (rb *ReportingEmailAttachmentBuilder) Build() ReportingEmailAttachment {
	return *rb.v
}

// Inline set the Inline property for ReportingEmailAttachmentBuilder.
func (rb *ReportingEmailAttachmentBuilder) Inline(inline bool) *ReportingEmailAttachmentBuilder {
	rb.v.Inline = &inline
	return rb
}

// Interval set the Interval property for ReportingEmailAttachmentBuilder.
func (rb *ReportingEmailAttachmentBuilder) Interval(interval Time) *ReportingEmailAttachmentBuilder {
	rb.v.Interval = &interval
	return rb
}

// Request set the Request property for ReportingEmailAttachmentBuilder.
func (rb *ReportingEmailAttachmentBuilder) Request(request HttpInputRequestDefinition) *ReportingEmailAttachmentBuilder {
	rb.v.Request = &request
	return rb
}

// Retries set the Retries property for ReportingEmailAttachmentBuilder.
func (rb *ReportingEmailAttachmentBuilder) Retries(retries int) *ReportingEmailAttachmentBuilder {
	rb.v.Retries = &retries
	return rb
}

// Url set the Url property for ReportingEmailAttachmentBuilder.
func (rb *ReportingEmailAttachmentBuilder) Url(url string) *ReportingEmailAttachmentBuilder {
	rb.v.Url = url
	return rb
}
