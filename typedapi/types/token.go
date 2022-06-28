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

// Token type.
type Token struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TokenBuilder holds Token struct and provides a builder API.
type TokenBuilder struct {
	v *Token
}

// NewToken provides a builder for the Token struct.
func NewToken() *TokenBuilder {
	r := TokenBuilder{
		&Token{},
	}

	return &r
}

// Build finalize the chain and returns the Token struct
func (rb *TokenBuilder) Build() Token {
	return *rb.v
}

// Name set the Name property for TokenBuilder.
func (rb *TokenBuilder) Name(name string) *TokenBuilder {
	rb.v.Name = name
	return rb
}

// Value set the Value property for TokenBuilder.
func (rb *TokenBuilder) Value(value string) *TokenBuilder {
	rb.v.Value = value
	return rb
}
