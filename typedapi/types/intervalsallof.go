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

// IntervalsAllOf type.
type IntervalsAllOf struct {
	Filter    *IntervalsFilter     `json:"filter,omitempty"`
	Intervals []IntervalsContainer `json:"intervals"`
	MaxGaps   *int                 `json:"max_gaps,omitempty"`
	Ordered   *bool                `json:"ordered,omitempty"`
}

// IntervalsAllOfBuilder holds IntervalsAllOf struct and provides a builder API.
type IntervalsAllOfBuilder struct {
	v *IntervalsAllOf
}

// NewIntervalsAllOf provides a builder for the IntervalsAllOf struct.
func NewIntervalsAllOf() *IntervalsAllOfBuilder {
	r := IntervalsAllOfBuilder{
		&IntervalsAllOf{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsAllOf struct
func (rb *IntervalsAllOfBuilder) Build() IntervalsAllOf {
	return *rb.v
}

// Filter set the Filter property for IntervalsAllOfBuilder.
func (rb *IntervalsAllOfBuilder) Filter(filter IntervalsFilter) *IntervalsAllOfBuilder {
	rb.v.Filter = &filter
	return rb
}

// Intervals set the Intervals property for IntervalsAllOfBuilder.
func (rb *IntervalsAllOfBuilder) Intervals(intervals ...IntervalsContainer) *IntervalsAllOfBuilder {
	rb.v.Intervals = append(rb.v.Intervals, intervals...)
	return rb
}

// MaxGaps set the MaxGaps property for IntervalsAllOfBuilder.
func (rb *IntervalsAllOfBuilder) MaxGaps(maxgaps int) *IntervalsAllOfBuilder {
	rb.v.MaxGaps = &maxgaps
	return rb
}

// Ordered set the Ordered property for IntervalsAllOfBuilder.
func (rb *IntervalsAllOfBuilder) Ordered(ordered bool) *IntervalsAllOfBuilder {
	rb.v.Ordered = &ordered
	return rb
}
