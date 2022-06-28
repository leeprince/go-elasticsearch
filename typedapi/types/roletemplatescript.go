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

// RoleTemplateScript holds the union for the following types:
//     RoleTemplateInlineScript
//     StoredScriptId
type RoleTemplateScript interface{}

// RoleTemplateScriptBuilder holds RoleTemplateScript struct and provides a builder API.
type RoleTemplateScriptBuilder struct {
	v RoleTemplateScript
}

// NewRoleTemplateScript provides a builder for the RoleTemplateScript struct.
func NewRoleTemplateScript() *RoleTemplateScriptBuilder {
	return &RoleTemplateScriptBuilder{}
}

// Build finalize the chain and returns the RoleTemplateScript struct
func (u *RoleTemplateScriptBuilder) Build() RoleTemplateScript {
	return u.v
}

// RoleTemplateInlineScript set the RoleTemplateInlineScript property for RoleTemplateScriptBuilder.
func (u *RoleTemplateScriptBuilder) RoleTemplateInlineScript(v RoleTemplateInlineScript) *RoleTemplateScriptBuilder {
	u.v = v
	return u
}

// StoredScriptId set the StoredScriptId property for RoleTemplateScriptBuilder.
func (u *RoleTemplateScriptBuilder) StoredScriptId(v StoredScriptId) *RoleTemplateScriptBuilder {
	u.v = v
	return u
}
