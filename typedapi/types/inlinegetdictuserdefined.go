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

// InlineGetDictUserDefined type.
type InlineGetDictUserDefined struct {
	Fields                   map[string]interface{} `json:"fields,omitempty"`
	Found                    bool                   `json:"found"`
	InlineGetDictUserDefined map[string]interface{} `json:"InlineGetDictUserDefined,omitempty"`
	PrimaryTerm_             *int64                 `json:"_primary_term,omitempty"`
	Routing_                 *Routing               `json:"_routing,omitempty"`
	SeqNo_                   *SequenceNumber        `json:"_seq_no,omitempty"`
	Source_                  map[string]interface{} `json:"_source"`
}

// InlineGetDictUserDefinedBuilder holds InlineGetDictUserDefined struct and provides a builder API.
type InlineGetDictUserDefinedBuilder struct {
	v *InlineGetDictUserDefined
}

// NewInlineGetDictUserDefined provides a builder for the InlineGetDictUserDefined struct.
func NewInlineGetDictUserDefined() *InlineGetDictUserDefinedBuilder {
	r := InlineGetDictUserDefinedBuilder{
		&InlineGetDictUserDefined{
			Fields:                   make(map[string]interface{}, 0),
			InlineGetDictUserDefined: make(map[string]interface{}, 0),
			Source_:                  make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InlineGetDictUserDefined struct
func (rb *InlineGetDictUserDefinedBuilder) Build() InlineGetDictUserDefined {
	return *rb.v
}

// Fields set the Fields property for InlineGetDictUserDefinedBuilder.
func (rb *InlineGetDictUserDefinedBuilder) Fields(value map[string]interface{}) *InlineGetDictUserDefinedBuilder {
	rb.v.Fields = value
	return rb
}

// Found set the Found property for InlineGetDictUserDefinedBuilder.
func (rb *InlineGetDictUserDefinedBuilder) Found(found bool) *InlineGetDictUserDefinedBuilder {
	rb.v.Found = found
	return rb
}

// InlineGetDictUserDefined set the InlineGetDictUserDefined property for InlineGetDictUserDefinedBuilder.
func (rb *InlineGetDictUserDefinedBuilder) InlineGetDictUserDefined(value map[string]interface{}) *InlineGetDictUserDefinedBuilder {
	rb.v.InlineGetDictUserDefined = value
	return rb
}

// PrimaryTerm_ set the PrimaryTerm_ property for InlineGetDictUserDefinedBuilder.
func (rb *InlineGetDictUserDefinedBuilder) PrimaryTerm_(primaryterm_ int64) *InlineGetDictUserDefinedBuilder {
	rb.v.PrimaryTerm_ = &primaryterm_
	return rb
}

// Routing_ set the Routing_ property for InlineGetDictUserDefinedBuilder.
func (rb *InlineGetDictUserDefinedBuilder) Routing_(routing_ Routing) *InlineGetDictUserDefinedBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

// SeqNo_ set the SeqNo_ property for InlineGetDictUserDefinedBuilder.
func (rb *InlineGetDictUserDefinedBuilder) SeqNo_(seqno_ SequenceNumber) *InlineGetDictUserDefinedBuilder {
	rb.v.SeqNo_ = &seqno_
	return rb
}

// Source_ set the Source_ property for InlineGetDictUserDefinedBuilder.
func (rb *InlineGetDictUserDefinedBuilder) Source_(value map[string]interface{}) *InlineGetDictUserDefinedBuilder {
	rb.v.Source_ = value
	return rb
}
