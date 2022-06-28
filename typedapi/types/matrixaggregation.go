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

// MatrixAggregation type.
type MatrixAggregation struct {
	Fields  *Fields           `json:"fields,omitempty"`
	Meta    *Metadata         `json:"meta,omitempty"`
	Missing map[Field]float64 `json:"missing,omitempty"`
	Name    *string           `json:"name,omitempty"`
}

// MatrixAggregationBuilder holds MatrixAggregation struct and provides a builder API.
type MatrixAggregationBuilder struct {
	v *MatrixAggregation
}

// NewMatrixAggregation provides a builder for the MatrixAggregation struct.
func NewMatrixAggregation() *MatrixAggregationBuilder {
	r := MatrixAggregationBuilder{
		&MatrixAggregation{
			Missing: make(map[Field]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MatrixAggregation struct
func (rb *MatrixAggregationBuilder) Build() MatrixAggregation {
	return *rb.v
}

// Fields set the Fields property for MatrixAggregationBuilder.
func (rb *MatrixAggregationBuilder) Fields(fields Fields) *MatrixAggregationBuilder {
	rb.v.Fields = &fields
	return rb
}

// Meta set the Meta property for MatrixAggregationBuilder.
func (rb *MatrixAggregationBuilder) Meta(meta Metadata) *MatrixAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Missing set the Missing property for MatrixAggregationBuilder.
func (rb *MatrixAggregationBuilder) Missing(value map[Field]float64) *MatrixAggregationBuilder {
	rb.v.Missing = value
	return rb
}

// Name set the Name property for MatrixAggregationBuilder.
func (rb *MatrixAggregationBuilder) Name(name string) *MatrixAggregationBuilder {
	rb.v.Name = &name
	return rb
}
