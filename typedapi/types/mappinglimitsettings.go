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

// MappingLimitSettings type.
type MappingLimitSettings struct {
	Depth           *MappingLimitSettingsDepth           `json:"depth,omitempty"`
	DimensionFields *MappingLimitSettingsDimensionFields `json:"dimension_fields,omitempty"`
	FieldNameLength *MappingLimitSettingsFieldNameLength `json:"field_name_length,omitempty"`
	IgnoreMalformed *bool                                `json:"ignore_malformed,omitempty"`
	NestedFields    *MappingLimitSettingsNestedFields    `json:"nested_fields,omitempty"`
	NestedObjects   *MappingLimitSettingsNestedObjects   `json:"nested_objects,omitempty"`
	TotalFields     *MappingLimitSettingsTotalFields     `json:"total_fields,omitempty"`
}

// MappingLimitSettingsBuilder holds MappingLimitSettings struct and provides a builder API.
type MappingLimitSettingsBuilder struct {
	v *MappingLimitSettings
}

// NewMappingLimitSettings provides a builder for the MappingLimitSettings struct.
func NewMappingLimitSettings() *MappingLimitSettingsBuilder {
	r := MappingLimitSettingsBuilder{
		&MappingLimitSettings{},
	}

	return &r
}

// Build finalize the chain and returns the MappingLimitSettings struct
func (rb *MappingLimitSettingsBuilder) Build() MappingLimitSettings {
	return *rb.v
}

// Depth set the Depth property for MappingLimitSettingsBuilder.
func (rb *MappingLimitSettingsBuilder) Depth(depth MappingLimitSettingsDepth) *MappingLimitSettingsBuilder {
	rb.v.Depth = &depth
	return rb
}

// DimensionFields set the DimensionFields property for MappingLimitSettingsBuilder.
func (rb *MappingLimitSettingsBuilder) DimensionFields(dimensionfields MappingLimitSettingsDimensionFields) *MappingLimitSettingsBuilder {
	rb.v.DimensionFields = &dimensionfields
	return rb
}

// FieldNameLength set the FieldNameLength property for MappingLimitSettingsBuilder.
func (rb *MappingLimitSettingsBuilder) FieldNameLength(fieldnamelength MappingLimitSettingsFieldNameLength) *MappingLimitSettingsBuilder {
	rb.v.FieldNameLength = &fieldnamelength
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for MappingLimitSettingsBuilder.
func (rb *MappingLimitSettingsBuilder) IgnoreMalformed(ignoremalformed bool) *MappingLimitSettingsBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// NestedFields set the NestedFields property for MappingLimitSettingsBuilder.
func (rb *MappingLimitSettingsBuilder) NestedFields(nestedfields MappingLimitSettingsNestedFields) *MappingLimitSettingsBuilder {
	rb.v.NestedFields = &nestedfields
	return rb
}

// NestedObjects set the NestedObjects property for MappingLimitSettingsBuilder.
func (rb *MappingLimitSettingsBuilder) NestedObjects(nestedobjects MappingLimitSettingsNestedObjects) *MappingLimitSettingsBuilder {
	rb.v.NestedObjects = &nestedobjects
	return rb
}

// TotalFields set the TotalFields property for MappingLimitSettingsBuilder.
func (rb *MappingLimitSettingsBuilder) TotalFields(totalfields MappingLimitSettingsTotalFields) *MappingLimitSettingsBuilder {
	rb.v.TotalFields = &totalfields
	return rb
}
