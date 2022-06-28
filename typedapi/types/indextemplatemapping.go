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

// IndexTemplateMapping type.
type IndexTemplateMapping struct {
	Aliases  map[IndexName]Alias `json:"aliases,omitempty"`
	Mappings *TypeMapping        `json:"mappings,omitempty"`
	Settings *IndexSettings      `json:"settings,omitempty"`
}

// IndexTemplateMappingBuilder holds IndexTemplateMapping struct and provides a builder API.
type IndexTemplateMappingBuilder struct {
	v *IndexTemplateMapping
}

// NewIndexTemplateMapping provides a builder for the IndexTemplateMapping struct.
func NewIndexTemplateMapping() *IndexTemplateMappingBuilder {
	r := IndexTemplateMappingBuilder{
		&IndexTemplateMapping{
			Aliases: make(map[IndexName]Alias, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexTemplateMapping struct
func (rb *IndexTemplateMappingBuilder) Build() IndexTemplateMapping {
	return *rb.v
}

// Aliases set the Aliases property for IndexTemplateMappingBuilder.
func (rb *IndexTemplateMappingBuilder) Aliases(value map[IndexName]Alias) *IndexTemplateMappingBuilder {
	rb.v.Aliases = value
	return rb
}

// Mappings set the Mappings property for IndexTemplateMappingBuilder.
func (rb *IndexTemplateMappingBuilder) Mappings(mappings TypeMapping) *IndexTemplateMappingBuilder {
	rb.v.Mappings = &mappings
	return rb
}

// Settings set the Settings property for IndexTemplateMappingBuilder.
func (rb *IndexTemplateMappingBuilder) Settings(settings IndexSettings) *IndexTemplateMappingBuilder {
	rb.v.Settings = &settings
	return rb
}
