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

// TermVectorsResult type.
type TermVectorsResult struct {
	Error       *ErrorCause          `json:"error,omitempty"`
	Found       *bool                `json:"found,omitempty"`
	Id_         Id                   `json:"_id"`
	Index_      IndexName            `json:"_index"`
	TermVectors map[Field]TermVector `json:"term_vectors,omitempty"`
	Took        *int64               `json:"took,omitempty"`
	Version_    *VersionNumber       `json:"_version,omitempty"`
}

// TermVectorsResultBuilder holds TermVectorsResult struct and provides a builder API.
type TermVectorsResultBuilder struct {
	v *TermVectorsResult
}

// NewTermVectorsResult provides a builder for the TermVectorsResult struct.
func NewTermVectorsResult() *TermVectorsResultBuilder {
	r := TermVectorsResultBuilder{
		&TermVectorsResult{
			TermVectors: make(map[Field]TermVector, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TermVectorsResult struct
func (rb *TermVectorsResultBuilder) Build() TermVectorsResult {
	return *rb.v
}

// Error set the Error property for TermVectorsResultBuilder.
func (rb *TermVectorsResultBuilder) Error(error ErrorCause) *TermVectorsResultBuilder {
	rb.v.Error = &error
	return rb
}

// Found set the Found property for TermVectorsResultBuilder.
func (rb *TermVectorsResultBuilder) Found(found bool) *TermVectorsResultBuilder {
	rb.v.Found = &found
	return rb
}

// Id_ set the Id_ property for TermVectorsResultBuilder.
func (rb *TermVectorsResultBuilder) Id_(id_ Id) *TermVectorsResultBuilder {
	rb.v.Id_ = id_
	return rb
}

// Index_ set the Index_ property for TermVectorsResultBuilder.
func (rb *TermVectorsResultBuilder) Index_(index_ IndexName) *TermVectorsResultBuilder {
	rb.v.Index_ = index_
	return rb
}

// TermVectors set the TermVectors property for TermVectorsResultBuilder.
func (rb *TermVectorsResultBuilder) TermVectors(value map[Field]TermVector) *TermVectorsResultBuilder {
	rb.v.TermVectors = value
	return rb
}

// Took set the Took property for TermVectorsResultBuilder.
func (rb *TermVectorsResultBuilder) Took(took int64) *TermVectorsResultBuilder {
	rb.v.Took = &took
	return rb
}

// Version_ set the Version_ property for TermVectorsResultBuilder.
func (rb *TermVectorsResultBuilder) Version_(version_ VersionNumber) *TermVectorsResultBuilder {
	rb.v.Version_ = &version_
	return rb
}
