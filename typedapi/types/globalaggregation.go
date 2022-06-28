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

// GlobalAggregation type.
type GlobalAggregation struct {
	Meta *Metadata `json:"meta,omitempty"`
	Name *string   `json:"name,omitempty"`
}

// GlobalAggregationBuilder holds GlobalAggregation struct and provides a builder API.
type GlobalAggregationBuilder struct {
	v *GlobalAggregation
}

// NewGlobalAggregation provides a builder for the GlobalAggregation struct.
func NewGlobalAggregation() *GlobalAggregationBuilder {
	r := GlobalAggregationBuilder{
		&GlobalAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the GlobalAggregation struct
func (rb *GlobalAggregationBuilder) Build() GlobalAggregation {
	return *rb.v
}

// Meta set the Meta property for GlobalAggregationBuilder.
func (rb *GlobalAggregationBuilder) Meta(meta Metadata) *GlobalAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for GlobalAggregationBuilder.
func (rb *GlobalAggregationBuilder) Name(name string) *GlobalAggregationBuilder {
	rb.v.Name = &name
	return rb
}
