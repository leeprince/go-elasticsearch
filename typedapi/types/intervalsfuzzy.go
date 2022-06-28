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

// IntervalsFuzzy type.
type IntervalsFuzzy struct {
	Analyzer       *string    `json:"analyzer,omitempty"`
	Fuzziness      *Fuzziness `json:"fuzziness,omitempty"`
	PrefixLength   *int       `json:"prefix_length,omitempty"`
	Term           string     `json:"term"`
	Transpositions *bool      `json:"transpositions,omitempty"`
	UseField       *Field     `json:"use_field,omitempty"`
}

// IntervalsFuzzyBuilder holds IntervalsFuzzy struct and provides a builder API.
type IntervalsFuzzyBuilder struct {
	v *IntervalsFuzzy
}

// NewIntervalsFuzzy provides a builder for the IntervalsFuzzy struct.
func NewIntervalsFuzzy() *IntervalsFuzzyBuilder {
	r := IntervalsFuzzyBuilder{
		&IntervalsFuzzy{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsFuzzy struct
func (rb *IntervalsFuzzyBuilder) Build() IntervalsFuzzy {
	return *rb.v
}

// Analyzer set the Analyzer property for IntervalsFuzzyBuilder.
func (rb *IntervalsFuzzyBuilder) Analyzer(analyzer string) *IntervalsFuzzyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Fuzziness set the Fuzziness property for IntervalsFuzzyBuilder.
func (rb *IntervalsFuzzyBuilder) Fuzziness(fuzziness Fuzziness) *IntervalsFuzzyBuilder {
	rb.v.Fuzziness = &fuzziness
	return rb
}

// PrefixLength set the PrefixLength property for IntervalsFuzzyBuilder.
func (rb *IntervalsFuzzyBuilder) PrefixLength(prefixlength int) *IntervalsFuzzyBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

// Term set the Term property for IntervalsFuzzyBuilder.
func (rb *IntervalsFuzzyBuilder) Term(term string) *IntervalsFuzzyBuilder {
	rb.v.Term = term
	return rb
}

// Transpositions set the Transpositions property for IntervalsFuzzyBuilder.
func (rb *IntervalsFuzzyBuilder) Transpositions(transpositions bool) *IntervalsFuzzyBuilder {
	rb.v.Transpositions = &transpositions
	return rb
}

// UseField set the UseField property for IntervalsFuzzyBuilder.
func (rb *IntervalsFuzzyBuilder) UseField(usefield Field) *IntervalsFuzzyBuilder {
	rb.v.UseField = &usefield
	return rb
}
