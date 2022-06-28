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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sampleraggregationexecutionhint"
)

// DiversifiedSamplerAggregation type.
type DiversifiedSamplerAggregation struct {
	ExecutionHint   *sampleraggregationexecutionhint.SamplerAggregationExecutionHint `json:"execution_hint,omitempty"`
	Field           *Field                                                           `json:"field,omitempty"`
	MaxDocsPerValue *int                                                             `json:"max_docs_per_value,omitempty"`
	Meta            *Metadata                                                        `json:"meta,omitempty"`
	Name            *string                                                          `json:"name,omitempty"`
	Script          *Script                                                          `json:"script,omitempty"`
	ShardSize       *int                                                             `json:"shard_size,omitempty"`
}

// DiversifiedSamplerAggregationBuilder holds DiversifiedSamplerAggregation struct and provides a builder API.
type DiversifiedSamplerAggregationBuilder struct {
	v *DiversifiedSamplerAggregation
}

// NewDiversifiedSamplerAggregation provides a builder for the DiversifiedSamplerAggregation struct.
func NewDiversifiedSamplerAggregation() *DiversifiedSamplerAggregationBuilder {
	r := DiversifiedSamplerAggregationBuilder{
		&DiversifiedSamplerAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the DiversifiedSamplerAggregation struct
func (rb *DiversifiedSamplerAggregationBuilder) Build() DiversifiedSamplerAggregation {
	return *rb.v
}

// ExecutionHint set the ExecutionHint property for DiversifiedSamplerAggregationBuilder.
func (rb *DiversifiedSamplerAggregationBuilder) ExecutionHint(executionhint sampleraggregationexecutionhint.SamplerAggregationExecutionHint) *DiversifiedSamplerAggregationBuilder {
	rb.v.ExecutionHint = &executionhint
	return rb
}

// Field set the Field property for DiversifiedSamplerAggregationBuilder.
func (rb *DiversifiedSamplerAggregationBuilder) Field(field Field) *DiversifiedSamplerAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// MaxDocsPerValue set the MaxDocsPerValue property for DiversifiedSamplerAggregationBuilder.
func (rb *DiversifiedSamplerAggregationBuilder) MaxDocsPerValue(maxdocspervalue int) *DiversifiedSamplerAggregationBuilder {
	rb.v.MaxDocsPerValue = &maxdocspervalue
	return rb
}

// Meta set the Meta property for DiversifiedSamplerAggregationBuilder.
func (rb *DiversifiedSamplerAggregationBuilder) Meta(meta Metadata) *DiversifiedSamplerAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for DiversifiedSamplerAggregationBuilder.
func (rb *DiversifiedSamplerAggregationBuilder) Name(name string) *DiversifiedSamplerAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Script set the Script property for DiversifiedSamplerAggregationBuilder.
func (rb *DiversifiedSamplerAggregationBuilder) Script(script Script) *DiversifiedSamplerAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// ShardSize set the ShardSize property for DiversifiedSamplerAggregationBuilder.
func (rb *DiversifiedSamplerAggregationBuilder) ShardSize(shardsize int) *DiversifiedSamplerAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}
