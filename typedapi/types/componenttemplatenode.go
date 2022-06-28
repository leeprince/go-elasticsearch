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

// ComponentTemplateNode type.
type ComponentTemplateNode struct {
	Meta_    *Metadata                `json:"_meta,omitempty"`
	Template ComponentTemplateSummary `json:"template"`
	Version  *VersionNumber           `json:"version,omitempty"`
}

// ComponentTemplateNodeBuilder holds ComponentTemplateNode struct and provides a builder API.
type ComponentTemplateNodeBuilder struct {
	v *ComponentTemplateNode
}

// NewComponentTemplateNode provides a builder for the ComponentTemplateNode struct.
func NewComponentTemplateNode() *ComponentTemplateNodeBuilder {
	r := ComponentTemplateNodeBuilder{
		&ComponentTemplateNode{},
	}

	return &r
}

// Build finalize the chain and returns the ComponentTemplateNode struct
func (rb *ComponentTemplateNodeBuilder) Build() ComponentTemplateNode {
	return *rb.v
}

// Meta_ set the Meta_ property for ComponentTemplateNodeBuilder.
func (rb *ComponentTemplateNodeBuilder) Meta_(meta_ Metadata) *ComponentTemplateNodeBuilder {
	rb.v.Meta_ = &meta_
	return rb
}

// Template set the Template property for ComponentTemplateNodeBuilder.
func (rb *ComponentTemplateNodeBuilder) Template(template ComponentTemplateSummary) *ComponentTemplateNodeBuilder {
	rb.v.Template = template
	return rb
}

// Version set the Version property for ComponentTemplateNodeBuilder.
func (rb *ComponentTemplateNodeBuilder) Version(version VersionNumber) *ComponentTemplateNodeBuilder {
	rb.v.Version = &version
	return rb
}
