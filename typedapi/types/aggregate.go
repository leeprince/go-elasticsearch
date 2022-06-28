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

// Aggregate holds the union for the following types:
//     AdjacencyMatrixAggregate
//     AutoDateHistogramAggregate
//     AvgAggregate
//     BoxPlotAggregate
//     BucketMetricValueAggregate
//     CardinalityAggregate
//     ChildrenAggregate
//     CompositeAggregate
//     CumulativeCardinalityAggregate
//     DateHistogramAggregate
//     DateRangeAggregate
//     DerivativeAggregate
//     DoubleTermsAggregate
//     ExtendedStatsAggregate
//     ExtendedStatsBucketAggregate
//     FilterAggregate
//     FiltersAggregate
//     GeoBoundsAggregate
//     GeoCentroidAggregate
//     GeoDistanceAggregate
//     GeoHashGridAggregate
//     GeoLineAggregate
//     GeoTileGridAggregate
//     GlobalAggregate
//     HdrPercentileRanksAggregate
//     HdrPercentilesAggregate
//     HistogramAggregate
//     InferenceAggregate
//     IpRangeAggregate
//     LongRareTermsAggregate
//     LongTermsAggregate
//     MatrixStatsAggregate
//     MaxAggregate
//     MedianAbsoluteDeviationAggregate
//     MinAggregate
//     MissingAggregate
//     MultiTermsAggregate
//     NestedAggregate
//     ParentAggregate
//     PercentilesBucketAggregate
//     RangeAggregate
//     RateAggregate
//     ReverseNestedAggregate
//     SamplerAggregate
//     ScriptedMetricAggregate
//     SignificantLongTermsAggregate
//     SignificantStringTermsAggregate
//     SimpleValueAggregate
//     StatsAggregate
//     StatsBucketAggregate
//     StringRareTermsAggregate
//     StringStatsAggregate
//     StringTermsAggregate
//     SumAggregate
//     TDigestPercentileRanksAggregate
//     TDigestPercentilesAggregate
//     TTestAggregate
//     TopHitsAggregate
//     TopMetricsAggregate
//     UnmappedRareTermsAggregate
//     UnmappedSamplerAggregate
//     UnmappedSignificantTermsAggregate
//     UnmappedTermsAggregate
//     ValueCountAggregate
//     VariableWidthHistogramAggregate
//     WeightedAvgAggregate
type Aggregate interface{}

// AggregateBuilder holds Aggregate struct and provides a builder API.
type AggregateBuilder struct {
	v Aggregate
}

// NewAggregate provides a builder for the Aggregate struct.
func NewAggregate() *AggregateBuilder {
	return &AggregateBuilder{}
}

// Build finalize the chain and returns the Aggregate struct
func (u *AggregateBuilder) Build() Aggregate {
	return u.v
}

