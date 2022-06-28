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

// IntervalsQuery type.
type IntervalsQuery struct {
	AllOf      *IntervalsAllOf    `json:"all_of,omitempty"`
	AnyOf      *IntervalsAnyOf    `json:"any_of,omitempty"`
	Boost      *float32           `json:"boost,omitempty"`
	Fuzzy      *IntervalsFuzzy    `json:"fuzzy,omitempty"`
	Match      *IntervalsMatch    `json:"match,omitempty"`
	Prefix     *IntervalsPrefix   `json:"prefix,omitempty"`
	QueryName_ *string            `json:"_name,omitempty"`
	Wildcard   *IntervalsWildcard `json:"wildcard,omitempty"`
}

// IntervalsQueryBuilder holds IntervalsQuery struct and provides a builder API.
type IntervalsQueryBuilder struct {
	v *IntervalsQuery
}

// NewIntervalsQuery provides a builder for the IntervalsQuery struct.
func NewIntervalsQuery() *IntervalsQueryBuilder {
	r := IntervalsQueryBuilder{
		&IntervalsQuery{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsQuery struct
func (rb *IntervalsQueryBuilder) Build() IntervalsQuery {
	return *rb.v
}

// AllOf set the AllOf property for IntervalsQueryBuilder.
func (rb *IntervalsQueryBuilder) AllOf(allof IntervalsAllOf) *IntervalsQueryBuilder {
	rb.v.AllOf = &allof
	return rb
}

// AnyOf set the AnyOf property for IntervalsQueryBuilder.
func (rb *IntervalsQueryBuilder) AnyOf(anyof IntervalsAnyOf) *IntervalsQueryBuilder {
	rb.v.AnyOf = &anyof
	return rb
}

// Boost set the Boost property for IntervalsQueryBuilder.
func (rb *IntervalsQueryBuilder) Boost(boost float32) *IntervalsQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Fuzzy set the Fuzzy property for IntervalsQueryBuilder.
func (rb *IntervalsQueryBuilder) Fuzzy(fuzzy IntervalsFuzzy) *IntervalsQueryBuilder {
	rb.v.Fuzzy = &fuzzy
	return rb
}

// Match set the Match property for IntervalsQueryBuilder.
func (rb *IntervalsQueryBuilder) Match(match IntervalsMatch) *IntervalsQueryBuilder {
	rb.v.Match = &match
	return rb
}

// Prefix set the Prefix property for IntervalsQueryBuilder.
func (rb *IntervalsQueryBuilder) Prefix(prefix IntervalsPrefix) *IntervalsQueryBuilder {
	rb.v.Prefix = &prefix
	return rb
}

// QueryName_ set the QueryName_ property for IntervalsQueryBuilder.
func (rb *IntervalsQueryBuilder) QueryName_(queryname_ string) *IntervalsQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Wildcard set the Wildcard property for IntervalsQueryBuilder.
func (rb *IntervalsQueryBuilder) Wildcard(wildcard IntervalsWildcard) *IntervalsQueryBuilder {
	rb.v.Wildcard = &wildcard
	return rb
}
