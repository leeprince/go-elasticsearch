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

// ResolveIndexItem type.
type ResolveIndexItem struct {
	Aliases    []string        `json:"aliases,omitempty"`
	Attributes []string        `json:"attributes"`
	DataStream *DataStreamName `json:"data_stream,omitempty"`
	Name       Name            `json:"name"`
}

// ResolveIndexItemBuilder holds ResolveIndexItem struct and provides a builder API.
type ResolveIndexItemBuilder struct {
	v *ResolveIndexItem
}

// NewResolveIndexItem provides a builder for the ResolveIndexItem struct.
func NewResolveIndexItem() *ResolveIndexItemBuilder {
	r := ResolveIndexItemBuilder{
		&ResolveIndexItem{},
	}

	return &r
}

// Build finalize the chain and returns the ResolveIndexItem struct
func (rb *ResolveIndexItemBuilder) Build() ResolveIndexItem {
	return *rb.v
}

// Aliases set the Aliases property for ResolveIndexItemBuilder.
func (rb *ResolveIndexItemBuilder) Aliases(aliases ...string) *ResolveIndexItemBuilder {
	rb.v.Aliases = append(rb.v.Aliases, aliases...)
	return rb
}

// Attributes set the Attributes property for ResolveIndexItemBuilder.
func (rb *ResolveIndexItemBuilder) Attributes(attributes ...string) *ResolveIndexItemBuilder {
	rb.v.Attributes = append(rb.v.Attributes, attributes...)
	return rb
}

// DataStream set the DataStream property for ResolveIndexItemBuilder.
func (rb *ResolveIndexItemBuilder) DataStream(datastream DataStreamName) *ResolveIndexItemBuilder {
	rb.v.DataStream = &datastream
	return rb
}

// Name set the Name property for ResolveIndexItemBuilder.
func (rb *ResolveIndexItemBuilder) Name(name Name) *ResolveIndexItemBuilder {
	rb.v.Name = name
	return rb
}
