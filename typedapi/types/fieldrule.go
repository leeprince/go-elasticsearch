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

// FieldRule type.
type FieldRule struct {
	Dn       *Names      `json:"dn,omitempty"`
	Groups   *Names      `json:"groups,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Realm    *Realm      `json:"realm,omitempty"`
	Username *Name       `json:"username,omitempty"`
}

// FieldRuleBuilder holds FieldRule struct and provides a builder API.
type FieldRuleBuilder struct {
	v *FieldRule
}

// NewFieldRule provides a builder for the FieldRule struct.
func NewFieldRule() *FieldRuleBuilder {
	r := FieldRuleBuilder{
		&FieldRule{},
	}

	return &r
}

// Build finalize the chain and returns the FieldRule struct
func (rb *FieldRuleBuilder) Build() FieldRule {
	return *rb.v
}

// Dn set the Dn property for FieldRuleBuilder.
func (rb *FieldRuleBuilder) Dn(dn Names) *FieldRuleBuilder {
	rb.v.Dn = &dn
	return rb
}

// Groups set the Groups property for FieldRuleBuilder.
func (rb *FieldRuleBuilder) Groups(groups Names) *FieldRuleBuilder {
	rb.v.Groups = &groups
	return rb
}

// Metadata set the Metadata property for FieldRuleBuilder.
func (rb *FieldRuleBuilder) Metadata(metadata interface{}) *FieldRuleBuilder {
	rb.v.Metadata = metadata
	return rb
}

// Realm set the Realm property for FieldRuleBuilder.
func (rb *FieldRuleBuilder) Realm(realm Realm) *FieldRuleBuilder {
	rb.v.Realm = &realm
	return rb
}

// Username set the Username property for FieldRuleBuilder.
func (rb *FieldRuleBuilder) Username(username Name) *FieldRuleBuilder {
	rb.v.Username = &username
	return rb
}
