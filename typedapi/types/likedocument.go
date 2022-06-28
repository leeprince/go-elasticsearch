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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/versiontype"
)

// LikeDocument type.
type LikeDocument struct {
	Doc              interface{}              `json:"doc,omitempty"`
	Fields           []Field                  `json:"fields,omitempty"`
	Id_              *Id                      `json:"_id,omitempty"`
	Index_           *IndexName               `json:"_index,omitempty"`
	PerFieldAnalyzer map[Field]string         `json:"per_field_analyzer,omitempty"`
	Routing          *Routing                 `json:"routing,omitempty"`
	Version          *VersionNumber           `json:"version,omitempty"`
	VersionType      *versiontype.VersionType `json:"version_type,omitempty"`
}

// LikeDocumentBuilder holds LikeDocument struct and provides a builder API.
type LikeDocumentBuilder struct {
	v *LikeDocument
}

// NewLikeDocument provides a builder for the LikeDocument struct.
func NewLikeDocument() *LikeDocumentBuilder {
	r := LikeDocumentBuilder{
		&LikeDocument{
			PerFieldAnalyzer: make(map[Field]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the LikeDocument struct
func (rb *LikeDocumentBuilder) Build() LikeDocument {
	return *rb.v
}

// Doc set the Doc property for LikeDocumentBuilder.
func (rb *LikeDocumentBuilder) Doc(doc interface{}) *LikeDocumentBuilder {
	rb.v.Doc = doc
	return rb
}

// Fields set the Fields property for LikeDocumentBuilder.
func (rb *LikeDocumentBuilder) Fields(fields ...Field) *LikeDocumentBuilder {
	rb.v.Fields = append(rb.v.Fields, fields...)
	return rb
}

// Id_ set the Id_ property for LikeDocumentBuilder.
func (rb *LikeDocumentBuilder) Id_(id_ Id) *LikeDocumentBuilder {
	rb.v.Id_ = &id_
	return rb
}

// Index_ set the Index_ property for LikeDocumentBuilder.
func (rb *LikeDocumentBuilder) Index_(index_ IndexName) *LikeDocumentBuilder {
	rb.v.Index_ = &index_
	return rb
}

// PerFieldAnalyzer set the PerFieldAnalyzer property for LikeDocumentBuilder.
func (rb *LikeDocumentBuilder) PerFieldAnalyzer(value map[Field]string) *LikeDocumentBuilder {
	rb.v.PerFieldAnalyzer = value
	return rb
}

// Routing set the Routing property for LikeDocumentBuilder.
func (rb *LikeDocumentBuilder) Routing(routing Routing) *LikeDocumentBuilder {
	rb.v.Routing = &routing
	return rb
}

// Version set the Version property for LikeDocumentBuilder.
func (rb *LikeDocumentBuilder) Version(version VersionNumber) *LikeDocumentBuilder {
	rb.v.Version = &version
	return rb
}

// VersionType set the VersionType property for LikeDocumentBuilder.
func (rb *LikeDocumentBuilder) VersionType(versiontype versiontype.VersionType) *LikeDocumentBuilder {
	rb.v.VersionType = &versiontype
	return rb
}
