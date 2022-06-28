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

// Jobs type.
type Jobs struct {
	All_ *AllJobs       `json:"_all,omitempty"`
	Jobs map[string]Job `json:"jobs,omitempty"`
}

// JobsBuilder holds Jobs struct and provides a builder API.
type JobsBuilder struct {
	v *Jobs
}

// NewJobs provides a builder for the Jobs struct.
func NewJobs() *JobsBuilder {
	r := JobsBuilder{
		&Jobs{
			Jobs: make(map[string]Job, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Jobs struct
func (rb *JobsBuilder) Build() Jobs {
	return *rb.v
}

// All_ set the All_ property for JobsBuilder.
func (rb *JobsBuilder) All_(all_ AllJobs) *JobsBuilder {
	rb.v.All_ = &all_
	return rb
}

// Jobs set the Jobs property for JobsBuilder.
func (rb *JobsBuilder) Jobs(value map[string]Job) *JobsBuilder {
	rb.v.Jobs = value
	return rb
}
