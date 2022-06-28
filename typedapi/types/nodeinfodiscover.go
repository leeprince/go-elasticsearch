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

// NodeInfoDiscover type.
type NodeInfoDiscover struct {
	SeedHosts string `json:"seed_hosts"`
}

// NodeInfoDiscoverBuilder holds NodeInfoDiscover struct and provides a builder API.
type NodeInfoDiscoverBuilder struct {
	v *NodeInfoDiscover
}

// NewNodeInfoDiscover provides a builder for the NodeInfoDiscover struct.
func NewNodeInfoDiscover() *NodeInfoDiscoverBuilder {
	r := NodeInfoDiscoverBuilder{
		&NodeInfoDiscover{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoDiscover struct
func (rb *NodeInfoDiscoverBuilder) Build() NodeInfoDiscover {
	return *rb.v
}

// SeedHosts set the SeedHosts property for NodeInfoDiscoverBuilder.
func (rb *NodeInfoDiscoverBuilder) SeedHosts(seedhosts string) *NodeInfoDiscoverBuilder {
	rb.v.SeedHosts = seedhosts
	return rb
}
