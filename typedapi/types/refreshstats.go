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

// RefreshStats type.
type RefreshStats struct {
	ExternalTotal             int64   `json:"external_total"`
	ExternalTotalTimeInMillis int64   `json:"external_total_time_in_millis"`
	Listeners                 int64   `json:"listeners"`
	Total                     int64   `json:"total"`
	TotalTime                 *string `json:"total_time,omitempty"`
	TotalTimeInMillis         int64   `json:"total_time_in_millis"`
}

// RefreshStatsBuilder holds RefreshStats struct and provides a builder API.
type RefreshStatsBuilder struct {
	v *RefreshStats
}

// NewRefreshStats provides a builder for the RefreshStats struct.
func NewRefreshStats() *RefreshStatsBuilder {
	r := RefreshStatsBuilder{
		&RefreshStats{},
	}

	return &r
}

// Build finalize the chain and returns the RefreshStats struct
func (rb *RefreshStatsBuilder) Build() RefreshStats {
	return *rb.v
}

// ExternalTotal set the ExternalTotal property for RefreshStatsBuilder.
func (rb *RefreshStatsBuilder) ExternalTotal(externaltotal int64) *RefreshStatsBuilder {
	rb.v.ExternalTotal = externaltotal
	return rb
}

// ExternalTotalTimeInMillis set the ExternalTotalTimeInMillis property for RefreshStatsBuilder.
func (rb *RefreshStatsBuilder) ExternalTotalTimeInMillis(externaltotaltimeinmillis int64) *RefreshStatsBuilder {
	rb.v.ExternalTotalTimeInMillis = externaltotaltimeinmillis
	return rb
}

// Listeners set the Listeners property for RefreshStatsBuilder.
func (rb *RefreshStatsBuilder) Listeners(listeners int64) *RefreshStatsBuilder {
	rb.v.Listeners = listeners
	return rb
}

// Total set the Total property for RefreshStatsBuilder.
func (rb *RefreshStatsBuilder) Total(total int64) *RefreshStatsBuilder {
	rb.v.Total = total
	return rb
}

// TotalTime set the TotalTime property for RefreshStatsBuilder.
func (rb *RefreshStatsBuilder) TotalTime(totaltime string) *RefreshStatsBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

// TotalTimeInMillis set the TotalTimeInMillis property for RefreshStatsBuilder.
func (rb *RefreshStatsBuilder) TotalTimeInMillis(totaltimeinmillis int64) *RefreshStatsBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
