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

// ByteNumberProperty type.
type ByteNumberProperty struct {
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
	NullValue        *byte                                      `json:"null_value,omitempty"`
	OnScriptError    *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	Script           *Script                                    `json:"script,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// ByteNumberPropertyBuilder holds ByteNumberProperty struct and provides a builder API.
type ByteNumberPropertyBuilder struct {
	v *ByteNumberProperty
}

// NewByteNumberProperty provides a builder for the ByteNumberProperty struct.
func NewByteNumberProperty() *ByteNumberPropertyBuilder {
	r := ByteNumberPropertyBuilder{
		&ByteNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "byte"

	return &r
}

// Build finalize the chain and returns the ByteNumberProperty struct
func (rb *ByteNumberPropertyBuilder) Build() ByteNumberProperty {
	return *rb.v
}

// Coerce set the Coerce property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Coerce(coerce bool) *ByteNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) CopyTo(copyto Fields) *ByteNumberPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) DocValues(docvalues bool) *ByteNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *ByteNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Fields(value map[PropertyName]Property) *ByteNumberPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *ByteNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *ByteNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Index(index bool) *ByteNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) LocalMetadata(localmetadata Metadata) *ByteNumberPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Meta(value map[string]string) *ByteNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) NullValue(nullvalue byte) *ByteNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// OnScriptError set the OnScriptError property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *ByteNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

// Properties set the Properties property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Properties(value map[PropertyName]Property) *ByteNumberPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Script set the Script property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Script(script Script) *ByteNumberPropertyBuilder {
	rb.v.Script = &script
	return rb
}

// Similarity set the Similarity property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Similarity(similarity string) *ByteNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) Store(store bool) *ByteNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesMetric set the TimeSeriesMetric property for ByteNumberPropertyBuilder.
func (rb *ByteNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *ByteNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}

// Type set the Type property for ByteNumberPropertyBuilder.
