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

// NodeInfoPath type.
type NodeInfoPath struct {
	Data []string `json:"data,omitempty"`
	Home string   `json:"home"`
	Logs string   `json:"logs"`
	Repo []string `json:"repo"`
}

// NodeInfoPathBuilder holds NodeInfoPath struct and provides a builder API.
type NodeInfoPathBuilder struct {
	v *NodeInfoPath
}

// NewNodeInfoPath provides a builder for the NodeInfoPath struct.
func NewNodeInfoPath() *NodeInfoPathBuilder {
	r := NodeInfoPathBuilder{
		&NodeInfoPath{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoPath struct
func (rb *NodeInfoPathBuilder) Build() NodeInfoPath {
	return *rb.v
}

// Data set the Data property for NodeInfoPathBuilder.
func (rb *NodeInfoPathBuilder) Data(data ...string) *NodeInfoPathBuilder {
	rb.v.Data = append(rb.v.Data, data...)
	return rb
}

// Home set the Home property for NodeInfoPathBuilder.
func (rb *NodeInfoPathBuilder) Home(home string) *NodeInfoPathBuilder {
	rb.v.Home = home
	return rb
}

// Logs set the Logs property for NodeInfoPathBuilder.
func (rb *NodeInfoPathBuilder) Logs(logs string) *NodeInfoPathBuilder {
	rb.v.Logs = logs
	return rb
}

// Repo set the Repo property for NodeInfoPathBuilder.
func (rb *NodeInfoPathBuilder) Repo(repo ...string) *NodeInfoPathBuilder {
	rb.v.Repo = append(rb.v.Repo, repo...)
	return rb
}
