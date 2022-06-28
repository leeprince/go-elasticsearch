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

// WildcardProperty type.
type WildcardProperty struct {
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	NullValue     *string                        `json:"null_value,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// WildcardPropertyBuilder holds WildcardProperty struct and provides a builder API.
type WildcardPropertyBuilder struct {
	v *WildcardProperty
}

// NewWildcardProperty provides a builder for the WildcardProperty struct.
func NewWildcardProperty() *WildcardPropertyBuilder {
	r := WildcardPropertyBuilder{
		&WildcardProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "wildcard"

	return &r
}

// Build finalize the chain and returns the WildcardProperty struct
func (rb *WildcardPropertyBuilder) Build() WildcardProperty {
	return *rb.v
}

// CopyTo set the CopyTo property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) CopyTo(copyto Fields) *WildcardPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) DocValues(docvalues bool) *WildcardPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *WildcardPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) Fields(value map[PropertyName]Property) *WildcardPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) IgnoreAbove(ignoreabove int) *WildcardPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) LocalMetadata(localmetadata Metadata) *WildcardPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) Meta(value map[string]string) *WildcardPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) NullValue(nullvalue string) *WildcardPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// Properties set the Properties property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) Properties(value map[PropertyName]Property) *WildcardPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) Similarity(similarity string) *WildcardPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for WildcardPropertyBuilder.
func (rb *WildcardPropertyBuilder) Store(store bool) *WildcardPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for WildcardPropertyBuilder.
