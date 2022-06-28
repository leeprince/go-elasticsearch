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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

// ScaledFloatNumberProperty type.
type ScaledFloatNumberProperty struct {
	Coerce           *bool                                      `json:"coerce,omitempty"`
	CopyTo           *Fields                                    `json:"copy_to,omitempty"`
	DocValues        *bool                                      `json:"doc_values,omitempty"`
	Dynamic          *dynamicmapping.DynamicMapping             `json:"dynamic,omitempty"`
	Fields           map[PropertyName]Property                  `json:"fields,omitempty"`
	IgnoreAbove      *int                                       `json:"ignore_above,omitempty"`
	IgnoreMalformed  *bool                                      `json:"ignore_malformed,omitempty"`
	Index            *bool                                      `json:"index,omitempty"`
	LocalMetadata    *Metadata                                  `json:"local_metadata,omitempty"`
	Meta             map[string]string                          `json:"meta,omitempty"`
	NullValue        *float64                                   `json:"null_value,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	ScalingFactor    *float64                                   `json:"scaling_factor,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// ScaledFloatNumberPropertyBuilder holds ScaledFloatNumberProperty struct and provides a builder API.
type ScaledFloatNumberPropertyBuilder struct {
	v *ScaledFloatNumberProperty
}

// NewScaledFloatNumberProperty provides a builder for the ScaledFloatNumberProperty struct.
func NewScaledFloatNumberProperty() *ScaledFloatNumberPropertyBuilder {
	r := ScaledFloatNumberPropertyBuilder{
		&ScaledFloatNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "scaled_float"

	return &r
}

// Build finalize the chain and returns the ScaledFloatNumberProperty struct
func (rb *ScaledFloatNumberPropertyBuilder) Build() ScaledFloatNumberProperty {
	return *rb.v
}

// Coerce set the Coerce property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) Coerce(coerce bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) CopyTo(copyto Fields) *ScaledFloatNumberPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) DocValues(docvalues bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *ScaledFloatNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) Fields(value map[PropertyName]Property) *ScaledFloatNumberPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *ScaledFloatNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) Index(index bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) LocalMetadata(localmetadata Metadata) *ScaledFloatNumberPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) Meta(value map[string]string) *ScaledFloatNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) NullValue(nullvalue float64) *ScaledFloatNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// Properties set the Properties property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) Properties(value map[PropertyName]Property) *ScaledFloatNumberPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// ScalingFactor set the ScalingFactor property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) ScalingFactor(scalingfactor float64) *ScaledFloatNumberPropertyBuilder {
	rb.v.ScalingFactor = &scalingfactor
	return rb
}

// Similarity set the Similarity property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) Similarity(similarity string) *ScaledFloatNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) Store(store bool) *ScaledFloatNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesMetric set the TimeSeriesMetric property for ScaledFloatNumberPropertyBuilder.
func (rb *ScaledFloatNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *ScaledFloatNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}

// Type set the Type property for ScaledFloatNumberPropertyBuilder.
