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

// JvmClasses type.
type JvmClasses struct {
	CurrentLoadedCount *int64 `json:"current_loaded_count,omitempty"`
	TotalLoadedCount   *int64 `json:"total_loaded_count,omitempty"`
	TotalUnloadedCount *int64 `json:"total_unloaded_count,omitempty"`
}

// JvmClassesBuilder holds JvmClasses struct and provides a builder API.
type JvmClassesBuilder struct {
	v *JvmClasses
}

// NewJvmClasses provides a builder for the JvmClasses struct.
func NewJvmClasses() *JvmClassesBuilder {
	r := JvmClassesBuilder{
		&JvmClasses{},
	}

	return &r
}

// Build finalize the chain and returns the JvmClasses struct
func (rb *JvmClassesBuilder) Build() JvmClasses {
	return *rb.v
}

// CurrentLoadedCount set the CurrentLoadedCount property for JvmClassesBuilder.
func (rb *JvmClassesBuilder) CurrentLoadedCount(currentloadedcount int64) *JvmClassesBuilder {
	rb.v.CurrentLoadedCount = &currentloadedcount
	return rb
}

// TotalLoadedCount set the TotalLoadedCount property for JvmClassesBuilder.
func (rb *JvmClassesBuilder) TotalLoadedCount(totalloadedcount int64) *JvmClassesBuilder {
	rb.v.TotalLoadedCount = &totalloadedcount
	return rb
}

// TotalUnloadedCount set the TotalUnloadedCount property for JvmClassesBuilder.
func (rb *JvmClassesBuilder) TotalUnloadedCount(totalunloadedcount int64) *JvmClassesBuilder {
	rb.v.TotalUnloadedCount = &totalunloadedcount
	return rb
}
