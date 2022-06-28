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

// NativeCode type.
type NativeCode struct {
	BuildHash string        `json:"build_hash"`
	Version   VersionString `json:"version"`
}

// NativeCodeBuilder holds NativeCode struct and provides a builder API.
type NativeCodeBuilder struct {
	v *NativeCode
}

// NewNativeCode provides a builder for the NativeCode struct.
func NewNativeCode() *NativeCodeBuilder {
	r := NativeCodeBuilder{
		&NativeCode{},
	}

	return &r
}

// Build finalize the chain and returns the NativeCode struct
func (rb *NativeCodeBuilder) Build() NativeCode {
	return *rb.v
}

// BuildHash set the BuildHash property for NativeCodeBuilder.
func (rb *NativeCodeBuilder) BuildHash(buildhash string) *NativeCodeBuilder {
	rb.v.BuildHash = buildhash
	return rb
}

// Version set the Version property for NativeCodeBuilder.
func (rb *NativeCodeBuilder) Version(version VersionString) *NativeCodeBuilder {
	rb.v.Version = version
	return rb
}
