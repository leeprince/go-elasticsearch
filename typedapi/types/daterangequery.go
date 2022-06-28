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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

// DateRangeQuery type.
type DateRangeQuery struct {
	Boost      *float32                     `json:"boost,omitempty"`
	Format     *DateFormat                  `json:"format,omitempty"`
	From       DateMath                     `json:"from,omitempty"`
	Gt         *DateMath                    `json:"gt,omitempty"`
	Gte        *DateMath                    `json:"gte,omitempty"`
	Lt         *DateMath                    `json:"lt,omitempty"`
	Lte        *DateMath                    `json:"lte,omitempty"`
	QueryName_ *string                      `json:"_name,omitempty"`
	Relation   *rangerelation.RangeRelation `json:"relation,omitempty"`
	TimeZone   *TimeZone                    `json:"time_zone,omitempty"`
	To         DateMath                     `json:"to,omitempty"`
}

// DateRangeQueryBuilder holds DateRangeQuery struct and provides a builder API.
type DateRangeQueryBuilder struct {
	v *DateRangeQuery
}

// NewDateRangeQuery provides a builder for the DateRangeQuery struct.
func NewDateRangeQuery() *DateRangeQueryBuilder {
	r := DateRangeQueryBuilder{
		&DateRangeQuery{},
	}

	return &r
}

// Build finalize the chain and returns the DateRangeQuery struct
func (rb *DateRangeQueryBuilder) Build() DateRangeQuery {
	return *rb.v
}

// Boost set the Boost property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) Boost(boost float32) *DateRangeQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// Format set the Format property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) Format(format DateFormat) *DateRangeQueryBuilder {
	rb.v.Format = &format
	return rb
}

// From set the From property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) From(from DateMath) *DateRangeQueryBuilder {
	rb.v.From = from
	return rb
}

// Gt set the Gt property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) Gt(gt DateMath) *DateRangeQueryBuilder {
	rb.v.Gt = &gt
	return rb
}

// Gte set the Gte property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) Gte(gte DateMath) *DateRangeQueryBuilder {
	rb.v.Gte = &gte
	return rb
}

// Lt set the Lt property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) Lt(lt DateMath) *DateRangeQueryBuilder {
	rb.v.Lt = &lt
	return rb
}

// Lte set the Lte property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) Lte(lte DateMath) *DateRangeQueryBuilder {
	rb.v.Lte = &lte
	return rb
}

// QueryName_ set the QueryName_ property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) QueryName_(queryname_ string) *DateRangeQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Relation set the Relation property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) Relation(relation rangerelation.RangeRelation) *DateRangeQueryBuilder {
	rb.v.Relation = &relation
	return rb
}

// TimeZone set the TimeZone property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) TimeZone(timezone TimeZone) *DateRangeQueryBuilder {
	rb.v.TimeZone = &timezone
	return rb
}

// To set the To property for DateRangeQueryBuilder.
func (rb *DateRangeQueryBuilder) To(to DateMath) *DateRangeQueryBuilder {
	rb.v.To = to
	return rb
}
