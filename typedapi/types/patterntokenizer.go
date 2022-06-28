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

// PatternTokenizer type.
type PatternTokenizer struct {
	Flags   string         `json:"flags"`
	Group   int            `json:"group"`
	Pattern string         `json:"pattern"`
	Type    string         `json:"type,omitempty"`
	Version *VersionString `json:"version,omitempty"`
}

// PatternTokenizerBuilder holds PatternTokenizer struct and provides a builder API.
type PatternTokenizerBuilder struct {
	v *PatternTokenizer
}

// NewPatternTokenizer provides a builder for the PatternTokenizer struct.
func NewPatternTokenizer() *PatternTokenizerBuilder {
	r := PatternTokenizerBuilder{
		&PatternTokenizer{},
	}

	r.v.Type = "pattern"

	return &r
}

// Build finalize the chain and returns the PatternTokenizer struct
func (rb *PatternTokenizerBuilder) Build() PatternTokenizer {
	return *rb.v
}

// Flags set the Flags property for PatternTokenizerBuilder.
func (rb *PatternTokenizerBuilder) Flags(flags string) *PatternTokenizerBuilder {
	rb.v.Flags = flags
	return rb
}

// Group set the Group property for PatternTokenizerBuilder.
func (rb *PatternTokenizerBuilder) Group(group int) *PatternTokenizerBuilder {
	rb.v.Group = group
	return rb
}

// Pattern set the Pattern property for PatternTokenizerBuilder.
func (rb *PatternTokenizerBuilder) Pattern(pattern string) *PatternTokenizerBuilder {
	rb.v.Pattern = pattern
	return rb
}

// Type set the Type property for PatternTokenizerBuilder.

// Version set the Version property for PatternTokenizerBuilder.
func (rb *PatternTokenizerBuilder) Version(version VersionString) *PatternTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
