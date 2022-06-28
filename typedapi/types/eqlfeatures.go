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

// EqlFeatures type.
type EqlFeatures struct {
	Event     uint                 `json:"event"`
	Join      uint                 `json:"join"`
	Joins     EqlFeaturesJoin      `json:"joins"`
	Keys      EqlFeaturesKeys      `json:"keys"`
	Pipes     EqlFeaturesPipes     `json:"pipes"`
	Sequence  uint                 `json:"sequence"`
	Sequences EqlFeaturesSequences `json:"sequences"`
}

// EqlFeaturesBuilder holds EqlFeatures struct and provides a builder API.
type EqlFeaturesBuilder struct {
	v *EqlFeatures
}

// NewEqlFeatures provides a builder for the EqlFeatures struct.
func NewEqlFeatures() *EqlFeaturesBuilder {
	r := EqlFeaturesBuilder{
		&EqlFeatures{},
	}

	return &r
}

// Build finalize the chain and returns the EqlFeatures struct
func (rb *EqlFeaturesBuilder) Build() EqlFeatures {
	return *rb.v
}

// Event set the Event property for EqlFeaturesBuilder.
func (rb *EqlFeaturesBuilder) Event(event uint) *EqlFeaturesBuilder {
	rb.v.Event = event
	return rb
}

// Join set the Join property for EqlFeaturesBuilder.
func (rb *EqlFeaturesBuilder) Join(join uint) *EqlFeaturesBuilder {
	rb.v.Join = join
	return rb
}

// Joins set the Joins property for EqlFeaturesBuilder.
func (rb *EqlFeaturesBuilder) Joins(joins EqlFeaturesJoin) *EqlFeaturesBuilder {
	rb.v.Joins = joins
	return rb
}

// Keys set the Keys property for EqlFeaturesBuilder.
func (rb *EqlFeaturesBuilder) Keys(keys EqlFeaturesKeys) *EqlFeaturesBuilder {
	rb.v.Keys = keys
	return rb
}

// Pipes set the Pipes property for EqlFeaturesBuilder.
func (rb *EqlFeaturesBuilder) Pipes(pipes EqlFeaturesPipes) *EqlFeaturesBuilder {
	rb.v.Pipes = pipes
	return rb
}

// Sequence set the Sequence property for EqlFeaturesBuilder.
func (rb *EqlFeaturesBuilder) Sequence(sequence uint) *EqlFeaturesBuilder {
	rb.v.Sequence = sequence
	return rb
}

// Sequences set the Sequences property for EqlFeaturesBuilder.
func (rb *EqlFeaturesBuilder) Sequences(sequences EqlFeaturesSequences) *EqlFeaturesBuilder {
	rb.v.Sequences = sequences
	return rb
}
