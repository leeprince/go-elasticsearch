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

// QueryContainer type.
type QueryContainer struct {
	Bool              *BoolQuery                       `json:"bool,omitempty"`
	Boosting          *BoostingQuery                   `json:"boosting,omitempty"`
	CombinedFields    *CombinedFieldsQuery             `json:"combined_fields,omitempty"`
	Common            map[Field]CommonTermsQuery       `json:"common,omitempty"`
	ConstantScore     *ConstantScoreQuery              `json:"constant_score,omitempty"`
	DisMax            *DisMaxQuery                     `json:"dis_max,omitempty"`
	DistanceFeature   *DistanceFeatureQuery            `json:"distance_feature,omitempty"`
	Exists            *ExistsQuery                     `json:"exists,omitempty"`
	FieldMaskingSpan  *SpanFieldMaskingQuery           `json:"field_masking_span,omitempty"`
	FunctionScore     *FunctionScoreQuery              `json:"function_score,omitempty"`
	Fuzzy             map[Field]FuzzyQuery             `json:"fuzzy,omitempty"`
	GeoBoundingBox    *GeoBoundingBoxQuery             `json:"geo_bounding_box,omitempty"`
	GeoDistance       *GeoDistanceQuery                `json:"geo_distance,omitempty"`
	GeoPolygon        *GeoPolygonQuery                 `json:"geo_polygon,omitempty"`
	GeoShape          *GeoShapeQuery                   `json:"geo_shape,omitempty"`
	HasChild          *HasChildQuery                   `json:"has_child,omitempty"`
	HasParent         *HasParentQuery                  `json:"has_parent,omitempty"`
	Ids               *IdsQuery                        `json:"ids,omitempty"`
	Intervals         map[Field]IntervalsQuery         `json:"intervals,omitempty"`
	Knn               *KnnQuery                        `json:"knn,omitempty"`
	Match             map[Field]MatchQuery             `json:"match,omitempty"`
	MatchAll          *MatchAllQuery                   `json:"match_all,omitempty"`
	MatchBoolPrefix   map[Field]MatchBoolPrefixQuery   `json:"match_bool_prefix,omitempty"`
	MatchNone         *MatchNoneQuery                  `json:"match_none,omitempty"`
	MatchPhrase       map[Field]MatchPhraseQuery       `json:"match_phrase,omitempty"`
	MatchPhrasePrefix map[Field]MatchPhrasePrefixQuery `json:"match_phrase_prefix,omitempty"`
	MoreLikeThis      *MoreLikeThisQuery               `json:"more_like_this,omitempty"`
	MultiMatch        *MultiMatchQuery                 `json:"multi_match,omitempty"`
	Nested            *NestedQuery                     `json:"nested,omitempty"`
	ParentId          *ParentIdQuery                   `json:"parent_id,omitempty"`
	Percolate         *PercolateQuery                  `json:"percolate,omitempty"`
	Pinned            *PinnedQuery                     `json:"pinned,omitempty"`
	Prefix            map[Field]PrefixQuery            `json:"prefix,omitempty"`
	QueryString       *QueryStringQuery                `json:"query_string,omitempty"`
	Range             map[Field]RangeQuery             `json:"range,omitempty"`
	RankFeature       *RankFeatureQuery                `json:"rank_feature,omitempty"`
	Regexp            map[Field]RegexpQuery            `json:"regexp,omitempty"`
	Script            *ScriptQuery                     `json:"script,omitempty"`
	ScriptScore       *ScriptScoreQuery                `json:"script_score,omitempty"`
	Shape             *ShapeQuery                      `json:"shape,omitempty"`
	SimpleQueryString *SimpleQueryStringQuery          `json:"simple_query_string,omitempty"`
	SpanContaining    *SpanContainingQuery             `json:"span_containing,omitempty"`
	SpanFirst         *SpanFirstQuery                  `json:"span_first,omitempty"`
	SpanMulti         *SpanMultiTermQuery              `json:"span_multi,omitempty"`
	SpanNear          *SpanNearQuery                   `json:"span_near,omitempty"`
	SpanNot           *SpanNotQuery                    `json:"span_not,omitempty"`
	SpanOr            *SpanOrQuery                     `json:"span_or,omitempty"`
	SpanTerm          map[Field]SpanTermQuery          `json:"span_term,omitempty"`
	SpanWithin        *SpanWithinQuery                 `json:"span_within,omitempty"`
	Term              map[Field]TermQuery              `json:"term,omitempty"`
	Terms             *TermsQuery                      `json:"terms,omitempty"`
	TermsSet          map[Field]TermsSetQuery          `json:"terms_set,omitempty"`
	Type              *TypeQuery                       `json:"type,omitempty"`
	Wildcard          map[Field]WildcardQuery          `json:"wildcard,omitempty"`
	Wrapper           *WrapperQuery                    `json:"wrapper,omitempty"`
}

