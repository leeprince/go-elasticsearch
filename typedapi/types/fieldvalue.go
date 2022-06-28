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

// FieldValue holds the union for the following types:
//     bool
//     float64
//     int64
//     string
type FieldValue interface{}

// FieldValueBuilder holds FieldValue struct and provides a builder API.
type FieldValueBuilder struct {
	v FieldValue
}

// NewFieldValue provides a builder for the FieldValue struct.
func NewFieldValue() *FieldValueBuilder {
	return &FieldValueBuilder{}
}

// Build finalize the chain and returns the FieldValue struct
func (u *FieldValueBuilder) Build() FieldValue {
	return u.v
}

// Bool set the Bool property for FieldValueBuilder.
func (u *FieldValueBuilder) Bool(v bool) *FieldValueBuilder {
	u.v = v
	return u
}

// Float64 set the Float64 property for FieldValueBuilder.
func (u *FieldValueBuilder) Float64(v float64) *FieldValueBuilder {
	u.v = v
	return u
}

// Int64 set the Int64 property for FieldValueBuilder.
func (u *FieldValueBuilder) Int64(v int64) *FieldValueBuilder {
	u.v = v
	return u
}

// String set the String property for FieldValueBuilder.
func (u *FieldValueBuilder) String(v string) *FieldValueBuilder {
	u.v = v
	return u
}
