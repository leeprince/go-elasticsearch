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

// CurrentNode type.
type CurrentNode struct {
	Attributes       map[string]string `json:"attributes"`
	Id               Id                `json:"id"`
	Name             Name              `json:"name"`
	TransportAddress TransportAddress  `json:"transport_address"`
	WeightRanking    int               `json:"weight_ranking"`
}

// CurrentNodeBuilder holds CurrentNode struct and provides a builder API.
type CurrentNodeBuilder struct {
	v *CurrentNode
}

// NewCurrentNode provides a builder for the CurrentNode struct.
func NewCurrentNode() *CurrentNodeBuilder {
	r := CurrentNodeBuilder{
		&CurrentNode{
			Attributes: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CurrentNode struct
func (rb *CurrentNodeBuilder) Build() CurrentNode {
	return *rb.v
}

// Attributes set the Attributes property for CurrentNodeBuilder.
func (rb *CurrentNodeBuilder) Attributes(value map[string]string) *CurrentNodeBuilder {
	rb.v.Attributes = value
	return rb
}

// Id set the Id property for CurrentNodeBuilder.
func (rb *CurrentNodeBuilder) Id(id Id) *CurrentNodeBuilder {
	rb.v.Id = id
	return rb
}

// Name set the Name property for CurrentNodeBuilder.
func (rb *CurrentNodeBuilder) Name(name Name) *CurrentNodeBuilder {
	rb.v.Name = name
	return rb
}

// TransportAddress set the TransportAddress property for CurrentNodeBuilder.
func (rb *CurrentNodeBuilder) TransportAddress(transportaddress TransportAddress) *CurrentNodeBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}

// WeightRanking set the WeightRanking property for CurrentNodeBuilder.
func (rb *CurrentNodeBuilder) WeightRanking(weightranking int) *CurrentNodeBuilder {
	rb.v.WeightRanking = weightranking
	return rb
}
