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

// SamplerAggregate type.
type SamplerAggregate struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Meta         *Metadata                   `json:"meta,omitempty"`
}

// SamplerAggregateBuilder holds SamplerAggregate struct and provides a builder API.
type SamplerAggregateBuilder struct {
	v *SamplerAggregate
}

// NewSamplerAggregate provides a builder for the SamplerAggregate struct.
func NewSamplerAggregate() *SamplerAggregateBuilder {
	r := SamplerAggregateBuilder{
		&SamplerAggregate{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SamplerAggregate struct
func (rb *SamplerAggregateBuilder) Build() SamplerAggregate {
	return *rb.v
}

// Aggregations set the Aggregations property for SamplerAggregateBuilder.
func (rb *SamplerAggregateBuilder) Aggregations(value map[AggregateName]Aggregate) *SamplerAggregateBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for SamplerAggregateBuilder.
func (rb *SamplerAggregateBuilder) DocCount(doccount int64) *SamplerAggregateBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Meta set the Meta property for SamplerAggregateBuilder.
func (rb *SamplerAggregateBuilder) Meta(meta Metadata) *SamplerAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
