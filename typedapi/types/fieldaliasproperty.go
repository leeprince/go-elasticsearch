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

// FieldAliasProperty type.
type FieldAliasProperty struct {
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Path          *Field                         `json:"path,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// FieldAliasPropertyBuilder holds FieldAliasProperty struct and provides a builder API.
type FieldAliasPropertyBuilder struct {
	v *FieldAliasProperty
}

// NewFieldAliasProperty provides a builder for the FieldAliasProperty struct.
func NewFieldAliasProperty() *FieldAliasPropertyBuilder {
	r := FieldAliasPropertyBuilder{
		&FieldAliasProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "alias"

	return &r
}

// Build finalize the chain and returns the FieldAliasProperty struct
func (rb *FieldAliasPropertyBuilder) Build() FieldAliasProperty {
	return *rb.v
}

// Dynamic set the Dynamic property for FieldAliasPropertyBuilder.
func (rb *FieldAliasPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *FieldAliasPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for FieldAliasPropertyBuilder.
func (rb *FieldAliasPropertyBuilder) Fields(value map[PropertyName]Property) *FieldAliasPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for FieldAliasPropertyBuilder.
func (rb *FieldAliasPropertyBuilder) IgnoreAbove(ignoreabove int) *FieldAliasPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for FieldAliasPropertyBuilder.
func (rb *FieldAliasPropertyBuilder) LocalMetadata(localmetadata Metadata) *FieldAliasPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for FieldAliasPropertyBuilder.
func (rb *FieldAliasPropertyBuilder) Meta(value map[string]string) *FieldAliasPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Path set the Path property for FieldAliasPropertyBuilder.
func (rb *FieldAliasPropertyBuilder) Path(path Field) *FieldAliasPropertyBuilder {
	rb.v.Path = &path
	return rb
}

// Properties set the Properties property for FieldAliasPropertyBuilder.
func (rb *FieldAliasPropertyBuilder) Properties(value map[PropertyName]Property) *FieldAliasPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Type set the Type property for FieldAliasPropertyBuilder.
