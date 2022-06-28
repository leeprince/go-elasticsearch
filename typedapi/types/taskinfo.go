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

// TaskInfo type.
type TaskInfo struct {
	Action             string            `json:"action"`
	Cancellable        bool              `json:"cancellable"`
	Description        *string           `json:"description,omitempty"`
	Headers            map[string]string `json:"headers"`
	Id                 int64             `json:"id"`
	Node               NodeId            `json:"node"`
	ParentTaskId       *TaskId           `json:"parent_task_id,omitempty"`
	RunningTime        *string           `json:"running_time,omitempty"`
	RunningTimeInNanos int64             `json:"running_time_in_nanos"`
	StartTimeInMillis  int64             `json:"start_time_in_millis"`
	Status             *TaskStatus       `json:"status,omitempty"`
	Type               string            `json:"type"`
}

// TaskInfoBuilder holds TaskInfo struct and provides a builder API.
type TaskInfoBuilder struct {
	v *TaskInfo
}

// NewTaskInfo provides a builder for the TaskInfo struct.
func NewTaskInfo() *TaskInfoBuilder {
	r := TaskInfoBuilder{
		&TaskInfo{
			Headers: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TaskInfo struct
func (rb *TaskInfoBuilder) Build() TaskInfo {
	return *rb.v
}

// Action set the Action property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) Action(action string) *TaskInfoBuilder {
	rb.v.Action = action
	return rb
}

// Cancellable set the Cancellable property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) Cancellable(cancellable bool) *TaskInfoBuilder {
	rb.v.Cancellable = cancellable
	return rb
}

// Description set the Description property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) Description(description string) *TaskInfoBuilder {
	rb.v.Description = &description
	return rb
}

// Headers set the Headers property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) Headers(value map[string]string) *TaskInfoBuilder {
	rb.v.Headers = value
	return rb
}

// Id set the Id property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) Id(id int64) *TaskInfoBuilder {
	rb.v.Id = id
	return rb
}

// Node set the Node property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) Node(node NodeId) *TaskInfoBuilder {
	rb.v.Node = node
	return rb
}

// ParentTaskId set the ParentTaskId property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) ParentTaskId(parenttaskid TaskId) *TaskInfoBuilder {
	rb.v.ParentTaskId = &parenttaskid
	return rb
}

// RunningTime set the RunningTime property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) RunningTime(runningtime string) *TaskInfoBuilder {
	rb.v.RunningTime = &runningtime
	return rb
}

// RunningTimeInNanos set the RunningTimeInNanos property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) RunningTimeInNanos(runningtimeinnanos int64) *TaskInfoBuilder {
	rb.v.RunningTimeInNanos = runningtimeinnanos
	return rb
}

// StartTimeInMillis set the StartTimeInMillis property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) StartTimeInMillis(starttimeinmillis int64) *TaskInfoBuilder {
	rb.v.StartTimeInMillis = starttimeinmillis
	return rb
}

// Status set the Status property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) Status(status TaskStatus) *TaskInfoBuilder {
	rb.v.Status = &status
	return rb
}

// Type set the Type property for TaskInfoBuilder.
func (rb *TaskInfoBuilder) Type_(type_ string) *TaskInfoBuilder {
	rb.v.Type = type_
	return rb
}
