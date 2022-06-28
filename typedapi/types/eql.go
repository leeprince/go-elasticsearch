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

// Eql type.
type Eql struct {
	Available bool             `json:"available"`
	Enabled   bool             `json:"enabled"`
	Features  EqlFeatures      `json:"features"`
	Queries   map[string]Query `json:"queries"`
}

// EqlBuilder holds Eql struct and provides a builder API.
type EqlBuilder struct {
	v *Eql
}

// NewEql provides a builder for the Eql struct.
func NewEql() *EqlBuilder {
	r := EqlBuilder{
		&Eql{
			Queries: make(map[string]Query, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Eql struct
func (rb *EqlBuilder) Build() Eql {
	return *rb.v
}

// Available set the Available property for EqlBuilder.
func (rb *EqlBuilder) Available(available bool) *EqlBuilder {
	rb.v.Available = available
	return rb
}

// Enabled set the Enabled property for EqlBuilder.
func (rb *EqlBuilder) Enabled(enabled bool) *EqlBuilder {
	rb.v.Enabled = enabled
	return rb
}

// Features set the Features property for EqlBuilder.
func (rb *EqlBuilder) Features(features EqlFeatures) *EqlBuilder {
	rb.v.Features = features
	return rb
}

// Queries set the Queries property for EqlBuilder.
func (rb *EqlBuilder) Queries(value map[string]Query) *EqlBuilder {
	rb.v.Queries = value
	return rb
}
