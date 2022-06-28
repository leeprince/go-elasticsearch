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

// SecurityRoles type.
type SecurityRoles struct {
	Dls    SecurityRolesDls    `json:"dls"`
	File   SecurityRolesFile   `json:"file"`
	Native SecurityRolesNative `json:"native"`
}

// SecurityRolesBuilder holds SecurityRoles struct and provides a builder API.
type SecurityRolesBuilder struct {
	v *SecurityRoles
}

// NewSecurityRoles provides a builder for the SecurityRoles struct.
func NewSecurityRoles() *SecurityRolesBuilder {
	r := SecurityRolesBuilder{
		&SecurityRoles{},
	}

	return &r
}

// Build finalize the chain and returns the SecurityRoles struct
func (rb *SecurityRolesBuilder) Build() SecurityRoles {
	return *rb.v
}

// Dls set the Dls property for SecurityRolesBuilder.
func (rb *SecurityRolesBuilder) Dls(dls SecurityRolesDls) *SecurityRolesBuilder {
	rb.v.Dls = dls
	return rb
}

// File set the File property for SecurityRolesBuilder.
func (rb *SecurityRolesBuilder) File(file SecurityRolesFile) *SecurityRolesBuilder {
	rb.v.File = file
	return rb
}

// Native set the Native property for SecurityRolesBuilder.
func (rb *SecurityRolesBuilder) Native(native SecurityRolesNative) *SecurityRolesBuilder {
	rb.v.Native = native
	return rb
}
