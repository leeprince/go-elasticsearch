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

// Feature type.
type Feature struct {
	Available      bool                   `json:"available"`
	Description    *string                `json:"description,omitempty"`
	Enabled        bool                   `json:"enabled"`
	NativeCodeInfo *NativeCodeInformation `json:"native_code_info,omitempty"`
}

// FeatureBuilder holds Feature struct and provides a builder API.
type FeatureBuilder struct {
	v *Feature
}

// NewFeature provides a builder for the Feature struct.
func NewFeature() *FeatureBuilder {
	r := FeatureBuilder{
		&Feature{},
	}

	return &r
}

// Build finalize the chain and returns the Feature struct
func (rb *FeatureBuilder) Build() Feature {
	return *rb.v
}

// Available set the Available property for FeatureBuilder.
func (rb *FeatureBuilder) Available(available bool) *FeatureBuilder {
	rb.v.Available = available
	return rb
}

// Description set the Description property for FeatureBuilder.
func (rb *FeatureBuilder) Description(description string) *FeatureBuilder {
	rb.v.Description = &description
	return rb
}

// Enabled set the Enabled property for FeatureBuilder.
func (rb *FeatureBuilder) Enabled(enabled bool) *FeatureBuilder {
	rb.v.Enabled = enabled
	return rb
}

// NativeCodeInfo set the NativeCodeInfo property for FeatureBuilder.
func (rb *FeatureBuilder) NativeCodeInfo(nativecodeinfo NativeCodeInformation) *FeatureBuilder {
	rb.v.NativeCodeInfo = &nativecodeinfo
	return rb
}
