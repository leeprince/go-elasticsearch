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

// TranslogStatus type.
type TranslogStatus struct {
	Percent           Percentage  `json:"percent"`
	Recovered         int64       `json:"recovered"`
	Total             int64       `json:"total"`
	TotalOnStart      int64       `json:"total_on_start"`
	TotalTime         *string     `json:"total_time,omitempty"`
	TotalTimeInMillis EpochMillis `json:"total_time_in_millis"`
}

// TranslogStatusBuilder holds TranslogStatus struct and provides a builder API.
type TranslogStatusBuilder struct {
	v *TranslogStatus
}

// NewTranslogStatus provides a builder for the TranslogStatus struct.
func NewTranslogStatus() *TranslogStatusBuilder {
	r := TranslogStatusBuilder{
		&TranslogStatus{},
	}

	return &r
}

// Build finalize the chain and returns the TranslogStatus struct
func (rb *TranslogStatusBuilder) Build() TranslogStatus {
	return *rb.v
}

// Percent set the Percent property for TranslogStatusBuilder.
func (rb *TranslogStatusBuilder) Percent(percent Percentage) *TranslogStatusBuilder {
	rb.v.Percent = percent
	return rb
}

// Recovered set the Recovered property for TranslogStatusBuilder.
func (rb *TranslogStatusBuilder) Recovered(recovered int64) *TranslogStatusBuilder {
	rb.v.Recovered = recovered
	return rb
}

// Total set the Total property for TranslogStatusBuilder.
func (rb *TranslogStatusBuilder) Total(total int64) *TranslogStatusBuilder {
	rb.v.Total = total
	return rb
}

// TotalOnStart set the TotalOnStart property for TranslogStatusBuilder.
func (rb *TranslogStatusBuilder) TotalOnStart(totalonstart int64) *TranslogStatusBuilder {
	rb.v.TotalOnStart = totalonstart
	return rb
}

// TotalTime set the TotalTime property for TranslogStatusBuilder.
func (rb *TranslogStatusBuilder) TotalTime(totaltime string) *TranslogStatusBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

// TotalTimeInMillis set the TotalTimeInMillis property for TranslogStatusBuilder.
func (rb *TranslogStatusBuilder) TotalTimeInMillis(totaltimeinmillis EpochMillis) *TranslogStatusBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
