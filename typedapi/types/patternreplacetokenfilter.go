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

// PatternReplaceTokenFilter type.
type PatternReplaceTokenFilter struct {
	Flags       string         `json:"flags"`
	Pattern     string         `json:"pattern"`
	Replacement string         `json:"replacement"`
	Type        string         `json:"type,omitempty"`
	Version     *VersionString `json:"version,omitempty"`
}

// PatternReplaceTokenFilterBuilder holds PatternReplaceTokenFilter struct and provides a builder API.
type PatternReplaceTokenFilterBuilder struct {
	v *PatternReplaceTokenFilter
}

// NewPatternReplaceTokenFilter provides a builder for the PatternReplaceTokenFilter struct.
func NewPatternReplaceTokenFilter() *PatternReplaceTokenFilterBuilder {
	r := PatternReplaceTokenFilterBuilder{
		&PatternReplaceTokenFilter{},
	}

	r.v.Type = "pattern_replace"

	return &r
}

// Build finalize the chain and returns the PatternReplaceTokenFilter struct
func (rb *PatternReplaceTokenFilterBuilder) Build() PatternReplaceTokenFilter {
	return *rb.v
}

// Flags set the Flags property for PatternReplaceTokenFilterBuilder.
func (rb *PatternReplaceTokenFilterBuilder) Flags(flags string) *PatternReplaceTokenFilterBuilder {
	rb.v.Flags = flags
	return rb
}

// Pattern set the Pattern property for PatternReplaceTokenFilterBuilder.
func (rb *PatternReplaceTokenFilterBuilder) Pattern(pattern string) *PatternReplaceTokenFilterBuilder {
	rb.v.Pattern = pattern
	return rb
}

// Replacement set the Replacement property for PatternReplaceTokenFilterBuilder.
func (rb *PatternReplaceTokenFilterBuilder) Replacement(replacement string) *PatternReplaceTokenFilterBuilder {
	rb.v.Replacement = replacement
	return rb
}

// Type set the Type property for PatternReplaceTokenFilterBuilder.

// Version set the Version property for PatternReplaceTokenFilterBuilder.
func (rb *PatternReplaceTokenFilterBuilder) Version(version VersionString) *PatternReplaceTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
