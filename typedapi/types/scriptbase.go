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

// ScriptBase type.
type ScriptBase struct {
	Params map[string]interface{} `json:"params,omitempty"`
}

// ScriptBaseBuilder holds ScriptBase struct and provides a builder API.
type ScriptBaseBuilder struct {
	v *ScriptBase
}

// NewScriptBase provides a builder for the ScriptBase struct.
func NewScriptBase() *ScriptBaseBuilder {
	r := ScriptBaseBuilder{
		&ScriptBase{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ScriptBase struct
func (rb *ScriptBaseBuilder) Build() ScriptBase {
	return *rb.v
}

// Params set the Params property for ScriptBaseBuilder.
func (rb *ScriptBaseBuilder) Params(value map[string]interface{}) *ScriptBaseBuilder {
	rb.v.Params = value
	return rb
}