// QueryContainerBuilder holds QueryContainer struct and provides a builder API.
type QueryContainerBuilder struct {
	v *QueryContainer
}

// NewQueryContainer provides a builder for the QueryContainer struct.
func NewQueryContainer() *QueryContainerBuilder {
	r := QueryContainerBuilder{
		&QueryContainer{
			Common:            make(map[Field]CommonTermsQuery, 0),
			Fuzzy:             make(map[Field]FuzzyQuery, 0),
			Intervals:         make(map[Field]IntervalsQuery, 0),
			Match:             make(map[Field]MatchQuery, 0),
			MatchBoolPrefix:   make(map[Field]MatchBoolPrefixQuery, 0),
			MatchPhrase:       make(map[Field]MatchPhraseQuery, 0),
			MatchPhrasePrefix: make(map[Field]MatchPhrasePrefixQuery, 0),
			Prefix:            make(map[Field]PrefixQuery, 0),
			Range:             make(map[Field]RangeQuery, 0),
			Regexp:            make(map[Field]RegexpQuery, 0),
			SpanTerm:          make(map[Field]SpanTermQuery, 0),
			Term:              make(map[Field]TermQuery, 0),
			TermsSet:          make(map[Field]TermsSetQuery, 0),
			Wildcard:          make(map[Field]WildcardQuery, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the QueryContainer struct
func (rb *QueryContainerBuilder) Build() QueryContainer {
	return *rb.v
}

// Bool set the Bool property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Bool(bool BoolQuery) *QueryContainerBuilder {
	rb.v.Bool = &bool
	return rb
}

// Boosting set the Boosting property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Boosting(boosting BoostingQuery) *QueryContainerBuilder {
	rb.v.Boosting = &boosting
	return rb
}

// CombinedFields set the CombinedFields property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) CombinedFields(combinedfields CombinedFieldsQuery) *QueryContainerBuilder {
	rb.v.CombinedFields = &combinedfields
	return rb
}

// Common set the Common property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Common(value map[Field]CommonTermsQuery) *QueryContainerBuilder {
	rb.v.Common = value
	return rb
}

// ConstantScore set the ConstantScore property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) ConstantScore(constantscore ConstantScoreQuery) *QueryContainerBuilder {
	rb.v.ConstantScore = &constantscore
	return rb
}

// DisMax set the DisMax property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) DisMax(dismax DisMaxQuery) *QueryContainerBuilder {
	rb.v.DisMax = &dismax
	return rb
}

// DistanceFeature set the DistanceFeature property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) DistanceFeature(distancefeature DistanceFeatureQuery) *QueryContainerBuilder {
	rb.v.DistanceFeature = &distancefeature
	return rb
}

// Exists set the Exists property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Exists(exists ExistsQuery) *QueryContainerBuilder {
	rb.v.Exists = &exists
	return rb
}

// FieldMaskingSpan set the FieldMaskingSpan property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) FieldMaskingSpan(fieldmaskingspan SpanFieldMaskingQuery) *QueryContainerBuilder {
	rb.v.FieldMaskingSpan = &fieldmaskingspan
	return rb
}

// FunctionScore set the FunctionScore property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) FunctionScore(functionscore FunctionScoreQuery) *QueryContainerBuilder {
	rb.v.FunctionScore = &functionscore
	return rb
}

