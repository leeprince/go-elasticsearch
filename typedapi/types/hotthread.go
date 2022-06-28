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

// HotThread type.
type HotThread struct {
	Hosts    []Host   `json:"hosts"`
	NodeId   Id       `json:"node_id"`
	NodeName Name     `json:"node_name"`
	Threads  []string `json:"threads"`
}

// HotThreadBuilder holds HotThread struct and provides a builder API.
type HotThreadBuilder struct {
	v *HotThread
}

// NewHotThread provides a builder for the HotThread struct.
func NewHotThread() *HotThreadBuilder {
	r := HotThreadBuilder{
		&HotThread{},
	}

	return &r
}

// Build finalize the chain and returns the HotThread struct
func (rb *HotThreadBuilder) Build() HotThread {
	return *rb.v
}

// Hosts set the Hosts property for HotThreadBuilder.
func (rb *HotThreadBuilder) Hosts(hosts ...Host) *HotThreadBuilder {
	rb.v.Hosts = append(rb.v.Hosts, hosts...)
	return rb
}

// NodeId set the NodeId property for HotThreadBuilder.
func (rb *HotThreadBuilder) NodeId(nodeid Id) *HotThreadBuilder {
	rb.v.NodeId = nodeid
	return rb
}

// NodeName set the NodeName property for HotThreadBuilder.
func (rb *HotThreadBuilder) NodeName(nodename Name) *HotThreadBuilder {
	rb.v.NodeName = nodename
	return rb
}

// Threads set the Threads property for HotThreadBuilder.
func (rb *HotThreadBuilder) Threads(threads ...string) *HotThreadBuilder {
	rb.v.Threads = append(rb.v.Threads, threads...)
	return rb
}
