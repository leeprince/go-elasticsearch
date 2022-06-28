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

// MlDataFrameAnalyticsJobs type.
type MlDataFrameAnalyticsJobs struct {
	All_           MlDataFrameAnalyticsJobsCount     `json:"_all"`
	AnalysisCounts *MlDataFrameAnalyticsJobsAnalysis `json:"analysis_counts,omitempty"`
	MemoryUsage    *MlDataFrameAnalyticsJobsMemory   `json:"memory_usage,omitempty"`
	Stopped        *MlDataFrameAnalyticsJobsCount    `json:"stopped,omitempty"`
}

// MlDataFrameAnalyticsJobsBuilder holds MlDataFrameAnalyticsJobs struct and provides a builder API.
type MlDataFrameAnalyticsJobsBuilder struct {
	v *MlDataFrameAnalyticsJobs
}

// NewMlDataFrameAnalyticsJobs provides a builder for the MlDataFrameAnalyticsJobs struct.
func NewMlDataFrameAnalyticsJobs() *MlDataFrameAnalyticsJobsBuilder {
	r := MlDataFrameAnalyticsJobsBuilder{
		&MlDataFrameAnalyticsJobs{},
	}

	return &r
}

// Build finalize the chain and returns the MlDataFrameAnalyticsJobs struct
func (rb *MlDataFrameAnalyticsJobsBuilder) Build() MlDataFrameAnalyticsJobs {
	return *rb.v
}

// All_ set the All_ property for MlDataFrameAnalyticsJobsBuilder.
func (rb *MlDataFrameAnalyticsJobsBuilder) All_(all_ MlDataFrameAnalyticsJobsCount) *MlDataFrameAnalyticsJobsBuilder {
	rb.v.All_ = all_
	return rb
}

// AnalysisCounts set the AnalysisCounts property for MlDataFrameAnalyticsJobsBuilder.
func (rb *MlDataFrameAnalyticsJobsBuilder) AnalysisCounts(analysiscounts MlDataFrameAnalyticsJobsAnalysis) *MlDataFrameAnalyticsJobsBuilder {
	rb.v.AnalysisCounts = &analysiscounts
	return rb
}

// MemoryUsage set the MemoryUsage property for MlDataFrameAnalyticsJobsBuilder.
func (rb *MlDataFrameAnalyticsJobsBuilder) MemoryUsage(memoryusage MlDataFrameAnalyticsJobsMemory) *MlDataFrameAnalyticsJobsBuilder {
	rb.v.MemoryUsage = &memoryusage
	return rb
}

// Stopped set the Stopped property for MlDataFrameAnalyticsJobsBuilder.
func (rb *MlDataFrameAnalyticsJobsBuilder) Stopped(stopped MlDataFrameAnalyticsJobsCount) *MlDataFrameAnalyticsJobsBuilder {
	rb.v.Stopped = &stopped
	return rb
}
