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

// TransformContainer type.
type TransformContainer struct {
	Chain  []TransformContainer `json:"chain,omitempty"`
	Script *ScriptTransform     `json:"script,omitempty"`
	Search *SearchTransform     `json:"search,omitempty"`
}

// TransformContainerBuilder holds TransformContainer struct and provides a builder API.
type TransformContainerBuilder struct {
	v *TransformContainer
}

// NewTransformContainer provides a builder for the TransformContainer struct.
func NewTransformContainer() *TransformContainerBuilder {
	r := TransformContainerBuilder{
		&TransformContainer{},
	}

	return &r
}

// Build finalize the chain and returns the TransformContainer struct
func (rb *TransformContainerBuilder) Build() TransformContainer {
	return *rb.v
}

// Chain set the Chain property for TransformContainerBuilder.
func (rb *TransformContainerBuilder) Chain(chain ...TransformContainer) *TransformContainerBuilder {
	rb.v.Chain = append(rb.v.Chain, chain...)
	return rb
}

// Script set the Script property for TransformContainerBuilder.
func (rb *TransformContainerBuilder) Script(script ScriptTransform) *TransformContainerBuilder {
	rb.v.Script = &script
	return rb
}

// Search set the Search property for TransformContainerBuilder.
func (rb *TransformContainerBuilder) Search(search SearchTransform) *TransformContainerBuilder {
	rb.v.Search = &search
	return rb
}
