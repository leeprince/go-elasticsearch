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

// ReindexTask type.
type ReindexTask struct {
	Action             string        `json:"action"`
	Cancellable        bool          `json:"cancellable"`
	Description        string        `json:"description"`
	Headers            HttpHeaders   `json:"headers"`
	Id                 int64         `json:"id"`
	Node               Name          `json:"node"`
	RunningTimeInNanos int64         `json:"running_time_in_nanos"`
	StartTimeInMillis  int64         `json:"start_time_in_millis"`
	Status             ReindexStatus `json:"status"`
	Type               string        `json:"type"`
}

// ReindexTaskBuilder holds ReindexTask struct and provides a builder API.
type ReindexTaskBuilder struct {
	v *ReindexTask
}

// NewReindexTask provides a builder for the ReindexTask struct.
func NewReindexTask() *ReindexTaskBuilder {
	r := ReindexTaskBuilder{
		&ReindexTask{},
	}

	return &r
}

// Build finalize the chain and returns the ReindexTask struct
func (rb *ReindexTaskBuilder) Build() ReindexTask {
	return *rb.v
}

// Action set the Action property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) Action(action string) *ReindexTaskBuilder {
	rb.v.Action = action
	return rb
}

// Cancellable set the Cancellable property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) Cancellable(cancellable bool) *ReindexTaskBuilder {
	rb.v.Cancellable = cancellable
	return rb
}

// Description set the Description property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) Description(description string) *ReindexTaskBuilder {
	rb.v.Description = description
	return rb
}

// Headers set the Headers property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) Headers(headers HttpHeaders) *ReindexTaskBuilder {
	rb.v.Headers = headers
	return rb
}

// Id set the Id property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) Id(id int64) *ReindexTaskBuilder {
	rb.v.Id = id
	return rb
}

// Node set the Node property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) Node(node Name) *ReindexTaskBuilder {
	rb.v.Node = node
	return rb
}

// RunningTimeInNanos set the RunningTimeInNanos property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) RunningTimeInNanos(runningtimeinnanos int64) *ReindexTaskBuilder {
	rb.v.RunningTimeInNanos = runningtimeinnanos
	return rb
}

// StartTimeInMillis set the StartTimeInMillis property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) StartTimeInMillis(starttimeinmillis int64) *ReindexTaskBuilder {
	rb.v.StartTimeInMillis = starttimeinmillis
	return rb
}

// Status set the Status property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) Status(status ReindexStatus) *ReindexTaskBuilder {
	rb.v.Status = status
	return rb
}

// Type set the Type property for ReindexTaskBuilder.
func (rb *ReindexTaskBuilder) Type_(type_ string) *ReindexTaskBuilder {
	rb.v.Type = type_
	return rb
}
