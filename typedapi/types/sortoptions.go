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

// SortOptions type.
type SortOptions struct {
	Doc_         *ScoreSort          `json:"_doc,omitempty"`
	GeoDistance_ *GeoDistanceSort    `json:"_geo_distance,omitempty"`
	Score_       *ScoreSort          `json:"_score,omitempty"`
	Script_      *ScriptSort         `json:"_script,omitempty"`
	SortOptions  map[Field]FieldSort `json:"SortOptions,omitempty"`
}

// SortOptionsBuilder holds SortOptions struct and provides a builder API.
type SortOptionsBuilder struct {
	v *SortOptions
}

// NewSortOptions provides a builder for the SortOptions struct.
func NewSortOptions() *SortOptionsBuilder {
	r := SortOptionsBuilder{
		&SortOptions{
			SortOptions: make(map[Field]FieldSort, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SortOptions struct
func (rb *SortOptionsBuilder) Build() SortOptions {
	return *rb.v
}

// Doc_ set the Doc_ property for SortOptionsBuilder.
func (rb *SortOptionsBuilder) Doc_(doc_ ScoreSort) *SortOptionsBuilder {
	rb.v.Doc_ = &doc_
	return rb
}

// GeoDistance_ set the GeoDistance_ property for SortOptionsBuilder.
func (rb *SortOptionsBuilder) GeoDistance_(geodistance_ GeoDistanceSort) *SortOptionsBuilder {
	rb.v.GeoDistance_ = &geodistance_
	return rb
}

// Score_ set the Score_ property for SortOptionsBuilder.
func (rb *SortOptionsBuilder) Score_(score_ ScoreSort) *SortOptionsBuilder {
	rb.v.Score_ = &score_
	return rb
}

// Script_ set the Script_ property for SortOptionsBuilder.
func (rb *SortOptionsBuilder) Script_(script_ ScriptSort) *SortOptionsBuilder {
	rb.v.Script_ = &script_
	return rb
}

// SortOptions set the SortOptions property for SortOptionsBuilder.
func (rb *SortOptionsBuilder) SortOptions(value map[Field]FieldSort) *SortOptionsBuilder {
	rb.v.SortOptions = value
	return rb
}
