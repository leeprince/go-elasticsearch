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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

// SignificantTermsAggregation type.
type SignificantTermsAggregation struct {
	BackgroundFilter  *QueryContainer                                              `json:"background_filter,omitempty"`
	ChiSquare         *ChiSquareHeuristic                                          `json:"chi_square,omitempty"`
	Exclude           *TermsExclude                                                `json:"exclude,omitempty"`
	ExecutionHint     *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field             *Field                                                       `json:"field,omitempty"`
	Gnd               *GoogleNormalizedDistanceHeuristic                           `json:"gnd,omitempty"`
	Include           *TermsInclude                                                `json:"include,omitempty"`
	Jlh               *EmptyObject                                                 `json:"jlh,omitempty"`
	Meta              *Metadata                                                    `json:"meta,omitempty"`
	MinDocCount       *int64                                                       `json:"min_doc_count,omitempty"`
	MutualInformation *MutualInformationHeuristic                                  `json:"mutual_information,omitempty"`
	Name              *string                                                      `json:"name,omitempty"`
	Percentage        *PercentageScoreHeuristic                                    `json:"percentage,omitempty"`
	ScriptHeuristic   *ScriptedHeuristic                                           `json:"script_heuristic,omitempty"`
	ShardMinDocCount  *int64                                                       `json:"shard_min_doc_count,omitempty"`
	ShardSize         *int                                                         `json:"shard_size,omitempty"`
	Size              *int                                                         `json:"size,omitempty"`
}

// SignificantTermsAggregationBuilder holds SignificantTermsAggregation struct and provides a builder API.
type SignificantTermsAggregationBuilder struct {
	v *SignificantTermsAggregation
}

// NewSignificantTermsAggregation provides a builder for the SignificantTermsAggregation struct.
func NewSignificantTermsAggregation() *SignificantTermsAggregationBuilder {
	r := SignificantTermsAggregationBuilder{
		&SignificantTermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SignificantTermsAggregation struct
func (rb *SignificantTermsAggregationBuilder) Build() SignificantTermsAggregation {
	return *rb.v
}

// BackgroundFilter set the BackgroundFilter property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) BackgroundFilter(backgroundfilter QueryContainer) *SignificantTermsAggregationBuilder {
	rb.v.BackgroundFilter = &backgroundfilter
	return rb
}

// ChiSquare set the ChiSquare property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) ChiSquare(chisquare ChiSquareHeuristic) *SignificantTermsAggregationBuilder {
	rb.v.ChiSquare = &chisquare
	return rb
}

// Exclude set the Exclude property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Exclude(exclude TermsExclude) *SignificantTermsAggregationBuilder {
	rb.v.Exclude = &exclude
	return rb
}

// ExecutionHint set the ExecutionHint property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *SignificantTermsAggregationBuilder {
	rb.v.ExecutionHint = &executionhint
	return rb
}

// Field set the Field property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Field(field Field) *SignificantTermsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Gnd set the Gnd property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Gnd(gnd GoogleNormalizedDistanceHeuristic) *SignificantTermsAggregationBuilder {
	rb.v.Gnd = &gnd
	return rb
}

// Include set the Include property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Include(include TermsInclude) *SignificantTermsAggregationBuilder {
	rb.v.Include = &include
	return rb
}

// Jlh set the Jlh property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Jlh(jlh EmptyObject) *SignificantTermsAggregationBuilder {
	rb.v.Jlh = &jlh
	return rb
}

// Meta set the Meta property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Meta(meta Metadata) *SignificantTermsAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// MinDocCount set the MinDocCount property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) MinDocCount(mindoccount int64) *SignificantTermsAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

// MutualInformation set the MutualInformation property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) MutualInformation(mutualinformation MutualInformationHeuristic) *SignificantTermsAggregationBuilder {
	rb.v.MutualInformation = &mutualinformation
	return rb
}

// Name set the Name property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Name(name string) *SignificantTermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Percentage set the Percentage property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Percentage(percentage PercentageScoreHeuristic) *SignificantTermsAggregationBuilder {
	rb.v.Percentage = &percentage
	return rb
}

// ScriptHeuristic set the ScriptHeuristic property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) ScriptHeuristic(scriptheuristic ScriptedHeuristic) *SignificantTermsAggregationBuilder {
	rb.v.ScriptHeuristic = &scriptheuristic
	return rb
}

// ShardMinDocCount set the ShardMinDocCount property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) ShardMinDocCount(shardmindoccount int64) *SignificantTermsAggregationBuilder {
	rb.v.ShardMinDocCount = &shardmindoccount
	return rb
}

// ShardSize set the ShardSize property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) ShardSize(shardsize int) *SignificantTermsAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

// Size set the Size property for SignificantTermsAggregationBuilder.
func (rb *SignificantTermsAggregationBuilder) Size(size int) *SignificantTermsAggregationBuilder {
	rb.v.Size = &size
	return rb
}
