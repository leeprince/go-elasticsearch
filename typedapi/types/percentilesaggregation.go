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

// PercentilesAggregation type.
type PercentilesAggregation struct {
	Field    *Field     `json:"field,omitempty"`
	Format   *string    `json:"format,omitempty"`
	Hdr      *HdrMethod `json:"hdr,omitempty"`
	Keyed    *bool      `json:"keyed,omitempty"`
	Missing  *Missing   `json:"missing,omitempty"`
	Percents []float64  `json:"percents,omitempty"`
	Script   *Script    `json:"script,omitempty"`
	Tdigest  *TDigest   `json:"tdigest,omitempty"`
}

// PercentilesAggregationBuilder holds PercentilesAggregation struct and provides a builder API.
type PercentilesAggregationBuilder struct {
	v *PercentilesAggregation
}

// NewPercentilesAggregation provides a builder for the PercentilesAggregation struct.
func NewPercentilesAggregation() *PercentilesAggregationBuilder {
	r := PercentilesAggregationBuilder{
		&PercentilesAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the PercentilesAggregation struct
func (rb *PercentilesAggregationBuilder) Build() PercentilesAggregation {
	return *rb.v
}

// Field set the Field property for PercentilesAggregationBuilder.
func (rb *PercentilesAggregationBuilder) Field(field Field) *PercentilesAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for PercentilesAggregationBuilder.
func (rb *PercentilesAggregationBuilder) Format(format string) *PercentilesAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Hdr set the Hdr property for PercentilesAggregationBuilder.
func (rb *PercentilesAggregationBuilder) Hdr(hdr HdrMethod) *PercentilesAggregationBuilder {
	rb.v.Hdr = &hdr
	return rb
}

// Keyed set the Keyed property for PercentilesAggregationBuilder.
func (rb *PercentilesAggregationBuilder) Keyed(keyed bool) *PercentilesAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

// Missing set the Missing property for PercentilesAggregationBuilder.
func (rb *PercentilesAggregationBuilder) Missing(missing Missing) *PercentilesAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Percents set the Percents property for PercentilesAggregationBuilder.
func (rb *PercentilesAggregationBuilder) Percents(percents ...float64) *PercentilesAggregationBuilder {
	rb.v.Percents = append(rb.v.Percents, percents...)
	return rb
}

// Script set the Script property for PercentilesAggregationBuilder.
func (rb *PercentilesAggregationBuilder) Script(script Script) *PercentilesAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// Tdigest set the Tdigest property for PercentilesAggregationBuilder.
func (rb *PercentilesAggregationBuilder) Tdigest(tdigest TDigest) *PercentilesAggregationBuilder {
	rb.v.Tdigest = &tdigest
	return rb
}
