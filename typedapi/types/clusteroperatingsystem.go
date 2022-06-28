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

// ClusterOperatingSystem type.
type ClusterOperatingSystem struct {
	AllocatedProcessors int                                  `json:"allocated_processors"`
	Architectures       []ClusterOperatingSystemArchitecture `json:"architectures,omitempty"`
	AvailableProcessors int                                  `json:"available_processors"`
	Mem                 OperatingSystemMemoryInfo            `json:"mem"`
	Names               []ClusterOperatingSystemName         `json:"names"`
	PrettyNames         []ClusterOperatingSystemPrettyName   `json:"pretty_names"`
}

// ClusterOperatingSystemBuilder holds ClusterOperatingSystem struct and provides a builder API.
type ClusterOperatingSystemBuilder struct {
	v *ClusterOperatingSystem
}

// NewClusterOperatingSystem provides a builder for the ClusterOperatingSystem struct.
func NewClusterOperatingSystem() *ClusterOperatingSystemBuilder {
	r := ClusterOperatingSystemBuilder{
		&ClusterOperatingSystem{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterOperatingSystem struct
func (rb *ClusterOperatingSystemBuilder) Build() ClusterOperatingSystem {
	return *rb.v
}

// AllocatedProcessors set the AllocatedProcessors property for ClusterOperatingSystemBuilder.
func (rb *ClusterOperatingSystemBuilder) AllocatedProcessors(allocatedprocessors int) *ClusterOperatingSystemBuilder {
	rb.v.AllocatedProcessors = allocatedprocessors
	return rb
}

// Architectures set the Architectures property for ClusterOperatingSystemBuilder.
func (rb *ClusterOperatingSystemBuilder) Architectures(architectures ...ClusterOperatingSystemArchitecture) *ClusterOperatingSystemBuilder {
	rb.v.Architectures = append(rb.v.Architectures, architectures...)
	return rb
}

// AvailableProcessors set the AvailableProcessors property for ClusterOperatingSystemBuilder.
func (rb *ClusterOperatingSystemBuilder) AvailableProcessors(availableprocessors int) *ClusterOperatingSystemBuilder {
	rb.v.AvailableProcessors = availableprocessors
	return rb
}

// Mem set the Mem property for ClusterOperatingSystemBuilder.
func (rb *ClusterOperatingSystemBuilder) Mem(mem OperatingSystemMemoryInfo) *ClusterOperatingSystemBuilder {
	rb.v.Mem = mem
	return rb
}

// Names set the Names property for ClusterOperatingSystemBuilder.
func (rb *ClusterOperatingSystemBuilder) Names(names ...ClusterOperatingSystemName) *ClusterOperatingSystemBuilder {
	rb.v.Names = append(rb.v.Names, names...)
	return rb
}

// PrettyNames set the PrettyNames property for ClusterOperatingSystemBuilder.
func (rb *ClusterOperatingSystemBuilder) PrettyNames(pretty_names ...ClusterOperatingSystemPrettyName) *ClusterOperatingSystemBuilder {
	rb.v.PrettyNames = append(rb.v.PrettyNames, pretty_names...)
	return rb
}
