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

// DateHistogramGrouping type.
type DateHistogramGrouping struct {
	CalendarInterval *Time   `json:"calendar_interval,omitempty"`
	Delay            *Time   `json:"delay,omitempty"`
	Field            Field   `json:"field"`
	FixedInterval    *Time   `json:"fixed_interval,omitempty"`
	Format           *string `json:"format,omitempty"`
	Interval         *Time   `json:"interval,omitempty"`
	TimeZone         *string `json:"time_zone,omitempty"`
}

// DateHistogramGroupingBuilder holds DateHistogramGrouping struct and provides a builder API.
type DateHistogramGroupingBuilder struct {
	v *DateHistogramGrouping
}

// NewDateHistogramGrouping provides a builder for the DateHistogramGrouping struct.
func NewDateHistogramGrouping() *DateHistogramGroupingBuilder {
	r := DateHistogramGroupingBuilder{
		&DateHistogramGrouping{},
	}

	return &r
}

// Build finalize the chain and returns the DateHistogramGrouping struct
func (rb *DateHistogramGroupingBuilder) Build() DateHistogramGrouping {
	return *rb.v
}

// CalendarInterval set the CalendarInterval property for DateHistogramGroupingBuilder.
func (rb *DateHistogramGroupingBuilder) CalendarInterval(calendarinterval Time) *DateHistogramGroupingBuilder {
	rb.v.CalendarInterval = &calendarinterval
	return rb
}

// Delay set the Delay property for DateHistogramGroupingBuilder.
func (rb *DateHistogramGroupingBuilder) Delay(delay Time) *DateHistogramGroupingBuilder {
	rb.v.Delay = &delay
	return rb
}

// Field set the Field property for DateHistogramGroupingBuilder.
func (rb *DateHistogramGroupingBuilder) Field(field Field) *DateHistogramGroupingBuilder {
	rb.v.Field = field
	return rb
}

// FixedInterval set the FixedInterval property for DateHistogramGroupingBuilder.
func (rb *DateHistogramGroupingBuilder) FixedInterval(fixedinterval Time) *DateHistogramGroupingBuilder {
	rb.v.FixedInterval = &fixedinterval
	return rb
}

// Format set the Format property for DateHistogramGroupingBuilder.
func (rb *DateHistogramGroupingBuilder) Format(format string) *DateHistogramGroupingBuilder {
	rb.v.Format = &format
	return rb
}

// Interval set the Interval property for DateHistogramGroupingBuilder.
func (rb *DateHistogramGroupingBuilder) Interval(interval Time) *DateHistogramGroupingBuilder {
	rb.v.Interval = &interval
	return rb
}

// TimeZone set the TimeZone property for DateHistogramGroupingBuilder.
func (rb *DateHistogramGroupingBuilder) TimeZone(timezone string) *DateHistogramGroupingBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
