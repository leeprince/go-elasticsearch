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

// Groupings type.
type Groupings struct {
	DateHistogram *DateHistogramGrouping `json:"date_histogram,omitempty"`
	Histogram     *HistogramGrouping     `json:"histogram,omitempty"`
	Terms         *TermsGrouping         `json:"terms,omitempty"`
}

// GroupingsBuilder holds Groupings struct and provides a builder API.
type GroupingsBuilder struct {
	v *Groupings
}

// NewGroupings provides a builder for the Groupings struct.
func NewGroupings() *GroupingsBuilder {
	r := GroupingsBuilder{
		&Groupings{},
	}

	return &r
}

// Build finalize the chain and returns the Groupings struct
func (rb *GroupingsBuilder) Build() Groupings {
	return *rb.v
}

// DateHistogram set the DateHistogram property for GroupingsBuilder.
func (rb *GroupingsBuilder) DateHistogram(datehistogram DateHistogramGrouping) *GroupingsBuilder {
	rb.v.DateHistogram = &datehistogram
	return rb
}

// Histogram set the Histogram property for GroupingsBuilder.
func (rb *GroupingsBuilder) Histogram(histogram HistogramGrouping) *GroupingsBuilder {
	rb.v.Histogram = &histogram
	return rb
}

// Terms set the Terms property for GroupingsBuilder.
func (rb *GroupingsBuilder) Terms(terms TermsGrouping) *GroupingsBuilder {
	rb.v.Terms = &terms
	return rb
}
