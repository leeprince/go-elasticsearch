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

// Normalizer holds the union for the following types:
//     CustomNormalizer
//     LowercaseNormalizer
type Normalizer interface{}

// NormalizerBuilder holds Normalizer struct and provides a builder API.
type NormalizerBuilder struct {
	v Normalizer
}

// NewNormalizer provides a builder for the Normalizer struct.
func NewNormalizer() *NormalizerBuilder {
	return &NormalizerBuilder{}
}

// Build finalize the chain and returns the Normalizer struct
func (u *NormalizerBuilder) Build() Normalizer {
	return u.v
}

// CustomNormalizer set the CustomNormalizer property for NormalizerBuilder.
func (u *NormalizerBuilder) CustomNormalizer(v CustomNormalizer) *NormalizerBuilder {
	u.v = v
	return u
}

// LowercaseNormalizer set the LowercaseNormalizer property for NormalizerBuilder.
func (u *NormalizerBuilder) LowercaseNormalizer(v LowercaseNormalizer) *NormalizerBuilder {
	u.v = v
	return u
}
