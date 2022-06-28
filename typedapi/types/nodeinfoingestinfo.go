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

// NodeInfoIngestInfo type.
type NodeInfoIngestInfo struct {
	Downloader NodeInfoIngestDownloader `json:"downloader"`
}

// NodeInfoIngestInfoBuilder holds NodeInfoIngestInfo struct and provides a builder API.
type NodeInfoIngestInfoBuilder struct {
	v *NodeInfoIngestInfo
}

// NewNodeInfoIngestInfo provides a builder for the NodeInfoIngestInfo struct.
func NewNodeInfoIngestInfo() *NodeInfoIngestInfoBuilder {
	r := NodeInfoIngestInfoBuilder{
		&NodeInfoIngestInfo{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoIngestInfo struct
func (rb *NodeInfoIngestInfoBuilder) Build() NodeInfoIngestInfo {
	return *rb.v
}

// Downloader set the Downloader property for NodeInfoIngestInfoBuilder.
func (rb *NodeInfoIngestInfoBuilder) Downloader(downloader NodeInfoIngestDownloader) *NodeInfoIngestInfoBuilder {
	rb.v.Downloader = downloader
	return rb
}
