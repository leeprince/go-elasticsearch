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

// DataframeAnalyticsStatsOutlierDetection type.
type DataframeAnalyticsStatsOutlierDetection struct {
	Parameters  OutlierDetectionParameters `json:"parameters"`
	Timestamp   DateString                 `json:"timestamp"`
	TimingStats TimingStats                `json:"timing_stats"`
}

// DataframeAnalyticsStatsOutlierDetectionBuilder holds DataframeAnalyticsStatsOutlierDetection struct and provides a builder API.
type DataframeAnalyticsStatsOutlierDetectionBuilder struct {
	v *DataframeAnalyticsStatsOutlierDetection
}

// NewDataframeAnalyticsStatsOutlierDetection provides a builder for the DataframeAnalyticsStatsOutlierDetection struct.
func NewDataframeAnalyticsStatsOutlierDetection() *DataframeAnalyticsStatsOutlierDetectionBuilder {
	r := DataframeAnalyticsStatsOutlierDetectionBuilder{
		&DataframeAnalyticsStatsOutlierDetection{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsStatsOutlierDetection struct
func (rb *DataframeAnalyticsStatsOutlierDetectionBuilder) Build() DataframeAnalyticsStatsOutlierDetection {
	return *rb.v
}

// Parameters set the Parameters property for DataframeAnalyticsStatsOutlierDetectionBuilder.
func (rb *DataframeAnalyticsStatsOutlierDetectionBuilder) Parameters(parameters OutlierDetectionParameters) *DataframeAnalyticsStatsOutlierDetectionBuilder {
	rb.v.Parameters = parameters
	return rb
}

// Timestamp set the Timestamp property for DataframeAnalyticsStatsOutlierDetectionBuilder.
func (rb *DataframeAnalyticsStatsOutlierDetectionBuilder) Timestamp(timestamp DateString) *DataframeAnalyticsStatsOutlierDetectionBuilder {
	rb.v.Timestamp = timestamp
	return rb
}

// TimingStats set the TimingStats property for DataframeAnalyticsStatsOutlierDetectionBuilder.
func (rb *DataframeAnalyticsStatsOutlierDetectionBuilder) TimingStats(timingstats TimingStats) *DataframeAnalyticsStatsOutlierDetectionBuilder {
	rb.v.TimingStats = timingstats
	return rb
}
