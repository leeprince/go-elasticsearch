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

// PercentilesBucketAggregate type.
type PercentilesBucketAggregate struct {
	Meta   *Metadata   `json:"meta,omitempty"`
	Values Percentiles `json:"values"`
}

// PercentilesBucketAggregateBuilder holds PercentilesBucketAggregate struct and provides a builder API.
type PercentilesBucketAggregateBuilder struct {
	v *PercentilesBucketAggregate
}

// NewPercentilesBucketAggregate provides a builder for the PercentilesBucketAggregate struct.
func NewPercentilesBucketAggregate() *PercentilesBucketAggregateBuilder {
	r := PercentilesBucketAggregateBuilder{
		&PercentilesBucketAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the PercentilesBucketAggregate struct
func (rb *PercentilesBucketAggregateBuilder) Build() PercentilesBucketAggregate {
	return *rb.v
}

// Meta set the Meta property for PercentilesBucketAggregateBuilder.
func (rb *PercentilesBucketAggregateBuilder) Meta(meta Metadata) *PercentilesBucketAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// Values set the Values property for PercentilesBucketAggregateBuilder.
func (rb *PercentilesBucketAggregateBuilder) Values(values Percentiles) *PercentilesBucketAggregateBuilder {
	rb.v.Values = values
	return rb
}
