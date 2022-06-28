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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/keeptypesmode"
)

// KeepTypesTokenFilter type.
type KeepTypesTokenFilter struct {
	Mode    *keeptypesmode.KeepTypesMode `json:"mode,omitempty"`
	Type    string                       `json:"type,omitempty"`
	Types   []string                     `json:"types,omitempty"`
	Version *VersionString               `json:"version,omitempty"`
}

// KeepTypesTokenFilterBuilder holds KeepTypesTokenFilter struct and provides a builder API.
type KeepTypesTokenFilterBuilder struct {
	v *KeepTypesTokenFilter
}

// NewKeepTypesTokenFilter provides a builder for the KeepTypesTokenFilter struct.
func NewKeepTypesTokenFilter() *KeepTypesTokenFilterBuilder {
	r := KeepTypesTokenFilterBuilder{
		&KeepTypesTokenFilter{},
	}

	r.v.Type = "keep_types"

	return &r
}

// Build finalize the chain and returns the KeepTypesTokenFilter struct
func (rb *KeepTypesTokenFilterBuilder) Build() KeepTypesTokenFilter {
	return *rb.v
}

// Mode set the Mode property for KeepTypesTokenFilterBuilder.
func (rb *KeepTypesTokenFilterBuilder) Mode(mode keeptypesmode.KeepTypesMode) *KeepTypesTokenFilterBuilder {
	rb.v.Mode = &mode
	return rb
}

// Type set the Type property for KeepTypesTokenFilterBuilder.

// Types set the Types property for KeepTypesTokenFilterBuilder.
func (rb *KeepTypesTokenFilterBuilder) Types(types ...string) *KeepTypesTokenFilterBuilder {
	rb.v.Types = append(rb.v.Types, types...)
	return rb
}

// Version set the Version property for KeepTypesTokenFilterBuilder.
func (rb *KeepTypesTokenFilterBuilder) Version(version VersionString) *KeepTypesTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
