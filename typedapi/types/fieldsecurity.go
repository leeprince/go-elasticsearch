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

// FieldSecurity type.
type FieldSecurity struct {
	Except *Fields `json:"except,omitempty"`
	Grant  Fields  `json:"grant"`
}

// FieldSecurityBuilder holds FieldSecurity struct and provides a builder API.
type FieldSecurityBuilder struct {
	v *FieldSecurity
}

// NewFieldSecurity provides a builder for the FieldSecurity struct.
func NewFieldSecurity() *FieldSecurityBuilder {
	r := FieldSecurityBuilder{
		&FieldSecurity{},
	}

	return &r
}

// Build finalize the chain and returns the FieldSecurity struct
func (rb *FieldSecurityBuilder) Build() FieldSecurity {
	return *rb.v
}

// Except set the Except property for FieldSecurityBuilder.
func (rb *FieldSecurityBuilder) Except(except Fields) *FieldSecurityBuilder {
	rb.v.Except = &except
	return rb
}

// Grant set the Grant property for FieldSecurityBuilder.
func (rb *FieldSecurityBuilder) Grant(grant Fields) *FieldSecurityBuilder {
	rb.v.Grant = grant
	return rb
}
