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

// CategorizeTextAnalyzer holds the union for the following types:
//     CustomCategorizeTextAnalyzer
//     string
type CategorizeTextAnalyzer interface{}

// CategorizeTextAnalyzerBuilder holds CategorizeTextAnalyzer struct and provides a builder API.
type CategorizeTextAnalyzerBuilder struct {
	v CategorizeTextAnalyzer
}

// NewCategorizeTextAnalyzer provides a builder for the CategorizeTextAnalyzer struct.
func NewCategorizeTextAnalyzer() *CategorizeTextAnalyzerBuilder {
	return &CategorizeTextAnalyzerBuilder{}
}

// Build finalize the chain and returns the CategorizeTextAnalyzer struct
func (u *CategorizeTextAnalyzerBuilder) Build() CategorizeTextAnalyzer {
	return u.v
}

// CustomCategorizeTextAnalyzer set the CustomCategorizeTextAnalyzer property for CategorizeTextAnalyzerBuilder.
func (u *CategorizeTextAnalyzerBuilder) CustomCategorizeTextAnalyzer(v CustomCategorizeTextAnalyzer) *CategorizeTextAnalyzerBuilder {
	u.v = v
	return u
}

// String set the String property for CategorizeTextAnalyzerBuilder.
func (u *CategorizeTextAnalyzerBuilder) String(v string) *CategorizeTextAnalyzerBuilder {
	u.v = v
	return u
}
