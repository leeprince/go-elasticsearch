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

// TokenDetail type.
type TokenDetail struct {
	Name   string                `json:"name"`
	Tokens []ExplainAnalyzeToken `json:"tokens"`
}

// TokenDetailBuilder holds TokenDetail struct and provides a builder API.
type TokenDetailBuilder struct {
	v *TokenDetail
}

// NewTokenDetail provides a builder for the TokenDetail struct.
func NewTokenDetail() *TokenDetailBuilder {
	r := TokenDetailBuilder{
		&TokenDetail{},
	}

	return &r
}

// Build finalize the chain and returns the TokenDetail struct
func (rb *TokenDetailBuilder) Build() TokenDetail {
	return *rb.v
}

// Name set the Name property for TokenDetailBuilder.
func (rb *TokenDetailBuilder) Name(name string) *TokenDetailBuilder {
	rb.v.Name = name
	return rb
}

// Tokens set the Tokens property for TokenDetailBuilder.
func (rb *TokenDetailBuilder) Tokens(tokens ...ExplainAnalyzeToken) *TokenDetailBuilder {
	rb.v.Tokens = append(rb.v.Tokens, tokens...)
	return rb
}
