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

// TrainedModelTreeNode type.
type TrainedModelTreeNode struct {
	DecisionType *string  `json:"decision_type,omitempty"`
	DefaultLeft  *bool    `json:"default_left,omitempty"`
	LeafValue    *float64 `json:"leaf_value,omitempty"`
	LeftChild    *int     `json:"left_child,omitempty"`
	NodeIndex    int      `json:"node_index"`
	RightChild   *int     `json:"right_child,omitempty"`
	SplitFeature *int     `json:"split_feature,omitempty"`
	SplitGain    *int     `json:"split_gain,omitempty"`
	Threshold    *float64 `json:"threshold,omitempty"`
}

// TrainedModelTreeNodeBuilder holds TrainedModelTreeNode struct and provides a builder API.
type TrainedModelTreeNodeBuilder struct {
	v *TrainedModelTreeNode
}

// NewTrainedModelTreeNode provides a builder for the TrainedModelTreeNode struct.
func NewTrainedModelTreeNode() *TrainedModelTreeNodeBuilder {
	r := TrainedModelTreeNodeBuilder{
		&TrainedModelTreeNode{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelTreeNode struct
func (rb *TrainedModelTreeNodeBuilder) Build() TrainedModelTreeNode {
	return *rb.v
}

// DecisionType set the DecisionType property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) DecisionType(decisiontype string) *TrainedModelTreeNodeBuilder {
	rb.v.DecisionType = &decisiontype
	return rb
}

// DefaultLeft set the DefaultLeft property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) DefaultLeft(defaultleft bool) *TrainedModelTreeNodeBuilder {
	rb.v.DefaultLeft = &defaultleft
	return rb
}

// LeafValue set the LeafValue property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) LeafValue(leafvalue float64) *TrainedModelTreeNodeBuilder {
	rb.v.LeafValue = &leafvalue
	return rb
}

// LeftChild set the LeftChild property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) LeftChild(leftchild int) *TrainedModelTreeNodeBuilder {
	rb.v.LeftChild = &leftchild
	return rb
}

// NodeIndex set the NodeIndex property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) NodeIndex(nodeindex int) *TrainedModelTreeNodeBuilder {
	rb.v.NodeIndex = nodeindex
	return rb
}

// RightChild set the RightChild property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) RightChild(rightchild int) *TrainedModelTreeNodeBuilder {
	rb.v.RightChild = &rightchild
	return rb
}

// SplitFeature set the SplitFeature property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) SplitFeature(splitfeature int) *TrainedModelTreeNodeBuilder {
	rb.v.SplitFeature = &splitfeature
	return rb
}

// SplitGain set the SplitGain property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) SplitGain(splitgain int) *TrainedModelTreeNodeBuilder {
	rb.v.SplitGain = &splitgain
	return rb
}

// Threshold set the Threshold property for TrainedModelTreeNodeBuilder.
func (rb *TrainedModelTreeNodeBuilder) Threshold(threshold float64) *TrainedModelTreeNodeBuilder {
	rb.v.Threshold = &threshold
	return rb
}
