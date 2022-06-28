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

// TermsInclude holds the union for the following types:
//     string
//     []string
//     TermsPartition
type TermsInclude interface{}

// TermsIncludeBuilder holds TermsInclude struct and provides a builder API.
type TermsIncludeBuilder struct {
	v TermsInclude
}

// NewTermsInclude provides a builder for the TermsInclude struct.
func NewTermsInclude() *TermsIncludeBuilder {
	return &TermsIncludeBuilder{}
}

// Build finalize the chain and returns the TermsInclude struct
func (u *TermsIncludeBuilder) Build() TermsInclude {
	return u.v
}

// String set the String property for TermsIncludeBuilder.
func (u *TermsIncludeBuilder) String(v string) *TermsIncludeBuilder {
	u.v = v
	return u
}

// Strings set the Strings property for TermsIncludeBuilder.
func (u *TermsIncludeBuilder) Strings(v []string) *TermsIncludeBuilder {
	u.v = v
	return u
}

// TermsPartition set the TermsPartition property for TermsIncludeBuilder.
func (u *TermsIncludeBuilder) TermsPartition(v TermsPartition) *TermsIncludeBuilder {
	u.v = v
	return u
}