// AdjacencyMatrixAggregate set the AdjacencyMatrixAggregate property for AggregateBuilder.
func (u *AggregateBuilder) AdjacencyMatrixAggregate(v AdjacencyMatrixAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// AutoDateHistogramAggregate set the AutoDateHistogramAggregate property for AggregateBuilder.
func (u *AggregateBuilder) AutoDateHistogramAggregate(v AutoDateHistogramAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// AvgAggregate set the AvgAggregate property for AggregateBuilder.
func (u *AggregateBuilder) AvgAggregate(v AvgAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// BoxPlotAggregate set the BoxPlotAggregate property for AggregateBuilder.
func (u *AggregateBuilder) BoxPlotAggregate(v BoxPlotAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// BucketMetricValueAggregate set the BucketMetricValueAggregate property for AggregateBuilder.
func (u *AggregateBuilder) BucketMetricValueAggregate(v BucketMetricValueAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// CardinalityAggregate set the CardinalityAggregate property for AggregateBuilder.
func (u *AggregateBuilder) CardinalityAggregate(v CardinalityAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// ChildrenAggregate set the ChildrenAggregate property for AggregateBuilder.
func (u *AggregateBuilder) ChildrenAggregate(v ChildrenAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// CompositeAggregate set the CompositeAggregate property for AggregateBuilder.
func (u *AggregateBuilder) CompositeAggregate(v CompositeAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// CumulativeCardinalityAggregate set the CumulativeCardinalityAggregate property for AggregateBuilder.
func (u *AggregateBuilder) CumulativeCardinalityAggregate(v CumulativeCardinalityAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// DateHistogramAggregate set the DateHistogramAggregate property for AggregateBuilder.
func (u *AggregateBuilder) DateHistogramAggregate(v DateHistogramAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// DateRangeAggregate set the DateRangeAggregate property for AggregateBuilder.
func (u *AggregateBuilder) DateRangeAggregate(v DateRangeAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// DerivativeAggregate set the DerivativeAggregate property for AggregateBuilder.
func (u *AggregateBuilder) DerivativeAggregate(v DerivativeAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// DoubleTermsAggregate set the DoubleTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) DoubleTermsAggregate(v DoubleTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// ExtendedStatsAggregate set the ExtendedStatsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) ExtendedStatsAggregate(v ExtendedStatsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// ExtendedStatsBucketAggregate set the ExtendedStatsBucketAggregate property for AggregateBuilder.
func (u *AggregateBuilder) ExtendedStatsBucketAggregate(v ExtendedStatsBucketAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// FilterAggregate set the FilterAggregate property for AggregateBuilder.
func (u *AggregateBuilder) FilterAggregate(v FilterAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// FiltersAggregate set the FiltersAggregate property for AggregateBuilder.
func (u *AggregateBuilder) FiltersAggregate(v FiltersAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// GeoBoundsAggregate set the GeoBoundsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) GeoBoundsAggregate(v GeoBoundsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// GeoCentroidAggregate set the GeoCentroidAggregate property for AggregateBuilder.
func (u *AggregateBuilder) GeoCentroidAggregate(v GeoCentroidAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// GeoDistanceAggregate set the GeoDistanceAggregate property for AggregateBuilder.
func (u *AggregateBuilder) GeoDistanceAggregate(v GeoDistanceAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// GeoHashGridAggregate set the GeoHashGridAggregate property for AggregateBuilder.
func (u *AggregateBuilder) GeoHashGridAggregate(v GeoHashGridAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// GeoLineAggregate set the GeoLineAggregate property for AggregateBuilder.
func (u *AggregateBuilder) GeoLineAggregate(v GeoLineAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// GeoTileGridAggregate set the GeoTileGridAggregate property for AggregateBuilder.
func (u *AggregateBuilder) GeoTileGridAggregate(v GeoTileGridAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// GlobalAggregate set the GlobalAggregate property for AggregateBuilder.
func (u *AggregateBuilder) GlobalAggregate(v GlobalAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// HdrPercentileRanksAggregate set the HdrPercentileRanksAggregate property for AggregateBuilder.
func (u *AggregateBuilder) HdrPercentileRanksAggregate(v HdrPercentileRanksAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// HdrPercentilesAggregate set the HdrPercentilesAggregate property for AggregateBuilder.
func (u *AggregateBuilder) HdrPercentilesAggregate(v HdrPercentilesAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// HistogramAggregate set the HistogramAggregate property for AggregateBuilder.
func (u *AggregateBuilder) HistogramAggregate(v HistogramAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// InferenceAggregate set the InferenceAggregate property for AggregateBuilder.
func (u *AggregateBuilder) InferenceAggregate(v InferenceAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// IpRangeAggregate set the IpRangeAggregate property for AggregateBuilder.
func (u *AggregateBuilder) IpRangeAggregate(v IpRangeAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// LongRareTermsAggregate set the LongRareTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) LongRareTermsAggregate(v LongRareTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// LongTermsAggregate set the LongTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) LongTermsAggregate(v LongTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// MatrixStatsAggregate set the MatrixStatsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) MatrixStatsAggregate(v MatrixStatsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// MaxAggregate set the MaxAggregate property for AggregateBuilder.
func (u *AggregateBuilder) MaxAggregate(v MaxAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// MedianAbsoluteDeviationAggregate set the MedianAbsoluteDeviationAggregate property for AggregateBuilder.
func (u *AggregateBuilder) MedianAbsoluteDeviationAggregate(v MedianAbsoluteDeviationAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// MinAggregate set the MinAggregate property for AggregateBuilder.
func (u *AggregateBuilder) MinAggregate(v MinAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// MissingAggregate set the MissingAggregate property for AggregateBuilder.
func (u *AggregateBuilder) MissingAggregate(v MissingAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// MultiTermsAggregate set the MultiTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) MultiTermsAggregate(v MultiTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// NestedAggregate set the NestedAggregate property for AggregateBuilder.
func (u *AggregateBuilder) NestedAggregate(v NestedAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// ParentAggregate set the ParentAggregate property for AggregateBuilder.
func (u *AggregateBuilder) ParentAggregate(v ParentAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// PercentilesBucketAggregate set the PercentilesBucketAggregate property for AggregateBuilder.
func (u *AggregateBuilder) PercentilesBucketAggregate(v PercentilesBucketAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// RangeAggregate set the RangeAggregate property for AggregateBuilder.
func (u *AggregateBuilder) RangeAggregate(v RangeAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// RateAggregate set the RateAggregate property for AggregateBuilder.
func (u *AggregateBuilder) RateAggregate(v RateAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// ReverseNestedAggregate set the ReverseNestedAggregate property for AggregateBuilder.
func (u *AggregateBuilder) ReverseNestedAggregate(v ReverseNestedAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// SamplerAggregate set the SamplerAggregate property for AggregateBuilder.
func (u *AggregateBuilder) SamplerAggregate(v SamplerAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// ScriptedMetricAggregate set the ScriptedMetricAggregate property for AggregateBuilder.
func (u *AggregateBuilder) ScriptedMetricAggregate(v ScriptedMetricAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// SignificantLongTermsAggregate set the SignificantLongTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) SignificantLongTermsAggregate(v SignificantLongTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// SignificantStringTermsAggregate set the SignificantStringTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) SignificantStringTermsAggregate(v SignificantStringTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// SimpleValueAggregate set the SimpleValueAggregate property for AggregateBuilder.
func (u *AggregateBuilder) SimpleValueAggregate(v SimpleValueAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// StatsAggregate set the StatsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) StatsAggregate(v StatsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// StatsBucketAggregate set the StatsBucketAggregate property for AggregateBuilder.
func (u *AggregateBuilder) StatsBucketAggregate(v StatsBucketAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// StringRareTermsAggregate set the StringRareTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) StringRareTermsAggregate(v StringRareTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// StringStatsAggregate set the StringStatsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) StringStatsAggregate(v StringStatsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// StringTermsAggregate set the StringTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) StringTermsAggregate(v StringTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// SumAggregate set the SumAggregate property for AggregateBuilder.
func (u *AggregateBuilder) SumAggregate(v SumAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// TDigestPercentileRanksAggregate set the TDigestPercentileRanksAggregate property for AggregateBuilder.
func (u *AggregateBuilder) TDigestPercentileRanksAggregate(v TDigestPercentileRanksAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// TDigestPercentilesAggregate set the TDigestPercentilesAggregate property for AggregateBuilder.
func (u *AggregateBuilder) TDigestPercentilesAggregate(v TDigestPercentilesAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// TTestAggregate set the TTestAggregate property for AggregateBuilder.
func (u *AggregateBuilder) TTestAggregate(v TTestAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// TopHitsAggregate set the TopHitsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) TopHitsAggregate(v TopHitsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// TopMetricsAggregate set the TopMetricsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) TopMetricsAggregate(v TopMetricsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// UnmappedRareTermsAggregate set the UnmappedRareTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) UnmappedRareTermsAggregate(v UnmappedRareTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// UnmappedSamplerAggregate set the UnmappedSamplerAggregate property for AggregateBuilder.
func (u *AggregateBuilder) UnmappedSamplerAggregate(v UnmappedSamplerAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// UnmappedSignificantTermsAggregate set the UnmappedSignificantTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) UnmappedSignificantTermsAggregate(v UnmappedSignificantTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// UnmappedTermsAggregate set the UnmappedTermsAggregate property for AggregateBuilder.
func (u *AggregateBuilder) UnmappedTermsAggregate(v UnmappedTermsAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// ValueCountAggregate set the ValueCountAggregate property for AggregateBuilder.
func (u *AggregateBuilder) ValueCountAggregate(v ValueCountAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// VariableWidthHistogramAggregate set the VariableWidthHistogramAggregate property for AggregateBuilder.
func (u *AggregateBuilder) VariableWidthHistogramAggregate(v VariableWidthHistogramAggregate) *AggregateBuilder {
	u.v = v
	return u
}

// WeightedAvgAggregate set the WeightedAvgAggregate property for AggregateBuilder.
func (u *AggregateBuilder) WeightedAvgAggregate(v WeightedAvgAggregate) *AggregateBuilder {
	u.v = v
	return u
}
