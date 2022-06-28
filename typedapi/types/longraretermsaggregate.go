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

// LongRareTermsAggregate type.
type LongRareTermsAggregate struct {
	Buckets BucketsLongRareTermsBucket `json:"buckets"`
	Meta    *Metadata                  `json:"meta,omitempty"`
}

// LongRareTermsAggregateBuilder holds LongRareTermsAggregate struct and provides a builder API.
type LongRareTermsAggregateBuilder struct {
	v *LongRareTermsAggregate
}

// NewLongRareTermsAggregate provides a builder for the LongRareTermsAggregate struct.
func NewLongRareTermsAggregate() *LongRareTermsAggregateBuilder {
	r := LongRareTermsAggregateBuilder{
		&LongRareTermsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the LongRareTermsAggregate struct
func (rb *LongRareTermsAggregateBuilder) Build() LongRareTermsAggregate {
	return *rb.v
}

// Buckets set the Buckets property for LongRareTermsAggregateBuilder.
func (rb *LongRareTermsAggregateBuilder) Buckets(buckets BucketsLongRareTermsBucket) *LongRareTermsAggregateBuilder {
	rb.v.Buckets = buckets
	return rb
}

// Meta set the Meta property for LongRareTermsAggregateBuilder.
func (rb *LongRareTermsAggregateBuilder) Meta(meta Metadata) *LongRareTermsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
