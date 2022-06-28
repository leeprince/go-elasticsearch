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

// NodeThreadPoolInfo type.
type NodeThreadPoolInfo struct {
	Core      *int    `json:"core,omitempty"`
	KeepAlive *string `json:"keep_alive,omitempty"`
	Max       *int    `json:"max,omitempty"`
	QueueSize int     `json:"queue_size"`
	Size      *int    `json:"size,omitempty"`
	Type      string  `json:"type"`
}

// NodeThreadPoolInfoBuilder holds NodeThreadPoolInfo struct and provides a builder API.
type NodeThreadPoolInfoBuilder struct {
	v *NodeThreadPoolInfo
}

// NewNodeThreadPoolInfo provides a builder for the NodeThreadPoolInfo struct.
func NewNodeThreadPoolInfo() *NodeThreadPoolInfoBuilder {
	r := NodeThreadPoolInfoBuilder{
		&NodeThreadPoolInfo{},
	}

	return &r
}

// Build finalize the chain and returns the NodeThreadPoolInfo struct
func (rb *NodeThreadPoolInfoBuilder) Build() NodeThreadPoolInfo {
	return *rb.v
}

// Core set the Core property for NodeThreadPoolInfoBuilder.
func (rb *NodeThreadPoolInfoBuilder) Core(core int) *NodeThreadPoolInfoBuilder {
	rb.v.Core = &core
	return rb
}

// KeepAlive set the KeepAlive property for NodeThreadPoolInfoBuilder.
func (rb *NodeThreadPoolInfoBuilder) KeepAlive(keepalive string) *NodeThreadPoolInfoBuilder {
	rb.v.KeepAlive = &keepalive
	return rb
}

// Max set the Max property for NodeThreadPoolInfoBuilder.
func (rb *NodeThreadPoolInfoBuilder) Max(max int) *NodeThreadPoolInfoBuilder {
	rb.v.Max = &max
	return rb
}

// QueueSize set the QueueSize property for NodeThreadPoolInfoBuilder.
func (rb *NodeThreadPoolInfoBuilder) QueueSize(queuesize int) *NodeThreadPoolInfoBuilder {
	rb.v.QueueSize = queuesize
	return rb
}

// Size set the Size property for NodeThreadPoolInfoBuilder.
func (rb *NodeThreadPoolInfoBuilder) Size(size int) *NodeThreadPoolInfoBuilder {
	rb.v.Size = &size
	return rb
}

// Type set the Type property for NodeThreadPoolInfoBuilder.
func (rb *NodeThreadPoolInfoBuilder) Type_(type_ string) *NodeThreadPoolInfoBuilder {
	rb.v.Type = type_
	return rb
}
