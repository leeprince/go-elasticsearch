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

// DropProcessor type.
type DropProcessor struct {
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// DropProcessorBuilder holds DropProcessor struct and provides a builder API.
type DropProcessorBuilder struct {
	v *DropProcessor
}

// NewDropProcessor provides a builder for the DropProcessor struct.
func NewDropProcessor() *DropProcessorBuilder {
	r := DropProcessorBuilder{
		&DropProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the DropProcessor struct
func (rb *DropProcessorBuilder) Build() DropProcessor {
	return *rb.v
}

// If set the If property for DropProcessorBuilder.
func (rb *DropProcessorBuilder) If_(if_ string) *DropProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for DropProcessorBuilder.
func (rb *DropProcessorBuilder) IgnoreFailure(ignorefailure bool) *DropProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// OnFailure set the OnFailure property for DropProcessorBuilder.
func (rb *DropProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *DropProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for DropProcessorBuilder.
func (rb *DropProcessorBuilder) Tag(tag string) *DropProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}
