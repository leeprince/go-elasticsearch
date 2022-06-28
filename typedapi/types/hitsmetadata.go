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

// HitsMetadata type.
type HitsMetadata struct {
	Hits     []Hit   `json:"hits"`
	MaxScore float64 `json:"max_score,omitempty"`
	// Total Total hit count information, present only if `track_total_hits` wasn't
	// `false` in the search request.
	Total int64 `json:"total,omitempty"`
}

// HitsMetadataBuilder holds HitsMetadata struct and provides a builder API.
type HitsMetadataBuilder struct {
	v *HitsMetadata
}

// NewHitsMetadata provides a builder for the HitsMetadata struct.
func NewHitsMetadata() *HitsMetadataBuilder {
	r := HitsMetadataBuilder{
		&HitsMetadata{},
	}

	return &r
}

// Build finalize the chain and returns the HitsMetadata struct
func (rb *HitsMetadataBuilder) Build() HitsMetadata {
	return *rb.v
}

// Hits set the Hits property for HitsMetadataBuilder.
func (rb *HitsMetadataBuilder) Hits(hits ...Hit) *HitsMetadataBuilder {
	rb.v.Hits = append(rb.v.Hits, hits...)
	return rb
}

// MaxScore set the MaxScore property for HitsMetadataBuilder.
func (rb *HitsMetadataBuilder) MaxScore(maxscore float64) *HitsMetadataBuilder {
	rb.v.MaxScore = maxscore
	return rb
}

// Total Total hit count information, present only if `track_total_hits` wasn't
// `false` in the search request.
func (rb *HitsMetadataBuilder) Total(arg int64) *HitsMetadataBuilder {
	rb.v.Total = arg
	return rb
}
