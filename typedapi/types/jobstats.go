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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jobstate"
)

// JobStats type.
type JobStats struct {
	AssignmentExplanation *string               `json:"assignment_explanation,omitempty"`
	DataCounts            DataCounts            `json:"data_counts"`
	Deleting              *bool                 `json:"deleting,omitempty"`
	ForecastsStats        JobForecastStatistics `json:"forecasts_stats"`
	JobId                 string                `json:"job_id"`
	ModelSizeStats        ModelSizeStats        `json:"model_size_stats"`
	Node                  *DiscoveryNode        `json:"node,omitempty"`
	OpenTime              *DateString           `json:"open_time,omitempty"`
	State                 jobstate.JobState     `json:"state"`
	TimingStats           JobTimingStats        `json:"timing_stats"`
}

// JobStatsBuilder holds JobStats struct and provides a builder API.
type JobStatsBuilder struct {
	v *JobStats
}

// NewJobStats provides a builder for the JobStats struct.
func NewJobStats() *JobStatsBuilder {
	r := JobStatsBuilder{
		&JobStats{},
	}

	return &r
}

// Build finalize the chain and returns the JobStats struct
func (rb *JobStatsBuilder) Build() JobStats {
	return *rb.v
}

// AssignmentExplanation set the AssignmentExplanation property for JobStatsBuilder.
func (rb *JobStatsBuilder) AssignmentExplanation(assignmentexplanation string) *JobStatsBuilder {
	rb.v.AssignmentExplanation = &assignmentexplanation
	return rb
}

// DataCounts set the DataCounts property for JobStatsBuilder.
func (rb *JobStatsBuilder) DataCounts(datacounts DataCounts) *JobStatsBuilder {
	rb.v.DataCounts = datacounts
	return rb
}

// Deleting set the Deleting property for JobStatsBuilder.
func (rb *JobStatsBuilder) Deleting(deleting bool) *JobStatsBuilder {
	rb.v.Deleting = &deleting
	return rb
}

// ForecastsStats set the ForecastsStats property for JobStatsBuilder.
func (rb *JobStatsBuilder) ForecastsStats(forecastsstats JobForecastStatistics) *JobStatsBuilder {
	rb.v.ForecastsStats = forecastsstats
	return rb
}

// JobId set the JobId property for JobStatsBuilder.
func (rb *JobStatsBuilder) JobId(jobid string) *JobStatsBuilder {
	rb.v.JobId = jobid
	return rb
}

// ModelSizeStats set the ModelSizeStats property for JobStatsBuilder.
func (rb *JobStatsBuilder) ModelSizeStats(modelsizestats ModelSizeStats) *JobStatsBuilder {
	rb.v.ModelSizeStats = modelsizestats
	return rb
}

// Node set the Node property for JobStatsBuilder.
func (rb *JobStatsBuilder) Node(node DiscoveryNode) *JobStatsBuilder {
	rb.v.Node = &node
	return rb
}

// OpenTime set the OpenTime property for JobStatsBuilder.
func (rb *JobStatsBuilder) OpenTime(opentime DateString) *JobStatsBuilder {
	rb.v.OpenTime = &opentime
	return rb
}

// State set the State property for JobStatsBuilder.
func (rb *JobStatsBuilder) State(state jobstate.JobState) *JobStatsBuilder {
	rb.v.State = state
	return rb
}

// TimingStats set the TimingStats property for JobStatsBuilder.
func (rb *JobStatsBuilder) TimingStats(timingstats JobTimingStats) *JobStatsBuilder {
	rb.v.TimingStats = timingstats
	return rb
}
