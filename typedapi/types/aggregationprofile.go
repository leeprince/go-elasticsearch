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

// AggregationProfile type.
type AggregationProfile struct {
	Breakdown   AggregationBreakdown     `json:"breakdown"`
	Children    []AggregationProfile     `json:"children,omitempty"`
	Debug       *AggregationProfileDebug `json:"debug,omitempty"`
	Description string                   `json:"description"`
	TimeInNanos int64                    `json:"time_in_nanos"`
	Type        string                   `json:"type"`
}

// AggregationProfileBuilder holds AggregationProfile struct and provides a builder API.
type AggregationProfileBuilder struct {
	v *AggregationProfile
}

// NewAggregationProfile provides a builder for the AggregationProfile struct.
func NewAggregationProfile() *AggregationProfileBuilder {
	r := AggregationProfileBuilder{
		&AggregationProfile{},
	}

	return &r
}

// Build finalize the chain and returns the AggregationProfile struct
func (rb *AggregationProfileBuilder) Build() AggregationProfile {
	return *rb.v
}

// Breakdown set the Breakdown property for AggregationProfileBuilder.
func (rb *AggregationProfileBuilder) Breakdown(breakdown AggregationBreakdown) *AggregationProfileBuilder {
	rb.v.Breakdown = breakdown
	return rb
}

// Children set the Children property for AggregationProfileBuilder.
func (rb *AggregationProfileBuilder) Children(children ...AggregationProfile) *AggregationProfileBuilder {
	rb.v.Children = append(rb.v.Children, children...)
	return rb
}

// Debug set the Debug property for AggregationProfileBuilder.
func (rb *AggregationProfileBuilder) Debug(debug AggregationProfileDebug) *AggregationProfileBuilder {
	rb.v.Debug = &debug
	return rb
}

// Description set the Description property for AggregationProfileBuilder.
func (rb *AggregationProfileBuilder) Description(description string) *AggregationProfileBuilder {
	rb.v.Description = description
	return rb
}

// TimeInNanos set the TimeInNanos property for AggregationProfileBuilder.
func (rb *AggregationProfileBuilder) TimeInNanos(timeinnanos int64) *AggregationProfileBuilder {
	rb.v.TimeInNanos = timeinnanos
	return rb
}

// Type set the Type property for AggregationProfileBuilder.
func (rb *AggregationProfileBuilder) Type_(type_ string) *AggregationProfileBuilder {
	rb.v.Type = type_
	return rb
}
