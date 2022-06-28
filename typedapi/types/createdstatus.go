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

// CreatedStatus type.
type CreatedStatus struct {
	Created bool `json:"created"`
}

// CreatedStatusBuilder holds CreatedStatus struct and provides a builder API.
type CreatedStatusBuilder struct {
	v *CreatedStatus
}

// NewCreatedStatus provides a builder for the CreatedStatus struct.
func NewCreatedStatus() *CreatedStatusBuilder {
	r := CreatedStatusBuilder{
		&CreatedStatus{},
	}

	return &r
}

// Build finalize the chain and returns the CreatedStatus struct
func (rb *CreatedStatusBuilder) Build() CreatedStatus {
	return *rb.v
}

// Created set the Created property for CreatedStatusBuilder.
func (rb *CreatedStatusBuilder) Created(created bool) *CreatedStatusBuilder {
	rb.v.Created = created
	return rb
}
