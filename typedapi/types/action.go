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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/actiontype"
)

// Action type.
type Action struct {
	ActionType             *actiontype.ActionType `json:"action_type,omitempty"`
	Condition              *ConditionContainer    `json:"condition,omitempty"`
	Email                  *EmailAction           `json:"email,omitempty"`
	Foreach                *string                `json:"foreach,omitempty"`
	Index                  *IndexAction           `json:"index,omitempty"`
	Logging                *LoggingAction         `json:"logging,omitempty"`
	MaxIterations          *int                   `json:"max_iterations,omitempty"`
	Name                   *Name                  `json:"name,omitempty"`
	Pagerduty              *PagerDutyAction       `json:"pagerduty,omitempty"`
	Slack                  *SlackAction           `json:"slack,omitempty"`
	ThrottlePeriod         *Time                  `json:"throttle_period,omitempty"`
	ThrottlePeriodInMillis *EpochMillis           `json:"throttle_period_in_millis,omitempty"`
	Transform              *TransformContainer    `json:"transform,omitempty"`
	Webhook                *WebhookAction         `json:"webhook,omitempty"`
}

// ActionBuilder holds Action struct and provides a builder API.
type ActionBuilder struct {
	v *Action
}

// NewAction provides a builder for the Action struct.
func NewAction() *ActionBuilder {
	r := ActionBuilder{
		&Action{},
	}

	return &r
}

// Build finalize the chain and returns the Action struct
func (rb *ActionBuilder) Build() Action {
	return *rb.v
}

// ActionType set the ActionType property for ActionBuilder.
func (rb *ActionBuilder) ActionType(actiontype actiontype.ActionType) *ActionBuilder {
	rb.v.ActionType = &actiontype
	return rb
}

// Condition set the Condition property for ActionBuilder.
func (rb *ActionBuilder) Condition(condition ConditionContainer) *ActionBuilder {
	rb.v.Condition = &condition
	return rb
}

// Email set the Email property for ActionBuilder.
func (rb *ActionBuilder) Email(email EmailAction) *ActionBuilder {
	rb.v.Email = &email
	return rb
}

// Foreach set the Foreach property for ActionBuilder.
func (rb *ActionBuilder) Foreach(foreach string) *ActionBuilder {
	rb.v.Foreach = &foreach
	return rb
}

// Index set the Index property for ActionBuilder.
func (rb *ActionBuilder) Index(index IndexAction) *ActionBuilder {
	rb.v.Index = &index
	return rb
}

// Logging set the Logging property for ActionBuilder.
func (rb *ActionBuilder) Logging(logging LoggingAction) *ActionBuilder {
	rb.v.Logging = &logging
	return rb
}

// MaxIterations set the MaxIterations property for ActionBuilder.
func (rb *ActionBuilder) MaxIterations(maxiterations int) *ActionBuilder {
	rb.v.MaxIterations = &maxiterations
	return rb
}

// Name set the Name property for ActionBuilder.
func (rb *ActionBuilder) Name(name Name) *ActionBuilder {
	rb.v.Name = &name
	return rb
}

// Pagerduty set the Pagerduty property for ActionBuilder.
func (rb *ActionBuilder) Pagerduty(pagerduty PagerDutyAction) *ActionBuilder {
	rb.v.Pagerduty = &pagerduty
	return rb
}

// Slack set the Slack property for ActionBuilder.
func (rb *ActionBuilder) Slack(slack SlackAction) *ActionBuilder {
	rb.v.Slack = &slack
	return rb
}

// ThrottlePeriod set the ThrottlePeriod property for ActionBuilder.
func (rb *ActionBuilder) ThrottlePeriod(throttleperiod Time) *ActionBuilder {
	rb.v.ThrottlePeriod = &throttleperiod
	return rb
}

// ThrottlePeriodInMillis set the ThrottlePeriodInMillis property for ActionBuilder.
func (rb *ActionBuilder) ThrottlePeriodInMillis(throttleperiodinmillis EpochMillis) *ActionBuilder {
	rb.v.ThrottlePeriodInMillis = &throttleperiodinmillis
	return rb
}

// Transform set the Transform property for ActionBuilder.
func (rb *ActionBuilder) Transform(transform TransformContainer) *ActionBuilder {
	rb.v.Transform = &transform
	return rb
}

// Webhook set the Webhook property for ActionBuilder.
func (rb *ActionBuilder) Webhook(webhook WebhookAction) *ActionBuilder {
	rb.v.Webhook = &webhook
	return rb
}
