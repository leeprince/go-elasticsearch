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

// ElisionTokenFilter type.
type ElisionTokenFilter struct {
	Articles     []string       `json:"articles"`
	ArticlesCase bool           `json:"articles_case"`
	Type         string         `json:"type,omitempty"`
	Version      *VersionString `json:"version,omitempty"`
}

// ElisionTokenFilterBuilder holds ElisionTokenFilter struct and provides a builder API.
type ElisionTokenFilterBuilder struct {
	v *ElisionTokenFilter
}

// NewElisionTokenFilter provides a builder for the ElisionTokenFilter struct.
func NewElisionTokenFilter() *ElisionTokenFilterBuilder {
	r := ElisionTokenFilterBuilder{
		&ElisionTokenFilter{},
	}

	r.v.Type = "elision"

	return &r
}

// Build finalize the chain and returns the ElisionTokenFilter struct
func (rb *ElisionTokenFilterBuilder) Build() ElisionTokenFilter {
	return *rb.v
}

// Articles set the Articles property for ElisionTokenFilterBuilder.
func (rb *ElisionTokenFilterBuilder) Articles(articles ...string) *ElisionTokenFilterBuilder {
	rb.v.Articles = append(rb.v.Articles, articles...)
	return rb
}

// ArticlesCase set the ArticlesCase property for ElisionTokenFilterBuilder.
func (rb *ElisionTokenFilterBuilder) ArticlesCase(articlescase bool) *ElisionTokenFilterBuilder {
	rb.v.ArticlesCase = articlescase
	return rb
}

// Type set the Type property for ElisionTokenFilterBuilder.

// Version set the Version property for ElisionTokenFilterBuilder.
func (rb *ElisionTokenFilterBuilder) Version(version VersionString) *ElisionTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
