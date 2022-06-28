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

// ShapeFieldQuery type.
type ShapeFieldQuery struct {
	IndexedShape *FieldLookup                       `json:"indexed_shape,omitempty"`
	Relation     *geoshaperelation.GeoShapeRelation `json:"relation,omitempty"`
	Shape        *GeoShape                          `json:"shape,omitempty"`
}

// ShapeFieldQueryBuilder holds ShapeFieldQuery struct and provides a builder API.
type ShapeFieldQueryBuilder struct {
	v *ShapeFieldQuery
}

// NewShapeFieldQuery provides a builder for the ShapeFieldQuery struct.
func NewShapeFieldQuery() *ShapeFieldQueryBuilder {
	r := ShapeFieldQueryBuilder{
		&ShapeFieldQuery{},
	}

	return &r
}

// Build finalize the chain and returns the ShapeFieldQuery struct
func (rb *ShapeFieldQueryBuilder) Build() ShapeFieldQuery {
	return *rb.v
}

// IndexedShape set the IndexedShape property for ShapeFieldQueryBuilder.
func (rb *ShapeFieldQueryBuilder) IndexedShape(indexedshape FieldLookup) *ShapeFieldQueryBuilder {
	rb.v.IndexedShape = &indexedshape
	return rb
}

// Relation set the Relation property for ShapeFieldQueryBuilder.
func (rb *ShapeFieldQueryBuilder) Relation(relation geoshaperelation.GeoShapeRelation) *ShapeFieldQueryBuilder {
	rb.v.Relation = &relation
	return rb
}

// Shape set the Shape property for ShapeFieldQueryBuilder.
func (rb *ShapeFieldQueryBuilder) Shape(shape GeoShape) *ShapeFieldQueryBuilder {
	rb.v.Shape = &shape
	return rb
}
