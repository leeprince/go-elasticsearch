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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actionstatusoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actiontype"
)

// ExecutionResultAction type.
type ExecutionResultAction struct {
	Email     *EmailResult                            `json:"email,omitempty"`
	Error     *ErrorCause                             `json:"error,omitempty"`
	Id        Id                                      `json:"id"`
	Index     *IndexResult                            `json:"index,omitempty"`
	Logging   *LoggingResult                          `json:"logging,omitempty"`
	Pagerduty *PagerDutyResult                        `json:"pagerduty,omitempty"`
	Reason    *string                                 `json:"reason,omitempty"`
	Slack     *SlackResult                            `json:"slack,omitempty"`
	Status    actionstatusoptions.ActionStatusOptions `json:"status"`
	Type      actiontype.ActionType                   `json:"type"`
	Webhook   *WebhookResult                          `json:"webhook,omitempty"`
}

// ExecutionResultActionBuilder holds ExecutionResultAction struct and provides a builder API.
type ExecutionResultActionBuilder struct {
	v *ExecutionResultAction
}

// NewExecutionResultAction provides a builder for the ExecutionResultAction struct.
func NewExecutionResultAction() *ExecutionResultActionBuilder {
	r := ExecutionResultActionBuilder{
		&ExecutionResultAction{},
	}

	return &r
}

// Build finalize the chain and returns the ExecutionResultAction struct
func (rb *ExecutionResultActionBuilder) Build() ExecutionResultAction {
	return *rb.v
}

// Email set the Email property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Email(email EmailResult) *ExecutionResultActionBuilder {
	rb.v.Email = &email
	return rb
}

// Error set the Error property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Error(error ErrorCause) *ExecutionResultActionBuilder {
	rb.v.Error = &error
	return rb
}

// Id set the Id property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Id(id Id) *ExecutionResultActionBuilder {
	rb.v.Id = id
	return rb
}

// Index set the Index property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Index(index IndexResult) *ExecutionResultActionBuilder {
	rb.v.Index = &index
	return rb
}

// Logging set the Logging property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Logging(logging LoggingResult) *ExecutionResultActionBuilder {
	rb.v.Logging = &logging
	return rb
}

// Pagerduty set the Pagerduty property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Pagerduty(pagerduty PagerDutyResult) *ExecutionResultActionBuilder {
	rb.v.Pagerduty = &pagerduty
	return rb
}

// Reason set the Reason property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Reason(reason string) *ExecutionResultActionBuilder {
	rb.v.Reason = &reason
	return rb
}

// Slack set the Slack property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Slack(slack SlackResult) *ExecutionResultActionBuilder {
	rb.v.Slack = &slack
	return rb
}

// Status set the Status property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Status(status actionstatusoptions.ActionStatusOptions) *ExecutionResultActionBuilder {
	rb.v.Status = status
	return rb
}

// Type set the Type property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Type_(type_ actiontype.ActionType) *ExecutionResultActionBuilder {
	rb.v.Type = type_
	return rb
}

// Webhook set the Webhook property for ExecutionResultActionBuilder.
func (rb *ExecutionResultActionBuilder) Webhook(webhook WebhookResult) *ExecutionResultActionBuilder {
	rb.v.Webhook = &webhook
	return rb
}
