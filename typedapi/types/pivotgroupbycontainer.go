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

// PivotGroupByContainer type.
type PivotGroupByContainer struct {
	DateHistogram *DateHistogramAggregation `json:"date_histogram,omitempty"`
	GeotileGrid   *GeoTileGridAggregation   `json:"geotile_grid,omitempty"`
	Histogram     *HistogramAggregation     `json:"histogram,omitempty"`
	Terms         *TermsAggregation         `json:"terms,omitempty"`
}

// PivotGroupByContainerBuilder holds PivotGroupByContainer struct and provides a builder API.
type PivotGroupByContainerBuilder struct {
	v *PivotGroupByContainer
}

// NewPivotGroupByContainer provides a builder for the PivotGroupByContainer struct.
func NewPivotGroupByContainer() *PivotGroupByContainerBuilder {
	r := PivotGroupByContainerBuilder{
		&PivotGroupByContainer{},
	}

	return &r
}

// Build finalize the chain and returns the PivotGroupByContainer struct
func (rb *PivotGroupByContainerBuilder) Build() PivotGroupByContainer {
	return *rb.v
}

// DateHistogram set the DateHistogram property for PivotGroupByContainerBuilder.
func (rb *PivotGroupByContainerBuilder) DateHistogram(datehistogram DateHistogramAggregation) *PivotGroupByContainerBuilder {
	rb.v.DateHistogram = &datehistogram
	return rb
}

// GeotileGrid set the GeotileGrid property for PivotGroupByContainerBuilder.
func (rb *PivotGroupByContainerBuilder) GeotileGrid(geotilegrid GeoTileGridAggregation) *PivotGroupByContainerBuilder {
	rb.v.GeotileGrid = &geotilegrid
	return rb
}

// Histogram set the Histogram property for PivotGroupByContainerBuilder.
func (rb *PivotGroupByContainerBuilder) Histogram(histogram HistogramAggregation) *PivotGroupByContainerBuilder {
	rb.v.Histogram = &histogram
	return rb
}

// Terms set the Terms property for PivotGroupByContainerBuilder.
func (rb *PivotGroupByContainerBuilder) Terms(terms TermsAggregation) *PivotGroupByContainerBuilder {
	rb.v.Terms = &terms
	return rb
}
