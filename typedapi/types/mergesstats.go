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

// MergesStats type.
type MergesStats struct {
	Current                    int64   `json:"current"`
	CurrentDocs                int64   `json:"current_docs"`
	CurrentSize                *string `json:"current_size,omitempty"`
	CurrentSizeInBytes         int64   `json:"current_size_in_bytes"`
	Total                      int64   `json:"total"`
	TotalAutoThrottle          *string `json:"total_auto_throttle,omitempty"`
	TotalAutoThrottleInBytes   int64   `json:"total_auto_throttle_in_bytes"`
	TotalDocs                  int64   `json:"total_docs"`
	TotalSize                  *string `json:"total_size,omitempty"`
	TotalSizeInBytes           int64   `json:"total_size_in_bytes"`
	TotalStoppedTime           *string `json:"total_stopped_time,omitempty"`
	TotalStoppedTimeInMillis   int64   `json:"total_stopped_time_in_millis"`
	TotalThrottledTime         *string `json:"total_throttled_time,omitempty"`
	TotalThrottledTimeInMillis int64   `json:"total_throttled_time_in_millis"`
	TotalTime                  *string `json:"total_time,omitempty"`
	TotalTimeInMillis          int64   `json:"total_time_in_millis"`
}

// MergesStatsBuilder holds MergesStats struct and provides a builder API.
type MergesStatsBuilder struct {
	v *MergesStats
}

// NewMergesStats provides a builder for the MergesStats struct.
func NewMergesStats() *MergesStatsBuilder {
	r := MergesStatsBuilder{
		&MergesStats{},
	}

	return &r
}

// Build finalize the chain and returns the MergesStats struct
func (rb *MergesStatsBuilder) Build() MergesStats {
	return *rb.v
}

// Current set the Current property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) Current(current int64) *MergesStatsBuilder {
	rb.v.Current = current
	return rb
}

// CurrentDocs set the CurrentDocs property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) CurrentDocs(currentdocs int64) *MergesStatsBuilder {
	rb.v.CurrentDocs = currentdocs
	return rb
}

// CurrentSize set the CurrentSize property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) CurrentSize(currentsize string) *MergesStatsBuilder {
	rb.v.CurrentSize = &currentsize
	return rb
}

// CurrentSizeInBytes set the CurrentSizeInBytes property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) CurrentSizeInBytes(currentsizeinbytes int64) *MergesStatsBuilder {
	rb.v.CurrentSizeInBytes = currentsizeinbytes
	return rb
}

// Total set the Total property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) Total(total int64) *MergesStatsBuilder {
	rb.v.Total = total
	return rb
}

// TotalAutoThrottle set the TotalAutoThrottle property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalAutoThrottle(totalautothrottle string) *MergesStatsBuilder {
	rb.v.TotalAutoThrottle = &totalautothrottle
	return rb
}

// TotalAutoThrottleInBytes set the TotalAutoThrottleInBytes property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalAutoThrottleInBytes(totalautothrottleinbytes int64) *MergesStatsBuilder {
	rb.v.TotalAutoThrottleInBytes = totalautothrottleinbytes
	return rb
}

// TotalDocs set the TotalDocs property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalDocs(totaldocs int64) *MergesStatsBuilder {
	rb.v.TotalDocs = totaldocs
	return rb
}

// TotalSize set the TotalSize property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalSize(totalsize string) *MergesStatsBuilder {
	rb.v.TotalSize = &totalsize
	return rb
}

// TotalSizeInBytes set the TotalSizeInBytes property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalSizeInBytes(totalsizeinbytes int64) *MergesStatsBuilder {
	rb.v.TotalSizeInBytes = totalsizeinbytes
	return rb
}

// TotalStoppedTime set the TotalStoppedTime property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalStoppedTime(totalstoppedtime string) *MergesStatsBuilder {
	rb.v.TotalStoppedTime = &totalstoppedtime
	return rb
}

// TotalStoppedTimeInMillis set the TotalStoppedTimeInMillis property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalStoppedTimeInMillis(totalstoppedtimeinmillis int64) *MergesStatsBuilder {
	rb.v.TotalStoppedTimeInMillis = totalstoppedtimeinmillis
	return rb
}

// TotalThrottledTime set the TotalThrottledTime property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalThrottledTime(totalthrottledtime string) *MergesStatsBuilder {
	rb.v.TotalThrottledTime = &totalthrottledtime
	return rb
}

// TotalThrottledTimeInMillis set the TotalThrottledTimeInMillis property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalThrottledTimeInMillis(totalthrottledtimeinmillis int64) *MergesStatsBuilder {
	rb.v.TotalThrottledTimeInMillis = totalthrottledtimeinmillis
	return rb
}

// TotalTime set the TotalTime property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalTime(totaltime string) *MergesStatsBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

// TotalTimeInMillis set the TotalTimeInMillis property for MergesStatsBuilder.
func (rb *MergesStatsBuilder) TotalTimeInMillis(totaltimeinmillis int64) *MergesStatsBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
