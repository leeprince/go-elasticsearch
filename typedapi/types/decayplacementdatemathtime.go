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

// DecayPlacementDateMathTime type.
type DecayPlacementDateMathTime struct {
	Decay  *float64  `json:"decay,omitempty"`
	Offset *Time     `json:"offset,omitempty"`
	Origin *DateMath `json:"origin,omitempty"`
	Scale  *Time     `json:"scale,omitempty"`
}

// DecayPlacementDateMathTimeBuilder holds DecayPlacementDateMathTime struct and provides a builder API.
type DecayPlacementDateMathTimeBuilder struct {
	v *DecayPlacementDateMathTime
}

// NewDecayPlacementDateMathTime provides a builder for the DecayPlacementDateMathTime struct.
func NewDecayPlacementDateMathTime() *DecayPlacementDateMathTimeBuilder {
	r := DecayPlacementDateMathTimeBuilder{
		&DecayPlacementDateMathTime{},
	}

	return &r
}

// Build finalize the chain and returns the DecayPlacementDateMathTime struct
func (rb *DecayPlacementDateMathTimeBuilder) Build() DecayPlacementDateMathTime {
	return *rb.v
}

// Decay set the Decay property for DecayPlacementDateMathTimeBuilder.
func (rb *DecayPlacementDateMathTimeBuilder) Decay(decay float64) *DecayPlacementDateMathTimeBuilder {
	rb.v.Decay = &decay
	return rb
}

// Offset set the Offset property for DecayPlacementDateMathTimeBuilder.
func (rb *DecayPlacementDateMathTimeBuilder) Offset(offset Time) *DecayPlacementDateMathTimeBuilder {
	rb.v.Offset = &offset
	return rb
}

// Origin set the Origin property for DecayPlacementDateMathTimeBuilder.
func (rb *DecayPlacementDateMathTimeBuilder) Origin(origin DateMath) *DecayPlacementDateMathTimeBuilder {
	rb.v.Origin = &origin
	return rb
}

// Scale set the Scale property for DecayPlacementDateMathTimeBuilder.
func (rb *DecayPlacementDateMathTimeBuilder) Scale(scale Time) *DecayPlacementDateMathTimeBuilder {
	rb.v.Scale = &scale
	return rb
}
