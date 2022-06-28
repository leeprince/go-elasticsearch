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

// NodeInfo type.
type NodeInfo struct {
	Aggregations map[string]NodeInfoAggregation `json:"aggregations,omitempty"`
	Attributes   map[string]string              `json:"attributes"`
	BuildFlavor  string                         `json:"build_flavor"`
	// BuildHash Short hash of the last git commit in this release.
	BuildHash string `json:"build_hash"`
	BuildType string `json:"build_type"`
	// Host The node’s host name.
	Host   Host            `json:"host"`
	Http   *NodeInfoHttp   `json:"http,omitempty"`
	Ingest *NodeInfoIngest `json:"ingest,omitempty"`
	// Ip The node’s IP address.
	Ip      Ip            `json:"ip"`
	Jvm     *NodeJvmInfo  `json:"jvm,omitempty"`
	Modules []PluginStats `json:"modules,omitempty"`
	// Name The node's name
	Name       Name                          `json:"name"`
	Network    *NodeInfoNetwork              `json:"network,omitempty"`
	Os         *NodeOperatingSystemInfo      `json:"os,omitempty"`
	Plugins    []PluginStats                 `json:"plugins,omitempty"`
	Process    *NodeProcessInfo              `json:"process,omitempty"`
	Roles      NodeRoles                     `json:"roles"`
	Settings   *NodeInfoSettings             `json:"settings,omitempty"`
	ThreadPool map[string]NodeThreadPoolInfo `json:"thread_pool,omitempty"`
	// TotalIndexingBuffer Total heap allowed to be used to hold recently indexed documents before they
	// must be written to disk. This size is a shared pool across all shards on this
	// node, and is controlled by Indexing Buffer settings.
	TotalIndexingBuffer *int64 `json:"total_indexing_buffer,omitempty"`
	// TotalIndexingBufferInBytes Same as total_indexing_buffer, but expressed in bytes.
	TotalIndexingBufferInBytes *ByteSize          `json:"total_indexing_buffer_in_bytes,omitempty"`
	Transport                  *NodeInfoTransport `json:"transport,omitempty"`
	// TransportAddress Host and port where transport HTTP connections are accepted.
	TransportAddress TransportAddress `json:"transport_address"`
	// Version Elasticsearch version running on this node.
	Version VersionString `json:"version"`
}

// NodeInfoBuilder holds NodeInfo struct and provides a builder API.
type NodeInfoBuilder struct {
	v *NodeInfo
}

// NewNodeInfo provides a builder for the NodeInfo struct.
func NewNodeInfo() *NodeInfoBuilder {
	r := NodeInfoBuilder{
		&NodeInfo{
			Aggregations: make(map[string]NodeInfoAggregation, 0),
			Attributes:   make(map[string]string, 0),
			ThreadPool:   make(map[string]NodeThreadPoolInfo, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfo struct
func (rb *NodeInfoBuilder) Build() NodeInfo {
	return *rb.v
}

// Aggregations set the Aggregations property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Aggregations(value map[string]NodeInfoAggregation) *NodeInfoBuilder {
	rb.v.Aggregations = value
	return rb
}

// Attributes set the Attributes property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Attributes(value map[string]string) *NodeInfoBuilder {
	rb.v.Attributes = value
	return rb
}

// BuildFlavor set the BuildFlavor property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) BuildFlavor(buildflavor string) *NodeInfoBuilder {
	rb.v.BuildFlavor = buildflavor
	return rb
}

// BuildHash Short hash of the last git commit in this release.
func (rb *NodeInfoBuilder) BuildHash(buildhash string) *NodeInfoBuilder {
	rb.v.BuildHash = buildhash
	return rb
}

// BuildType set the BuildType property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) BuildType(buildtype string) *NodeInfoBuilder {
	rb.v.BuildType = buildtype
	return rb
}

// Host The node’s host name.
func (rb *NodeInfoBuilder) Host(host Host) *NodeInfoBuilder {
	rb.v.Host = host
	return rb
}

// Http set the Http property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Http(http NodeInfoHttp) *NodeInfoBuilder {
	rb.v.Http = &http
	return rb
}

// Ingest set the Ingest property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Ingest(ingest NodeInfoIngest) *NodeInfoBuilder {
	rb.v.Ingest = &ingest
	return rb
}

// Ip The node’s IP address.
func (rb *NodeInfoBuilder) Ip(ip Ip) *NodeInfoBuilder {
	rb.v.Ip = ip
	return rb
}

// Jvm set the Jvm property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Jvm(jvm NodeJvmInfo) *NodeInfoBuilder {
	rb.v.Jvm = &jvm
	return rb
}

// Modules set the Modules property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Modules(modules ...PluginStats) *NodeInfoBuilder {
	rb.v.Modules = append(rb.v.Modules, modules...)
	return rb
}

// Name The node's name
func (rb *NodeInfoBuilder) Name(name Name) *NodeInfoBuilder {
	rb.v.Name = name
	return rb
}

// Network set the Network property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Network(network NodeInfoNetwork) *NodeInfoBuilder {
	rb.v.Network = &network
	return rb
}

// Os set the Os property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Os(os NodeOperatingSystemInfo) *NodeInfoBuilder {
	rb.v.Os = &os
	return rb
}

// Plugins set the Plugins property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Plugins(plugins ...PluginStats) *NodeInfoBuilder {
	rb.v.Plugins = append(rb.v.Plugins, plugins...)
	return rb
}

// Process set the Process property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Process(process NodeProcessInfo) *NodeInfoBuilder {
	rb.v.Process = &process
	return rb
}

// Roles set the Roles property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Roles(roles NodeRoles) *NodeInfoBuilder {
	rb.v.Roles = roles
	return rb
}

// Settings set the Settings property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Settings(settings NodeInfoSettings) *NodeInfoBuilder {
	rb.v.Settings = &settings
	return rb
}

// ThreadPool set the ThreadPool property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) ThreadPool(value map[string]NodeThreadPoolInfo) *NodeInfoBuilder {
	rb.v.ThreadPool = value
	return rb
}

// TotalIndexingBuffer Total heap allowed to be used to hold recently indexed documents before they
// must be written to disk. This size is a shared pool across all shards on this
// node, and is controlled by Indexing Buffer settings.
func (rb *NodeInfoBuilder) TotalIndexingBuffer(totalindexingbuffer int64) *NodeInfoBuilder {
	rb.v.TotalIndexingBuffer = &totalindexingbuffer
	return rb
}

// TotalIndexingBufferInBytes Same as total_indexing_buffer, but expressed in bytes.
func (rb *NodeInfoBuilder) TotalIndexingBufferInBytes(totalindexingbufferinbytes ByteSize) *NodeInfoBuilder {
	rb.v.TotalIndexingBufferInBytes = &totalindexingbufferinbytes
	return rb
}

// Transport set the Transport property for NodeInfoBuilder.
func (rb *NodeInfoBuilder) Transport(transport NodeInfoTransport) *NodeInfoBuilder {
	rb.v.Transport = &transport
	return rb
}

// TransportAddress Host and port where transport HTTP connections are accepted.
func (rb *NodeInfoBuilder) TransportAddress(transportaddress TransportAddress) *NodeInfoBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}

// Version Elasticsearch version running on this node.
func (rb *NodeInfoBuilder) Version(version VersionString) *NodeInfoBuilder {
	rb.v.Version = version
	return rb
}
