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

// TermsAggregateBaseVoid type.
type TermsAggregateBaseVoid struct {
	Buckets                 BucketsVoid `json:"buckets"`
	DocCountErrorUpperBound *int64      `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    *Metadata   `json:"meta,omitempty"`
	SumOtherDocCount        int64       `json:"sum_other_doc_count"`
}

// TermsAggregateBaseVoidBuilder holds TermsAggregateBaseVoid struct and provides a builder API.
type TermsAggregateBaseVoidBuilder struct {
	v *TermsAggregateBaseVoid
}

// NewTermsAggregateBaseVoid provides a builder for the TermsAggregateBaseVoid struct.
func NewTermsAggregateBaseVoid() *TermsAggregateBaseVoidBuilder {
	r := TermsAggregateBaseVoidBuilder{
		&TermsAggregateBaseVoid{},
	}

	return &r
}

// Build finalize the chain and returns the TermsAggregateBaseVoid struct
func (rb *TermsAggregateBaseVoidBuilder) Build() TermsAggregateBaseVoid {
	return *rb.v
}

// Buckets set the Buckets property for TermsAggregateBaseVoidBuilder.
func (rb *TermsAggregateBaseVoidBuilder) Buckets(buckets BucketsVoid) *TermsAggregateBaseVoidBuilder {
	rb.v.Buckets = buckets
	return rb
}

// DocCountErrorUpperBound set the DocCountErrorUpperBound property for TermsAggregateBaseVoidBuilder.
func (rb *TermsAggregateBaseVoidBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *TermsAggregateBaseVoidBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

// Meta set the Meta property for TermsAggregateBaseVoidBuilder.
func (rb *TermsAggregateBaseVoidBuilder) Meta(meta Metadata) *TermsAggregateBaseVoidBuilder {
	rb.v.Meta = &meta
	return rb
}

// SumOtherDocCount set the SumOtherDocCount property for TermsAggregateBaseVoidBuilder.
func (rb *TermsAggregateBaseVoidBuilder) SumOtherDocCount(sumotherdoccount int64) *TermsAggregateBaseVoidBuilder {
	rb.v.SumOtherDocCount = sumotherdoccount
	return rb
}
