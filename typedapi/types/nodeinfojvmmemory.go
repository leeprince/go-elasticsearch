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

// NodeInfoJvmMemory type.
type NodeInfoJvmMemory struct {
	DirectMax          *ByteSize `json:"direct_max,omitempty"`
	DirectMaxInBytes   int64     `json:"direct_max_in_bytes"`
	HeapInit           *ByteSize `json:"heap_init,omitempty"`
	HeapInitInBytes    int64     `json:"heap_init_in_bytes"`
	HeapMax            *ByteSize `json:"heap_max,omitempty"`
	HeapMaxInBytes     int64     `json:"heap_max_in_bytes"`
	NonHeapInit        *ByteSize `json:"non_heap_init,omitempty"`
	NonHeapInitInBytes int64     `json:"non_heap_init_in_bytes"`
	NonHeapMax         *ByteSize `json:"non_heap_max,omitempty"`
	NonHeapMaxInBytes  int64     `json:"non_heap_max_in_bytes"`
}

// NodeInfoJvmMemoryBuilder holds NodeInfoJvmMemory struct and provides a builder API.
type NodeInfoJvmMemoryBuilder struct {
	v *NodeInfoJvmMemory
}

// NewNodeInfoJvmMemory provides a builder for the NodeInfoJvmMemory struct.
func NewNodeInfoJvmMemory() *NodeInfoJvmMemoryBuilder {
	r := NodeInfoJvmMemoryBuilder{
		&NodeInfoJvmMemory{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoJvmMemory struct
func (rb *NodeInfoJvmMemoryBuilder) Build() NodeInfoJvmMemory {
	return *rb.v
}

// DirectMax set the DirectMax property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) DirectMax(directmax ByteSize) *NodeInfoJvmMemoryBuilder {
	rb.v.DirectMax = &directmax
	return rb
}

// DirectMaxInBytes set the DirectMaxInBytes property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) DirectMaxInBytes(directmaxinbytes int64) *NodeInfoJvmMemoryBuilder {
	rb.v.DirectMaxInBytes = directmaxinbytes
	return rb
}

// HeapInit set the HeapInit property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) HeapInit(heapinit ByteSize) *NodeInfoJvmMemoryBuilder {
	rb.v.HeapInit = &heapinit
	return rb
}

// HeapInitInBytes set the HeapInitInBytes property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) HeapInitInBytes(heapinitinbytes int64) *NodeInfoJvmMemoryBuilder {
	rb.v.HeapInitInBytes = heapinitinbytes
	return rb
}

// HeapMax set the HeapMax property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) HeapMax(heapmax ByteSize) *NodeInfoJvmMemoryBuilder {
	rb.v.HeapMax = &heapmax
	return rb
}

// HeapMaxInBytes set the HeapMaxInBytes property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) HeapMaxInBytes(heapmaxinbytes int64) *NodeInfoJvmMemoryBuilder {
	rb.v.HeapMaxInBytes = heapmaxinbytes
	return rb
}

// NonHeapInit set the NonHeapInit property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) NonHeapInit(nonheapinit ByteSize) *NodeInfoJvmMemoryBuilder {
	rb.v.NonHeapInit = &nonheapinit
	return rb
}

// NonHeapInitInBytes set the NonHeapInitInBytes property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) NonHeapInitInBytes(nonheapinitinbytes int64) *NodeInfoJvmMemoryBuilder {
	rb.v.NonHeapInitInBytes = nonheapinitinbytes
	return rb
}

// NonHeapMax set the NonHeapMax property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) NonHeapMax(nonheapmax ByteSize) *NodeInfoJvmMemoryBuilder {
	rb.v.NonHeapMax = &nonheapmax
	return rb
}

// NonHeapMaxInBytes set the NonHeapMaxInBytes property for NodeInfoJvmMemoryBuilder.
func (rb *NodeInfoJvmMemoryBuilder) NonHeapMaxInBytes(nonheapmaxinbytes int64) *NodeInfoJvmMemoryBuilder {
	rb.v.NonHeapMaxInBytes = nonheapmaxinbytes
	return rb
}
