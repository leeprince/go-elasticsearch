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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/minimuminterval"
)

// AutoDateHistogramAggregation type.
type AutoDateHistogramAggregation struct {
	Buckets         *int                             `json:"buckets,omitempty"`
	Field           *Field                           `json:"field,omitempty"`
	Format          *string                          `json:"format,omitempty"`
	Meta            *Metadata                        `json:"meta,omitempty"`
	MinimumInterval *minimuminterval.MinimumInterval `json:"minimum_interval,omitempty"`
	Missing         *DateString                      `json:"missing,omitempty"`
	Name            *string                          `json:"name,omitempty"`
	Offset          *string                          `json:"offset,omitempty"`
	Params          map[string]interface{}           `json:"params,omitempty"`
	Script          *Script                          `json:"script,omitempty"`
	TimeZone        *string                          `json:"time_zone,omitempty"`
}

// AutoDateHistogramAggregationBuilder holds AutoDateHistogramAggregation struct and provides a builder API.
type AutoDateHistogramAggregationBuilder struct {
	v *AutoDateHistogramAggregation
}

// NewAutoDateHistogramAggregation provides a builder for the AutoDateHistogramAggregation struct.
func NewAutoDateHistogramAggregation() *AutoDateHistogramAggregationBuilder {
	r := AutoDateHistogramAggregationBuilder{
		&AutoDateHistogramAggregation{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AutoDateHistogramAggregation struct
func (rb *AutoDateHistogramAggregationBuilder) Build() AutoDateHistogramAggregation {
	return *rb.v
}

// Buckets set the Buckets property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Buckets(buckets int) *AutoDateHistogramAggregationBuilder {
	rb.v.Buckets = &buckets
	return rb
}

// Field set the Field property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Field(field Field) *AutoDateHistogramAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// Format set the Format property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Format(format string) *AutoDateHistogramAggregationBuilder {
	rb.v.Format = &format
	return rb
}

// Meta set the Meta property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Meta(meta Metadata) *AutoDateHistogramAggregationBuilder {
	rb.v.Meta = &meta
	return rb
}

// MinimumInterval set the MinimumInterval property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) MinimumInterval(minimuminterval minimuminterval.MinimumInterval) *AutoDateHistogramAggregationBuilder {
	rb.v.MinimumInterval = &minimuminterval
	return rb
}

// Missing set the Missing property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Missing(missing DateString) *AutoDateHistogramAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Name set the Name property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Name(name string) *AutoDateHistogramAggregationBuilder {
	rb.v.Name = &name
	return rb
}

// Offset set the Offset property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Offset(offset string) *AutoDateHistogramAggregationBuilder {
	rb.v.Offset = &offset
	return rb
}

// Params set the Params property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Params(value map[string]interface{}) *AutoDateHistogramAggregationBuilder {
	rb.v.Params = value
	return rb
}

// Script set the Script property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) Script(script Script) *AutoDateHistogramAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// TimeZone set the TimeZone property for AutoDateHistogramAggregationBuilder.
func (rb *AutoDateHistogramAggregationBuilder) TimeZone(timezone string) *AutoDateHistogramAggregationBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
