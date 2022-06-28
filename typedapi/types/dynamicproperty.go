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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/onscripterror"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termvectoroption"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeseriesmetrictype"
)

// DynamicProperty type.
type DynamicProperty struct {
	Analyzer             *string                                    `json:"analyzer,omitempty"`
	Boost                *float64                                   `json:"boost,omitempty"`
	Coerce               *bool                                      `json:"coerce,omitempty"`
	CopyTo               *Fields                                    `json:"copy_to,omitempty"`
	DocValues            *bool                                      `json:"doc_values,omitempty"`
	Dynamic              *dynamicmapping.DynamicMapping             `json:"dynamic,omitempty"`
	EagerGlobalOrdinals  *bool                                      `json:"eager_global_ordinals,omitempty"`
	Enabled              *bool                                      `json:"enabled,omitempty"`
	Fields               map[PropertyName]Property                  `json:"fields,omitempty"`
	Format               *string                                    `json:"format,omitempty"`
	IgnoreAbove          *int                                       `json:"ignore_above,omitempty"`
	IgnoreMalformed      *bool                                      `json:"ignore_malformed,omitempty"`
	Index                *bool                                      `json:"index,omitempty"`
	IndexOptions         *indexoptions.IndexOptions                 `json:"index_options,omitempty"`
	IndexPhrases         *bool                                      `json:"index_phrases,omitempty"`
	IndexPrefixes        *TextIndexPrefixes                         `json:"index_prefixes,omitempty"`
	LocalMetadata        *Metadata                                  `json:"local_metadata,omitempty"`
	Locale               *string                                    `json:"locale,omitempty"`
	Meta                 map[string]string                          `json:"meta,omitempty"`
	Norms                *bool                                      `json:"norms,omitempty"`
	NullValue            *FieldValue                                `json:"null_value,omitempty"`
	OnScriptError        *onscripterror.OnScriptError               `json:"on_script_error,omitempty"`
	PositionIncrementGap *int                                       `json:"position_increment_gap,omitempty"`
	PrecisionStep        *int                                       `json:"precision_step,omitempty"`
	Properties           map[PropertyName]Property                  `json:"properties,omitempty"`
	Script               *Script                                    `json:"script,omitempty"`
	SearchAnalyzer       *string                                    `json:"search_analyzer,omitempty"`
	SearchQuoteAnalyzer  *string                                    `json:"search_quote_analyzer,omitempty"`
	Similarity           *string                                    `json:"similarity,omitempty"`
	Store                *bool                                      `json:"store,omitempty"`
	TermVector           *termvectoroption.TermVectorOption         `json:"term_vector,omitempty"`
	TimeSeriesMetric     *timeseriesmetrictype.TimeSeriesMetricType `json:"time_series_metric,omitempty"`
	Type                 string                                     `json:"type,omitempty"`
}

// DynamicPropertyBuilder holds DynamicProperty struct and provides a builder API.
type DynamicPropertyBuilder struct {
	v *DynamicProperty
}

// NewDynamicProperty provides a builder for the DynamicProperty struct.
func NewDynamicProperty() *DynamicPropertyBuilder {
	r := DynamicPropertyBuilder{
		&DynamicProperty{
			Fields:     make(map[PropertyName]Property, 0),
			Meta:       make(map[string]string, 0),
			Properties: make(map[PropertyName]Property, 0),
		},
	}

	r.v.Type = "{dynamic_property}"

	return &r
}

// Build finalize the chain and returns the DynamicProperty struct
func (rb *DynamicPropertyBuilder) Build() DynamicProperty {
	return *rb.v
}

// Analyzer set the Analyzer property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Analyzer(analyzer string) *DynamicPropertyBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

// Boost set the Boost property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Boost(boost float64) *DynamicPropertyBuilder {
	rb.v.Boost = &boost
	return rb
}

// Coerce set the Coerce property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Coerce(coerce bool) *DynamicPropertyBuilder {
	rb.v.Coerce = &coerce
	return rb
}

// CopyTo set the CopyTo property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) CopyTo(copyto Fields) *DynamicPropertyBuilder {
	rb.v.CopyTo = &copyto
	return rb
}

// DocValues set the DocValues property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) DocValues(docvalues bool) *DynamicPropertyBuilder {
	rb.v.DocValues = &docvalues
	return rb
}

// Dynamic set the Dynamic property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Dynamic(dynamic dynamicmapping.DynamicMapping) *DynamicPropertyBuilder {
	rb.v.Dynamic = &dynamic
	return rb
}

// EagerGlobalOrdinals set the EagerGlobalOrdinals property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) EagerGlobalOrdinals(eagerglobalordinals bool) *DynamicPropertyBuilder {
	rb.v.EagerGlobalOrdinals = &eagerglobalordinals
	return rb
}

