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

// UsageStatsShards type.
type UsageStatsShards struct {
	Routing                 ShardRouting `json:"routing"`
	Stats                   ShardsStats  `json:"stats"`
	TrackingId              string       `json:"tracking_id"`
	TrackingStartedAtMillis EpochMillis  `json:"tracking_started_at_millis"`
}

// UsageStatsShardsBuilder holds UsageStatsShards struct and provides a builder API.
type UsageStatsShardsBuilder struct {
	v *UsageStatsShards
}

// NewUsageStatsShards provides a builder for the UsageStatsShards struct.
func NewUsageStatsShards() *UsageStatsShardsBuilder {
	r := UsageStatsShardsBuilder{
		&UsageStatsShards{},
	}

	return &r
}

// Build finalize the chain and returns the UsageStatsShards struct
func (rb *UsageStatsShardsBuilder) Build() UsageStatsShards {
	return *rb.v
}

// Routing set the Routing property for UsageStatsShardsBuilder.
func (rb *UsageStatsShardsBuilder) Routing(routing ShardRouting) *UsageStatsShardsBuilder {
	rb.v.Routing = routing
	return rb
}

// Stats set the Stats property for UsageStatsShardsBuilder.
func (rb *UsageStatsShardsBuilder) Stats(stats ShardsStats) *UsageStatsShardsBuilder {
	rb.v.Stats = stats
	return rb
}

// TrackingId set the TrackingId property for UsageStatsShardsBuilder.
func (rb *UsageStatsShardsBuilder) TrackingId(trackingid string) *UsageStatsShardsBuilder {
	rb.v.TrackingId = trackingid
	return rb
}

// TrackingStartedAtMillis set the TrackingStartedAtMillis property for UsageStatsShardsBuilder.
func (rb *UsageStatsShardsBuilder) TrackingStartedAtMillis(trackingstartedatmillis EpochMillis) *UsageStatsShardsBuilder {
	rb.v.TrackingStartedAtMillis = trackingstartedatmillis
	return rb
}
