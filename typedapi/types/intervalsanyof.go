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

// IntervalsAnyOf type.
type IntervalsAnyOf struct {
	Filter    *IntervalsFilter     `json:"filter,omitempty"`
	Intervals []IntervalsContainer `json:"intervals"`
}

// IntervalsAnyOfBuilder holds IntervalsAnyOf struct and provides a builder API.
type IntervalsAnyOfBuilder struct {
	v *IntervalsAnyOf
}

// NewIntervalsAnyOf provides a builder for the IntervalsAnyOf struct.
func NewIntervalsAnyOf() *IntervalsAnyOfBuilder {
	r := IntervalsAnyOfBuilder{
		&IntervalsAnyOf{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsAnyOf struct
func (rb *IntervalsAnyOfBuilder) Build() IntervalsAnyOf {
	return *rb.v
}

// Filter set the Filter property for IntervalsAnyOfBuilder.
func (rb *IntervalsAnyOfBuilder) Filter(filter IntervalsFilter) *IntervalsAnyOfBuilder {
	rb.v.Filter = &filter
	return rb
}

// Intervals set the Intervals property for IntervalsAnyOfBuilder.
func (rb *IntervalsAnyOfBuilder) Intervals(intervals ...IntervalsContainer) *IntervalsAnyOfBuilder {
	rb.v.Intervals = append(rb.v.Intervals, intervals...)
	return rb
}
