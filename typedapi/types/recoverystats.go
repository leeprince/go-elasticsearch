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

// RecoveryStats type.
type RecoveryStats struct {
	CurrentAsSource      int64   `json:"current_as_source"`
	CurrentAsTarget      int64   `json:"current_as_target"`
	ThrottleTime         *string `json:"throttle_time,omitempty"`
	ThrottleTimeInMillis int64   `json:"throttle_time_in_millis"`
}

// RecoveryStatsBuilder holds RecoveryStats struct and provides a builder API.
type RecoveryStatsBuilder struct {
	v *RecoveryStats
}

// NewRecoveryStats provides a builder for the RecoveryStats struct.
func NewRecoveryStats() *RecoveryStatsBuilder {
	r := RecoveryStatsBuilder{
		&RecoveryStats{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryStats struct
func (rb *RecoveryStatsBuilder) Build() RecoveryStats {
	return *rb.v
}

// CurrentAsSource set the CurrentAsSource property for RecoveryStatsBuilder.
func (rb *RecoveryStatsBuilder) CurrentAsSource(currentassource int64) *RecoveryStatsBuilder {
	rb.v.CurrentAsSource = currentassource
	return rb
}

// CurrentAsTarget set the CurrentAsTarget property for RecoveryStatsBuilder.
func (rb *RecoveryStatsBuilder) CurrentAsTarget(currentastarget int64) *RecoveryStatsBuilder {
	rb.v.CurrentAsTarget = currentastarget
	return rb
}

// ThrottleTime set the ThrottleTime property for RecoveryStatsBuilder.
func (rb *RecoveryStatsBuilder) ThrottleTime(throttletime string) *RecoveryStatsBuilder {
	rb.v.ThrottleTime = &throttletime
	return rb
}

// ThrottleTimeInMillis set the ThrottleTimeInMillis property for RecoveryStatsBuilder.
func (rb *RecoveryStatsBuilder) ThrottleTimeInMillis(throttletimeinmillis int64) *RecoveryStatsBuilder {
	rb.v.ThrottleTimeInMillis = throttletimeinmillis
	return rb
}
