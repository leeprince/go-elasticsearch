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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

// SortProcessor type.
type SortProcessor struct {
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Order         sortorder.SortOrder  `json:"order"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   Field                `json:"target_field"`
}

// SortProcessorBuilder holds SortProcessor struct and provides a builder API.
type SortProcessorBuilder struct {
	v *SortProcessor
}

// NewSortProcessor provides a builder for the SortProcessor struct.
func NewSortProcessor() *SortProcessorBuilder {
	r := SortProcessorBuilder{
		&SortProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the SortProcessor struct
func (rb *SortProcessorBuilder) Build() SortProcessor {
	return *rb.v
}

// Field set the Field property for SortProcessorBuilder.
func (rb *SortProcessorBuilder) Field(field Field) *SortProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for SortProcessorBuilder.
func (rb *SortProcessorBuilder) If_(if_ string) *SortProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for SortProcessorBuilder.
func (rb *SortProcessorBuilder) IgnoreFailure(ignorefailure bool) *SortProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// OnFailure set the OnFailure property for SortProcessorBuilder.
func (rb *SortProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *SortProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Order set the Order property for SortProcessorBuilder.
func (rb *SortProcessorBuilder) Order(order sortorder.SortOrder) *SortProcessorBuilder {
	rb.v.Order = order
	return rb
}

// Tag set the Tag property for SortProcessorBuilder.
func (rb *SortProcessorBuilder) Tag(tag string) *SortProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for SortProcessorBuilder.
func (rb *SortProcessorBuilder) TargetField(targetfield Field) *SortProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
