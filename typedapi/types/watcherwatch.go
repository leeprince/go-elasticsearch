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

// WatcherWatch type.
type WatcherWatch struct {
	Action    map[Name]Counter    `json:"action,omitempty"`
	Condition map[Name]Counter    `json:"condition,omitempty"`
	Input     map[Name]Counter    `json:"input"`
	Trigger   WatcherWatchTrigger `json:"trigger"`
}

// WatcherWatchBuilder holds WatcherWatch struct and provides a builder API.
type WatcherWatchBuilder struct {
	v *WatcherWatch
}

// NewWatcherWatch provides a builder for the WatcherWatch struct.
func NewWatcherWatch() *WatcherWatchBuilder {
	r := WatcherWatchBuilder{
		&WatcherWatch{
			Action:    make(map[Name]Counter, 0),
			Condition: make(map[Name]Counter, 0),
			Input:     make(map[Name]Counter, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the WatcherWatch struct
func (rb *WatcherWatchBuilder) Build() WatcherWatch {
	return *rb.v
}

// Action set the Action property for WatcherWatchBuilder.
func (rb *WatcherWatchBuilder) Action(value map[Name]Counter) *WatcherWatchBuilder {
	rb.v.Action = value
	return rb
}

// Condition set the Condition property for WatcherWatchBuilder.
func (rb *WatcherWatchBuilder) Condition(value map[Name]Counter) *WatcherWatchBuilder {
	rb.v.Condition = value
	return rb
}

// Input set the Input property for WatcherWatchBuilder.
func (rb *WatcherWatchBuilder) Input(value map[Name]Counter) *WatcherWatchBuilder {
	rb.v.Input = value
	return rb
}

// Trigger set the Trigger property for WatcherWatchBuilder.
func (rb *WatcherWatchBuilder) Trigger(trigger WatcherWatchTrigger) *WatcherWatchBuilder {
	rb.v.Trigger = trigger
	return rb
}