// Enabled set the Enabled property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Enabled(enabled bool) *DynamicPropertyBuilder {
	rb.v.Enabled = &enabled
	return rb
}

// Fields set the Fields property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Fields(value map[PropertyName]Property) *DynamicPropertyBuilder {
	rb.v.Fields = value
	return rb
}

// Format set the Format property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Format(format string) *DynamicPropertyBuilder {
	rb.v.Format = &format
	return rb
}

// IgnoreAbove set the IgnoreAbove property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) IgnoreAbove(ignoreabove int) *DynamicPropertyBuilder {
	rb.v.IgnoreAbove = &ignoreabove
	return rb
}

// IgnoreMalformed set the IgnoreMalformed property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) IgnoreMalformed(ignoremalformed bool) *DynamicPropertyBuilder {
	rb.v.IgnoreMalformed = &ignoremalformed
	return rb
}

// Index set the Index property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Index(index bool) *DynamicPropertyBuilder {
	rb.v.Index = &index
	return rb
}

// IndexOptions set the IndexOptions property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) IndexOptions(indexoptions indexoptions.IndexOptions) *DynamicPropertyBuilder {
	rb.v.IndexOptions = &indexoptions
	return rb
}

// IndexPhrases set the IndexPhrases property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) IndexPhrases(indexphrases bool) *DynamicPropertyBuilder {
	rb.v.IndexPhrases = &indexphrases
	return rb
}

// IndexPrefixes set the IndexPrefixes property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) IndexPrefixes(indexprefixes TextIndexPrefixes) *DynamicPropertyBuilder {
	rb.v.IndexPrefixes = &indexprefixes
	return rb
}

// LocalMetadata set the LocalMetadata property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) LocalMetadata(localmetadata Metadata) *DynamicPropertyBuilder {
	rb.v.LocalMetadata = &localmetadata
	return rb
}

// Locale set the Locale property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Locale(locale string) *DynamicPropertyBuilder {
	rb.v.Locale = &locale
	return rb
}

// Meta set the Meta property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Meta(value map[string]string) *DynamicPropertyBuilder {
	rb.v.Meta = value
	return rb
}

// Norms set the Norms property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Norms(norms bool) *DynamicPropertyBuilder {
	rb.v.Norms = &norms
	return rb
}

// NullValue set the NullValue property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) NullValue(nullvalue FieldValue) *DynamicPropertyBuilder {
	rb.v.NullValue = &nullvalue
	return rb
}

// OnScriptError set the OnScriptError property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) OnScriptError(onscripterror onscripterror.OnScriptError) *DynamicPropertyBuilder {
	rb.v.OnScriptError = &onscripterror
	return rb
}

// PositionIncrementGap set the PositionIncrementGap property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) PositionIncrementGap(positionincrementgap int) *DynamicPropertyBuilder {
	rb.v.PositionIncrementGap = &positionincrementgap
	return rb
}

// PrecisionStep set the PrecisionStep property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) PrecisionStep(precisionstep int) *DynamicPropertyBuilder {
	rb.v.PrecisionStep = &precisionstep
	return rb
}

// Properties set the Properties property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Properties(value map[PropertyName]Property) *DynamicPropertyBuilder {
	rb.v.Properties = value
	return rb
}

// Script set the Script property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Script(script Script) *DynamicPropertyBuilder {
	rb.v.Script = &script
	return rb
}

// SearchAnalyzer set the SearchAnalyzer property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) SearchAnalyzer(searchanalyzer string) *DynamicPropertyBuilder {
	rb.v.SearchAnalyzer = &searchanalyzer
	return rb
}

// SearchQuoteAnalyzer set the SearchQuoteAnalyzer property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) SearchQuoteAnalyzer(searchquoteanalyzer string) *DynamicPropertyBuilder {
	rb.v.SearchQuoteAnalyzer = &searchquoteanalyzer
	return rb
}

// Similarity set the Similarity property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Similarity(similarity string) *DynamicPropertyBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// Store set the Store property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) Store(store bool) *DynamicPropertyBuilder {
	rb.v.Store = &store
	return rb
}

// TermVector set the TermVector property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) TermVector(termvector termvectoroption.TermVectorOption) *DynamicPropertyBuilder {
	rb.v.TermVector = &termvector
	return rb
}

// TimeSeriesMetric set the TimeSeriesMetric property for DynamicPropertyBuilder.
func (rb *DynamicPropertyBuilder) TimeSeriesMetric(timeseriesmetric timeseriesmetrictype.TimeSeriesMetricType) *DynamicPropertyBuilder {
	rb.v.TimeSeriesMetric = &timeseriesmetric
	return rb
}

// Type set the Type property for DynamicPropertyBuilder.
