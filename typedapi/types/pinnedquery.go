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

// PinnedQuery type.
type PinnedQuery struct {
	Boost      *float32        `json:"boost,omitempty"`
	Docs       []PinnedDoc     `json:"docs,omitempty"`
	Ids        []Id            `json:"ids,omitempty"`
	Organic    *QueryContainer `json:"organic,omitempty"`
	QueryName_ *string         `json:"_name,omitempty"`
}

// PinnedQueryBuilder holds PinnedQuery struct and provides a builder API.
type PinnedQueryBuilder struct {
	v *PinnedQuery
}

// NewPinnedQuery provides a builder for the PinnedQuery struct.
func NewPinnedQuery() *PinnedQueryBuilder {
	r := PinnedQueryBuilder{
		&PinnedQuery{},
	}

	return &r
}

// Build finalize the chain and returns the PinnedQuery struct
func (rb *PinnedQueryBuilder) Build() PinnedQuery {
	return *rb.v
}

// Boost set the Boost property for PinnedQueryBuilder.
func (rb *PinnedQueryBuilder) Boost(boost float32) *PinnedQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Docs set the Docs property for PinnedQueryBuilder.
func (rb *PinnedQueryBuilder) Docs(docs ...PinnedDoc) *PinnedQueryBuilder {
	rb.v.Docs = append(rb.v.Docs, docs...)
	return rb
}

// Ids set the Ids property for PinnedQueryBuilder.
func (rb *PinnedQueryBuilder) Ids(ids ...Id) *PinnedQueryBuilder {
	rb.v.Ids = append(rb.v.Ids, ids...)
	return rb
}

// Organic set the Organic property for PinnedQueryBuilder.
func (rb *PinnedQueryBuilder) Organic(organic QueryContainer) *PinnedQueryBuilder {
	rb.v.Organic = &organic
	return rb
}

// QueryName_ set the QueryName_ property for PinnedQueryBuilder.
func (rb *PinnedQueryBuilder) QueryName_(queryname_ string) *PinnedQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
