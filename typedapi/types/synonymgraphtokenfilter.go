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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/synonymformat"
)

// SynonymGraphTokenFilter type.
type SynonymGraphTokenFilter struct {
	Expand       *bool                        `json:"expand,omitempty"`
	Format       *synonymformat.SynonymFormat `json:"format,omitempty"`
	Lenient      *bool                        `json:"lenient,omitempty"`
	Synonyms     []string                     `json:"synonyms,omitempty"`
	SynonymsPath *string                      `json:"synonyms_path,omitempty"`
	Tokenizer    *string                      `json:"tokenizer,omitempty"`
	Type         string                       `json:"type,omitempty"`
	Updateable   *bool                        `json:"updateable,omitempty"`
	Version      *VersionString               `json:"version,omitempty"`
}

// SynonymGraphTokenFilterBuilder holds SynonymGraphTokenFilter struct and provides a builder API.
type SynonymGraphTokenFilterBuilder struct {
	v *SynonymGraphTokenFilter
}

// NewSynonymGraphTokenFilter provides a builder for the SynonymGraphTokenFilter struct.
func NewSynonymGraphTokenFilter() *SynonymGraphTokenFilterBuilder {
	r := SynonymGraphTokenFilterBuilder{
		&SynonymGraphTokenFilter{},
	}

	r.v.Type = "synonym_graph"

	return &r
}

// Build finalize the chain and returns the SynonymGraphTokenFilter struct
func (rb *SynonymGraphTokenFilterBuilder) Build() SynonymGraphTokenFilter {
	return *rb.v
}

// Expand set the Expand property for SynonymGraphTokenFilterBuilder.
func (rb *SynonymGraphTokenFilterBuilder) Expand(expand bool) *SynonymGraphTokenFilterBuilder {
	rb.v.Expand = &expand
	return rb
}

// Format set the Format property for SynonymGraphTokenFilterBuilder.
func (rb *SynonymGraphTokenFilterBuilder) Format(format synonymformat.SynonymFormat) *SynonymGraphTokenFilterBuilder {
	rb.v.Format = &format
	return rb
}

// Lenient set the Lenient property for SynonymGraphTokenFilterBuilder.
func (rb *SynonymGraphTokenFilterBuilder) Lenient(lenient bool) *SynonymGraphTokenFilterBuilder {
	rb.v.Lenient = &lenient
	return rb
}

// Synonyms set the Synonyms property for SynonymGraphTokenFilterBuilder.
func (rb *SynonymGraphTokenFilterBuilder) Synonyms(synonyms ...string) *SynonymGraphTokenFilterBuilder {
	rb.v.Synonyms = append(rb.v.Synonyms, synonyms...)
	return rb
}

// SynonymsPath set the SynonymsPath property for SynonymGraphTokenFilterBuilder.
func (rb *SynonymGraphTokenFilterBuilder) SynonymsPath(synonymspath string) *SynonymGraphTokenFilterBuilder {
	rb.v.SynonymsPath = &synonymspath
	return rb
}

// Tokenizer set the Tokenizer property for SynonymGraphTokenFilterBuilder.
func (rb *SynonymGraphTokenFilterBuilder) Tokenizer(tokenizer string) *SynonymGraphTokenFilterBuilder {
	rb.v.Tokenizer = &tokenizer
	return rb
}

// Type set the Type property for SynonymGraphTokenFilterBuilder.

// Updateable set the Updateable property for SynonymGraphTokenFilterBuilder.
func (rb *SynonymGraphTokenFilterBuilder) Updateable(updateable bool) *SynonymGraphTokenFilterBuilder {
	rb.v.Updateable = &updateable
	return rb
}

// Version set the Version property for SynonymGraphTokenFilterBuilder.
func (rb *SynonymGraphTokenFilterBuilder) Version(version VersionString) *SynonymGraphTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
