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

// RangeQuery holds the union for the following types:
//     DateRangeQuery
//     NumberRangeQuery
type RangeQuery interface{}

// RangeQueryBuilder holds RangeQuery struct and provides a builder API.
type RangeQueryBuilder struct {
	v RangeQuery
}

// NewRangeQuery provides a builder for the RangeQuery struct.
func NewRangeQuery() *RangeQueryBuilder {
	return &RangeQueryBuilder{}
}

// Build finalize the chain and returns the RangeQuery struct
func (u *RangeQueryBuilder) Build() RangeQuery {
	return u.v
}

// DateRangeQuery set the DateRangeQuery property for RangeQueryBuilder.
func (u *RangeQueryBuilder) DateRangeQuery(v DateRangeQuery) *RangeQueryBuilder {
	u.v = v
	return u
}

// NumberRangeQuery set the NumberRangeQuery property for RangeQueryBuilder.
func (u *RangeQueryBuilder) NumberRangeQuery(v NumberRangeQuery) *RangeQueryBuilder {
	u.v = v
	return u
}
