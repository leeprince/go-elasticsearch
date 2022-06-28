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

// ShardStatistics type.
type ShardStatistics struct {
	Failed     uint           `json:"failed"`
	Failures   []ShardFailure `json:"failures,omitempty"`
	Skipped    *uint          `json:"skipped,omitempty"`
	Successful uint           `json:"successful"`
	Total      uint           `json:"total"`
}

// ShardStatisticsBuilder holds ShardStatistics struct and provides a builder API.
type ShardStatisticsBuilder struct {
	v *ShardStatistics
}

// NewShardStatistics provides a builder for the ShardStatistics struct.
func NewShardStatistics() *ShardStatisticsBuilder {
	r := ShardStatisticsBuilder{
		&ShardStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the ShardStatistics struct
func (rb *ShardStatisticsBuilder) Build() ShardStatistics {
	return *rb.v
}

// Failed set the Failed property for ShardStatisticsBuilder.
func (rb *ShardStatisticsBuilder) Failed(failed uint) *ShardStatisticsBuilder {
	rb.v.Failed = failed
	return rb
}

// Failures set the Failures property for ShardStatisticsBuilder.
func (rb *ShardStatisticsBuilder) Failures(failures ...ShardFailure) *ShardStatisticsBuilder {
	rb.v.Failures = append(rb.v.Failures, failures...)
	return rb
}

// Skipped set the Skipped property for ShardStatisticsBuilder.
func (rb *ShardStatisticsBuilder) Skipped(skipped uint) *ShardStatisticsBuilder {
	rb.v.Skipped = &skipped
	return rb
}

// Successful set the Successful property for ShardStatisticsBuilder.
func (rb *ShardStatisticsBuilder) Successful(successful uint) *ShardStatisticsBuilder {
	rb.v.Successful = successful
	return rb
}

// Total set the Total property for ShardStatisticsBuilder.
func (rb *ShardStatisticsBuilder) Total(total uint) *ShardStatisticsBuilder {
	rb.v.Total = total
	return rb
}
