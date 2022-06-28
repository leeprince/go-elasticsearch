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

// HistogramProperty type.
type HistogramProperty struct {
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// HistogramPropertyBuilder holds HistogramProperty struct and provides a builder API.
type HistogramPropertyBuilder struct {
	v *HistogramProperty
}

// NewHistogramProperty provides a builder for the HistogramProperty struct.
func NewHistogramProperty() *HistogramPropertyBuilder {
	r := HistogramPropertyBuilder{
		&HistogramProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "histogram"

	return &r
}

// Build finalize the chain and returns the HistogramProperty struct
func (rb *HistogramPropertyBuilder) Build() HistogramProperty {
	return *rb.v
}

// Dynamic set the Dynamic property for HistogramPropertyBuilder.
func (rb *HistogramPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *HistogramPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for HistogramPropertyBuilder.
func (rb *HistogramPropertyBuilder) Fields(value map[PropertyName]Property) *HistogramPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for HistogramPropertyBuilder.
func (rb *HistogramPropertyBuilder) IgnoreAbove(ignoreabove int) *HistogramPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for HistogramPropertyBuilder.
func (rb *HistogramPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *HistogramPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// LocalMetadata set the LocalMetadata property for HistogramPropertyBuilder.
func (rb *HistogramPropertyBuilder) LocalMetadata(localmetadata Metadata) *HistogramPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for HistogramPropertyBuilder.
func (rb *HistogramPropertyBuilder) Meta(value map[string]string) *HistogramPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for HistogramPropertyBuilder.
func (rb *HistogramPropertyBuilder) Properties(value map[PropertyName]Property) *HistogramPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Type set the Type property for HistogramPropertyBuilder.
