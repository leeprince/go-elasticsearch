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

// DoubleTermsAggregate type.
type DoubleTermsAggregate struct {
	Buckets                 BucketsDoubleTermsBucket `json:"buckets"`
	DocCountErrorUpperBound *int64                   `json:"doc_count_error_upper_bound,omitempty"`
	Meta                    *Metadata                `json:"meta,omitempty"`
	SumOtherDocCount        int64                    `json:"sum_other_doc_count"`
}

// DoubleTermsAggregateBuilder holds DoubleTermsAggregate struct and provides a builder API.
type DoubleTermsAggregateBuilder struct {
	v *DoubleTermsAggregate
}

// NewDoubleTermsAggregate provides a builder for the DoubleTermsAggregate struct.
func NewDoubleTermsAggregate() *DoubleTermsAggregateBuilder {
	r := DoubleTermsAggregateBuilder{
		&DoubleTermsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the DoubleTermsAggregate struct
func (rb *DoubleTermsAggregateBuilder) Build() DoubleTermsAggregate {
	return *rb.v
}

// Buckets set the Buckets property for DoubleTermsAggregateBuilder.
func (rb *DoubleTermsAggregateBuilder) Buckets(buckets BucketsDoubleTermsBucket) *DoubleTermsAggregateBuilder {
	rb.v.Buckets = buckets
	return rb
}

// DocCountErrorUpperBound set the DocCountErrorUpperBound property for DoubleTermsAggregateBuilder.
func (rb *DoubleTermsAggregateBuilder) DocCountErrorUpperBound(doccounterrorupperbound int64) *DoubleTermsAggregateBuilder {
	rb.v.DocCountErrorUpperBound = &doccounterrorupperbound
	return rb
}

// Meta set the Meta property for DoubleTermsAggregateBuilder.
func (rb *DoubleTermsAggregateBuilder) Meta(meta Metadata) *DoubleTermsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// SumOtherDocCount set the SumOtherDocCount property for DoubleTermsAggregateBuilder.
func (rb *DoubleTermsAggregateBuilder) SumOtherDocCount(sumotherdoccount int64) *DoubleTermsAggregateBuilder {
	rb.v.SumOtherDocCount = sumotherdoccount
	return rb
}
