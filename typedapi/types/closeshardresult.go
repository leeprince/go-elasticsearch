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

// CloseShardResult type.
type CloseShardResult struct {
	Failures []ShardFailure `json:"failures"`
}

// CloseShardResultBuilder holds CloseShardResult struct and provides a builder API.
type CloseShardResultBuilder struct {
	v *CloseShardResult
}

// NewCloseShardResult provides a builder for the CloseShardResult struct.
func NewCloseShardResult() *CloseShardResultBuilder {
	r := CloseShardResultBuilder{
		&CloseShardResult{},
	}

	return &r
}

// Build finalize the chain and returns the CloseShardResult struct
func (rb *CloseShardResultBuilder) Build() CloseShardResult {
	return *rb.v
}

// Failures set the Failures property for CloseShardResultBuilder.
func (rb *CloseShardResultBuilder) Failures(failures ...ShardFailure) *CloseShardResultBuilder {
	rb.v.Failures = append(rb.v.Failures, failures...)
	return rb
}
