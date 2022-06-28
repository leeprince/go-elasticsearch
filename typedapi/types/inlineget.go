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

// InlineGet type.
type InlineGet struct {
	Fields       map[string]interface{} `json:"fields,omitempty"`
	Found        bool                   `json:"found"`
	Metadata     map[string]interface{} `json:"metadata,omitempty"`
	PrimaryTerm_ *int64                 `json:"_primary_term,omitempty"`
	Routing_     *Routing               `json:"_routing,omitempty"`
	SeqNo_       *SequenceNumber        `json:"_seq_no,omitempty"`
	Source_      interface{}            `json:"_source,omitempty"`
}

// InlineGetBuilder holds InlineGet struct and provides a builder API.
type InlineGetBuilder struct {
	v *InlineGet
}

// NewInlineGet provides a builder for the InlineGet struct.
func NewInlineGet() *InlineGetBuilder {
	r := InlineGetBuilder{
		&InlineGet{
			Fields:   make(map[string]interface{}, 0),
			Metadata: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InlineGet struct
func (rb *InlineGetBuilder) Build() InlineGet {
	return *rb.v
}

// Fields set the Fields property for InlineGetBuilder.
func (rb *InlineGetBuilder) Fields(value map[string]interface{}) *InlineGetBuilder {
	rb.v.Fields = value
	return rb
}

// Found set the Found property for InlineGetBuilder.
func (rb *InlineGetBuilder) Found(found bool) *InlineGetBuilder {
	rb.v.Found = found
	return rb
}

// Metadata set the Metadata property for InlineGetBuilder.
func (rb *InlineGetBuilder) Metadata(value map[string]interface{}) *InlineGetBuilder {
	rb.v.Metadata = value
	return rb
}

// PrimaryTerm_ set the PrimaryTerm_ property for InlineGetBuilder.
func (rb *InlineGetBuilder) PrimaryTerm_(primaryterm_ int64) *InlineGetBuilder {
	rb.v.PrimaryTerm_ = &primaryterm_
	return rb
}

// Routing_ set the Routing_ property for InlineGetBuilder.
func (rb *InlineGetBuilder) Routing_(routing_ Routing) *InlineGetBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

// SeqNo_ set the SeqNo_ property for InlineGetBuilder.
func (rb *InlineGetBuilder) SeqNo_(seqno_ SequenceNumber) *InlineGetBuilder {
	rb.v.SeqNo_ = &seqno_
	return rb
}

// Source_ set the Source_ property for InlineGetBuilder.
func (rb *InlineGetBuilder) Source_(source_ interface{}) *InlineGetBuilder {
	rb.v.Source_ = source_
	return rb
}
