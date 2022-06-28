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

// ShardProfile type.
type ShardProfile struct {
	Aggregations []AggregationProfile `json:"aggregations"`
	Fetch        *FetchProfile        `json:"fetch,omitempty"`
	Id           string               `json:"id"`
	Searches     []SearchProfile      `json:"searches"`
}

// ShardProfileBuilder holds ShardProfile struct and provides a builder API.
type ShardProfileBuilder struct {
	v *ShardProfile
}

// NewShardProfile provides a builder for the ShardProfile struct.
func NewShardProfile() *ShardProfileBuilder {
	r := ShardProfileBuilder{
		&ShardProfile{},
	}

	return &r
}

// Build finalize the chain and returns the ShardProfile struct
func (rb *ShardProfileBuilder) Build() ShardProfile {
	return *rb.v
}

// Aggregations set the Aggregations property for ShardProfileBuilder.
func (rb *ShardProfileBuilder) Aggregations(aggregations ...AggregationProfile) *ShardProfileBuilder {
	rb.v.Aggregations = append(rb.v.Aggregations, aggregations...)
	return rb
}

// Fetch set the Fetch property for ShardProfileBuilder.
func (rb *ShardProfileBuilder) Fetch(fetch FetchProfile) *ShardProfileBuilder {
	rb.v.Fetch = &fetch
	return rb
}

// Id set the Id property for ShardProfileBuilder.
func (rb *ShardProfileBuilder) Id(id string) *ShardProfileBuilder {
	rb.v.Id = id
	return rb
}

// Searches set the Searches property for ShardProfileBuilder.
func (rb *ShardProfileBuilder) Searches(searches ...SearchProfile) *ShardProfileBuilder {
	rb.v.Searches = append(rb.v.Searches, searches...)
	return rb
}
