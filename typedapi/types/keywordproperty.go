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

// KeywordProperty type.
type KeywordProperty struct {
	Boost                    *float64                       `json:"boost,omitempty"`
	CopyTo                   *Fields                        `json:"copy_to,omitempty"`
	DocValues                *bool                          `json:"doc_values,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EagerGlobalOrdinals      *bool                          `json:"eager_global_ordinals,omitempty"`
	Fields                   map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`
	IndexOptions             *indexoptions.IndexOptions     `json:"index_options,omitempty"`
	LocalMetadata            *Metadata                      `json:"local_metadata,omitempty"`
	Meta                     map[string]string              `json:"meta,omitempty"`
	Normalizer               *string                        `json:"normalizer,omitempty"`
	Norms                    *bool                          `json:"norms,omitempty"`
	NullValue                *string                        `json:"null_value,omitempty"`
	Properties               map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity               *string                        `json:"similarity,omitempty"`
	SplitQueriesOnWhitespace *bool                          `json:"split_queries_on_whitespace,omitempty"`
	Store                    *bool                          `json:"store,omitempty"`
	// TimeSeriesDimension [experimental] For internal use by Elastic only. Marks the field as a time
	// series dimension. Defaults to false.
	TimeSeriesDimension *bool  `json:"time_series_dimension,omitempty"`
	Type                string `json:"type,omitempty"`
}

// KeywordPropertyBuilder holds KeywordProperty struct and provides a builder API.
type KeywordPropertyBuilder struct {
	v *KeywordProperty
}

// NewKeywordProperty provides a builder for the KeywordProperty struct.
func NewKeywordProperty() *KeywordPropertyBuilder {
	r := KeywordPropertyBuilder{
		&KeywordProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "keyword"

	return &r
}

// Build finalize the chain and returns the KeywordProperty struct
func (rb *KeywordPropertyBuilder) Build() KeywordProperty {
	return *rb.v
}

// Boost set the Boost property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Boost(boost float64) *KeywordPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// CopyTo set the CopyTo property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) CopyTo(copyto Fields) *KeywordPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) DocValues(docvalues bool) *KeywordPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *KeywordPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// EagerGlobalOrdinals set the EagerGlobalOrdinals property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) EagerGlobalOrdinals(eagerglobalordinals bool) *KeywordPropertyBuilder {
	rb.v.EagerGlobalOrdinals = &eagerglobalordinals
	return rb
}

// Fields set the Fields property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Fields(value map[PropertyName]Property) *KeywordPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) IgnoreAbove(ignoreabove int) *KeywordPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// Index set the Index property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Index(index bool) *KeywordPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// IndexOptions set the IndexOptions property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) IndexOptions(indexoptions indexoptions.IndexOptions) *KeywordPropertyBuilder {
	rb.v.IndexOptions = &indexoptions
	return rb
}

// LocalMetadata set the LocalMetadata property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) LocalMetadata(localmetadata Metadata) *KeywordPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Meta(value map[string]string) *KeywordPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Normalizer set the Normalizer property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Normalizer(normalizer string) *KeywordPropertyBuilder {
	rb.v.Normalizer = &normalizer
	return rb
}

// Norms set the Norms property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Norms(norms bool) *KeywordPropertyBuilder {
	rb.v.Norms = &norms
	return rb
}

// NullValue set the NullValue property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) NullValue(nullvalue string) *KeywordPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// Properties set the Properties property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Properties(value map[PropertyName]Property) *KeywordPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Similarity(similarity string) *KeywordPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// SplitQueriesOnWhitespace set the SplitQueriesOnWhitespace property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) SplitQueriesOnWhitespace(splitqueriesonwhitespace bool) *KeywordPropertyBuilder {
	rb.v.SplitQueriesOnWhitespace = &splitqueriesonwhitespace
	return rb
}

// Store set the Store property for KeywordPropertyBuilder.
func (rb *KeywordPropertyBuilder) Store(store bool) *KeywordPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesDimension [experimental] For internal use by Elastic only. Marks the field as a time
// series dimension. Defaults to false.
func (rb *KeywordPropertyBuilder) TimeSeriesDimension(timeseriesdimension bool) *KeywordPropertyBuilder {
	rb.v.TimeSeriesDimension = &timeseriesdimension
	return rb
}

// Type set the Type property for KeywordPropertyBuilder.
