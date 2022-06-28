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

// PercolateQuery type.
type PercolateQuery struct {
	Boost      *float32       `json:"boost,omitempty"`
	Document   interface{}    `json:"document,omitempty"`
	Documents  []interface{}  `json:"documents,omitempty"`
	Field      Field          `json:"field"`
	Id         *Id            `json:"id,omitempty"`
	Index      *IndexName     `json:"index,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Preference *string        `json:"preference,omitempty"`
	QueryName_ *string        `json:"_name,omitempty"`
	Routing    *Routing       `json:"routing,omitempty"`
	Version    *VersionNumber `json:"version,omitempty"`
}

// PercolateQueryBuilder holds PercolateQuery struct and provides a builder API.
type PercolateQueryBuilder struct {
	v *PercolateQuery
}

// NewPercolateQuery provides a builder for the PercolateQuery struct.
func NewPercolateQuery() *PercolateQueryBuilder {
	r := PercolateQueryBuilder{
		&PercolateQuery{},
	}

	return &r
}

// Build finalize the chain and returns the PercolateQuery struct
func (rb *PercolateQueryBuilder) Build() PercolateQuery {
	return *rb.v
}

// Boost set the Boost property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Boost(boost float32) *PercolateQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Document set the Document property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Document(document interface{}) *PercolateQueryBuilder {
	rb.v.Document = document
	return rb
}

// Documents set the Documents property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Documents(documents ...interface{}) *PercolateQueryBuilder {
	rb.v.Documents = append(rb.v.Documents, documents...)
	return rb
}

// Field set the Field property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Field(field Field) *PercolateQueryBuilder {
	rb.v.Field = field
	return rb
}

// Id set the Id property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Id(id Id) *PercolateQueryBuilder {
	rb.v.Id = &id
	return rb
}

// Index set the Index property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Index(index IndexName) *PercolateQueryBuilder {
	rb.v.Index = &index
	return rb
}

// Name set the Name property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Name(name string) *PercolateQueryBuilder {
	rb.v.Name = &name
	return rb
}

// Preference set the Preference property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Preference(preference string) *PercolateQueryBuilder {
	rb.v.Preference = &preference
	return rb
}

// QueryName_ set the QueryName_ property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) QueryName_(queryname_ string) *PercolateQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Routing set the Routing property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Routing(routing Routing) *PercolateQueryBuilder {
	rb.v.Routing = &routing
	return rb
}

// Version set the Version property for PercolateQueryBuilder.
func (rb *PercolateQueryBuilder) Version(version VersionNumber) *PercolateQueryBuilder {
	rb.v.Version = &version
	return rb
}
