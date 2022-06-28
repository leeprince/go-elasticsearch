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

// DatafeedTimingStats type.
type DatafeedTimingStats struct {
	AverageSearchTimePerBucketMs          *int    `json:"average_search_time_per_bucket_ms,omitempty"`
	BucketCount                           int64   `json:"bucket_count"`
	ExponentialAverageSearchTimePerHourMs float64 `json:"exponential_average_search_time_per_hour_ms"`
	JobId                                 Id      `json:"job_id"`
	SearchCount                           int64   `json:"search_count"`
	TotalSearchTimeMs                     float64 `json:"total_search_time_ms"`
}

// DatafeedTimingStatsBuilder holds DatafeedTimingStats struct and provides a builder API.
type DatafeedTimingStatsBuilder struct {
	v *DatafeedTimingStats
}

// NewDatafeedTimingStats provides a builder for the DatafeedTimingStats struct.
func NewDatafeedTimingStats() *DatafeedTimingStatsBuilder {
	r := DatafeedTimingStatsBuilder{
		&DatafeedTimingStats{},
	}

	return &r
}

// Build finalize the chain and returns the DatafeedTimingStats struct
func (rb *DatafeedTimingStatsBuilder) Build() DatafeedTimingStats {
	return *rb.v
}

// AverageSearchTimePerBucketMs set the AverageSearchTimePerBucketMs property for DatafeedTimingStatsBuilder.
func (rb *DatafeedTimingStatsBuilder) AverageSearchTimePerBucketMs(averagesearchtimeperbucketms int) *DatafeedTimingStatsBuilder {
	rb.v.AverageSearchTimePerBucketMs = &averagesearchtimeperbucketms
	return rb
}

// BucketCount set the BucketCount property for DatafeedTimingStatsBuilder.
func (rb *DatafeedTimingStatsBuilder) BucketCount(bucketcount int64) *DatafeedTimingStatsBuilder {
	rb.v.BucketCount = bucketcount
	return rb
}

// ExponentialAverageSearchTimePerHourMs set the ExponentialAverageSearchTimePerHourMs property for DatafeedTimingStatsBuilder.
func (rb *DatafeedTimingStatsBuilder) ExponentialAverageSearchTimePerHourMs(exponentialaveragesearchtimeperhourms float64) *DatafeedTimingStatsBuilder {
	rb.v.ExponentialAverageSearchTimePerHourMs = exponentialaveragesearchtimeperhourms
	return rb
}

// JobId set the JobId property for DatafeedTimingStatsBuilder.
func (rb *DatafeedTimingStatsBuilder) JobId(jobid Id) *DatafeedTimingStatsBuilder {
	rb.v.JobId = jobid
	return rb
}

// SearchCount set the SearchCount property for DatafeedTimingStatsBuilder.
func (rb *DatafeedTimingStatsBuilder) SearchCount(searchcount int64) *DatafeedTimingStatsBuilder {
	rb.v.SearchCount = searchcount
	return rb
}

// TotalSearchTimeMs set the TotalSearchTimeMs property for DatafeedTimingStatsBuilder.
func (rb *DatafeedTimingStatsBuilder) TotalSearchTimeMs(totalsearchtimems float64) *DatafeedTimingStatsBuilder {
	rb.v.TotalSearchTimeMs = totalsearchtimems
	return rb
}
