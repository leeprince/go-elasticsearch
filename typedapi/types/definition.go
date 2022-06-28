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

// Definition type.
type Definition struct {
	// Preprocessors Collection of preprocessors
	Preprocessors []Preprocessor `json:"preprocessors,omitempty"`
	// TrainedModel The definition of the trained model.
	TrainedModel TrainedModel `json:"trained_model"`
}

// DefinitionBuilder holds Definition struct and provides a builder API.
type DefinitionBuilder struct {
	v *Definition
}

// NewDefinition provides a builder for the Definition struct.
func NewDefinition() *DefinitionBuilder {
	r := DefinitionBuilder{
		&Definition{},
	}

	return &r
}

// Build finalize the chain and returns the Definition struct
func (rb *DefinitionBuilder) Build() Definition {
	return *rb.v
}

// Preprocessors Collection of preprocessors
func (rb *DefinitionBuilder) Preprocessors(preprocessors ...Preprocessor) *DefinitionBuilder {
	rb.v.Preprocessors = append(rb.v.Preprocessors, preprocessors...)
	return rb
}

// TrainedModel The definition of the trained model.
func (rb *DefinitionBuilder) TrainedModel(trainedmodel TrainedModel) *DefinitionBuilder {
	rb.v.TrainedModel = trainedmodel
	return rb
}
