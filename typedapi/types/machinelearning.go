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

// MachineLearning type.
type MachineLearning struct {
	Available              bool                     `json:"available"`
	DataFrameAnalyticsJobs MlDataFrameAnalyticsJobs `json:"data_frame_analytics_jobs"`
	Datafeeds              map[string]Datafeed      `json:"datafeeds"`
	Enabled                bool                     `json:"enabled"`
	Inference              MlInference              `json:"inference"`
	Jobs                   Jobs                     `json:"jobs"`
	NodeCount              int                      `json:"node_count"`
}

// MachineLearningBuilder holds MachineLearning struct and provides a builder API.
type MachineLearningBuilder struct {
	v *MachineLearning
}

// NewMachineLearning provides a builder for the MachineLearning struct.
func NewMachineLearning() *MachineLearningBuilder {
	r := MachineLearningBuilder{
		&MachineLearning{
			Datafeeds: make(map[string]Datafeed, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MachineLearning struct
func (rb *MachineLearningBuilder) Build() MachineLearning {
	return *rb.v
}

// Available set the Available property for MachineLearningBuilder.
func (rb *MachineLearningBuilder) Available(available bool) *MachineLearningBuilder {
	rb.v.Available = available
	return rb
}

// DataFrameAnalyticsJobs set the DataFrameAnalyticsJobs property for MachineLearningBuilder.
func (rb *MachineLearningBuilder) DataFrameAnalyticsJobs(dataframeanalyticsjobs MlDataFrameAnalyticsJobs) *MachineLearningBuilder {
	rb.v.DataFrameAnalyticsJobs = dataframeanalyticsjobs
	return rb
}

// Datafeeds set the Datafeeds property for MachineLearningBuilder.
func (rb *MachineLearningBuilder) Datafeeds(value map[string]Datafeed) *MachineLearningBuilder {
	rb.v.Datafeeds = value
	return rb
}

// Enabled set the Enabled property for MachineLearningBuilder.
func (rb *MachineLearningBuilder) Enabled(enabled bool) *MachineLearningBuilder {
	rb.v.Enabled = enabled
	return rb
}

// Inference set the Inference property for MachineLearningBuilder.
func (rb *MachineLearningBuilder) Inference(inference MlInference) *MachineLearningBuilder {
	rb.v.Inference = inference
	return rb
}

// Jobs set the Jobs property for MachineLearningBuilder.
func (rb *MachineLearningBuilder) Jobs(jobs Jobs) *MachineLearningBuilder {
	rb.v.Jobs = jobs
	return rb
}

// NodeCount set the NodeCount property for MachineLearningBuilder.
func (rb *MachineLearningBuilder) NodeCount(nodecount int) *MachineLearningBuilder {
	rb.v.NodeCount = nodecount
	return rb
}
