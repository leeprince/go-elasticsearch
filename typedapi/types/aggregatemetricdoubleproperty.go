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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// AggregateMetricDoubleProperty type.
type AggregateMetricDoubleProperty struct {
	DefaultMetric string                         `json:"default_metric"`
	Dynamic       *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields        map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove   *int                           `json:"ignore_above,omitempty"`
	LocalMetadata *Metadata                      `json:"local_metadata,omitempty"`
	Meta          map[string]string              `json:"meta,omitempty"`
	Metrics       []string                       `json:"metrics"`
	Properties    map[PropertyName]Property      `json:"properties,omitempty"`
	Type          string                         `json:"type,omitempty"`
}

// AggregateMetricDoublePropertyBuilder holds AggregateMetricDoubleProperty struct and provides a builder API.
type AggregateMetricDoublePropertyBuilder struct {
	v *AggregateMetricDoubleProperty
}

// NewAggregateMetricDoubleProperty provides a builder for the AggregateMetricDoubleProperty struct.
func NewAggregateMetricDoubleProperty() *AggregateMetricDoublePropertyBuilder {
	r := AggregateMetricDoublePropertyBuilder{
		&AggregateMetricDoubleProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "aggregate_metric_double"

	return &r
}

// Build finalize the chain and returns the AggregateMetricDoubleProperty struct
func (rb *AggregateMetricDoublePropertyBuilder) Build() AggregateMetricDoubleProperty {
	return *rb.v
}

// DefaultMetric set the DefaultMetric property for AggregateMetricDoublePropertyBuilder.
func (rb *AggregateMetricDoublePropertyBuilder) DefaultMetric(defaultmetric string) *AggregateMetricDoublePropertyBuilder {
	rb.v.DefaultMetric = defaultmetric
	return rb
}

// Dynamic set the Dynamic property for AggregateMetricDoublePropertyBuilder.
func (rb *AggregateMetricDoublePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *AggregateMetricDoublePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for AggregateMetricDoublePropertyBuilder.
func (rb *AggregateMetricDoublePropertyBuilder) Fields(value map[PropertyName]Property) *AggregateMetricDoublePropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for AggregateMetricDoublePropertyBuilder.
func (rb *AggregateMetricDoublePropertyBuilder) IgnoreAbove(ignoreabove int) *AggregateMetricDoublePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for AggregateMetricDoublePropertyBuilder.
func (rb *AggregateMetricDoublePropertyBuilder) LocalMetadata(localmetadata Metadata) *AggregateMetricDoublePropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for AggregateMetricDoublePropertyBuilder.
func (rb *AggregateMetricDoublePropertyBuilder) Meta(value map[string]string) *AggregateMetricDoublePropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Metrics set the Metrics property for AggregateMetricDoublePropertyBuilder.
func (rb *AggregateMetricDoublePropertyBuilder) Metrics(metrics ...string) *AggregateMetricDoublePropertyBuilder {
	rb.v.Metrics = append(rb.v.Metrics, metrics...)
	return rb
}

// Properties set the Properties property for AggregateMetricDoublePropertyBuilder.
func (rb *AggregateMetricDoublePropertyBuilder) Properties(value map[PropertyName]Property) *AggregateMetricDoublePropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Type set the Type property for AggregateMetricDoublePropertyBuilder.
