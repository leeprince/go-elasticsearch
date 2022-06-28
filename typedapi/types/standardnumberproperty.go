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

// StandardNumberProperty type.
type StandardNumberProperty struct {
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
	OnScriptError    *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	Script           *Script                                    `json:"script,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
}

// StandardNumberPropertyBuilder holds StandardNumberProperty struct and provides a builder API.
type StandardNumberPropertyBuilder struct {
	v *StandardNumberProperty
}

// NewStandardNumberProperty provides a builder for the StandardNumberProperty struct.
func NewStandardNumberProperty() *StandardNumberPropertyBuilder {
	r := StandardNumberPropertyBuilder{
		&StandardNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the StandardNumberProperty struct
func (rb *StandardNumberPropertyBuilder) Build() StandardNumberProperty {
	return *rb.v
}

// Coerce set the Coerce property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Coerce(coerce bool) *StandardNumberPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) CopyTo(copyto Fields) *StandardNumberPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) DocValues(docvalues bool) *StandardNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *StandardNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Fields(value map[PropertyName]Property) *StandardNumberPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *StandardNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *StandardNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Index(index bool) *StandardNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) LocalMetadata(localmetadata Metadata) *StandardNumberPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Meta(value map[string]string) *StandardNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// OnScriptError set the OnScriptError property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *StandardNumberPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

// Properties set the Properties property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Properties(value map[PropertyName]Property) *StandardNumberPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Script set the Script property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Script(script Script) *StandardNumberPropertyBuilder {
	rb.v.Script = &script
	return rb
}

// Similarity set the Similarity property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Similarity(similarity string) *StandardNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) Store(store bool) *StandardNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesMetric set the TimeSeriesMetric property for StandardNumberPropertyBuilder.
func (rb *StandardNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *StandardNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}
