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

// ActivationState type.
type ActivationState struct {
	Active    bool      `json:"active"`
	Timestamp Timestamp `json:"timestamp"`
}

// ActivationStateBuilder holds ActivationState struct and provides a builder API.
type ActivationStateBuilder struct {
	v *ActivationState
}

// NewActivationState provides a builder for the ActivationState struct.
func NewActivationState() *ActivationStateBuilder {
	r := ActivationStateBuilder{
		&ActivationState{},
	}

	return &r
}

// Build finalize the chain and returns the ActivationState struct
func (rb *ActivationStateBuilder) Build() ActivationState {
	return *rb.v
}

// Active set the Active property for ActivationStateBuilder.
func (rb *ActivationStateBuilder) Active(active bool) *ActivationStateBuilder {
	rb.v.Active = active
	return rb
}

// Timestamp set the Timestamp property for ActivationStateBuilder.
func (rb *ActivationStateBuilder) Timestamp(timestamp Timestamp) *ActivationStateBuilder {
	rb.v.Timestamp = timestamp
	return rb
}
