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

// Realm type.
type Realm struct {
	Available                 bool         `json:"available"`
	Cache                     []RealmCache `json:"cache,omitempty"`
	Enabled                   bool         `json:"enabled"`
	HasAuthorizationRealms    []bool       `json:"has_authorization_realms,omitempty"`
	HasDefaultUsernamePattern []bool       `json:"has_default_username_pattern,omitempty"`
	HasTruststore             []bool       `json:"has_truststore,omitempty"`
	IsAuthenticationDelegated []bool       `json:"is_authentication_delegated,omitempty"`
	Name                      []string     `json:"name,omitempty"`
	Order                     []int64      `json:"order,omitempty"`
	Size                      []int64      `json:"size,omitempty"`
}

// RealmBuilder holds Realm struct and provides a builder API.
type RealmBuilder struct {
	v *Realm
}

// NewRealm provides a builder for the Realm struct.
func NewRealm() *RealmBuilder {
	r := RealmBuilder{
		&Realm{},
	}

	return &r
}

// Build finalize the chain and returns the Realm struct
func (rb *RealmBuilder) Build() Realm {
	return *rb.v
}

// Available set the Available property for RealmBuilder.
func (rb *RealmBuilder) Available(available bool) *RealmBuilder {
	rb.v.Available = available
	return rb
}

// Cache set the Cache property for RealmBuilder.
func (rb *RealmBuilder) Cache(cache ...RealmCache) *RealmBuilder {
	rb.v.Cache = append(rb.v.Cache, cache...)
	return rb
}

// Enabled set the Enabled property for RealmBuilder.
func (rb *RealmBuilder) Enabled(enabled bool) *RealmBuilder {
	rb.v.Enabled = enabled
	return rb
}

// HasAuthorizationRealms set the HasAuthorizationRealms property for RealmBuilder.
func (rb *RealmBuilder) HasAuthorizationRealms(has_authorization_realms ...bool) *RealmBuilder {
	rb.v.HasAuthorizationRealms = append(rb.v.HasAuthorizationRealms, has_authorization_realms...)
	return rb
}

// HasDefaultUsernamePattern set the HasDefaultUsernamePattern property for RealmBuilder.
func (rb *RealmBuilder) HasDefaultUsernamePattern(has_default_username_pattern ...bool) *RealmBuilder {
	rb.v.HasDefaultUsernamePattern = append(rb.v.HasDefaultUsernamePattern, has_default_username_pattern...)
	return rb
}

// HasTruststore set the HasTruststore property for RealmBuilder.
func (rb *RealmBuilder) HasTruststore(has_truststore ...bool) *RealmBuilder {
	rb.v.HasTruststore = append(rb.v.HasTruststore, has_truststore...)
	return rb
}

// IsAuthenticationDelegated set the IsAuthenticationDelegated property for RealmBuilder.
func (rb *RealmBuilder) IsAuthenticationDelegated(is_authentication_delegated ...bool) *RealmBuilder {
	rb.v.IsAuthenticationDelegated = append(rb.v.IsAuthenticationDelegated, is_authentication_delegated...)
	return rb
}

// Name set the Name property for RealmBuilder.
func (rb *RealmBuilder) Name(name ...string) *RealmBuilder {
	rb.v.Name = append(rb.v.Name, name...)
	return rb
}

// Order set the Order property for RealmBuilder.
func (rb *RealmBuilder) Order(order ...int64) *RealmBuilder {
	rb.v.Order = append(rb.v.Order, order...)
	return rb
}

// Size set the Size property for RealmBuilder.
func (rb *RealmBuilder) Size(size ...int64) *RealmBuilder {
	rb.v.Size = append(rb.v.Size, size...)
	return rb
}
