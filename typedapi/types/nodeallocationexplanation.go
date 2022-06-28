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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/decision"
)

// NodeAllocationExplanation type.
type NodeAllocationExplanation struct {
	Deciders         []AllocationDecision `json:"deciders"`
	NodeAttributes   map[string]string    `json:"node_attributes"`
	NodeDecision     decision.Decision    `json:"node_decision"`
	NodeId           Id                   `json:"node_id"`
	NodeName         Name                 `json:"node_name"`
	Store            *AllocationStore     `json:"store,omitempty"`
	TransportAddress TransportAddress     `json:"transport_address"`
	WeightRanking    int                  `json:"weight_ranking"`
}

// NodeAllocationExplanationBuilder holds NodeAllocationExplanation struct and provides a builder API.
type NodeAllocationExplanationBuilder struct {
	v *NodeAllocationExplanation
}

// NewNodeAllocationExplanation provides a builder for the NodeAllocationExplanation struct.
func NewNodeAllocationExplanation() *NodeAllocationExplanationBuilder {
	r := NodeAllocationExplanationBuilder{
		&NodeAllocationExplanation{
			NodeAttributes: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeAllocationExplanation struct
func (rb *NodeAllocationExplanationBuilder) Build() NodeAllocationExplanation {
	return *rb.v
}

// Deciders set the Deciders property for NodeAllocationExplanationBuilder.
func (rb *NodeAllocationExplanationBuilder) Deciders(deciders ...AllocationDecision) *NodeAllocationExplanationBuilder {
	rb.v.Deciders = append(rb.v.Deciders, deciders...)
	return rb
}

// NodeAttributes set the NodeAttributes property for NodeAllocationExplanationBuilder.
func (rb *NodeAllocationExplanationBuilder) NodeAttributes(value map[string]string) *NodeAllocationExplanationBuilder {
	rb.v.NodeAttributes = value
	return rb
}

// NodeDecision set the NodeDecision property for NodeAllocationExplanationBuilder.
func (rb *NodeAllocationExplanationBuilder) NodeDecision(nodedecision decision.Decision) *NodeAllocationExplanationBuilder {
	rb.v.NodeDecision = nodedecision
	return rb
}

// NodeId set the NodeId property for NodeAllocationExplanationBuilder.
func (rb *NodeAllocationExplanationBuilder) NodeId(nodeid Id) *NodeAllocationExplanationBuilder {
	rb.v.NodeId = nodeid
	return rb
}

// NodeName set the NodeName property for NodeAllocationExplanationBuilder.
func (rb *NodeAllocationExplanationBuilder) NodeName(nodename Name) *NodeAllocationExplanationBuilder {
	rb.v.NodeName = nodename
	return rb
}

// Store set the Store property for NodeAllocationExplanationBuilder.
func (rb *NodeAllocationExplanationBuilder) Store(store AllocationStore) *NodeAllocationExplanationBuilder {
	rb.v.Store = &store
	return rb
}

// TransportAddress set the TransportAddress property for NodeAllocationExplanationBuilder.
func (rb *NodeAllocationExplanationBuilder) TransportAddress(transportaddress TransportAddress) *NodeAllocationExplanationBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}

// WeightRanking set the WeightRanking property for NodeAllocationExplanationBuilder.
func (rb *NodeAllocationExplanationBuilder) WeightRanking(weightranking int) *NodeAllocationExplanationBuilder {
	rb.v.WeightRanking = weightranking
	return rb
}
