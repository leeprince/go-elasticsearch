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

// PublishedClusterStates type.
type PublishedClusterStates struct {
	CompatibleDiffs   *int64 `json:"compatible_diffs,omitempty"`
	FullStates        *int64 `json:"full_states,omitempty"`
	IncompatibleDiffs *int64 `json:"incompatible_diffs,omitempty"`
}

// PublishedClusterStatesBuilder holds PublishedClusterStates struct and provides a builder API.
type PublishedClusterStatesBuilder struct {
	v *PublishedClusterStates
}

// NewPublishedClusterStates provides a builder for the PublishedClusterStates struct.
func NewPublishedClusterStates() *PublishedClusterStatesBuilder {
	r := PublishedClusterStatesBuilder{
		&PublishedClusterStates{},
	}

	return &r
}

// Build finalize the chain and returns the PublishedClusterStates struct
func (rb *PublishedClusterStatesBuilder) Build() PublishedClusterStates {
	return *rb.v
}

// CompatibleDiffs set the CompatibleDiffs property for PublishedClusterStatesBuilder.
func (rb *PublishedClusterStatesBuilder) CompatibleDiffs(compatiblediffs int64) *PublishedClusterStatesBuilder {
	rb.v.CompatibleDiffs = &compatiblediffs
	return rb
}

// FullStates set the FullStates property for PublishedClusterStatesBuilder.
func (rb *PublishedClusterStatesBuilder) FullStates(fullstates int64) *PublishedClusterStatesBuilder {
	rb.v.FullStates = &fullstates
	return rb
}

// IncompatibleDiffs set the IncompatibleDiffs property for PublishedClusterStatesBuilder.
func (rb *PublishedClusterStatesBuilder) IncompatibleDiffs(incompatiblediffs int64) *PublishedClusterStatesBuilder {
	rb.v.IncompatibleDiffs = &incompatiblediffs
	return rb
}
