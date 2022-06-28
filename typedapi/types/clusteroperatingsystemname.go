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

// ClusterOperatingSystemName type.
type ClusterOperatingSystemName struct {
	Count int  `json:"count"`
	Name  Name `json:"name"`
}

// ClusterOperatingSystemNameBuilder holds ClusterOperatingSystemName struct and provides a builder API.
type ClusterOperatingSystemNameBuilder struct {
	v *ClusterOperatingSystemName
}

// NewClusterOperatingSystemName provides a builder for the ClusterOperatingSystemName struct.
func NewClusterOperatingSystemName() *ClusterOperatingSystemNameBuilder {
	r := ClusterOperatingSystemNameBuilder{
		&ClusterOperatingSystemName{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterOperatingSystemName struct
func (rb *ClusterOperatingSystemNameBuilder) Build() ClusterOperatingSystemName {
	return *rb.v
}

// Count set the Count property for ClusterOperatingSystemNameBuilder.
func (rb *ClusterOperatingSystemNameBuilder) Count(count int) *ClusterOperatingSystemNameBuilder {
	rb.v.Count = count
	return rb
}

// Name set the Name property for ClusterOperatingSystemNameBuilder.
func (rb *ClusterOperatingSystemNameBuilder) Name(name Name) *ClusterOperatingSystemNameBuilder {
	rb.v.Name = name
	return rb
}
