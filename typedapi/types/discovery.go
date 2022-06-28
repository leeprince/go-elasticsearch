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

// Discovery type.
type Discovery struct {
	ClusterApplierStats     *ClusterAppliedStats          `json:"cluster_applier_stats,omitempty"`
	ClusterStateQueue       *ClusterStateQueue            `json:"cluster_state_queue,omitempty"`
	ClusterStateUpdate      map[string]ClusterStateUpdate `json:"cluster_state_update,omitempty"`
	PublishedClusterStates  *PublishedClusterStates       `json:"published_cluster_states,omitempty"`
	SerializedClusterStates *SerializedClusterState       `json:"serialized_cluster_states,omitempty"`
}

// DiscoveryBuilder holds Discovery struct and provides a builder API.
type DiscoveryBuilder struct {
	v *Discovery
}

// NewDiscovery provides a builder for the Discovery struct.
func NewDiscovery() *DiscoveryBuilder {
	r := DiscoveryBuilder{
		&Discovery{
			ClusterStateUpdate: make(map[string]ClusterStateUpdate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Discovery struct
func (rb *DiscoveryBuilder) Build() Discovery {
	return *rb.v
}

// ClusterApplierStats set the ClusterApplierStats property for DiscoveryBuilder.
func (rb *DiscoveryBuilder) ClusterApplierStats(clusterapplierstats ClusterAppliedStats) *DiscoveryBuilder {
	rb.v.ClusterApplierStats = &clusterapplierstats
	return rb
}

// ClusterStateQueue set the ClusterStateQueue property for DiscoveryBuilder.
func (rb *DiscoveryBuilder) ClusterStateQueue(clusterstatequeue ClusterStateQueue) *DiscoveryBuilder {
	rb.v.ClusterStateQueue = &clusterstatequeue
	return rb
}

// ClusterStateUpdate set the ClusterStateUpdate property for DiscoveryBuilder.
func (rb *DiscoveryBuilder) ClusterStateUpdate(value map[string]ClusterStateUpdate) *DiscoveryBuilder {
	rb.v.ClusterStateUpdate = value
	return rb
}

// PublishedClusterStates set the PublishedClusterStates property for DiscoveryBuilder.
func (rb *DiscoveryBuilder) PublishedClusterStates(publishedclusterstates PublishedClusterStates) *DiscoveryBuilder {
	rb.v.PublishedClusterStates = &publishedclusterstates
	return rb
}

// SerializedClusterStates set the SerializedClusterStates property for DiscoveryBuilder.
func (rb *DiscoveryBuilder) SerializedClusterStates(serializedclusterstates SerializedClusterState) *DiscoveryBuilder {
	rb.v.SerializedClusterStates = &serializedclusterstates
	return rb
}
