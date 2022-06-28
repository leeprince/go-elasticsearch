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

// MultiTermsAggregation type.
type MultiTermsAggregation struct {
	Meta  *Metadata         `json:"meta,omitempty"`
	Name  *string           `json:"name,omitempty"`
	Terms []MultiTermLookup `json:"terms"`
}

// MultiTermsAggregationBuilder holds MultiTermsAggregation struct and provides a builder API.
type MultiTermsAggregationBuilder struct {
	v *MultiTermsAggregation
}

// NewMultiTermsAggregation provides a builder for the MultiTermsAggregation struct.
func NewMultiTermsAggregation() *MultiTermsAggregationBuilder {
	r := MultiTermsAggregationBuilder{
		&MultiTermsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MultiTermsAggregation struct
func (rb *MultiTermsAggregationBuilder) Build() MultiTermsAggregation {
	return *rb.v
}

// Meta set the Meta property for MultiTermsAggregationBuilder.
func (rb *MultiTermsAggregationBuilder) Meta(meta Metadata) *MultiTermsAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// Name set the Name property for MultiTermsAggregationBuilder.
func (rb *MultiTermsAggregationBuilder) Name(name string) *MultiTermsAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Terms set the Terms property for MultiTermsAggregationBuilder.
func (rb *MultiTermsAggregationBuilder) Terms(terms ...MultiTermLookup) *MultiTermsAggregationBuilder {
	rb.v.Terms = append(rb.v.Terms, terms...)
	return rb
}
