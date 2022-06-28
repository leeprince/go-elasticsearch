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

// NodesCredentials type.
type NodesCredentials struct {
	// FileTokens File-backed tokens collected from all nodes
	FileTokens map[string]NodesCredentialsFileToken `json:"file_tokens"`
	// NodeStats General status showing how nodes respond to the above collection request
	NodeStats NodeStatistics `json:"_nodes"`
}

// NodesCredentialsBuilder holds NodesCredentials struct and provides a builder API.
type NodesCredentialsBuilder struct {
	v *NodesCredentials
}

// NewNodesCredentials provides a builder for the NodesCredentials struct.
func NewNodesCredentials() *NodesCredentialsBuilder {
	r := NodesCredentialsBuilder{
		&NodesCredentials{
			FileTokens: make(map[string]NodesCredentialsFileToken, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodesCredentials struct
func (rb *NodesCredentialsBuilder) Build() NodesCredentials {
	return *rb.v
}

// FileTokens File-backed tokens collected from all nodes
func (rb *NodesCredentialsBuilder) FileTokens(value map[string]NodesCredentialsFileToken) *NodesCredentialsBuilder {
	rb.v.FileTokens = value
	return rb
}

// NodeStats General status showing how nodes respond to the above collection request
func (rb *NodesCredentialsBuilder) NodeStats(nodestats NodeStatistics) *NodesCredentialsBuilder {
	rb.v.NodeStats = nodestats
	return rb
}
