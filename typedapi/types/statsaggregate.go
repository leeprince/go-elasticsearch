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

// StatsAggregate type.
type StatsAggregate struct {
	Avg         float64   `json:"avg,omitempty"`
	AvgAsString *string   `json:"avg_as_string,omitempty"`
	Count       int64     `json:"count"`
	Max         float64   `json:"max,omitempty"`
	MaxAsString *string   `json:"max_as_string,omitempty"`
	Meta        *Metadata `json:"meta,omitempty"`
	Min         float64   `json:"min,omitempty"`
	MinAsString *string   `json:"min_as_string,omitempty"`
	Sum         float64   `json:"sum"`
	SumAsString *string   `json:"sum_as_string,omitempty"`
}

// StatsAggregateBuilder holds StatsAggregate struct and provides a builder API.
type StatsAggregateBuilder struct {
	v *StatsAggregate
}

// NewStatsAggregate provides a builder for the StatsAggregate struct.
func NewStatsAggregate() *StatsAggregateBuilder {
	r := StatsAggregateBuilder{
		&StatsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the StatsAggregate struct
func (rb *StatsAggregateBuilder) Build() StatsAggregate {
	return *rb.v
}

// Avg set the Avg property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) Avg(avg float64) *StatsAggregateBuilder {
	rb.v.Avg = avg
	return rb
}

// AvgAsString set the AvgAsString property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) AvgAsString(avgasstring string) *StatsAggregateBuilder {
	rb.v.AvgAsString = &avgasstring
	return rb
}

// Count set the Count property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) Count(count int64) *StatsAggregateBuilder {
	rb.v.Count = count
	return rb
}

// Max set the Max property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) Max(max float64) *StatsAggregateBuilder {
	rb.v.Max = max
	return rb
}

// MaxAsString set the MaxAsString property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) MaxAsString(maxasstring string) *StatsAggregateBuilder {
	rb.v.MaxAsString = &maxasstring
	return rb
}

// Meta set the Meta property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) Meta(meta Metadata) *StatsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// Min set the Min property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) Min(min float64) *StatsAggregateBuilder {
	rb.v.Min = min
	return rb
}

// MinAsString set the MinAsString property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) MinAsString(minasstring string) *StatsAggregateBuilder {
	rb.v.MinAsString = &minasstring
	return rb
}

// Sum set the Sum property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) Sum(sum float64) *StatsAggregateBuilder {
	rb.v.Sum = sum
	return rb
}

// SumAsString set the SumAsString property for StatsAggregateBuilder.
func (rb *StatsAggregateBuilder) SumAsString(sumasstring string) *StatsAggregateBuilder {
	rb.v.SumAsString = &sumasstring
	return rb
}
