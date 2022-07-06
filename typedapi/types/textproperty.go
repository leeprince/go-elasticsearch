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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexoptions"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termvectoroption"
)

// TextProperty type.
type TextProperty struct {
	Analyzer                 *string                            `json:"analyzer,omitempty"`
	Boost                    *float64                           `json:"boost,omitempty"`
	CopyTo                   *Fields                            `json:"copy_to,omitempty"`
	Dynamic                  *dynamicmapping.DynamicMapping     `json:"dynamic,omitempty"`
	EagerGlobalOrdinals      *bool                              `json:"eager_global_ordinals,omitempty"`
	Fielddata                *bool                              `json:"fielddata,omitempty"`
	FielddataFrequencyFilter *FielddataFrequencyFilter          `json:"fielddata_frequency_filter,omitempty"`
	Fields                   map[PropertyName]Property          `json:"fields,omitempty"`
	IgnoreAbove              *int                               `json:"ignore_above,omitempty"`
	Index                    *bool                              `json:"index,omitempty"`
	IndexOptions             *indexoptions.IndexOptions         `json:"index_options,omitempty"`
	IndexPhrases             *bool                              `json:"index_phrases,omitempty"`
	IndexPrefixes            *TextIndexPrefixes                 `json:"index_prefixes,omitempty"`
	LocalMetadata            *Metadata                          `json:"local_metadata,omitempty"`
	Meta                     map[string]string                  `json:"meta,omitempty"`
	Norms                    *bool                              `json:"norms,omitempty"`
	PositionIncrementGap     *int                               `json:"position_increment_gap,omitempty"`
	Properties               map[PropertyName]Property          `json:"properties,omitempty"`
	SearchAnalyzer           *string                            `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer      *string                            `json:"search_quote_analyzer,omitempty"`
	Similarity               *string                            `json:"similarity,omitempty"`
	Store                    *bool                              `json:"store,omitempty"`
	TermVector               *termvectoroption.TermVectorOption `json:"term_vector,omitempty"`
	Type                     string                             `json:"type,omitempty"`
}

// TextPropertyBuilder holds TextProperty struct and provides a builder API.
type TextPropertyBuilder struct {
	v *TextProperty
}

// NewTextProperty provides a builder for the TextProperty struct.
func NewTextProperty() *TextPropertyBuilder {
	r := TextPropertyBuilder{
		&TextProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "text"

	return &r
}

// Build finalize the chain and returns the TextProperty struct
func (rb *TextPropertyBuilder) Build() TextProperty {
	return *rb.v
}

// Analyzer set the Analyzer property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Analyzer(analyzer string) *TextPropertyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Boost set the Boost property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Boost(boost float64) *TextPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// CopyTo set the CopyTo property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) CopyTo(copyto Fields) *TextPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// Dynamic set the Dynamic property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *TextPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// EagerGlobalOrdinals set the EagerGlobalOrdinals property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) EagerGlobalOrdinals(eagerglobalordinals bool) *TextPropertyBuilder {
	rb.v.EagerGlobalOrdinals = &eagerglobalordinals
	return rb
}

// Fielddata set the Fielddata property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Fielddata(fielddata bool) *TextPropertyBuilder {
	rb.v.Fielddata = &fielddata
	return rb
}

// FielddataFrequencyFilter set the FielddataFrequencyFilter property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) FielddataFrequencyFilter(fielddatafrequencyfilter FielddataFrequencyFilter) *TextPropertyBuilder {
	rb.v.FielddataFrequencyFilter = &fielddatafrequencyfilter
	return rb
}

// Fields set the Fields property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Fields(value map[PropertyName]Property) *TextPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// IgnoreAbove set the IgnoreAbove property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) IgnoreAbove(ignoreabove int) *TextPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// Index set the Index property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Index(index bool) *TextPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// IndexOptions set the IndexOptions property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) IndexOptions(indexoptions indexoptions.IndexOptions) *TextPropertyBuilder {
	rb.v.IndexOptions = &indexoptions
	return rb
}

// IndexPhrases set the IndexPhrases property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) IndexPhrases(indexphrases bool) *TextPropertyBuilder {
	rb.v.IndexPhrases = &indexphrases
	return rb
}

// IndexPrefixes set the IndexPrefixes property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) IndexPrefixes(indexprefixes TextIndexPrefixes) *TextPropertyBuilder {
	rb.v.IndexPrefixes = &indexprefixes
	return rb
}

// LocalMetadata set the LocalMetadata property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) LocalMetadata(localmetadata Metadata) *TextPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Meta set the Meta property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Meta(value map[string]string) *TextPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Norms set the Norms property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Norms(norms bool) *TextPropertyBuilder {
	rb.v.Norms = &norms
	return rb
}

// PositionIncrementGap set the PositionIncrementGap property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) PositionIncrementGap(positionincrementgap int) *TextPropertyBuilder {
	rb.v.PositionIncrementGap = &positionincrementgap
	return rb
}

// Properties set the Properties property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Properties(value map[PropertyName]Property) *TextPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// SearchAnalyzer set the SearchAnalyzer property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) SearchAnalyzer(searchanalyzer string) *TextPropertyBuilder {
	rb.v.SearchAnalyzer = &searchanalyzer
	return rb
}

// SearchQuoteAnalyzer set the SearchQuoteAnalyzer property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) SearchQuoteAnalyzer(searchquoteanalyzer string) *TextPropertyBuilder {
	rb.v.SearchQuoteAnalyzer = &searchquoteanalyzer
	return rb
}

// Similarity set the Similarity property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Similarity(similarity string) *TextPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) Store(store bool) *TextPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TermVector set the TermVector property for TextPropertyBuilder.
func (rb *TextPropertyBuilder) TermVector(termvector termvectoroption.TermVectorOption) *TextPropertyBuilder {
	rb.v.TermVector = &termvector
	return rb
}

// Type set the Type property for TextPropertyBuilder.
