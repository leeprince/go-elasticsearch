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

// ShortNumberProperty type.
type ShortNumberProperty struct {
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

// ShortNumberPropertyBuilder holds ShortNumberProperty struct and provides a builder API.
type ShortNumberPropertyBuilder struct {
	v *ShortNumberProperty
}

// NewShortNumberProperty provides a builder for the ShortNumberProperty struct.
func NewShortNumberProperty() *ShortNumberPropertyBuilder {
	r := ShortNumberPropertyBuilder{
		&ShortNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "short"

	return &r
}

// Build finalize the chain and returns the ShortNumberProperty struct
func (rb *ShortNumberPropertyBuilder) Build() ShortNumberProperty {
	return *rb.v
}

// Coerce set the Coerce property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Coerce(coerce bool) *ShortNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) CopyTo(copyto Fields) *ShortNumberPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) DocValues(docvalues bool) *ShortNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *ShortNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Fields(value map[PropertyName]Property) *ShortNumberPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *ShortNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *ShortNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Index(index bool) *ShortNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) LocalMetadata(localmetadata Metadata) *ShortNumberPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Meta(value map[string]string) *ShortNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) NullValue(nullvalue int) *ShortNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// OnScriptError set the OnScriptError property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *ShortNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

// Properties set the Properties property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Properties(value map[PropertyName]Property) *ShortNumberPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Script set the Script property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Script(script Script) *ShortNumberPropertyBuilder {
	rb.v.Script = &script
	return rb
}

// Similarity set the Similarity property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Similarity(similarity string) *ShortNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) Store(store bool) *ShortNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesMetric set the TimeSeriesMetric property for ShortNumberPropertyBuilder.
func (rb *ShortNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *ShortNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}

// Type set the Type property for ShortNumberPropertyBuilder.
