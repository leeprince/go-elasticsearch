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

// IndexTemplate type.
type IndexTemplate struct {
	AllowAutoCreate *bool                                 `json:"allow_auto_create,omitempty"`
	ComposedOf      []Name                                `json:"composed_of"`
	DataStream      *IndexTemplateDataStreamConfiguration `json:"data_stream,omitempty"`
	IndexPatterns   Names                                 `json:"index_patterns"`
	Meta_           *Metadata                             `json:"_meta,omitempty"`
	Priority        *int64                                `json:"priority,omitempty"`
	Template        *IndexTemplateSummary                 `json:"template,omitempty"`
	Version         *VersionNumber                        `json:"version,omitempty"`
}

// IndexTemplateBuilder holds IndexTemplate struct and provides a builder API.
type IndexTemplateBuilder struct {
	v *IndexTemplate
}

// NewIndexTemplate provides a builder for the IndexTemplate struct.
func NewIndexTemplate() *IndexTemplateBuilder {
	r := IndexTemplateBuilder{
		&IndexTemplate{},
	}

	return &r
}

// Build finalize the chain and returns the IndexTemplate struct
func (rb *IndexTemplateBuilder) Build() IndexTemplate {
	return *rb.v
}

// AllowAutoCreate set the AllowAutoCreate property for IndexTemplateBuilder.
func (rb *IndexTemplateBuilder) AllowAutoCreate(allowautocreate bool) *IndexTemplateBuilder {
	rb.v.AllowAutoCreate = &allowautocreate
	return rb
}

// ComposedOf set the ComposedOf property for IndexTemplateBuilder.
func (rb *IndexTemplateBuilder) ComposedOf(composed_of ...Name) *IndexTemplateBuilder {
	rb.v.ComposedOf = append(rb.v.ComposedOf, composed_of...)
	return rb
}

// DataStream set the DataStream property for IndexTemplateBuilder.
func (rb *IndexTemplateBuilder) DataStream(datastream IndexTemplateDataStreamConfiguration) *IndexTemplateBuilder {
	rb.v.DataStream = &datastream
	return rb
}

// IndexPatterns set the IndexPatterns property for IndexTemplateBuilder.
func (rb *IndexTemplateBuilder) IndexPatterns(indexpatterns Names) *IndexTemplateBuilder {
	rb.v.IndexPatterns = indexpatterns
	return rb
}

// Meta_ set the Meta_ property for IndexTemplateBuilder.
func (rb *IndexTemplateBuilder) Meta_(meta_ Metadata) *IndexTemplateBuilder {
	rb.v.Meta_ = &meta_
	return rb
}

// Priority set the Priority property for IndexTemplateBuilder.
func (rb *IndexTemplateBuilder) Priority(priority int64) *IndexTemplateBuilder {
	rb.v.Priority = &priority
	return rb
}

// Template set the Template property for IndexTemplateBuilder.
func (rb *IndexTemplateBuilder) Template(template IndexTemplateSummary) *IndexTemplateBuilder {
	rb.v.Template = &template
	return rb
}

// Version set the Version property for IndexTemplateBuilder.
func (rb *IndexTemplateBuilder) Version(version VersionNumber) *IndexTemplateBuilder {
	rb.v.Version = &version
	return rb
}
