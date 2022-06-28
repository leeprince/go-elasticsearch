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

// ExtendedStatsAggregate type.
type ExtendedStatsAggregate struct {
	Avg                        float64                          `json:"avg,omitempty"`
	AvgAsString                *string                          `json:"avg_as_string,omitempty"`
	Count                      int64                            `json:"count"`
	Max                        float64                          `json:"max,omitempty"`
	MaxAsString                *string                          `json:"max_as_string,omitempty"`
	Meta                       *Metadata                        `json:"meta,omitempty"`
	Min                        float64                          `json:"min,omitempty"`
	MinAsString                *string                          `json:"min_as_string,omitempty"`
	StdDeviation               float64                          `json:"std_deviation,omitempty"`
	StdDeviationAsString       *string                          `json:"std_deviation_as_string,omitempty"`
	StdDeviationBounds         *StandardDeviationBounds         `json:"std_deviation_bounds,omitempty"`
	StdDeviationBoundsAsString *StandardDeviationBoundsAsString `json:"std_deviation_bounds_as_string,omitempty"`
	Sum                        float64                          `json:"sum"`
	SumAsString                *string                          `json:"sum_as_string,omitempty"`
	SumOfSquares               float64                          `json:"sum_of_squares,omitempty"`
	SumOfSquaresAsString       *string                          `json:"sum_of_squares_as_string,omitempty"`
	Variance                   float64                          `json:"variance,omitempty"`
	VarianceAsString           *string                          `json:"variance_as_string,omitempty"`
	VariancePopulation         float64                          `json:"variance_population,omitempty"`
	VariancePopulationAsString *string                          `json:"variance_population_as_string,omitempty"`
	VarianceSampling           float64                          `json:"variance_sampling,omitempty"`
	VarianceSamplingAsString   *string                          `json:"variance_sampling_as_string,omitempty"`
}

// ExtendedStatsAggregateBuilder holds ExtendedStatsAggregate struct and provides a builder API.
type ExtendedStatsAggregateBuilder struct {
	v *ExtendedStatsAggregate
}

// NewExtendedStatsAggregate provides a builder for the ExtendedStatsAggregate struct.
func NewExtendedStatsAggregate() *ExtendedStatsAggregateBuilder {
	r := ExtendedStatsAggregateBuilder{
		&ExtendedStatsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the ExtendedStatsAggregate struct
func (rb *ExtendedStatsAggregateBuilder) Build() ExtendedStatsAggregate {
	return *rb.v
}

// Avg set the Avg property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) Avg(avg float64) *ExtendedStatsAggregateBuilder {
	rb.v.Avg = avg
	return rb
}

// AvgAsString set the AvgAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) AvgAsString(avgasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.AvgAsString = &avgasstring
	return rb
}

// Count set the Count property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) Count(count int64) *ExtendedStatsAggregateBuilder {
	rb.v.Count = count
	return rb
}

// Max set the Max property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) Max(max float64) *ExtendedStatsAggregateBuilder {
	rb.v.Max = max
	return rb
}

// MaxAsString set the MaxAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) MaxAsString(maxasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.MaxAsString = &maxasstring
	return rb
}

// Meta set the Meta property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) Meta(meta Metadata) *ExtendedStatsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// Min set the Min property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) Min(min float64) *ExtendedStatsAggregateBuilder {
	rb.v.Min = min
	return rb
}

// MinAsString set the MinAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) MinAsString(minasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.MinAsString = &minasstring
	return rb
}

// StdDeviation set the StdDeviation property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) StdDeviation(stddeviation float64) *ExtendedStatsAggregateBuilder {
	rb.v.StdDeviation = stddeviation
	return rb
}

// StdDeviationAsString set the StdDeviationAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) StdDeviationAsString(stddeviationasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.StdDeviationAsString = &stddeviationasstring
	return rb
}

// StdDeviationBounds set the StdDeviationBounds property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) StdDeviationBounds(stddeviationbounds StandardDeviationBounds) *ExtendedStatsAggregateBuilder {
	rb.v.StdDeviationBounds = &stddeviationbounds
	return rb
}

// StdDeviationBoundsAsString set the StdDeviationBoundsAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) StdDeviationBoundsAsString(stddeviationboundsasstring StandardDeviationBoundsAsString) *ExtendedStatsAggregateBuilder {
	rb.v.StdDeviationBoundsAsString = &stddeviationboundsasstring
	return rb
}

// Sum set the Sum property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) Sum(sum float64) *ExtendedStatsAggregateBuilder {
	rb.v.Sum = sum
	return rb
}

// SumAsString set the SumAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) SumAsString(sumasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.SumAsString = &sumasstring
	return rb
}

// SumOfSquares set the SumOfSquares property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) SumOfSquares(sumofsquares float64) *ExtendedStatsAggregateBuilder {
	rb.v.SumOfSquares = sumofsquares
	return rb
}

// SumOfSquaresAsString set the SumOfSquaresAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) SumOfSquaresAsString(sumofsquaresasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.SumOfSquaresAsString = &sumofsquaresasstring
	return rb
}

// Variance set the Variance property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) Variance(variance float64) *ExtendedStatsAggregateBuilder {
	rb.v.Variance = variance
	return rb
}

// VarianceAsString set the VarianceAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) VarianceAsString(varianceasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.VarianceAsString = &varianceasstring
	return rb
}

// VariancePopulation set the VariancePopulation property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) VariancePopulation(variancepopulation float64) *ExtendedStatsAggregateBuilder {
	rb.v.VariancePopulation = variancepopulation
	return rb
}

// VariancePopulationAsString set the VariancePopulationAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) VariancePopulationAsString(variancepopulationasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.VariancePopulationAsString = &variancepopulationasstring
	return rb
}

// VarianceSampling set the VarianceSampling property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) VarianceSampling(variancesampling float64) *ExtendedStatsAggregateBuilder {
	rb.v.VarianceSampling = variancesampling
	return rb
}

// VarianceSamplingAsString set the VarianceSamplingAsString property for ExtendedStatsAggregateBuilder.
func (rb *ExtendedStatsAggregateBuilder) VarianceSamplingAsString(variancesamplingasstring string) *ExtendedStatsAggregateBuilder {
	rb.v.VarianceSamplingAsString = &variancesamplingasstring
	return rb
}
