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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/healthstatus"
)

// IndexHealthStats type.
type IndexHealthStats struct {
	ActivePrimaryShards int                         `json:"active_primary_shards"`
	ActiveShards        int                         `json:"active_shards"`
	InitializingShards  int                         `json:"initializing_shards"`
	NumberOfReplicas    int                         `json:"number_of_replicas"`
	NumberOfShards      int                         `json:"number_of_shards"`
	RelocatingShards    int                         `json:"relocating_shards"`
	Shards              map[string]ShardHealthStats `json:"shards,omitempty"`
	Status              healthstatus.HealthStatus   `json:"status"`
	UnassignedShards    int                         `json:"unassigned_shards"`
}

// IndexHealthStatsBuilder holds IndexHealthStats struct and provides a builder API.
type IndexHealthStatsBuilder struct {
	v *IndexHealthStats
}

// NewIndexHealthStats provides a builder for the IndexHealthStats struct.
func NewIndexHealthStats() *IndexHealthStatsBuilder {
	r := IndexHealthStatsBuilder{
		&IndexHealthStats{
			Shards: make(map[string]ShardHealthStats, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexHealthStats struct
func (rb *IndexHealthStatsBuilder) Build() IndexHealthStats {
	return *rb.v
}

// ActivePrimaryShards set the ActivePrimaryShards property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) ActivePrimaryShards(activeprimaryshards int) *IndexHealthStatsBuilder {
	rb.v.ActivePrimaryShards = activeprimaryshards
	return rb
}

// ActiveShards set the ActiveShards property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) ActiveShards(activeshards int) *IndexHealthStatsBuilder {
	rb.v.ActiveShards = activeshards
	return rb
}

// InitializingShards set the InitializingShards property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) InitializingShards(initializingshards int) *IndexHealthStatsBuilder {
	rb.v.InitializingShards = initializingshards
	return rb
}

// NumberOfReplicas set the NumberOfReplicas property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) NumberOfReplicas(numberofreplicas int) *IndexHealthStatsBuilder {
	rb.v.NumberOfReplicas = numberofreplicas
	return rb
}

// NumberOfShards set the NumberOfShards property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) NumberOfShards(numberofshards int) *IndexHealthStatsBuilder {
	rb.v.NumberOfShards = numberofshards
	return rb
}

// RelocatingShards set the RelocatingShards property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) RelocatingShards(relocatingshards int) *IndexHealthStatsBuilder {
	rb.v.RelocatingShards = relocatingshards
	return rb
}

// Shards set the Shards property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) Shards(value map[string]ShardHealthStats) *IndexHealthStatsBuilder {
	rb.v.Shards = value
	return rb
}

// Status set the Status property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) Status(status healthstatus.HealthStatus) *IndexHealthStatsBuilder {
	rb.v.Status = status
	return rb
}

// UnassignedShards set the UnassignedShards property for IndexHealthStatsBuilder.
func (rb *IndexHealthStatsBuilder) UnassignedShards(unassignedshards int) *IndexHealthStatsBuilder {
	rb.v.UnassignedShards = unassignedshards
	return rb
}
