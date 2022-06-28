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

// GetResult type.
type GetResult struct {
	Fields       map[string]interface{} `json:"fields,omitempty"`
	Found        bool                   `json:"found"`
	Id_          Id                     `json:"_id"`
	Index_       IndexName              `json:"_index"`
	PrimaryTerm_ *int64                 `json:"_primary_term,omitempty"`
	Routing_     *string                `json:"_routing,omitempty"`
	SeqNo_       *SequenceNumber        `json:"_seq_no,omitempty"`
	Source_      interface{}            `json:"_source,omitempty"`
	Version_     *VersionNumber         `json:"_version,omitempty"`
}

// GetResultBuilder holds GetResult struct and provides a builder API.
type GetResultBuilder struct {
	v *GetResult
}

// NewGetResult provides a builder for the GetResult struct.
func NewGetResult() *GetResultBuilder {
	r := GetResultBuilder{
		&GetResult{
			Fields: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GetResult struct
func (rb *GetResultBuilder) Build() GetResult {
	return *rb.v
}

// Fields set the Fields property for GetResultBuilder.
func (rb *GetResultBuilder) Fields(value map[string]interface{}) *GetResultBuilder {
	rb.v.Fields = value
	return rb
}

// Found set the Found property for GetResultBuilder.
func (rb *GetResultBuilder) Found(found bool) *GetResultBuilder {
	rb.v.Found = found
	return rb
}

// Id_ set the Id_ property for GetResultBuilder.
func (rb *GetResultBuilder) Id_(id_ Id) *GetResultBuilder {
	rb.v.Id_ = id_
	return rb
}

// Index_ set the Index_ property for GetResultBuilder.
func (rb *GetResultBuilder) Index_(index_ IndexName) *GetResultBuilder {
	rb.v.Index_ = index_
	return rb
}

// PrimaryTerm_ set the PrimaryTerm_ property for GetResultBuilder.
func (rb *GetResultBuilder) PrimaryTerm_(primaryterm_ int64) *GetResultBuilder {
	rb.v.PrimaryTerm_ = &primaryterm_
	return rb
}

// Routing_ set the Routing_ property for GetResultBuilder.
func (rb *GetResultBuilder) Routing_(routing_ string) *GetResultBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

// SeqNo_ set the SeqNo_ property for GetResultBuilder.
func (rb *GetResultBuilder) SeqNo_(seqno_ SequenceNumber) *GetResultBuilder {
	rb.v.SeqNo_ = &seqno_
	return rb
}

// Source_ set the Source_ property for GetResultBuilder.
func (rb *GetResultBuilder) Source_(source_ interface{}) *GetResultBuilder {
	rb.v.Source_ = source_
	return rb
}

// Version_ set the Version_ property for GetResultBuilder.
func (rb *GetResultBuilder) Version_(version_ VersionNumber) *GetResultBuilder {
	rb.v.Version_ = &version_
	return rb
}
