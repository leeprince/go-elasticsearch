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

// NodeInfoSettings type.
type NodeInfoSettings struct {
	Action       *NodeInfoAction           `json:"action,omitempty"`
	Bootstrap    *NodeInfoBootstrap        `json:"bootstrap,omitempty"`
	Client       NodeInfoClient            `json:"client"`
	Cluster      NodeInfoSettingsCluster   `json:"cluster"`
	Discovery    *NodeInfoDiscover         `json:"discovery,omitempty"`
	Http         NodeInfoSettingsHttp      `json:"http"`
	Ingest       *NodeInfoSettingsIngest   `json:"ingest,omitempty"`
	Network      *NodeInfoSettingsNetwork  `json:"network,omitempty"`
	Node         NodeInfoSettingsNode      `json:"node"`
	Path         NodeInfoPath              `json:"path"`
	Repositories *NodeInfoRepositories     `json:"repositories,omitempty"`
	Script       *NodeInfoScript           `json:"script,omitempty"`
	Search       *NodeInfoSearch           `json:"search,omitempty"`
	Transport    NodeInfoSettingsTransport `json:"transport"`
	Xpack        *NodeInfoXpack            `json:"xpack,omitempty"`
}

// NodeInfoSettingsBuilder holds NodeInfoSettings struct and provides a builder API.
type NodeInfoSettingsBuilder struct {
	v *NodeInfoSettings
}

// NewNodeInfoSettings provides a builder for the NodeInfoSettings struct.
func NewNodeInfoSettings() *NodeInfoSettingsBuilder {
	r := NodeInfoSettingsBuilder{
		&NodeInfoSettings{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettings struct
func (rb *NodeInfoSettingsBuilder) Build() NodeInfoSettings {
	return *rb.v
}

// Action set the Action property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Action(action NodeInfoAction) *NodeInfoSettingsBuilder {
	rb.v.Action = &action
	return rb
}

// Bootstrap set the Bootstrap property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Bootstrap(bootstrap NodeInfoBootstrap) *NodeInfoSettingsBuilder {
	rb.v.Bootstrap = &bootstrap
	return rb
}

// Client set the Client property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Client(client NodeInfoClient) *NodeInfoSettingsBuilder {
	rb.v.Client = client
	return rb
}

// Cluster set the Cluster property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Cluster(cluster NodeInfoSettingsCluster) *NodeInfoSettingsBuilder {
	rb.v.Cluster = cluster
	return rb
}

// Discovery set the Discovery property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Discovery(discovery NodeInfoDiscover) *NodeInfoSettingsBuilder {
	rb.v.Discovery = &discovery
	return rb
}

// Http set the Http property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Http(http NodeInfoSettingsHttp) *NodeInfoSettingsBuilder {
	rb.v.Http = http
	return rb
}

// Ingest set the Ingest property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Ingest(ingest NodeInfoSettingsIngest) *NodeInfoSettingsBuilder {
	rb.v.Ingest = &ingest
	return rb
}

// Network set the Network property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Network(network NodeInfoSettingsNetwork) *NodeInfoSettingsBuilder {
	rb.v.Network = &network
	return rb
}

// Node set the Node property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Node(node NodeInfoSettingsNode) *NodeInfoSettingsBuilder {
	rb.v.Node = node
	return rb
}

// Path set the Path property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Path(path NodeInfoPath) *NodeInfoSettingsBuilder {
	rb.v.Path = path
	return rb
}

// Repositories set the Repositories property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Repositories(repositories NodeInfoRepositories) *NodeInfoSettingsBuilder {
	rb.v.Repositories = &repositories
	return rb
}

// Script set the Script property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Script(script NodeInfoScript) *NodeInfoSettingsBuilder {
	rb.v.Script = &script
	return rb
}

// Search set the Search property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Search(search NodeInfoSearch) *NodeInfoSettingsBuilder {
	rb.v.Search = &search
	return rb
}

// Transport set the Transport property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Transport(transport NodeInfoSettingsTransport) *NodeInfoSettingsBuilder {
	rb.v.Transport = transport
	return rb
}

// Xpack set the Xpack property for NodeInfoSettingsBuilder.
func (rb *NodeInfoSettingsBuilder) Xpack(xpack NodeInfoXpack) *NodeInfoSettingsBuilder {
	rb.v.Xpack = &xpack
	return rb
}
