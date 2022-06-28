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

// TypeMapping type.
type TypeMapping struct {
	AllField             *AllField                      `json:"all_field,omitempty"`
	DataStreamTimestamp_ *DataStreamTimestamp           `json:"_data_stream_timestamp,omitempty"`
	DateDetection        *bool                          `json:"date_detection,omitempty"`
	Dynamic              *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	DynamicDateFormats   []string                       `json:"dynamic_date_formats,omitempty"`
	DynamicTemplates     []map[string]DynamicTemplate   `json:"dynamic_templates,omitempty"`
	Enabled              *bool                          `json:"enabled,omitempty"`
	FieldNames_          *FieldNamesField               `json:"_field_names,omitempty"`
	IndexField           *IndexField                    `json:"index_field,omitempty"`
	Meta_                *Metadata                      `json:"_meta,omitempty"`
	NumericDetection     *bool                          `json:"numeric_detection,omitempty"`
	Properties           map[PropertyName]Property      `json:"properties,omitempty"`
	Routing_             *RoutingField                  `json:"_routing,omitempty"`
	Runtime              map[string]RuntimeField        `json:"runtime,omitempty"`
	Size_                *SizeField                     `json:"_size,omitempty"`
	Source_              *SourceField                   `json:"_source,omitempty"`
}

// TypeMappingBuilder holds TypeMapping struct and provides a builder API.
type TypeMappingBuilder struct {
	v *TypeMapping
}

// NewTypeMapping provides a builder for the TypeMapping struct.
func NewTypeMapping() *TypeMappingBuilder {
	r := TypeMappingBuilder{
		&TypeMapping{
			Properties: make(map[PropertyName]Property, 0),
			Runtime:    make(map[string]RuntimeField, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TypeMapping struct
func (rb *TypeMappingBuilder) Build() TypeMapping {
	return *rb.v
}

// AllField set the AllField property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) AllField(allfield AllField) *TypeMappingBuilder {
	rb.v.AllField = &allfield
	return rb
}

// DataStreamTimestamp_ set the DataStreamTimestamp_ property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) DataStreamTimestamp_(datastreamtimestamp_ DataStreamTimestamp) *TypeMappingBuilder {
	rb.v.DataStreamTimestamp_ = &datastreamtimestamp_
	return rb
}

// DateDetection set the DateDetection property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) DateDetection(datedetection bool) *TypeMappingBuilder {
	rb.v.DateDetection = &datedetection
	return rb
}

// Dynamic set the Dynamic property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *TypeMappingBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// DynamicDateFormats set the DynamicDateFormats property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) DynamicDateFormats(dynamic_date_formats ...string) *TypeMappingBuilder {
	rb.v.DynamicDateFormats = append(rb.v.DynamicDateFormats, dynamic_date_formats...)
	return rb
}

// DynamicTemplates set the DynamicTemplates property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) DynamicTemplates(arg []map[string]DynamicTemplate) *TypeMappingBuilder {
	rb.v.DynamicTemplates = arg
	return rb
}

// Enabled set the Enabled property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) Enabled(enabled bool) *TypeMappingBuilder {
	rb.v.Enabled = &enabled
	return rb
}

// FieldNames_ set the FieldNames_ property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) FieldNames_(fieldnames_ FieldNamesField) *TypeMappingBuilder {
	rb.v.FieldNames_ = &fieldnames_
	return rb
}

// IndexField set the IndexField property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) IndexField(indexfield IndexField) *TypeMappingBuilder {
	rb.v.IndexField = &indexfield
	return rb
}

// Meta_ set the Meta_ property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) Meta_(meta_ Metadata) *TypeMappingBuilder {
	rb.v.Meta_ = &meta_
	return rb
}

// NumericDetection set the NumericDetection property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) NumericDetection(numericdetection bool) *TypeMappingBuilder {
	rb.v.NumericDetection = &numericdetection
	return rb
}

// Properties set the Properties property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) Properties(value map[PropertyName]Property) *TypeMappingBuilder {
	rb.v.Properties = value
	return rb
}

// Routing_ set the Routing_ property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) Routing_(routing_ RoutingField) *TypeMappingBuilder {
	rb.v.Routing_ = &routing_
	return rb
}

// Runtime set the Runtime property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) Runtime(value map[string]RuntimeField) *TypeMappingBuilder {
	rb.v.Runtime = value
	return rb
}

// Size_ set the Size_ property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) Size_(size_ SizeField) *TypeMappingBuilder {
	rb.v.Size_ = &size_
	return rb
}

// Source_ set the Source_ property for TypeMappingBuilder.
func (rb *TypeMappingBuilder) Source_(source_ SourceField) *TypeMappingBuilder {
	rb.v.Source_ = &source_
	return rb
}
