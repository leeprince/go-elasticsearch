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

// FunctionScoreContainer type.
type FunctionScoreContainer struct {
	Exp              *DecayFunction                 `json:"exp,omitempty"`
	FieldValueFactor *FieldValueFactorScoreFunction `json:"field_value_factor,omitempty"`
	Filter           *QueryContainer                `json:"filter,omitempty"`
	Gauss            *DecayFunction                 `json:"gauss,omitempty"`
	Linear           *DecayFunction                 `json:"linear,omitempty"`
	RandomScore      *RandomScoreFunction           `json:"random_score,omitempty"`
	ScriptScore      *ScriptScoreFunction           `json:"script_score,omitempty"`
	Weight           *float64                       `json:"weight,omitempty"`
}

// FunctionScoreContainerBuilder holds FunctionScoreContainer struct and provides a builder API.
type FunctionScoreContainerBuilder struct {
	v *FunctionScoreContainer
}

// NewFunctionScoreContainer provides a builder for the FunctionScoreContainer struct.
func NewFunctionScoreContainer() *FunctionScoreContainerBuilder {
	r := FunctionScoreContainerBuilder{
		&FunctionScoreContainer{},
	}

	return &r
}

// Build finalize the chain and returns the FunctionScoreContainer struct
func (rb *FunctionScoreContainerBuilder) Build() FunctionScoreContainer {
	return *rb.v
}

// Exp set the Exp property for FunctionScoreContainerBuilder.
func (rb *FunctionScoreContainerBuilder) Exp(exp DecayFunction) *FunctionScoreContainerBuilder {
	rb.v.Exp = &exp
	return rb
}

// FieldValueFactor set the FieldValueFactor property for FunctionScoreContainerBuilder.
func (rb *FunctionScoreContainerBuilder) FieldValueFactor(fieldvaluefactor FieldValueFactorScoreFunction) *FunctionScoreContainerBuilder {
	rb.v.FieldValueFactor = &fieldvaluefactor
	return rb
}

// Filter set the Filter property for FunctionScoreContainerBuilder.
func (rb *FunctionScoreContainerBuilder) Filter(filter QueryContainer) *FunctionScoreContainerBuilder {
	rb.v.Filter = &filter
	return rb
}

// Gauss set the Gauss property for FunctionScoreContainerBuilder.
func (rb *FunctionScoreContainerBuilder) Gauss(gauss DecayFunction) *FunctionScoreContainerBuilder {
	rb.v.Gauss = &gauss
	return rb
}

// Linear set the Linear property for FunctionScoreContainerBuilder.
func (rb *FunctionScoreContainerBuilder) Linear(linear DecayFunction) *FunctionScoreContainerBuilder {
	rb.v.Linear = &linear
	return rb
}

// RandomScore set the RandomScore property for FunctionScoreContainerBuilder.
func (rb *FunctionScoreContainerBuilder) RandomScore(randomscore RandomScoreFunction) *FunctionScoreContainerBuilder {
	rb.v.RandomScore = &randomscore
	return rb
}

// ScriptScore set the ScriptScore property for FunctionScoreContainerBuilder.
func (rb *FunctionScoreContainerBuilder) ScriptScore(scriptscore ScriptScoreFunction) *FunctionScoreContainerBuilder {
	rb.v.ScriptScore = &scriptscore
	return rb
}

// Weight set the Weight property for FunctionScoreContainerBuilder.
func (rb *FunctionScoreContainerBuilder) Weight(weight float64) *FunctionScoreContainerBuilder {
	rb.v.Weight = &weight
	return rb
}
