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

// DateNanosProperty type.
type DateNanosProperty struct {
	Boost           *float64                       `json:"boost,omitempty"`
	CopyTo          *Fields                        `json:"copy_to,omitempty"`
	DocValues       *bool                          `json:"doc_values,omitempty"`
	Dynamic         *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields          map[PropertyName]Property      `json:"fields,omitempty"`
	Format          *string                        `json:"format,omitempty"`
	IgnoreAbove     *int                           `json:"ignore_above,omitempty"`
	IgnoreMalformed *bool                          `json:"ignore_malformed,omitempty"`
	Index           *bool                          `json:"index,omitempty"`
	LocalMetadata   *Metadata                      `json:"local_metadata,omitempty"`
	Meta            map[string]string              `json:"meta,omitempty"`
	NullValue       *DateString                    `json:"null_value,omitempty"`
	PrecisionStep   *int                           `json:"precision_step,omitempty"`
	Properties      map[PropertyName]Property      `json:"properties,omitempty"`
	Similarity      *string                        `json:"similarity,omitempty"`
	Store           *bool                          `json:"store,omitempty"`
	Type            string                         `json:"type,omitempty"`
}

// DateNanosPropertyBuilder holds DateNanosProperty struct and provides a builder API.
type DateNanosPropertyBuilder struct {
	v *DateNanosProperty
}

// NewDateNanosProperty provides a builder for the DateNanosProperty struct.
func NewDateNanosProperty() *DateNanosPropertyBuilder {
	r := DateNanosPropertyBuilder{
		&DateNanosProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "date_nanos"

	return &r
}

// Build finalize the chain and returns the DateNanosProperty struct
func (rb *DateNanosPropertyBuilder) Build() DateNanosProperty {
	return *rb.v
}

// Boost set the Boost property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Boost(boost float64) *DateNanosPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// CopyTo set the CopyTo property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) CopyTo(copyto Fields) *DateNanosPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) DocValues(docvalues bool) *DateNanosPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DateNanosPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Fields(value map[PropertyName]Property) *DateNanosPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// Format set the Format property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Format(format string) *DateNanosPropertyBuilder {
	rb.v.Format = &format
	return rb
}

// IgnoreAbove set the IgnoreAbove property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) IgnoreAbove(ignoreabove int) *DateNanosPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *DateNanosPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Index(index bool) *DateNanosPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// LocalMetadata set the LocalMetadata property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) LocalMetadata(localmetadata Metadata) *DateNanosPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Meta(value map[string]string) *DateNanosPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// NullValue set the NullValue property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) NullValue(nullvalue DateString) *DateNanosPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// PrecisionStep set the PrecisionStep property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) PrecisionStep(precisionstep int) *DateNanosPropertyBuilder {
	rb.v.PrecisionStep = &precisionstep
	return rb
}

// Properties set the Properties property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Properties(value map[PropertyName]Property) *DateNanosPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Similarity set the Similarity property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Similarity(similarity string) *DateNanosPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for DateNanosPropertyBuilder.
func (rb *DateNanosPropertyBuilder) Store(store bool) *DateNanosPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for DateNanosPropertyBuilder.
