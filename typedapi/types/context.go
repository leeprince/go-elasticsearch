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

// Context holds the union for the following types:
//     GeoLocation
//     string
type Context interface{}

// ContextBuilder holds Context struct and provides a builder API.
type ContextBuilder struct {
	v Context
}

// NewContext provides a builder for the Context struct.
func NewContext() *ContextBuilder {
	return &ContextBuilder{}
}

// Build finalize the chain and returns the Context struct
func (u *ContextBuilder) Build() Context {
	return u.v
}

// GeoLocation set the GeoLocation property for ContextBuilder.
func (u *ContextBuilder) GeoLocation(v GeoLocation) *ContextBuilder {
	u.v = v
	return u
}

// String set the String property for ContextBuilder.
func (u *ContextBuilder) String(v string) *ContextBuilder {
	u.v = v
	return u
}
