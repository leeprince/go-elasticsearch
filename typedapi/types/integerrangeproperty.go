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

// IntegerRangeProperty type.
type IntegerRangeProperty struct {
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

// IntegerRangePropertyBuilder holds IntegerRangeProperty struct and provides a builder API.
type IntegerRangePropertyBuilder struct {
	v *IntegerRangeProperty
}

// NewIntegerRangeProperty provides a builder for the IntegerRangeProperty struct.
func NewIntegerRangeProperty() *IntegerRangePropertyBuilder {
	r := IntegerRangePropertyBuilder{
		&IntegerRangeProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "integer_range"

	return &r
}

// Build finalize the chain and returns the IntegerRangeProperty struct
func (rb *IntegerRangePropertyBuilder) Build() IntegerRangeProperty {
	return *rb.v
}

// Boost set the Boost property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Boost(boost float64) *IntegerRangePropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// Coerce set the Coerce property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Coerce(coerce bool) *IntegerRangePropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) CopyTo(copyto Fields) *IntegerRangePropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) DocValues(docvalues bool) *IntegerRangePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *IntegerRangePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Fields(value map[PropertyName]Property) *IntegerRangePropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) IgnoreAbove(ignoreabove int) *IntegerRangePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// Index set the Index property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Index(index bool) *IntegerRangePropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) LocalMetadata(localmetadata Metadata) *IntegerRangePropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Meta(value map[string]string) *IntegerRangePropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Properties(value map[PropertyName]Property) *IntegerRangePropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Similarity(similarity string) *IntegerRangePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for IntegerRangePropertyBuilder.
func (rb *IntegerRangePropertyBuilder) Store(store bool) *IntegerRangePropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for IntegerRangePropertyBuilder.
