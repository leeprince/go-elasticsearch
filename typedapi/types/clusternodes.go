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

// ClusterNodes type.
type ClusterNodes struct {
	// Count Contains counts for nodes selected by the request’s node filters.
	Count ClusterNodeCount `json:"count"`
	// DiscoveryTypes Contains statistics about the discovery types used by selected nodes.
	DiscoveryTypes map[string]int `json:"discovery_types"`
	// Fs Contains statistics about file stores by selected nodes.
	Fs               ClusterFileSystem `json:"fs"`
	IndexingPressure IndexingPressure  `json:"indexing_pressure"`
	Ingest           ClusterIngest     `json:"ingest"`
	// Jvm Contains statistics about the Java Virtual Machines (JVMs) used by selected
	// nodes.
	Jvm ClusterJvm `json:"jvm"`
	// NetworkTypes Contains statistics about the transport and HTTP networks used by selected
	// nodes.
	NetworkTypes ClusterNetworkTypes `json:"network_types"`
	// Os Contains statistics about the operating systems used by selected nodes.
	Os ClusterOperatingSystem `json:"os"`
	// PackagingTypes Contains statistics about Elasticsearch distributions installed on selected
	// nodes.
	PackagingTypes []NodePackagingType `json:"packaging_types"`
	// Plugins Contains statistics about installed plugins and modules by selected nodes.
	Plugins []PluginStats `json:"plugins"`
	// Process Contains statistics about processes used by selected nodes.
	Process ClusterProcess `json:"process"`
	// Versions Array of Elasticsearch versions used on selected nodes.
	Versions []VersionString `json:"versions"`
}

// ClusterNodesBuilder holds ClusterNodes struct and provides a builder API.
type ClusterNodesBuilder struct {
	v *ClusterNodes
}

// NewClusterNodes provides a builder for the ClusterNodes struct.
func NewClusterNodes() *ClusterNodesBuilder {
	r := ClusterNodesBuilder{
		&ClusterNodes{
			DiscoveryTypes: make(map[string]int, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ClusterNodes struct
func (rb *ClusterNodesBuilder) Build() ClusterNodes {
	return *rb.v
}

// Count Contains counts for nodes selected by the request’s node filters.
func (rb *ClusterNodesBuilder) Count(count ClusterNodeCount) *ClusterNodesBuilder {
	rb.v.Count = count
	return rb
}

// DiscoveryTypes Contains statistics about the discovery types used by selected nodes.
func (rb *ClusterNodesBuilder) DiscoveryTypes(value map[string]int) *ClusterNodesBuilder {
	rb.v.DiscoveryTypes = value
	return rb
}

// Fs Contains statistics about file stores by selected nodes.
func (rb *ClusterNodesBuilder) Fs(fs ClusterFileSystem) *ClusterNodesBuilder {
	rb.v.Fs = fs
	return rb
}

// IndexingPressure set the IndexingPressure property for ClusterNodesBuilder.
func (rb *ClusterNodesBuilder) IndexingPressure(indexingpressure IndexingPressure) *ClusterNodesBuilder {
	rb.v.IndexingPressure = indexingpressure
	return rb
}

// Ingest set the Ingest property for ClusterNodesBuilder.
func (rb *ClusterNodesBuilder) Ingest(ingest ClusterIngest) *ClusterNodesBuilder {
	rb.v.Ingest = ingest
	return rb
}

// Jvm Contains statistics about the Java Virtual Machines (JVMs) used by selected
// nodes.
func (rb *ClusterNodesBuilder) Jvm(jvm ClusterJvm) *ClusterNodesBuilder {
	rb.v.Jvm = jvm
	return rb
}

// NetworkTypes Contains statistics about the transport and HTTP networks used by selected
// nodes.
func (rb *ClusterNodesBuilder) NetworkTypes(networktypes ClusterNetworkTypes) *ClusterNodesBuilder {
	rb.v.NetworkTypes = networktypes
	return rb
}

// Os Contains statistics about the operating systems used by selected nodes.
func (rb *ClusterNodesBuilder) Os(os ClusterOperatingSystem) *ClusterNodesBuilder {
	rb.v.Os = os
	return rb
}

// PackagingTypes Contains statistics about Elasticsearch distributions installed on selected
// nodes.
func (rb *ClusterNodesBuilder) PackagingTypes(packaging_types ...NodePackagingType) *ClusterNodesBuilder {
	rb.v.PackagingTypes = append(rb.v.PackagingTypes, packaging_types...)
	return rb
}

// Plugins Contains statistics about installed plugins and modules by selected nodes.
func (rb *ClusterNodesBuilder) Plugins(plugins ...PluginStats) *ClusterNodesBuilder {
	rb.v.Plugins = append(rb.v.Plugins, plugins...)
	return rb
}

// Process Contains statistics about processes used by selected nodes.
func (rb *ClusterNodesBuilder) Process(process ClusterProcess) *ClusterNodesBuilder {
	rb.v.Process = process
	return rb
}

// Versions Array of Elasticsearch versions used on selected nodes.
func (rb *ClusterNodesBuilder) Versions(versions ...VersionString) *ClusterNodesBuilder {
	rb.v.Versions = append(rb.v.Versions, versions...)
	return rb
}
