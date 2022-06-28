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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/useragentproperty"
)

// UserAgentProcessor type.
type UserAgentProcessor struct {
	Field         Field                                 `json:"field"`
	If            *string                               `json:"if,omitempty"`
	IgnoreFailure *bool                                 `json:"ignore_failure,omitempty"`
	IgnoreMissing bool                                  `json:"ignore_missing"`
	OnFailure     []ProcessorContainer                  `json:"on_failure,omitempty"`
	Options       []useragentproperty.UserAgentProperty `json:"options"`
	RegexFile     string                                `json:"regex_file"`
	Tag           *string                               `json:"tag,omitempty"`
	TargetField   Field                                 `json:"target_field"`
}

// UserAgentProcessorBuilder holds UserAgentProcessor struct and provides a builder API.
type UserAgentProcessorBuilder struct {
	v *UserAgentProcessor
}

// NewUserAgentProcessor provides a builder for the UserAgentProcessor struct.
func NewUserAgentProcessor() *UserAgentProcessorBuilder {
	r := UserAgentProcessorBuilder{
		&UserAgentProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the UserAgentProcessor struct
func (rb *UserAgentProcessorBuilder) Build() UserAgentProcessor {
	return *rb.v
}

// Field set the Field property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) Field(field Field) *UserAgentProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) If_(if_ string) *UserAgentProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) IgnoreFailure(ignorefailure bool) *UserAgentProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) IgnoreMissing(ignoremissing bool) *UserAgentProcessorBuilder {
	rb.v.IgnoreMissing = ignoremissing
	return rb
}

// OnFailure set the OnFailure property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *UserAgentProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Options set the Options property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) Options(options ...useragentproperty.UserAgentProperty) *UserAgentProcessorBuilder {
	rb.v.Options = append(rb.v.Options, options...)
	return rb
}

// RegexFile set the RegexFile property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) RegexFile(regexfile string) *UserAgentProcessorBuilder {
	rb.v.RegexFile = regexfile
	return rb
}

// Tag set the Tag property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) Tag(tag string) *UserAgentProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetField set the TargetField property for UserAgentProcessorBuilder.
func (rb *UserAgentProcessorBuilder) TargetField(targetfield Field) *UserAgentProcessorBuilder {
	rb.v.TargetField = targetfield
	return rb
}
