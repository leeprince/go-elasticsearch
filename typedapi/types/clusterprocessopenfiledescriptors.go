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

// ClusterProcessOpenFileDescriptors type.
type ClusterProcessOpenFileDescriptors struct {
	Avg int64 `json:"avg"`
	Max int64 `json:"max"`
	Min int64 `json:"min"`
}

// ClusterProcessOpenFileDescriptorsBuilder holds ClusterProcessOpenFileDescriptors struct and provides a builder API.
type ClusterProcessOpenFileDescriptorsBuilder struct {
	v *ClusterProcessOpenFileDescriptors
}

// NewClusterProcessOpenFileDescriptors provides a builder for the ClusterProcessOpenFileDescriptors struct.
func NewClusterProcessOpenFileDescriptors() *ClusterProcessOpenFileDescriptorsBuilder {
	r := ClusterProcessOpenFileDescriptorsBuilder{
		&ClusterProcessOpenFileDescriptors{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterProcessOpenFileDescriptors struct
func (rb *ClusterProcessOpenFileDescriptorsBuilder) Build() ClusterProcessOpenFileDescriptors {
	return *rb.v
}

// Avg set the Avg property for ClusterProcessOpenFileDescriptorsBuilder.
func (rb *ClusterProcessOpenFileDescriptorsBuilder) Avg(avg int64) *ClusterProcessOpenFileDescriptorsBuilder {
	rb.v.Avg = avg
	return rb
}

// Max set the Max property for ClusterProcessOpenFileDescriptorsBuilder.
func (rb *ClusterProcessOpenFileDescriptorsBuilder) Max(max int64) *ClusterProcessOpenFileDescriptorsBuilder {
	rb.v.Max = max
	return rb
}

// Min set the Min property for ClusterProcessOpenFileDescriptorsBuilder.
func (rb *ClusterProcessOpenFileDescriptorsBuilder) Min(min int64) *ClusterProcessOpenFileDescriptorsBuilder {
	rb.v.Min = min
	return rb
}
