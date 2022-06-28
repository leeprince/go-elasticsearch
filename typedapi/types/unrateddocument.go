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

// UnratedDocument type.
type UnratedDocument struct {
	Id_    Id        `json:"_id"`
	Index_ IndexName `json:"_index"`
}

// UnratedDocumentBuilder holds UnratedDocument struct and provides a builder API.
type UnratedDocumentBuilder struct {
	v *UnratedDocument
}

// NewUnratedDocument provides a builder for the UnratedDocument struct.
func NewUnratedDocument() *UnratedDocumentBuilder {
	r := UnratedDocumentBuilder{
		&UnratedDocument{},
	}

	return &r
}

// Build finalize the chain and returns the UnratedDocument struct
func (rb *UnratedDocumentBuilder) Build() UnratedDocument {
	return *rb.v
}

// Id_ set the Id_ property for UnratedDocumentBuilder.
func (rb *UnratedDocumentBuilder) Id_(id_ Id) *UnratedDocumentBuilder {
	rb.v.Id_ = id_
	return rb
}

// Index_ set the Index_ property for UnratedDocumentBuilder.
func (rb *UnratedDocumentBuilder) Index_(index_ IndexName) *UnratedDocumentBuilder {
	rb.v.Index_ = index_
	return rb
}
