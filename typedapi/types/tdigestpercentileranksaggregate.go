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

// TDigestPercentileRanksAggregate type.
type TDigestPercentileRanksAggregate struct {
	Meta   *Metadata   `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

// TDigestPercentileRanksAggregateBuilder holds TDigestPercentileRanksAggregate struct and provides a builder API.
type TDigestPercentileRanksAggregateBuilder struct {
	v *TDigestPercentileRanksAggregate
}

// NewTDigestPercentileRanksAggregate provides a builder for the TDigestPercentileRanksAggregate struct.
func NewTDigestPercentileRanksAggregate() *TDigestPercentileRanksAggregateBuilder {
	r := TDigestPercentileRanksAggregateBuilder{
		&TDigestPercentileRanksAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the TDigestPercentileRanksAggregate struct
func (rb *TDigestPercentileRanksAggregateBuilder) Build() TDigestPercentileRanksAggregate {
	return *rb.v
}

// Meta set the Meta property for TDigestPercentileRanksAggregateBuilder.
func (rb *TDigestPercentileRanksAggregateBuilder) Meta(meta Metadata) *TDigestPercentileRanksAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// Values set the Values property for TDigestPercentileRanksAggregateBuilder.
func (rb *TDigestPercentileRanksAggregateBuilder) Values(values Percentiles) *TDigestPercentileRanksAggregateBuilder {
	rb.v.Values = values
	return rb
}
