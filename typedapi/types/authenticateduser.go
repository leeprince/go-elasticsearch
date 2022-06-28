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

// AuthenticatedUser type.
type AuthenticatedUser struct {
	AuthenticationProvider *AuthenticationProvider `json:"authentication_provider,omitempty"`
	AuthenticationRealm    UserRealm               `json:"authentication_realm"`
	AuthenticationType     string                  `json:"authentication_type"`
	Email                  string                  `json:"email,omitempty"`
	Enabled                bool                    `json:"enabled"`
	FullName               Name                    `json:"full_name,omitempty"`
	LookupRealm            UserRealm               `json:"lookup_realm"`
	Metadata               Metadata                `json:"metadata"`
	Roles                  []string                `json:"roles"`
	Username               Username                `json:"username"`
}

// AuthenticatedUserBuilder holds AuthenticatedUser struct and provides a builder API.
type AuthenticatedUserBuilder struct {
	v *AuthenticatedUser
}

// NewAuthenticatedUser provides a builder for the AuthenticatedUser struct.
func NewAuthenticatedUser() *AuthenticatedUserBuilder {
	r := AuthenticatedUserBuilder{
		&AuthenticatedUser{},
	}

	return &r
}

// Build finalize the chain and returns the AuthenticatedUser struct
func (rb *AuthenticatedUserBuilder) Build() AuthenticatedUser {
	return *rb.v
}

// AuthenticationProvider set the AuthenticationProvider property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) AuthenticationProvider(authenticationprovider AuthenticationProvider) *AuthenticatedUserBuilder {
	rb.v.AuthenticationProvider = &authenticationprovider
	return rb
}

// AuthenticationRealm set the AuthenticationRealm property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) AuthenticationRealm(authenticationrealm UserRealm) *AuthenticatedUserBuilder {
	rb.v.AuthenticationRealm = authenticationrealm
	return rb
}

// AuthenticationType set the AuthenticationType property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) AuthenticationType(authenticationtype string) *AuthenticatedUserBuilder {
	rb.v.AuthenticationType = authenticationtype
	return rb
}

// Email set the Email property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) Email(email string) *AuthenticatedUserBuilder {
	rb.v.Email = email
	return rb
}

// Enabled set the Enabled property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) Enabled(enabled bool) *AuthenticatedUserBuilder {
	rb.v.Enabled = enabled
	return rb
}

// FullName set the FullName property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) FullName(fullname Name) *AuthenticatedUserBuilder {
	rb.v.FullName = fullname
	return rb
}

// LookupRealm set the LookupRealm property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) LookupRealm(lookuprealm UserRealm) *AuthenticatedUserBuilder {
	rb.v.LookupRealm = lookuprealm
	return rb
}

// Metadata set the Metadata property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) Metadata(metadata Metadata) *AuthenticatedUserBuilder {
	rb.v.Metadata = metadata
	return rb
}

// Roles set the Roles property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) Roles(roles ...string) *AuthenticatedUserBuilder {
	rb.v.Roles = append(rb.v.Roles, roles...)
	return rb
}

// Username set the Username property for AuthenticatedUserBuilder.
func (rb *AuthenticatedUserBuilder) Username(username Username) *AuthenticatedUserBuilder {
	rb.v.Username = username
	return rb
}
