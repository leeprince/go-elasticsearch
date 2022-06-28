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

// TermsAggregateBaseStringTermsBucket type.
type TermsAggregateBaseStringTermsBucket struct {
	Buckets                 BucketsStringTermsBucket `json:"buckets"`
	DocCountErrorUpperBound *int64                   `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    *Metadata                `json:"meta,omitempty"`
	SumOtherDocCount        int64                    `json:"sum_other_doc_count"`
}

// TermsAggregateBaseStringTermsBucketBuilder holds TermsAggregateBaseStringTermsBucket struct and provides a builder API.
type TermsAggregateBaseStringTermsBucketBuilder struct {
	v *TermsAggregateBaseStringTermsBucket
}

// NewTermsAggregateBaseStringTermsBucket provides a builder for the TermsAggregateBaseStringTermsBucket struct.
func NewTermsAggregateBaseStringTermsBucket() *TermsAggregateBaseStringTermsBucketBuilder {
	r := TermsAggregateBaseStringTermsBucketBuilder{
		&TermsAggregateBaseStringTermsBucket{},
	}

	return &r
}

// Build finalize the chain and returns the TermsAggregateBaseStringTermsBucket struct
func (rb *TermsAggregateBaseStringTermsBucketBuilder) Build() TermsAggregateBaseStringTermsBucket {
	return *rb.v
}

// Buckets set the Buckets property for TermsAggregateBaseStringTermsBucketBuilder.
func (rb *TermsAggregateBaseStringTermsBucketBuilder) Buckets(buckets BucketsStringTermsBucket) *TermsAggregateBaseStringTermsBucketBuilder {
	rb.v.Buckets = buckets
	return rb
}

// DocCountErrorUpperBound set the DocCountErrorUpperBound property for TermsAggregateBaseStringTermsBucketBuilder.
func (rb *TermsAggregateBaseStringTermsBucketBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *TermsAggregateBaseStringTermsBucketBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

// Meta set the Meta property for TermsAggregateBaseStringTermsBucketBuilder.
func (rb *TermsAggregateBaseStringTermsBucketBuilder) Meta(meta Metadata) *TermsAggregateBaseStringTermsBucketBuilder {
	rb.v.Meta = &meta
	return rb
}

// SumOtherDocCount set the SumOtherDocCount property for TermsAggregateBaseStringTermsBucketBuilder.
func (rb *TermsAggregateBaseStringTermsBucketBuilder) SumOtherDocCount(sumotherdoccount int64) *TermsAggregateBaseStringTermsBucketBuilder {
	rb.v.SumOtherDocCount = sumotherdoccount
	return rb
}
