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

// DocValuesPropertyBase type.
type DocValuesPropertyBase struct {
	CopyTo        *Fields                        `json:"copy_to,omitempty"`
	DocValues     *bool                          `json:"doc_values,omitempty"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Store         *bool                          `json:"store,omitempty"`
}

// DocValuesPropertyBaseBuilder holds DocValuesPropertyBase struct and provides a builder API.
type DocValuesPropertyBaseBuilder struct {
	v *DocValuesPropertyBase
}

// NewDocValuesPropertyBase provides a builder for the DocValuesPropertyBase struct.
func NewDocValuesPropertyBase() *DocValuesPropertyBaseBuilder {
	r := DocValuesPropertyBaseBuilder{
		&DocValuesPropertyBase{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DocValuesPropertyBase struct
func (rb *DocValuesPropertyBaseBuilder) Build() DocValuesPropertyBase {
	return *rb.v
}

// CopyTo set the CopyTo property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) CopyTo(copyto Fields) *DocValuesPropertyBaseBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) DocValues(docvalues bool) *DocValuesPropertyBaseBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DocValuesPropertyBaseBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) Fields(value map[PropertyName]Property) *DocValuesPropertyBaseBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) IgnoreAbove(ignoreabove int) *DocValuesPropertyBaseBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) LocalMetadata(localmetadata Metadata) *DocValuesPropertyBaseBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) Meta(value map[string]string) *DocValuesPropertyBaseBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) Properties(value map[PropertyName]Property) *DocValuesPropertyBaseBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) Similarity(similarity string) *DocValuesPropertyBaseBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for DocValuesPropertyBaseBuilder.
func (rb *DocValuesPropertyBaseBuilder) Store(store bool) *DocValuesPropertyBaseBuilder {
	rb.v.Store = &store
	return rb
}
