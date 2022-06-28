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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
)

// GeoDecayFunction type.
type GeoDecayFunction struct {
	GeoDecayFunction map[Field]DecayPlacementGeoLocationDistance `json:"GeoDecayFunction,omitempty"`
	MultiValueMode   *multivaluemode.MultiValueMode              `json:"multi_value_mode,omitempty"`
}

// GeoDecayFunctionBuilder holds GeoDecayFunction struct and provides a builder API.
type GeoDecayFunctionBuilder struct {
	v *GeoDecayFunction
}

// NewGeoDecayFunction provides a builder for the GeoDecayFunction struct.
func NewGeoDecayFunction() *GeoDecayFunctionBuilder {
	r := GeoDecayFunctionBuilder{
		&GeoDecayFunction{
			GeoDecayFunction: make(map[Field]DecayPlacementGeoLocationDistance, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GeoDecayFunction struct
func (rb *GeoDecayFunctionBuilder) Build() GeoDecayFunction {
	return *rb.v
}

// GeoDecayFunction set the GeoDecayFunction property for GeoDecayFunctionBuilder.
func (rb *GeoDecayFunctionBuilder) GeoDecayFunction(value map[Field]DecayPlacementGeoLocationDistance) *GeoDecayFunctionBuilder {
	rb.v.GeoDecayFunction = value
	return rb
}

// MultiValueMode set the MultiValueMode property for GeoDecayFunctionBuilder.
func (rb *GeoDecayFunctionBuilder) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *GeoDecayFunctionBuilder {
	rb.v.MultiValueMode = &multivaluemode
	return rb
}
