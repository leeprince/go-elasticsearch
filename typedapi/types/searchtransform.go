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

// SearchTransform type.
type SearchTransform struct {
	Request SearchInputRequestDefinition `json:"request"`
	Timeout Time                         `json:"timeout"`
}

// SearchTransformBuilder holds SearchTransform struct and provides a builder API.
type SearchTransformBuilder struct {
	v *SearchTransform
}

// NewSearchTransform provides a builder for the SearchTransform struct.
func NewSearchTransform() *SearchTransformBuilder {
	r := SearchTransformBuilder{
		&SearchTransform{},
	}

	return &r
}

// Build finalize the chain and returns the SearchTransform struct
func (rb *SearchTransformBuilder) Build() SearchTransform {
	return *rb.v
}

// Request set the Request property for SearchTransformBuilder.
func (rb *SearchTransformBuilder) Request(request SearchInputRequestDefinition) *SearchTransformBuilder {
	rb.v.Request = request
	return rb
}

// Timeout set the Timeout property for SearchTransformBuilder.
func (rb *SearchTransformBuilder) Timeout(timeout Time) *SearchTransformBuilder {
	rb.v.Timeout = timeout
	return rb
}
