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

// DistanceFeatureQueryBaseDateMathTime type.
type DistanceFeatureQueryBaseDateMathTime struct {
	Boost      *float32 `json:"boost,omitempty"`
	Field      Field    `json:"field"`
	Origin     DateMath `json:"origin"`
	Pivot      Time     `json:"pivot"`
	QueryName_ *string  `json:"_name,omitempty"`
}

// DistanceFeatureQueryBaseDateMathTimeBuilder holds DistanceFeatureQueryBaseDateMathTime struct and provides a builder API.
type DistanceFeatureQueryBaseDateMathTimeBuilder struct {
	v *DistanceFeatureQueryBaseDateMathTime
}

// NewDistanceFeatureQueryBaseDateMathTime provides a builder for the DistanceFeatureQueryBaseDateMathTime struct.
func NewDistanceFeatureQueryBaseDateMathTime() *DistanceFeatureQueryBaseDateMathTimeBuilder {
	r := DistanceFeatureQueryBaseDateMathTimeBuilder{
		&DistanceFeatureQueryBaseDateMathTime{},
	}

	return &r
}

// Build finalize the chain and returns the DistanceFeatureQueryBaseDateMathTime struct
func (rb *DistanceFeatureQueryBaseDateMathTimeBuilder) Build() DistanceFeatureQueryBaseDateMathTime {
	return *rb.v
}

// Boost set the Boost property for DistanceFeatureQueryBaseDateMathTimeBuilder.
func (rb *DistanceFeatureQueryBaseDateMathTimeBuilder) Boost(boost float32) *DistanceFeatureQueryBaseDateMathTimeBuilder {
	rb.v.Boost = &boost
	return rb
}

// Field set the Field property for DistanceFeatureQueryBaseDateMathTimeBuilder.
func (rb *DistanceFeatureQueryBaseDateMathTimeBuilder) Field(field Field) *DistanceFeatureQueryBaseDateMathTimeBuilder {
	rb.v.Field = field
	return rb
}

// Origin set the Origin property for DistanceFeatureQueryBaseDateMathTimeBuilder.
func (rb *DistanceFeatureQueryBaseDateMathTimeBuilder) Origin(origin DateMath) *DistanceFeatureQueryBaseDateMathTimeBuilder {
	rb.v.Origin = origin
	return rb
}

// Pivot set the Pivot property for DistanceFeatureQueryBaseDateMathTimeBuilder.
func (rb *DistanceFeatureQueryBaseDateMathTimeBuilder) Pivot(pivot Time) *DistanceFeatureQueryBaseDateMathTimeBuilder {
	rb.v.Pivot = pivot
	return rb
}

// QueryName_ set the QueryName_ property for DistanceFeatureQueryBaseDateMathTimeBuilder.
func (rb *DistanceFeatureQueryBaseDateMathTimeBuilder) QueryName_(queryname_ string) *DistanceFeatureQueryBaseDateMathTimeBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
