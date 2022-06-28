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

// CsvProcessor type.
type CsvProcessor struct {
	Description   *string              `json:"description,omitempty"`
	EmptyValue    interface{}          `json:"empty_value,omitempty"`
	Field         Field                `json:"field"`
	If            *string              `json:"if,omitempty"`
	IgnoreFailure *bool                `json:"ignore_failure,omitempty"`
	IgnoreMissing *bool                `json:"ignore_missing,omitempty"`
	OnFailure     []ProcessorContainer `json:"on_failure,omitempty"`
	Quote         *string              `json:"quote,omitempty"`
	Separator     *string              `json:"separator,omitempty"`
	Tag           *string              `json:"tag,omitempty"`
	TargetFields  Fields               `json:"target_fields"`
	Trim          bool                 `json:"trim"`
}

// CsvProcessorBuilder holds CsvProcessor struct and provides a builder API.
type CsvProcessorBuilder struct {
	v *CsvProcessor
}

// NewCsvProcessor provides a builder for the CsvProcessor struct.
func NewCsvProcessor() *CsvProcessorBuilder {
	r := CsvProcessorBuilder{
		&CsvProcessor{},
	}

	return &r
}

// Build finalize the chain and returns the CsvProcessor struct
func (rb *CsvProcessorBuilder) Build() CsvProcessor {
	return *rb.v
}

// Description set the Description property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) Description(description string) *CsvProcessorBuilder {
	rb.v.Description = &description
	return rb
}

// EmptyValue set the EmptyValue property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) EmptyValue(emptyvalue interface{}) *CsvProcessorBuilder {
	rb.v.EmptyValue = emptyvalue
	return rb
}

// Field set the Field property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) Field(field Field) *CsvProcessorBuilder {
	rb.v.Field = field
	return rb
}

// If set the If property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) If_(if_ string) *CsvProcessorBuilder {
	rb.v.If = &if_
	return rb
}

// IgnoreFailure set the IgnoreFailure property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) IgnoreFailure(ignorefailure bool) *CsvProcessorBuilder {
	rb.v.IgnoreFailure = &ignorefailure
	return rb
}

// IgnoreMissing set the IgnoreMissing property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) IgnoreMissing(ignoremissing bool) *CsvProcessorBuilder {
	rb.v.IgnoreMissing = &ignoremissing
	return rb
}

// OnFailure set the OnFailure property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) OnFailure(on_failure ...ProcessorContainer) *CsvProcessorBuilder {
	rb.v.OnFailure = append(rb.v.OnFailure, on_failure...)
	return rb
}

// Quote set the Quote property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) Quote(quote string) *CsvProcessorBuilder {
	rb.v.Quote = &quote
	return rb
}

// Separator set the Separator property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) Separator(separator string) *CsvProcessorBuilder {
	rb.v.Separator = &separator
	return rb
}

// Tag set the Tag property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) Tag(tag string) *CsvProcessorBuilder {
	rb.v.Tag = &tag
	return rb
}

// TargetFields set the TargetFields property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) TargetFields(targetfields Fields) *CsvProcessorBuilder {
	rb.v.TargetFields = targetfields
	return rb
}

// Trim set the Trim property for CsvProcessorBuilder.
func (rb *CsvProcessorBuilder) Trim(trim bool) *CsvProcessorBuilder {
	rb.v.Trim = trim
	return rb
}
