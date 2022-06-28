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

// Watcher type.
type Watcher struct {
	Available bool           `json:"available"`
	Count     Counter        `json:"count"`
	Enabled   bool           `json:"enabled"`
	Execution WatcherActions `json:"execution"`
	Watch     WatcherWatch   `json:"watch"`
}

// WatcherBuilder holds Watcher struct and provides a builder API.
type WatcherBuilder struct {
	v *Watcher
}

// NewWatcher provides a builder for the Watcher struct.
func NewWatcher() *WatcherBuilder {
	r := WatcherBuilder{
		&Watcher{},
	}

	return &r
}

// Build finalize the chain and returns the Watcher struct
func (rb *WatcherBuilder) Build() Watcher {
	return *rb.v
}

// Available set the Available property for WatcherBuilder.
func (rb *WatcherBuilder) Available(available bool) *WatcherBuilder {
	rb.v.Available = available
	return rb
}

// Count set the Count property for WatcherBuilder.
func (rb *WatcherBuilder) Count(count Counter) *WatcherBuilder {
	rb.v.Count = count
	return rb
}

// Enabled set the Enabled property for WatcherBuilder.
func (rb *WatcherBuilder) Enabled(enabled bool) *WatcherBuilder {
	rb.v.Enabled = enabled
	return rb
}

// Execution set the Execution property for WatcherBuilder.
func (rb *WatcherBuilder) Execution(execution WatcherActions) *WatcherBuilder {
	rb.v.Execution = execution
	return rb
}

// Watch set the Watch property for WatcherBuilder.
func (rb *WatcherBuilder) Watch(watch WatcherWatch) *WatcherBuilder {
	rb.v.Watch = watch
	return rb
}
