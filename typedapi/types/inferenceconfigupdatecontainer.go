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

// InferenceConfigUpdateContainer type.
type InferenceConfigUpdateContainer struct {
	// Classification Classification configuration for inference.
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`
	// FillMask Fill mask configuration for inference.
	FillMask *FillMaskInferenceUpdateOptions `json:"fill_mask,omitempty"`
	// Ner Named entity recognition configuration for inference.
	Ner *NerInferenceUpdateOptions `json:"ner,omitempty"`
	// PassThrough Pass through configuration for inference.
	PassThrough *PassThroughInferenceUpdateOptions `json:"pass_through,omitempty"`
	// QuestionAnswering Question answering configuration for inference
	QuestionAnswering *QuestionAnsweringInferenceUpdateOptions `json:"question_answering,omitempty"`
	// Regression Regression configuration for inference.
	Regression *RegressionInferenceOptions `json:"regression,omitempty"`
	// TextClassification Text classification configuration for inference.
	TextClassification *TextClassificationInferenceUpdateOptions `json:"text_classification,omitempty"`
	// TextEmbedding Text embedding configuration for inference.
	TextEmbedding *TextEmbeddingInferenceUpdateOptions `json:"text_embedding,omitempty"`
	// ZeroShotClassification Zeroshot classification configuration for inference.
	ZeroShotClassification *ZeroShotClassificationInferenceUpdateOptions `json:"zero_shot_classification,omitempty"`
}

// InferenceConfigUpdateContainerBuilder holds InferenceConfigUpdateContainer struct and provides a builder API.
type InferenceConfigUpdateContainerBuilder struct {
	v *InferenceConfigUpdateContainer
}

// NewInferenceConfigUpdateContainer provides a builder for the InferenceConfigUpdateContainer struct.
func NewInferenceConfigUpdateContainer() *InferenceConfigUpdateContainerBuilder {
	r := InferenceConfigUpdateContainerBuilder{
		&InferenceConfigUpdateContainer{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceConfigUpdateContainer struct
func (rb *InferenceConfigUpdateContainerBuilder) Build() InferenceConfigUpdateContainer {
	return *rb.v
}

// Classification Classification configuration for inference.
func (rb *InferenceConfigUpdateContainerBuilder) Classification(classification ClassificationInferenceOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.Classification = &classification
	return rb
}

// FillMask Fill mask configuration for inference.
func (rb *InferenceConfigUpdateContainerBuilder) FillMask(fillmask FillMaskInferenceUpdateOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.FillMask = &fillmask
	return rb
}

// Ner Named entity recognition configuration for inference.
func (rb *InferenceConfigUpdateContainerBuilder) Ner(ner NerInferenceUpdateOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.Ner = &ner
	return rb
}

// PassThrough Pass through configuration for inference.
func (rb *InferenceConfigUpdateContainerBuilder) PassThrough(passthrough PassThroughInferenceUpdateOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.PassThrough = &passthrough
	return rb
}

// QuestionAnswering Question answering configuration for inference
func (rb *InferenceConfigUpdateContainerBuilder) QuestionAnswering(questionanswering QuestionAnsweringInferenceUpdateOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.QuestionAnswering = &questionanswering
	return rb
}

// Regression Regression configuration for inference.
func (rb *InferenceConfigUpdateContainerBuilder) Regression(regression RegressionInferenceOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.Regression = &regression
	return rb
}

// TextClassification Text classification configuration for inference.
func (rb *InferenceConfigUpdateContainerBuilder) TextClassification(textclassification TextClassificationInferenceUpdateOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.TextClassification = &textclassification
	return rb
}

// TextEmbedding Text embedding configuration for inference.
func (rb *InferenceConfigUpdateContainerBuilder) TextEmbedding(textembedding TextEmbeddingInferenceUpdateOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.TextEmbedding = &textembedding
	return rb
}

// ZeroShotClassification Zeroshot classification configuration for inference.
func (rb *InferenceConfigUpdateContainerBuilder) ZeroShotClassification(zeroshotclassification ZeroShotClassificationInferenceUpdateOptions) *InferenceConfigUpdateContainerBuilder {
	rb.v.ZeroShotClassification = &zeroshotclassification
	return rb
}
