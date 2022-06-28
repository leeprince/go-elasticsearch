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

// ShardsStats type.
type ShardsStats struct {
	Done         int64 `json:"done"`
	Failed       int64 `json:"failed"`
	Finalizing   int64 `json:"finalizing"`
	Initializing int64 `json:"initializing"`
	Started      int64 `json:"started"`
	Total        int64 `json:"total"`
}

// ShardsStatsBuilder holds ShardsStats struct and provides a builder API.
type ShardsStatsBuilder struct {
	v *ShardsStats
}

// NewShardsStats provides a builder for the ShardsStats struct.
func NewShardsStats() *ShardsStatsBuilder {
	r := ShardsStatsBuilder{
		&ShardsStats{},
	}

	return &r
}

// Build finalize the chain and returns the ShardsStats struct
func (rb *ShardsStatsBuilder) Build() ShardsStats {
	return *rb.v
}

// Done set the Done property for ShardsStatsBuilder.
func (rb *ShardsStatsBuilder) Done(done int64) *ShardsStatsBuilder {
	rb.v.Done = done
	return rb
}

// Failed set the Failed property for ShardsStatsBuilder.
func (rb *ShardsStatsBuilder) Failed(failed int64) *ShardsStatsBuilder {
	rb.v.Failed = failed
	return rb
}

// Finalizing set the Finalizing property for ShardsStatsBuilder.
func (rb *ShardsStatsBuilder) Finalizing(finalizing int64) *ShardsStatsBuilder {
	rb.v.Finalizing = finalizing
	return rb
}

// Initializing set the Initializing property for ShardsStatsBuilder.
func (rb *ShardsStatsBuilder) Initializing(initializing int64) *ShardsStatsBuilder {
	rb.v.Initializing = initializing
	return rb
}

// Started set the Started property for ShardsStatsBuilder.
func (rb *ShardsStatsBuilder) Started(started int64) *ShardsStatsBuilder {
	rb.v.Started = started
	return rb
}

// Total set the Total property for ShardsStatsBuilder.
func (rb *ShardsStatsBuilder) Total(total int64) *ShardsStatsBuilder {
	rb.v.Total = total
	return rb
}
