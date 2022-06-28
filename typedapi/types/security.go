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

// Security type.
type Security struct {
	Anonymous          FeatureToggle          `json:"anonymous"`
	ApiKeyService      FeatureToggle          `json:"api_key_service"`
	Audit              Audit                  `json:"audit"`
	Available          bool                   `json:"available"`
	Enabled            bool                   `json:"enabled"`
	Fips140            FeatureToggle          `json:"fips_140"`
	Ipfilter           IpFilter               `json:"ipfilter"`
	OperatorPrivileges Base                   `json:"operator_privileges"`
	Realms             map[string]Realm       `json:"realms"`
	RoleMapping        map[string]RoleMapping `json:"role_mapping"`
	Roles              SecurityRoles          `json:"roles"`
	Ssl                Ssl                    `json:"ssl"`
	SystemKey          *FeatureToggle         `json:"system_key,omitempty"`
	TokenService       FeatureToggle          `json:"token_service"`
}

// SecurityBuilder holds Security struct and provides a builder API.
type SecurityBuilder struct {
	v *Security
}

// NewSecurity provides a builder for the Security struct.
func NewSecurity() *SecurityBuilder {
	r := SecurityBuilder{
		&Security{
			Realms:      make(map[string]Realm, 0),
			RoleMapping: make(map[string]RoleMapping, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Security struct
func (rb *SecurityBuilder) Build() Security {
	return *rb.v
}

// Anonymous set the Anonymous property for SecurityBuilder.
func (rb *SecurityBuilder) Anonymous(anonymous FeatureToggle) *SecurityBuilder {
	rb.v.Anonymous = anonymous
	return rb
}

// ApiKeyService set the ApiKeyService property for SecurityBuilder.
func (rb *SecurityBuilder) ApiKeyService(apikeyservice FeatureToggle) *SecurityBuilder {
	rb.v.ApiKeyService = apikeyservice
	return rb
}

// Audit set the Audit property for SecurityBuilder.
func (rb *SecurityBuilder) Audit(audit Audit) *SecurityBuilder {
	rb.v.Audit = audit
	return rb
}

// Available set the Available property for SecurityBuilder.
func (rb *SecurityBuilder) Available(available bool) *SecurityBuilder {
	rb.v.Available = available
	return rb
}

// Enabled set the Enabled property for SecurityBuilder.
func (rb *SecurityBuilder) Enabled(enabled bool) *SecurityBuilder {
	rb.v.Enabled = enabled
	return rb
}

// Fips140 set the Fips140 property for SecurityBuilder.
func (rb *SecurityBuilder) Fips140(fips140 FeatureToggle) *SecurityBuilder {
	rb.v.Fips140 = fips140
	return rb
}

// Ipfilter set the Ipfilter property for SecurityBuilder.
func (rb *SecurityBuilder) Ipfilter(ipfilter IpFilter) *SecurityBuilder {
	rb.v.Ipfilter = ipfilter
	return rb
}

// OperatorPrivileges set the OperatorPrivileges property for SecurityBuilder.
func (rb *SecurityBuilder) OperatorPrivileges(operatorprivileges Base) *SecurityBuilder {
	rb.v.OperatorPrivileges = operatorprivileges
	return rb
}

// Realms set the Realms property for SecurityBuilder.
func (rb *SecurityBuilder) Realms(value map[string]Realm) *SecurityBuilder {
	rb.v.Realms = value
	return rb
}

// RoleMapping set the RoleMapping property for SecurityBuilder.
func (rb *SecurityBuilder) RoleMapping(value map[string]RoleMapping) *SecurityBuilder {
	rb.v.RoleMapping = value
	return rb
}

// Roles set the Roles property for SecurityBuilder.
func (rb *SecurityBuilder) Roles(roles SecurityRoles) *SecurityBuilder {
	rb.v.Roles = roles
	return rb
}

// Ssl set the Ssl property for SecurityBuilder.
func (rb *SecurityBuilder) Ssl(ssl Ssl) *SecurityBuilder {
	rb.v.Ssl = ssl
	return rb
}

// SystemKey set the SystemKey property for SecurityBuilder.
func (rb *SecurityBuilder) SystemKey(systemkey FeatureToggle) *SecurityBuilder {
	rb.v.SystemKey = &systemkey
	return rb
}

// TokenService set the TokenService property for SecurityBuilder.
func (rb *SecurityBuilder) TokenService(tokenservice FeatureToggle) *SecurityBuilder {
	rb.v.TokenService = tokenservice
	return rb
}
