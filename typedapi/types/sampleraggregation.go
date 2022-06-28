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

// SamplerAggregation type.
type SamplerAggregation struct {
	Meta      *Metadata `json:"meta,omitempty"`
	Name      *string   `json:"name,omitempty"`
	ShardSize *int      `json:"shard_size,omitempty"`
}

// SamplerAggregationBuilder holds SamplerAggregation struct and provides a builder API.
type SamplerAggregationBuilder struct {
	v *SamplerAggregation
}

// NewSamplerAggregation provides a builder for the SamplerAggregation struct.
func NewSamplerAggregation() *SamplerAggregationBuilder {
	r := SamplerAggregationBuilder{
		&SamplerAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SamplerAggregation struct
func (rb *SamplerAggregationBuilder) Build() SamplerAggregation {
	return *rb.v
}

// Meta set the Meta property for SamplerAggregationBuilder.
func (rb *SamplerAggregationBuilder) Meta(meta Metadata) *SamplerAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for SamplerAggregationBuilder.
func (rb *SamplerAggregationBuilder) Name(name string) *SamplerAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// ShardSize set the ShardSize property for SamplerAggregationBuilder.
func (rb *SamplerAggregationBuilder) ShardSize(shardsize int) *SamplerAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}
