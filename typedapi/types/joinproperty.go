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

// JoinProperty type.
type JoinProperty struct {
	Dynamic       *dynamicmapping.DynamicMapping  `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property       `json:"fields,omitempty"`
	IgnoreAbove   *int                            `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                       `json:"local_metadata,omitempty"`
	Meta          map[string]string               `json:"meta,omitempty"`
	Properties    map[PropertyName]Property       `json:"properties,omitempty"`
	Relations     map[RelationName][]RelationName `json:"relations,omitempty"`
	Type          string                          `json:"type,omitempty"`
}

// JoinPropertyBuilder holds JoinProperty struct and provides a builder API.
type JoinPropertyBuilder struct {
	v *JoinProperty
}

// NewJoinProperty provides a builder for the JoinProperty struct.
func NewJoinProperty() *JoinPropertyBuilder {
	r := JoinPropertyBuilder{
		&JoinProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
			Relations:  make(map[RelationName][]RelationName, 0),
		},
	}

	r.v.Type = "join"

	return &r
}

// Build finalize the chain and returns the JoinProperty struct
func (rb *JoinPropertyBuilder) Build() JoinProperty {
	return *rb.v
}

// Dynamic set the Dynamic property for JoinPropertyBuilder.
func (rb *JoinPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *JoinPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for JoinPropertyBuilder.
func (rb *JoinPropertyBuilder) Fields(value map[PropertyName]Property) *JoinPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for JoinPropertyBuilder.
func (rb *JoinPropertyBuilder) IgnoreAbove(ignoreabove int) *JoinPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for JoinPropertyBuilder.
func (rb *JoinPropertyBuilder) LocalMetadata(localmetadata Metadata) *JoinPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for JoinPropertyBuilder.
func (rb *JoinPropertyBuilder) Meta(value map[string]string) *JoinPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for JoinPropertyBuilder.
func (rb *JoinPropertyBuilder) Properties(value map[PropertyName]Property) *JoinPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Relations set the Relations property for JoinPropertyBuilder.
//TODO: Union in map Not implemented yet: UnionOf
func (rb *JoinPropertyBuilder) Relations(value map[RelationName][]RelationName) *JoinPropertyBuilder {
	rb.v.Relations = value
	return rb
}

// Type set the Type property for JoinPropertyBuilder.
