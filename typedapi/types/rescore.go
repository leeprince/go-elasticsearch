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

// Rescore type.
type Rescore struct {
	Query      RescoreQuery `json:"query"`
	WindowSize *int         `json:"window_size,omitempty"`
}

// RescoreBuilder holds Rescore struct and provides a builder API.
type RescoreBuilder struct {
	v *Rescore
}

// NewRescore provides a builder for the Rescore struct.
func NewRescore() *RescoreBuilder {
	r := RescoreBuilder{
		&Rescore{},
	}

	return &r
}

// Build finalize the chain and returns the Rescore struct
func (rb *RescoreBuilder) Build() Rescore {
	return *rb.v
}

// Query set the Query property for RescoreBuilder.
func (rb *RescoreBuilder) Query(query RescoreQuery) *RescoreBuilder {
	rb.v.Query = query
	return rb
}

// WindowSize set the WindowSize property for RescoreBuilder.
func (rb *RescoreBuilder) WindowSize(windowsize int) *RescoreBuilder {
	rb.v.WindowSize = &windowsize
	return rb
}
