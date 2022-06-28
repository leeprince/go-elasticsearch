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

// AnalyzerDetail type.
type AnalyzerDetail struct {
	Name   string                `json:"name"`
	Tokens []ExplainAnalyzeToken `json:"tokens"`
}

// AnalyzerDetailBuilder holds AnalyzerDetail struct and provides a builder API.
type AnalyzerDetailBuilder struct {
	v *AnalyzerDetail
}

// NewAnalyzerDetail provides a builder for the AnalyzerDetail struct.
func NewAnalyzerDetail() *AnalyzerDetailBuilder {
	r := AnalyzerDetailBuilder{
		&AnalyzerDetail{},
	}

	return &r
}

// Build finalize the chain and returns the AnalyzerDetail struct
func (rb *AnalyzerDetailBuilder) Build() AnalyzerDetail {
	return *rb.v
}

// Name set the Name property for AnalyzerDetailBuilder.
func (rb *AnalyzerDetailBuilder) Name(name string) *AnalyzerDetailBuilder {
	rb.v.Name = name
	return rb
}

// Tokens set the Tokens property for AnalyzerDetailBuilder.
func (rb *AnalyzerDetailBuilder) Tokens(tokens ...ExplainAnalyzeToken) *AnalyzerDetailBuilder {
	rb.v.Tokens = append(rb.v.Tokens, tokens...)
	return rb
}
