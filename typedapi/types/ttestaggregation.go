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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/ttesttype"
)

// TTestAggregation type.
type TTestAggregation struct {
	A    *TestPopulation      `json:"a,omitempty"`
	B    *TestPopulation      `json:"b,omitempty"`
	Meta *Metadata            `json:"meta,omitempty"`
	Name *string              `json:"name,omitempty"`
	Type *ttesttype.TTestType `json:"type,omitempty"`
}

// TTestAggregationBuilder holds TTestAggregation struct and provides a builder API.
type TTestAggregationBuilder struct {
	v *TTestAggregation
}

// NewTTestAggregation provides a builder for the TTestAggregation struct.
func NewTTestAggregation() *TTestAggregationBuilder {
	r := TTestAggregationBuilder{
		&TTestAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the TTestAggregation struct
func (rb *TTestAggregationBuilder) Build() TTestAggregation {
	return *rb.v
}

// A set the A property for TTestAggregationBuilder.
func (rb *TTestAggregationBuilder) A(a TestPopulation) *TTestAggregationBuilder {
	rb.v.A = &a
	return rb
}

// B set the B property for TTestAggregationBuilder.
func (rb *TTestAggregationBuilder) B(b TestPopulation) *TTestAggregationBuilder {
	rb.v.B = &b
	return rb
}

// Meta set the Meta property for TTestAggregationBuilder.
func (rb *TTestAggregationBuilder) Meta(meta Metadata) *TTestAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for TTestAggregationBuilder.
func (rb *TTestAggregationBuilder) Name(name string) *TTestAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Type set the Type property for TTestAggregationBuilder.
func (rb *TTestAggregationBuilder) Type_(type_ ttesttype.TTestType) *TTestAggregationBuilder {
	rb.v.Type = &type_
	return rb
}
