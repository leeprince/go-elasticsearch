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

// DenseVectorProperty type.
type DenseVectorProperty struct {
	Dims          int                            `json:"dims"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	Index         *bool                          `json:"index,omitempty"`
	IndexOptions  *DenseVectorIndexOptions       `json:"index_options,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity    *string                        `json:"similarity,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// DenseVectorPropertyBuilder holds DenseVectorProperty struct and provides a builder API.
type DenseVectorPropertyBuilder struct {
	v *DenseVectorProperty
}

// NewDenseVectorProperty provides a builder for the DenseVectorProperty struct.
func NewDenseVectorProperty() *DenseVectorPropertyBuilder {
	r := DenseVectorPropertyBuilder{
		&DenseVectorProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "dense_vector"

	return &r
}

// Build finalize the chain and returns the DenseVectorProperty struct
func (rb *DenseVectorPropertyBuilder) Build() DenseVectorProperty {
	return *rb.v
}

// Dims set the Dims property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) Dims(dims int) *DenseVectorPropertyBuilder {
	rb.v.Dims = dims
	return rb
}

// Dynamic set the Dynamic property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DenseVectorPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) Fields(value map[PropertyName]Property) *DenseVectorPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) IgnoreAbove(ignoreabove int) *DenseVectorPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// Index set the Index property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) Index(index bool) *DenseVectorPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// IndexOptions set the IndexOptions property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) IndexOptions(indexoptions DenseVectorIndexOptions) *DenseVectorPropertyBuilder {
	rb.v.IndexOptions = &indexoptions
	return rb
}

// LocalMetadata set the LocalMetadata property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) LocalMetadata(localmetadata Metadata) *DenseVectorPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) Meta(value map[string]string) *DenseVectorPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) Properties(value map[PropertyName]Property) *DenseVectorPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for DenseVectorPropertyBuilder.
func (rb *DenseVectorPropertyBuilder) Similarity(similarity string) *DenseVectorPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Type set the Type property for DenseVectorPropertyBuilder.
