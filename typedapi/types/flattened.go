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

// Flattened type.
type Flattened struct {
	Available  bool `json:"available"`
	Enabled    bool `json:"enabled"`
	FieldCount int  `json:"field_count"`
}

// FlattenedBuilder holds Flattened struct and provides a builder API.
type FlattenedBuilder struct {
	v *Flattened
}

// NewFlattened provides a builder for the Flattened struct.
func NewFlattened() *FlattenedBuilder {
	r := FlattenedBuilder{
		&Flattened{},
	}

	return &r
}

// Build finalize the chain and returns the Flattened struct
func (rb *FlattenedBuilder) Build() Flattened {
	return *rb.v
}

// Available set the Available property for FlattenedBuilder.
func (rb *FlattenedBuilder) Available(available bool) *FlattenedBuilder {
	rb.v.Available = available
	return rb
}

// Enabled set the Enabled property for FlattenedBuilder.
func (rb *FlattenedBuilder) Enabled(enabled bool) *FlattenedBuilder {
	rb.v.Enabled = enabled
	return rb
}

// FieldCount set the FieldCount property for FlattenedBuilder.
func (rb *FlattenedBuilder) FieldCount(fieldcount int) *FlattenedBuilder {
	rb.v.FieldCount = fieldcount
	return rb
}
