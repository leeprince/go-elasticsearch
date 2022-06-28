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

// UsageStatsIndex type.
type UsageStatsIndex struct {
	Shards []UsageStatsShards `json:"shards"`
}

// UsageStatsIndexBuilder holds UsageStatsIndex struct and provides a builder API.
type UsageStatsIndexBuilder struct {
	v *UsageStatsIndex
}

// NewUsageStatsIndex provides a builder for the UsageStatsIndex struct.
func NewUsageStatsIndex() *UsageStatsIndexBuilder {
	r := UsageStatsIndexBuilder{
		&UsageStatsIndex{},
	}

	return &r
}

// Build finalize the chain and returns the UsageStatsIndex struct
func (rb *UsageStatsIndexBuilder) Build() UsageStatsIndex {
	return *rb.v
}

// Shards set the Shards property for UsageStatsIndexBuilder.
func (rb *UsageStatsIndexBuilder) Shards(shards ...UsageStatsShards) *UsageStatsIndexBuilder {
	rb.v.Shards = append(rb.v.Shards, shards...)
	return rb
}
