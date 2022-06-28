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

// CharFilterDefinition holds the union for the following types:
//     HtmlStripCharFilter
//     IcuNormalizationCharFilter
//     KuromojiIterationMarkCharFilter
//     MappingCharFilter
//     PatternReplaceCharFilter
type CharFilterDefinition interface{}

// CharFilterDefinitionBuilder holds CharFilterDefinition struct and provides a builder API.
type CharFilterDefinitionBuilder struct {
	v CharFilterDefinition
}

// NewCharFilterDefinition provides a builder for the CharFilterDefinition struct.
func NewCharFilterDefinition() *CharFilterDefinitionBuilder {
	return &CharFilterDefinitionBuilder{}
}

// Build finalize the chain and returns the CharFilterDefinition struct
func (u *CharFilterDefinitionBuilder) Build() CharFilterDefinition {
	return u.v
}

// HtmlStripCharFilter set the HtmlStripCharFilter property for CharFilterDefinitionBuilder.
func (u *CharFilterDefinitionBuilder) HtmlStripCharFilter(v HtmlStripCharFilter) *CharFilterDefinitionBuilder {
	u.v = v
	return u
}

// IcuNormalizationCharFilter set the IcuNormalizationCharFilter property for CharFilterDefinitionBuilder.
func (u *CharFilterDefinitionBuilder) IcuNormalizationCharFilter(v IcuNormalizationCharFilter) *CharFilterDefinitionBuilder {
	u.v = v
	return u
}

// KuromojiIterationMarkCharFilter set the KuromojiIterationMarkCharFilter property for CharFilterDefinitionBuilder.
func (u *CharFilterDefinitionBuilder) KuromojiIterationMarkCharFilter(v KuromojiIterationMarkCharFilter) *CharFilterDefinitionBuilder {
	u.v = v
	return u
}

// MappingCharFilter set the MappingCharFilter property for CharFilterDefinitionBuilder.
func (u *CharFilterDefinitionBuilder) MappingCharFilter(v MappingCharFilter) *CharFilterDefinitionBuilder {
	u.v = v
	return u
}

// PatternReplaceCharFilter set the PatternReplaceCharFilter property for CharFilterDefinitionBuilder.
func (u *CharFilterDefinitionBuilder) PatternReplaceCharFilter(v PatternReplaceCharFilter) *CharFilterDefinitionBuilder {
	u.v = v
	return u
}
