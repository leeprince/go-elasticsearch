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

// DataframeAnalyticsStatsContainer type.
type DataframeAnalyticsStatsContainer struct {
	// ClassificationStats An object containing information about the classification analysis job.
	ClassificationStats *DataframeAnalyticsStatsHyperparameters `json:"classification_stats,omitempty"`
	// OutlierDetectionStats An object containing information about the outlier detection job.
	OutlierDetectionStats *DataframeAnalyticsStatsOutlierDetection `json:"outlier_detection_stats,omitempty"`
	// RegressionStats An object containing information about the regression analysis.
	RegressionStats *DataframeAnalyticsStatsHyperparameters `json:"regression_stats,omitempty"`
}

// DataframeAnalyticsStatsContainerBuilder holds DataframeAnalyticsStatsContainer struct and provides a builder API.
type DataframeAnalyticsStatsContainerBuilder struct {
	v *DataframeAnalyticsStatsContainer
}

// NewDataframeAnalyticsStatsContainer provides a builder for the DataframeAnalyticsStatsContainer struct.
func NewDataframeAnalyticsStatsContainer() *DataframeAnalyticsStatsContainerBuilder {
	r := DataframeAnalyticsStatsContainerBuilder{
		&DataframeAnalyticsStatsContainer{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsStatsContainer struct
func (rb *DataframeAnalyticsStatsContainerBuilder) Build() DataframeAnalyticsStatsContainer {
	return *rb.v
}

// ClassificationStats An object containing information about the classification analysis job.
func (rb *DataframeAnalyticsStatsContainerBuilder) ClassificationStats(classificationstats DataframeAnalyticsStatsHyperparameters) *DataframeAnalyticsStatsContainerBuilder {
	rb.v.ClassificationStats = &classificationstats
	return rb
}

// OutlierDetectionStats An object containing information about the outlier detection job.
func (rb *DataframeAnalyticsStatsContainerBuilder) OutlierDetectionStats(outlierdetectionstats DataframeAnalyticsStatsOutlierDetection) *DataframeAnalyticsStatsContainerBuilder {
	rb.v.OutlierDetectionStats = &outlierdetectionstats
	return rb
}

// RegressionStats An object containing information about the regression analysis.
func (rb *DataframeAnalyticsStatsContainerBuilder) RegressionStats(regressionstats DataframeAnalyticsStatsHyperparameters) *DataframeAnalyticsStatsContainerBuilder {
	rb.v.RegressionStats = &regressionstats
	return rb
}
