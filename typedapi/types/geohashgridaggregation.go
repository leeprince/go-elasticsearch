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

// GeoHashGridAggregation type.
type GeoHashGridAggregation struct {
	Bounds    *GeoBounds        `json:"bounds,omitempty"`
	Field     *Field            `json:"field,omitempty"`
	Meta      *Metadata         `json:"meta,omitempty"`
	Name      *string           `json:"name,omitempty"`
	Precision *GeoHashPrecision `json:"precision,omitempty"`
	ShardSize *int              `json:"shard_size,omitempty"`
	Size      *int              `json:"size,omitempty"`
}

// GeoHashGridAggregationBuilder holds GeoHashGridAggregation struct and provides a builder API.
type GeoHashGridAggregationBuilder struct {
	v *GeoHashGridAggregation
}

// NewGeoHashGridAggregation provides a builder for the GeoHashGridAggregation struct.
func NewGeoHashGridAggregation() *GeoHashGridAggregationBuilder {
	r := GeoHashGridAggregationBuilder{
		&GeoHashGridAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the GeoHashGridAggregation struct
func (rb *GeoHashGridAggregationBuilder) Build() GeoHashGridAggregation {
	return *rb.v
}

// Bounds set the Bounds property for GeoHashGridAggregationBuilder.
func (rb *GeoHashGridAggregationBuilder) Bounds(bounds GeoBounds) *GeoHashGridAggregationBuilder {
	rb.v.Bounds = &bounds
	return rb
}

// Field set the Field property for GeoHashGridAggregationBuilder.
func (rb *GeoHashGridAggregationBuilder) Field(field Field) *GeoHashGridAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Meta set the Meta property for GeoHashGridAggregationBuilder.
func (rb *GeoHashGridAggregationBuilder) Meta(meta Metadata) *GeoHashGridAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for GeoHashGridAggregationBuilder.
func (rb *GeoHashGridAggregationBuilder) Name(name string) *GeoHashGridAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Precision set the Precision property for GeoHashGridAggregationBuilder.
func (rb *GeoHashGridAggregationBuilder) Precision(precision GeoHashPrecision) *GeoHashGridAggregationBuilder {
	rb.v.Precision = &precision
	return rb
}

// ShardSize set the ShardSize property for GeoHashGridAggregationBuilder.
func (rb *GeoHashGridAggregationBuilder) ShardSize(shardsize int) *GeoHashGridAggregationBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

// Size set the Size property for GeoHashGridAggregationBuilder.
func (rb *GeoHashGridAggregationBuilder) Size(size int) *GeoHashGridAggregationBuilder {
	rb.v.Size = &size
	return rb
}
