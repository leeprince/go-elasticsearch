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

// IndexState type.
type IndexState struct {
	Aliases    map[IndexName]Alias `json:"aliases,omitempty"`
	DataStream *DataStreamName     `json:"data_stream,omitempty"`
	// Defaults Default settings, included when the request's `include_default` is `true`.
	Defaults *IndexSettings `json:"defaults,omitempty"`
	Mappings *TypeMapping   `json:"mappings,omitempty"`
	Settings *IndexSettings `json:"settings,omitempty"`
}

// IndexStateBuilder holds IndexState struct and provides a builder API.
type IndexStateBuilder struct {
	v *IndexState
}

// NewIndexState provides a builder for the IndexState struct.
func NewIndexState() *IndexStateBuilder {
	r := IndexStateBuilder{
		&IndexState{
			Aliases: make(map[IndexName]Alias, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IndexState struct
func (rb *IndexStateBuilder) Build() IndexState {
	return *rb.v
}

// Aliases set the Aliases property for IndexStateBuilder.
func (rb *IndexStateBuilder) Aliases(value map[IndexName]Alias) *IndexStateBuilder {
	rb.v.Aliases = value
	return rb
}

// DataStream set the DataStream property for IndexStateBuilder.
func (rb *IndexStateBuilder) DataStream(datastream DataStreamName) *IndexStateBuilder {
	rb.v.DataStream = &datastream
	return rb
}

// Defaults Default settings, included when the request's `include_default` is `true`.
func (rb *IndexStateBuilder) Defaults(defaults IndexSettings) *IndexStateBuilder {
	rb.v.Defaults = &defaults
	return rb
}

// Mappings set the Mappings property for IndexStateBuilder.
func (rb *IndexStateBuilder) Mappings(mappings TypeMapping) *IndexStateBuilder {
	rb.v.Mappings = &mappings
	return rb
}

// Settings set the Settings property for IndexStateBuilder.
func (rb *IndexStateBuilder) Settings(settings IndexSettings) *IndexStateBuilder {
	rb.v.Settings = &settings
	return rb
}
