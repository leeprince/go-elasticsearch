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

// ByteSize holds the union for the following types:
//     int64
//     string
type ByteSize interface{}

// ByteSizeBuilder holds ByteSize struct and provides a builder API.
type ByteSizeBuilder struct {
	v ByteSize
}

// NewByteSize provides a builder for the ByteSize struct.
func NewByteSize() *ByteSizeBuilder {
	return &ByteSizeBuilder{}
}

// Build finalize the chain and returns the ByteSize struct
func (u *ByteSizeBuilder) Build() ByteSize {
	return u.v
}

// Int64 set the Int64 property for ByteSizeBuilder.
func (u *ByteSizeBuilder) Int64(v int64) *ByteSizeBuilder {
	u.v = v
	return u
}

// String set the String property for ByteSizeBuilder.
func (u *ByteSizeBuilder) String(v string) *ByteSizeBuilder {
	u.v = v
	return u
}
