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

// RoleMappingRule type.
type RoleMappingRule struct {
	All    []RoleMappingRule `json:"all,omitempty"`
	Any    []RoleMappingRule `json:"any,omitempty"`
	Except *RoleMappingRule  `json:"except,omitempty"`
	Field  *FieldRule        `json:"field,omitempty"`
}

// RoleMappingRuleBuilder holds RoleMappingRule struct and provides a builder API.
type RoleMappingRuleBuilder struct {
	v *RoleMappingRule
}

// NewRoleMappingRule provides a builder for the RoleMappingRule struct.
func NewRoleMappingRule() *RoleMappingRuleBuilder {
	r := RoleMappingRuleBuilder{
		&RoleMappingRule{},
	}

	return &r
}

// Build finalize the chain and returns the RoleMappingRule struct
func (rb *RoleMappingRuleBuilder) Build() RoleMappingRule {
	return *rb.v
}

// All set the All property for RoleMappingRuleBuilder.
func (rb *RoleMappingRuleBuilder) All(all ...RoleMappingRule) *RoleMappingRuleBuilder {
	rb.v.All = append(rb.v.All, all...)
	return rb
}

// Any set the Any property for RoleMappingRuleBuilder.
func (rb *RoleMappingRuleBuilder) Any(any ...RoleMappingRule) *RoleMappingRuleBuilder {
	rb.v.Any = append(rb.v.Any, any...)
	return rb
}

// Except set the Except property for RoleMappingRuleBuilder.
func (rb *RoleMappingRuleBuilder) Except(except RoleMappingRule) *RoleMappingRuleBuilder {
	rb.v.Except = &except
	return rb
}

// Field set the Field property for RoleMappingRuleBuilder.
func (rb *RoleMappingRuleBuilder) Field(field FieldRule) *RoleMappingRuleBuilder {
	rb.v.Field = &field
	return rb
}
