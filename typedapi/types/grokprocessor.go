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

// GrokProcessor type.
type GrokProcessor struct {
	Field              Field                `json:"field"`
	If                 *string              `json:"if,omitempty"`
	IgnoreFailure      *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing      *bool                `json:"ignore_missing,omitempty"`
	OnFailure          []ProcessorContainer `json:"on_failure,omitempty"`
	PatternDefinitions map[string]string    `json:"pattern_definitions"`
	Patterns           []string             `json:"patterns"`
	Tag                *string              `json:"tag,omitempty"`
	TraceMatch         *bool                `json:"trace_match,omitempty"`
}

// GrokProcessorBuilder holds GrokProcessor struct and provides a builder API.
type GrokProcessorBuilder struct {
	v *GrokProcessor
}

// NewGrokProcessor provides a builder for the GrokProcessor struct.
func NewGrokProcessor() *GrokProcessorBuilder {
	r := GrokProcessorBuilder{
		&GrokProcessor{
			PatternDefinitions: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the GrokProcessor struct
func (rb *GrokProcessorBuilder) Build() GrokProcessor {
	return *rb.v
}

// Field set the Field property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) Field(field Field) *GrokProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) If_(if_ string) *GrokProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) IgnoreFailure(ignorefailure bool) *GrokProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) IgnoreMissing(ignoremissing bool) *GrokProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// OnFailure set the OnFailure property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *GrokProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// PatternDefinitions set the PatternDefinitions property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) PatternDefinitions(value map[string]string) *GrokProcessorBuilder {
	rb.v.PatternDefinitions = value
	return rb
}

// Patterns set the Patterns property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) Patterns(patterns ...string) *GrokProcessorBuilder {
	rb.v.Patterns = append(rb.v.Patterns, patterns...)
	return rb
}

// Tag set the Tag property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) Tag(tag string) *GrokProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TraceMatch set the TraceMatch property for GrokProcessorBuilder.
func (rb *GrokProcessorBuilder) TraceMatch(tracematch bool) *GrokProcessorBuilder {
	rb.v.TraceMatch = &tracematch
	return rb
}
