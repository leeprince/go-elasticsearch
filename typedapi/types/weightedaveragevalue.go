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

// WeightedAverageValue type.
type WeightedAverageValue struct {
	Field   *Field   `json:"field,omitempty"`
	Missing *float64 `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// WeightedAverageValueBuilder holds WeightedAverageValue struct and provides a builder API.
type WeightedAverageValueBuilder struct {
	v *WeightedAverageValue
}

// NewWeightedAverageValue provides a builder for the WeightedAverageValue struct.
func NewWeightedAverageValue() *WeightedAverageValueBuilder {
	r := WeightedAverageValueBuilder{
		&WeightedAverageValue{},
	}

	return &r
}

// Build finalize the chain and returns the WeightedAverageValue struct
func (rb *WeightedAverageValueBuilder) Build() WeightedAverageValue {
	return *rb.v
}

// Field set the Field property for WeightedAverageValueBuilder.
func (rb *WeightedAverageValueBuilder) Field(field Field) *WeightedAverageValueBuilder {
	rb.v.Field = &field
	return rb
}

// Missing set the Missing property for WeightedAverageValueBuilder.
func (rb *WeightedAverageValueBuilder) Missing(missing float64) *WeightedAverageValueBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for WeightedAverageValueBuilder.
func (rb *WeightedAverageValueBuilder) Script(script Script) *WeightedAverageValueBuilder {
	rb.v.Script = &script
	return rb
}
