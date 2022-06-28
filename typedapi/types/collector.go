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

// Collector type.
type Collector struct {
	Children    []Collector `json:"children,omitempty"`
	Name        string      `json:"name"`
	Reason      string      `json:"reason"`
	TimeInNanos int64       `json:"time_in_nanos"`
}

// CollectorBuilder holds Collector struct and provides a builder API.
type CollectorBuilder struct {
	v *Collector
}

// NewCollector provides a builder for the Collector struct.
func NewCollector() *CollectorBuilder {
	r := CollectorBuilder{
		&Collector{},
	}

	return &r
}

// Build finalize the chain and returns the Collector struct
func (rb *CollectorBuilder) Build() Collector {
	return *rb.v
}

// Children set the Children property for CollectorBuilder.
func (rb *CollectorBuilder) Children(children ...Collector) *CollectorBuilder {
	rb.v.Children = append(rb.v.Children, children...)
	return rb
}

// Name set the Name property for CollectorBuilder.
func (rb *CollectorBuilder) Name(name string) *CollectorBuilder {
	rb.v.Name = name
	return rb
}

// Reason set the Reason property for CollectorBuilder.
func (rb *CollectorBuilder) Reason(reason string) *CollectorBuilder {
	rb.v.Reason = reason
	return rb
}

// TimeInNanos set the TimeInNanos property for CollectorBuilder.
func (rb *CollectorBuilder) TimeInNanos(timeinnanos int64) *CollectorBuilder {
	rb.v.TimeInNanos = timeinnanos
	return rb
}
