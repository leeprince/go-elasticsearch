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

// LimitTokenCountTokenFilter type.
type LimitTokenCountTokenFilter struct {
	ConsumeAllTokens bool           `json:"consume_all_tokens"`
	MaxTokenCount    int            `json:"max_token_count"`
	Type             string         `json:"type,omitempty"`
	Version          *VersionString `json:"version,omitempty"`
}

// LimitTokenCountTokenFilterBuilder holds LimitTokenCountTokenFilter struct and provides a builder API.
type LimitTokenCountTokenFilterBuilder struct {
	v *LimitTokenCountTokenFilter
}

// NewLimitTokenCountTokenFilter provides a builder for the LimitTokenCountTokenFilter struct.
func NewLimitTokenCountTokenFilter() *LimitTokenCountTokenFilterBuilder {
	r := LimitTokenCountTokenFilterBuilder{
		&LimitTokenCountTokenFilter{},
	}

	r.v.Type = "limit"

	return &r
}

// Build finalize the chain and returns the LimitTokenCountTokenFilter struct
func (rb *LimitTokenCountTokenFilterBuilder) Build() LimitTokenCountTokenFilter {
	return *rb.v
}

// ConsumeAllTokens set the ConsumeAllTokens property for LimitTokenCountTokenFilterBuilder.
func (rb *LimitTokenCountTokenFilterBuilder) ConsumeAllTokens(consumealltokens bool) *LimitTokenCountTokenFilterBuilder {
	rb.v.ConsumeAllTokens = consumealltokens
	return rb
}

// MaxTokenCount set the MaxTokenCount property for LimitTokenCountTokenFilterBuilder.
func (rb *LimitTokenCountTokenFilterBuilder) MaxTokenCount(maxtokencount int) *LimitTokenCountTokenFilterBuilder {
	rb.v.MaxTokenCount = maxtokencount
	return rb
}

// Type set the Type property for LimitTokenCountTokenFilterBuilder.

// Version set the Version property for LimitTokenCountTokenFilterBuilder.
func (rb *LimitTokenCountTokenFilterBuilder) Version(version VersionString) *LimitTokenCountTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
