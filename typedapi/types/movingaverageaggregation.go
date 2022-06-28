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

// MovingAverageAggregation holds the union for the following types:
//     EwmaMovingAverageAggregation
//     HoltMovingAverageAggregation
//     HoltWintersMovingAverageAggregation
//     LinearMovingAverageAggregation
//     SimpleMovingAverageAggregation
type MovingAverageAggregation interface{}

// MovingAverageAggregationBuilder holds MovingAverageAggregation struct and provides a builder API.
type MovingAverageAggregationBuilder struct {
	v MovingAverageAggregation
}

// NewMovingAverageAggregation provides a builder for the MovingAverageAggregation struct.
func NewMovingAverageAggregation() *MovingAverageAggregationBuilder {
	return &MovingAverageAggregationBuilder{}
}

// Build finalize the chain and returns the MovingAverageAggregation struct
func (u *MovingAverageAggregationBuilder) Build() MovingAverageAggregation {
	return u.v
}

// EwmaMovingAverageAggregation set the EwmaMovingAverageAggregation property for MovingAverageAggregationBuilder.
func (u *MovingAverageAggregationBuilder) EwmaMovingAverageAggregation(v EwmaMovingAverageAggregation) *MovingAverageAggregationBuilder {
	u.v = v
	return u
}

// HoltMovingAverageAggregation set the HoltMovingAverageAggregation property for MovingAverageAggregationBuilder.
func (u *MovingAverageAggregationBuilder) HoltMovingAverageAggregation(v HoltMovingAverageAggregation) *MovingAverageAggregationBuilder {
	u.v = v
	return u
}

// HoltWintersMovingAverageAggregation set the HoltWintersMovingAverageAggregation property for MovingAverageAggregationBuilder.
func (u *MovingAverageAggregationBuilder) HoltWintersMovingAverageAggregation(v HoltWintersMovingAverageAggregation) *MovingAverageAggregationBuilder {
	u.v = v
	return u
}

// LinearMovingAverageAggregation set the LinearMovingAverageAggregation property for MovingAverageAggregationBuilder.
func (u *MovingAverageAggregationBuilder) LinearMovingAverageAggregation(v LinearMovingAverageAggregation) *MovingAverageAggregationBuilder {
	u.v = v
	return u
}

// SimpleMovingAverageAggregation set the SimpleMovingAverageAggregation property for MovingAverageAggregationBuilder.
func (u *MovingAverageAggregationBuilder) SimpleMovingAverageAggregation(v SimpleMovingAverageAggregation) *MovingAverageAggregationBuilder {
	u.v = v
	return u
}
