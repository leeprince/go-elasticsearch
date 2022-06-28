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

// Analytics type.
type Analytics struct {
	Available bool                `json:"available"`
	Enabled   bool                `json:"enabled"`
	Stats     AnalyticsStatistics `json:"stats"`
}

// AnalyticsBuilder holds Analytics struct and provides a builder API.
type AnalyticsBuilder struct {
	v *Analytics
}

// NewAnalytics provides a builder for the Analytics struct.
func NewAnalytics() *AnalyticsBuilder {
	r := AnalyticsBuilder{
		&Analytics{},
	}

	return &r
}

// Build finalize the chain and returns the Analytics struct
func (rb *AnalyticsBuilder) Build() Analytics {
	return *rb.v
}

// Available set the Available property for AnalyticsBuilder.
func (rb *AnalyticsBuilder) Available(available bool) *AnalyticsBuilder {
	rb.v.Available = available
	return rb
}

// Enabled set the Enabled property for AnalyticsBuilder.
func (rb *AnalyticsBuilder) Enabled(enabled bool) *AnalyticsBuilder {
	rb.v.Enabled = enabled
	return rb
}

// Stats set the Stats property for AnalyticsBuilder.
func (rb *AnalyticsBuilder) Stats(stats AnalyticsStatistics) *AnalyticsBuilder {
	rb.v.Stats = stats
	return rb
}
