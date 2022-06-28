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

// SignificantTextAggregation type.
type SignificantTextAggregation struct {
	BackgroundFilter    *QueryContainer                                              `json:"background_filter,omitempty"`
	ChiSquare           *ChiSquareHeuristic                                          `json:"chi_square,omitempty"`
	Exclude             *TermsExclude                                                `json:"exclude,omitempty"`
	ExecutionHint       *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field               *Field                                                       `json:"field,omitempty"`
	FilterDuplicateText *bool                                                        `json:"filter_duplicate_text,omitempty"`
	Gnd                 *GoogleNormalizedDistanceHeuristic                           `json:"gnd,omitempty"`
	Include             []string                                                     `json:"include,omitempty"`
	Jlh                 *EmptyObject                                                 `json:"jlh,omitempty"`
	Meta                *Metadata                                                    `json:"meta,omitempty"`
	MinDocCount         *int64                                                       `json:"min_doc_count,omitempty"`
	MutualInformation   *MutualInformationHeuristic                                  `json:"mutual_information,omitempty"`
	Name                *string                                                      `json:"name,omitempty"`
	Percentage          *PercentageScoreHeuristic                                    `json:"percentage,omitempty"`
	ScriptHeuristic     *ScriptedHeuristic                                           `json:"script_heuristic,omitempty"`
	ShardMinDocCount    *int64                                                       `json:"shard_min_doc_count,omitempty"`
	ShardSize           *int                                                         `json:"shard_size,omitempty"`
	Size                *int                                                         `json:"size,omitempty"`
	SourceFields        *Fields                                                      `json:"source_fields,omitempty"`
}

// SignificantTextAggregationBuilder holds SignificantTextAggregation struct and provides a builder API.
type SignificantTextAggregationBuilder struct {
	v *SignificantTextAggregation
}

// NewSignificantTextAggregation provides a builder for the SignificantTextAggregation struct.
func NewSignificantTextAggregation() *SignificantTextAggregationBuilder {
	r := SignificantTextAggregationBuilder{
		&SignificantTextAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the SignificantTextAggregation struct
func (rb *SignificantTextAggregationBuilder) Build() SignificantTextAggregation {
	return *rb.v
}

// BackgroundFilter set the BackgroundFilter property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) BackgroundFilter(backgroundfilter QueryContainer) *SignificantTextAggregationBuilder {
	rb.v.BackgroundFilter = &backgroundfilter
	return rb
}

// ChiSquare set the ChiSquare property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) ChiSquare(chisquare ChiSquareHeuristic) *SignificantTextAggregationBuilder {
	rb.v.ChiSquare = &chisquare
	return rb
}

// Exclude set the Exclude property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Exclude(exclude TermsExclude) *SignificantTextAggregationBuilder {
	rb.v.Exclude = &exclude
	return rb
}

// ExecutionHint set the ExecutionHint property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *SignificantTextAggregationBuilder {
	rb.v.ExecutionHint = &executionhint
	return rb
}

// Field set the Field property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Field(field Field) *SignificantTextAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// FilterDuplicateText set the FilterDuplicateText property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) FilterDuplicateText(filterduplicatetext bool) *SignificantTextAggregationBuilder {
	rb.v.FilterDuplicateText = &filterduplicatetext
	return rb
}

// Gnd set the Gnd property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Gnd(gnd GoogleNormalizedDistanceHeuristic) *SignificantTextAggregationBuilder {
	rb.v.Gnd = &gnd
	return rb
}

// Include set the Include property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Include(arg []string) *SignificantTextAggregationBuilder {
	rb.v.Include = arg
	return rb
}

// Jlh set the Jlh property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Jlh(jlh EmptyObject) *SignificantTextAggregationBuilder {
	rb.v.Jlh = &jlh
	return rb
}

// Meta set the Meta property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Meta(meta Metadata) *SignificantTextAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// MinDocCount set the MinDocCount property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) MinDocCount(mindoccount int64) *SignificantTextAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

// MutualInformation set the MutualInformation property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) MutualInformation(mutualinformation MutualInformationHeuristic) *SignificantTextAggregationBuilder {
	rb.v.MutualInformation = &mutualinformation
	return rb
}

// Name set the Name property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Name(name string) *SignificantTextAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Percentage set the Percentage property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Percentage(percentage PercentageScoreHeuristic) *SignificantTextAggregationBuilder {
	rb.v.Percentage = &percentage
	return rb
}

// ScriptHeuristic set the ScriptHeuristic property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) ScriptHeuristic(scriptheuristic ScriptedHeuristic) *SignificantTextAggregationBuilder {
	rb.v.ScriptHeuristic = &scriptheuristic
	return rb
}

// ShardMinDocCount set the ShardMinDocCount property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) ShardMinDocCount(shardmindoccount int64) *SignificantTextAggregationBuilder {
	rb.v.ShardMinDocCount = &shardmindoccount
	return rb
}

// ShardSize set the ShardSize property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) ShardSize(shardsize int) *SignificantTextAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

// Size set the Size property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) Size(size int) *SignificantTextAggregationBuilder {
	rb.v.Size = &size
	return rb
}

// SourceFields set the SourceFields property for SignificantTextAggregationBuilder.
func (rb *SignificantTextAggregationBuilder) SourceFields(sourcefields Fields) *SignificantTextAggregationBuilder {
	rb.v.SourceFields = &sourcefields
	return rb
}
