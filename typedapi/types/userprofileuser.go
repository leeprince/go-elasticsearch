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

// UserProfileUser type.
type UserProfileUser struct {
	Email       string   `json:"email,omitempty"`
	FullName    Name     `json:"full_name,omitempty"`
	RealmDomain *Name    `json:"realm_domain,omitempty"`
	RealmName   Name     `json:"realm_name"`
	Roles       []string `json:"roles"`
	Username    Username `json:"username"`
}

// UserProfileUserBuilder holds UserProfileUser struct and provides a builder API.
type UserProfileUserBuilder struct {
	v *UserProfileUser
}

// NewUserProfileUser provides a builder for the UserProfileUser struct.
func NewUserProfileUser() *UserProfileUserBuilder {
	r := UserProfileUserBuilder{
		&UserProfileUser{},
	}

	return &r
}

// Build finalize the chain and returns the UserProfileUser struct
func (rb *UserProfileUserBuilder) Build() UserProfileUser {
	return *rb.v
}

// Email set the Email property for UserProfileUserBuilder.
func (rb *UserProfileUserBuilder) Email(email string) *UserProfileUserBuilder {
	rb.v.Email = email
	return rb
}

// FullName set the FullName property for UserProfileUserBuilder.
func (rb *UserProfileUserBuilder) FullName(fullname Name) *UserProfileUserBuilder {
	rb.v.FullName = fullname
	return rb
}

// RealmDomain set the RealmDomain property for UserProfileUserBuilder.
func (rb *UserProfileUserBuilder) RealmDomain(realmdomain Name) *UserProfileUserBuilder {
	rb.v.RealmDomain = &realmdomain
	return rb
}

// RealmName set the RealmName property for UserProfileUserBuilder.
func (rb *UserProfileUserBuilder) RealmName(realmname Name) *UserProfileUserBuilder {
	rb.v.RealmName = realmname
	return rb
}

// Roles set the Roles property for UserProfileUserBuilder.
func (rb *UserProfileUserBuilder) Roles(roles ...string) *UserProfileUserBuilder {
	rb.v.Roles = append(rb.v.Roles, roles...)
	return rb
}

// Username set the Username property for UserProfileUserBuilder.
func (rb *UserProfileUserBuilder) Username(username Username) *UserProfileUserBuilder {
	rb.v.Username = username
	return rb
}
