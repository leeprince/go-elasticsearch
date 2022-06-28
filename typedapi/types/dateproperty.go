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

// DateProperty type.
type DateProperty struct {
	Boost           *float64                       `json:"boost,omitempty"`
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fielddata       *NumericFielddata              `json:"fielddata,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	Format          *string                        `json:"format,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	Index           *bool                          `json:"index,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Locale          *string                        `json:"locale,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	NullValue       *DateString                    `json:"null_value,omitempty"`
	PrecisionStep   *int                           `json:"precision_step,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// DatePropertyBuilder holds DateProperty struct and provides a builder API.
type DatePropertyBuilder struct {
	v *DateProperty
}

// NewDateProperty provides a builder for the DateProperty struct.
func NewDateProperty() *DatePropertyBuilder {
	r := DatePropertyBuilder{
		&DateProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "date"

	return &r
}

// Build finalize the chain and returns the DateProperty struct
func (rb *DatePropertyBuilder) Build() DateProperty {
	return *rb.v
}

// Boost set the Boost property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Boost(boost float64) *DatePropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// CopyTo set the CopyTo property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) CopyTo(copyto Fields) *DatePropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) DocValues(docvalues bool) *DatePropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DatePropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fielddata set the Fielddata property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Fielddata(fielddata NumericFielddata) *DatePropertyBuilder {
	rb.v.Fielddata = &fielddata
	return rb
}

// Fields set the Fields property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Fields(value map[PropertyName]Property) *DatePropertyBuilder {
	rb.v.Fields = value
	return rb
}

// Format set the Format property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Format(format string) *DatePropertyBuilder {
	rb.v.Format = &format
	return rb
}

// IgnoreAbove set the IgnoreAbove property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) IgnoreAbove(ignoreabove int) *DatePropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) IgnoreMalformed(ignoremalformed bool) *DatePropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Index(index bool) *DatePropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) LocalMetadata(localmetadata Metadata) *DatePropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Locale set the Locale property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Locale(locale string) *DatePropertyBuilder {
	rb.v.Locale = &locale
	return rb
}

// Meta set the Meta property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Meta(value map[string]string) *DatePropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) NullValue(nullvalue DateString) *DatePropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// PrecisionStep set the PrecisionStep property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) PrecisionStep(precisionstep int) *DatePropertyBuilder {
	rb.v.PrecisionStep = &precisionstep
	return rb
}

// Properties set the Properties property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Properties(value map[PropertyName]Property) *DatePropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Similarity(similarity string) *DatePropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for DatePropertyBuilder.
func (rb *DatePropertyBuilder) Store(store bool) *DatePropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for DatePropertyBuilder.
