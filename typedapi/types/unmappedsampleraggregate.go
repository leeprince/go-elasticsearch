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

// UnmappedSamplerAggregate type.
type UnmappedSamplerAggregate struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Meta         *Metadata                   `json:"meta,omitempty"`
}

// UnmappedSamplerAggregateBuilder holds UnmappedSamplerAggregate struct and provides a builder API.
type UnmappedSamplerAggregateBuilder struct {
	v *UnmappedSamplerAggregate
}

// NewUnmappedSamplerAggregate provides a builder for the UnmappedSamplerAggregate struct.
func NewUnmappedSamplerAggregate() *UnmappedSamplerAggregateBuilder {
	r := UnmappedSamplerAggregateBuilder{
		&UnmappedSamplerAggregate{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the UnmappedSamplerAggregate struct
func (rb *UnmappedSamplerAggregateBuilder) Build() UnmappedSamplerAggregate {
	return *rb.v
}

// Aggregations set the Aggregations property for UnmappedSamplerAggregateBuilder.
func (rb *UnmappedSamplerAggregateBuilder) Aggregations(value map[AggregateName]Aggregate) *UnmappedSamplerAggregateBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for UnmappedSamplerAggregateBuilder.
func (rb *UnmappedSamplerAggregateBuilder) DocCount(doccount int64) *UnmappedSamplerAggregateBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Meta set the Meta property for UnmappedSamplerAggregateBuilder.
func (rb *UnmappedSamplerAggregateBuilder) Meta(meta Metadata) *UnmappedSamplerAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
