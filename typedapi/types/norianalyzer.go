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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noridecompoundmode"
)

// NoriAnalyzer type.
type NoriAnalyzer struct {
	DecompoundMode *noridecompoundmode.NoriDecompoundMode `json:"decompound_mode,omitempty"`
	Stoptags       []string                               `json:"stoptags,omitempty"`
	Type           string                                 `json:"type,omitempty"`
	UserDictionary *string                                `json:"user_dictionary,omitempty"`
	Version        *VersionString                         `json:"version,omitempty"`
}

// NoriAnalyzerBuilder holds NoriAnalyzer struct and provides a builder API.
type NoriAnalyzerBuilder struct {
	v *NoriAnalyzer
}

// NewNoriAnalyzer provides a builder for the NoriAnalyzer struct.
func NewNoriAnalyzer() *NoriAnalyzerBuilder {
	r := NoriAnalyzerBuilder{
		&NoriAnalyzer{},
	}

	r.v.Type = "nori"

	return &r
}

// Build finalize the chain and returns the NoriAnalyzer struct
func (rb *NoriAnalyzerBuilder) Build() NoriAnalyzer {
	return *rb.v
}

// DecompoundMode set the DecompoundMode property for NoriAnalyzerBuilder.
func (rb *NoriAnalyzerBuilder) DecompoundMode(decompoundmode noridecompoundmode.NoriDecompoundMode) *NoriAnalyzerBuilder {
	rb.v.DecompoundMode = &decompoundmode
	return rb
}

// Stoptags set the Stoptags property for NoriAnalyzerBuilder.
func (rb *NoriAnalyzerBuilder) Stoptags(stoptags ...string) *NoriAnalyzerBuilder {
	rb.v.Stoptags = append(rb.v.Stoptags, stoptags...)
	return rb
}

// Type set the Type property for NoriAnalyzerBuilder.

// UserDictionary set the UserDictionary property for NoriAnalyzerBuilder.
func (rb *NoriAnalyzerBuilder) UserDictionary(userdictionary string) *NoriAnalyzerBuilder {
	rb.v.UserDictionary = &userdictionary
	return rb
}

// Version set the Version property for NoriAnalyzerBuilder.
func (rb *NoriAnalyzerBuilder) Version(version VersionString) *NoriAnalyzerBuilder {
	rb.v.Version = &version
	return rb
}
