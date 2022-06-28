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

// InferenceConfigCreateContainer type.
type InferenceConfigCreateContainer struct {
	// Classification Classification configuration for inference.
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`
	// FillMask Fill mask configuration for inference.
	FillMask *FillMaskInferenceOptions `json:"fill_mask,omitempty"`
	// Ner Named entity recognition configuration for inference.
	Ner *NerInferenceOptions `json:"ner,omitempty"`
	// PassThrough Pass through configuration for inference.
	PassThrough *PassThroughInferenceOptions `json:"pass_through,omitempty"`
	// QuestionAnswering Question answering configuration for inference.
	QuestionAnswering *QuestionAnsweringInferenceOptions `json:"question_answering,omitempty"`
	// Regression Regression configuration for inference.
	Regression *RegressionInferenceOptions `json:"regression,omitempty"`
	// TextClassification Text classification configuration for inference.
	TextClassification *TextClassificationInferenceOptions `json:"text_classification,omitempty"`
	// TextEmbedding Text embedding configuration for inference.
	TextEmbedding *TextEmbeddingInferenceOptions `json:"text_embedding,omitempty"`
	// ZeroShotClassification Zeroshot classification configuration for inference.
	ZeroShotClassification *ZeroShotClassificationInferenceOptions `json:"zero_shot_classification,omitempty"`
}

// InferenceConfigCreateContainerBuilder holds InferenceConfigCreateContainer struct and provides a builder API.
type InferenceConfigCreateContainerBuilder struct {
	v *InferenceConfigCreateContainer
}

// NewInferenceConfigCreateContainer provides a builder for the InferenceConfigCreateContainer struct.
func NewInferenceConfigCreateContainer() *InferenceConfigCreateContainerBuilder {
	r := InferenceConfigCreateContainerBuilder{
		&InferenceConfigCreateContainer{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceConfigCreateContainer struct
func (rb *InferenceConfigCreateContainerBuilder) Build() InferenceConfigCreateContainer {
	return *rb.v
}

// Classification Classification configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) Classification(classification ClassificationInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.Classification = &classification
	return rb
}

// FillMask Fill mask configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) FillMask(fillmask FillMaskInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.FillMask = &fillmask
	return rb
}

// Ner Named entity recognition configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) Ner(ner NerInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.Ner = &ner
	return rb
}

// PassThrough Pass through configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) PassThrough(passthrough PassThroughInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.PassThrough = &passthrough
	return rb
}

// QuestionAnswering Question answering configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) QuestionAnswering(questionanswering QuestionAnsweringInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.QuestionAnswering = &questionanswering
	return rb
}

// Regression Regression configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) Regression(regression RegressionInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.Regression = &regression
	return rb
}

// TextClassification Text classification configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) TextClassification(textclassification TextClassificationInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.TextClassification = &textclassification
	return rb
}

// TextEmbedding Text embedding configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) TextEmbedding(textembedding TextEmbeddingInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.TextEmbedding = &textembedding
	return rb
}

// ZeroShotClassification Zeroshot classification configuration for inference.
func (rb *InferenceConfigCreateContainerBuilder) ZeroShotClassification(zeroshotclassification ZeroShotClassificationInferenceOptions) *InferenceConfigCreateContainerBuilder {
	rb.v.ZeroShotClassification = &zeroshotclassification
	return rb
}
