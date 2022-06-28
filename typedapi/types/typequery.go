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

// TypeQuery type.
type TypeQuery struct {
	Boost      *float32 `json:"boost,omitempty"`
	QueryName_ *string  `json:"_name,omitempty"`
	Value      string   `json:"value"`
}

// TypeQueryBuilder holds TypeQuery struct and provides a builder API.
type TypeQueryBuilder struct {
	v *TypeQuery
}

// NewTypeQuery provides a builder for the TypeQuery struct.
func NewTypeQuery() *TypeQueryBuilder {
	r := TypeQueryBuilder{
		&TypeQuery{},
	}

	return &r
}

// Build finalize the chain and returns the TypeQuery struct
func (rb *TypeQueryBuilder) Build() TypeQuery {
	return *rb.v
}

// Boost set the Boost property for TypeQueryBuilder.
func (rb *TypeQueryBuilder) Boost(boost float32) *TypeQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// QueryName_ set the QueryName_ property for TypeQueryBuilder.
func (rb *TypeQueryBuilder) QueryName_(queryname_ string) *TypeQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Value set the Value property for TypeQueryBuilder.
func (rb *TypeQueryBuilder) Value(value string) *TypeQueryBuilder {
	rb.v.Value = value
	return rb
}
