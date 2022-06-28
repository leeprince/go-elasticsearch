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

// Watch type.
type Watch struct {
	Actions                map[IndexName]Action `json:"actions"`
	Condition              ConditionContainer   `json:"condition"`
	Input                  InputContainer       `json:"input"`
	Metadata               *Metadata            `json:"metadata,omitempty"`
	Status                 *WatchStatus         `json:"status,omitempty"`
	ThrottlePeriod         *string              `json:"throttle_period,omitempty"`
	ThrottlePeriodInMillis *int64               `json:"throttle_period_in_millis,omitempty"`
	Transform              *TransformContainer  `json:"transform,omitempty"`
	Trigger                TriggerContainer     `json:"trigger"`
}

// WatchBuilder holds Watch struct and provides a builder API.
type WatchBuilder struct {
	v *Watch
}

// NewWatch provides a builder for the Watch struct.
func NewWatch() *WatchBuilder {
	r := WatchBuilder{
		&Watch{
			Actions: make(map[IndexName]Action, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Watch struct
func (rb *WatchBuilder) Build() Watch {
	return *rb.v
}

// Actions set the Actions property for WatchBuilder.
func (rb *WatchBuilder) Actions(value map[IndexName]Action) *WatchBuilder {
	rb.v.Actions = value
	return rb
}

// Condition set the Condition property for WatchBuilder.
func (rb *WatchBuilder) Condition(condition ConditionContainer) *WatchBuilder {
	rb.v.Condition = condition
	return rb
}

// Input set the Input property for WatchBuilder.
func (rb *WatchBuilder) Input(input InputContainer) *WatchBuilder {
	rb.v.Input = input
	return rb
}

// Metadata set the Metadata property for WatchBuilder.
func (rb *WatchBuilder) Metadata(metadata Metadata) *WatchBuilder {
	rb.v.Metadata = &metadata
	return rb
}

// Status set the Status property for WatchBuilder.
func (rb *WatchBuilder) Status(status WatchStatus) *WatchBuilder {
	rb.v.Status = &status
	return rb
}

// ThrottlePeriod set the ThrottlePeriod property for WatchBuilder.
func (rb *WatchBuilder) ThrottlePeriod(throttleperiod string) *WatchBuilder {
	rb.v.ThrottlePeriod = &throttleperiod
	return rb
}

// ThrottlePeriodInMillis set the ThrottlePeriodInMillis property for WatchBuilder.
func (rb *WatchBuilder) ThrottlePeriodInMillis(throttleperiodinmillis int64) *WatchBuilder {
	rb.v.ThrottlePeriodInMillis = &throttleperiodinmillis
	return rb
}

// Transform set the Transform property for WatchBuilder.
func (rb *WatchBuilder) Transform(transform TransformContainer) *WatchBuilder {
	rb.v.Transform = &transform
	return rb
}

// Trigger set the Trigger property for WatchBuilder.
func (rb *WatchBuilder) Trigger(trigger TriggerContainer) *WatchBuilder {
	rb.v.Trigger = trigger
	return rb
}
