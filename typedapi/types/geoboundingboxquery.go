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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoexecution"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geovalidationmethod"
)

// GeoBoundingBoxQuery type.
type GeoBoundingBoxQuery struct {
	Boost               *float32                                 `json:"boost,omitempty"`
	GeoBoundingBoxQuery map[Field]GeoBounds                      `json:"GeoBoundingBoxQuery,omitempty"`
	IgnoreUnmapped      *bool                                    `json:"ignore_unmapped,omitempty"`
	QueryName_          *string                                  `json:"_name,omitempty"`
	Type                *geoexecution.GeoExecution               `json:"type,omitempty"`
	ValidationMethod    *geovalidationmethod.GeoValidationMethod `json:"validation_method,omitempty"`
}

// GeoBoundingBoxQueryBuilder holds GeoBoundingBoxQuery struct and provides a builder API.
type GeoBoundingBoxQueryBuilder struct {
	v *GeoBoundingBoxQuery
}

// NewGeoBoundingBoxQuery provides a builder for the GeoBoundingBoxQuery struct.
func NewGeoBoundingBoxQuery() *GeoBoundingBoxQueryBuilder {
	r := GeoBoundingBoxQueryBuilder{
		&GeoBoundingBoxQuery{
			GeoBoundingBoxQuery: make(map[Field]GeoBounds, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoBoundingBoxQuery struct
func (rb *GeoBoundingBoxQueryBuilder) Build() GeoBoundingBoxQuery {
	return *rb.v
}

// Boost set the Boost property for GeoBoundingBoxQueryBuilder.
func (rb *GeoBoundingBoxQueryBuilder) Boost(boost float32) *GeoBoundingBoxQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// GeoBoundingBoxQuery set the GeoBoundingBoxQuery property for GeoBoundingBoxQueryBuilder.
func (rb *GeoBoundingBoxQueryBuilder) GeoBoundingBoxQuery(value map[Field]GeoBounds) *GeoBoundingBoxQueryBuilder {
	rb.v.GeoBoundingBoxQuery = value
	return rb
}

// IgnoreUnmapped set the IgnoreUnmapped property for GeoBoundingBoxQueryBuilder.
func (rb *GeoBoundingBoxQueryBuilder) IgnoreUnmapped(ignoreunmapped bool) *GeoBoundingBoxQueryBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

// QueryName_ set the QueryName_ property for GeoBoundingBoxQueryBuilder.
func (rb *GeoBoundingBoxQueryBuilder) QueryName_(queryname_ string) *GeoBoundingBoxQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Type set the Type property for GeoBoundingBoxQueryBuilder.
func (rb *GeoBoundingBoxQueryBuilder) Type_(type_ geoexecution.GeoExecution) *GeoBoundingBoxQueryBuilder {
	rb.v.Type = &type_
	return rb
}

// ValidationMethod set the ValidationMethod property for GeoBoundingBoxQueryBuilder.
func (rb *GeoBoundingBoxQueryBuilder) ValidationMethod(validationmethod geovalidationmethod.GeoValidationMethod) *GeoBoundingBoxQueryBuilder {
	rb.v.ValidationMethod = &validationmethod
	return rb
}
