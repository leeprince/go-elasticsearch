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

// GeoDistanceFeatureQuery type.
type GeoDistanceFeatureQuery struct {
	Boost      *float32    `json:"boost,omitempty"`
	Field      Field       `json:"field"`
	Origin     GeoLocation `json:"origin"`
	Pivot      Distance    `json:"pivot"`
	QueryName_ *string     `json:"_name,omitempty"`
}

// GeoDistanceFeatureQueryBuilder holds GeoDistanceFeatureQuery struct and provides a builder API.
type GeoDistanceFeatureQueryBuilder struct {
	v *GeoDistanceFeatureQuery
}

// NewGeoDistanceFeatureQuery provides a builder for the GeoDistanceFeatureQuery struct.
func NewGeoDistanceFeatureQuery() *GeoDistanceFeatureQueryBuilder {
	r := GeoDistanceFeatureQueryBuilder{
		&GeoDistanceFeatureQuery{},
	}

	return &r
}

// Build finalize the chain and returns the GeoDistanceFeatureQuery struct
func (rb *GeoDistanceFeatureQueryBuilder) Build() GeoDistanceFeatureQuery {
	return *rb.v
}

// Boost set the Boost property for GeoDistanceFeatureQueryBuilder.
func (rb *GeoDistanceFeatureQueryBuilder) Boost(boost float32) *GeoDistanceFeatureQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Field set the Field property for GeoDistanceFeatureQueryBuilder.
func (rb *GeoDistanceFeatureQueryBuilder) Field(field Field) *GeoDistanceFeatureQueryBuilder {
	rb.v.Field = field
	return rb
}

// Origin set the Origin property for GeoDistanceFeatureQueryBuilder.
func (rb *GeoDistanceFeatureQueryBuilder) Origin(origin GeoLocation) *GeoDistanceFeatureQueryBuilder {
	rb.v.Origin = origin
	return rb
}

// Pivot set the Pivot property for GeoDistanceFeatureQueryBuilder.
func (rb *GeoDistanceFeatureQueryBuilder) Pivot(pivot Distance) *GeoDistanceFeatureQueryBuilder {
	rb.v.Pivot = pivot
	return rb
}

// QueryName_ set the QueryName_ property for GeoDistanceFeatureQueryBuilder.
func (rb *GeoDistanceFeatureQueryBuilder) QueryName_(queryname_ string) *GeoDistanceFeatureQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
