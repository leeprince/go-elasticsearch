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

// PressureMemory type.
type PressureMemory struct {
	AllInBytes                            *int64 `json:"all_in_bytes,omitempty"`
	CombinedCoordinatingAndPrimaryInBytes *int64 `json:"combined_coordinating_and_primary_in_bytes,omitempty"`
	CoordinatingInBytes                   *int64 `json:"coordinating_in_bytes,omitempty"`
	CoordinatingRejections                *int64 `json:"coordinating_rejections,omitempty"`
	PrimaryInBytes                        *int64 `json:"primary_in_bytes,omitempty"`
	PrimaryRejections                     *int64 `json:"primary_rejections,omitempty"`
	ReplicaInBytes                        *int64 `json:"replica_in_bytes,omitempty"`
	ReplicaRejections                     *int64 `json:"replica_rejections,omitempty"`
}

// PressureMemoryBuilder holds PressureMemory struct and provides a builder API.
type PressureMemoryBuilder struct {
	v *PressureMemory
}

// NewPressureMemory provides a builder for the PressureMemory struct.
func NewPressureMemory() *PressureMemoryBuilder {
	r := PressureMemoryBuilder{
		&PressureMemory{},
	}

	return &r
}

// Build finalize the chain and returns the PressureMemory struct
func (rb *PressureMemoryBuilder) Build() PressureMemory {
	return *rb.v
}

// AllInBytes set the AllInBytes property for PressureMemoryBuilder.
func (rb *PressureMemoryBuilder) AllInBytes(allinbytes int64) *PressureMemoryBuilder {
	rb.v.AllInBytes = &allinbytes
	return rb
}

// CombinedCoordinatingAndPrimaryInBytes set the CombinedCoordinatingAndPrimaryInBytes property for PressureMemoryBuilder.
func (rb *PressureMemoryBuilder) CombinedCoordinatingAndPrimaryInBytes(combinedcoordinatingandprimaryinbytes int64) *PressureMemoryBuilder {
	rb.v.CombinedCoordinatingAndPrimaryInBytes = &combinedcoordinatingandprimaryinbytes
	return rb
}

// CoordinatingInBytes set the CoordinatingInBytes property for PressureMemoryBuilder.
func (rb *PressureMemoryBuilder) CoordinatingInBytes(coordinatinginbytes int64) *PressureMemoryBuilder {
	rb.v.CoordinatingInBytes = &coordinatinginbytes
	return rb
}

// CoordinatingRejections set the CoordinatingRejections property for PressureMemoryBuilder.
func (rb *PressureMemoryBuilder) CoordinatingRejections(coordinatingrejections int64) *PressureMemoryBuilder {
	rb.v.CoordinatingRejections = &coordinatingrejections
	return rb
}

// PrimaryInBytes set the PrimaryInBytes property for PressureMemoryBuilder.
func (rb *PressureMemoryBuilder) PrimaryInBytes(primaryinbytes int64) *PressureMemoryBuilder {
	rb.v.PrimaryInBytes = &primaryinbytes
	return rb
}

// PrimaryRejections set the PrimaryRejections property for PressureMemoryBuilder.
func (rb *PressureMemoryBuilder) PrimaryRejections(primaryrejections int64) *PressureMemoryBuilder {
	rb.v.PrimaryRejections = &primaryrejections
	return rb
}

// ReplicaInBytes set the ReplicaInBytes property for PressureMemoryBuilder.
func (rb *PressureMemoryBuilder) ReplicaInBytes(replicainbytes int64) *PressureMemoryBuilder {
	rb.v.ReplicaInBytes = &replicainbytes
	return rb
}

// ReplicaRejections set the ReplicaRejections property for PressureMemoryBuilder.
func (rb *PressureMemoryBuilder) ReplicaRejections(replicarejections int64) *PressureMemoryBuilder {
	rb.v.ReplicaRejections = &replicarejections
	return rb
}
