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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
)

// GeoShapeFieldQuery type.
type GeoShapeFieldQuery struct {
	IndexedShape *FieldLookup                       `json:"indexed_shape,omitempty"`
	Relation     *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`
	Shape        *GeoShape                          `json:"shape,omitempty"`
}

// GeoShapeFieldQueryBuilder holds GeoShapeFieldQuery struct and provides a builder API.
type GeoShapeFieldQueryBuilder struct {
	v *GeoShapeFieldQuery
}

// NewGeoShapeFieldQuery provides a builder for the GeoShapeFieldQuery struct.
func NewGeoShapeFieldQuery() *GeoShapeFieldQueryBuilder {
	r := GeoShapeFieldQueryBuilder{
		&GeoShapeFieldQuery{},
	}

	return &r
}

// Build finalize the chain and returns the GeoShapeFieldQuery struct
func (rb *GeoShapeFieldQueryBuilder) Build() GeoShapeFieldQuery {
	return *rb.v
}

// IndexedShape set the IndexedShape property for GeoShapeFieldQueryBuilder.
func (rb *GeoShapeFieldQueryBuilder) IndexedShape(indexedshape FieldLookup) *GeoShapeFieldQueryBuilder {
	rb.v.IndexedShape = &indexedshape
	return rb
}

// Relation set the Relation property for GeoShapeFieldQueryBuilder.
func (rb *GeoShapeFieldQueryBuilder) Relation(relation geoshaperelation.GeoShapeRelation) *GeoShapeFieldQueryBuilder {
	rb.v.Relation = &relation
	return rb
}

// Shape set the Shape property for GeoShapeFieldQueryBuilder.
func (rb *GeoShapeFieldQueryBuilder) Shape(shape GeoShape) *GeoShapeFieldQueryBuilder {
	rb.v.Shape = &shape
	return rb
}
