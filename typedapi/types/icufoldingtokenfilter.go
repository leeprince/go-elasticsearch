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

// IcuFoldingTokenFilter type.
type IcuFoldingTokenFilter struct {
	Type             string         `json:"type,omitempty"`
	UnicodeSetFilter string         `json:"unicode_set_filter"`
	Version          *VersionString `json:"version,omitempty"`
}

// IcuFoldingTokenFilterBuilder holds IcuFoldingTokenFilter struct and provides a builder API.
type IcuFoldingTokenFilterBuilder struct {
	v *IcuFoldingTokenFilter
}

// NewIcuFoldingTokenFilter provides a builder for the IcuFoldingTokenFilter struct.
func NewIcuFoldingTokenFilter() *IcuFoldingTokenFilterBuilder {
	r := IcuFoldingTokenFilterBuilder{
		&IcuFoldingTokenFilter{},
	}

	r.v.Type = "icu_folding"

	return &r
}

// Build finalize the chain and returns the IcuFoldingTokenFilter struct
func (rb *IcuFoldingTokenFilterBuilder) Build() IcuFoldingTokenFilter {
	return *rb.v
}

// Type set the Type property for IcuFoldingTokenFilterBuilder.

// UnicodeSetFilter set the UnicodeSetFilter property for IcuFoldingTokenFilterBuilder.
func (rb *IcuFoldingTokenFilterBuilder) UnicodeSetFilter(unicodesetfilter string) *IcuFoldingTokenFilterBuilder {
	rb.v.UnicodeSetFilter = unicodesetfilter
	return rb
}

// Version set the Version property for IcuFoldingTokenFilterBuilder.
func (rb *IcuFoldingTokenFilterBuilder) Version(version VersionString) *IcuFoldingTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
