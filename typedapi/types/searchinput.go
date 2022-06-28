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

// SearchInput type.
type SearchInput struct {
	Extract []string                     `json:"extract,omitempty"`
	Request SearchInputRequestDefinition `json:"request"`
	Timeout *Time                        `json:"timeout,omitempty"`
}

// SearchInputBuilder holds SearchInput struct and provides a builder API.
type SearchInputBuilder struct {
	v *SearchInput
}

// NewSearchInput provides a builder for the SearchInput struct.
func NewSearchInput() *SearchInputBuilder {
	r := SearchInputBuilder{
		&SearchInput{},
	}

	return &r
}

// Build finalize the chain and returns the SearchInput struct
func (rb *SearchInputBuilder) Build() SearchInput {
	return *rb.v
}

// Extract set the Extract property for SearchInputBuilder.
func (rb *SearchInputBuilder) Extract(extract ...string) *SearchInputBuilder {
	rb.v.Extract = append(rb.v.Extract, extract...)
	return rb
}

// Request set the Request property for SearchInputBuilder.
func (rb *SearchInputBuilder) Request(request SearchInputRequestDefinition) *SearchInputBuilder {
	rb.v.Request = request
	return rb
}

// Timeout set the Timeout property for SearchInputBuilder.
func (rb *SearchInputBuilder) Timeout(timeout Time) *SearchInputBuilder {
	rb.v.Timeout = &timeout
	return rb
}
