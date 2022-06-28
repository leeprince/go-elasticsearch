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

// IndexSettingsTimeSeries type.
type IndexSettingsTimeSeries struct {
	EndTime   *DateOrEpochMillis `json:"end_time,omitempty"`
	StartTime *DateOrEpochMillis `json:"start_time,omitempty"`
}

// IndexSettingsTimeSeriesBuilder holds IndexSettingsTimeSeries struct and provides a builder API.
type IndexSettingsTimeSeriesBuilder struct {
	v *IndexSettingsTimeSeries
}

// NewIndexSettingsTimeSeries provides a builder for the IndexSettingsTimeSeries struct.
func NewIndexSettingsTimeSeries() *IndexSettingsTimeSeriesBuilder {
	r := IndexSettingsTimeSeriesBuilder{
		&IndexSettingsTimeSeries{},
	}

	return &r
}

// Build finalize the chain and returns the IndexSettingsTimeSeries struct
func (rb *IndexSettingsTimeSeriesBuilder) Build() IndexSettingsTimeSeries {
	return *rb.v
}

// EndTime set the EndTime property for IndexSettingsTimeSeriesBuilder.
func (rb *IndexSettingsTimeSeriesBuilder) EndTime(endtime DateOrEpochMillis) *IndexSettingsTimeSeriesBuilder {
	rb.v.EndTime = &endtime
	return rb
}

// StartTime set the StartTime property for IndexSettingsTimeSeriesBuilder.
func (rb *IndexSettingsTimeSeriesBuilder) StartTime(starttime DateOrEpochMillis) *IndexSettingsTimeSeriesBuilder {
	rb.v.StartTime = &starttime
	return rb
}
