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

// HunspellTokenFilter type.
type HunspellTokenFilter struct {
	Dedup       bool           `json:"dedup"`
	Dictionary  string         `json:"dictionary"`
	Locale      string         `json:"locale"`
	LongestOnly bool           `json:"longest_only"`
	Type        string         `json:"type,omitempty"`
	Version     *VersionString `json:"version,omitempty"`
}

// HunspellTokenFilterBuilder holds HunspellTokenFilter struct and provides a builder API.
type HunspellTokenFilterBuilder struct {
	v *HunspellTokenFilter
}

// NewHunspellTokenFilter provides a builder for the HunspellTokenFilter struct.
func NewHunspellTokenFilter() *HunspellTokenFilterBuilder {
	r := HunspellTokenFilterBuilder{
		&HunspellTokenFilter{},
	}

	r.v.Type = "hunspell"

	return &r
}

// Build finalize the chain and returns the HunspellTokenFilter struct
func (rb *HunspellTokenFilterBuilder) Build() HunspellTokenFilter {
	return *rb.v
}

// Dedup set the Dedup property for HunspellTokenFilterBuilder.
func (rb *HunspellTokenFilterBuilder) Dedup(dedup bool) *HunspellTokenFilterBuilder {
	rb.v.Dedup = dedup
	return rb
}

// Dictionary set the Dictionary property for HunspellTokenFilterBuilder.
func (rb *HunspellTokenFilterBuilder) Dictionary(dictionary string) *HunspellTokenFilterBuilder {
	rb.v.Dictionary = dictionary
	return rb
}

// Locale set the Locale property for HunspellTokenFilterBuilder.
func (rb *HunspellTokenFilterBuilder) Locale(locale string) *HunspellTokenFilterBuilder {
	rb.v.Locale = locale
	return rb
}

// LongestOnly set the LongestOnly property for HunspellTokenFilterBuilder.
func (rb *HunspellTokenFilterBuilder) LongestOnly(longestonly bool) *HunspellTokenFilterBuilder {
	rb.v.LongestOnly = longestonly
	return rb
}

// Type set the Type property for HunspellTokenFilterBuilder.

// Version set the Version property for HunspellTokenFilterBuilder.
func (rb *HunspellTokenFilterBuilder) Version(version VersionString) *HunspellTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
