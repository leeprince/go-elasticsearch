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

// ConditionTokenFilter type.
type ConditionTokenFilter struct {
	Filter  []string       `json:"filter"`
	Script  Script         `json:"script"`
	Type    string         `json:"type,omitempty"`
	Version *VersionString `json:"version,omitempty"`
}

// ConditionTokenFilterBuilder holds ConditionTokenFilter struct and provides a builder API.
type ConditionTokenFilterBuilder struct {
	v *ConditionTokenFilter
}

// NewConditionTokenFilter provides a builder for the ConditionTokenFilter struct.
func NewConditionTokenFilter() *ConditionTokenFilterBuilder {
	r := ConditionTokenFilterBuilder{
		&ConditionTokenFilter{},
	}

	r.v.Type = "condition"

	return &r
}

// Build finalize the chain and returns the ConditionTokenFilter struct
func (rb *ConditionTokenFilterBuilder) Build() ConditionTokenFilter {
	return *rb.v
}

// Filter set the Filter property for ConditionTokenFilterBuilder.
func (rb *ConditionTokenFilterBuilder) Filter(filter ...string) *ConditionTokenFilterBuilder {
	rb.v.Filter = append(rb.v.Filter, filter...)
	return rb
}

// Script set the Script property for ConditionTokenFilterBuilder.
func (rb *ConditionTokenFilterBuilder) Script(script Script) *ConditionTokenFilterBuilder {
	rb.v.Script = script
	return rb
}

// Type set the Type property for ConditionTokenFilterBuilder.

// Version set the Version property for ConditionTokenFilterBuilder.
func (rb *ConditionTokenFilterBuilder) Version(version VersionString) *ConditionTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
