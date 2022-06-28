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

// IngestTotal type.
type IngestTotal struct {
	Count        *int64                      `json:"count,omitempty"`
	Current      *int64                      `json:"current,omitempty"`
	Failed       *int64                      `json:"failed,omitempty"`
	Processors   []map[string]KeyedProcessor `json:"processors,omitempty"`
	TimeInMillis *int64                      `json:"time_in_millis,omitempty"`
}

// IngestTotalBuilder holds IngestTotal struct and provides a builder API.
type IngestTotalBuilder struct {
	v *IngestTotal
}

// NewIngestTotal provides a builder for the IngestTotal struct.
func NewIngestTotal() *IngestTotalBuilder {
	r := IngestTotalBuilder{
		&IngestTotal{},
	}

	return &r
}

// Build finalize the chain and returns the IngestTotal struct
func (rb *IngestTotalBuilder) Build() IngestTotal {
	return *rb.v
}

// Count set the Count property for IngestTotalBuilder.
func (rb *IngestTotalBuilder) Count(count int64) *IngestTotalBuilder {
	rb.v.Count = &count
	return rb
}

// Current set the Current property for IngestTotalBuilder.
func (rb *IngestTotalBuilder) Current(current int64) *IngestTotalBuilder {
	rb.v.Current = &current
	return rb
}

// Failed set the Failed property for IngestTotalBuilder.
func (rb *IngestTotalBuilder) Failed(failed int64) *IngestTotalBuilder {
	rb.v.Failed = &failed
	return rb
}

// Processors set the Processors property for IngestTotalBuilder.
func (rb *IngestTotalBuilder) Processors(value ...map[string]KeyedProcessor) *IngestTotalBuilder {
	rb.v.Processors = append(rb.v.Processors, value...)
	return rb
}

// TimeInMillis set the TimeInMillis property for IngestTotalBuilder.
func (rb *IngestTotalBuilder) TimeInMillis(timeinmillis int64) *IngestTotalBuilder {
	rb.v.TimeInMillis = &timeinmillis
	return rb
}
