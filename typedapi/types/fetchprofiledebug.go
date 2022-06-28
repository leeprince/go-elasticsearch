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

// FetchProfileDebug type.
type FetchProfileDebug struct {
	FastPath     *int     `json:"fast_path,omitempty"`
	StoredFields []string `json:"stored_fields,omitempty"`
}

// FetchProfileDebugBuilder holds FetchProfileDebug struct and provides a builder API.
type FetchProfileDebugBuilder struct {
	v *FetchProfileDebug
}

// NewFetchProfileDebug provides a builder for the FetchProfileDebug struct.
func NewFetchProfileDebug() *FetchProfileDebugBuilder {
	r := FetchProfileDebugBuilder{
		&FetchProfileDebug{},
	}

	return &r
}

// Build finalize the chain and returns the FetchProfileDebug struct
func (rb *FetchProfileDebugBuilder) Build() FetchProfileDebug {
	return *rb.v
}

// FastPath set the FastPath property for FetchProfileDebugBuilder.
func (rb *FetchProfileDebugBuilder) FastPath(fastpath int) *FetchProfileDebugBuilder {
	rb.v.FastPath = &fastpath
	return rb
}

// StoredFields set the StoredFields property for FetchProfileDebugBuilder.
func (rb *FetchProfileDebugBuilder) StoredFields(stored_fields ...string) *FetchProfileDebugBuilder {
	rb.v.StoredFields = append(rb.v.StoredFields, stored_fields...)
	return rb
}
