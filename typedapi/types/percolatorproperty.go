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

// PercolatorProperty type.
type PercolatorProperty struct {
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// PercolatorPropertyBuilder holds PercolatorProperty struct and provides a builder API.
type PercolatorPropertyBuilder struct {
	v *PercolatorProperty
}

// NewPercolatorProperty provides a builder for the PercolatorProperty struct.
func NewPercolatorProperty() *PercolatorPropertyBuilder {
	r := PercolatorPropertyBuilder{
		&PercolatorProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "percolator"

	return &r
}

// Build finalize the chain and returns the PercolatorProperty struct
func (rb *PercolatorPropertyBuilder) Build() PercolatorProperty {
	return *rb.v
}

// Dynamic set the Dynamic property for PercolatorPropertyBuilder.
func (rb *PercolatorPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *PercolatorPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for PercolatorPropertyBuilder.
func (rb *PercolatorPropertyBuilder) Fields(value map[PropertyName]Property) *PercolatorPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for PercolatorPropertyBuilder.
func (rb *PercolatorPropertyBuilder) IgnoreAbove(ignoreabove int) *PercolatorPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for PercolatorPropertyBuilder.
func (rb *PercolatorPropertyBuilder) LocalMetadata(localmetadata Metadata) *PercolatorPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for PercolatorPropertyBuilder.
func (rb *PercolatorPropertyBuilder) Meta(value map[string]string) *PercolatorPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for PercolatorPropertyBuilder.
func (rb *PercolatorPropertyBuilder) Properties(value map[PropertyName]Property) *PercolatorPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Type set the Type property for PercolatorPropertyBuilder.
