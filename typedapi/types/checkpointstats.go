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

// CheckpointStats type.
type CheckpointStats struct {
	Checkpoint           int64              `json:"checkpoint"`
	CheckpointProgress   *TransformProgress `json:"checkpoint_progress,omitempty"`
	TimeUpperBound       *DateString        `json:"time_upper_bound,omitempty"`
	TimeUpperBoundMillis *EpochMillis       `json:"time_upper_bound_millis,omitempty"`
	Timestamp            *DateString        `json:"timestamp,omitempty"`
	TimestampMillis      *EpochMillis       `json:"timestamp_millis,omitempty"`
}

// CheckpointStatsBuilder holds CheckpointStats struct and provides a builder API.
type CheckpointStatsBuilder struct {
	v *CheckpointStats
}

// NewCheckpointStats provides a builder for the CheckpointStats struct.
func NewCheckpointStats() *CheckpointStatsBuilder {
	r := CheckpointStatsBuilder{
		&CheckpointStats{},
	}

	return &r
}

// Build finalize the chain and returns the CheckpointStats struct
func (rb *CheckpointStatsBuilder) Build() CheckpointStats {
	return *rb.v
}

// Checkpoint set the Checkpoint property for CheckpointStatsBuilder.
func (rb *CheckpointStatsBuilder) Checkpoint(checkpoint int64) *CheckpointStatsBuilder {
	rb.v.Checkpoint = checkpoint
	return rb
}

// CheckpointProgress set the CheckpointProgress property for CheckpointStatsBuilder.
func (rb *CheckpointStatsBuilder) CheckpointProgress(checkpointprogress TransformProgress) *CheckpointStatsBuilder {
	rb.v.CheckpointProgress = &checkpointprogress
	return rb
}

// TimeUpperBound set the TimeUpperBound property for CheckpointStatsBuilder.
func (rb *CheckpointStatsBuilder) TimeUpperBound(timeupperbound DateString) *CheckpointStatsBuilder {
	rb.v.TimeUpperBound = &timeupperbound
	return rb
}

// TimeUpperBoundMillis set the TimeUpperBoundMillis property for CheckpointStatsBuilder.
func (rb *CheckpointStatsBuilder) TimeUpperBoundMillis(timeupperboundmillis EpochMillis) *CheckpointStatsBuilder {
	rb.v.TimeUpperBoundMillis = &timeupperboundmillis
	return rb
}

// Timestamp set the Timestamp property for CheckpointStatsBuilder.
func (rb *CheckpointStatsBuilder) Timestamp(timestamp DateString) *CheckpointStatsBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

// TimestampMillis set the TimestampMillis property for CheckpointStatsBuilder.
func (rb *CheckpointStatsBuilder) TimestampMillis(timestampmillis EpochMillis) *CheckpointStatsBuilder {
	rb.v.TimestampMillis = &timestampmillis
	return rb
}
