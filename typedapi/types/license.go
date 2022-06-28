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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensetype"
)

// License type.
type License struct {
	ExpiryDateInMillis EpochMillis             `json:"expiry_date_in_millis"`
	IssueDateInMillis  EpochMillis             `json:"issue_date_in_millis"`
	IssuedTo           string                  `json:"issued_to"`
	Issuer             string                  `json:"issuer"`
	MaxNodes           int64                   `json:"max_nodes,omitempty"`
	MaxResourceUnits   *int64                  `json:"max_resource_units,omitempty"`
	Signature          string                  `json:"signature"`
	StartDateInMillis  *EpochMillis            `json:"start_date_in_millis,omitempty"`
	Type               licensetype.LicenseType `json:"type"`
	Uid                string                  `json:"uid"`
}

// LicenseBuilder holds License struct and provides a builder API.
type LicenseBuilder struct {
	v *License
}

// NewLicense provides a builder for the License struct.
func NewLicense() *LicenseBuilder {
	r := LicenseBuilder{
		&License{},
	}

	return &r
}

// Build finalize the chain and returns the License struct
func (rb *LicenseBuilder) Build() License {
	return *rb.v
}

// ExpiryDateInMillis set the ExpiryDateInMillis property for LicenseBuilder.
func (rb *LicenseBuilder) ExpiryDateInMillis(expirydateinmillis EpochMillis) *LicenseBuilder {
	rb.v.ExpiryDateInMillis = expirydateinmillis
	return rb
}

// IssueDateInMillis set the IssueDateInMillis property for LicenseBuilder.
func (rb *LicenseBuilder) IssueDateInMillis(issuedateinmillis EpochMillis) *LicenseBuilder {
	rb.v.IssueDateInMillis = issuedateinmillis
	return rb
}

// IssuedTo set the IssuedTo property for LicenseBuilder.
func (rb *LicenseBuilder) IssuedTo(issuedto string) *LicenseBuilder {
	rb.v.IssuedTo = issuedto
	return rb
}

// Issuer set the Issuer property for LicenseBuilder.
func (rb *LicenseBuilder) Issuer(issuer string) *LicenseBuilder {
	rb.v.Issuer = issuer
	return rb
}

// MaxNodes set the MaxNodes property for LicenseBuilder.
func (rb *LicenseBuilder) MaxNodes(maxnodes int64) *LicenseBuilder {
	rb.v.MaxNodes = maxnodes
	return rb
}

// MaxResourceUnits set the MaxResourceUnits property for LicenseBuilder.
func (rb *LicenseBuilder) MaxResourceUnits(maxresourceunits int64) *LicenseBuilder {
	rb.v.MaxResourceUnits = &maxresourceunits
	return rb
}

// Signature set the Signature property for LicenseBuilder.
func (rb *LicenseBuilder) Signature(signature string) *LicenseBuilder {
	rb.v.Signature = signature
	return rb
}

// StartDateInMillis set the StartDateInMillis property for LicenseBuilder.
func (rb *LicenseBuilder) StartDateInMillis(startdateinmillis EpochMillis) *LicenseBuilder {
	rb.v.StartDateInMillis = &startdateinmillis
	return rb
}

// Type set the Type property for LicenseBuilder.
func (rb *LicenseBuilder) Type_(type_ licensetype.LicenseType) *LicenseBuilder {
	rb.v.Type = type_
	return rb
}

// Uid set the Uid property for LicenseBuilder.
func (rb *LicenseBuilder) Uid(uid string) *LicenseBuilder {
	rb.v.Uid = uid
	return rb
}
