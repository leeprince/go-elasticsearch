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

// UnmappedTermsAggregate type.
type UnmappedTermsAggregate struct {
	Buckets                 BucketsVoid `json:"buckets"`
	DocCountErrorUpperBound *int64      `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    *Metadata   `json:"meta,omitempty"`
	SumOtherDocCount        int64       `json:"sum_other_doc_count"`
}

// UnmappedTermsAggregateBuilder holds UnmappedTermsAggregate struct and provides a builder API.
type UnmappedTermsAggregateBuilder struct {
	v *UnmappedTermsAggregate
}

// NewUnmappedTermsAggregate provides a builder for the UnmappedTermsAggregate struct.
func NewUnmappedTermsAggregate() *UnmappedTermsAggregateBuilder {
	r := UnmappedTermsAggregateBuilder{
		&UnmappedTermsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the UnmappedTermsAggregate struct
func (rb *UnmappedTermsAggregateBuilder) Build() UnmappedTermsAggregate {
	return *rb.v
}

// Buckets set the Buckets property for UnmappedTermsAggregateBuilder.
func (rb *UnmappedTermsAggregateBuilder) Buckets(buckets BucketsVoid) *UnmappedTermsAggregateBuilder {
	rb.v.Buckets = buckets
	return rb
}

// DocCountErrorUpperBound set the DocCountErrorUpperBound property for UnmappedTermsAggregateBuilder.
func (rb *UnmappedTermsAggregateBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *UnmappedTermsAggregateBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

// Meta set the Meta property for UnmappedTermsAggregateBuilder.
func (rb *UnmappedTermsAggregateBuilder) Meta(meta Metadata) *UnmappedTermsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// SumOtherDocCount set the SumOtherDocCount property for UnmappedTermsAggregateBuilder.
func (rb *UnmappedTermsAggregateBuilder) SumOtherDocCount(sumotherdoccount int64) *UnmappedTermsAggregateBuilder {
	rb.v.SumOtherDocCount = sumotherdoccount
	return rb
}
