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

// TaskFailure type.
type TaskFailure struct {
	NodeId NodeId     `json:"node_id"`
	Reason ErrorCause `json:"reason"`
	Status string     `json:"status"`
	TaskId int64      `json:"task_id"`
}

// TaskFailureBuilder holds TaskFailure struct and provides a builder API.
type TaskFailureBuilder struct {
	v *TaskFailure
}

// NewTaskFailure provides a builder for the TaskFailure struct.
func NewTaskFailure() *TaskFailureBuilder {
	r := TaskFailureBuilder{
		&TaskFailure{},
	}

	return &r
}

// Build finalize the chain and returns the TaskFailure struct
func (rb *TaskFailureBuilder) Build() TaskFailure {
	return *rb.v
}

// NodeId set the NodeId property for TaskFailureBuilder.
func (rb *TaskFailureBuilder) NodeId(nodeid NodeId) *TaskFailureBuilder {
	rb.v.NodeId = nodeid
	return rb
}

// Reason set the Reason property for TaskFailureBuilder.
func (rb *TaskFailureBuilder) Reason(reason ErrorCause) *TaskFailureBuilder {
	rb.v.Reason = reason
	return rb
}

// Status set the Status property for TaskFailureBuilder.
func (rb *TaskFailureBuilder) Status(status string) *TaskFailureBuilder {
	rb.v.Status = status
	return rb
}

// TaskId set the TaskId property for TaskFailureBuilder.
func (rb *TaskFailureBuilder) TaskId(taskid int64) *TaskFailureBuilder {
	rb.v.TaskId = taskid
	return rb
}
