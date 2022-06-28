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

// HoltLinearModelSettings type.
type HoltLinearModelSettings struct {
	Alpha *float32 `json:"alpha,omitempty"`
	Beta  *float32 `json:"beta,omitempty"`
}

// HoltLinearModelSettingsBuilder holds HoltLinearModelSettings struct and provides a builder API.
type HoltLinearModelSettingsBuilder struct {
	v *HoltLinearModelSettings
}

// NewHoltLinearModelSettings provides a builder for the HoltLinearModelSettings struct.
func NewHoltLinearModelSettings() *HoltLinearModelSettingsBuilder {
	r := HoltLinearModelSettingsBuilder{
		&HoltLinearModelSettings{},
	}

	return &r
}

// Build finalize the chain and returns the HoltLinearModelSettings struct
func (rb *HoltLinearModelSettingsBuilder) Build() HoltLinearModelSettings {
	return *rb.v
}

// Alpha set the Alpha property for HoltLinearModelSettingsBuilder.
func (rb *HoltLinearModelSettingsBuilder) Alpha(alpha float32) *HoltLinearModelSettingsBuilder {
	rb.v.Alpha = &alpha
	return rb
}

// Beta set the Beta property for HoltLinearModelSettingsBuilder.
func (rb *HoltLinearModelSettingsBuilder) Beta(beta float32) *HoltLinearModelSettingsBuilder {
	rb.v.Beta = &beta
	return rb
}
