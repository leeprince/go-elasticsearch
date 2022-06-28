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

// Document type.
type Document struct {
	Id_     *Id         `json:"_id,omitempty"`
	Index_  *IndexName  `json:"_index,omitempty"`
	Source_ interface{} `json:"_source,omitempty"`
}

// DocumentBuilder holds Document struct and provides a builder API.
type DocumentBuilder struct {
	v *Document
}

// NewDocument provides a builder for the Document struct.
func NewDocument() *DocumentBuilder {
	r := DocumentBuilder{
		&Document{},
	}

	return &r
}

// Build finalize the chain and returns the Document struct
func (rb *DocumentBuilder) Build() Document {
	return *rb.v
}

// Id_ set the Id_ property for DocumentBuilder.
func (rb *DocumentBuilder) Id_(id_ Id) *DocumentBuilder {
	rb.v.Id_ = &id_
	return rb
}

// Index_ set the Index_ property for DocumentBuilder.
func (rb *DocumentBuilder) Index_(index_ IndexName) *DocumentBuilder {
	rb.v.Index_ = &index_
	return rb
}

// Source_ set the Source_ property for DocumentBuilder.
func (rb *DocumentBuilder) Source_(source_ interface{}) *DocumentBuilder {
	rb.v.Source_ = source_
	return rb
}
