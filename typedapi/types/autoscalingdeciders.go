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

// AutoscalingDeciders type.
type AutoscalingDeciders struct {
	CurrentCapacity  AutoscalingCapacity           `json:"current_capacity"`
	CurrentNodes     []AutoscalingNode             `json:"current_nodes"`
	Deciders         map[string]AutoscalingDecider `json:"deciders"`
	RequiredCapacity AutoscalingCapacity           `json:"required_capacity"`
}

// AutoscalingDecidersBuilder holds AutoscalingDeciders struct and provides a builder API.
type AutoscalingDecidersBuilder struct {
	v *AutoscalingDeciders
}

// NewAutoscalingDeciders provides a builder for the AutoscalingDeciders struct.
func NewAutoscalingDeciders() *AutoscalingDecidersBuilder {
	r := AutoscalingDecidersBuilder{
		&AutoscalingDeciders{
			Deciders: make(map[string]AutoscalingDecider, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AutoscalingDeciders struct
func (rb *AutoscalingDecidersBuilder) Build() AutoscalingDeciders {
	return *rb.v
}

// CurrentCapacity set the CurrentCapacity property for AutoscalingDecidersBuilder.
func (rb *AutoscalingDecidersBuilder) CurrentCapacity(currentcapacity AutoscalingCapacity) *AutoscalingDecidersBuilder {
	rb.v.CurrentCapacity = currentcapacity
	return rb
}

// CurrentNodes set the CurrentNodes property for AutoscalingDecidersBuilder.
func (rb *AutoscalingDecidersBuilder) CurrentNodes(current_nodes ...AutoscalingNode) *AutoscalingDecidersBuilder {
	rb.v.CurrentNodes = append(rb.v.CurrentNodes, current_nodes...)
	return rb
}

// Deciders set the Deciders property for AutoscalingDecidersBuilder.
func (rb *AutoscalingDecidersBuilder) Deciders(value map[string]AutoscalingDecider) *AutoscalingDecidersBuilder {
	rb.v.Deciders = value
	return rb
}

// RequiredCapacity set the RequiredCapacity property for AutoscalingDecidersBuilder.
func (rb *AutoscalingDecidersBuilder) RequiredCapacity(requiredcapacity AutoscalingCapacity) *AutoscalingDecidersBuilder {
	rb.v.RequiredCapacity = requiredcapacity
	return rb
}
