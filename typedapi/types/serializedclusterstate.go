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

// SerializedClusterState type.
type SerializedClusterState struct {
	Diffs      *SerializedClusterStateDetail `json:"diffs,omitempty"`
	FullStates *SerializedClusterStateDetail `json:"full_states,omitempty"`
}

// SerializedClusterStateBuilder holds SerializedClusterState struct and provides a builder API.
type SerializedClusterStateBuilder struct {
	v *SerializedClusterState
}

// NewSerializedClusterState provides a builder for the SerializedClusterState struct.
func NewSerializedClusterState() *SerializedClusterStateBuilder {
	r := SerializedClusterStateBuilder{
		&SerializedClusterState{},
	}

	return &r
}

// Build finalize the chain and returns the SerializedClusterState struct
func (rb *SerializedClusterStateBuilder) Build() SerializedClusterState {
	return *rb.v
}

// Diffs set the Diffs property for SerializedClusterStateBuilder.
func (rb *SerializedClusterStateBuilder) Diffs(diffs SerializedClusterStateDetail) *SerializedClusterStateBuilder {
	rb.v.Diffs = &diffs
	return rb
}

// FullStates set the FullStates property for SerializedClusterStateBuilder.
func (rb *SerializedClusterStateBuilder) FullStates(fullstates SerializedClusterStateDetail) *SerializedClusterStateBuilder {
	rb.v.FullStates = &fullstates
	return rb
}
