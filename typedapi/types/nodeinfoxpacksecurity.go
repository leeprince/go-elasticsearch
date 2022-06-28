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

// NodeInfoXpackSecurity type.
type NodeInfoXpackSecurity struct {
	Authc     *NodeInfoXpackSecurityAuthc `json:"authc,omitempty"`
	Enabled   string                      `json:"enabled"`
	Http      NodeInfoXpackSecuritySsl    `json:"http"`
	Transport *NodeInfoXpackSecuritySsl   `json:"transport,omitempty"`
}

// NodeInfoXpackSecurityBuilder holds NodeInfoXpackSecurity struct and provides a builder API.
type NodeInfoXpackSecurityBuilder struct {
	v *NodeInfoXpackSecurity
}

// NewNodeInfoXpackSecurity provides a builder for the NodeInfoXpackSecurity struct.
func NewNodeInfoXpackSecurity() *NodeInfoXpackSecurityBuilder {
	r := NodeInfoXpackSecurityBuilder{
		&NodeInfoXpackSecurity{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoXpackSecurity struct
func (rb *NodeInfoXpackSecurityBuilder) Build() NodeInfoXpackSecurity {
	return *rb.v
}

// Authc set the Authc property for NodeInfoXpackSecurityBuilder.
func (rb *NodeInfoXpackSecurityBuilder) Authc(authc NodeInfoXpackSecurityAuthc) *NodeInfoXpackSecurityBuilder {
	rb.v.Authc = &authc
	return rb
}

// Enabled set the Enabled property for NodeInfoXpackSecurityBuilder.
func (rb *NodeInfoXpackSecurityBuilder) Enabled(enabled string) *NodeInfoXpackSecurityBuilder {
	rb.v.Enabled = enabled
	return rb
}

// Http set the Http property for NodeInfoXpackSecurityBuilder.
func (rb *NodeInfoXpackSecurityBuilder) Http(http NodeInfoXpackSecuritySsl) *NodeInfoXpackSecurityBuilder {
	rb.v.Http = http
	return rb
}

// Transport set the Transport property for NodeInfoXpackSecurityBuilder.
func (rb *NodeInfoXpackSecurityBuilder) Transport(transport NodeInfoXpackSecuritySsl) *NodeInfoXpackSecurityBuilder {
	rb.v.Transport = &transport
	return rb
}
