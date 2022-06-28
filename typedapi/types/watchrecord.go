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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/executionstatus"
)

// WatchRecord type.
type WatchRecord struct {
	Condition    ConditionContainer              `json:"condition"`
	Input        InputContainer                  `json:"input"`
	Messages     []string                        `json:"messages"`
	Metadata     *Metadata                       `json:"metadata,omitempty"`
	Node         string                          `json:"node"`
	Result       ExecutionResult                 `json:"result"`
	State        executionstatus.ExecutionStatus `json:"state"`
	Status       *WatchStatus                    `json:"status,omitempty"`
	TriggerEvent TriggerEventResult              `json:"trigger_event"`
	User         Username                        `json:"user"`
	WatchId      Id                              `json:"watch_id"`
}

// WatchRecordBuilder holds WatchRecord struct and provides a builder API.
type WatchRecordBuilder struct {
	v *WatchRecord
}

// NewWatchRecord provides a builder for the WatchRecord struct.
func NewWatchRecord() *WatchRecordBuilder {
	r := WatchRecordBuilder{
		&WatchRecord{},
	}

	return &r
}

// Build finalize the chain and returns the WatchRecord struct
func (rb *WatchRecordBuilder) Build() WatchRecord {
	return *rb.v
}

// Condition set the Condition property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) Condition(condition ConditionContainer) *WatchRecordBuilder {
	rb.v.Condition = condition
	return rb
}

// Input set the Input property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) Input(input InputContainer) *WatchRecordBuilder {
	rb.v.Input = input
	return rb
}

// Messages set the Messages property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) Messages(messages ...string) *WatchRecordBuilder {
	rb.v.Messages = append(rb.v.Messages, messages...)
	return rb
}

// Metadata set the Metadata property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) Metadata(metadata Metadata) *WatchRecordBuilder {
	rb.v.Metadata = &metadata
	return rb
}

// Node set the Node property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) Node(node string) *WatchRecordBuilder {
	rb.v.Node = node
	return rb
}

// Result set the Result property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) Result(result ExecutionResult) *WatchRecordBuilder {
	rb.v.Result = result
	return rb
}

// State set the State property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) State(state executionstatus.ExecutionStatus) *WatchRecordBuilder {
	rb.v.State = state
	return rb
}

// Status set the Status property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) Status(status WatchStatus) *WatchRecordBuilder {
	rb.v.Status = &status
	return rb
}

// TriggerEvent set the TriggerEvent property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) TriggerEvent(triggerevent TriggerEventResult) *WatchRecordBuilder {
	rb.v.TriggerEvent = triggerevent
	return rb
}

// User set the User property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) User(user Username) *WatchRecordBuilder {
	rb.v.User = user
	return rb
}

// WatchId set the WatchId property for WatchRecordBuilder.
func (rb *WatchRecordBuilder) WatchId(watchid Id) *WatchRecordBuilder {
	rb.v.WatchId = watchid
	return rb
}
