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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// GeoLineAggregation type.
type GeoLineAggregation struct {
	IncludeSort *bool                `json:"include_sort,omitempty"`
	Point       GeoLinePoint         `json:"point"`
	Size        *int                 `json:"size,omitempty"`
	Sort        GeoLineSort          `json:"sort"`
	SortOrder   *sortorder.SortOrder `json:"sort_order,omitempty"`
}

// GeoLineAggregationBuilder holds GeoLineAggregation struct and provides a builder API.
type GeoLineAggregationBuilder struct {
	v *GeoLineAggregation
}

// NewGeoLineAggregation provides a builder for the GeoLineAggregation struct.
func NewGeoLineAggregation() *GeoLineAggregationBuilder {
	r := GeoLineAggregationBuilder{
		&GeoLineAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the GeoLineAggregation struct
func (rb *GeoLineAggregationBuilder) Build() GeoLineAggregation {
	return *rb.v
}

// IncludeSort set the IncludeSort property for GeoLineAggregationBuilder.
func (rb *GeoLineAggregationBuilder) IncludeSort(includesort bool) *GeoLineAggregationBuilder {
	rb.v.IncludeSort = &includesort
	return rb
}

// Point set the Point property for GeoLineAggregationBuilder.
func (rb *GeoLineAggregationBuilder) Point(point GeoLinePoint) *GeoLineAggregationBuilder {
	rb.v.Point = point
	return rb
}

// Size set the Size property for GeoLineAggregationBuilder.
func (rb *GeoLineAggregationBuilder) Size(size int) *GeoLineAggregationBuilder {
	rb.v.Size = &size
	return rb
}

// Sort set the Sort property for GeoLineAggregationBuilder.
func (rb *GeoLineAggregationBuilder) Sort(sort GeoLineSort) *GeoLineAggregationBuilder {
	rb.v.Sort = sort
	return rb
}

// SortOrder set the SortOrder property for GeoLineAggregationBuilder.
func (rb *GeoLineAggregationBuilder) SortOrder(sortorder sortorder.SortOrder) *GeoLineAggregationBuilder {
	rb.v.SortOrder = &sortorder
	return rb
}
