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

// IndexSettingsLifecycleStep type.
type IndexSettingsLifecycleStep struct {
	// WaitTimeThreshold Time to wait for the cluster to resolve allocation issues during an ILM
	// shrink action. Must be greater than 1h (1 hour).
	// See Shard allocation for shrink.
	WaitTimeThreshold *Time `json:"wait_time_threshold,omitempty"`
}

// IndexSettingsLifecycleStepBuilder holds IndexSettingsLifecycleStep struct and provides a builder API.
type IndexSettingsLifecycleStepBuilder struct {
	v *IndexSettingsLifecycleStep
}

// NewIndexSettingsLifecycleStep provides a builder for the IndexSettingsLifecycleStep struct.
func NewIndexSettingsLifecycleStep() *IndexSettingsLifecycleStepBuilder {
	r := IndexSettingsLifecycleStepBuilder{
		&IndexSettingsLifecycleStep{},
	}

	return &r
}

// Build finalize the chain and returns the IndexSettingsLifecycleStep struct
func (rb *IndexSettingsLifecycleStepBuilder) Build() IndexSettingsLifecycleStep {
	return *rb.v
}

// WaitTimeThreshold Time to wait for the cluster to resolve allocation issues during an ILM
// shrink action. Must be greater than 1h (1 hour).
// See Shard allocation for shrink.
func (rb *IndexSettingsLifecycleStepBuilder) WaitTimeThreshold(waittimethreshold Time) *IndexSettingsLifecycleStepBuilder {
	rb.v.WaitTimeThreshold = &waittimethreshold
	return rb
}
