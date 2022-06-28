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

// SpanQuery type.
type SpanQuery struct {
	FieldMaskingSpan *SpanFieldMaskingQuery  `json:"field_masking_span,omitempty"`
	SpanContaining   *SpanContainingQuery    `json:"span_containing,omitempty"`
	SpanFirst        *SpanFirstQuery         `json:"span_first,omitempty"`
	SpanGap          *SpanGapQuery           `json:"span_gap,omitempty"`
	SpanMulti        *SpanMultiTermQuery     `json:"span_multi,omitempty"`
	SpanNear         *SpanNearQuery          `json:"span_near,omitempty"`
	SpanNot          *SpanNotQuery           `json:"span_not,omitempty"`
	SpanOr           *SpanOrQuery            `json:"span_or,omitempty"`
	SpanTerm         map[Field]SpanTermQuery `json:"span_term,omitempty"`
	SpanWithin       *SpanWithinQuery        `json:"span_within,omitempty"`
}

// SpanQueryBuilder holds SpanQuery struct and provides a builder API.
type SpanQueryBuilder struct {
	v *SpanQuery
}

// NewSpanQuery provides a builder for the SpanQuery struct.
func NewSpanQuery() *SpanQueryBuilder {
	r := SpanQueryBuilder{
		&SpanQuery{
			SpanTerm: make(map[Field]SpanTermQuery, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SpanQuery struct
func (rb *SpanQueryBuilder) Build() SpanQuery {
	return *rb.v
}

// FieldMaskingSpan set the FieldMaskingSpan property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) FieldMaskingSpan(fieldmaskingspan SpanFieldMaskingQuery) *SpanQueryBuilder {
	rb.v.FieldMaskingSpan = &fieldmaskingspan
	return rb
}

// SpanContaining set the SpanContaining property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanContaining(spancontaining SpanContainingQuery) *SpanQueryBuilder {
	rb.v.SpanContaining = &spancontaining
	return rb
}

// SpanFirst set the SpanFirst property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanFirst(spanfirst SpanFirstQuery) *SpanQueryBuilder {
	rb.v.SpanFirst = &spanfirst
	return rb
}

// SpanGap set the SpanGap property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanGap(spangap SpanGapQuery) *SpanQueryBuilder {
	rb.v.SpanGap = &spangap
	return rb
}

// SpanMulti set the SpanMulti property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanMulti(spanmulti SpanMultiTermQuery) *SpanQueryBuilder {
	rb.v.SpanMulti = &spanmulti
	return rb
}

// SpanNear set the SpanNear property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanNear(spannear SpanNearQuery) *SpanQueryBuilder {
	rb.v.SpanNear = &spannear
	return rb
}

// SpanNot set the SpanNot property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanNot(spannot SpanNotQuery) *SpanQueryBuilder {
	rb.v.SpanNot = &spannot
	return rb
}

// SpanOr set the SpanOr property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanOr(spanor SpanOrQuery) *SpanQueryBuilder {
	rb.v.SpanOr = &spanor
	return rb
}

// SpanTerm set the SpanTerm property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanTerm(value map[Field]SpanTermQuery) *SpanQueryBuilder {
	rb.v.SpanTerm = value
	return rb
}

// SpanWithin set the SpanWithin property for SpanQueryBuilder.
func (rb *SpanQueryBuilder) SpanWithin(spanwithin SpanWithinQuery) *SpanQueryBuilder {
	rb.v.SpanWithin = &spanwithin
	return rb
}
