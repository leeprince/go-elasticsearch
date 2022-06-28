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

// Missing holds the union for the following types:
//     bool
//     float64
//     int
//     string
type Missing interface{}

// MissingBuilder holds Missing struct and provides a builder API.
type MissingBuilder struct {
	v Missing
}

// NewMissing provides a builder for the Missing struct.
func NewMissing() *MissingBuilder {
	return &MissingBuilder{}
}

// Build finalize the chain and returns the Missing struct
func (u *MissingBuilder) Build() Missing {
	return u.v
}

// Bool set the Bool property for MissingBuilder.
func (u *MissingBuilder) Bool(v bool) *MissingBuilder {
	u.v = v
	return u
}

// Float64 set the Float64 property for MissingBuilder.
func (u *MissingBuilder) Float64(v float64) *MissingBuilder {
	u.v = v
	return u
}

// Int set the Int property for MissingBuilder.
func (u *MissingBuilder) Int(v int) *MissingBuilder {
	u.v = v
	return u
}

// String set the String property for MissingBuilder.
func (u *MissingBuilder) String(v string) *MissingBuilder {
	u.v = v
	return u
}
