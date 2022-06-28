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

// EwmaModelSettings type.
type EwmaModelSettings struct {
	Alpha *float32 `json:"alpha,omitempty"`
}

// EwmaModelSettingsBuilder holds EwmaModelSettings struct and provides a builder API.
type EwmaModelSettingsBuilder struct {
	v *EwmaModelSettings
}

// NewEwmaModelSettings provides a builder for the EwmaModelSettings struct.
func NewEwmaModelSettings() *EwmaModelSettingsBuilder {
	r := EwmaModelSettingsBuilder{
		&EwmaModelSettings{},
	}

	return &r
}

// Build finalize the chain and returns the EwmaModelSettings struct
func (rb *EwmaModelSettingsBuilder) Build() EwmaModelSettings {
	return *rb.v
}

// Alpha set the Alpha property for EwmaModelSettingsBuilder.
func (rb *EwmaModelSettingsBuilder) Alpha(alpha float32) *EwmaModelSettingsBuilder {
	rb.v.Alpha = &alpha
	return rb
}
