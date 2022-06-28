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

// MatrixStatsAggregate type.
type MatrixStatsAggregate struct {
	DocCount int64               `json:"doc_count"`
	Fields   []MatrixStatsFields `json:"fields"`
	Meta     *Metadata           `json:"meta,omitempty"`
}

// MatrixStatsAggregateBuilder holds MatrixStatsAggregate struct and provides a builder API.
type MatrixStatsAggregateBuilder struct {
	v *MatrixStatsAggregate
}

// NewMatrixStatsAggregate provides a builder for the MatrixStatsAggregate struct.
func NewMatrixStatsAggregate() *MatrixStatsAggregateBuilder {
	r := MatrixStatsAggregateBuilder{
		&MatrixStatsAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the MatrixStatsAggregate struct
func (rb *MatrixStatsAggregateBuilder) Build() MatrixStatsAggregate {
	return *rb.v
}

// DocCount set the DocCount property for MatrixStatsAggregateBuilder.
func (rb *MatrixStatsAggregateBuilder) DocCount(doccount int64) *MatrixStatsAggregateBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Fields set the Fields property for MatrixStatsAggregateBuilder.
func (rb *MatrixStatsAggregateBuilder) Fields(fields ...MatrixStatsFields) *MatrixStatsAggregateBuilder {
	rb.v.Fields = append(rb.v.Fields, fields...)
	return rb
}

// Meta set the Meta property for MatrixStatsAggregateBuilder.
func (rb *MatrixStatsAggregateBuilder) Meta(meta Metadata) *MatrixStatsAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}
