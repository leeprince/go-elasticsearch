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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geoshaperelation"
)

// EnrichProcessor type.
type EnrichProcessor struct {
	Field         Field                              `json:"field"`
	If            *string                            `json:"if,omitempty"`
	IgnoreFailure *bool                              `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                              `json:"ignore_missing,omitempty"`
	MaxMatches    *int                               `json:"max_matches,omitempty"`
	OnFailure     []ProcessorContainer               `json:"on_failure,omitempty"`
	Override      *bool                              `json:"override,omitempty"`
	PolicyName    string                             `json:"policy_name"`
	ShapeRelation *geoshaperelation.GeoShapeRelation `json:"shape_relation,omitempty"`
	Tag           *string                            `json:"tag,omitempty"`
	TargetField   Field                              `json:"target_field"`
}

// EnrichProcessorBuilder holds EnrichProcessor struct and provides a builder API.
type EnrichProcessorBuilder struct {
	v *EnrichProcessor
}

// NewEnrichProcessor provides a builder for the EnrichProcessor struct.
func NewEnrichProcessor() *EnrichProcessorBuilder {
	r := EnrichProcessorBuilder{
		&EnrichProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the EnrichProcessor struct
func (rb *EnrichProcessorBuilder) Build() EnrichProcessor {
	return *rb.v
}

// Field set the Field property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) Field(field Field) *EnrichProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) If_(if_ string) *EnrichProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) IgnoreFailure(ignorefailure bool) *EnrichProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) IgnoreMissing(ignoremissing bool) *EnrichProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// MaxMatches set the MaxMatches property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) MaxMatches(maxmatches int) *EnrichProcessorBuilder {
	rb.v.MaxMatches = &maxmatches
	return rb
}

// OnFailure set the OnFailure property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *EnrichProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Override set the Override property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) Override(override bool) *EnrichProcessorBuilder {
	rb.v.Override = &override
	return rb
}

// PolicyName set the PolicyName property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) PolicyName(policyname string) *EnrichProcessorBuilder {
	rb.v.PolicyName = policyname
	return rb
}

// ShapeRelation set the ShapeRelation property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) ShapeRelation(shaperelation geoshaperelation.GeoShapeRelation) *EnrichProcessorBuilder {
	rb.v.ShapeRelation = &shaperelation
	return rb
}

// Tag set the Tag property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) Tag(tag string) *EnrichProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for EnrichProcessorBuilder.
func (rb *EnrichProcessorBuilder) TargetField(targetfield Field) *EnrichProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
