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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shapetype"
)

// CircleProcessor type.
type CircleProcessor struct {
	ErrorDistance float64              `json:"error_distance"`
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing bool                 `json:"ignore_missing"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	ShapeType     shapetype.ShapeType  `json:"shape_type"`
	Tag           *string              `json:"tag,omitempty"`
	TargetField   Field                `json:"target_field"`
}

// CircleProcessorBuilder holds CircleProcessor struct and provides a builder API.
type CircleProcessorBuilder struct {
	v *CircleProcessor
}

// NewCircleProcessor provides a builder for the CircleProcessor struct.
func NewCircleProcessor() *CircleProcessorBuilder {
	r := CircleProcessorBuilder{
		&CircleProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the CircleProcessor struct
func (rb *CircleProcessorBuilder) Build() CircleProcessor {
	return *rb.v
}

// ErrorDistance set the ErrorDistance property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) ErrorDistance(errordistance float64) *CircleProcessorBuilder {
	rb.v.ErrorDistance = errordistance
	return rb
}

// Field set the Field property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) Field(field Field) *CircleProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) If_(if_ string) *CircleProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) IgnoreFailure(ignorefailure bool) *CircleProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) IgnoreMissing(ignoremissing bool) *CircleProcessorBuilder {
	rb.v.IgnoreMissing = ignoremissing
	return rb
}

// OnFailure set the OnFailure property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *CircleProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// ShapeType set the ShapeType property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) ShapeType(shapetype shapetype.ShapeType) *CircleProcessorBuilder {
	rb.v.ShapeType = shapetype
	return rb
}

// Tag set the Tag property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) Tag(tag string) *CircleProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for CircleProcessorBuilder.
func (rb *CircleProcessorBuilder) TargetField(targetfield Field) *CircleProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
