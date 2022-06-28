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

// CompletionContext type.
type CompletionContext struct {
	Boost      *float64           `json:"boost,omitempty"`
	Context    Context            `json:"context"`
	Neighbours []GeoHashPrecision `json:"neighbours,omitempty"`
	Precision  *GeoHashPrecision  `json:"precision,omitempty"`
	Prefix     *bool              `json:"prefix,omitempty"`
}

// CompletionContextBuilder holds CompletionContext struct and provides a builder API.
type CompletionContextBuilder struct {
	v *CompletionContext
}

// NewCompletionContext provides a builder for the CompletionContext struct.
func NewCompletionContext() *CompletionContextBuilder {
	r := CompletionContextBuilder{
		&CompletionContext{},
	}

	return &r
}

// Build finalize the chain and returns the CompletionContext struct
func (rb *CompletionContextBuilder) Build() CompletionContext {
	return *rb.v
}

// Boost set the Boost property for CompletionContextBuilder.
func (rb *CompletionContextBuilder) Boost(boost float64) *CompletionContextBuilder {
	rb.v.Boost = &boost
	return rb
}

// Context set the Context property for CompletionContextBuilder.
func (rb *CompletionContextBuilder) Context(context Context) *CompletionContextBuilder {
	rb.v.Context = context
	return rb
}

// Neighbours set the Neighbours property for CompletionContextBuilder.
func (rb *CompletionContextBuilder) Neighbours(neighbours ...GeoHashPrecision) *CompletionContextBuilder {
	rb.v.Neighbours = append(rb.v.Neighbours, neighbours...)
	return rb
}

// Precision set the Precision property for CompletionContextBuilder.
func (rb *CompletionContextBuilder) Precision(precision GeoHashPrecision) *CompletionContextBuilder {
	rb.v.Precision = &precision
	return rb
}

// Prefix set the Prefix property for CompletionContextBuilder.
func (rb *CompletionContextBuilder) Prefix(prefix bool) *CompletionContextBuilder {
	rb.v.Prefix = &prefix
	return rb
}
