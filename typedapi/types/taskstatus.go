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

// TaskStatus type.
type TaskStatus struct {
	Batches              int64    `json:"batches"`
	Canceled             *string  `json:"canceled,omitempty"`
	Created              int64    `json:"created"`
	Deleted              int64    `json:"deleted"`
	Failures             []string `json:"failures,omitempty"`
	Noops                int64    `json:"noops"`
	RequestsPerSecond    float32  `json:"requests_per_second"`
	Retries              Retries  `json:"retries"`
	Throttled            *Time    `json:"throttled,omitempty"`
	ThrottledMillis      int64    `json:"throttled_millis"`
	ThrottledUntil       *Time    `json:"throttled_until,omitempty"`
	ThrottledUntilMillis int64    `json:"throttled_until_millis"`
	TimedOut             *bool    `json:"timed_out,omitempty"`
	Took                 *int64   `json:"took,omitempty"`
	Total                int64    `json:"total"`
	Updated              int64    `json:"updated"`
	VersionConflicts     int64    `json:"version_conflicts"`
}

// TaskStatusBuilder holds TaskStatus struct and provides a builder API.
type TaskStatusBuilder struct {
	v *TaskStatus
}

// NewTaskStatus provides a builder for the TaskStatus struct.
func NewTaskStatus() *TaskStatusBuilder {
	r := TaskStatusBuilder{
		&TaskStatus{},
	}

	return &r
}

// Build finalize the chain and returns the TaskStatus struct
func (rb *TaskStatusBuilder) Build() TaskStatus {
	return *rb.v
}

// Batches set the Batches property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Batches(batches int64) *TaskStatusBuilder {
	rb.v.Batches = batches
	return rb
}

// Canceled set the Canceled property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Canceled(canceled string) *TaskStatusBuilder {
	rb.v.Canceled = &canceled
	return rb
}

// Created set the Created property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Created(created int64) *TaskStatusBuilder {
	rb.v.Created = created
	return rb
}

// Deleted set the Deleted property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Deleted(deleted int64) *TaskStatusBuilder {
	rb.v.Deleted = deleted
	return rb
}

// Failures set the Failures property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Failures(failures ...string) *TaskStatusBuilder {
	rb.v.Failures = append(rb.v.Failures, failures...)
	return rb
}

// Noops set the Noops property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Noops(noops int64) *TaskStatusBuilder {
	rb.v.Noops = noops
	return rb
}

// RequestsPerSecond set the RequestsPerSecond property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) RequestsPerSecond(requestspersecond float32) *TaskStatusBuilder {
	rb.v.RequestsPerSecond = requestspersecond
	return rb
}

// Retries set the Retries property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Retries(retries Retries) *TaskStatusBuilder {
	rb.v.Retries = retries
	return rb
}

// Throttled set the Throttled property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Throttled(throttled Time) *TaskStatusBuilder {
	rb.v.Throttled = &throttled
	return rb
}

// ThrottledMillis set the ThrottledMillis property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) ThrottledMillis(throttledmillis int64) *TaskStatusBuilder {
	rb.v.ThrottledMillis = throttledmillis
	return rb
}

// ThrottledUntil set the ThrottledUntil property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) ThrottledUntil(throttleduntil Time) *TaskStatusBuilder {
	rb.v.ThrottledUntil = &throttleduntil
	return rb
}

// ThrottledUntilMillis set the ThrottledUntilMillis property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) ThrottledUntilMillis(throttleduntilmillis int64) *TaskStatusBuilder {
	rb.v.ThrottledUntilMillis = throttleduntilmillis
	return rb
}

// TimedOut set the TimedOut property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) TimedOut(timedout bool) *TaskStatusBuilder {
	rb.v.TimedOut = &timedout
	return rb
}

// Took set the Took property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Took(took int64) *TaskStatusBuilder {
	rb.v.Took = &took
	return rb
}

// Total set the Total property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Total(total int64) *TaskStatusBuilder {
	rb.v.Total = total
	return rb
}

// Updated set the Updated property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) Updated(updated int64) *TaskStatusBuilder {
	rb.v.Updated = updated
	return rb
}

// VersionConflicts set the VersionConflicts property for TaskStatusBuilder.
func (rb *TaskStatusBuilder) VersionConflicts(versionconflicts int64) *TaskStatusBuilder {
	rb.v.VersionConflicts = versionconflicts
	return rb
}
