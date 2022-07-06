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

// ScriptedMetricAggregate type.
type ScriptedMetricAggregate struct {
	Meta  *Metadata   `json:"meta,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// ScriptedMetricAggregateBuilder holds ScriptedMetricAggregate struct and provides a builder API.
type ScriptedMetricAggregateBuilder struct {
	v *ScriptedMetricAggregate
}

// NewScriptedMetricAggregate provides a builder for the ScriptedMetricAggregate struct.
func NewScriptedMetricAggregate() *ScriptedMetricAggregateBuilder {
	r := ScriptedMetricAggregateBuilder{
		&ScriptedMetricAggregate{},
	}

	return &r
}

// Build finalize the chain and returns the ScriptedMetricAggregate struct
func (rb *ScriptedMetricAggregateBuilder) Build() ScriptedMetricAggregate {
	return *rb.v
}

// Meta set the Meta property for ScriptedMetricAggregateBuilder.
func (rb *ScriptedMetricAggregateBuilder) Meta(meta Metadata) *ScriptedMetricAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// Value set the Value property for ScriptedMetricAggregateBuilder.
func (rb *ScriptedMetricAggregateBuilder) Value(value interface{}) *ScriptedMetricAggregateBuilder {
	rb.v.Value = value
	return rb
}
