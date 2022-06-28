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

// FielddataFrequencyFilter type.
type FielddataFrequencyFilter struct {
	Max            float64 `json:"max"`
	Min            float64 `json:"min"`
	MinSegmentSize int     `json:"min_segment_size"`
}

// FielddataFrequencyFilterBuilder holds FielddataFrequencyFilter struct and provides a builder API.
type FielddataFrequencyFilterBuilder struct {
	v *FielddataFrequencyFilter
}

// NewFielddataFrequencyFilter provides a builder for the FielddataFrequencyFilter struct.
func NewFielddataFrequencyFilter() *FielddataFrequencyFilterBuilder {
	r := FielddataFrequencyFilterBuilder{
		&FielddataFrequencyFilter{},
	}

	return &r
}

// Build finalize the chain and returns the FielddataFrequencyFilter struct
func (rb *FielddataFrequencyFilterBuilder) Build() FielddataFrequencyFilter {
	return *rb.v
}

// Max set the Max property for FielddataFrequencyFilterBuilder.
func (rb *FielddataFrequencyFilterBuilder) Max(max float64) *FielddataFrequencyFilterBuilder {
	rb.v.Max = max
	return rb
}

// Min set the Min property for FielddataFrequencyFilterBuilder.
func (rb *FielddataFrequencyFilterBuilder) Min(min float64) *FielddataFrequencyFilterBuilder {
	rb.v.Min = min
	return rb
}

// MinSegmentSize set the MinSegmentSize property for FielddataFrequencyFilterBuilder.
func (rb *FielddataFrequencyFilterBuilder) MinSegmentSize(minsegmentsize int) *FielddataFrequencyFilterBuilder {
	rb.v.MinSegmentSize = minsegmentsize
	return rb
}
