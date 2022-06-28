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

// ComponentTemplate type.
type ComponentTemplate struct {
	ComponentTemplate ComponentTemplateNode `json:"component_template"`
	Name              Name                  `json:"name"`
}

// ComponentTemplateBuilder holds ComponentTemplate struct and provides a builder API.
type ComponentTemplateBuilder struct {
	v *ComponentTemplate
}

// NewComponentTemplate provides a builder for the ComponentTemplate struct.
func NewComponentTemplate() *ComponentTemplateBuilder {
	r := ComponentTemplateBuilder{
		&ComponentTemplate{},
	}

	return &r
}

// Build finalize the chain and returns the ComponentTemplate struct
func (rb *ComponentTemplateBuilder) Build() ComponentTemplate {
	return *rb.v
}

// ComponentTemplate set the ComponentTemplate property for ComponentTemplateBuilder.
func (rb *ComponentTemplateBuilder) ComponentTemplate(componenttemplate ComponentTemplateNode) *ComponentTemplateBuilder {
	rb.v.ComponentTemplate = componenttemplate
	return rb
}

// Name set the Name property for ComponentTemplateBuilder.
func (rb *ComponentTemplateBuilder) Name(name Name) *ComponentTemplateBuilder {
	rb.v.Name = name
	return rb
}
