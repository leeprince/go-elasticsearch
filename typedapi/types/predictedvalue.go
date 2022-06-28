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

// PredictedValue holds the union for the following types:
//     bool
//     float64
//     int
//     string
type PredictedValue interface{}

// PredictedValueBuilder holds PredictedValue struct and provides a builder API.
type PredictedValueBuilder struct {
	v PredictedValue
}

// NewPredictedValue provides a builder for the PredictedValue struct.
func NewPredictedValue() *PredictedValueBuilder {
	return &PredictedValueBuilder{}
}

// Build finalize the chain and returns the PredictedValue struct
func (u *PredictedValueBuilder) Build() PredictedValue {
	return u.v
}

// Bool set the Bool property for PredictedValueBuilder.
func (u *PredictedValueBuilder) Bool(v bool) *PredictedValueBuilder {
	u.v = v
	return u
}

// Float64 set the Float64 property for PredictedValueBuilder.
func (u *PredictedValueBuilder) Float64(v float64) *PredictedValueBuilder {
	u.v = v
	return u
}

// Int set the Int property for PredictedValueBuilder.
func (u *PredictedValueBuilder) Int(v int) *PredictedValueBuilder {
	u.v = v
	return u
}

// String set the String property for PredictedValueBuilder.
func (u *PredictedValueBuilder) String(v string) *PredictedValueBuilder {
	u.v = v
	return u
}
