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

// AggregationContainer type.
type AggregationContainer struct {
	AdjacencyMatrix *AdjacencyMatrixAggregation `json:"adjacency_matrix,omitempty"`
	// Aggregations Sub-aggregations for this aggregation. Only applies to bucket aggregations.
	Aggregations            map[string]AggregationContainer     `json:"aggregations,omitempty"`
	AutoDateHistogram       *AutoDateHistogramAggregation       `json:"auto_date_histogram,omitempty"`
	Avg                     *AverageAggregation                 `json:"avg,omitempty"`
	AvgBucket               *AverageBucketAggregation           `json:"avg_bucket,omitempty"`
	Boxplot                 *BoxplotAggregation                 `json:"boxplot,omitempty"`
	BucketCorrelation       *BucketCorrelationAggregation       `json:"bucket_correlation,omitempty"`
	BucketCountKsTest       *BucketKsAggregation                `json:"bucket_count_ks_test,omitempty"`
	BucketScript            *BucketScriptAggregation            `json:"bucket_script,omitempty"`
	BucketSelector          *BucketSelectorAggregation          `json:"bucket_selector,omitempty"`
	BucketSort              *BucketSortAggregation              `json:"bucket_sort,omitempty"`
	Cardinality             *CardinalityAggregation             `json:"cardinality,omitempty"`
	CategorizeText          *CategorizeTextAggregation          `json:"categorize_text,omitempty"`
	Children                *ChildrenAggregation                `json:"children,omitempty"`
	Composite               *CompositeAggregation               `json:"composite,omitempty"`
	CumulativeCardinality   *CumulativeCardinalityAggregation   `json:"cumulative_cardinality,omitempty"`
	CumulativeSum           *CumulativeSumAggregation           `json:"cumulative_sum,omitempty"`
	DateHistogram           *DateHistogramAggregation           `json:"date_histogram,omitempty"`
	DateRange               *DateRangeAggregation               `json:"date_range,omitempty"`
	Derivative              *DerivativeAggregation              `json:"derivative,omitempty"`
	DiversifiedSampler      *DiversifiedSamplerAggregation      `json:"diversified_sampler,omitempty"`
	ExtendedStats           *ExtendedStatsAggregation           `json:"extended_stats,omitempty"`
	ExtendedStatsBucket     *ExtendedStatsBucketAggregation     `json:"extended_stats_bucket,omitempty"`
	Filter                  *QueryContainer                     `json:"filter,omitempty"`
	Filters                 *FiltersAggregation                 `json:"filters,omitempty"`
	GeoBounds               *GeoBoundsAggregation               `json:"geo_bounds,omitempty"`
	GeoCentroid             *GeoCentroidAggregation             `json:"geo_centroid,omitempty"`
	GeoDistance             *GeoDistanceAggregation             `json:"geo_distance,omitempty"`
	GeoLine                 *GeoLineAggregation                 `json:"geo_line,omitempty"`
	GeohashGrid             *GeoHashGridAggregation             `json:"geohash_grid,omitempty"`
	GeohexGrid              *GeohexGridAggregation              `json:"geohex_grid,omitempty"`
	GeotileGrid             *GeoTileGridAggregation             `json:"geotile_grid,omitempty"`
	Global                  *GlobalAggregation                  `json:"global,omitempty"`
	Histogram               *HistogramAggregation               `json:"histogram,omitempty"`
	Inference               *InferenceAggregation               `json:"inference,omitempty"`
	IpRange                 *IpRangeAggregation                 `json:"ip_range,omitempty"`
	Line                    *GeoLineAggregation                 `json:"line,omitempty"`
	MatrixStats             *MatrixStatsAggregation             `json:"matrix_stats,omitempty"`
	Max                     *MaxAggregation                     `json:"max,omitempty"`
	MaxBucket               *MaxBucketAggregation               `json:"max_bucket,omitempty"`
	MedianAbsoluteDeviation *MedianAbsoluteDeviationAggregation `json:"median_absolute_deviation,omitempty"`
	Meta                    *Metadata                           `json:"meta,omitempty"`
	Min                     *MinAggregation                     `json:"min,omitempty"`
	MinBucket               *MinBucketAggregation               `json:"min_bucket,omitempty"`
	Missing                 *MissingAggregation                 `json:"missing,omitempty"`
	MovingAvg               *MovingAverageAggregation           `json:"moving_avg,omitempty"`
	MovingFn                *MovingFunctionAggregation          `json:"moving_fn,omitempty"`
	MovingPercentiles       *MovingPercentilesAggregation       `json:"moving_percentiles,omitempty"`
	MultiTerms              *MultiTermsAggregation              `json:"multi_terms,omitempty"`
	Nested                  *NestedAggregation                  `json:"nested,omitempty"`
	Normalize               *NormalizeAggregation               `json:"normalize,omitempty"`
	Parent                  *ParentAggregation                  `json:"parent,omitempty"`
	PercentileRanks         *PercentileRanksAggregation         `json:"percentile_ranks,omitempty"`
	Percentiles             *PercentilesAggregation             `json:"percentiles,omitempty"`
	PercentilesBucket       *PercentilesBucketAggregation       `json:"percentiles_bucket,omitempty"`
	Range                   *RangeAggregation                   `json:"range,omitempty"`
	RareTerms               *RareTermsAggregation               `json:"rare_terms,omitempty"`
	Rate                    *RateAggregation                    `json:"rate,omitempty"`
	ReverseNested           *ReverseNestedAggregation           `json:"reverse_nested,omitempty"`
	Sampler                 *SamplerAggregation                 `json:"sampler,omitempty"`
	ScriptedMetric          *ScriptedMetricAggregation          `json:"scripted_metric,omitempty"`
	SerialDiff              *SerialDifferencingAggregation      `json:"serial_diff,omitempty"`
	SignificantTerms        *SignificantTermsAggregation        `json:"significant_terms,omitempty"`
	SignificantText         *SignificantTextAggregation         `json:"significant_text,omitempty"`
	Stats                   *StatsAggregation                   `json:"stats,omitempty"`
	StatsBucket             *StatsBucketAggregation             `json:"stats_bucket,omitempty"`
	StringStats             *StringStatsAggregation             `json:"string_stats,omitempty"`
	Sum                     *SumAggregation                     `json:"sum,omitempty"`
	SumBucket               *SumBucketAggregation               `json:"sum_bucket,omitempty"`
	TTest                   *TTestAggregation                   `json:"t_test,omitempty"`
	Terms                   *TermsAggregation                   `json:"terms,omitempty"`
	TopHits                 *TopHitsAggregation                 `json:"top_hits,omitempty"`
	TopMetrics              *TopMetricsAggregation              `json:"top_metrics,omitempty"`
	ValueCount              *ValueCountAggregation              `json:"value_count,omitempty"`
	VariableWidthHistogram  *VariableWidthHistogramAggregation  `json:"variable_width_histogram,omitempty"`
	WeightedAvg             *WeightedAverageAggregation         `json:"weighted_avg,omitempty"`
}

