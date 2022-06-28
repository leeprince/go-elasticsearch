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

// BoostingQuery type.
type BoostingQuery struct {
	Boost         *float32        `json:"boost,omitempty"`
	Negative      *QueryContainer `json:"negative,omitempty"`
	NegativeBoost float64         `json:"negative_boost"`
	Positive      *QueryContainer `json:"positive,omitempty"`
	QueryName_    *string         `json:"_name,omitempty"`
}

// BoostingQueryBuilder holds BoostingQuery struct and provides a builder API.
type BoostingQueryBuilder struct {
	v *BoostingQuery
}

// NewBoostingQuery provides a builder for the BoostingQuery struct.
func NewBoostingQuery() *BoostingQueryBuilder {
	r := BoostingQueryBuilder{
		&BoostingQuery{},
	}

	return &r
}

// Build finalize the chain and returns the BoostingQuery struct
func (rb *BoostingQueryBuilder) Build() BoostingQuery {
	return *rb.v
}

// Boost set the Boost property for BoostingQueryBuilder.
func (rb *BoostingQueryBuilder) Boost(boost float32) *BoostingQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Negative set the Negative property for BoostingQueryBuilder.
func (rb *BoostingQueryBuilder) Negative(negative QueryContainer) *BoostingQueryBuilder {
	rb.v.Negative = &negative
	return rb
}

// NegativeBoost set the NegativeBoost property for BoostingQueryBuilder.
func (rb *BoostingQueryBuilder) NegativeBoost(negativeboost float64) *BoostingQueryBuilder {
	rb.v.NegativeBoost = negativeboost
	return rb
}

// Positive set the Positive property for BoostingQueryBuilder.
func (rb *BoostingQueryBuilder) Positive(positive QueryContainer) *BoostingQueryBuilder {
	rb.v.Positive = &positive
	return rb
}

// QueryName_ set the QueryName_ property for BoostingQueryBuilder.
func (rb *BoostingQueryBuilder) QueryName_(queryname_ string) *BoostingQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
