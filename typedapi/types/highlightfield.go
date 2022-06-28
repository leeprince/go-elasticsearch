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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/boundaryscanner"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterfragmenter"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlighterorder"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertagsschema"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/highlightertype"
)

// HighlightField type.
type HighlightField struct {
	BoundaryChars         *string                                      `json:"boundary_chars,omitempty"`
	BoundaryMaxScan       *int                                         `json:"boundary_max_scan,omitempty"`
	BoundaryScanner       *boundaryscanner.BoundaryScanner             `json:"boundary_scanner,omitempty"`
	BoundaryScannerLocale *string                                      `json:"boundary_scanner_locale,omitempty"`
	ForceSource           *bool                                        `json:"force_source,omitempty"`
	FragmentOffset        *int                                         `json:"fragment_offset,omitempty"`
	FragmentSize          *int                                         `json:"fragment_size,omitempty"`
	Fragmenter            *highlighterfragmenter.HighlighterFragmenter `json:"fragmenter,omitempty"`
	HighlightFilter       *bool                                        `json:"highlight_filter,omitempty"`
	HighlightQuery        *QueryContainer                              `json:"highlight_query,omitempty"`
	MatchedFields         *Fields                                      `json:"matched_fields,omitempty"`
	MaxAnalyzedOffset     *int                                         `json:"max_analyzed_offset,omitempty"`
	MaxFragmentLength     *int                                         `json:"max_fragment_length,omitempty"`
	NoMatchSize           *int                                         `json:"no_match_size,omitempty"`
	NumberOfFragments     *int                                         `json:"number_of_fragments,omitempty"`
	Options               map[string]interface{}                       `json:"options,omitempty"`
	Order                 *highlighterorder.HighlighterOrder           `json:"order,omitempty"`
	PhraseLimit           *int                                         `json:"phrase_limit,omitempty"`
	PostTags              []string                                     `json:"post_tags,omitempty"`
	PreTags               []string                                     `json:"pre_tags,omitempty"`
	RequireFieldMatch     *bool                                        `json:"require_field_match,omitempty"`
	TagsSchema            *highlightertagsschema.HighlighterTagsSchema `json:"tags_schema,omitempty"`
	Type                  *highlightertype.HighlighterType             `json:"type,omitempty"`
}

// HighlightFieldBuilder holds HighlightField struct and provides a builder API.
type HighlightFieldBuilder struct {
	v *HighlightField
}

// NewHighlightField provides a builder for the HighlightField struct.
func NewHighlightField() *HighlightFieldBuilder {
	r := HighlightFieldBuilder{
		&HighlightField{
			Options: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the HighlightField struct
func (rb *HighlightFieldBuilder) Build() HighlightField {
	return *rb.v
}

// BoundaryChars set the BoundaryChars property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) BoundaryChars(boundarychars string) *HighlightFieldBuilder {
	rb.v.BoundaryChars = &boundarychars
	return rb
}

// BoundaryMaxScan set the BoundaryMaxScan property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) BoundaryMaxScan(boundarymaxscan int) *HighlightFieldBuilder {
	rb.v.BoundaryMaxScan = &boundarymaxscan
	return rb
}

// BoundaryScanner set the BoundaryScanner property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) BoundaryScanner(boundaryscanner boundaryscanner.BoundaryScanner) *HighlightFieldBuilder {
	rb.v.BoundaryScanner = &boundaryscanner
	return rb
}

// BoundaryScannerLocale set the BoundaryScannerLocale property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) BoundaryScannerLocale(boundaryscannerlocale string) *HighlightFieldBuilder {
	rb.v.BoundaryScannerLocale = &boundaryscannerlocale
	return rb
}

// ForceSource set the ForceSource property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) ForceSource(forcesource bool) *HighlightFieldBuilder {
	rb.v.ForceSource = &forcesource
	return rb
}

// FragmentOffset set the FragmentOffset property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) FragmentOffset(fragmentoffset int) *HighlightFieldBuilder {
	rb.v.FragmentOffset = &fragmentoffset
	return rb
}

// FragmentSize set the FragmentSize property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) FragmentSize(fragmentsize int) *HighlightFieldBuilder {
	rb.v.FragmentSize = &fragmentsize
	return rb
}

// Fragmenter set the Fragmenter property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) Fragmenter(fragmenter highlighterfragmenter.HighlighterFragmenter) *HighlightFieldBuilder {
	rb.v.Fragmenter = &fragmenter
	return rb
}

// HighlightFilter set the HighlightFilter property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) HighlightFilter(highlightfilter bool) *HighlightFieldBuilder {
	rb.v.HighlightFilter = &highlightfilter
	return rb
}

// HighlightQuery set the HighlightQuery property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) HighlightQuery(highlightquery QueryContainer) *HighlightFieldBuilder {
	rb.v.HighlightQuery = &highlightquery
	return rb
}

// MatchedFields set the MatchedFields property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) MatchedFields(matchedfields Fields) *HighlightFieldBuilder {
	rb.v.MatchedFields = &matchedfields
	return rb
}

// MaxAnalyzedOffset set the MaxAnalyzedOffset property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) MaxAnalyzedOffset(maxanalyzedoffset int) *HighlightFieldBuilder {
	rb.v.MaxAnalyzedOffset = &maxanalyzedoffset
	return rb
}

// MaxFragmentLength set the MaxFragmentLength property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) MaxFragmentLength(maxfragmentlength int) *HighlightFieldBuilder {
	rb.v.MaxFragmentLength = &maxfragmentlength
	return rb
}

// NoMatchSize set the NoMatchSize property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) NoMatchSize(nomatchsize int) *HighlightFieldBuilder {
	rb.v.NoMatchSize = &nomatchsize
	return rb
}

// NumberOfFragments set the NumberOfFragments property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) NumberOfFragments(numberoffragments int) *HighlightFieldBuilder {
	rb.v.NumberOfFragments = &numberoffragments
	return rb
}

// Options set the Options property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) Options(value map[string]interface{}) *HighlightFieldBuilder {
	rb.v.Options = value
	return rb
}

// Order set the Order property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) Order(order highlighterorder.HighlighterOrder) *HighlightFieldBuilder {
	rb.v.Order = &order
	return rb
}

// PhraseLimit set the PhraseLimit property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) PhraseLimit(phraselimit int) *HighlightFieldBuilder {
	rb.v.PhraseLimit = &phraselimit
	return rb
}

// PostTags set the PostTags property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) PostTags(post_tags ...string) *HighlightFieldBuilder {
	rb.v.PostTags = append(rb.v.PostTags, post_tags...)
	return rb
}

// PreTags set the PreTags property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) PreTags(pre_tags ...string) *HighlightFieldBuilder {
	rb.v.PreTags = append(rb.v.PreTags, pre_tags...)
	return rb
}

// RequireFieldMatch set the RequireFieldMatch property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) RequireFieldMatch(requirefieldmatch bool) *HighlightFieldBuilder {
	rb.v.RequireFieldMatch = &requirefieldmatch
	return rb
}

// TagsSchema set the TagsSchema property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) TagsSchema(tagsschema highlightertagsschema.HighlighterTagsSchema) *HighlightFieldBuilder {
	rb.v.TagsSchema = &tagsschema
	return rb
}

// Type set the Type property for HighlightFieldBuilder.
func (rb *HighlightFieldBuilder) Type_(type_ highlightertype.HighlighterType) *HighlightFieldBuilder {
	rb.v.Type = &type_
	return rb
}
