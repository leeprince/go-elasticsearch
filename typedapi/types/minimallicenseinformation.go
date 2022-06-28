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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensestatus"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/licensetype"
)

// MinimalLicenseInformation type.
type MinimalLicenseInformation struct {
	ExpiryDateInMillis EpochMillis                 `json:"expiry_date_in_millis"`
	Mode               licensetype.LicenseType     `json:"mode"`
	Status             licensestatus.LicenseStatus `json:"status"`
	Type               licensetype.LicenseType     `json:"type"`
	Uid                string                      `json:"uid"`
}

// MinimalLicenseInformationBuilder holds MinimalLicenseInformation struct and provides a builder API.
type MinimalLicenseInformationBuilder struct {
	v *MinimalLicenseInformation
}

// NewMinimalLicenseInformation provides a builder for the MinimalLicenseInformation struct.
func NewMinimalLicenseInformation() *MinimalLicenseInformationBuilder {
	r := MinimalLicenseInformationBuilder{
		&MinimalLicenseInformation{},
	}

	return &r
}

// Build finalize the chain and returns the MinimalLicenseInformation struct
func (rb *MinimalLicenseInformationBuilder) Build() MinimalLicenseInformation {
	return *rb.v
}

// ExpiryDateInMillis set the ExpiryDateInMillis property for MinimalLicenseInformationBuilder.
func (rb *MinimalLicenseInformationBuilder) ExpiryDateInMillis(expirydateinmillis EpochMillis) *MinimalLicenseInformationBuilder {
	rb.v.ExpiryDateInMillis = expirydateinmillis
	return rb
}

// Mode set the Mode property for MinimalLicenseInformationBuilder.
func (rb *MinimalLicenseInformationBuilder) Mode(mode licensetype.LicenseType) *MinimalLicenseInformationBuilder {
	rb.v.Mode = mode
	return rb
}

// Status set the Status property for MinimalLicenseInformationBuilder.
func (rb *MinimalLicenseInformationBuilder) Status(status licensestatus.LicenseStatus) *MinimalLicenseInformationBuilder {
	rb.v.Status = status
	return rb
}

// Type set the Type property for MinimalLicenseInformationBuilder.
func (rb *MinimalLicenseInformationBuilder) Type_(type_ licensetype.LicenseType) *MinimalLicenseInformationBuilder {
	rb.v.Type = type_
	return rb
}

// Uid set the Uid property for MinimalLicenseInformationBuilder.
func (rb *MinimalLicenseInformationBuilder) Uid(uid string) *MinimalLicenseInformationBuilder {
	rb.v.Uid = uid
	return rb
}
