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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/simplequerystringflag"
)

// SimpleQueryStringFlags holds the union for the following types:
//     simplequerystringflag.SimpleQueryStringFlag
//     string
type SimpleQueryStringFlags interface{}

// SimpleQueryStringFlagsBuilder holds SimpleQueryStringFlags struct and provides a builder API.
type SimpleQueryStringFlagsBuilder struct {
	v SimpleQueryStringFlags
}

// NewSimpleQueryStringFlags provides a builder for the SimpleQueryStringFlags struct.
func NewSimpleQueryStringFlags() *SimpleQueryStringFlagsBuilder {
	return &SimpleQueryStringFlagsBuilder{}
}

// Build finalize the chain and returns the SimpleQueryStringFlags struct
func (u *SimpleQueryStringFlagsBuilder) Build() SimpleQueryStringFlags {
	return u.v
}

// SimpleQueryStringFlag set the SimpleQueryStringFlag property for SimpleQueryStringFlagsBuilder.
func (u *SimpleQueryStringFlagsBuilder) SimpleQueryStringFlag(v simplequerystringflag.SimpleQueryStringFlag) *SimpleQueryStringFlagsBuilder {
	u.v = v
	return u
}

// String set the String property for SimpleQueryStringFlagsBuilder.
func (u *SimpleQueryStringFlagsBuilder) String(v string) *SimpleQueryStringFlagsBuilder {
	u.v = v
	return u
}
