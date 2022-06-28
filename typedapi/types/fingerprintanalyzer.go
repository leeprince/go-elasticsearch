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

// FingerprintAnalyzer type.
type FingerprintAnalyzer struct {
	MaxOutputSize    int            `json:"max_output_size"`
	PreserveOriginal bool           `json:"preserve_original"`
	Separator        string         `json:"separator"`
	Stopwords        *StopWords     `json:"stopwords,omitempty"`
	StopwordsPath    *string        `json:"stopwords_path,omitempty"`
	Type             string         `json:"type,omitempty"`
	Version          *VersionString `json:"version,omitempty"`
}

// FingerprintAnalyzerBuilder holds FingerprintAnalyzer struct and provides a builder API.
type FingerprintAnalyzerBuilder struct {
	v *FingerprintAnalyzer
}

// NewFingerprintAnalyzer provides a builder for the FingerprintAnalyzer struct.
func NewFingerprintAnalyzer() *FingerprintAnalyzerBuilder {
	r := FingerprintAnalyzerBuilder{
		&FingerprintAnalyzer{},
	}

	r.v.Type = "fingerprint"

	return &r
}

// Build finalize the chain and returns the FingerprintAnalyzer struct
func (rb *FingerprintAnalyzerBuilder) Build() FingerprintAnalyzer {
	return *rb.v
}

// MaxOutputSize set the MaxOutputSize property for FingerprintAnalyzerBuilder.
func (rb *FingerprintAnalyzerBuilder) MaxOutputSize(maxoutputsize int) *FingerprintAnalyzerBuilder {
	rb.v.MaxOutputSize = maxoutputsize
	return rb
}

// PreserveOriginal set the PreserveOriginal property for FingerprintAnalyzerBuilder.
func (rb *FingerprintAnalyzerBuilder) PreserveOriginal(preserveoriginal bool) *FingerprintAnalyzerBuilder {
	rb.v.PreserveOriginal = preserveoriginal
	return rb
}

// Separator set the Separator property for FingerprintAnalyzerBuilder.
func (rb *FingerprintAnalyzerBuilder) Separator(separator string) *FingerprintAnalyzerBuilder {
	rb.v.Separator = separator
	return rb
}

// Stopwords set the Stopwords property for FingerprintAnalyzerBuilder.
func (rb *FingerprintAnalyzerBuilder) Stopwords(stopwords StopWords) *FingerprintAnalyzerBuilder {
	rb.v.Stopwords = &stopwords
	return rb
}

// StopwordsPath set the StopwordsPath property for FingerprintAnalyzerBuilder.
func (rb *FingerprintAnalyzerBuilder) StopwordsPath(stopwordspath string) *FingerprintAnalyzerBuilder {
	rb.v.StopwordsPath = &stopwordspath
	return rb
}

// Type set the Type property for FingerprintAnalyzerBuilder.

// Version set the Version property for FingerprintAnalyzerBuilder.
func (rb *FingerprintAnalyzerBuilder) Version(version VersionString) *FingerprintAnalyzerBuilder {
	rb.v.Version = &version
	return rb
}
