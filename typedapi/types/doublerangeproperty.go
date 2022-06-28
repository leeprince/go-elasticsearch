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

// DoubleRangeProperty type.
type DoubleRangeProperty struct {
	Boost         *float64                       `json:"boost,omitempty"`
	Coerce        *bool                          `json:"coerce,omitempty"`
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	Index         *bool                          `json:"index,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// DoubleRangePropertyBuilder holds DoubleRangeProperty struct and provides a builder API.
type DoubleRangePropertyBuilder struct {
	v *DoubleRangeProperty
}

// NewDoubleRangeProperty provides a builder for the DoubleRangeProperty struct.
func NewDoubleRangeProperty() *DoubleRangePropertyBuilder {
	r := DoubleRangePropertyBuilder{
		&DoubleRangeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "double_range"

	return &r
}

// Build finalize the chain and returns the DoubleRangeProperty struct
func (rb *DoubleRangePropertyBuilder) Build() DoubleRangeProperty {
	return *rb.v
}

// Boost set the Boost property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Boost(boost float64) *DoubleRangePropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// Coerce set the Coerce property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Coerce(coerce bool) *DoubleRangePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) CopyTo(copyto Fields) *DoubleRangePropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) DocValues(docvalues bool) *DoubleRangePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DoubleRangePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Fields(value map[PropertyName]Property) *DoubleRangePropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) IgnoreAbove(ignoreabove int) *DoubleRangePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// Index set the Index property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Index(index bool) *DoubleRangePropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) LocalMetadata(localmetadata Metadata) *DoubleRangePropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Meta(value map[string]string) *DoubleRangePropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Properties(value map[PropertyName]Property) *DoubleRangePropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Similarity(similarity string) *DoubleRangePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for DoubleRangePropertyBuilder.
func (rb *DoubleRangePropertyBuilder) Store(store bool) *DoubleRangePropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for DoubleRangePropertyBuilder.
