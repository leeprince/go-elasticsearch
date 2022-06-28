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

// RankFeaturesProperty type.
type RankFeaturesProperty struct {
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// RankFeaturesPropertyBuilder holds RankFeaturesProperty struct and provides a builder API.
type RankFeaturesPropertyBuilder struct {
	v *RankFeaturesProperty
}

// NewRankFeaturesProperty provides a builder for the RankFeaturesProperty struct.
func NewRankFeaturesProperty() *RankFeaturesPropertyBuilder {
	r := RankFeaturesPropertyBuilder{
		&RankFeaturesProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "rank_features"

	return &r
}

// Build finalize the chain and returns the RankFeaturesProperty struct
func (rb *RankFeaturesPropertyBuilder) Build() RankFeaturesProperty {
	return *rb.v
}

// Dynamic set the Dynamic property for RankFeaturesPropertyBuilder.
func (rb *RankFeaturesPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *RankFeaturesPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for RankFeaturesPropertyBuilder.
func (rb *RankFeaturesPropertyBuilder) Fields(value map[PropertyName]Property) *RankFeaturesPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for RankFeaturesPropertyBuilder.
func (rb *RankFeaturesPropertyBuilder) IgnoreAbove(ignoreabove int) *RankFeaturesPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for RankFeaturesPropertyBuilder.
func (rb *RankFeaturesPropertyBuilder) LocalMetadata(localmetadata Metadata) *RankFeaturesPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for RankFeaturesPropertyBuilder.
func (rb *RankFeaturesPropertyBuilder) Meta(value map[string]string) *RankFeaturesPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Properties set the Properties property for RankFeaturesPropertyBuilder.
func (rb *RankFeaturesPropertyBuilder) Properties(value map[PropertyName]Property) *RankFeaturesPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Type set the Type property for RankFeaturesPropertyBuilder.
