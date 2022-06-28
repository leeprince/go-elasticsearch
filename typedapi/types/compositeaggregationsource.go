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

// CompositeAggregationSource type.
type CompositeAggregationSource struct {
	DateHistogram *DateHistogramAggregation `json:"date_histogram,omitempty"`
	GeotileGrid   *GeoTileGridAggregation   `json:"geotile_grid,omitempty"`
	Histogram     *HistogramAggregation     `json:"histogram,omitempty"`
	Terms         *TermsAggregation         `json:"terms,omitempty"`
}

// CompositeAggregationSourceBuilder holds CompositeAggregationSource struct and provides a builder API.
type CompositeAggregationSourceBuilder struct {
	v *CompositeAggregationSource
}

// NewCompositeAggregationSource provides a builder for the CompositeAggregationSource struct.
func NewCompositeAggregationSource() *CompositeAggregationSourceBuilder {
	r := CompositeAggregationSourceBuilder{
		&CompositeAggregationSource{},
	}

	return &r
}

// Build finalize the chain and returns the CompositeAggregationSource struct
func (rb *CompositeAggregationSourceBuilder) Build() CompositeAggregationSource {
	return *rb.v
}

// DateHistogram set the DateHistogram property for CompositeAggregationSourceBuilder.
func (rb *CompositeAggregationSourceBuilder) DateHistogram(datehistogram DateHistogramAggregation) *CompositeAggregationSourceBuilder {
	rb.v.DateHistogram = &datehistogram
	return rb
}

// GeotileGrid set the GeotileGrid property for CompositeAggregationSourceBuilder.
func (rb *CompositeAggregationSourceBuilder) GeotileGrid(geotilegrid GeoTileGridAggregation) *CompositeAggregationSourceBuilder {
	rb.v.GeotileGrid = &geotilegrid
	return rb
}

// Histogram set the Histogram property for CompositeAggregationSourceBuilder.
func (rb *CompositeAggregationSourceBuilder) Histogram(histogram HistogramAggregation) *CompositeAggregationSourceBuilder {
	rb.v.Histogram = &histogram
	return rb
}

// Terms set the Terms property for CompositeAggregationSourceBuilder.
func (rb *CompositeAggregationSourceBuilder) Terms(terms TermsAggregation) *CompositeAggregationSourceBuilder {
	rb.v.Terms = &terms
	return rb
}
