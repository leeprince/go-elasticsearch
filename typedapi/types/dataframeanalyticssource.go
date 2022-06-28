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

// DataframeAnalyticsSource type.
type DataframeAnalyticsSource struct {
	// Index Index or indices on which to perform the analysis. It can be a single index
	// or index pattern as well as an array of indices or patterns. NOTE: If your
	// source indices contain documents with the same IDs, only the document that is
	// indexed last appears in the destination index.
	Index Indices `json:"index"`
	// Query The Elasticsearch query domain-specific language (DSL). This value
	// corresponds to the query object in an Elasticsearch search POST body. All the
	// options that are supported by Elasticsearch can be used, as this object is
	// passed verbatim to Elasticsearch. By default, this property has the following
	// value: {"match_all": {}}.
	Query *QueryContainer `json:"query,omitempty"`
	// RuntimeMappings Definitions of runtime fields that will become part of the mapping of the
	// destination index.
	RuntimeMappings *RuntimeFields `json:"runtime_mappings,omitempty"`
	// Source_ Specify `includes` and/or `excludes patterns to select which fields will be
	// present in the destination. Fields that are excluded cannot be included in
	// the analysis.
	Source_ *DataframeAnalysisAnalyzedFields `json:"_source,omitempty"`
}

// DataframeAnalyticsSourceBuilder holds DataframeAnalyticsSource struct and provides a builder API.
type DataframeAnalyticsSourceBuilder struct {
	v *DataframeAnalyticsSource
}

// NewDataframeAnalyticsSource provides a builder for the DataframeAnalyticsSource struct.
func NewDataframeAnalyticsSource() *DataframeAnalyticsSourceBuilder {
	r := DataframeAnalyticsSourceBuilder{
		&DataframeAnalyticsSource{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeAnalyticsSource struct
func (rb *DataframeAnalyticsSourceBuilder) Build() DataframeAnalyticsSource {
	return *rb.v
}

// Index Index or indices on which to perform the analysis. It can be a single index
// or index pattern as well as an array of indices or patterns. NOTE: If your
// source indices contain documents with the same IDs, only the document that is
// indexed last appears in the destination index.
func (rb *DataframeAnalyticsSourceBuilder) Index(index Indices) *DataframeAnalyticsSourceBuilder {
	rb.v.Index = index
	return rb
}

// Query The Elasticsearch query domain-specific language (DSL). This value
// corresponds to the query object in an Elasticsearch search POST body. All the
// options that are supported by Elasticsearch can be used, as this object is
// passed verbatim to Elasticsearch. By default, this property has the following
// value: {"match_all": {}}.
func (rb *DataframeAnalyticsSourceBuilder) Query(query QueryContainer) *DataframeAnalyticsSourceBuilder {
	rb.v.Query = &query
	return rb
}

// RuntimeMappings Definitions of runtime fields that will become part of the mapping of the
// destination index.
func (rb *DataframeAnalyticsSourceBuilder) RuntimeMappings(runtimemappings RuntimeFields) *DataframeAnalyticsSourceBuilder {
	rb.v.RuntimeMappings = &runtimemappings
	return rb
}

// Source_ Specify `includes` and/or `excludes patterns to select which fields will be
// present in the destination. Fields that are excluded cannot be included in
// the analysis.
func (rb *DataframeAnalyticsSourceBuilder) Source_(source_ DataframeAnalysisAnalyzedFields) *DataframeAnalyticsSourceBuilder {
	rb.v.Source_ = &source_
	return rb
}
