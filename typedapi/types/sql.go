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

// Sql type.
type Sql struct {
	Available bool             `json:"available"`
	Enabled   bool             `json:"enabled"`
	Features  map[string]int   `json:"features"`
	Queries   map[string]Query `json:"queries"`
}

// SqlBuilder holds Sql struct and provides a builder API.
type SqlBuilder struct {
	v *Sql
}

// NewSql provides a builder for the Sql struct.
func NewSql() *SqlBuilder {
	r := SqlBuilder{
		&Sql{
			Features: make(map[string]int, 0),
			Queries:  make(map[string]Query, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Sql struct
func (rb *SqlBuilder) Build() Sql {
	return *rb.v
}

// Available set the Available property for SqlBuilder.
func (rb *SqlBuilder) Available(available bool) *SqlBuilder {
	rb.v.Available = available
	return rb
}

// Enabled set the Enabled property for SqlBuilder.
func (rb *SqlBuilder) Enabled(enabled bool) *SqlBuilder {
	rb.v.Enabled = enabled
	return rb
}

// Features set the Features property for SqlBuilder.
func (rb *SqlBuilder) Features(value map[string]int) *SqlBuilder {
	rb.v.Features = value
	return rb
}

// Queries set the Queries property for SqlBuilder.
func (rb *SqlBuilder) Queries(value map[string]Query) *SqlBuilder {
	rb.v.Queries = value
	return rb
}
