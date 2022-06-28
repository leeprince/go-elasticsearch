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

// Script holds the union for the following types:
//     InlineScript
//     StoredScriptId
type Script interface{}

// ScriptBuilder holds Script struct and provides a builder API.
type ScriptBuilder struct {
	v Script
}

// NewScript provides a builder for the Script struct.
func NewScript() *ScriptBuilder {
	return &ScriptBuilder{}
}

// Build finalize the chain and returns the Script struct
func (u *ScriptBuilder) Build() Script {
	return u.v
}

// InlineScript set the InlineScript property for ScriptBuilder.
func (u *ScriptBuilder) InlineScript(v InlineScript) *ScriptBuilder {
	u.v = v
	return u
}

// StoredScriptId set the StoredScriptId property for ScriptBuilder.
func (u *ScriptBuilder) StoredScriptId(v StoredScriptId) *ScriptBuilder {
	u.v = v
	return u
}
