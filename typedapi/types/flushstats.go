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

// FlushStats type.
type FlushStats struct {
	Periodic          int64   `json:"periodic"`
	Total             int64   `json:"total"`
	TotalTime         *string `json:"total_time,omitempty"`
	TotalTimeInMillis int64   `json:"total_time_in_millis"`
}

// FlushStatsBuilder holds FlushStats struct and provides a builder API.
type FlushStatsBuilder struct {
	v *FlushStats
}

// NewFlushStats provides a builder for the FlushStats struct.
func NewFlushStats() *FlushStatsBuilder {
	r := FlushStatsBuilder{
		&FlushStats{},
	}

	return &r
}

// Build finalize the chain and returns the FlushStats struct
func (rb *FlushStatsBuilder) Build() FlushStats {
	return *rb.v
}

// Periodic set the Periodic property for FlushStatsBuilder.
func (rb *FlushStatsBuilder) Periodic(periodic int64) *FlushStatsBuilder {
	rb.v.Periodic = periodic
	return rb
}

// Total set the Total property for FlushStatsBuilder.
func (rb *FlushStatsBuilder) Total(total int64) *FlushStatsBuilder {
	rb.v.Total = total
	return rb
}

// TotalTime set the TotalTime property for FlushStatsBuilder.
func (rb *FlushStatsBuilder) TotalTime(totaltime string) *FlushStatsBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

// TotalTimeInMillis set the TotalTimeInMillis property for FlushStatsBuilder.
func (rb *FlushStatsBuilder) TotalTimeInMillis(totaltimeinmillis int64) *FlushStatsBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
