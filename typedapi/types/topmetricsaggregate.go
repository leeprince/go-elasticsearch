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

// TopMetricsAggregate type.
type TopMetricsAggregate struct {
	Meta *Metadata    `json:"meta,omitempty"`
	Top  []TopMetrics `json:"top"`
}

// TopMetricsAggregateBuilder holds TopMetricsAggregate struct and provides a builder API.
type TopMetricsAggregateBuilder struct {
	v *TopMetricsAggregate
}

// NewTopMetricsAggregate provides a builder for the TopMetricsAggregate struct.
func NewTopMetricsAggregate() *TopMetricsAggregateBuilder {
	r := TopMetricsAggregateBuilder{
		&TopMetricsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the TopMetricsAggregate struct
func (rb *TopMetricsAggregateBuilder) Build() TopMetricsAggregate {
	return *rb.v
}

// Meta set the Meta property for TopMetricsAggregateBuilder.
func (rb *TopMetricsAggregateBuilder) Meta(meta Metadata) *TopMetricsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// Top set the Top property for TopMetricsAggregateBuilder.
func (rb *TopMetricsAggregateBuilder) Top(top ...TopMetrics) *TopMetricsAggregateBuilder {
	rb.v.Top = append(rb.v.Top, top...)
	return rb
}
