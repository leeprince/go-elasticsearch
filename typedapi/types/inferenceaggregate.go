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

// InferenceAggregate type.
type InferenceAggregate struct {
	Data              map[string]interface{}       `json:"data,omitempty"`
	FeatureImportance []InferenceFeatureImportance `json:"feature_importance,omitempty"`
	Meta              *Metadata                    `json:"meta,omitempty"`
	TopClasses        []InferenceTopClassEntry     `json:"top_classes,omitempty"`
	Value             *FieldValue                  `json:"value,omitempty"`
	Warning           *string                      `json:"warning,omitempty"`
}

// InferenceAggregateBuilder holds InferenceAggregate struct and provides a builder API.
type InferenceAggregateBuilder struct {
	v *InferenceAggregate
}

// NewInferenceAggregate provides a builder for the InferenceAggregate struct.
func NewInferenceAggregate() *InferenceAggregateBuilder {
	r := InferenceAggregateBuilder{
		&InferenceAggregate{
			Data: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InferenceAggregate struct
func (rb *InferenceAggregateBuilder) Build() InferenceAggregate {
	return *rb.v
}

// Data set the Data property for InferenceAggregateBuilder.
func (rb *InferenceAggregateBuilder) Data(value map[string]interface{}) *InferenceAggregateBuilder {
	rb.v.Data = value
	return rb
}

// FeatureImportance set the FeatureImportance property for InferenceAggregateBuilder.
func (rb *InferenceAggregateBuilder) FeatureImportance(feature_importance ...InferenceFeatureImportance) *InferenceAggregateBuilder {
	rb.v.FeatureImportance = append(rb.v.FeatureImportance, feature_importance...)
	return rb
}

// Meta set the Meta property for InferenceAggregateBuilder.
func (rb *InferenceAggregateBuilder) Meta(meta Metadata) *InferenceAggregateBuilder {
	rb.v.Meta = &meta
	return rb
}

// TopClasses set the TopClasses property for InferenceAggregateBuilder.
func (rb *InferenceAggregateBuilder) TopClasses(top_classes ...InferenceTopClassEntry) *InferenceAggregateBuilder {
	rb.v.TopClasses = append(rb.v.TopClasses, top_classes...)
	return rb
}

// Value set the Value property for InferenceAggregateBuilder.
func (rb *InferenceAggregateBuilder) Value(value FieldValue) *InferenceAggregateBuilder {
	rb.v.Value = &value
	return rb
}

// Warning set the Warning property for InferenceAggregateBuilder.
func (rb *InferenceAggregateBuilder) Warning(warning string) *InferenceAggregateBuilder {
	rb.v.Warning = &warning
	return rb
}
