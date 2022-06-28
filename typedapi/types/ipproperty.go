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

// IpProperty type.
type IpProperty struct {
	Boost           *float64                       `json:"boost,omitempty"`
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	Index           *bool                          `json:"index,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	NullValue       *string                        `json:"null_value,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// IpPropertyBuilder holds IpProperty struct and provides a builder API.
type IpPropertyBuilder struct {
	v *IpProperty
}

// NewIpProperty provides a builder for the IpProperty struct.
func NewIpProperty() *IpPropertyBuilder {
	r := IpPropertyBuilder{
		&IpProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "ip"

	return &r
}

// Build finalize the chain and returns the IpProperty struct
func (rb *IpPropertyBuilder) Build() IpProperty {
	return *rb.v
}

// Boost set the Boost property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) Boost(boost float64) *IpPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// CopyTo set the CopyTo property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) CopyTo(copyto Fields) *IpPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) DocValues(docvalues bool) *IpPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *IpPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) Fields(value map[PropertyName]Property) *IpPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) IgnoreAbove(ignoreabove int) *IpPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *IpPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) Index(index bool) *IpPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) LocalMetadata(localmetadata Metadata) *IpPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) Meta(value map[string]string) *IpPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) NullValue(nullvalue string) *IpPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// Properties set the Properties property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) Properties(value map[PropertyName]Property) *IpPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) Similarity(similarity string) *IpPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for IpPropertyBuilder.
func (rb *IpPropertyBuilder) Store(store bool) *IpPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for IpPropertyBuilder.
