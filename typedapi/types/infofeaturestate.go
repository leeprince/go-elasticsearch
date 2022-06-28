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

// InfoFeatureState type.
type InfoFeatureState struct {
	FeatureName string  `json:"feature_name"`
	Indices     Indices `json:"indices"`
}

// InfoFeatureStateBuilder holds InfoFeatureState struct and provides a builder API.
type InfoFeatureStateBuilder struct {
	v *InfoFeatureState
}

// NewInfoFeatureState provides a builder for the InfoFeatureState struct.
func NewInfoFeatureState() *InfoFeatureStateBuilder {
	r := InfoFeatureStateBuilder{
		&InfoFeatureState{},
	}

	return &r
}

// Build finalize the chain and returns the InfoFeatureState struct
func (rb *InfoFeatureStateBuilder) Build() InfoFeatureState {
	return *rb.v
}

// FeatureName set the FeatureName property for InfoFeatureStateBuilder.
func (rb *InfoFeatureStateBuilder) FeatureName(featurename string) *InfoFeatureStateBuilder {
	rb.v.FeatureName = featurename
	return rb
}

// Indices set the Indices property for InfoFeatureStateBuilder.
func (rb *InfoFeatureStateBuilder) Indices(indices Indices) *InfoFeatureStateBuilder {
	rb.v.Indices = indices
	return rb
}
