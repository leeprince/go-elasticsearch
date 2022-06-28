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

// RollupJobSummaryField type.
type RollupJobSummaryField struct {
	Agg              string  `json:"agg"`
	CalendarInterval *Time   `json:"calendar_interval,omitempty"`
	TimeZone         *string `json:"time_zone,omitempty"`
}

// RollupJobSummaryFieldBuilder holds RollupJobSummaryField struct and provides a builder API.
type RollupJobSummaryFieldBuilder struct {
	v *RollupJobSummaryField
}

// NewRollupJobSummaryField provides a builder for the RollupJobSummaryField struct.
func NewRollupJobSummaryField() *RollupJobSummaryFieldBuilder {
	r := RollupJobSummaryFieldBuilder{
		&RollupJobSummaryField{},
	}

	return &r
}

// Build finalize the chain and returns the RollupJobSummaryField struct
func (rb *RollupJobSummaryFieldBuilder) Build() RollupJobSummaryField {
	return *rb.v
}

// Agg set the Agg property for RollupJobSummaryFieldBuilder.
func (rb *RollupJobSummaryFieldBuilder) Agg(agg string) *RollupJobSummaryFieldBuilder {
	rb.v.Agg = agg
	return rb
}

// CalendarInterval set the CalendarInterval property for RollupJobSummaryFieldBuilder.
func (rb *RollupJobSummaryFieldBuilder) CalendarInterval(calendarinterval Time) *RollupJobSummaryFieldBuilder {
	rb.v.CalendarInterval = &calendarinterval
	return rb
}

// TimeZone set the TimeZone property for RollupJobSummaryFieldBuilder.
func (rb *RollupJobSummaryFieldBuilder) TimeZone(timezone string) *RollupJobSummaryFieldBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
