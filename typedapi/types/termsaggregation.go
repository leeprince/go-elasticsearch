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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationexecutionhint"
)

// TermsAggregation type.
type TermsAggregation struct {
	CollectMode           *termsaggregationcollectmode.TermsAggregationCollectMode     `json:"collect_mode,omitempty"`
	Exclude               *TermsExclude                                                `json:"exclude,omitempty"`
	ExecutionHint         *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field                 *Field                                                       `json:"field,omitempty"`
	Format                *string                                                      `json:"format,omitempty"`
	Include               *TermsInclude                                                `json:"include,omitempty"`
	Meta                  *Metadata                                                    `json:"meta,omitempty"`
	MinDocCount           *int                                                         `json:"min_doc_count,omitempty"`
	Missing               *Missing                                                     `json:"missing,omitempty"`
	MissingBucket         *bool                                                        `json:"missing_bucket,omitempty"`
	MissingOrder          *missingorder.MissingOrder                                   `json:"missing_order,omitempty"`
	Name                  *string                                                      `json:"name,omitempty"`
	Order                 *TermsAggregationOrder                                       `json:"order,omitempty"`
	Script                *Script                                                      `json:"script,omitempty"`
	ShardSize             *int                                                         `json:"shard_size,omitempty"`
	ShowTermDocCountError *bool                                                        `json:"show_term_doc_count_error,omitempty"`
	Size                  *int                                                         `json:"size,omitempty"`
	ValueType             *string                                                      `json:"value_type,omitempty"`
}

// TermsAggregationBuilder holds TermsAggregation struct and provides a builder API.
type TermsAggregationBuilder struct {
	v *TermsAggregation
}

// NewTermsAggregation provides a builder for the TermsAggregation struct.
func NewTermsAggregation() *TermsAggregationBuilder {
	r := TermsAggregationBuilder{
		&TermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the TermsAggregation struct
func (rb *TermsAggregationBuilder) Build() TermsAggregation {
	return *rb.v
}

// CollectMode set the CollectMode property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *TermsAggregationBuilder {
	rb.v.CollectMode = &collectmode
	return rb
}

// Exclude set the Exclude property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Exclude(exclude TermsExclude) *TermsAggregationBuilder {
	rb.v.Exclude = &exclude
	return rb
}

// ExecutionHint set the ExecutionHint property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) ExecutionHint(executionhint termsaggregationexecutionhint.TermsAggregationExecutionHint) *TermsAggregationBuilder {
	rb.v.ExecutionHint = &executionhint
	return rb
}

// Field set the Field property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Field(field Field) *TermsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Format(format string) *TermsAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Include set the Include property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Include(include TermsInclude) *TermsAggregationBuilder {
	rb.v.Include = &include
	return rb
}

// Meta set the Meta property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Meta(meta Metadata) *TermsAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// MinDocCount set the MinDocCount property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) MinDocCount(mindoccount int) *TermsAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

// Missing set the Missing property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Missing(missing Missing) *TermsAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// MissingBucket set the MissingBucket property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) MissingBucket(missingbucket bool) *TermsAggregationBuilder {
	rb.v.MissingBucket = &missingbucket
	return rb
}

// MissingOrder set the MissingOrder property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) MissingOrder(missingorder missingorder.MissingOrder) *TermsAggregationBuilder {
	rb.v.MissingOrder = &missingorder
	return rb
}

// Name set the Name property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Name(name string) *TermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Order set the Order property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Order(order TermsAggregationOrder) *TermsAggregationBuilder {
	rb.v.Order = &order
	return rb
}

// Script set the Script property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Script(script Script) *TermsAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// ShardSize set the ShardSize property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) ShardSize(shardsize int) *TermsAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

// ShowTermDocCountError set the ShowTermDocCountError property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) ShowTermDocCountError(showtermdoccounterror bool) *TermsAggregationBuilder {
	rb.v.ShowTermDocCountError = &showtermdoccounterror
	return rb
}

// Size set the Size property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) Size(size int) *TermsAggregationBuilder {
	rb.v.Size = &size
	return rb
}

// ValueType set the ValueType property for TermsAggregationBuilder.
func (rb *TermsAggregationBuilder) ValueType(valuetype string) *TermsAggregationBuilder {
	rb.v.ValueType = &valuetype
	return rb
}
