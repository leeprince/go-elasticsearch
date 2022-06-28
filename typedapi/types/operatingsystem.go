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

// OperatingSystem type.
type OperatingSystem struct {
	Cgroup    *Cgroup              `json:"cgroup,omitempty"`
	Cpu       *Cpu                 `json:"cpu,omitempty"`
	Mem       *ExtendedMemoryStats `json:"mem,omitempty"`
	Swap      *MemoryStats         `json:"swap,omitempty"`
	Timestamp *int64               `json:"timestamp,omitempty"`
}

// OperatingSystemBuilder holds OperatingSystem struct and provides a builder API.
type OperatingSystemBuilder struct {
	v *OperatingSystem
}

// NewOperatingSystem provides a builder for the OperatingSystem struct.
func NewOperatingSystem() *OperatingSystemBuilder {
	r := OperatingSystemBuilder{
		&OperatingSystem{},
	}

	return &r
}

// Build finalize the chain and returns the OperatingSystem struct
func (rb *OperatingSystemBuilder) Build() OperatingSystem {
	return *rb.v
}

// Cgroup set the Cgroup property for OperatingSystemBuilder.
func (rb *OperatingSystemBuilder) Cgroup(cgroup Cgroup) *OperatingSystemBuilder {
	rb.v.Cgroup = &cgroup
	return rb
}

// Cpu set the Cpu property for OperatingSystemBuilder.
func (rb *OperatingSystemBuilder) Cpu(cpu Cpu) *OperatingSystemBuilder {
	rb.v.Cpu = &cpu
	return rb
}

// Mem set the Mem property for OperatingSystemBuilder.
func (rb *OperatingSystemBuilder) Mem(mem ExtendedMemoryStats) *OperatingSystemBuilder {
	rb.v.Mem = &mem
	return rb
}

// Swap set the Swap property for OperatingSystemBuilder.
func (rb *OperatingSystemBuilder) Swap(swap MemoryStats) *OperatingSystemBuilder {
	rb.v.Swap = &swap
	return rb
}

// Timestamp set the Timestamp property for OperatingSystemBuilder.
func (rb *OperatingSystemBuilder) Timestamp(timestamp int64) *OperatingSystemBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}
