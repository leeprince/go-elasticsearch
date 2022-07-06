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

// HourAndMinute type.
type HourAndMinute struct {
	Hour   []int `json:"hour"`
	Minute []int `json:"minute"`
}

// HourAndMinuteBuilder holds HourAndMinute struct and provides a builder API.
type HourAndMinuteBuilder struct {
	v *HourAndMinute
}

// NewHourAndMinute provides a builder for the HourAndMinute struct.
func NewHourAndMinute() *HourAndMinuteBuilder {
	r := HourAndMinuteBuilder{
		&HourAndMinute{},
	}

	return &r
}

// Build finalize the chain and returns the HourAndMinute struct
func (rb *HourAndMinuteBuilder) Build() HourAndMinute {
	return *rb.v
}

// Hour set the Hour property for HourAndMinuteBuilder.
func (rb *HourAndMinuteBuilder) Hour(hour ...int) *HourAndMinuteBuilder {
	rb.v.Hour = append(rb.v.Hour, hour...)
	return rb
}

// Minute set the Minute property for HourAndMinuteBuilder.
func (rb *HourAndMinuteBuilder) Minute(minute ...int) *HourAndMinuteBuilder {
	rb.v.Minute = append(rb.v.Minute, minute...)
	return rb
}
