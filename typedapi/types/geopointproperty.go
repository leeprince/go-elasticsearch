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
)

// GeoPointProperty type.
type GeoPointProperty struct {
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	IgnoreZValue    *bool                          `json:"ignore_z_value,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	NullValue       *GeoLocation                   `json:"null_value,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// GeoPointPropertyBuilder holds GeoPointProperty struct and provides a builder API.
type GeoPointPropertyBuilder struct {
	v *GeoPointProperty
}

// NewGeoPointProperty provides a builder for the GeoPointProperty struct.
func NewGeoPointProperty() *GeoPointPropertyBuilder {
	r := GeoPointPropertyBuilder{
		&GeoPointProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "geo_point"

	return &r
}

// Build finalize the chain and returns the GeoPointProperty struct
func (rb *GeoPointPropertyBuilder) Build() GeoPointProperty {
	return *rb.v
}

// CopyTo set the CopyTo property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) CopyTo(copyto Fields) *GeoPointPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) DocValues(docvalues bool) *GeoPointPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *GeoPointPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) Fields(value map[PropertyName]Property) *GeoPointPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) IgnoreAbove(ignoreabove int) *GeoPointPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *GeoPointPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// IgnoreZValue set the IgnoreZValue property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) IgnoreZValue(ignorezvalue bool) *GeoPointPropertyBuilder {
	rb.v.IgnoreZValue = &ignorezvalue
	return rb
}

// LocalMetadata set the LocalMetadata property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) LocalMetadata(localmetadata Metadata) *GeoPointPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) Meta(value map[string]string) *GeoPointPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) NullValue(nullvalue GeoLocation) *GeoPointPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// Properties set the Properties property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) Properties(value map[PropertyName]Property) *GeoPointPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) Similarity(similarity string) *GeoPointPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for GeoPointPropertyBuilder.
func (rb *GeoPointPropertyBuilder) Store(store bool) *GeoPointPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for GeoPointPropertyBuilder.
