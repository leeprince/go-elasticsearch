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

// StandardDeviationBounds type.
type StandardDeviationBounds struct {
	Lower           float64 `json:"lower,omitempty"`
	LowerPopulation float64 `json:"lower_population,omitempty"`
	LowerSampling   float64 `json:"lower_sampling,omitempty"`
	Upper           float64 `json:"upper,omitempty"`
	UpperPopulation float64 `json:"upper_population,omitempty"`
	UpperSampling   float64 `json:"upper_sampling,omitempty"`
}

// StandardDeviationBoundsBuilder holds StandardDeviationBounds struct and provides a builder API.
type StandardDeviationBoundsBuilder struct {
	v *StandardDeviationBounds
}

// NewStandardDeviationBounds provides a builder for the StandardDeviationBounds struct.
func NewStandardDeviationBounds() *StandardDeviationBoundsBuilder {
	r := StandardDeviationBoundsBuilder{
		&StandardDeviationBounds{},
	}

	return &r
}

// Build finalize the chain and returns the StandardDeviationBounds struct
func (rb *StandardDeviationBoundsBuilder) Build() StandardDeviationBounds {
	return *rb.v
}

// Lower set the Lower property for StandardDeviationBoundsBuilder.
func (rb *StandardDeviationBoundsBuilder) Lower(lower float64) *StandardDeviationBoundsBuilder {
	rb.v.Lower = lower
	return rb
}

// LowerPopulation set the LowerPopulation property for StandardDeviationBoundsBuilder.
func (rb *StandardDeviationBoundsBuilder) LowerPopulation(lowerpopulation float64) *StandardDeviationBoundsBuilder {
	rb.v.LowerPopulation = lowerpopulation
	return rb
}

// LowerSampling set the LowerSampling property for StandardDeviationBoundsBuilder.
func (rb *StandardDeviationBoundsBuilder) LowerSampling(lowersampling float64) *StandardDeviationBoundsBuilder {
	rb.v.LowerSampling = lowersampling
	return rb
}

// Upper set the Upper property for StandardDeviationBoundsBuilder.
func (rb *StandardDeviationBoundsBuilder) Upper(upper float64) *StandardDeviationBoundsBuilder {
	rb.v.Upper = upper
	return rb
}

// UpperPopulation set the UpperPopulation property for StandardDeviationBoundsBuilder.
func (rb *StandardDeviationBoundsBuilder) UpperPopulation(upperpopulation float64) *StandardDeviationBoundsBuilder {
	rb.v.UpperPopulation = upperpopulation
	return rb
}

// UpperSampling set the UpperSampling property for StandardDeviationBoundsBuilder.
func (rb *StandardDeviationBoundsBuilder) UpperSampling(uppersampling float64) *StandardDeviationBoundsBuilder {
	rb.v.UpperSampling = uppersampling
	return rb
}
