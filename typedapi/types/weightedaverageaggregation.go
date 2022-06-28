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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/valuetype"
)

// WeightedAverageAggregation type.
type WeightedAverageAggregation struct {
	Format    *string               `json:"format,omitempty"`
	Meta      *Metadata             `json:"meta,omitempty"`
	Name      *string               `json:"name,omitempty"`
	Value     *WeightedAverageValue `json:"value,omitempty"`
	ValueType *valuetype.ValueType  `json:"value_type,omitempty"`
	Weight    *WeightedAverageValue `json:"weight,omitempty"`
}

// WeightedAverageAggregationBuilder holds WeightedAverageAggregation struct and provides a builder API.
type WeightedAverageAggregationBuilder struct {
	v *WeightedAverageAggregation
}

// NewWeightedAverageAggregation provides a builder for the WeightedAverageAggregation struct.
func NewWeightedAverageAggregation() *WeightedAverageAggregationBuilder {
	r := WeightedAverageAggregationBuilder{
		&WeightedAverageAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the WeightedAverageAggregation struct
func (rb *WeightedAverageAggregationBuilder) Build() WeightedAverageAggregation {
	return *rb.v
}

// Format set the Format property for WeightedAverageAggregationBuilder.
func (rb *WeightedAverageAggregationBuilder) Format(format string) *WeightedAverageAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Meta set the Meta property for WeightedAverageAggregationBuilder.
func (rb *WeightedAverageAggregationBuilder) Meta(meta Metadata) *WeightedAverageAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for WeightedAverageAggregationBuilder.
func (rb *WeightedAverageAggregationBuilder) Name(name string) *WeightedAverageAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Value set the Value property for WeightedAverageAggregationBuilder.
func (rb *WeightedAverageAggregationBuilder) Value(value WeightedAverageValue) *WeightedAverageAggregationBuilder {
	rb.v.Value = &value
	return rb
}

// ValueType set the ValueType property for WeightedAverageAggregationBuilder.
func (rb *WeightedAverageAggregationBuilder) ValueType(valuetype valuetype.ValueType) *WeightedAverageAggregationBuilder {
	rb.v.ValueType = &valuetype
	return rb
}

// Weight set the Weight property for WeightedAverageAggregationBuilder.
func (rb *WeightedAverageAggregationBuilder) Weight(weight WeightedAverageValue) *WeightedAverageAggregationBuilder {
	rb.v.Weight = &weight
	return rb
}
