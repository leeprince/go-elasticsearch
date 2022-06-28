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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

// IntegerNumberProperty type.
type IntegerNumberProperty struct {
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
	NullValue        *int                                       `json:"null_value,omitempty"`
	OnScriptError    *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	Script           *Script                                    `json:"script,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// IntegerNumberPropertyBuilder holds IntegerNumberProperty struct and provides a builder API.
type IntegerNumberPropertyBuilder struct {
	v *IntegerNumberProperty
}

// NewIntegerNumberProperty provides a builder for the IntegerNumberProperty struct.
func NewIntegerNumberProperty() *IntegerNumberPropertyBuilder {
	r := IntegerNumberPropertyBuilder{
		&IntegerNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "integer"

	return &r
}

// Build finalize the chain and returns the IntegerNumberProperty struct
func (rb *IntegerNumberPropertyBuilder) Build() IntegerNumberProperty {
	return *rb.v
}

// Coerce set the Coerce property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Coerce(coerce bool) *IntegerNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) CopyTo(copyto Fields) *IntegerNumberPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) DocValues(docvalues bool) *IntegerNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *IntegerNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Fields(value map[PropertyName]Property) *IntegerNumberPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *IntegerNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *IntegerNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Index(index bool) *IntegerNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) LocalMetadata(localmetadata Metadata) *IntegerNumberPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Meta(value map[string]string) *IntegerNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) NullValue(nullvalue int) *IntegerNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// OnScriptError set the OnScriptError property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *IntegerNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

// Properties set the Properties property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Properties(value map[PropertyName]Property) *IntegerNumberPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Script set the Script property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Script(script Script) *IntegerNumberPropertyBuilder {
	rb.v.Script = &script
	return rb
}

// Similarity set the Similarity property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Similarity(similarity string) *IntegerNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) Store(store bool) *IntegerNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesMetric set the TimeSeriesMetric property for IntegerNumberPropertyBuilder.
func (rb *IntegerNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *IntegerNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}

// Type set the Type property for IntegerNumberPropertyBuilder.
