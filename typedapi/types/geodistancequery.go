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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

// GeoDistanceQuery type.
type GeoDistanceQuery struct {
	Boost            *float32                                 `json:"boost,omitempty"`
	Distance         *Distance                                `json:"distance,omitempty"`
	DistanceType     *geodistancetype.GeoDistanceType         `json:"distance_type,omitempty"`
	GeoDistanceQuery map[Field]GeoLocation                    `json:"GeoDistanceQuery,omitempty"`
	QueryName_       *string                                  `json:"_name,omitempty"`
	ValidationMethod *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

// GeoDistanceQueryBuilder holds GeoDistanceQuery struct and provides a builder API.
type GeoDistanceQueryBuilder struct {
	v *GeoDistanceQuery
}

// NewGeoDistanceQuery provides a builder for the GeoDistanceQuery struct.
func NewGeoDistanceQuery() *GeoDistanceQueryBuilder {
	r := GeoDistanceQueryBuilder{
		&GeoDistanceQuery{
			GeoDistanceQuery: make(map[Field]GeoLocation, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoDistanceQuery struct
func (rb *GeoDistanceQueryBuilder) Build() GeoDistanceQuery {
	return *rb.v
}

// Boost set the Boost property for GeoDistanceQueryBuilder.
func (rb *GeoDistanceQueryBuilder) Boost(boost float32) *GeoDistanceQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Distance set the Distance property for GeoDistanceQueryBuilder.
func (rb *GeoDistanceQueryBuilder) Distance(distance Distance) *GeoDistanceQueryBuilder {
	rb.v.Distance = &distance
	return rb
}

// DistanceType set the DistanceType property for GeoDistanceQueryBuilder.
func (rb *GeoDistanceQueryBuilder) DistanceType(distancetype geodistancetype.GeoDistanceType) *GeoDistanceQueryBuilder {
	rb.v.DistanceType = &distancetype
	return rb
}

// GeoDistanceQuery set the GeoDistanceQuery property for GeoDistanceQueryBuilder.
func (rb *GeoDistanceQueryBuilder) GeoDistanceQuery(value map[Field]GeoLocation) *GeoDistanceQueryBuilder {
	rb.v.GeoDistanceQuery = value
	return rb
}

// QueryName_ set the QueryName_ property for GeoDistanceQueryBuilder.
func (rb *GeoDistanceQueryBuilder) QueryName_(queryname_ string) *GeoDistanceQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// ValidationMethod set the ValidationMethod property for GeoDistanceQueryBuilder.
func (rb *GeoDistanceQueryBuilder) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *GeoDistanceQueryBuilder {
	rb.v.ValidationMethod = &validationmethod
	return rb
}
