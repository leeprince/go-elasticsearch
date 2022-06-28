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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/quantifier"
)

// ArrayCompareCondition type.
type ArrayCompareCondition struct {
	ArrayPath  string                `json:"array_path"`
	Comparison string                `json:"comparison"`
	Path       string                `json:"path"`
	Quantifier quantifier.Quantifier `json:"quantifier"`
	Value      interface{}           `json:"value,omitempty"`
}

// ArrayCompareConditionBuilder holds ArrayCompareCondition struct and provides a builder API.
type ArrayCompareConditionBuilder struct {
	v *ArrayCompareCondition
}

// NewArrayCompareCondition provides a builder for the ArrayCompareCondition struct.
func NewArrayCompareCondition() *ArrayCompareConditionBuilder {
	r := ArrayCompareConditionBuilder{
		&ArrayCompareCondition{},
	}

	return &r
}

// Build finalize the chain and returns the ArrayCompareCondition struct
func (rb *ArrayCompareConditionBuilder) Build() ArrayCompareCondition {
	return *rb.v
}

// ArrayPath set the ArrayPath property for ArrayCompareConditionBuilder.
func (rb *ArrayCompareConditionBuilder) ArrayPath(arraypath string) *ArrayCompareConditionBuilder {
	rb.v.ArrayPath = arraypath
	return rb
}

// Comparison set the Comparison property for ArrayCompareConditionBuilder.
func (rb *ArrayCompareConditionBuilder) Comparison(comparison string) *ArrayCompareConditionBuilder {
	rb.v.Comparison = comparison
	return rb
}

// Path set the Path property for ArrayCompareConditionBuilder.
func (rb *ArrayCompareConditionBuilder) Path(path string) *ArrayCompareConditionBuilder {
	rb.v.Path = path
	return rb
}

// Quantifier set the Quantifier property for ArrayCompareConditionBuilder.
func (rb *ArrayCompareConditionBuilder) Quantifier(quantifier quantifier.Quantifier) *ArrayCompareConditionBuilder {
	rb.v.Quantifier = quantifier
	return rb
}

// Value set the Value property for ArrayCompareConditionBuilder.
func (rb *ArrayCompareConditionBuilder) Value(value interface{}) *ArrayCompareConditionBuilder {
	rb.v.Value = value
	return rb
}
