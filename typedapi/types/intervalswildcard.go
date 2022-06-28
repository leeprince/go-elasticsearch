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

// IntervalsWildcard type.
type IntervalsWildcard struct {
	Analyzer *string `json:"analyzer,omitempty"`
	Pattern  string  `json:"pattern"`
	UseField *Field  `json:"use_field,omitempty"`
}

// IntervalsWildcardBuilder holds IntervalsWildcard struct and provides a builder API.
type IntervalsWildcardBuilder struct {
	v *IntervalsWildcard
}

// NewIntervalsWildcard provides a builder for the IntervalsWildcard struct.
func NewIntervalsWildcard() *IntervalsWildcardBuilder {
	r := IntervalsWildcardBuilder{
		&IntervalsWildcard{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsWildcard struct
func (rb *IntervalsWildcardBuilder) Build() IntervalsWildcard {
	return *rb.v
}

// Analyzer set the Analyzer property for IntervalsWildcardBuilder.
func (rb *IntervalsWildcardBuilder) Analyzer(analyzer string) *IntervalsWildcardBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Pattern set the Pattern property for IntervalsWildcardBuilder.
func (rb *IntervalsWildcardBuilder) Pattern(pattern string) *IntervalsWildcardBuilder {
	rb.v.Pattern = pattern
	return rb
}

// UseField set the UseField property for IntervalsWildcardBuilder.
func (rb *IntervalsWildcardBuilder) UseField(usefield Field) *IntervalsWildcardBuilder {
	rb.v.UseField = &usefield
	return rb
}
