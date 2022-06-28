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

// SuggestFuzziness type.
type SuggestFuzziness struct {
	Fuzziness      *Fuzziness `json:"fuzziness,omitempty"`
	MinLength      *int       `json:"min_length,omitempty"`
	PrefixLength   *int       `json:"prefix_length,omitempty"`
	Transpositions *bool      `json:"transpositions,omitempty"`
	UnicodeAware   *bool      `json:"unicode_aware,omitempty"`
}

// SuggestFuzzinessBuilder holds SuggestFuzziness struct and provides a builder API.
type SuggestFuzzinessBuilder struct {
	v *SuggestFuzziness
}

// NewSuggestFuzziness provides a builder for the SuggestFuzziness struct.
func NewSuggestFuzziness() *SuggestFuzzinessBuilder {
	r := SuggestFuzzinessBuilder{
		&SuggestFuzziness{},
	}

	return &r
}

// Build finalize the chain and returns the SuggestFuzziness struct
func (rb *SuggestFuzzinessBuilder) Build() SuggestFuzziness {
	return *rb.v
}

// Fuzziness set the Fuzziness property for SuggestFuzzinessBuilder.
func (rb *SuggestFuzzinessBuilder) Fuzziness(fuzziness Fuzziness) *SuggestFuzzinessBuilder {
	rb.v.Fuzziness = &fuzziness
	return rb
}

// MinLength set the MinLength property for SuggestFuzzinessBuilder.
func (rb *SuggestFuzzinessBuilder) MinLength(minlength int) *SuggestFuzzinessBuilder {
	rb.v.MinLength = &minlength
	return rb
}

// PrefixLength set the PrefixLength property for SuggestFuzzinessBuilder.
func (rb *SuggestFuzzinessBuilder) PrefixLength(prefixlength int) *SuggestFuzzinessBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

// Transpositions set the Transpositions property for SuggestFuzzinessBuilder.
func (rb *SuggestFuzzinessBuilder) Transpositions(transpositions bool) *SuggestFuzzinessBuilder {
	rb.v.Transpositions = &transpositions
	return rb
}

// UnicodeAware set the UnicodeAware property for SuggestFuzzinessBuilder.
func (rb *SuggestFuzzinessBuilder) UnicodeAware(unicodeaware bool) *SuggestFuzzinessBuilder {
	rb.v.UnicodeAware = &unicodeaware
	return rb
}
