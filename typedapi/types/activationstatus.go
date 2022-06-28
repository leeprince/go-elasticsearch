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

// ActivationStatus type.
type ActivationStatus struct {
	Actions Actions         `json:"actions"`
	State   ActivationState `json:"state"`
	Version VersionNumber   `json:"version"`
}

// ActivationStatusBuilder holds ActivationStatus struct and provides a builder API.
type ActivationStatusBuilder struct {
	v *ActivationStatus
}

// NewActivationStatus provides a builder for the ActivationStatus struct.
func NewActivationStatus() *ActivationStatusBuilder {
	r := ActivationStatusBuilder{
		&ActivationStatus{},
	}

	return &r
}

// Build finalize the chain and returns the ActivationStatus struct
func (rb *ActivationStatusBuilder) Build() ActivationStatus {
	return *rb.v
}

// Actions set the Actions property for ActivationStatusBuilder.
func (rb *ActivationStatusBuilder) Actions(actions Actions) *ActivationStatusBuilder {
	rb.v.Actions = actions
	return rb
}

// State set the State property for ActivationStatusBuilder.
func (rb *ActivationStatusBuilder) State(state ActivationState) *ActivationStatusBuilder {
	rb.v.State = state
	return rb
}

// Version set the Version property for ActivationStatusBuilder.
func (rb *ActivationStatusBuilder) Version(version VersionNumber) *ActivationStatusBuilder {
	rb.v.Version = version
	return rb
}
