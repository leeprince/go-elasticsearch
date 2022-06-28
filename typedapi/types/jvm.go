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

// Jvm type.
type Jvm struct {
	BufferPools    map[string]NodeBufferPool `json:"buffer_pools,omitempty"`
	Classes        *JvmClasses               `json:"classes,omitempty"`
	Gc             *GarbageCollector         `json:"gc,omitempty"`
	Mem            *JvmMemoryStats           `json:"mem,omitempty"`
	Threads        *JvmThreads               `json:"threads,omitempty"`
	Timestamp      *int64                    `json:"timestamp,omitempty"`
	Uptime         *string                   `json:"uptime,omitempty"`
	UptimeInMillis *int64                    `json:"uptime_in_millis,omitempty"`
}

// JvmBuilder holds Jvm struct and provides a builder API.
type JvmBuilder struct {
	v *Jvm
}

// NewJvm provides a builder for the Jvm struct.
func NewJvm() *JvmBuilder {
	r := JvmBuilder{
		&Jvm{
			BufferPools: make(map[string]NodeBufferPool, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Jvm struct
func (rb *JvmBuilder) Build() Jvm {
	return *rb.v
}

// BufferPools set the BufferPools property for JvmBuilder.
func (rb *JvmBuilder) BufferPools(value map[string]NodeBufferPool) *JvmBuilder {
	rb.v.BufferPools = value
	return rb
}

// Classes set the Classes property for JvmBuilder.
func (rb *JvmBuilder) Classes(classes JvmClasses) *JvmBuilder {
	rb.v.Classes = &classes
	return rb
}

// Gc set the Gc property for JvmBuilder.
func (rb *JvmBuilder) Gc(gc GarbageCollector) *JvmBuilder {
	rb.v.Gc = &gc
	return rb
}

// Mem set the Mem property for JvmBuilder.
func (rb *JvmBuilder) Mem(mem JvmMemoryStats) *JvmBuilder {
	rb.v.Mem = &mem
	return rb
}

// Threads set the Threads property for JvmBuilder.
func (rb *JvmBuilder) Threads(threads JvmThreads) *JvmBuilder {
	rb.v.Threads = &threads
	return rb
}

// Timestamp set the Timestamp property for JvmBuilder.
func (rb *JvmBuilder) Timestamp(timestamp int64) *JvmBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

// Uptime set the Uptime property for JvmBuilder.
func (rb *JvmBuilder) Uptime(uptime string) *JvmBuilder {
	rb.v.Uptime = &uptime
	return rb
}

// UptimeInMillis set the UptimeInMillis property for JvmBuilder.
func (rb *JvmBuilder) UptimeInMillis(uptimeinmillis int64) *JvmBuilder {
	rb.v.UptimeInMillis = &uptimeinmillis
	return rb
}
