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

// DenseVectorIndexOptions type.
type DenseVectorIndexOptions struct {
	EfConstruction int    `json:"ef_construction"`
	M              int    `json:"m"`
	Type           string `json:"type"`
}

// DenseVectorIndexOptionsBuilder holds DenseVectorIndexOptions struct and provides a builder API.
type DenseVectorIndexOptionsBuilder struct {
	v *DenseVectorIndexOptions
}

// NewDenseVectorIndexOptions provides a builder for the DenseVectorIndexOptions struct.
func NewDenseVectorIndexOptions() *DenseVectorIndexOptionsBuilder {
	r := DenseVectorIndexOptionsBuilder{
		&DenseVectorIndexOptions{},
	}

	return &r
}

// Build finalize the chain and returns the DenseVectorIndexOptions struct
func (rb *DenseVectorIndexOptionsBuilder) Build() DenseVectorIndexOptions {
	return *rb.v
}

// EfConstruction set the EfConstruction property for DenseVectorIndexOptionsBuilder.
func (rb *DenseVectorIndexOptionsBuilder) EfConstruction(efconstruction int) *DenseVectorIndexOptionsBuilder {
	rb.v.EfConstruction = efconstruction
	return rb
}

// M set the M property for DenseVectorIndexOptionsBuilder.
func (rb *DenseVectorIndexOptionsBuilder) M(m int) *DenseVectorIndexOptionsBuilder {
	rb.v.M = m
	return rb
}

// Type set the Type property for DenseVectorIndexOptionsBuilder.
func (rb *DenseVectorIndexOptionsBuilder) Type_(type_ string) *DenseVectorIndexOptionsBuilder {
	rb.v.Type = type_
	return rb
}
