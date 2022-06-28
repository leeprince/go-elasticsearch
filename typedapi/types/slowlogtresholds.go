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

// SlowlogTresholds type.
type SlowlogTresholds struct {
	Fetch *SlowlogTresholdLevels `json:"fetch,omitempty"`
	// Index The indexing slow log, similar in functionality to the search slow log. The
	// log file name ends with `_index_indexing_slowlog.json`.
	// Log and the thresholds are configured in the same way as the search slowlog.
	Index *SlowlogTresholdLevels `json:"index,omitempty"`
	Query *SlowlogTresholdLevels `json:"query,omitempty"`
}

// SlowlogTresholdsBuilder holds SlowlogTresholds struct and provides a builder API.
type SlowlogTresholdsBuilder struct {
	v *SlowlogTresholds
}

// NewSlowlogTresholds provides a builder for the SlowlogTresholds struct.
func NewSlowlogTresholds() *SlowlogTresholdsBuilder {
	r := SlowlogTresholdsBuilder{
		&SlowlogTresholds{},
	}

	return &r
}

// Build finalize the chain and returns the SlowlogTresholds struct
func (rb *SlowlogTresholdsBuilder) Build() SlowlogTresholds {
	return *rb.v
}

// Fetch set the Fetch property for SlowlogTresholdsBuilder.
func (rb *SlowlogTresholdsBuilder) Fetch(fetch SlowlogTresholdLevels) *SlowlogTresholdsBuilder {
	rb.v.Fetch = &fetch
	return rb
}

// Index The indexing slow log, similar in functionality to the search slow log. The
// log file name ends with `_index_indexing_slowlog.json`.
// Log and the thresholds are configured in the same way as the search slowlog.
func (rb *SlowlogTresholdsBuilder) Index(index SlowlogTresholdLevels) *SlowlogTresholdsBuilder {
	rb.v.Index = &index
	return rb
}

// Query set the Query property for SlowlogTresholdsBuilder.
func (rb *SlowlogTresholdsBuilder) Query(query SlowlogTresholdLevels) *SlowlogTresholdsBuilder {
	rb.v.Query = &query
	return rb
}
