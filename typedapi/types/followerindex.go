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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/followerindexstatus"
)

// FollowerIndex type.
type FollowerIndex struct {
	FollowerIndex IndexName                               `json:"follower_index"`
	LeaderIndex   IndexName                               `json:"leader_index"`
	Parameters    *FollowerIndexParameters                `json:"parameters,omitempty"`
	RemoteCluster Name                                    `json:"remote_cluster"`
	Status        followerindexstatus.FollowerIndexStatus `json:"status"`
}

// FollowerIndexBuilder holds FollowerIndex struct and provides a builder API.
type FollowerIndexBuilder struct {
	v *FollowerIndex
}

// NewFollowerIndex provides a builder for the FollowerIndex struct.
func NewFollowerIndex() *FollowerIndexBuilder {
	r := FollowerIndexBuilder{
		&FollowerIndex{},
	}

	return &r
}

// Build finalize the chain and returns the FollowerIndex struct
func (rb *FollowerIndexBuilder) Build() FollowerIndex {
	return *rb.v
}

// FollowerIndex set the FollowerIndex property for FollowerIndexBuilder.
func (rb *FollowerIndexBuilder) FollowerIndex(followerindex IndexName) *FollowerIndexBuilder {
	rb.v.FollowerIndex = followerindex
	return rb
}

// LeaderIndex set the LeaderIndex property for FollowerIndexBuilder.
func (rb *FollowerIndexBuilder) LeaderIndex(leaderindex IndexName) *FollowerIndexBuilder {
	rb.v.LeaderIndex = leaderindex
	return rb
}

// Parameters set the Parameters property for FollowerIndexBuilder.
func (rb *FollowerIndexBuilder) Parameters(parameters FollowerIndexParameters) *FollowerIndexBuilder {
	rb.v.Parameters = &parameters
	return rb
}

// RemoteCluster set the RemoteCluster property for FollowerIndexBuilder.
func (rb *FollowerIndexBuilder) RemoteCluster(remotecluster Name) *FollowerIndexBuilder {
	rb.v.RemoteCluster = remotecluster
	return rb
}

// Status set the Status property for FollowerIndexBuilder.
func (rb *FollowerIndexBuilder) Status(status followerindexstatus.FollowerIndexStatus) *FollowerIndexBuilder {
	rb.v.Status = status
	return rb
}
