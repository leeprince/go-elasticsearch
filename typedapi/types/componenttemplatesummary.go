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

// ComponentTemplateSummary type.
type ComponentTemplateSummary struct {
	Aliases  map[string]AliasDefinition  `json:"aliases,omitempty"`
	Mappings *TypeMapping                `json:"mappings,omitempty"`
	Meta_    *Metadata                   `json:"_meta,omitempty"`
	Settings map[IndexName]IndexSettings `json:"settings"`
	Version  *VersionNumber              `json:"version,omitempty"`
}

// ComponentTemplateSummaryBuilder holds ComponentTemplateSummary struct and provides a builder API.
type ComponentTemplateSummaryBuilder struct {
	v *ComponentTemplateSummary
}

// NewComponentTemplateSummary provides a builder for the ComponentTemplateSummary struct.
func NewComponentTemplateSummary() *ComponentTemplateSummaryBuilder {
	r := ComponentTemplateSummaryBuilder{
		&ComponentTemplateSummary{
			Aliases:  make(map[string]AliasDefinition, 0),
			Settings: make(map[IndexName]IndexSettings, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ComponentTemplateSummary struct
func (rb *ComponentTemplateSummaryBuilder) Build() ComponentTemplateSummary {
	return *rb.v
}

// Aliases set the Aliases property for ComponentTemplateSummaryBuilder.
func (rb *ComponentTemplateSummaryBuilder) Aliases(value map[string]AliasDefinition) *ComponentTemplateSummaryBuilder {
	rb.v.Aliases = value
	return rb
}

// Mappings set the Mappings property for ComponentTemplateSummaryBuilder.
func (rb *ComponentTemplateSummaryBuilder) Mappings(mappings TypeMapping) *ComponentTemplateSummaryBuilder {
	rb.v.Mappings = &mappings
	return rb
}

// Meta_ set the Meta_ property for ComponentTemplateSummaryBuilder.
func (rb *ComponentTemplateSummaryBuilder) Meta_(meta_ Metadata) *ComponentTemplateSummaryBuilder {
	rb.v.Meta_ = &meta_
	return rb
}

// Settings set the Settings property for ComponentTemplateSummaryBuilder.
func (rb *ComponentTemplateSummaryBuilder) Settings(value map[IndexName]IndexSettings) *ComponentTemplateSummaryBuilder {
	rb.v.Settings = value
	return rb
}

// Version set the Version property for ComponentTemplateSummaryBuilder.
func (rb *ComponentTemplateSummaryBuilder) Version(version VersionNumber) *ComponentTemplateSummaryBuilder {
	rb.v.Version = &version
	return rb
}
