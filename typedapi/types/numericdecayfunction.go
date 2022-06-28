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

// NumericDecayFunction type.
type NumericDecayFunction struct {
	MultiValueMode       *multivaluemode.MultiValueMode       `json:"multi_value_mode,omitempty"`
	NumericDecayFunction map[Field]DecayPlacementdoubledouble `json:"NumericDecayFunction,omitempty"`
}

// NumericDecayFunctionBuilder holds NumericDecayFunction struct and provides a builder API.
type NumericDecayFunctionBuilder struct {
	v *NumericDecayFunction
}

// NewNumericDecayFunction provides a builder for the NumericDecayFunction struct.
func NewNumericDecayFunction() *NumericDecayFunctionBuilder {
	r := NumericDecayFunctionBuilder{
		&NumericDecayFunction{
			NumericDecayFunction: make(map[Field]DecayPlacementdoubledouble, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NumericDecayFunction struct
func (rb *NumericDecayFunctionBuilder) Build() NumericDecayFunction {
	return *rb.v
}

// MultiValueMode set the MultiValueMode property for NumericDecayFunctionBuilder.
func (rb *NumericDecayFunctionBuilder) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *NumericDecayFunctionBuilder {
	rb.v.MultiValueMode = &multivaluemode
	return rb
}

// NumericDecayFunction set the NumericDecayFunction property for NumericDecayFunctionBuilder.
func (rb *NumericDecayFunctionBuilder) NumericDecayFunction(value map[Field]DecayPlacementdoubledouble) *NumericDecayFunctionBuilder {
	rb.v.NumericDecayFunction = value
	return rb
}
