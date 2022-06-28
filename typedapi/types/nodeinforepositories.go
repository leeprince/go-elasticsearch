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

// NodeInfoRepositories type.
type NodeInfoRepositories struct {
	Url NodeInfoRepositoriesUrl `json:"url"`
}

// NodeInfoRepositoriesBuilder holds NodeInfoRepositories struct and provides a builder API.
type NodeInfoRepositoriesBuilder struct {
	v *NodeInfoRepositories
}

// NewNodeInfoRepositories provides a builder for the NodeInfoRepositories struct.
func NewNodeInfoRepositories() *NodeInfoRepositoriesBuilder {
	r := NodeInfoRepositoriesBuilder{
		&NodeInfoRepositories{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoRepositories struct
func (rb *NodeInfoRepositoriesBuilder) Build() NodeInfoRepositories {
	return *rb.v
}

// Url set the Url property for NodeInfoRepositoriesBuilder.
func (rb *NodeInfoRepositoriesBuilder) Url(url NodeInfoRepositoriesUrl) *NodeInfoRepositoriesBuilder {
	rb.v.Url = url
	return rb
}
