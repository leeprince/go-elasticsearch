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

// Checkpointing type.
type Checkpointing struct {
	ChangesLastDetectedAt         *int64           `json:"changes_last_detected_at,omitempty"`
	ChangesLastDetectedAtDateTime *DateString      `json:"changes_last_detected_at_date_time,omitempty"`
	Last                          CheckpointStats  `json:"last"`
	LastSearchTime                *int64           `json:"last_search_time,omitempty"`
	Next                          *CheckpointStats `json:"next,omitempty"`
	OperationsBehind              *int64           `json:"operations_behind,omitempty"`
}

// CheckpointingBuilder holds Checkpointing struct and provides a builder API.
type CheckpointingBuilder struct {
	v *Checkpointing
}

// NewCheckpointing provides a builder for the Checkpointing struct.
func NewCheckpointing() *CheckpointingBuilder {
	r := CheckpointingBuilder{
		&Checkpointing{},
	}

	return &r
}

// Build finalize the chain and returns the Checkpointing struct
func (rb *CheckpointingBuilder) Build() Checkpointing {
	return *rb.v
}

// ChangesLastDetectedAt set the ChangesLastDetectedAt property for CheckpointingBuilder.
func (rb *CheckpointingBuilder) ChangesLastDetectedAt(changeslastdetectedat int64) *CheckpointingBuilder {
	rb.v.ChangesLastDetectedAt = &changeslastdetectedat
	return rb
}

// ChangesLastDetectedAtDateTime set the ChangesLastDetectedAtDateTime property for CheckpointingBuilder.
func (rb *CheckpointingBuilder) ChangesLastDetectedAtDateTime(changeslastdetectedatdatetime DateString) *CheckpointingBuilder {
	rb.v.ChangesLastDetectedAtDateTime = &changeslastdetectedatdatetime
	return rb
}

// Last set the Last property for CheckpointingBuilder.
func (rb *CheckpointingBuilder) Last(last CheckpointStats) *CheckpointingBuilder {
	rb.v.Last = last
	return rb
}

// LastSearchTime set the LastSearchTime property for CheckpointingBuilder.
func (rb *CheckpointingBuilder) LastSearchTime(lastsearchtime int64) *CheckpointingBuilder {
	rb.v.LastSearchTime = &lastsearchtime
	return rb
}

// Next set the Next property for CheckpointingBuilder.
func (rb *CheckpointingBuilder) Next(next CheckpointStats) *CheckpointingBuilder {
	rb.v.Next = &next
	return rb
}

// OperationsBehind set the OperationsBehind property for CheckpointingBuilder.
func (rb *CheckpointingBuilder) OperationsBehind(operationsbehind int64) *CheckpointingBuilder {
	rb.v.OperationsBehind = &operationsbehind
	return rb
}
