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

// ClusterJvm type.
type ClusterJvm struct {
	MaxUptimeInMillis int64               `json:"max_uptime_in_millis"`
	Mem               ClusterJvmMemory    `json:"mem"`
	Threads           int64               `json:"threads"`
	Versions          []ClusterJvmVersion `json:"versions"`
}

// ClusterJvmBuilder holds ClusterJvm struct and provides a builder API.
type ClusterJvmBuilder struct {
	v *ClusterJvm
}

// NewClusterJvm provides a builder for the ClusterJvm struct.
func NewClusterJvm() *ClusterJvmBuilder {
	r := ClusterJvmBuilder{
		&ClusterJvm{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterJvm struct
func (rb *ClusterJvmBuilder) Build() ClusterJvm {
	return *rb.v
}

// MaxUptimeInMillis set the MaxUptimeInMillis property for ClusterJvmBuilder.
func (rb *ClusterJvmBuilder) MaxUptimeInMillis(maxuptimeinmillis int64) *ClusterJvmBuilder {
	rb.v.MaxUptimeInMillis = maxuptimeinmillis
	return rb
}

// Mem set the Mem property for ClusterJvmBuilder.
func (rb *ClusterJvmBuilder) Mem(mem ClusterJvmMemory) *ClusterJvmBuilder {
	rb.v.Mem = mem
	return rb
}

// Threads set the Threads property for ClusterJvmBuilder.
func (rb *ClusterJvmBuilder) Threads(threads int64) *ClusterJvmBuilder {
	rb.v.Threads = threads
	return rb
}

// Versions set the Versions property for ClusterJvmBuilder.
func (rb *ClusterJvmBuilder) Versions(versions ...ClusterJvmVersion) *ClusterJvmBuilder {
	rb.v.Versions = append(rb.v.Versions, versions...)
	return rb
}
