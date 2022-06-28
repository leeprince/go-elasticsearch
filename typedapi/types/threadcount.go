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

// ThreadCount type.
type ThreadCount struct {
	Active    *int64 `json:"active,omitempty"`
	Completed *int64 `json:"completed,omitempty"`
	Largest   *int64 `json:"largest,omitempty"`
	Queue     *int64 `json:"queue,omitempty"`
	Rejected  *int64 `json:"rejected,omitempty"`
	Threads   *int64 `json:"threads,omitempty"`
}

// ThreadCountBuilder holds ThreadCount struct and provides a builder API.
type ThreadCountBuilder struct {
	v *ThreadCount
}

// NewThreadCount provides a builder for the ThreadCount struct.
func NewThreadCount() *ThreadCountBuilder {
	r := ThreadCountBuilder{
		&ThreadCount{},
	}

	return &r
}

// Build finalize the chain and returns the ThreadCount struct
func (rb *ThreadCountBuilder) Build() ThreadCount {
	return *rb.v
}

// Active set the Active property for ThreadCountBuilder.
func (rb *ThreadCountBuilder) Active(active int64) *ThreadCountBuilder {
	rb.v.Active = &active
	return rb
}

// Completed set the Completed property for ThreadCountBuilder.
func (rb *ThreadCountBuilder) Completed(completed int64) *ThreadCountBuilder {
	rb.v.Completed = &completed
	return rb
}

// Largest set the Largest property for ThreadCountBuilder.
func (rb *ThreadCountBuilder) Largest(largest int64) *ThreadCountBuilder {
	rb.v.Largest = &largest
	return rb
}

// Queue set the Queue property for ThreadCountBuilder.
func (rb *ThreadCountBuilder) Queue(queue int64) *ThreadCountBuilder {
	rb.v.Queue = &queue
	return rb
}

// Rejected set the Rejected property for ThreadCountBuilder.
func (rb *ThreadCountBuilder) Rejected(rejected int64) *ThreadCountBuilder {
	rb.v.Rejected = &rejected
	return rb
}

// Threads set the Threads property for ThreadCountBuilder.
func (rb *ThreadCountBuilder) Threads(threads int64) *ThreadCountBuilder {
	rb.v.Threads = &threads
	return rb
}
