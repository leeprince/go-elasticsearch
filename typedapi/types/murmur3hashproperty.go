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

// Murmur3HashProperty type.
type Murmur3HashProperty struct {
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// Murmur3HashPropertyBuilder holds Murmur3HashProperty struct and provides a builder API.
type Murmur3HashPropertyBuilder struct {
	v *Murmur3HashProperty
}

// NewMurmur3HashProperty provides a builder for the Murmur3HashProperty struct.
func NewMurmur3HashProperty() *Murmur3HashPropertyBuilder {
	r := Murmur3HashPropertyBuilder{
		&Murmur3HashProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "murmur3"

	return &r
}

// Build finalize the chain and returns the Murmur3HashProperty struct
func (rb *Murmur3HashPropertyBuilder) Build() Murmur3HashProperty {
	return *rb.v
}

// CopyTo set the CopyTo property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) CopyTo(copyto Fields) *Murmur3HashPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) DocValues(docvalues bool) *Murmur3HashPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *Murmur3HashPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) Fields(value map[PropertyName]Property) *Murmur3HashPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) IgnoreAbove(ignoreabove int) *Murmur3HashPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) LocalMetadata(localmetadata Metadata) *Murmur3HashPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) Meta(value map[string]string) *Murmur3HashPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) Properties(value map[PropertyName]Property) *Murmur3HashPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) Similarity(similarity string) *Murmur3HashPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for Murmur3HashPropertyBuilder.
func (rb *Murmur3HashPropertyBuilder) Store(store bool) *Murmur3HashPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for Murmur3HashPropertyBuilder.
