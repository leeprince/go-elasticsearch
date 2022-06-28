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

// Analyzer holds the union for the following types:
//     CustomAnalyzer
//     DutchAnalyzer
//     FingerprintAnalyzer
//     IcuAnalyzer
//     KeywordAnalyzer
//     KuromojiAnalyzer
//     LanguageAnalyzer
//     NoriAnalyzer
//     PatternAnalyzer
//     SimpleAnalyzer
//     SnowballAnalyzer
//     StandardAnalyzer
//     StopAnalyzer
//     WhitespaceAnalyzer
type Analyzer interface{}

// AnalyzerBuilder holds Analyzer struct and provides a builder API.
type AnalyzerBuilder struct {
	v Analyzer
}

// NewAnalyzer provides a builder for the Analyzer struct.
func NewAnalyzer() *AnalyzerBuilder {
	return &AnalyzerBuilder{}
}

// Build finalize the chain and returns the Analyzer struct
func (u *AnalyzerBuilder) Build() Analyzer {
	return u.v
}

// CustomAnalyzer set the CustomAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) CustomAnalyzer(v CustomAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// DutchAnalyzer set the DutchAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) DutchAnalyzer(v DutchAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// FingerprintAnalyzer set the FingerprintAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) FingerprintAnalyzer(v FingerprintAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// IcuAnalyzer set the IcuAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) IcuAnalyzer(v IcuAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// KeywordAnalyzer set the KeywordAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) KeywordAnalyzer(v KeywordAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// KuromojiAnalyzer set the KuromojiAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) KuromojiAnalyzer(v KuromojiAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// LanguageAnalyzer set the LanguageAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) LanguageAnalyzer(v LanguageAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// NoriAnalyzer set the NoriAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) NoriAnalyzer(v NoriAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// PatternAnalyzer set the PatternAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) PatternAnalyzer(v PatternAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// SimpleAnalyzer set the SimpleAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) SimpleAnalyzer(v SimpleAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// SnowballAnalyzer set the SnowballAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) SnowballAnalyzer(v SnowballAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// StandardAnalyzer set the StandardAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) StandardAnalyzer(v StandardAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// StopAnalyzer set the StopAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) StopAnalyzer(v StopAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}

// WhitespaceAnalyzer set the WhitespaceAnalyzer property for AnalyzerBuilder.
func (u *AnalyzerBuilder) WhitespaceAnalyzer(v WhitespaceAnalyzer) *AnalyzerBuilder {
	u.v = v
	return u
}
