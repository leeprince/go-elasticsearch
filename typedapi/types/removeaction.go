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

// RemoveAction type.
type RemoveAction struct {
	Alias     *IndexAlias  `json:"alias,omitempty"`
	Aliases   []IndexAlias `json:"aliases,omitempty"`
	Index     *IndexName   `json:"index,omitempty"`
	Indices   *Indices     `json:"indices,omitempty"`
	MustExist *bool        `json:"must_exist,omitempty"`
}

// RemoveActionBuilder holds RemoveAction struct and provides a builder API.
type RemoveActionBuilder struct {
	v *RemoveAction
}

// NewRemoveAction provides a builder for the RemoveAction struct.
func NewRemoveAction() *RemoveActionBuilder {
	r := RemoveActionBuilder{
		&RemoveAction{},
	}

	return &r
}

// Build finalize the chain and returns the RemoveAction struct
func (rb *RemoveActionBuilder) Build() RemoveAction {
	return *rb.v
}

// Alias set the Alias property for RemoveActionBuilder.
func (rb *RemoveActionBuilder) Alias(alias IndexAlias) *RemoveActionBuilder {
	rb.v.Alias = &alias
	return rb
}

// Aliases set the Aliases property for RemoveActionBuilder.
func (rb *RemoveActionBuilder) Aliases(arg []IndexAlias) *RemoveActionBuilder {
	rb.v.Aliases = arg
	return rb
}

// Index set the Index property for RemoveActionBuilder.
func (rb *RemoveActionBuilder) Index(index IndexName) *RemoveActionBuilder {
	rb.v.Index = &index
	return rb
}

// Indices set the Indices property for RemoveActionBuilder.
func (rb *RemoveActionBuilder) Indices(indices Indices) *RemoveActionBuilder {
	rb.v.Indices = &indices
	return rb
}

// MustExist set the MustExist property for RemoveActionBuilder.
func (rb *RemoveActionBuilder) MustExist(mustexist bool) *RemoveActionBuilder {
	rb.v.MustExist = &mustexist
	return rb
}
