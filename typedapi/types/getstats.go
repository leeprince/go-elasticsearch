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

// GetStats type.
type GetStats struct {
	Current             int64   `json:"current"`
	ExistsTime          *string `json:"exists_time,omitempty"`
	ExistsTimeInMillis  int64   `json:"exists_time_in_millis"`
	ExistsTotal         int64   `json:"exists_total"`
	MissingTime         *string `json:"missing_time,omitempty"`
	MissingTimeInMillis int64   `json:"missing_time_in_millis"`
	MissingTotal        int64   `json:"missing_total"`
	Time                *string `json:"time,omitempty"`
	TimeInMillis        int64   `json:"time_in_millis"`
	Total               int64   `json:"total"`
}

// GetStatsBuilder holds GetStats struct and provides a builder API.
type GetStatsBuilder struct {
	v *GetStats
}

// NewGetStats provides a builder for the GetStats struct.
func NewGetStats() *GetStatsBuilder {
	r := GetStatsBuilder{
		&GetStats{},
	}

	return &r
}

// Build finalize the chain and returns the GetStats struct
func (rb *GetStatsBuilder) Build() GetStats {
	return *rb.v
}

// Current set the Current property for GetStatsBuilder.
func (rb *GetStatsBuilder) Current(current int64) *GetStatsBuilder {
	rb.v.Current = current
	return rb
}

// ExistsTime set the ExistsTime property for GetStatsBuilder.
func (rb *GetStatsBuilder) ExistsTime(existstime string) *GetStatsBuilder {
	rb.v.ExistsTime = &existstime
	return rb
}

// ExistsTimeInMillis set the ExistsTimeInMillis property for GetStatsBuilder.
func (rb *GetStatsBuilder) ExistsTimeInMillis(existstimeinmillis int64) *GetStatsBuilder {
	rb.v.ExistsTimeInMillis = existstimeinmillis
	return rb
}

// ExistsTotal set the ExistsTotal property for GetStatsBuilder.
func (rb *GetStatsBuilder) ExistsTotal(existstotal int64) *GetStatsBuilder {
	rb.v.ExistsTotal = existstotal
	return rb
}

// MissingTime set the MissingTime property for GetStatsBuilder.
func (rb *GetStatsBuilder) MissingTime(missingtime string) *GetStatsBuilder {
	rb.v.MissingTime = &missingtime
	return rb
}

// MissingTimeInMillis set the MissingTimeInMillis property for GetStatsBuilder.
func (rb *GetStatsBuilder) MissingTimeInMillis(missingtimeinmillis int64) *GetStatsBuilder {
	rb.v.MissingTimeInMillis = missingtimeinmillis
	return rb
}

// MissingTotal set the MissingTotal property for GetStatsBuilder.
func (rb *GetStatsBuilder) MissingTotal(missingtotal int64) *GetStatsBuilder {
	rb.v.MissingTotal = missingtotal
	return rb
}

// Time set the Time property for GetStatsBuilder.
func (rb *GetStatsBuilder) Time(time string) *GetStatsBuilder {
	rb.v.Time = &time
	return rb
}

// TimeInMillis set the TimeInMillis property for GetStatsBuilder.
func (rb *GetStatsBuilder) TimeInMillis(timeinmillis int64) *GetStatsBuilder {
	rb.v.TimeInMillis = timeinmillis
	return rb
}

// Total set the Total property for GetStatsBuilder.
func (rb *GetStatsBuilder) Total(total int64) *GetStatsBuilder {
	rb.v.Total = total
	return rb
}
