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

// EmailAttachmentContainer type.
type EmailAttachmentContainer struct {
	Data      *DataEmailAttachment      `json:"data,omitempty"`
	Http      *HttpEmailAttachment      `json:"http,omitempty"`
	Reporting *ReportingEmailAttachment `json:"reporting,omitempty"`
}

// EmailAttachmentContainerBuilder holds EmailAttachmentContainer struct and provides a builder API.
type EmailAttachmentContainerBuilder struct {
	v *EmailAttachmentContainer
}

// NewEmailAttachmentContainer provides a builder for the EmailAttachmentContainer struct.
func NewEmailAttachmentContainer() *EmailAttachmentContainerBuilder {
	r := EmailAttachmentContainerBuilder{
		&EmailAttachmentContainer{},
	}

	return &r
}

// Build finalize the chain and returns the EmailAttachmentContainer struct
func (rb *EmailAttachmentContainerBuilder) Build() EmailAttachmentContainer {
	return *rb.v
}

// Data set the Data property for EmailAttachmentContainerBuilder.
func (rb *EmailAttachmentContainerBuilder) Data(data DataEmailAttachment) *EmailAttachmentContainerBuilder {
	rb.v.Data = &data
	return rb
}

// Http set the Http property for EmailAttachmentContainerBuilder.
func (rb *EmailAttachmentContainerBuilder) Http(http HttpEmailAttachment) *EmailAttachmentContainerBuilder {
	rb.v.Http = &http
	return rb
}

// Reporting set the Reporting property for EmailAttachmentContainerBuilder.
func (rb *EmailAttachmentContainerBuilder) Reporting(reporting ReportingEmailAttachment) *EmailAttachmentContainerBuilder {
	rb.v.Reporting = &reporting
	return rb
}
