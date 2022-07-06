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

// MutualInformationHeuristic type.
type MutualInformationHeuristic struct {
	BackgroundIsSuperset *bool `json:"background_is_superset,omitempty"`
	IncludeNegatives     *bool `json:"include_negatives,omitempty"`
}

// MutualInformationHeuristicBuilder holds MutualInformationHeuristic struct and provides a builder API.
type MutualInformationHeuristicBuilder struct {
	v *MutualInformationHeuristic
}

// NewMutualInformationHeuristic provides a builder for the MutualInformationHeuristic struct.
func NewMutualInformationHeuristic() *MutualInformationHeuristicBuilder {
	r := MutualInformationHeuristicBuilder{
		&MutualInformationHeuristic{},
	}

	return &r
}

// Build finalize the chain and returns the MutualInformationHeuristic struct
func (rb *MutualInformationHeuristicBuilder) Build() MutualInformationHeuristic {
	return *rb.v
}

// BackgroundIsSuperset set the BackgroundIsSuperset property for MutualInformationHeuristicBuilder.
func (rb *MutualInformationHeuristicBuilder) BackgroundIsSuperset(backgroundissuperset bool) *MutualInformationHeuristicBuilder {
	rb.v.BackgroundIsSuperset = &backgroundissuperset
	return rb
}

// IncludeNegatives set the IncludeNegatives property for MutualInformationHeuristicBuilder.
func (rb *MutualInformationHeuristicBuilder) IncludeNegatives(includenegatives bool) *MutualInformationHeuristicBuilder {
	rb.v.IncludeNegatives = &includenegatives
	return rb
}
