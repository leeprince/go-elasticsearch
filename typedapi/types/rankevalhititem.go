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

// RankEvalHitItem type.
type RankEvalHitItem struct {
	Hit    RankEvalHit `json:"hit"`
	Rating float64     `json:"rating,omitempty"`
}

// RankEvalHitItemBuilder holds RankEvalHitItem struct and provides a builder API.
type RankEvalHitItemBuilder struct {
	v *RankEvalHitItem
}

// NewRankEvalHitItem provides a builder for the RankEvalHitItem struct.
func NewRankEvalHitItem() *RankEvalHitItemBuilder {
	r := RankEvalHitItemBuilder{
		&RankEvalHitItem{},
	}

	return &r
}

// Build finalize the chain and returns the RankEvalHitItem struct
func (rb *RankEvalHitItemBuilder) Build() RankEvalHitItem {
	return *rb.v
}

// Hit set the Hit property for RankEvalHitItemBuilder.
func (rb *RankEvalHitItemBuilder) Hit(hit RankEvalHit) *RankEvalHitItemBuilder {
	rb.v.Hit = hit
	return rb
}

// Rating set the Rating property for RankEvalHitItemBuilder.
func (rb *RankEvalHitItemBuilder) Rating(rating float64) *RankEvalHitItemBuilder {
	rb.v.Rating = rating
	return rb
}