// AggregationContainerBuilder holds AggregationContainer struct and provides a builder API.
type AggregationContainerBuilder struct {
	v *AggregationContainer
}

// NewAggregationContainer provides a builder for the AggregationContainer struct.
func NewAggregationContainer() *AggregationContainerBuilder {
	r := AggregationContainerBuilder{
		&AggregationContainer{
			Aggregations: make(map[string]AggregationContainer, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AggregationContainer struct
func (rb *AggregationContainerBuilder) Build() AggregationContainer {
	return *rb.v
}

// AdjacencyMatrix set the AdjacencyMatrix property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) AdjacencyMatrix(adjacencymatrix AdjacencyMatrixAggregation) *AggregationContainerBuilder {
	rb.v.AdjacencyMatrix = &adjacencymatrix
	return rb
}

// Aggregations Sub-aggregations for this aggregation. Only applies to bucket aggregations.
func (rb *AggregationContainerBuilder) Aggregations(value map[string]AggregationContainer) *AggregationContainerBuilder {
	rb.v.Aggregations = value
	return rb
}

// AutoDateHistogram set the AutoDateHistogram property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) AutoDateHistogram(autodatehistogram AutoDateHistogramAggregation) *AggregationContainerBuilder {
	rb.v.AutoDateHistogram = &autodatehistogram
	return rb
}

// Avg set the Avg property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Avg(avg AverageAggregation) *AggregationContainerBuilder {
	rb.v.Avg = &avg
	return rb
}

// AvgBucket set the AvgBucket property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) AvgBucket(avgbucket AverageBucketAggregation) *AggregationContainerBuilder {
	rb.v.AvgBucket = &avgbucket
	return rb
}

// Boxplot set the Boxplot property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Boxplot(boxplot BoxplotAggregation) *AggregationContainerBuilder {
	rb.v.Boxplot = &boxplot
	return rb
}

