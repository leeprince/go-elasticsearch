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

// NestedProperty type.
type NestedProperty struct {
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Enabled         *bool                          `json:"enabled,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IncludeInParent *bool                          `json:"include_in_parent,omitempty"`
	IncludeInRoot   *bool                          `json:"include_in_root,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// NestedPropertyBuilder holds NestedProperty struct and provides a builder API.
type NestedPropertyBuilder struct {
	v *NestedProperty
}

// NewNestedProperty provides a builder for the NestedProperty struct.
func NewNestedProperty() *NestedPropertyBuilder {
	r := NestedPropertyBuilder{
		&NestedProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "nested"

	return &r
}

// Build finalize the chain and returns the NestedProperty struct
func (rb *NestedPropertyBuilder) Build() NestedProperty {
	return *rb.v
}

// CopyTo set the CopyTo property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) CopyTo(copyto Fields) *NestedPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// Dynamic set the Dynamic property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *NestedPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Enabled set the Enabled property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) Enabled(enabled bool) *NestedPropertyBuilder {
	rb.v.Enabled = &enabled
	return rb
}

// Fields set the Fields property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) Fields(value map[PropertyName]Property) *NestedPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) IgnoreAbove(ignoreabove int) *NestedPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IncludeInParent set the IncludeInParent property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) IncludeInParent(includeinparent bool) *NestedPropertyBuilder {
	rb.v.IncludeInParent = &includeinparent
	return rb
}

// IncludeInRoot set the IncludeInRoot property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) IncludeInRoot(includeinroot bool) *NestedPropertyBuilder {
	rb.v.IncludeInRoot = &includeinroot
	return rb
}

// LocalMetadata set the LocalMetadata property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) LocalMetadata(localmetadata Metadata) *NestedPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) Meta(value map[string]string) *NestedPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) Properties(value map[PropertyName]Property) *NestedPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) Similarity(similarity string) *NestedPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for NestedPropertyBuilder.
func (rb *NestedPropertyBuilder) Store(store bool) *NestedPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for NestedPropertyBuilder.
