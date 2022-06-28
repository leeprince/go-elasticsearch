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

// SpanWithinQuery type.
type SpanWithinQuery struct {
	Big        *SpanQuery `json:"big,omitempty"`
	Boost      *float32   `json:"boost,omitempty"`
	Little     *SpanQuery `json:"little,omitempty"`
	QueryName_ *string    `json:"_name,omitempty"`
}

// SpanWithinQueryBuilder holds SpanWithinQuery struct and provides a builder API.
type SpanWithinQueryBuilder struct {
	v *SpanWithinQuery
}

// NewSpanWithinQuery provides a builder for the SpanWithinQuery struct.
func NewSpanWithinQuery() *SpanWithinQueryBuilder {
	r := SpanWithinQueryBuilder{
		&SpanWithinQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SpanWithinQuery struct
func (rb *SpanWithinQueryBuilder) Build() SpanWithinQuery {
	return *rb.v
}

// Big set the Big property for SpanWithinQueryBuilder.
func (rb *SpanWithinQueryBuilder) Big(big SpanQuery) *SpanWithinQueryBuilder {
	rb.v.Big = &big
	return rb
}

// Boost set the Boost property for SpanWithinQueryBuilder.
func (rb *SpanWithinQueryBuilder) Boost(boost float32) *SpanWithinQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Little set the Little property for SpanWithinQueryBuilder.
func (rb *SpanWithinQueryBuilder) Little(little SpanQuery) *SpanWithinQueryBuilder {
	rb.v.Little = &little
	return rb
}

// QueryName_ set the QueryName_ property for SpanWithinQueryBuilder.
func (rb *SpanWithinQueryBuilder) QueryName_(queryname_ string) *SpanWithinQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
