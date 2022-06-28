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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/emailpriority"
)

// EmailAction type.
type EmailAction struct {
	Attachments map[string]EmailAttachmentContainer `json:"attachments,omitempty"`
	Bcc         []string                            `json:"bcc,omitempty"`
	Body        *EmailBody                          `json:"body,omitempty"`
	Cc          []string                            `json:"cc,omitempty"`
	From        *string                             `json:"from,omitempty"`
	Id          *Id                                 `json:"id,omitempty"`
	Priority    *emailpriority.EmailPriority        `json:"priority,omitempty"`
	ReplyTo     []string                            `json:"reply_to,omitempty"`
	SentDate    *DateString                         `json:"sent_date,omitempty"`
	Subject     string                              `json:"subject"`
	To          []string                            `json:"to"`
}

// EmailActionBuilder holds EmailAction struct and provides a builder API.
type EmailActionBuilder struct {
	v *EmailAction
}

// NewEmailAction provides a builder for the EmailAction struct.
func NewEmailAction() *EmailActionBuilder {
	r := EmailActionBuilder{
		&EmailAction{
			Attachments: make(map[string]EmailAttachmentContainer, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the EmailAction struct
func (rb *EmailActionBuilder) Build() EmailAction {
	return *rb.v
}

// Attachments set the Attachments property for EmailActionBuilder.
func (rb *EmailActionBuilder) Attachments(value map[string]EmailAttachmentContainer) *EmailActionBuilder {
	rb.v.Attachments = value
	return rb
}

// Bcc set the Bcc property for EmailActionBuilder.
func (rb *EmailActionBuilder) Bcc(bcc ...string) *EmailActionBuilder {
	rb.v.Bcc = append(rb.v.Bcc, bcc...)
	return rb
}

// Body set the Body property for EmailActionBuilder.
func (rb *EmailActionBuilder) Body(body EmailBody) *EmailActionBuilder {
	rb.v.Body = &body
	return rb
}

// Cc set the Cc property for EmailActionBuilder.
func (rb *EmailActionBuilder) Cc(cc ...string) *EmailActionBuilder {
	rb.v.Cc = append(rb.v.Cc, cc...)
	return rb
}

// From set the From property for EmailActionBuilder.
func (rb *EmailActionBuilder) From(from string) *EmailActionBuilder {
	rb.v.From = &from
	return rb
}

// Id set the Id property for EmailActionBuilder.
func (rb *EmailActionBuilder) Id(id Id) *EmailActionBuilder {
	rb.v.Id = &id
	return rb
}

// Priority set the Priority property for EmailActionBuilder.
func (rb *EmailActionBuilder) Priority(priority emailpriority.EmailPriority) *EmailActionBuilder {
	rb.v.Priority = &priority
	return rb
}

// ReplyTo set the ReplyTo property for EmailActionBuilder.
func (rb *EmailActionBuilder) ReplyTo(reply_to ...string) *EmailActionBuilder {
	rb.v.ReplyTo = append(rb.v.ReplyTo, reply_to...)
	return rb
}

// SentDate set the SentDate property for EmailActionBuilder.
func (rb *EmailActionBuilder) SentDate(sentdate DateString) *EmailActionBuilder {
	rb.v.SentDate = &sentdate
	return rb
}

// Subject set the Subject property for EmailActionBuilder.
func (rb *EmailActionBuilder) Subject(subject string) *EmailActionBuilder {
	rb.v.Subject = subject
	return rb
}

// To set the To property for EmailActionBuilder.
func (rb *EmailActionBuilder) To(to ...string) *EmailActionBuilder {
	rb.v.To = append(rb.v.To, to...)
	return rb
}
