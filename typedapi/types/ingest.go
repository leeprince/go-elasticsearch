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

// Ingest type.
type Ingest struct {
	Pipelines map[string]IngestTotal `json:"pipelines,omitempty"`
	Total     *IngestTotal           `json:"total,omitempty"`
}

// IngestBuilder holds Ingest struct and provides a builder API.
type IngestBuilder struct {
	v *Ingest
}

// NewIngest provides a builder for the Ingest struct.
func NewIngest() *IngestBuilder {
	r := IngestBuilder{
		&Ingest{
			Pipelines: make(map[string]IngestTotal, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Ingest struct
func (rb *IngestBuilder) Build() Ingest {
	return *rb.v
}

// Pipelines set the Pipelines property for IngestBuilder.
func (rb *IngestBuilder) Pipelines(value map[string]IngestTotal) *IngestBuilder {
	rb.v.Pipelines = value
	return rb
}

// Total set the Total property for IngestBuilder.
func (rb *IngestBuilder) Total(total IngestTotal) *IngestBuilder {
	rb.v.Total = &total
	return rb
}
