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

// AggregationBreakdown type.
type AggregationBreakdown struct {
	BuildAggregation        int64  `json:"build_aggregation"`
	BuildAggregationCount   int64  `json:"build_aggregation_count"`
	BuildLeafCollector      int64  `json:"build_leaf_collector"`
	BuildLeafCollectorCount int64  `json:"build_leaf_collector_count"`
	Collect                 int64  `json:"collect"`
	CollectCount            int64  `json:"collect_count"`
	Initialize              int64  `json:"initialize"`
	InitializeCount         int64  `json:"initialize_count"`
	PostCollection          *int64 `json:"post_collection,omitempty"`
	PostCollectionCount     *int64 `json:"post_collection_count,omitempty"`
	Reduce                  int64  `json:"reduce"`
	ReduceCount             int64  `json:"reduce_count"`
}

// AggregationBreakdownBuilder holds AggregationBreakdown struct and provides a builder API.
type AggregationBreakdownBuilder struct {
	v *AggregationBreakdown
}

// NewAggregationBreakdown provides a builder for the AggregationBreakdown struct.
func NewAggregationBreakdown() *AggregationBreakdownBuilder {
	r := AggregationBreakdownBuilder{
		&AggregationBreakdown{},
	}

	return &r
}

// Build finalize the chain and returns the AggregationBreakdown struct
func (rb *AggregationBreakdownBuilder) Build() AggregationBreakdown {
	return *rb.v
}

// BuildAggregation set the BuildAggregation property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) BuildAggregation(buildaggregation int64) *AggregationBreakdownBuilder {
	rb.v.BuildAggregation = buildaggregation
	return rb
}

// BuildAggregationCount set the BuildAggregationCount property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) BuildAggregationCount(buildaggregationcount int64) *AggregationBreakdownBuilder {
	rb.v.BuildAggregationCount = buildaggregationcount
	return rb
}

// BuildLeafCollector set the BuildLeafCollector property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) BuildLeafCollector(buildleafcollector int64) *AggregationBreakdownBuilder {
	rb.v.BuildLeafCollector = buildleafcollector
	return rb
}

// BuildLeafCollectorCount set the BuildLeafCollectorCount property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) BuildLeafCollectorCount(buildleafcollectorcount int64) *AggregationBreakdownBuilder {
	rb.v.BuildLeafCollectorCount = buildleafcollectorcount
	return rb
}

// Collect set the Collect property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) Collect(collect int64) *AggregationBreakdownBuilder {
	rb.v.Collect = collect
	return rb
}

// CollectCount set the CollectCount property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) CollectCount(collectcount int64) *AggregationBreakdownBuilder {
	rb.v.CollectCount = collectcount
	return rb
}

// Initialize set the Initialize property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) Initialize(initialize int64) *AggregationBreakdownBuilder {
	rb.v.Initialize = initialize
	return rb
}

// InitializeCount set the InitializeCount property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) InitializeCount(initializecount int64) *AggregationBreakdownBuilder {
	rb.v.InitializeCount = initializecount
	return rb
}

// PostCollection set the PostCollection property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) PostCollection(postcollection int64) *AggregationBreakdownBuilder {
	rb.v.PostCollection = &postcollection
	return rb
}

// PostCollectionCount set the PostCollectionCount property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) PostCollectionCount(postcollectioncount int64) *AggregationBreakdownBuilder {
	rb.v.PostCollectionCount = &postcollectioncount
	return rb
}

// Reduce set the Reduce property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) Reduce(reduce int64) *AggregationBreakdownBuilder {
	rb.v.Reduce = reduce
	return rb
}

// ReduceCount set the ReduceCount property for AggregationBreakdownBuilder.
func (rb *AggregationBreakdownBuilder) ReduceCount(reducecount int64) *AggregationBreakdownBuilder {
	rb.v.ReduceCount = reducecount
	return rb
}
