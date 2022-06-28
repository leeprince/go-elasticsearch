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

// PropertyBase type.
type PropertyBase struct {
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
}

// PropertyBaseBuilder holds PropertyBase struct and provides a builder API.
type PropertyBaseBuilder struct {
	v *PropertyBase
}

// NewPropertyBase provides a builder for the PropertyBase struct.
func NewPropertyBase() *PropertyBaseBuilder {
	r := PropertyBaseBuilder{
		&PropertyBase{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the PropertyBase struct
func (rb *PropertyBaseBuilder) Build() PropertyBase {
	return *rb.v
}

// Dynamic set the Dynamic property for PropertyBaseBuilder.
func (rb *PropertyBaseBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *PropertyBaseBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for PropertyBaseBuilder.
func (rb *PropertyBaseBuilder) Fields(value map[PropertyName]Property) *PropertyBaseBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for PropertyBaseBuilder.
func (rb *PropertyBaseBuilder) IgnoreAbove(ignoreabove int) *PropertyBaseBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for PropertyBaseBuilder.
func (rb *PropertyBaseBuilder) LocalMetadata(localmetadata Metadata) *PropertyBaseBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for PropertyBaseBuilder.
func (rb *PropertyBaseBuilder) Meta(value map[string]string) *PropertyBaseBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for PropertyBaseBuilder.
func (rb *PropertyBaseBuilder) Properties(value map[PropertyName]Property) *PropertyBaseBuilder {
	rb.v.Properties = value
	return rb
}
