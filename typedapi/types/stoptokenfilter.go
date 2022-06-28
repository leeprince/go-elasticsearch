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

// StopTokenFilter type.
type StopTokenFilter struct {
	IgnoreCase     *bool          `json:"ignore_case,omitempty"`
	RemoveTrailing *bool          `json:"remove_trailing,omitempty"`
	Stopwords      StopWords      `json:"stopwords"`
	StopwordsPath  *string        `json:"stopwords_path,omitempty"`
	Type           string         `json:"type,omitempty"`
	Version        *VersionString `json:"version,omitempty"`
}

// StopTokenFilterBuilder holds StopTokenFilter struct and provides a builder API.
type StopTokenFilterBuilder struct {
	v *StopTokenFilter
}

// NewStopTokenFilter provides a builder for the StopTokenFilter struct.
func NewStopTokenFilter() *StopTokenFilterBuilder {
	r := StopTokenFilterBuilder{
		&StopTokenFilter{},
	}

	r.v.Type = "stop"

	return &r
}

// Build finalize the chain and returns the StopTokenFilter struct
func (rb *StopTokenFilterBuilder) Build() StopTokenFilter {
	return *rb.v
}

// IgnoreCase set the IgnoreCase property for StopTokenFilterBuilder.
func (rb *StopTokenFilterBuilder) IgnoreCase(ignorecase bool) *StopTokenFilterBuilder {
	rb.v.IgnoreCase = &ignorecase
	return rb
}

// RemoveTrailing set the RemoveTrailing property for StopTokenFilterBuilder.
func (rb *StopTokenFilterBuilder) RemoveTrailing(removetrailing bool) *StopTokenFilterBuilder {
	rb.v.RemoveTrailing = &removetrailing
	return rb
}

// Stopwords set the Stopwords property for StopTokenFilterBuilder.
func (rb *StopTokenFilterBuilder) Stopwords(stopwords StopWords) *StopTokenFilterBuilder {
	rb.v.Stopwords = stopwords
	return rb
}

// StopwordsPath set the StopwordsPath property for StopTokenFilterBuilder.
func (rb *StopTokenFilterBuilder) StopwordsPath(stopwordspath string) *StopTokenFilterBuilder {
	rb.v.StopwordsPath = &stopwordspath
	return rb
}

// Type set the Type property for StopTokenFilterBuilder.

// Version set the Version property for StopTokenFilterBuilder.
func (rb *StopTokenFilterBuilder) Version(version VersionString) *StopTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
