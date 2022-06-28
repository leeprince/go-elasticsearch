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

// InferenceProcessor type.
type InferenceProcessor struct {
	FieldMap        map[Field]interface{} `json:"field_map,omitempty"`
	If              *string               `json:"if,omitempty"`
	IgnoreFailure   *bool                 `json:"ignore_failure,omitempty"`
	InferenceConfig *InferenceConfig      `json:"inference_config,omitempty"`
	ModelId         Id                    `json:"model_id"`
	OnFailure       []ProcessorContainer  `json:"on_failure,omitempty"`
	Tag             *string               `json:"tag,omitempty"`
	TargetField     Field                 `json:"target_field"`
}

// InferenceProcessorBuilder holds InferenceProcessor struct and provides a builder API.
type InferenceProcessorBuilder struct {
	v *InferenceProcessor
}

// NewInferenceProcessor provides a builder for the InferenceProcessor struct.
func NewInferenceProcessor() *InferenceProcessorBuilder {
	r := InferenceProcessorBuilder{
		&InferenceProcessor{
			FieldMap: make(map[Field]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InferenceProcessor struct
func (rb *InferenceProcessorBuilder) Build() InferenceProcessor {
	return *rb.v
}

// FieldMap set the FieldMap property for InferenceProcessorBuilder.
func (rb *InferenceProcessorBuilder) FieldMap(value map[Field]interface{}) *InferenceProcessorBuilder {
	rb.v.FieldMap = value
	return rb
}

// If set the If property for InferenceProcessorBuilder.
func (rb *InferenceProcessorBuilder) If_(if_ string) *InferenceProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for InferenceProcessorBuilder.
func (rb *InferenceProcessorBuilder) IgnoreFailure(ignorefailure bool) *InferenceProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// InferenceConfig set the InferenceConfig property for InferenceProcessorBuilder.
func (rb *InferenceProcessorBuilder) InferenceConfig(inferenceconfig InferenceConfig) *InferenceProcessorBuilder {
	rb.v.InferenceConfig = &inferenceconfig
	return rb
}

// ModelId set the ModelId property for InferenceProcessorBuilder.
func (rb *InferenceProcessorBuilder) ModelId(modelid Id) *InferenceProcessorBuilder {
	rb.v.ModelId = modelid
	return rb
}

// OnFailure set the OnFailure property for InferenceProcessorBuilder.
func (rb *InferenceProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *InferenceProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for InferenceProcessorBuilder.
func (rb *InferenceProcessorBuilder) Tag(tag string) *InferenceProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for InferenceProcessorBuilder.
func (rb *InferenceProcessorBuilder) TargetField(targetfield Field) *InferenceProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
