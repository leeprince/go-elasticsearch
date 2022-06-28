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

// RollupJob type.
type RollupJob struct {
	Config RollupJobConfiguration `json:"config"`
	Stats  RollupJobStats         `json:"stats"`
	Status RollupJobStatus        `json:"status"`
}

// RollupJobBuilder holds RollupJob struct and provides a builder API.
type RollupJobBuilder struct {
	v *RollupJob
}

// NewRollupJob provides a builder for the RollupJob struct.
func NewRollupJob() *RollupJobBuilder {
	r := RollupJobBuilder{
		&RollupJob{},
	}

	return &r
}

// Build finalize the chain and returns the RollupJob struct
func (rb *RollupJobBuilder) Build() RollupJob {
	return *rb.v
}

// Config set the Config property for RollupJobBuilder.
func (rb *RollupJobBuilder) Config(config RollupJobConfiguration) *RollupJobBuilder {
	rb.v.Config = config
	return rb
}

// Stats set the Stats property for RollupJobBuilder.
func (rb *RollupJobBuilder) Stats(stats RollupJobStats) *RollupJobBuilder {
	rb.v.Stats = stats
	return rb
}

// Status set the Status property for RollupJobBuilder.
func (rb *RollupJobBuilder) Status(status RollupJobStatus) *RollupJobBuilder {
	rb.v.Status = status
	return rb
}
