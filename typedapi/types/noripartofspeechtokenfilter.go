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

// NoriPartOfSpeechTokenFilter type.
type NoriPartOfSpeechTokenFilter struct {
	Stoptags []string       `json:"stoptags"`
	Type     string         `json:"type,omitempty"`
	Version  *VersionString `json:"version,omitempty"`
}

// NoriPartOfSpeechTokenFilterBuilder holds NoriPartOfSpeechTokenFilter struct and provides a builder API.
type NoriPartOfSpeechTokenFilterBuilder struct {
	v *NoriPartOfSpeechTokenFilter
}

// NewNoriPartOfSpeechTokenFilter provides a builder for the NoriPartOfSpeechTokenFilter struct.
func NewNoriPartOfSpeechTokenFilter() *NoriPartOfSpeechTokenFilterBuilder {
	r := NoriPartOfSpeechTokenFilterBuilder{
		&NoriPartOfSpeechTokenFilter{},
	}

	r.v.Type = "nori_part_of_speech"

	return &r
}

// Build finalize the chain and returns the NoriPartOfSpeechTokenFilter struct
func (rb *NoriPartOfSpeechTokenFilterBuilder) Build() NoriPartOfSpeechTokenFilter {
	return *rb.v
}

// Stoptags set the Stoptags property for NoriPartOfSpeechTokenFilterBuilder.
func (rb *NoriPartOfSpeechTokenFilterBuilder) Stoptags(stoptags ...string) *NoriPartOfSpeechTokenFilterBuilder {
	rb.v.Stoptags = append(rb.v.Stoptags, stoptags...)
	return rb
}

// Type set the Type property for NoriPartOfSpeechTokenFilterBuilder.

// Version set the Version property for NoriPartOfSpeechTokenFilterBuilder.
func (rb *NoriPartOfSpeechTokenFilterBuilder) Version(version VersionString) *NoriPartOfSpeechTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
