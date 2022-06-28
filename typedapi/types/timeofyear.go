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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/month"
)

// TimeOfYear type.
type TimeOfYear struct {
	At  []string      `json:"at"`
	Int []month.Month `json:"int"`
	On  []int         `json:"on"`
}

// TimeOfYearBuilder holds TimeOfYear struct and provides a builder API.
type TimeOfYearBuilder struct {
	v *TimeOfYear
}

// NewTimeOfYear provides a builder for the TimeOfYear struct.
func NewTimeOfYear() *TimeOfYearBuilder {
	r := TimeOfYearBuilder{
		&TimeOfYear{},
	}

	return &r
}

// Build finalize the chain and returns the TimeOfYear struct
func (rb *TimeOfYearBuilder) Build() TimeOfYear {
	return *rb.v
}

// At set the At property for TimeOfYearBuilder.
func (rb *TimeOfYearBuilder) At(at ...string) *TimeOfYearBuilder {
	rb.v.At = append(rb.v.At, at...)
	return rb
}

// Int set the Int property for TimeOfYearBuilder.
func (rb *TimeOfYearBuilder) Int(int ...month.Month) *TimeOfYearBuilder {
	rb.v.Int = append(rb.v.Int, int...)
	return rb
}

// On set the On property for TimeOfYearBuilder.
func (rb *TimeOfYearBuilder) On(on ...int) *TimeOfYearBuilder {
	rb.v.On = append(rb.v.On, on...)
	return rb
}
