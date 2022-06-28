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

// RankEvalMetric type.
type RankEvalMetric struct {
	Dcg                    *RankEvalMetricDiscountedCumulativeGain `json:"dcg,omitempty"`
	ExpectedReciprocalRank *RankEvalMetricExpectedReciprocalRank   `json:"expected_reciprocal_rank,omitempty"`
	MeanReciprocalRank     *RankEvalMetricMeanReciprocalRank       `json:"mean_reciprocal_rank,omitempty"`
	Precision              *RankEvalMetricPrecision                `json:"precision,omitempty"`
	Recall                 *RankEvalMetricRecall                   `json:"recall,omitempty"`
}

// RankEvalMetricBuilder holds RankEvalMetric struct and provides a builder API.
type RankEvalMetricBuilder struct {
	v *RankEvalMetric
}

// NewRankEvalMetric provides a builder for the RankEvalMetric struct.
func NewRankEvalMetric() *RankEvalMetricBuilder {
	r := RankEvalMetricBuilder{
		&RankEvalMetric{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalMetric struct
func (rb *RankEvalMetricBuilder) Build() RankEvalMetric {
	return *rb.v
}

// Dcg set the Dcg property for RankEvalMetricBuilder.
func (rb *RankEvalMetricBuilder) Dcg(dcg RankEvalMetricDiscountedCumulativeGain) *RankEvalMetricBuilder {
	rb.v.Dcg = &dcg
	return rb
}

// ExpectedReciprocalRank set the ExpectedReciprocalRank property for RankEvalMetricBuilder.
func (rb *RankEvalMetricBuilder) ExpectedReciprocalRank(expectedreciprocalrank RankEvalMetricExpectedReciprocalRank) *RankEvalMetricBuilder {
	rb.v.ExpectedReciprocalRank = &expectedreciprocalrank
	return rb
}

// MeanReciprocalRank set the MeanReciprocalRank property for RankEvalMetricBuilder.
func (rb *RankEvalMetricBuilder) MeanReciprocalRank(meanreciprocalrank RankEvalMetricMeanReciprocalRank) *RankEvalMetricBuilder {
	rb.v.MeanReciprocalRank = &meanreciprocalrank
	return rb
}

// Precision set the Precision property for RankEvalMetricBuilder.
func (rb *RankEvalMetricBuilder) Precision(precision RankEvalMetricPrecision) *RankEvalMetricBuilder {
	rb.v.Precision = &precision
	return rb
}

// Recall set the Recall property for RankEvalMetricBuilder.
func (rb *RankEvalMetricBuilder) Recall(recall RankEvalMetricRecall) *RankEvalMetricBuilder {
	rb.v.Recall = &recall
	return rb
}
