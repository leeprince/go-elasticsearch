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

// NodeJvmInfo type.
type NodeJvmInfo struct {
	BundledJdk                            bool              `json:"bundled_jdk"`
	GcCollectors                          []string          `json:"gc_collectors"`
	InputArguments                        []string          `json:"input_arguments"`
	Mem                                   NodeInfoJvmMemory `json:"mem"`
	MemoryPools                           []string          `json:"memory_pools"`
	Pid                                   int               `json:"pid"`
	StartTimeInMillis                     int64             `json:"start_time_in_millis"`
	UsingBundledJdk                       bool              `json:"using_bundled_jdk"`
	UsingCompressedOrdinaryObjectPointers string            `json:"using_compressed_ordinary_object_pointers,omitempty"`
	Version                               VersionString     `json:"version"`
	VmName                                Name              `json:"vm_name"`
	VmVendor                              string            `json:"vm_vendor"`
	VmVersion                             VersionString     `json:"vm_version"`
}

// NodeJvmInfoBuilder holds NodeJvmInfo struct and provides a builder API.
type NodeJvmInfoBuilder struct {
	v *NodeJvmInfo
}

// NewNodeJvmInfo provides a builder for the NodeJvmInfo struct.
func NewNodeJvmInfo() *NodeJvmInfoBuilder {
	r := NodeJvmInfoBuilder{
		&NodeJvmInfo{},
	}

	return &r
}

// Build finalize the chain and returns the NodeJvmInfo struct
func (rb *NodeJvmInfoBuilder) Build() NodeJvmInfo {
	return *rb.v
}

// BundledJdk set the BundledJdk property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) BundledJdk(bundledjdk bool) *NodeJvmInfoBuilder {
	rb.v.BundledJdk = bundledjdk
	return rb
}

// GcCollectors set the GcCollectors property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) GcCollectors(gc_collectors ...string) *NodeJvmInfoBuilder {
	rb.v.GcCollectors = append(rb.v.GcCollectors, gc_collectors...)
	return rb
}

// InputArguments set the InputArguments property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) InputArguments(input_arguments ...string) *NodeJvmInfoBuilder {
	rb.v.InputArguments = append(rb.v.InputArguments, input_arguments...)
	return rb
}

// Mem set the Mem property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) Mem(mem NodeInfoJvmMemory) *NodeJvmInfoBuilder {
	rb.v.Mem = mem
	return rb
}

// MemoryPools set the MemoryPools property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) MemoryPools(memory_pools ...string) *NodeJvmInfoBuilder {
	rb.v.MemoryPools = append(rb.v.MemoryPools, memory_pools...)
	return rb
}

// Pid set the Pid property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) Pid(pid int) *NodeJvmInfoBuilder {
	rb.v.Pid = pid
	return rb
}

// StartTimeInMillis set the StartTimeInMillis property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) StartTimeInMillis(starttimeinmillis int64) *NodeJvmInfoBuilder {
	rb.v.StartTimeInMillis = starttimeinmillis
	return rb
}

// UsingBundledJdk set the UsingBundledJdk property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) UsingBundledJdk(usingbundledjdk bool) *NodeJvmInfoBuilder {
	rb.v.UsingBundledJdk = usingbundledjdk
	return rb
}

// UsingCompressedOrdinaryObjectPointers set the UsingCompressedOrdinaryObjectPointers property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) UsingCompressedOrdinaryObjectPointers(arg string) *NodeJvmInfoBuilder {
	rb.v.UsingCompressedOrdinaryObjectPointers = arg
	return rb
}

// Version set the Version property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) Version(version VersionString) *NodeJvmInfoBuilder {
	rb.v.Version = version
	return rb
}

// VmName set the VmName property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) VmName(vmname Name) *NodeJvmInfoBuilder {
	rb.v.VmName = vmname
	return rb
}

// VmVendor set the VmVendor property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) VmVendor(vmvendor string) *NodeJvmInfoBuilder {
	rb.v.VmVendor = vmvendor
	return rb
}

// VmVersion set the VmVersion property for NodeJvmInfoBuilder.
func (rb *NodeJvmInfoBuilder) VmVersion(vmversion VersionString) *NodeJvmInfoBuilder {
	rb.v.VmVersion = vmversion
	return rb
}