// BucketCorrelation set the BucketCorrelation property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) BucketCorrelation(bucketcorrelation BucketCorrelationAggregation) *AggregationContainerBuilder {
	rb.v.BucketCorrelation = &bucketcorrelation
	return rb
}

// BucketCountKsTest set the BucketCountKsTest property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) BucketCountKsTest(bucketcountkstest BucketKsAggregation) *AggregationContainerBuilder {
	rb.v.BucketCountKsTest = &bucketcountkstest
	return rb
}

// BucketScript set the BucketScript property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) BucketScript(bucketscript BucketScriptAggregation) *AggregationContainerBuilder {
	rb.v.BucketScript = &bucketscript
	return rb
}

// BucketSelector set the BucketSelector property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) BucketSelector(bucketselector BucketSelectorAggregation) *AggregationContainerBuilder {
	rb.v.BucketSelector = &bucketselector
	return rb
}

// BucketSort set the BucketSort property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) BucketSort(bucketsort BucketSortAggregation) *AggregationContainerBuilder {
	rb.v.BucketSort = &bucketsort
	return rb
}

// Cardinality set the Cardinality property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Cardinality(cardinality CardinalityAggregation) *AggregationContainerBuilder {
	rb.v.Cardinality = &cardinality
	return rb
}

// CategorizeText set the CategorizeText property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) CategorizeText(categorizetext CategorizeTextAggregation) *AggregationContainerBuilder {
	rb.v.CategorizeText = &categorizetext
	return rb
}

// Children set the Children property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Children(children ChildrenAggregation) *AggregationContainerBuilder {
	rb.v.Children = &children
	return rb
}

// Composite set the Composite property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Composite(composite CompositeAggregation) *AggregationContainerBuilder {
	rb.v.Composite = &composite
	return rb
}

// CumulativeCardinality set the CumulativeCardinality property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) CumulativeCardinality(cumulativecardinality CumulativeCardinalityAggregation) *AggregationContainerBuilder {
	rb.v.CumulativeCardinality = &cumulativecardinality
	return rb
}

// CumulativeSum set the CumulativeSum property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) CumulativeSum(cumulativesum CumulativeSumAggregation) *AggregationContainerBuilder {
	rb.v.CumulativeSum = &cumulativesum
	return rb
}

// DateHistogram set the DateHistogram property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) DateHistogram(datehistogram DateHistogramAggregation) *AggregationContainerBuilder {
	rb.v.DateHistogram = &datehistogram
	return rb
}

// DateRange set the DateRange property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) DateRange(daterange DateRangeAggregation) *AggregationContainerBuilder {
	rb.v.DateRange = &daterange
	return rb
}

// Derivative set the Derivative property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Derivative(derivative DerivativeAggregation) *AggregationContainerBuilder {
	rb.v.Derivative = &derivative
	return rb
}

// DiversifiedSampler set the DiversifiedSampler property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) DiversifiedSampler(diversifiedsampler DiversifiedSamplerAggregation) *AggregationContainerBuilder {
	rb.v.DiversifiedSampler = &diversifiedsampler
	return rb
}

// ExtendedStats set the ExtendedStats property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) ExtendedStats(extendedstats ExtendedStatsAggregation) *AggregationContainerBuilder {
	rb.v.ExtendedStats = &extendedstats
	return rb
}

// ExtendedStatsBucket set the ExtendedStatsBucket property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) ExtendedStatsBucket(extendedstatsbucket ExtendedStatsBucketAggregation) *AggregationContainerBuilder {
	rb.v.ExtendedStatsBucket = &extendedstatsbucket
	return rb
}

// Filter set the Filter property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Filter(filter QueryContainer) *AggregationContainerBuilder {
	rb.v.Filter = &filter
	return rb
}

// Filters set the Filters property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Filters(filters FiltersAggregation) *AggregationContainerBuilder {
	rb.v.Filters = &filters
	return rb
}

// GeoBounds set the GeoBounds property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) GeoBounds(geobounds GeoBoundsAggregation) *AggregationContainerBuilder {
	rb.v.GeoBounds = &geobounds
	return rb
}

// GeoCentroid set the GeoCentroid property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) GeoCentroid(geocentroid GeoCentroidAggregation) *AggregationContainerBuilder {
	rb.v.GeoCentroid = &geocentroid
	return rb
}

// GeoDistance set the GeoDistance property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) GeoDistance(geodistance GeoDistanceAggregation) *AggregationContainerBuilder {
	rb.v.GeoDistance = &geodistance
	return rb
}

