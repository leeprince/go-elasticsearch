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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geostrategy"
)

// GeoShapeProperty type.
type GeoShapeProperty struct {
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
	Strategy        *geostrategy.GeoStrategy       `json:"strategy,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// GeoShapePropertyBuilder holds GeoShapeProperty struct and provides a builder API.
type GeoShapePropertyBuilder struct {
	v *GeoShapeProperty
}

// NewGeoShapeProperty provides a builder for the GeoShapeProperty struct.
func NewGeoShapeProperty() *GeoShapePropertyBuilder {
	r := GeoShapePropertyBuilder{
		&GeoShapeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "geo_shape"

	return &r
}

// Build finalize the chain and returns the GeoShapeProperty struct
func (rb *GeoShapePropertyBuilder) Build() GeoShapeProperty {
	return *rb.v
}

// Coerce set the Coerce property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Coerce(coerce bool) *GeoShapePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) CopyTo(copyto Fields) *GeoShapePropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) DocValues(docvalues bool) *GeoShapePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *GeoShapePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Fields(value map[PropertyName]Property) *GeoShapePropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) IgnoreAbove(ignoreabove int) *GeoShapePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) IgnoreMalformed(ignoremalformed bool) *GeoShapePropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// IgnoreZValue set the IgnoreZValue property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) IgnoreZValue(ignorezvalue bool) *GeoShapePropertyBuilder {
	rb.v.IgnoreZValue = &ignorezvalue
	return rb
}

// LocalMetadata set the LocalMetadata property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) LocalMetadata(localmetadata Metadata) *GeoShapePropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Meta(value map[string]string) *GeoShapePropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Orientation set the Orientation property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Orientation(orientation geoorientation.GeoOrientation) *GeoShapePropertyBuilder {
	rb.v.Orientation = &orientation
	return rb
}

// Properties set the Properties property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Properties(value map[PropertyName]Property) *GeoShapePropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Similarity(similarity string) *GeoShapePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Store(store bool) *GeoShapePropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Strategy set the Strategy property for GeoShapePropertyBuilder.
func (rb *GeoShapePropertyBuilder) Strategy(strategy geostrategy.GeoStrategy) *GeoShapePropertyBuilder {
	rb.v.Strategy = &strategy
	return rb
}

// Type set the Type property for GeoShapePropertyBuilder.
