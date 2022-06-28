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

// PredicateTokenFilter type.
type PredicateTokenFilter struct {
	Script  Script         `json:"script"`
	Type    string         `json:"type,omitempty"`
	Version *VersionString `json:"version,omitempty"`
}

// PredicateTokenFilterBuilder holds PredicateTokenFilter struct and provides a builder API.
type PredicateTokenFilterBuilder struct {
	v *PredicateTokenFilter
}

// NewPredicateTokenFilter provides a builder for the PredicateTokenFilter struct.
func NewPredicateTokenFilter() *PredicateTokenFilterBuilder {
	r := PredicateTokenFilterBuilder{
		&PredicateTokenFilter{},
	}

	r.v.Type = "predicate_token_filter"

	return &r
}

// Build finalize the chain and returns the PredicateTokenFilter struct
func (rb *PredicateTokenFilterBuilder) Build() PredicateTokenFilter {
	return *rb.v
}

// Script set the Script property for PredicateTokenFilterBuilder.
func (rb *PredicateTokenFilterBuilder) Script(script Script) *PredicateTokenFilterBuilder {
	rb.v.Script = script
	return rb
}

// Type set the Type property for PredicateTokenFilterBuilder.

// Version set the Version property for PredicateTokenFilterBuilder.
func (rb *PredicateTokenFilterBuilder) Version(version VersionString) *PredicateTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
