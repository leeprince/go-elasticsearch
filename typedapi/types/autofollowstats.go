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

// AutoFollowStats type.
type AutoFollowStats struct {
	AutoFollowedClusters                     []AutoFollowedCluster `json:"auto_followed_clusters"`
	NumberOfFailedFollowIndices              int64                 `json:"number_of_failed_follow_indices"`
	NumberOfFailedRemoteClusterStateRequests int64                 `json:"number_of_failed_remote_cluster_state_requests"`
	NumberOfSuccessfulFollowIndices          int64                 `json:"number_of_successful_follow_indices"`
	RecentAutoFollowErrors                   []ErrorCause          `json:"recent_auto_follow_errors"`
}

// AutoFollowStatsBuilder holds AutoFollowStats struct and provides a builder API.
type AutoFollowStatsBuilder struct {
	v *AutoFollowStats
}

// NewAutoFollowStats provides a builder for the AutoFollowStats struct.
func NewAutoFollowStats() *AutoFollowStatsBuilder {
	r := AutoFollowStatsBuilder{
		&AutoFollowStats{},
	}

	return &r
}

// Build finalize the chain and returns the AutoFollowStats struct
func (rb *AutoFollowStatsBuilder) Build() AutoFollowStats {
	return *rb.v
}

// AutoFollowedClusters set the AutoFollowedClusters property for AutoFollowStatsBuilder.
func (rb *AutoFollowStatsBuilder) AutoFollowedClusters(auto_followed_clusters ...AutoFollowedCluster) *AutoFollowStatsBuilder {
	rb.v.AutoFollowedClusters = append(rb.v.AutoFollowedClusters, auto_followed_clusters...)
	return rb
}

// NumberOfFailedFollowIndices set the NumberOfFailedFollowIndices property for AutoFollowStatsBuilder.
func (rb *AutoFollowStatsBuilder) NumberOfFailedFollowIndices(numberoffailedfollowindices int64) *AutoFollowStatsBuilder {
	rb.v.NumberOfFailedFollowIndices = numberoffailedfollowindices
	return rb
}

// NumberOfFailedRemoteClusterStateRequests set the NumberOfFailedRemoteClusterStateRequests property for AutoFollowStatsBuilder.
func (rb *AutoFollowStatsBuilder) NumberOfFailedRemoteClusterStateRequests(numberoffailedremoteclusterstaterequests int64) *AutoFollowStatsBuilder {
	rb.v.NumberOfFailedRemoteClusterStateRequests = numberoffailedremoteclusterstaterequests
	return rb
}

// NumberOfSuccessfulFollowIndices set the NumberOfSuccessfulFollowIndices property for AutoFollowStatsBuilder.
func (rb *AutoFollowStatsBuilder) NumberOfSuccessfulFollowIndices(numberofsuccessfulfollowindices int64) *AutoFollowStatsBuilder {
	rb.v.NumberOfSuccessfulFollowIndices = numberofsuccessfulfollowindices
	return rb
}

// RecentAutoFollowErrors set the RecentAutoFollowErrors property for AutoFollowStatsBuilder.
func (rb *AutoFollowStatsBuilder) RecentAutoFollowErrors(recent_auto_follow_errors ...ErrorCause) *AutoFollowStatsBuilder {
	rb.v.RecentAutoFollowErrors = append(rb.v.RecentAutoFollowErrors, recent_auto_follow_errors...)
	return rb
}
