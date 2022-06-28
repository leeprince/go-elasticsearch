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

// BoolQuery type.
type BoolQuery struct {
	Boost              *float32            `json:"boost,omitempty"`
	Filter             []QueryContainer    `json:"filter,omitempty"`
	MinimumShouldMatch *MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	Must               []QueryContainer    `json:"must,omitempty"`
	MustNot            []QueryContainer    `json:"must_not,omitempty"`
	QueryName_         *string             `json:"_name,omitempty"`
	Should             []QueryContainer    `json:"should,omitempty"`
}

// BoolQueryBuilder holds BoolQuery struct and provides a builder API.
type BoolQueryBuilder struct {
	v *BoolQuery
}

// NewBoolQuery provides a builder for the BoolQuery struct.
func NewBoolQuery() *BoolQueryBuilder {
	r := BoolQueryBuilder{
		&BoolQuery{},
	}

	return &r
}

// Build finalize the chain and returns the BoolQuery struct
func (rb *BoolQueryBuilder) Build() BoolQuery {
	return *rb.v
}

// Boost set the Boost property for BoolQueryBuilder.
func (rb *BoolQueryBuilder) Boost(boost float32) *BoolQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Filter set the Filter property for BoolQueryBuilder.
func (rb *BoolQueryBuilder) Filter(arg []QueryContainer) *BoolQueryBuilder {
	rb.v.Filter = arg
	return rb
}

// MinimumShouldMatch set the MinimumShouldMatch property for BoolQueryBuilder.
func (rb *BoolQueryBuilder) MinimumShouldMatch(minimumshouldmatch MinimumShouldMatch) *BoolQueryBuilder {
	rb.v.MinimumShouldMatch = &minimumshouldmatch
	return rb
}

// Must set the Must property for BoolQueryBuilder.
func (rb *BoolQueryBuilder) Must(arg []QueryContainer) *BoolQueryBuilder {
	rb.v.Must = arg
	return rb
}

// MustNot set the MustNot property for BoolQueryBuilder.
func (rb *BoolQueryBuilder) MustNot(arg []QueryContainer) *BoolQueryBuilder {
	rb.v.MustNot = arg
	return rb
}

// QueryName_ set the QueryName_ property for BoolQueryBuilder.
func (rb *BoolQueryBuilder) QueryName_(queryname_ string) *BoolQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Should set the Should property for BoolQueryBuilder.
func (rb *BoolQueryBuilder) Should(arg []QueryContainer) *BoolQueryBuilder {
	rb.v.Should = arg
	return rb
}
