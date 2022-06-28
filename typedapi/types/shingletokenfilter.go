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

// ShingleTokenFilter type.
type ShingleTokenFilter struct {
	FillerToken                *string        `json:"filler_token,omitempty"`
	MaxShingleSize             string         `json:"max_shingle_size,omitempty"`
	MinShingleSize             string         `json:"min_shingle_size,omitempty"`
	OutputUnigrams             *bool          `json:"output_unigrams,omitempty"`
	OutputUnigramsIfNoShingles *bool          `json:"output_unigrams_if_no_shingles,omitempty"`
	TokenSeparator             *string        `json:"token_separator,omitempty"`
	Type                       string         `json:"type,omitempty"`
	Version                    *VersionString `json:"version,omitempty"`
}

// ShingleTokenFilterBuilder holds ShingleTokenFilter struct and provides a builder API.
type ShingleTokenFilterBuilder struct {
	v *ShingleTokenFilter
}

// NewShingleTokenFilter provides a builder for the ShingleTokenFilter struct.
func NewShingleTokenFilter() *ShingleTokenFilterBuilder {
	r := ShingleTokenFilterBuilder{
		&ShingleTokenFilter{},
	}

	r.v.Type = "shingle"

	return &r
}

// Build finalize the chain and returns the ShingleTokenFilter struct
func (rb *ShingleTokenFilterBuilder) Build() ShingleTokenFilter {
	return *rb.v
}

// FillerToken set the FillerToken property for ShingleTokenFilterBuilder.
func (rb *ShingleTokenFilterBuilder) FillerToken(fillertoken string) *ShingleTokenFilterBuilder {
	rb.v.FillerToken = &fillertoken
	return rb
}

// MaxShingleSize set the MaxShingleSize property for ShingleTokenFilterBuilder.
func (rb *ShingleTokenFilterBuilder) MaxShingleSize(arg string) *ShingleTokenFilterBuilder {
	rb.v.MaxShingleSize = arg
	return rb
}

// MinShingleSize set the MinShingleSize property for ShingleTokenFilterBuilder.
func (rb *ShingleTokenFilterBuilder) MinShingleSize(arg string) *ShingleTokenFilterBuilder {
	rb.v.MinShingleSize = arg
	return rb
}

// OutputUnigrams set the OutputUnigrams property for ShingleTokenFilterBuilder.
func (rb *ShingleTokenFilterBuilder) OutputUnigrams(outputunigrams bool) *ShingleTokenFilterBuilder {
	rb.v.OutputUnigrams = &outputunigrams
	return rb
}

// OutputUnigramsIfNoShingles set the OutputUnigramsIfNoShingles property for ShingleTokenFilterBuilder.
func (rb *ShingleTokenFilterBuilder) OutputUnigramsIfNoShingles(outputunigramsifnoshingles bool) *ShingleTokenFilterBuilder {
	rb.v.OutputUnigramsIfNoShingles = &outputunigramsifnoshingles
	return rb
}

// TokenSeparator set the TokenSeparator property for ShingleTokenFilterBuilder.
func (rb *ShingleTokenFilterBuilder) TokenSeparator(tokenseparator string) *ShingleTokenFilterBuilder {
	rb.v.TokenSeparator = &tokenseparator
	return rb
}

// Type set the Type property for ShingleTokenFilterBuilder.

// Version set the Version property for ShingleTokenFilterBuilder.
func (rb *ShingleTokenFilterBuilder) Version(version VersionString) *ShingleTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
