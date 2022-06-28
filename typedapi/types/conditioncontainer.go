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

// ConditionContainer type.
type ConditionContainer struct {
	Always       *AlwaysCondition       `json:"always,omitempty"`
	ArrayCompare *ArrayCompareCondition `json:"array_compare,omitempty"`
	Compare      *CompareCondition      `json:"compare,omitempty"`
	Never        *NeverCondition        `json:"never,omitempty"`
	Script       *ScriptCondition       `json:"script,omitempty"`
}

// ConditionContainerBuilder holds ConditionContainer struct and provides a builder API.
type ConditionContainerBuilder struct {
	v *ConditionContainer
}

// NewConditionContainer provides a builder for the ConditionContainer struct.
func NewConditionContainer() *ConditionContainerBuilder {
	r := ConditionContainerBuilder{
		&ConditionContainer{},
	}

	return &r
}

// Build finalize the chain and returns the ConditionContainer struct
func (rb *ConditionContainerBuilder) Build() ConditionContainer {
	return *rb.v
}

// Always set the Always property for ConditionContainerBuilder.
func (rb *ConditionContainerBuilder) Always(always AlwaysCondition) *ConditionContainerBuilder {
	rb.v.Always = &always
	return rb
}

// ArrayCompare set the ArrayCompare property for ConditionContainerBuilder.
func (rb *ConditionContainerBuilder) ArrayCompare(arraycompare ArrayCompareCondition) *ConditionContainerBuilder {
	rb.v.ArrayCompare = &arraycompare
	return rb
}

// Compare set the Compare property for ConditionContainerBuilder.
func (rb *ConditionContainerBuilder) Compare(compare CompareCondition) *ConditionContainerBuilder {
	rb.v.Compare = &compare
	return rb
}

// Never set the Never property for ConditionContainerBuilder.
func (rb *ConditionContainerBuilder) Never(never NeverCondition) *ConditionContainerBuilder {
	rb.v.Never = &never
	return rb
}

// Script set the Script property for ConditionContainerBuilder.
func (rb *ConditionContainerBuilder) Script(script ScriptCondition) *ConditionContainerBuilder {
	rb.v.Script = &script
	return rb
}
