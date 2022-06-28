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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/calendarinterval"
)

// DateHistogramAggregation type.
type DateHistogramAggregation struct {
	CalendarInterval *calendarinterval.CalendarInterval `json:"calendar_interval,omitempty"`
	ExtendedBounds   *ExtendedBoundsFieldDateMath       `json:"extended_bounds,omitempty"`
	Field            *Field                             `json:"field,omitempty"`
	FixedInterval    *Time                              `json:"fixed_interval,omitempty"`
	Format           *string                            `json:"format,omitempty"`
	HardBounds       *ExtendedBoundsFieldDateMath       `json:"hard_bounds,omitempty"`
	Interval         *Time                              `json:"interval,omitempty"`
	Keyed            *bool                              `json:"keyed,omitempty"`
	Meta             *Metadata                          `json:"meta,omitempty"`
	MinDocCount      *int                               `json:"min_doc_count,omitempty"`
	Missing          *DateString                        `json:"missing,omitempty"`
	Name             *string                            `json:"name,omitempty"`
	Offset           *Time                              `json:"offset,omitempty"`
	Order            *HistogramOrder                    `json:"order,omitempty"`
	Params           map[string]interface{}             `json:"params,omitempty"`
	Script           *Script                            `json:"script,omitempty"`
	TimeZone         *string                            `json:"time_zone,omitempty"`
}

// DateHistogramAggregationBuilder holds DateHistogramAggregation struct and provides a builder API.
type DateHistogramAggregationBuilder struct {
	v *DateHistogramAggregation
}

// NewDateHistogramAggregation provides a builder for the DateHistogramAggregation struct.
func NewDateHistogramAggregation() *DateHistogramAggregationBuilder {
	r := DateHistogramAggregationBuilder{
		&DateHistogramAggregation{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DateHistogramAggregation struct
func (rb *DateHistogramAggregationBuilder) Build() DateHistogramAggregation {
	return *rb.v
}

// CalendarInterval set the CalendarInterval property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) CalendarInterval(calendarinterval calendarinterval.CalendarInterval) *DateHistogramAggregationBuilder {
	rb.v.CalendarInterval = &calendarinterval
	return rb
}

// ExtendedBounds set the ExtendedBounds property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) ExtendedBounds(extendedbounds ExtendedBoundsFieldDateMath) *DateHistogramAggregationBuilder {
	rb.v.ExtendedBounds = &extendedbounds
	return rb
}

// Field set the Field property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Field(field Field) *DateHistogramAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// FixedInterval set the FixedInterval property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) FixedInterval(fixedinterval Time) *DateHistogramAggregationBuilder {
	rb.v.FixedInterval = &fixedinterval
	return rb
}

// Format set the Format property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Format(format string) *DateHistogramAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// HardBounds set the HardBounds property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) HardBounds(hardbounds ExtendedBoundsFieldDateMath) *DateHistogramAggregationBuilder {
	rb.v.HardBounds = &hardbounds
	return rb
}

// Interval set the Interval property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Interval(interval Time) *DateHistogramAggregationBuilder {
	rb.v.Interval = &interval
	return rb
}

// Keyed set the Keyed property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Keyed(keyed bool) *DateHistogramAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

// Meta set the Meta property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Meta(meta Metadata) *DateHistogramAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// MinDocCount set the MinDocCount property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) MinDocCount(mindoccount int) *DateHistogramAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

// Missing set the Missing property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Missing(missing DateString) *DateHistogramAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Name set the Name property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Name(name string) *DateHistogramAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Offset set the Offset property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Offset(offset Time) *DateHistogramAggregationBuilder {
	rb.v.Offset = &offset
	return rb
}

// Order set the Order property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Order(order HistogramOrder) *DateHistogramAggregationBuilder {
	rb.v.Order = &order
	return rb
}

// Params set the Params property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Params(value map[string]interface{}) *DateHistogramAggregationBuilder {
	rb.v.Params = value
	return rb
}

// Script set the Script property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) Script(script Script) *DateHistogramAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// TimeZone set the TimeZone property for DateHistogramAggregationBuilder.
func (rb *DateHistogramAggregationBuilder) TimeZone(timezone string) *DateHistogramAggregationBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
