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

// SmoothingModelContainer type.
type SmoothingModelContainer struct {
	Laplace             *LaplaceSmoothingModel             `json:"laplace,omitempty"`
	LinearInterpolation *LinearInterpolationSmoothingModel `json:"linear_interpolation,omitempty"`
	StupidBackoff       *StupidBackoffSmoothingModel       `json:"stupid_backoff,omitempty"`
}

// SmoothingModelContainerBuilder holds SmoothingModelContainer struct and provides a builder API.
type SmoothingModelContainerBuilder struct {
	v *SmoothingModelContainer
}

// NewSmoothingModelContainer provides a builder for the SmoothingModelContainer struct.
func NewSmoothingModelContainer() *SmoothingModelContainerBuilder {
	r := SmoothingModelContainerBuilder{
		&SmoothingModelContainer{},
	}

	return &r
}

// Build finalize the chain and returns the SmoothingModelContainer struct
func (rb *SmoothingModelContainerBuilder) Build() SmoothingModelContainer {
	return *rb.v
}

// Laplace set the Laplace property for SmoothingModelContainerBuilder.
func (rb *SmoothingModelContainerBuilder) Laplace(laplace LaplaceSmoothingModel) *SmoothingModelContainerBuilder {
	rb.v.Laplace = &laplace
	return rb
}

// LinearInterpolation set the LinearInterpolation property for SmoothingModelContainerBuilder.
func (rb *SmoothingModelContainerBuilder) LinearInterpolation(linearinterpolation LinearInterpolationSmoothingModel) *SmoothingModelContainerBuilder {
	rb.v.LinearInterpolation = &linearinterpolation
	return rb
}

// StupidBackoff set the StupidBackoff property for SmoothingModelContainerBuilder.
func (rb *SmoothingModelContainerBuilder) StupidBackoff(stupidbackoff StupidBackoffSmoothingModel) *SmoothingModelContainerBuilder {
	rb.v.StupidBackoff = &stupidbackoff
	return rb
}
