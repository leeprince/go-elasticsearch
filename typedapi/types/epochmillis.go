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

// EpochMillis holds the union for the following types:
//     int64
//     string
type EpochMillis interface{}

// EpochMillisBuilder holds EpochMillis struct and provides a builder API.
type EpochMillisBuilder struct {
	v EpochMillis
}

// NewEpochMillis provides a builder for the EpochMillis struct.
func NewEpochMillis() *EpochMillisBuilder {
	return &EpochMillisBuilder{}
}

// Build finalize the chain and returns the EpochMillis struct
func (u *EpochMillisBuilder) Build() EpochMillis {
	return u.v
}

// Int64 set the Int64 property for EpochMillisBuilder.
func (u *EpochMillisBuilder) Int64(v int64) *EpochMillisBuilder {
	u.v = v
	return u
}

// String set the String property for EpochMillisBuilder.
func (u *EpochMillisBuilder) String(v string) *EpochMillisBuilder {
	u.v = v
	return u
}
