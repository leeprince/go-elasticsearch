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

// ScriptedMetricAggregation type.
type ScriptedMetricAggregation struct {
	CombineScript *Script                `json:"combine_script,omitempty"`
	Field         *Field                 `json:"field,omitempty"`
	InitScript    *Script                `json:"init_script,omitempty"`
	MapScript     *Script                `json:"map_script,omitempty"`
	Missing       *Missing               `json:"missing,omitempty"`
	Params        map[string]interface{} `json:"params,omitempty"`
	ReduceScript  *Script                `json:"reduce_script,omitempty"`
	Script        *Script                `json:"script,omitempty"`
}

// ScriptedMetricAggregationBuilder holds ScriptedMetricAggregation struct and provides a builder API.
type ScriptedMetricAggregationBuilder struct {
	v *ScriptedMetricAggregation
}

// NewScriptedMetricAggregation provides a builder for the ScriptedMetricAggregation struct.
func NewScriptedMetricAggregation() *ScriptedMetricAggregationBuilder {
	r := ScriptedMetricAggregationBuilder{
		&ScriptedMetricAggregation{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ScriptedMetricAggregation struct
func (rb *ScriptedMetricAggregationBuilder) Build() ScriptedMetricAggregation {
	return *rb.v
}

// CombineScript set the CombineScript property for ScriptedMetricAggregationBuilder.
func (rb *ScriptedMetricAggregationBuilder) CombineScript(combinescript Script) *ScriptedMetricAggregationBuilder {
	rb.v.CombineScript = &combinescript
	return rb
}

// Field set the Field property for ScriptedMetricAggregationBuilder.
func (rb *ScriptedMetricAggregationBuilder) Field(field Field) *ScriptedMetricAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// InitScript set the InitScript property for ScriptedMetricAggregationBuilder.
func (rb *ScriptedMetricAggregationBuilder) InitScript(initscript Script) *ScriptedMetricAggregationBuilder {
	rb.v.InitScript = &initscript
	return rb
}

// MapScript set the MapScript property for ScriptedMetricAggregationBuilder.
func (rb *ScriptedMetricAggregationBuilder) MapScript(mapscript Script) *ScriptedMetricAggregationBuilder {
	rb.v.MapScript = &mapscript
	return rb
}

// Missing set the Missing property for ScriptedMetricAggregationBuilder.
func (rb *ScriptedMetricAggregationBuilder) Missing(missing Missing) *ScriptedMetricAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Params set the Params property for ScriptedMetricAggregationBuilder.
func (rb *ScriptedMetricAggregationBuilder) Params(value map[string]interface{}) *ScriptedMetricAggregationBuilder {
	rb.v.Params = value
	return rb
}

// ReduceScript set the ReduceScript property for ScriptedMetricAggregationBuilder.
func (rb *ScriptedMetricAggregationBuilder) ReduceScript(reducescript Script) *ScriptedMetricAggregationBuilder {
	rb.v.ReduceScript = &reducescript
	return rb
}

// Script set the Script property for ScriptedMetricAggregationBuilder.
func (rb *ScriptedMetricAggregationBuilder) Script(script Script) *ScriptedMetricAggregationBuilder {
	rb.v.Script = &script
	return rb
}
