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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
)

// CompletionProperty type.
type CompletionProperty struct {
	Analyzer                   *string                        `json:"analyzer,omitempty"`
	Contexts                   []SuggestContext               `json:"contexts,omitempty"`
	CopyTo                     *Fields                        `json:"copy_to,omitempty"`
	DocValues                  *bool                          `json:"doc_values,omitempty"`
	Dynamic                    *dynamicmapping.DynamicMapping `json:"dynamic,omitempty"`
	Fields                     map[PropertyName]Property      `json:"fields,omitempty"`
	IgnoreAbove                *int                           `json:"ignore_above,omitempty"`
	LocalMetadata              *Metadata                      `json:"local_metadata,omitempty"`
	MaxInputLength             *int                           `json:"max_input_length,omitempty"`
	Meta                       map[string]string              `json:"meta,omitempty"`
	PreservePositionIncrements *bool                          `json:"preserve_position_increments,omitempty"`
	PreserveSeparators         *bool                          `json:"preserve_separators,omitempty"`
	Properties                 map[PropertyName]Property      `json:"properties,omitempty"`
	SearchAnalyzer             *string                        `json:"search_analyzer,omitempty"`
	Similarity                 *string                        `json:"similarity,omitempty"`
	Store                      *bool                          `json:"store,omitempty"`
	Type                       string                         `json:"type,omitempty"`
}

// CompletionPropertyBuilder holds CompletionProperty struct and provides a builder API.
type CompletionPropertyBuilder struct {
	v *CompletionProperty
}

// NewCompletionProperty provides a builder for the CompletionProperty struct.
func NewCompletionProperty() *CompletionPropertyBuilder {
	r := CompletionPropertyBuilder{
		&CompletionProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "completion"

	return &r
}

// Build finalize the chain and returns the CompletionProperty struct
func (rb *CompletionPropertyBuilder) Build() CompletionProperty {
	return *rb.v
}

// Analyzer set the Analyzer property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) Analyzer(analyzer string) *CompletionPropertyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Contexts set the Contexts property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) Contexts(contexts ...SuggestContext) *CompletionPropertyBuilder {
	rb.v.Contexts = append(rb.v.Contexts, contexts...)
	return rb
}

// CopyTo set the CopyTo property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) CopyTo(copyto Fields) *CompletionPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) DocValues(docvalues bool) *CompletionPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *CompletionPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// Fields set the Fields property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) Fields(value map[PropertyName]Property) *CompletionPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) IgnoreAbove(ignoreabove int) *CompletionPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// LocalMetadata set the LocalMetadata property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) LocalMetadata(localmetadata Metadata) *CompletionPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// MaxInputLength set the MaxInputLength property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) MaxInputLength(maxinputlength int) *CompletionPropertyBuilder {
	rb.v.MaxInputLength = &maxinputlength
	return rb
}

// Meta set the Meta property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) Meta(value map[string]string) *CompletionPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// PreservePositionIncrements set the PreservePositionIncrements property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) PreservePositionIncrements(preservepositionincrements bool) *CompletionPropertyBuilder {
	rb.v.PreservePositionIncrements = &preservepositionincrements
	return rb
}

// PreserveSeparators set the PreserveSeparators property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) PreserveSeparators(preserveseparators bool) *CompletionPropertyBuilder {
	rb.v.PreserveSeparators = &preserveseparators
	return rb
}

// Properties set the Properties property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) Properties(value map[PropertyName]Property) *CompletionPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// SearchAnalyzer set the SearchAnalyzer property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) SearchAnalyzer(searchanalyzer string) *CompletionPropertyBuilder {
	rb.v.SearchAnalyzer = &searchanalyzer
	return rb
}

// Similarity set the Similarity property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) Similarity(similarity string) *CompletionPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for CompletionPropertyBuilder.
func (rb *CompletionPropertyBuilder) Store(store bool) *CompletionPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// Type set the Type property for CompletionPropertyBuilder.
