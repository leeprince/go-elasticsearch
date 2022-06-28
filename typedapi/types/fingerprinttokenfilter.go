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

// FingerprintTokenFilter type.
type FingerprintTokenFilter struct {
	MaxOutputSize int            `json:"max_output_size"`
	Separator     string         `json:"separator"`
	Type          string         `json:"type,omitempty"`
	Version       *VersionString `json:"version,omitempty"`
}

// FingerprintTokenFilterBuilder holds FingerprintTokenFilter struct and provides a builder API.
type FingerprintTokenFilterBuilder struct {
	v *FingerprintTokenFilter
}

// NewFingerprintTokenFilter provides a builder for the FingerprintTokenFilter struct.
func NewFingerprintTokenFilter() *FingerprintTokenFilterBuilder {
	r := FingerprintTokenFilterBuilder{
		&FingerprintTokenFilter{},
	}

	r.v.Type = "fingerprint"

	return &r
}

// Build finalize the chain and returns the FingerprintTokenFilter struct
func (rb *FingerprintTokenFilterBuilder) Build() FingerprintTokenFilter {
	return *rb.v
}

// MaxOutputSize set the MaxOutputSize property for FingerprintTokenFilterBuilder.
func (rb *FingerprintTokenFilterBuilder) MaxOutputSize(maxoutputsize int) *FingerprintTokenFilterBuilder {
	rb.v.MaxOutputSize = maxoutputsize
	return rb
}

// Separator set the Separator property for FingerprintTokenFilterBuilder.
func (rb *FingerprintTokenFilterBuilder) Separator(separator string) *FingerprintTokenFilterBuilder {
	rb.v.Separator = separator
	return rb
}

// Type set the Type property for FingerprintTokenFilterBuilder.

// Version set the Version property for FingerprintTokenFilterBuilder.
func (rb *FingerprintTokenFilterBuilder) Version(version VersionString) *FingerprintTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