// GeoLine set the GeoLine property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) GeoLine(geoline GeoLineAggregation) *AggregationContainerBuilder {
	rb.v.GeoLine = &geoline
	return rb
}

// GeohashGrid set the GeohashGrid property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) GeohashGrid(geohashgrid GeoHashGridAggregation) *AggregationContainerBuilder {
	rb.v.GeohashGrid = &geohashgrid
	return rb
}

// GeohexGrid set the GeohexGrid property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) GeohexGrid(geohexgrid GeohexGridAggregation) *AggregationContainerBuilder {
	rb.v.GeohexGrid = &geohexgrid
	return rb
}

// GeotileGrid set the GeotileGrid property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) GeotileGrid(geotilegrid GeoTileGridAggregation) *AggregationContainerBuilder {
	rb.v.GeotileGrid = &geotilegrid
	return rb
}

// Global set the Global property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Global(global GlobalAggregation) *AggregationContainerBuilder {
	rb.v.Global = &global
	return rb
}

// Histogram set the Histogram property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Histogram(histogram HistogramAggregation) *AggregationContainerBuilder {
	rb.v.Histogram = &histogram
	return rb
}

// Inference set the Inference property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Inference(inference InferenceAggregation) *AggregationContainerBuilder {
	rb.v.Inference = &inference
	return rb
}

// IpRange set the IpRange property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) IpRange(iprange IpRangeAggregation) *AggregationContainerBuilder {
	rb.v.IpRange = &iprange
	return rb
}

// Line set the Line property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Line(line GeoLineAggregation) *AggregationContainerBuilder {
	rb.v.Line = &line
	return rb
}

// MatrixStats set the MatrixStats property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) MatrixStats(matrixstats MatrixStatsAggregation) *AggregationContainerBuilder {
	rb.v.MatrixStats = &matrixstats
	return rb
}

// Max set the Max property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Max(max MaxAggregation) *AggregationContainerBuilder {
	rb.v.Max = &max
	return rb
}

// MaxBucket set the MaxBucket property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) MaxBucket(maxbucket MaxBucketAggregation) *AggregationContainerBuilder {
	rb.v.MaxBucket = &maxbucket
	return rb
}

// MedianAbsoluteDeviation set the MedianAbsoluteDeviation property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) MedianAbsoluteDeviation(medianabsolutedeviation MedianAbsoluteDeviationAggregation) *AggregationContainerBuilder {
	rb.v.MedianAbsoluteDeviation = &medianabsolutedeviation
	return rb
}

// Meta set the Meta property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Meta(meta Metadata) *AggregationContainerBuilder {
	rb.v.Meta = &meta
	return rb
}

// Min set the Min property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Min(min MinAggregation) *AggregationContainerBuilder {
	rb.v.Min = &min
	return rb
}

// MinBucket set the MinBucket property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) MinBucket(minbucket MinBucketAggregation) *AggregationContainerBuilder {
	rb.v.MinBucket = &minbucket
	return rb
}

// Missing set the Missing property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Missing(missing MissingAggregation) *AggregationContainerBuilder {
	rb.v.Missing = &missing
	return rb
}

// MovingAvg set the MovingAvg property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) MovingAvg(movingavg MovingAverageAggregation) *AggregationContainerBuilder {
	rb.v.MovingAvg = &movingavg
	return rb
}

// MovingFn set the MovingFn property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) MovingFn(movingfn MovingFunctionAggregation) *AggregationContainerBuilder {
	rb.v.MovingFn = &movingfn
	return rb
}

// MovingPercentiles set the MovingPercentiles property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) MovingPercentiles(movingpercentiles MovingPercentilesAggregation) *AggregationContainerBuilder {
	rb.v.MovingPercentiles = &movingpercentiles
	return rb
}

// MultiTerms set the MultiTerms property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) MultiTerms(multiterms MultiTermsAggregation) *AggregationContainerBuilder {
	rb.v.MultiTerms = &multiterms
	return rb
}

// Nested set the Nested property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Nested(nested NestedAggregation) *AggregationContainerBuilder {
	rb.v.Nested = &nested
	return rb
}

// Normalize set the Normalize property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Normalize(normalize NormalizeAggregation) *AggregationContainerBuilder {
	rb.v.Normalize = &normalize
	return rb
}

// Parent set the Parent property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Parent(parent ParentAggregation) *AggregationContainerBuilder {
	rb.v.Parent = &parent
	return rb
}

// PercentileRanks set the PercentileRanks property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) PercentileRanks(percentileranks PercentileRanksAggregation) *AggregationContainerBuilder {
	rb.v.PercentileRanks = &percentileranks
	return rb
}

