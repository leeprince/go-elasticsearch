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

// Base type.
type Base struct {
	Available bool `json:"available"`
	Enabled   bool `json:"enabled"`
}

// BaseBuilder holds Base struct and provides a builder API.
type BaseBuilder struct {
	v *Base
}

// NewBase provides a builder for the Base struct.
func NewBase() *BaseBuilder {
	r := BaseBuilder{
		&Base{},
	}

	return &r
}

// Build finalize the chain and returns the Base struct
func (rb *BaseBuilder) Build() Base {
	return *rb.v
}

// Available set the Available property for BaseBuilder.
func (rb *BaseBuilder) Available(available bool) *BaseBuilder {
	rb.v.Available = available
	return rb
}

// Enabled set the Enabled property for BaseBuilder.
func (rb *BaseBuilder) Enabled(enabled bool) *BaseBuilder {
	rb.v.Enabled = enabled
	return rb
}
