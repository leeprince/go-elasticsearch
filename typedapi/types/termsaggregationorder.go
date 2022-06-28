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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// TermsAggregationOrder holds the union for the following types:
//     map[Field]sortorder.SortOrder
//     []map[Field]sortorder.SortOrder
type TermsAggregationOrder interface{}

// TermsAggregationOrderBuilder holds TermsAggregationOrder struct and provides a builder API.
type TermsAggregationOrderBuilder struct {
	v TermsAggregationOrder
}

// NewTermsAggregationOrder provides a builder for the TermsAggregationOrder struct.
func NewTermsAggregationOrder() *TermsAggregationOrderBuilder {
	return &TermsAggregationOrderBuilder{}
}

// Build finalize the chain and returns the TermsAggregationOrder struct
func (u *TermsAggregationOrderBuilder) Build() TermsAggregationOrder {
	return u.v
}

// Map set the Map property for TermsAggregationOrderBuilder.
func (u *TermsAggregationOrderBuilder) Map(v map[Field]sortorder.SortOrder) *TermsAggregationOrderBuilder {
	u.v = v
	return u
}

// SortOrders set the SortOrders property for TermsAggregationOrderBuilder.
func (u *TermsAggregationOrderBuilder) SortOrders(v []map[Field]sortorder.SortOrder) *TermsAggregationOrderBuilder {
	u.v = v
	return u
}
