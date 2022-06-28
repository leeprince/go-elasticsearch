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

// PatternCaptureTokenFilter type.
type PatternCaptureTokenFilter struct {
	Patterns         []string       `json:"patterns"`
	PreserveOriginal bool           `json:"preserve_original"`
	Type             string         `json:"type,omitempty"`
	Version          *VersionString `json:"version,omitempty"`
}

// PatternCaptureTokenFilterBuilder holds PatternCaptureTokenFilter struct and provides a builder API.
type PatternCaptureTokenFilterBuilder struct {
	v *PatternCaptureTokenFilter
}

// NewPatternCaptureTokenFilter provides a builder for the PatternCaptureTokenFilter struct.
func NewPatternCaptureTokenFilter() *PatternCaptureTokenFilterBuilder {
	r := PatternCaptureTokenFilterBuilder{
		&PatternCaptureTokenFilter{},
	}

	r.v.Type = "pattern_capture"

	return &r
}

// Build finalize the chain and returns the PatternCaptureTokenFilter struct
func (rb *PatternCaptureTokenFilterBuilder) Build() PatternCaptureTokenFilter {
	return *rb.v
}

// Patterns set the Patterns property for PatternCaptureTokenFilterBuilder.
func (rb *PatternCaptureTokenFilterBuilder) Patterns(patterns ...string) *PatternCaptureTokenFilterBuilder {
	rb.v.Patterns = append(rb.v.Patterns, patterns...)
	return rb
}

// PreserveOriginal set the PreserveOriginal property for PatternCaptureTokenFilterBuilder.
func (rb *PatternCaptureTokenFilterBuilder) PreserveOriginal(preserveoriginal bool) *PatternCaptureTokenFilterBuilder {
	rb.v.PreserveOriginal = preserveoriginal
	return rb
}

// Type set the Type property for PatternCaptureTokenFilterBuilder.

// Version set the Version property for PatternCaptureTokenFilterBuilder.
func (rb *PatternCaptureTokenFilterBuilder) Version(version VersionString) *PatternCaptureTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
