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

// UnsignedLongNumberProperty type.
type UnsignedLongNumberProperty struct {
	CopyTo           *Fields                                    `json:"copy_to,omitempty"`
	DocValues        *bool                                      `json:"doc_values,omitempty"`
	Dynamic          *dynamicmapping.DynamicMapping             `json:"dynamic,omitempty"`
	Fields           map[PropertyName]Property                  `json:"fields,omitempty"`
	IgnoreAbove      *int                                       `json:"ignore_above,omitempty"`
	IgnoreMalformed  *bool                                      `json:"ignore_malformed,omitempty"`
	Index            *bool                                      `json:"index,omitempty"`
	LocalMetadata    *Metadata                                  `json:"local_metadata,omitempty"`
	Meta             map[string]string                          `json:"meta,omitempty"`
	NullValue        *uint64                                    `json:"null_value,omitempty"`
	Properties       map[PropertyName]Property                  `json:"properties,omitempty"`
	Similarity       *string                                    `json:"similarity,omitempty"`
	Store            *bool                                      `json:"store,omitempty"`
	TimeSeriesMetric *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type             string                                     `json:"type,omitempty"`
}

// UnsignedLongNumberPropertyBuilder holds UnsignedLongNumberProperty struct and provides a builder API.
type UnsignedLongNumberPropertyBuilder struct {
	v *UnsignedLongNumberProperty
}

// NewUnsignedLongNumberProperty provides a builder for the UnsignedLongNumberProperty struct.
func NewUnsignedLongNumberProperty() *UnsignedLongNumberPropertyBuilder {
	r := UnsignedLongNumberPropertyBuilder{
		&UnsignedLongNumberProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "unsigned_long"

	return &r
}

// Build finalize the chain and returns the UnsignedLongNumberProperty struct
func (rb *UnsignedLongNumberPropertyBuilder) Build() UnsignedLongNumberProperty {
	return *rb.v
}

// CopyTo set the CopyTo property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) CopyTo(copyto Fields) *UnsignedLongNumberPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) DocValues(docvalues bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *UnsignedLongNumberPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) Fields(value map[PropertyName]Property) *UnsignedLongNumberPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) IgnoreAbove(ignoreabove int) *UnsignedLongNumberPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) Index(index bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) LocalMetadata(localmetadata Metadata) *UnsignedLongNumberPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) Meta(value map[string]string) *UnsignedLongNumberPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) NullValue(nullvalue uint64) *UnsignedLongNumberPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// Properties set the Properties property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) Properties(value map[PropertyName]Property) *UnsignedLongNumberPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) Similarity(similarity string) *UnsignedLongNumberPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) Store(store bool) *UnsignedLongNumberPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeriesMetric set the TimeSeriesMetric property for UnsignedLongNumberPropertyBuilder.
func (rb *UnsignedLongNumberPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *UnsignedLongNumberPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}

// Type set the Type property for UnsignedLongNumberPropertyBuilder.
