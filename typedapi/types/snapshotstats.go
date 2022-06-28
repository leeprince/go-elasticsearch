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

// SnapshotStats type.
type SnapshotStats struct {
	Incremental       FileCountSnapshotStats `json:"incremental"`
	StartTimeInMillis int64                  `json:"start_time_in_millis"`
	TimeInMillis      int64                  `json:"time_in_millis"`
	Total             FileCountSnapshotStats `json:"total"`
}

// SnapshotStatsBuilder holds SnapshotStats struct and provides a builder API.
type SnapshotStatsBuilder struct {
	v *SnapshotStats
}

// NewSnapshotStats provides a builder for the SnapshotStats struct.
func NewSnapshotStats() *SnapshotStatsBuilder {
	r := SnapshotStatsBuilder{
		&SnapshotStats{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotStats struct
func (rb *SnapshotStatsBuilder) Build() SnapshotStats {
	return *rb.v
}

// Incremental set the Incremental property for SnapshotStatsBuilder.
func (rb *SnapshotStatsBuilder) Incremental(incremental FileCountSnapshotStats) *SnapshotStatsBuilder {
	rb.v.Incremental = incremental
	return rb
}

// StartTimeInMillis set the StartTimeInMillis property for SnapshotStatsBuilder.
func (rb *SnapshotStatsBuilder) StartTimeInMillis(starttimeinmillis int64) *SnapshotStatsBuilder {
	rb.v.StartTimeInMillis = starttimeinmillis
	return rb
}

// TimeInMillis set the TimeInMillis property for SnapshotStatsBuilder.
func (rb *SnapshotStatsBuilder) TimeInMillis(timeinmillis int64) *SnapshotStatsBuilder {
	rb.v.TimeInMillis = timeinmillis
	return rb
}

// Total set the Total property for SnapshotStatsBuilder.
func (rb *SnapshotStatsBuilder) Total(total FileCountSnapshotStats) *SnapshotStatsBuilder {
	rb.v.Total = total
	return rb
}
