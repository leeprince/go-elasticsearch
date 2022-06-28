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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// GeoDistanceSort type.
type GeoDistanceSort struct {
	DistanceType    *geodistancetype.GeoDistanceType `json:"distance_type,omitempty"`
	GeoDistanceSort map[Field][]GeoLocation          `json:"GeoDistanceSort,omitempty"`
	IgnoreUnmapped  *bool                            `json:"ignore_unmapped,omitempty"`
	Mode            *sortmode.SortMode               `json:"mode,omitempty"`
	Order           *sortorder.SortOrder             `json:"order,omitempty"`
	Unit            *distanceunit.DistanceUnit       `json:"unit,omitempty"`
}

// GeoDistanceSortBuilder holds GeoDistanceSort struct and provides a builder API.
type GeoDistanceSortBuilder struct {
	v *GeoDistanceSort
}

// NewGeoDistanceSort provides a builder for the GeoDistanceSort struct.
func NewGeoDistanceSort() *GeoDistanceSortBuilder {
	r := GeoDistanceSortBuilder{
		&GeoDistanceSort{
			GeoDistanceSort: make(map[Field][]GeoLocation, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoDistanceSort struct
func (rb *GeoDistanceSortBuilder) Build() GeoDistanceSort {
	return *rb.v
}

// DistanceType set the DistanceType property for GeoDistanceSortBuilder.
func (rb *GeoDistanceSortBuilder) DistanceType(distancetype geodistancetype.GeoDistanceType) *GeoDistanceSortBuilder {
	rb.v.DistanceType = &distancetype
	return rb
}

// GeoDistanceSort set the GeoDistanceSort property for GeoDistanceSortBuilder.
//TODO: Union in map Not implemented yet: UnionOf
func (rb *GeoDistanceSortBuilder) GeoDistanceSort(value map[Field][]GeoLocation) *GeoDistanceSortBuilder {
	rb.v.GeoDistanceSort = value
	return rb
}

// IgnoreUnmapped set the IgnoreUnmapped property for GeoDistanceSortBuilder.
func (rb *GeoDistanceSortBuilder) IgnoreUnmapped(ignoreunmapped bool) *GeoDistanceSortBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

// Mode set the Mode property for GeoDistanceSortBuilder.
func (rb *GeoDistanceSortBuilder) Mode(mode sortmode.SortMode) *GeoDistanceSortBuilder {
	rb.v.Mode = &mode
	return rb
}

// Order set the Order property for GeoDistanceSortBuilder.
func (rb *GeoDistanceSortBuilder) Order(order sortorder.SortOrder) *GeoDistanceSortBuilder {
	rb.v.Order = &order
	return rb
}

// Unit set the Unit property for GeoDistanceSortBuilder.
func (rb *GeoDistanceSortBuilder) Unit(unit distanceunit.DistanceUnit) *GeoDistanceSortBuilder {
	rb.v.Unit = &unit
	return rb
}
