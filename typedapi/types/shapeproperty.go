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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoorientation"
)

// ShapeProperty type.
type ShapeProperty struct {
	Coerce          *bool                          `json:"coerce,omitempty"`
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	Orientation     *geoorientation.GeoOrientation `json:"orientation,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// ShapePropertyBuilder holds ShapeProperty struct and provides a builder API.
type ShapePropertyBuilder struct {
	v *ShapeProperty
}

// NewShapeProperty provides a builder for the ShapeProperty struct.
func NewShapeProperty() *ShapePropertyBuilder {
	r := ShapePropertyBuilder{
		&ShapeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "shape"

	return &r
}

// Build finalize the chain and returns the ShapeProperty struct
func (rb *ShapePropertyBuilder) Build() ShapeProperty {
	return *rb.v
}

// Coerce set the Coerce property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) Coerce(coerce bool) *ShapePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) CopyTo(copyto Fields) *ShapePropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) DocValues(docvalues bool) *ShapePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *ShapePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) Fields(value map[PropertyName]Property) *ShapePropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) IgnoreAbove(ignoreabove int) *ShapePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) IgnoreMalformed(ignoremalformed bool) *ShapePropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// IgnoreZValue set the IgnoreZValue property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) IgnoreZValue(ignorezvalue bool) *ShapePropertyBuilder {
	rb.v.IgnoreZValue = &ignorezvalue
	return rb
}

// LocalMetadata set the LocalMetadata property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) LocalMetadata(localmetadata Metadata) *ShapePropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) Meta(value map[string]string) *ShapePropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Orientation set the Orientation property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) Orientation(orientation geoorientation.GeoOrientation) *ShapePropertyBuilder {
	rb.v.Orientation = &orientation
	return rb
}

// Properties set the Properties property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) Properties(value map[PropertyName]Property) *ShapePropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) Similarity(similarity string) *ShapePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for ShapePropertyBuilder.
func (rb *ShapePropertyBuilder) Store(store bool) *ShapePropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for ShapePropertyBuilder.
