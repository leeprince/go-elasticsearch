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

// CompositeAggregation type.
type CompositeAggregation struct {
	After   map[string]string                       `json:"after,omitempty"`
	Meta    *Metadata                               `json:"meta,omitempty"`
	Name    *string                                 `json:"name,omitempty"`
	Size    *int                                    `json:"size,omitempty"`
	Sources []map[string]CompositeAggregationSource `json:"sources,omitempty"`
}

// CompositeAggregationBuilder holds CompositeAggregation struct and provides a builder API.
type CompositeAggregationBuilder struct {
	v *CompositeAggregation
}

// NewCompositeAggregation provides a builder for the CompositeAggregation struct.
func NewCompositeAggregation() *CompositeAggregationBuilder {
	r := CompositeAggregationBuilder{
		&CompositeAggregation{
			After: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CompositeAggregation struct
func (rb *CompositeAggregationBuilder) Build() CompositeAggregation {
	return *rb.v
}

// After set the After property for CompositeAggregationBuilder.
//TODO: Union in map Not implemented yet: UnionOf
func (rb *CompositeAggregationBuilder) After(value map[string]string) *CompositeAggregationBuilder {
	rb.v.After = value
	return rb
}

// Meta set the Meta property for CompositeAggregationBuilder.
func (rb *CompositeAggregationBuilder) Meta(meta Metadata) *CompositeAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for CompositeAggregationBuilder.
func (rb *CompositeAggregationBuilder) Name(name string) *CompositeAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Size set the Size property for CompositeAggregationBuilder.
func (rb *CompositeAggregationBuilder) Size(size int) *CompositeAggregationBuilder {
	rb.v.Size = &size
	return rb
}

// Sources set the Sources property for CompositeAggregationBuilder.
func (rb *CompositeAggregationBuilder) Sources(value ...map[string]CompositeAggregationSource) *CompositeAggregationBuilder {
	rb.v.Sources = append(rb.v.Sources, value...)
	return rb
}
