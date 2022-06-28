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

// Influence type.
type Influence struct {
	InfluencerFieldName   string   `json:"influencer_field_name"`
	InfluencerFieldValues []string `json:"influencer_field_values"`
}

// InfluenceBuilder holds Influence struct and provides a builder API.
type InfluenceBuilder struct {
	v *Influence
}

// NewInfluence provides a builder for the Influence struct.
func NewInfluence() *InfluenceBuilder {
	r := InfluenceBuilder{
		&Influence{},
	}

	return &r
}

// Build finalize the chain and returns the Influence struct
func (rb *InfluenceBuilder) Build() Influence {
	return *rb.v
}

// InfluencerFieldName set the InfluencerFieldName property for InfluenceBuilder.
func (rb *InfluenceBuilder) InfluencerFieldName(influencerfieldname string) *InfluenceBuilder {
	rb.v.InfluencerFieldName = influencerfieldname
	return rb
}

// InfluencerFieldValues set the InfluencerFieldValues property for InfluenceBuilder.
func (rb *InfluenceBuilder) InfluencerFieldValues(influencer_field_values ...string) *InfluenceBuilder {
	rb.v.InfluencerFieldValues = append(rb.v.InfluencerFieldValues, influencer_field_values...)
	return rb
}
