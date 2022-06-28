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

// SlackMessage type.
type SlackMessage struct {
	Attachments        []SlackAttachment       `json:"attachments"`
	DynamicAttachments *SlackDynamicAttachment `json:"dynamic_attachments,omitempty"`
	From               string                  `json:"from"`
	Icon               *string                 `json:"icon,omitempty"`
	Text               string                  `json:"text"`
	To                 []string                `json:"to"`
}

// SlackMessageBuilder holds SlackMessage struct and provides a builder API.
type SlackMessageBuilder struct {
	v *SlackMessage
}

// NewSlackMessage provides a builder for the SlackMessage struct.
func NewSlackMessage() *SlackMessageBuilder {
	r := SlackMessageBuilder{
		&SlackMessage{},
	}

	return &r
}

// Build finalize the chain and returns the SlackMessage struct
func (rb *SlackMessageBuilder) Build() SlackMessage {
	return *rb.v
}

// Attachments set the Attachments property for SlackMessageBuilder.
func (rb *SlackMessageBuilder) Attachments(attachments ...SlackAttachment) *SlackMessageBuilder {
	rb.v.Attachments = append(rb.v.Attachments, attachments...)
	return rb
}

// DynamicAttachments set the DynamicAttachments property for SlackMessageBuilder.
func (rb *SlackMessageBuilder) DynamicAttachments(dynamicattachments SlackDynamicAttachment) *SlackMessageBuilder {
	rb.v.DynamicAttachments = &dynamicattachments
	return rb
}

// From set the From property for SlackMessageBuilder.
func (rb *SlackMessageBuilder) From(from string) *SlackMessageBuilder {
	rb.v.From = from
	return rb
}

// Icon set the Icon property for SlackMessageBuilder.
func (rb *SlackMessageBuilder) Icon(icon string) *SlackMessageBuilder {
	rb.v.Icon = &icon
	return rb
}

// Text set the Text property for SlackMessageBuilder.
func (rb *SlackMessageBuilder) Text(text string) *SlackMessageBuilder {
	rb.v.Text = text
	return rb
}

// To set the To property for SlackMessageBuilder.
func (rb *SlackMessageBuilder) To(to ...string) *SlackMessageBuilder {
	rb.v.To = append(rb.v.To, to...)
	return rb
}