// Percentiles set the Percentiles property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Percentiles(percentiles PercentilesAggregation) *AggregationContainerBuilder {
	rb.v.Percentiles = &percentiles
	return rb
}

// PercentilesBucket set the PercentilesBucket property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) PercentilesBucket(percentilesbucket PercentilesBucketAggregation) *AggregationContainerBuilder {
	rb.v.PercentilesBucket = &percentilesbucket
	return rb
}

// Range set the Range property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Range_(range_ RangeAggregation) *AggregationContainerBuilder {
	rb.v.Range = &range_
	return rb
}

// RareTerms set the RareTerms property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) RareTerms(rareterms RareTermsAggregation) *AggregationContainerBuilder {
	rb.v.RareTerms = &rareterms
	return rb
}

// Rate set the Rate property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Rate(rate RateAggregation) *AggregationContainerBuilder {
	rb.v.Rate = &rate
	return rb
}

// ReverseNested set the ReverseNested property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) ReverseNested(reversenested ReverseNestedAggregation) *AggregationContainerBuilder {
	rb.v.ReverseNested = &reversenested
	return rb
}

// Sampler set the Sampler property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Sampler(sampler SamplerAggregation) *AggregationContainerBuilder {
	rb.v.Sampler = &sampler
	return rb
}

// ScriptedMetric set the ScriptedMetric property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) ScriptedMetric(scriptedmetric ScriptedMetricAggregation) *AggregationContainerBuilder {
	rb.v.ScriptedMetric = &scriptedmetric
	return rb
}

// SerialDiff set the SerialDiff property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) SerialDiff(serialdiff SerialDifferencingAggregation) *AggregationContainerBuilder {
	rb.v.SerialDiff = &serialdiff
	return rb
}

// SignificantTerms set the SignificantTerms property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) SignificantTerms(significantterms SignificantTermsAggregation) *AggregationContainerBuilder {
	rb.v.SignificantTerms = &significantterms
	return rb
}

// SignificantText set the SignificantText property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) SignificantText(significanttext SignificantTextAggregation) *AggregationContainerBuilder {
	rb.v.SignificantText = &significanttext
	return rb
}

// Stats set the Stats property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Stats(stats StatsAggregation) *AggregationContainerBuilder {
	rb.v.Stats = &stats
	return rb
}

// StatsBucket set the StatsBucket property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) StatsBucket(statsbucket StatsBucketAggregation) *AggregationContainerBuilder {
	rb.v.StatsBucket = &statsbucket
	return rb
}

// StringStats set the StringStats property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) StringStats(stringstats StringStatsAggregation) *AggregationContainerBuilder {
	rb.v.StringStats = &stringstats
	return rb
}

// Sum set the Sum property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Sum(sum SumAggregation) *AggregationContainerBuilder {
	rb.v.Sum = &sum
	return rb
}

// SumBucket set the SumBucket property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) SumBucket(sumbucket SumBucketAggregation) *AggregationContainerBuilder {
	rb.v.SumBucket = &sumbucket
	return rb
}

// TTest set the TTest property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) TTest(ttest TTestAggregation) *AggregationContainerBuilder {
	rb.v.TTest = &ttest
	return rb
}

// Terms set the Terms property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) Terms(terms TermsAggregation) *AggregationContainerBuilder {
	rb.v.Terms = &terms
	return rb
}

// TopHits set the TopHits property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) TopHits(tophits TopHitsAggregation) *AggregationContainerBuilder {
	rb.v.TopHits = &tophits
	return rb
}

// TopMetrics set the TopMetrics property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) TopMetrics(topmetrics TopMetricsAggregation) *AggregationContainerBuilder {
	rb.v.TopMetrics = &topmetrics
	return rb
}

// ValueCount set the ValueCount property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) ValueCount(valuecount ValueCountAggregation) *AggregationContainerBuilder {
	rb.v.ValueCount = &valuecount
	return rb
}

// VariableWidthHistogram set the VariableWidthHistogram property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) VariableWidthHistogram(variablewidthhistogram VariableWidthHistogramAggregation) *AggregationContainerBuilder {
	rb.v.VariableWidthHistogram = &variablewidthhistogram
	return rb
}

// WeightedAvg set the WeightedAvg property for AggregationContainerBuilder.
func (rb *AggregationContainerBuilder) WeightedAvg(weightedavg WeightedAverageAggregation) *AggregationContainerBuilder {
	rb.v.WeightedAvg = &weightedavg
	return rb
}
