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

// Slm type.
type Slm struct {
	Available   bool        `json:"available"`
	Enabled     bool        `json:"enabled"`
	PolicyCount *int        `json:"policy_count,omitempty"`
	PolicyStats *Statistics `json:"policy_stats,omitempty"`
}

// SlmBuilder holds Slm struct and provides a builder API.
type SlmBuilder struct {
	v *Slm
}

// NewSlm provides a builder for the Slm struct.
func NewSlm() *SlmBuilder {
	r := SlmBuilder{
		&Slm{},
	}

	return &r
}

// Build finalize the chain and returns the Slm struct
func (rb *SlmBuilder) Build() Slm {
	return *rb.v
}

// Available set the Available property for SlmBuilder.
func (rb *SlmBuilder) Available(available bool) *SlmBuilder {
	rb.v.Available = available
	return rb
}

// Enabled set the Enabled property for SlmBuilder.
func (rb *SlmBuilder) Enabled(enabled bool) *SlmBuilder {
	rb.v.Enabled = enabled
	return rb
}

// PolicyCount set the PolicyCount property for SlmBuilder.
func (rb *SlmBuilder) PolicyCount(policycount int) *SlmBuilder {
	rb.v.PolicyCount = &policycount
	return rb
}

// PolicyStats set the PolicyStats property for SlmBuilder.
func (rb *SlmBuilder) PolicyStats(policystats Statistics) *SlmBuilder {
	rb.v.PolicyStats = &policystats
	return rb
}
