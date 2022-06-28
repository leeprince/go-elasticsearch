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

// ChainInput type.
type ChainInput struct {
	Inputs []InputContainer `json:"inputs"`
}

// ChainInputBuilder holds ChainInput struct and provides a builder API.
type ChainInputBuilder struct {
	v *ChainInput
}

// NewChainInput provides a builder for the ChainInput struct.
func NewChainInput() *ChainInputBuilder {
	r := ChainInputBuilder{
		&ChainInput{},
	}

	return &r
}

// Build finalize the chain and returns the ChainInput struct
func (rb *ChainInputBuilder) Build() ChainInput {
	return *rb.v
}

// Inputs set the Inputs property for ChainInputBuilder.
func (rb *ChainInputBuilder) Inputs(inputs ...InputContainer) *ChainInputBuilder {
	rb.v.Inputs = append(rb.v.Inputs, inputs...)
	return rb
}
