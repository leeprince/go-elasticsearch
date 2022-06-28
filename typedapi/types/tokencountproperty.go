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

// TokenCountProperty type.
type TokenCountProperty struct {
	Analyzer                 *string                        `json:"analyzer,omitempty"`
	Boost                    *float64                       `json:"boost,omitempty"`
	CopyTo                   *Fields                        `json:"copy_to,omitempty"`
	DocValues                *bool                          `json:"doc_values,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	EnablePositionIncrements *bool                          `json:"enable_position_increments,omitempty"`
	Fields                   map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove              *int                           `json:"ignore_above,omitempty"`
	Index                    *bool                          `json:"index,omitempty"`
	LocalMetadata            *Metadata                      `json:"local_metadata,omitempty"`
	Meta                     map[string]string              `json:"meta,omitempty"`
	NullValue                *float64                       `json:"null_value,omitempty"`
	Properties               map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity               *string                        `json:"similarity,omitempty"`
	Store                    *bool                          `json:"store,omitempty"`
	Type                     string                         `json:"type,omitempty"`
}

// TokenCountPropertyBuilder holds TokenCountProperty struct and provides a builder API.
type TokenCountPropertyBuilder struct {
	v *TokenCountProperty
}

// NewTokenCountProperty provides a builder for the TokenCountProperty struct.
func NewTokenCountProperty() *TokenCountPropertyBuilder {
	r := TokenCountPropertyBuilder{
		&TokenCountProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "token_count"

	return &r
}

// Build finalize the chain and returns the TokenCountProperty struct
func (rb *TokenCountPropertyBuilder) Build() TokenCountProperty {
	return *rb.v
}

// Analyzer set the Analyzer property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Analyzer(analyzer string) *TokenCountPropertyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Boost set the Boost property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Boost(boost float64) *TokenCountPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// CopyTo set the CopyTo property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) CopyTo(copyto Fields) *TokenCountPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) DocValues(docvalues bool) *TokenCountPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *TokenCountPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// EnablePositionIncrements set the EnablePositionIncrements property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) EnablePositionIncrements(enablepositionincrements bool) *TokenCountPropertyBuilder {
	rb.v.EnablePositionIncrements = &enablepositionincrements
	return rb
}

// Fields set the Fields property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Fields(value map[PropertyName]Property) *TokenCountPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) IgnoreAbove(ignoreabove int) *TokenCountPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// Index set the Index property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Index(index bool) *TokenCountPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) LocalMetadata(localmetadata Metadata) *TokenCountPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Meta(value map[string]string) *TokenCountPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) NullValue(nullvalue float64) *TokenCountPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// Properties set the Properties property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Properties(value map[PropertyName]Property) *TokenCountPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Similarity(similarity string) *TokenCountPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for TokenCountPropertyBuilder.
func (rb *TokenCountPropertyBuilder) Store(store bool) *TokenCountPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for TokenCountPropertyBuilder.
