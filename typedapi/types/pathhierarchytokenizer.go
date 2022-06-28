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

// PathHierarchyTokenizer type.
type PathHierarchyTokenizer struct {
	BufferSize  int            `json:"buffer_size"`
	Delimiter   string         `json:"delimiter"`
	Replacement string         `json:"replacement"`
	Reverse     bool           `json:"reverse"`
	Skip        int            `json:"skip"`
	Type        string         `json:"type,omitempty"`
	Version     *VersionString `json:"version,omitempty"`
}

// PathHierarchyTokenizerBuilder holds PathHierarchyTokenizer struct and provides a builder API.
type PathHierarchyTokenizerBuilder struct {
	v *PathHierarchyTokenizer
}

// NewPathHierarchyTokenizer provides a builder for the PathHierarchyTokenizer struct.
func NewPathHierarchyTokenizer() *PathHierarchyTokenizerBuilder {
	r := PathHierarchyTokenizerBuilder{
		&PathHierarchyTokenizer{},
	}

	r.v.Type = "path_hierarchy"

	return &r
}

// Build finalize the chain and returns the PathHierarchyTokenizer struct
func (rb *PathHierarchyTokenizerBuilder) Build() PathHierarchyTokenizer {
	return *rb.v
}

// BufferSize set the BufferSize property for PathHierarchyTokenizerBuilder.
func (rb *PathHierarchyTokenizerBuilder) BufferSize(buffersize int) *PathHierarchyTokenizerBuilder {
	rb.v.BufferSize = buffersize
	return rb
}

// Delimiter set the Delimiter property for PathHierarchyTokenizerBuilder.
func (rb *PathHierarchyTokenizerBuilder) Delimiter(delimiter string) *PathHierarchyTokenizerBuilder {
	rb.v.Delimiter = delimiter
	return rb
}

// Replacement set the Replacement property for PathHierarchyTokenizerBuilder.
func (rb *PathHierarchyTokenizerBuilder) Replacement(replacement string) *PathHierarchyTokenizerBuilder {
	rb.v.Replacement = replacement
	return rb
}

// Reverse set the Reverse property for PathHierarchyTokenizerBuilder.
func (rb *PathHierarchyTokenizerBuilder) Reverse(reverse bool) *PathHierarchyTokenizerBuilder {
	rb.v.Reverse = reverse
	return rb
}

// Skip set the Skip property for PathHierarchyTokenizerBuilder.
func (rb *PathHierarchyTokenizerBuilder) Skip(skip int) *PathHierarchyTokenizerBuilder {
	rb.v.Skip = skip
	return rb
}

// Type set the Type property for PathHierarchyTokenizerBuilder.

// Version set the Version property for PathHierarchyTokenizerBuilder.
func (rb *PathHierarchyTokenizerBuilder) Version(version VersionString) *PathHierarchyTokenizerBuilder {
	rb.v.Version = &version
	return rb
}
