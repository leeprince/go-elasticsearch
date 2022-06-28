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

// WatcherWatchTrigger type.
type WatcherWatchTrigger struct {
	All_     Counter                      `json:"_all"`
	Schedule *WatcherWatchTriggerSchedule `json:"schedule,omitempty"`
}

// WatcherWatchTriggerBuilder holds WatcherWatchTrigger struct and provides a builder API.
type WatcherWatchTriggerBuilder struct {
	v *WatcherWatchTrigger
}

// NewWatcherWatchTrigger provides a builder for the WatcherWatchTrigger struct.
func NewWatcherWatchTrigger() *WatcherWatchTriggerBuilder {
	r := WatcherWatchTriggerBuilder{
		&WatcherWatchTrigger{},
	}

	return &r
}

// Build finalize the chain and returns the WatcherWatchTrigger struct
func (rb *WatcherWatchTriggerBuilder) Build() WatcherWatchTrigger {
	return *rb.v
}

// All_ set the All_ property for WatcherWatchTriggerBuilder.
func (rb *WatcherWatchTriggerBuilder) All_(all_ Counter) *WatcherWatchTriggerBuilder {
	rb.v.All_ = all_
	return rb
}

// Schedule set the Schedule property for WatcherWatchTriggerBuilder.
func (rb *WatcherWatchTriggerBuilder) Schedule(schedule WatcherWatchTriggerSchedule) *WatcherWatchTriggerBuilder {
	rb.v.Schedule = &schedule
	return rb
}