// Fuzzy set the Fuzzy property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Fuzzy(value map[Field]FuzzyQuery) *QueryContainerBuilder {
	rb.v.Fuzzy = value
	return rb
}

// GeoBoundingBox set the GeoBoundingBox property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) GeoBoundingBox(geoboundingbox GeoBoundingBoxQuery) *QueryContainerBuilder {
	rb.v.GeoBoundingBox = &geoboundingbox
	return rb
}

// GeoDistance set the GeoDistance property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) GeoDistance(geodistance GeoDistanceQuery) *QueryContainerBuilder {
	rb.v.GeoDistance = &geodistance
	return rb
}

// GeoPolygon set the GeoPolygon property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) GeoPolygon(geopolygon GeoPolygonQuery) *QueryContainerBuilder {
	rb.v.GeoPolygon = &geopolygon
	return rb
}

// GeoShape set the GeoShape property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) GeoShape(geoshape GeoShapeQuery) *QueryContainerBuilder {
	rb.v.GeoShape = &geoshape
	return rb
}

// HasChild set the HasChild property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) HasChild(haschild HasChildQuery) *QueryContainerBuilder {
	rb.v.HasChild = &haschild
	return rb
}

// HasParent set the HasParent property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) HasParent(hasparent HasParentQuery) *QueryContainerBuilder {
	rb.v.HasParent = &hasparent
	return rb
}

// Ids set the Ids property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Ids(ids IdsQuery) *QueryContainerBuilder {
	rb.v.Ids = &ids
	return rb
}

// Intervals set the Intervals property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Intervals(value map[Field]IntervalsQuery) *QueryContainerBuilder {
	rb.v.Intervals = value
	return rb
}

// Knn set the Knn property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Knn(knn KnnQuery) *QueryContainerBuilder {
	rb.v.Knn = &knn
	return rb
}

// Match set the Match property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Match(value map[Field]MatchQuery) *QueryContainerBuilder {
	rb.v.Match = value
	return rb
}

// MatchAll set the MatchAll property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) MatchAll(matchall MatchAllQuery) *QueryContainerBuilder {
	rb.v.MatchAll = &matchall
	return rb
}

// MatchBoolPrefix set the MatchBoolPrefix property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) MatchBoolPrefix(value map[Field]MatchBoolPrefixQuery) *QueryContainerBuilder {
	rb.v.MatchBoolPrefix = value
	return rb
}

// MatchNone set the MatchNone property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) MatchNone(matchnone MatchNoneQuery) *QueryContainerBuilder {
	rb.v.MatchNone = &matchnone
	return rb
}

// MatchPhrase set the MatchPhrase property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) MatchPhrase(value map[Field]MatchPhraseQuery) *QueryContainerBuilder {
	rb.v.MatchPhrase = value
	return rb
}

// MatchPhrasePrefix set the MatchPhrasePrefix property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) MatchPhrasePrefix(value map[Field]MatchPhrasePrefixQuery) *QueryContainerBuilder {
	rb.v.MatchPhrasePrefix = value
	return rb
}

// MoreLikeThis set the MoreLikeThis property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) MoreLikeThis(morelikethis MoreLikeThisQuery) *QueryContainerBuilder {
	rb.v.MoreLikeThis = &morelikethis
	return rb
}

// MultiMatch set the MultiMatch property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) MultiMatch(multimatch MultiMatchQuery) *QueryContainerBuilder {
	rb.v.MultiMatch = &multimatch
	return rb
}

// Nested set the Nested property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Nested(nested NestedQuery) *QueryContainerBuilder {
	rb.v.Nested = &nested
	return rb
}

// ParentId set the ParentId property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) ParentId(parentid ParentIdQuery) *QueryContainerBuilder {
	rb.v.ParentId = &parentid
	return rb
}

// Percolate set the Percolate property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Percolate(percolate PercolateQuery) *QueryContainerBuilder {
	rb.v.Percolate = &percolate
	return rb
}

// Pinned set the Pinned property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Pinned(pinned PinnedQuery) *QueryContainerBuilder {
	rb.v.Pinned = &pinned
	return rb
}

