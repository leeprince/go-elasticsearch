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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/day"
)

// TimeOfWeek type.
type TimeOfWeek struct {
	At []string  `json:"at"`
	On []day.Day `json:"on"`
}

// TimeOfWeekBuilder holds TimeOfWeek struct and provides a builder API.
type TimeOfWeekBuilder struct {
	v *TimeOfWeek
}

// NewTimeOfWeek provides a builder for the TimeOfWeek struct.
func NewTimeOfWeek() *TimeOfWeekBuilder {
	r := TimeOfWeekBuilder{
		&TimeOfWeek{},
	}

	return &r
}

// Build finalize the chain and returns the TimeOfWeek struct
func (rb *TimeOfWeekBuilder) Build() TimeOfWeek {
	return *rb.v
}

// At set the At property for TimeOfWeekBuilder.
func (rb *TimeOfWeekBuilder) At(at ...string) *TimeOfWeekBuilder {
	rb.v.At = append(rb.v.At, at...)
	return rb
}

// On set the On property for TimeOfWeekBuilder.
func (rb *TimeOfWeekBuilder) On(on ...day.Day) *TimeOfWeekBuilder {
	rb.v.On = append(rb.v.On, on...)
	return rb
}
