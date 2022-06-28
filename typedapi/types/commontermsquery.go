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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

// CommonTermsQuery type.
type CommonTermsQuery struct {
	Analyzer           *string             `json:"analyzer,omitempty"`
	Boost              *float32            `json:"boost,omitempty"`
	CutoffFrequency    *float64            `json:"cutoff_frequency,omitempty"`
	HighFreqOperator   *operator.Operator  `json:"high_freq_operator,omitempty"`
	LowFreqOperator    *operator.Operator  `json:"low_freq_operator,omitempty"`
	MinimumShouldMatch *MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	Query              string              `json:"query"`
	QueryName_         *string             `json:"_name,omitempty"`
}

// CommonTermsQueryBuilder holds CommonTermsQuery struct and provides a builder API.
type CommonTermsQueryBuilder struct {
	v *CommonTermsQuery
}

// NewCommonTermsQuery provides a builder for the CommonTermsQuery struct.
func NewCommonTermsQuery() *CommonTermsQueryBuilder {
	r := CommonTermsQueryBuilder{
		&CommonTermsQuery{},
	}

	return &r
}

// Build finalize the chain and returns the CommonTermsQuery struct
func (rb *CommonTermsQueryBuilder) Build() CommonTermsQuery {
	return *rb.v
}

// Analyzer set the Analyzer property for CommonTermsQueryBuilder.
func (rb *CommonTermsQueryBuilder) Analyzer(analyzer string) *CommonTermsQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Boost set the Boost property for CommonTermsQueryBuilder.
func (rb *CommonTermsQueryBuilder) Boost(boost float32) *CommonTermsQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// CutoffFrequency set the CutoffFrequency property for CommonTermsQueryBuilder.
func (rb *CommonTermsQueryBuilder) CutoffFrequency(cutofffrequency float64) *CommonTermsQueryBuilder {
	rb.v.CutoffFrequency = &cutofffrequency
	return rb
}

// HighFreqOperator set the HighFreqOperator property for CommonTermsQueryBuilder.
func (rb *CommonTermsQueryBuilder) HighFreqOperator(highfreqoperator operator.Operator) *CommonTermsQueryBuilder {
	rb.v.HighFreqOperator = &highfreqoperator
	return rb
}

// LowFreqOperator set the LowFreqOperator property for CommonTermsQueryBuilder.
func (rb *CommonTermsQueryBuilder) LowFreqOperator(lowfreqoperator operator.Operator) *CommonTermsQueryBuilder {
	rb.v.LowFreqOperator = &lowfreqoperator
	return rb
}

// MinimumShouldMatch set the MinimumShouldMatch property for CommonTermsQueryBuilder.
func (rb *CommonTermsQueryBuilder) MinimumShouldMatch(minimumshouldmatch MinimumShouldMatch) *CommonTermsQueryBuilder {
	rb.v.MinimumShouldMatch = &minimumshouldmatch
	return rb
}

// Query set the Query property for CommonTermsQueryBuilder.
func (rb *CommonTermsQueryBuilder) Query(query string) *CommonTermsQueryBuilder {
	rb.v.Query = query
	return rb
}

// QueryName_ set the QueryName_ property for CommonTermsQueryBuilder.
func (rb *CommonTermsQueryBuilder) QueryName_(queryname_ string) *CommonTermsQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
