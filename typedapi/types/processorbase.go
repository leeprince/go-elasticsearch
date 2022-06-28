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

// ProcessorBase type.
type ProcessorBase struct {
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
}

// ProcessorBaseBuilder holds ProcessorBase struct and provides a builder API.
type ProcessorBaseBuilder struct {
	v *ProcessorBase
}

// NewProcessorBase provides a builder for the ProcessorBase struct.
func NewProcessorBase() *ProcessorBaseBuilder {
	r := ProcessorBaseBuilder{
		&ProcessorBase{},
	}

	return &r
}

// Build finalize the chain and returns the ProcessorBase struct
func (rb *ProcessorBaseBuilder) Build() ProcessorBase {
	return *rb.v
}

// If set the If property for ProcessorBaseBuilder.
func (rb *ProcessorBaseBuilder) If_(if_ string) *ProcessorBaseBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for ProcessorBaseBuilder.
func (rb *ProcessorBaseBuilder) IgnoreFailure(ignorefailure bool) *ProcessorBaseBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// OnFailure set the OnFailure property for ProcessorBaseBuilder.
func (rb *ProcessorBaseBuilder) OnFailure(on_failure ...ProcessorContainer) *ProcessorBaseBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Tag set the Tag property for ProcessorBaseBuilder.
func (rb *ProcessorBaseBuilder) Tag(tag string) *ProcessorBaseBuilder {
	rb.v.Tag = &tag
	return rb
}
