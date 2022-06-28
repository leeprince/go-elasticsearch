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

// AdjacencyMatrixAggregation type.
type AdjacencyMatrixAggregation struct {
	Filters map[string]QueryContainer `json:"filters,omitempty"`
	Meta    *Metadata                 `json:"meta,omitempty"`
	Name    *string                   `json:"name,omitempty"`
}

// AdjacencyMatrixAggregationBuilder holds AdjacencyMatrixAggregation struct and provides a builder API.
type AdjacencyMatrixAggregationBuilder struct {
	v *AdjacencyMatrixAggregation
}

// NewAdjacencyMatrixAggregation provides a builder for the AdjacencyMatrixAggregation struct.
func NewAdjacencyMatrixAggregation() *AdjacencyMatrixAggregationBuilder {
	r := AdjacencyMatrixAggregationBuilder{
		&AdjacencyMatrixAggregation{
			Filters: make(map[string]QueryContainer, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AdjacencyMatrixAggregation struct
func (rb *AdjacencyMatrixAggregationBuilder) Build() AdjacencyMatrixAggregation {
	return *rb.v
}

// Filters set the Filters property for AdjacencyMatrixAggregationBuilder.
func (rb *AdjacencyMatrixAggregationBuilder) Filters(value map[string]QueryContainer) *AdjacencyMatrixAggregationBuilder {
	rb.v.Filters = value
	return rb
}

// Meta set the Meta property for AdjacencyMatrixAggregationBuilder.
func (rb *AdjacencyMatrixAggregationBuilder) Meta(meta Metadata) *AdjacencyMatrixAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for AdjacencyMatrixAggregationBuilder.
func (rb *AdjacencyMatrixAggregationBuilder) Name(name string) *AdjacencyMatrixAggregationBuilder {
	rb.v.Name = &name
	return rb
}
