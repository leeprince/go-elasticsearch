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

// ScriptTransform type.
type ScriptTransform struct {
	Id     *string                `json:"id,omitempty"`
	Lang   *string                `json:"lang,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	Source *string                `json:"source,omitempty"`
}

// ScriptTransformBuilder holds ScriptTransform struct and provides a builder API.
type ScriptTransformBuilder struct {
	v *ScriptTransform
}

// NewScriptTransform provides a builder for the ScriptTransform struct.
func NewScriptTransform() *ScriptTransformBuilder {
	r := ScriptTransformBuilder{
		&ScriptTransform{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ScriptTransform struct
func (rb *ScriptTransformBuilder) Build() ScriptTransform {
	return *rb.v
}

// Id set the Id property for ScriptTransformBuilder.
func (rb *ScriptTransformBuilder) Id(id string) *ScriptTransformBuilder {
	rb.v.Id = &id
	return rb
}

// Lang set the Lang property for ScriptTransformBuilder.
func (rb *ScriptTransformBuilder) Lang(lang string) *ScriptTransformBuilder {
	rb.v.Lang = &lang
	return rb
}

// Params set the Params property for ScriptTransformBuilder.
func (rb *ScriptTransformBuilder) Params(value map[string]interface{}) *ScriptTransformBuilder {
	rb.v.Params = value
	return rb
}

// Source set the Source property for ScriptTransformBuilder.
func (rb *ScriptTransformBuilder) Source(source string) *ScriptTransformBuilder {
	rb.v.Source = &source
	return rb
}