// Prefix set the Prefix property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Prefix(value map[Field]PrefixQuery) *QueryContainerBuilder {
	rb.v.Prefix = value
	return rb
}

// QueryString set the QueryString property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) QueryString(querystring QueryStringQuery) *QueryContainerBuilder {
	rb.v.QueryString = &querystring
	return rb
}

// Range set the Range property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Range(value map[Field]RangeQuery) *QueryContainerBuilder {
	rb.v.Range = value
	return rb
}

// RankFeature set the RankFeature property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) RankFeature(rankfeature RankFeatureQuery) *QueryContainerBuilder {
	rb.v.RankFeature = &rankfeature
	return rb
}

// Regexp set the Regexp property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Regexp(value map[Field]RegexpQuery) *QueryContainerBuilder {
	rb.v.Regexp = value
	return rb
}

// Script set the Script property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Script(script ScriptQuery) *QueryContainerBuilder {
	rb.v.Script = &script
	return rb
}

// ScriptScore set the ScriptScore property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) ScriptScore(scriptscore ScriptScoreQuery) *QueryContainerBuilder {
	rb.v.ScriptScore = &scriptscore
	return rb
}

// Shape set the Shape property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Shape(shape ShapeQuery) *QueryContainerBuilder {
	rb.v.Shape = &shape
	return rb
}

// SimpleQueryString set the SimpleQueryString property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SimpleQueryString(simplequerystring SimpleQueryStringQuery) *QueryContainerBuilder {
	rb.v.SimpleQueryString = &simplequerystring
	return rb
}

// SpanContaining set the SpanContaining property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SpanContaining(spancontaining SpanContainingQuery) *QueryContainerBuilder {
	rb.v.SpanContaining = &spancontaining
	return rb
}

// SpanFirst set the SpanFirst property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SpanFirst(spanfirst SpanFirstQuery) *QueryContainerBuilder {
	rb.v.SpanFirst = &spanfirst
	return rb
}

// SpanMulti set the SpanMulti property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SpanMulti(spanmulti SpanMultiTermQuery) *QueryContainerBuilder {
	rb.v.SpanMulti = &spanmulti
	return rb
}

// SpanNear set the SpanNear property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SpanNear(spannear SpanNearQuery) *QueryContainerBuilder {
	rb.v.SpanNear = &spannear
	return rb
}

// SpanNot set the SpanNot property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SpanNot(spannot SpanNotQuery) *QueryContainerBuilder {
	rb.v.SpanNot = &spannot
	return rb
}

// SpanOr set the SpanOr property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SpanOr(spanor SpanOrQuery) *QueryContainerBuilder {
	rb.v.SpanOr = &spanor
	return rb
}

// SpanTerm set the SpanTerm property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SpanTerm(value map[Field]SpanTermQuery) *QueryContainerBuilder {
	rb.v.SpanTerm = value
	return rb
}

// SpanWithin set the SpanWithin property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) SpanWithin(spanwithin SpanWithinQuery) *QueryContainerBuilder {
	rb.v.SpanWithin = &spanwithin
	return rb
}

// Term set the Term property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Term(value map[Field]TermQuery) *QueryContainerBuilder {
	rb.v.Term = value
	return rb
}

// Terms set the Terms property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Terms(terms TermsQuery) *QueryContainerBuilder {
	rb.v.Terms = &terms
	return rb
}

// TermsSet set the TermsSet property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) TermsSet(value map[Field]TermsSetQuery) *QueryContainerBuilder {
	rb.v.TermsSet = value
	return rb
}

// Type set the Type property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Type_(type_ TypeQuery) *QueryContainerBuilder {
	rb.v.Type = &type_
	return rb
}

// Wildcard set the Wildcard property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Wildcard(value map[Field]WildcardQuery) *QueryContainerBuilder {
	rb.v.Wildcard = value
	return rb
}

// Wrapper set the Wrapper property for QueryContainerBuilder.
func (rb *QueryContainerBuilder) Wrapper(wrapper WrapperQuery) *QueryContainerBuilder {
	rb.v.Wrapper = &wrapper
	return rb
}
