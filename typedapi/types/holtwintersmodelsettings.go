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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/holtwinterstype"
)

// HoltWintersModelSettings type.
type HoltWintersModelSettings struct {
	Alpha  *float32                         `json:"alpha,omitempty"`
	Beta   *float32                         `json:"beta,omitempty"`
	Gamma  *float32                         `json:"gamma,omitempty"`
	Pad    *bool                            `json:"pad,omitempty"`
	Period *int                             `json:"period,omitempty"`
	Type   *holtwinterstype.HoltWintersType `json:"type,omitempty"`
}

// HoltWintersModelSettingsBuilder holds HoltWintersModelSettings struct and provides a builder API.
type HoltWintersModelSettingsBuilder struct {
	v *HoltWintersModelSettings
}

// NewHoltWintersModelSettings provides a builder for the HoltWintersModelSettings struct.
func NewHoltWintersModelSettings() *HoltWintersModelSettingsBuilder {
	r := HoltWintersModelSettingsBuilder{
		&HoltWintersModelSettings{},
	}

	return &r
}

// Build finalize the chain and returns the HoltWintersModelSettings struct
func (rb *HoltWintersModelSettingsBuilder) Build() HoltWintersModelSettings {
	return *rb.v
}

// Alpha set the Alpha property for HoltWintersModelSettingsBuilder.
func (rb *HoltWintersModelSettingsBuilder) Alpha(alpha float32) *HoltWintersModelSettingsBuilder {
	rb.v.Alpha = &alpha
	return rb
}

// Beta set the Beta property for HoltWintersModelSettingsBuilder.
func (rb *HoltWintersModelSettingsBuilder) Beta(beta float32) *HoltWintersModelSettingsBuilder {
	rb.v.Beta = &beta
	return rb
}

// Gamma set the Gamma property for HoltWintersModelSettingsBuilder.
func (rb *HoltWintersModelSettingsBuilder) Gamma(gamma float32) *HoltWintersModelSettingsBuilder {
	rb.v.Gamma = &gamma
	return rb
}

// Pad set the Pad property for HoltWintersModelSettingsBuilder.
func (rb *HoltWintersModelSettingsBuilder) Pad(pad bool) *HoltWintersModelSettingsBuilder {
	rb.v.Pad = &pad
	return rb
}

// Period set the Period property for HoltWintersModelSettingsBuilder.
func (rb *HoltWintersModelSettingsBuilder) Period(period int) *HoltWintersModelSettingsBuilder {
	rb.v.Period = &period
	return rb
}

// Type set the Type property for HoltWintersModelSettingsBuilder.
func (rb *HoltWintersModelSettingsBuilder) Type_(type_ holtwinterstype.HoltWintersType) *HoltWintersModelSettingsBuilder {
	rb.v.Type = &type_
	return rb
}
