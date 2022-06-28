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

// WatchStatus type.
type WatchStatus struct {
	Actions          Actions         `json:"actions"`
	ExecutionState   *string         `json:"execution_state,omitempty"`
	LastChecked      *DateString     `json:"last_checked,omitempty"`
	LastMetCondition *DateString     `json:"last_met_condition,omitempty"`
	State            ActivationState `json:"state"`
	Version          VersionNumber   `json:"version"`
}

// WatchStatusBuilder holds WatchStatus struct and provides a builder API.
type WatchStatusBuilder struct {
	v *WatchStatus
}

// NewWatchStatus provides a builder for the WatchStatus struct.
func NewWatchStatus() *WatchStatusBuilder {
	r := WatchStatusBuilder{
		&WatchStatus{},
	}

	return &r
}

// Build finalize the chain and returns the WatchStatus struct
func (rb *WatchStatusBuilder) Build() WatchStatus {
	return *rb.v
}

// Actions set the Actions property for WatchStatusBuilder.
func (rb *WatchStatusBuilder) Actions(actions Actions) *WatchStatusBuilder {
	rb.v.Actions = actions
	return rb
}

// ExecutionState set the ExecutionState property for WatchStatusBuilder.
func (rb *WatchStatusBuilder) ExecutionState(executionstate string) *WatchStatusBuilder {
	rb.v.ExecutionState = &executionstate
	return rb
}

// LastChecked set the LastChecked property for WatchStatusBuilder.
func (rb *WatchStatusBuilder) LastChecked(lastchecked DateString) *WatchStatusBuilder {
	rb.v.LastChecked = &lastchecked
	return rb
}

// LastMetCondition set the LastMetCondition property for WatchStatusBuilder.
func (rb *WatchStatusBuilder) LastMetCondition(lastmetcondition DateString) *WatchStatusBuilder {
	rb.v.LastMetCondition = &lastmetcondition
	return rb
}

// State set the State property for WatchStatusBuilder.
func (rb *WatchStatusBuilder) State(state ActivationState) *WatchStatusBuilder {
	rb.v.State = state
	return rb
}

// Version set the Version property for WatchStatusBuilder.
func (rb *WatchStatusBuilder) Version(version VersionNumber) *WatchStatusBuilder {
	rb.v.Version = version
	return rb
}
