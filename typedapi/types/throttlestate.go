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

// ThrottleState type.
type ThrottleState struct {
	Reason    string     `json:"reason"`
	Timestamp DateString `json:"timestamp"`
}

// ThrottleStateBuilder holds ThrottleState struct and provides a builder API.
type ThrottleStateBuilder struct {
	v *ThrottleState
}

// NewThrottleState provides a builder for the ThrottleState struct.
func NewThrottleState() *ThrottleStateBuilder {
	r := ThrottleStateBuilder{
		&ThrottleState{},
	}

	return &r
}

// Build finalize the chain and returns the ThrottleState struct
func (rb *ThrottleStateBuilder) Build() ThrottleState {
	return *rb.v
}

// Reason set the Reason property for ThrottleStateBuilder.
func (rb *ThrottleStateBuilder) Reason(reason string) *ThrottleStateBuilder {
	rb.v.Reason = reason
	return rb
}

// Timestamp set the Timestamp property for ThrottleStateBuilder.
func (rb *ThrottleStateBuilder) Timestamp(timestamp DateString) *ThrottleStateBuilder {
	rb.v.Timestamp = timestamp
	return rb
}
