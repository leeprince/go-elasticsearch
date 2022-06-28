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

// RangePropertyBase type.
type RangePropertyBase struct {
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
}

// RangePropertyBaseBuilder holds RangePropertyBase struct and provides a builder API.
type RangePropertyBaseBuilder struct {
	v *RangePropertyBase
}

// NewRangePropertyBase provides a builder for the RangePropertyBase struct.
func NewRangePropertyBase() *RangePropertyBaseBuilder {
	r := RangePropertyBaseBuilder{
		&RangePropertyBase{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RangePropertyBase struct
func (rb *RangePropertyBaseBuilder) Build() RangePropertyBase {
	return *rb.v
}

// Boost set the Boost property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Boost(boost float64) *RangePropertyBaseBuilder {
	rb.v.Boost = &boost
	return rb
}

// Coerce set the Coerce property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Coerce(coerce bool) *RangePropertyBaseBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) CopyTo(copyto Fields) *RangePropertyBaseBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) DocValues(docvalues bool) *RangePropertyBaseBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *RangePropertyBaseBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Fields(value map[PropertyName]Property) *RangePropertyBaseBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) IgnoreAbove(ignoreabove int) *RangePropertyBaseBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// Index set the Index property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Index(index bool) *RangePropertyBaseBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) LocalMetadata(localmetadata Metadata) *RangePropertyBaseBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Meta(value map[string]string) *RangePropertyBaseBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Properties(value map[PropertyName]Property) *RangePropertyBaseBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Similarity(similarity string) *RangePropertyBaseBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for RangePropertyBaseBuilder.
func (rb *RangePropertyBaseBuilder) Store(store bool) *RangePropertyBaseBuilder {
	rb.v.Store = &store
	return rb
}
