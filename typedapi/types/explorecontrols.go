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

// ExploreControls type.
type ExploreControls struct {
	SampleDiversity *SampleDiversity `json:"sample_diversity,omitempty"`
	SampleSize      *int             `json:"sample_size,omitempty"`
	Timeout         *Time            `json:"timeout,omitempty"`
	UseSignificance bool             `json:"use_significance"`
}

// ExploreControlsBuilder holds ExploreControls struct and provides a builder API.
type ExploreControlsBuilder struct {
	v *ExploreControls
}

// NewExploreControls provides a builder for the ExploreControls struct.
func NewExploreControls() *ExploreControlsBuilder {
	r := ExploreControlsBuilder{
		&ExploreControls{},
	}

	return &r
}

// Build finalize the chain and returns the ExploreControls struct
func (rb *ExploreControlsBuilder) Build() ExploreControls {
	return *rb.v
}

// SampleDiversity set the SampleDiversity property for ExploreControlsBuilder.
func (rb *ExploreControlsBuilder) SampleDiversity(samplediversity SampleDiversity) *ExploreControlsBuilder {
	rb.v.SampleDiversity = &samplediversity
	return rb
}

// SampleSize set the SampleSize property for ExploreControlsBuilder.
func (rb *ExploreControlsBuilder) SampleSize(samplesize int) *ExploreControlsBuilder {
	rb.v.SampleSize = &samplesize
	return rb
}

// Timeout set the Timeout property for ExploreControlsBuilder.
func (rb *ExploreControlsBuilder) Timeout(timeout Time) *ExploreControlsBuilder {
	rb.v.Timeout = &timeout
	return rb
}

// UseSignificance set the UseSignificance property for ExploreControlsBuilder.
func (rb *ExploreControlsBuilder) UseSignificance(usesignificance bool) *ExploreControlsBuilder {
	rb.v.UseSignificance = usesignificance
	return rb
}
