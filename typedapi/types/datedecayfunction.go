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

// DateDecayFunction type.
type DateDecayFunction struct {
	DateDecayFunction map[Field]DecayPlacementDateMathTime `json:"DateDecayFunction,omitempty"`
	MultiValueMode    *multivaluemode.MultiValueMode       `json:"multi_value_mode,omitempty"`
}

// DateDecayFunctionBuilder holds DateDecayFunction struct and provides a builder API.
type DateDecayFunctionBuilder struct {
	v *DateDecayFunction
}

// NewDateDecayFunction provides a builder for the DateDecayFunction struct.
func NewDateDecayFunction() *DateDecayFunctionBuilder {
	r := DateDecayFunctionBuilder{
		&DateDecayFunction{
			DateDecayFunction: make(map[Field]DecayPlacementDateMathTime, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DateDecayFunction struct
func (rb *DateDecayFunctionBuilder) Build() DateDecayFunction {
	return *rb.v
}

// DateDecayFunction set the DateDecayFunction property for DateDecayFunctionBuilder.
func (rb *DateDecayFunctionBuilder) DateDecayFunction(value map[Field]DecayPlacementDateMathTime) *DateDecayFunctionBuilder {
	rb.v.DateDecayFunction = value
	return rb
}

// MultiValueMode set the MultiValueMode property for DateDecayFunctionBuilder.
func (rb *DateDecayFunctionBuilder) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *DateDecayFunctionBuilder {
	rb.v.MultiValueMode = &multivaluemode
	return rb
}
