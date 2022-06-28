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

// SnapshotInfo type.
type SnapshotInfo struct {
	DataStreams        []string                   `json:"data_streams"`
	Duration           *Time                      `json:"duration,omitempty"`
	DurationInMillis   *EpochMillis               `json:"duration_in_millis,omitempty"`
	EndTime            *Time                      `json:"end_time,omitempty"`
	EndTimeInMillis    *EpochMillis               `json:"end_time_in_millis,omitempty"`
	Failures           []SnapshotShardFailure     `json:"failures,omitempty"`
	FeatureStates      []InfoFeatureState         `json:"feature_states,omitempty"`
	IncludeGlobalState *bool                      `json:"include_global_state,omitempty"`
	IndexDetails       map[IndexName]IndexDetails `json:"index_details,omitempty"`
	Indices            []IndexName                `json:"indices"`
	Metadata           *Metadata                  `json:"metadata,omitempty"`
	Reason             *string                    `json:"reason,omitempty"`
	Repository         *Name                      `json:"repository,omitempty"`
	Shards             *ShardStatistics           `json:"shards,omitempty"`
	Snapshot           Name                       `json:"snapshot"`
	StartTime          *Time                      `json:"start_time,omitempty"`
	StartTimeInMillis  *EpochMillis               `json:"start_time_in_millis,omitempty"`
	State              *string                    `json:"state,omitempty"`
	Uuid               Uuid                       `json:"uuid"`
	Version            *VersionString             `json:"version,omitempty"`
	VersionId          *VersionNumber             `json:"version_id,omitempty"`
}

// SnapshotInfoBuilder holds SnapshotInfo struct and provides a builder API.
type SnapshotInfoBuilder struct {
	v *SnapshotInfo
}

// NewSnapshotInfo provides a builder for the SnapshotInfo struct.
func NewSnapshotInfo() *SnapshotInfoBuilder {
	r := SnapshotInfoBuilder{
		&SnapshotInfo{
			IndexDetails: make(map[IndexName]IndexDetails, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotInfo struct
func (rb *SnapshotInfoBuilder) Build() SnapshotInfo {
	return *rb.v
}

// DataStreams set the DataStreams property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) DataStreams(data_streams ...string) *SnapshotInfoBuilder {
	rb.v.DataStreams = append(rb.v.DataStreams, data_streams...)
	return rb
}

// Duration set the Duration property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Duration(duration Time) *SnapshotInfoBuilder {
	rb.v.Duration = &duration
	return rb
}

// DurationInMillis set the DurationInMillis property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) DurationInMillis(durationinmillis EpochMillis) *SnapshotInfoBuilder {
	rb.v.DurationInMillis = &durationinmillis
	return rb
}

// EndTime set the EndTime property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) EndTime(endtime Time) *SnapshotInfoBuilder {
	rb.v.EndTime = &endtime
	return rb
}

// EndTimeInMillis set the EndTimeInMillis property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) EndTimeInMillis(endtimeinmillis EpochMillis) *SnapshotInfoBuilder {
	rb.v.EndTimeInMillis = &endtimeinmillis
	return rb
}

// Failures set the Failures property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Failures(failures ...SnapshotShardFailure) *SnapshotInfoBuilder {
	rb.v.Failures = append(rb.v.Failures, failures...)
	return rb
}

// FeatureStates set the FeatureStates property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) FeatureStates(feature_states ...InfoFeatureState) *SnapshotInfoBuilder {
	rb.v.FeatureStates = append(rb.v.FeatureStates, feature_states...)
	return rb
}

// IncludeGlobalState set the IncludeGlobalState property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) IncludeGlobalState(includeglobalstate bool) *SnapshotInfoBuilder {
	rb.v.IncludeGlobalState = &includeglobalstate
	return rb
}

// IndexDetails set the IndexDetails property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) IndexDetails(value map[IndexName]IndexDetails) *SnapshotInfoBuilder {
	rb.v.IndexDetails = value
	return rb
}

// Indices set the Indices property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Indices(indices ...IndexName) *SnapshotInfoBuilder {
	rb.v.Indices = append(rb.v.Indices, indices...)
	return rb
}

// Metadata set the Metadata property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Metadata(metadata Metadata) *SnapshotInfoBuilder {
	rb.v.Metadata = &metadata
	return rb
}

// Reason set the Reason property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Reason(reason string) *SnapshotInfoBuilder {
	rb.v.Reason = &reason
	return rb
}

// Repository set the Repository property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Repository(repository Name) *SnapshotInfoBuilder {
	rb.v.Repository = &repository
	return rb
}

// Shards set the Shards property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Shards(shards ShardStatistics) *SnapshotInfoBuilder {
	rb.v.Shards = &shards
	return rb
}

// Snapshot set the Snapshot property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Snapshot(snapshot Name) *SnapshotInfoBuilder {
	rb.v.Snapshot = snapshot
	return rb
}

// StartTime set the StartTime property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) StartTime(starttime Time) *SnapshotInfoBuilder {
	rb.v.StartTime = &starttime
	return rb
}

// StartTimeInMillis set the StartTimeInMillis property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) StartTimeInMillis(starttimeinmillis EpochMillis) *SnapshotInfoBuilder {
	rb.v.StartTimeInMillis = &starttimeinmillis
	return rb
}

// State set the State property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) State(state string) *SnapshotInfoBuilder {
	rb.v.State = &state
	return rb
}

// Uuid set the Uuid property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Uuid(uuid Uuid) *SnapshotInfoBuilder {
	rb.v.Uuid = uuid
	return rb
}

// Version set the Version property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) Version(version VersionString) *SnapshotInfoBuilder {
	rb.v.Version = &version
	return rb
}

// VersionId set the VersionId property for SnapshotInfoBuilder.
func (rb *SnapshotInfoBuilder) VersionId(versionid VersionNumber) *SnapshotInfoBuilder {
	rb.v.VersionId = &versionid
	return rb
}
