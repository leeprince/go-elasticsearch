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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
)

// FlattenedProperty type.
type FlattenedProperty struct {
	Boost                    *float64                       `json:"boost,omitempty"`
	DepthLimit               *int                           `json:"depth_limit,omitempty"`
	DocValues                *bool                          `json:"doc_values,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals      *bool                          `json:"eager_global_ordinals,omitempty"`
	Fields                   map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`
	IndexOptions             *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	LocalMetadata            *Metadata                      `json:"local_metadata,omitempty"`
	Meta                     map[string]string              `json:"meta,omitempty"`
	NullValue                *string                        `json:"null_value,omitempty"`
	Properties               map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity               *string                        `json:"similarity,omitempty"`
	SplitQueriesOnWhitespace *bool                          `json:"split_queries_on_whitespace,omitempty"`
	Type                     string                         `json:"type,omitempty"`
}

// FlattenedPropertyBuilder holds FlattenedProperty struct and provides a builder API.
type FlattenedPropertyBuilder struct {
	v *FlattenedProperty
}

// NewFlattenedProperty provides a builder for the FlattenedProperty struct.
func NewFlattenedProperty() *FlattenedPropertyBuilder {
	r := FlattenedPropertyBuilder{
		&FlattenedProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "flattened"

	return &r
}

// Build finalize the chain and returns the FlattenedProperty struct
func (rb *FlattenedPropertyBuilder) Build() FlattenedProperty {
	return *rb.v
}

// Boost set the Boost property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) Boost(boost float64) *FlattenedPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// DepthLimit set the DepthLimit property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) DepthLimit(depthlimit int) *FlattenedPropertyBuilder {
	rb.v.DepthLimit = &depthlimit
	return rb
}

// DocValues set the DocValues property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) DocValues(docvalues bool) *FlattenedPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *FlattenedPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// EagerGlobalOrdinals set the EagerGlobalOrdinals property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) EagerGlobalOrdinals(eagerglobalordinals bool) *FlattenedPropertyBuilder {
	rb.v.EagerGlobalOrdinals = &eagerglobalordinals
	return rb
}

// Fields set the Fields property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) Fields(value map[PropertyName]Property) *FlattenedPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) IgnoreAbove(ignoreabove int) *FlattenedPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// Index set the Index property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) Index(index bool) *FlattenedPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// IndexOptions set the IndexOptions property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) IndexOptions(indexoptions indexoptions.IndexOptions) *FlattenedPropertyBuilder {
	rb.v.IndexOptions = &indexoptions
	return rb
}

// LocalMetadata set the LocalMetadata property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) LocalMetadata(localmetadata Metadata) *FlattenedPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) Meta(value map[string]string) *FlattenedPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) NullValue(nullvalue string) *FlattenedPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// Properties set the Properties property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) Properties(value map[PropertyName]Property) *FlattenedPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) Similarity(similarity string) *FlattenedPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// SplitQueriesOnWhitespace set the SplitQueriesOnWhitespace property for FlattenedPropertyBuilder.
func (rb *FlattenedPropertyBuilder) SplitQueriesOnWhitespace(splitqueriesonwhitespace bool) *FlattenedPropertyBuilder {
	rb.v.SplitQueriesOnWhitespace = &splitqueriesonwhitespace
	return rb
}

// Type set the Type property for FlattenedPropertyBuilder.
