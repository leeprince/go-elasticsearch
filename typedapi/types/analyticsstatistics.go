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

// AnalyticsStatistics type.
type AnalyticsStatistics struct {
	BoxplotUsage               int64  `json:"boxplot_usage"`
	CumulativeCardinalityUsage int64  `json:"cumulative_cardinality_usage"`
	MovingPercentilesUsage     int64  `json:"moving_percentiles_usage"`
	MultiTermsUsage            *int64 `json:"multi_terms_usage,omitempty"`
	NormalizeUsage             int64  `json:"normalize_usage"`
	RateUsage                  int64  `json:"rate_usage"`
	StringStatsUsage           int64  `json:"string_stats_usage"`
	TTestUsage                 int64  `json:"t_test_usage"`
	TopMetricsUsage            int64  `json:"top_metrics_usage"`
}

// AnalyticsStatisticsBuilder holds AnalyticsStatistics struct and provides a builder API.
type AnalyticsStatisticsBuilder struct {
	v *AnalyticsStatistics
}

// NewAnalyticsStatistics provides a builder for the AnalyticsStatistics struct.
func NewAnalyticsStatistics() *AnalyticsStatisticsBuilder {
	r := AnalyticsStatisticsBuilder{
		&AnalyticsStatistics{},
	}

	return &r
}

// Build finalize the chain and returns the AnalyticsStatistics struct
func (rb *AnalyticsStatisticsBuilder) Build() AnalyticsStatistics {
	return *rb.v
}

// BoxplotUsage set the BoxplotUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) BoxplotUsage(boxplotusage int64) *AnalyticsStatisticsBuilder {
	rb.v.BoxplotUsage = boxplotusage
	return rb
}

// CumulativeCardinalityUsage set the CumulativeCardinalityUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) CumulativeCardinalityUsage(cumulativecardinalityusage int64) *AnalyticsStatisticsBuilder {
	rb.v.CumulativeCardinalityUsage = cumulativecardinalityusage
	return rb
}

// MovingPercentilesUsage set the MovingPercentilesUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) MovingPercentilesUsage(movingpercentilesusage int64) *AnalyticsStatisticsBuilder {
	rb.v.MovingPercentilesUsage = movingpercentilesusage
	return rb
}

// MultiTermsUsage set the MultiTermsUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) MultiTermsUsage(multitermsusage int64) *AnalyticsStatisticsBuilder {
	rb.v.MultiTermsUsage = &multitermsusage
	return rb
}

// NormalizeUsage set the NormalizeUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) NormalizeUsage(normalizeusage int64) *AnalyticsStatisticsBuilder {
	rb.v.NormalizeUsage = normalizeusage
	return rb
}

// RateUsage set the RateUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) RateUsage(rateusage int64) *AnalyticsStatisticsBuilder {
	rb.v.RateUsage = rateusage
	return rb
}

// StringStatsUsage set the StringStatsUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) StringStatsUsage(stringstatsusage int64) *AnalyticsStatisticsBuilder {
	rb.v.StringStatsUsage = stringstatsusage
	return rb
}

// TTestUsage set the TTestUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) TTestUsage(ttestusage int64) *AnalyticsStatisticsBuilder {
	rb.v.TTestUsage = ttestusage
	return rb
}

// TopMetricsUsage set the TopMetricsUsage property for AnalyticsStatisticsBuilder.
func (rb *AnalyticsStatisticsBuilder) TopMetricsUsage(topmetricsusage int64) *AnalyticsStatisticsBuilder {
	rb.v.TopMetricsUsage = topmetricsusage
	return rb
}
