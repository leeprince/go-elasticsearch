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

// ObjectProperty type.
type ObjectProperty struct {
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled       *bool                          `json:"enabled,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// ObjectPropertyBuilder holds ObjectProperty struct and provides a builder API.
type ObjectPropertyBuilder struct {
	v *ObjectProperty
}

// NewObjectProperty provides a builder for the ObjectProperty struct.
func NewObjectProperty() *ObjectPropertyBuilder {
	r := ObjectPropertyBuilder{
		&ObjectProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "object"

	return &r
}

// Build finalize the chain and returns the ObjectProperty struct
func (rb *ObjectPropertyBuilder) Build() ObjectProperty {
	return *rb.v
}

// CopyTo set the CopyTo property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) CopyTo(copyto Fields) *ObjectPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// Dynamic set the Dynamic property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *ObjectPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Enabled set the Enabled property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) Enabled(enabled bool) *ObjectPropertyBuilder {
	rb.v.Enabled = &enabled
	return rb
}

// Fields set the Fields property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) Fields(value map[PropertyName]Property) *ObjectPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) IgnoreAbove(ignoreabove int) *ObjectPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) LocalMetadata(localmetadata Metadata) *ObjectPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) Meta(value map[string]string) *ObjectPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) Properties(value map[PropertyName]Property) *ObjectPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) Similarity(similarity string) *ObjectPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for ObjectPropertyBuilder.
func (rb *ObjectPropertyBuilder) Store(store bool) *ObjectPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for ObjectPropertyBuilder.
