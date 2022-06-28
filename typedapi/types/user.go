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

// User type.
type User struct {
	Email    string   `json:"email,omitempty"`
	Enabled  bool     `json:"enabled"`
	FullName Name     `json:"full_name,omitempty"`
	Metadata Metadata `json:"metadata"`
	Roles    []string `json:"roles"`
	Username Username `json:"username"`
}

// UserBuilder holds User struct and provides a builder API.
type UserBuilder struct {
	v *User
}

// NewUser provides a builder for the User struct.
func NewUser() *UserBuilder {
	r := UserBuilder{
		&User{},
	}

	return &r
}

// Build finalize the chain and returns the User struct
func (rb *UserBuilder) Build() User {
	return *rb.v
}

// Email set the Email property for UserBuilder.
func (rb *UserBuilder) Email(email string) *UserBuilder {
	rb.v.Email = email
	return rb
}

// Enabled set the Enabled property for UserBuilder.
func (rb *UserBuilder) Enabled(enabled bool) *UserBuilder {
	rb.v.Enabled = enabled
	return rb
}

// FullName set the FullName property for UserBuilder.
func (rb *UserBuilder) FullName(fullname Name) *UserBuilder {
	rb.v.FullName = fullname
	return rb
}

// Metadata set the Metadata property for UserBuilder.
func (rb *UserBuilder) Metadata(metadata Metadata) *UserBuilder {
	rb.v.Metadata = metadata
	return rb
}

// Roles set the Roles property for UserBuilder.
func (rb *UserBuilder) Roles(roles ...string) *UserBuilder {
	rb.v.Roles = append(rb.v.Roles, roles...)
	return rb
}

// Username set the Username property for UserBuilder.
func (rb *UserBuilder) Username(username Username) *UserBuilder {
	rb.v.Username = username
	return rb
}
