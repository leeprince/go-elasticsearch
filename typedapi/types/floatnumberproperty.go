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

// FloatNumberProperty type.
type FloatNumberProperty struct {
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
	NullValue        *float32                                   `json:"null_value,omitempty"`
	OnScriptError    *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	Script           *Script                                    `json:"script,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// FloatNumberPropertyBuilder holds FloatNumberProperty struct and provides a builder API.
type FloatNumberPropertyBuilder struct {
	v *FloatNumberProperty
}

// NewFloatNumberProperty provides a builder for the FloatNumberProperty struct.
func NewFloatNumberProperty() *FloatNumberPropertyBuilder {
	r := FloatNumberPropertyBuilder{
		&FloatNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "float"

	return &r
}

// Build finalize the chain and returns the FloatNumberProperty struct
func (rb *FloatNumberPropertyBuilder) Build() FloatNumberProperty {
	return *rb.v
}

// Coerce set the Coerce property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Coerce(coerce bool) *FloatNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) CopyTo(copyto Fields) *FloatNumberPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) DocValues(docvalues bool) *FloatNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *FloatNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Fields(value map[PropertyName]Property) *FloatNumberPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *FloatNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *FloatNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Index(index bool) *FloatNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) LocalMetadata(localmetadata Metadata) *FloatNumberPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Meta(value map[string]string) *FloatNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) NullValue(nullvalue float32) *FloatNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// OnScriptError set the OnScriptError property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *FloatNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

// Properties set the Properties property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Properties(value map[PropertyName]Property) *FloatNumberPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Script set the Script property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Script(script Script) *FloatNumberPropertyBuilder {
	rb.v.Script = &script
	return rb
}

// Similarity set the Similarity property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Similarity(similarity string) *FloatNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) Store(store bool) *FloatNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesMetric set the TimeSeriesMetric property for FloatNumberPropertyBuilder.
func (rb *FloatNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *FloatNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}

// Type set the Type property for FloatNumberPropertyBuilder.
