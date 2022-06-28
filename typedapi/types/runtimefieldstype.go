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

// RuntimeFieldsType type.
type RuntimeFieldsType struct {
	CharsMax        int64    `json:"chars_max"`
	CharsTotal      int64    `json:"chars_total"`
	Count           int64    `json:"count"`
	DocMax          int64    `json:"doc_max"`
	DocTotal        int64    `json:"doc_total"`
	IndexCount      int64    `json:"index_count"`
	Lang            []string `json:"lang"`
	LinesMax        int64    `json:"lines_max"`
	LinesTotal      int64    `json:"lines_total"`
	Name            Field    `json:"name"`
	ScriptlessCount int64    `json:"scriptless_count"`
	ShadowedCount   int64    `json:"shadowed_count"`
	SourceMax       int64    `json:"source_max"`
	SourceTotal     int64    `json:"source_total"`
}

// RuntimeFieldsTypeBuilder holds RuntimeFieldsType struct and provides a builder API.
type RuntimeFieldsTypeBuilder struct {
	v *RuntimeFieldsType
}

// NewRuntimeFieldsType provides a builder for the RuntimeFieldsType struct.
func NewRuntimeFieldsType() *RuntimeFieldsTypeBuilder {
	r := RuntimeFieldsTypeBuilder{
		&RuntimeFieldsType{},
	}

	return &r
}

// Build finalize the chain and returns the RuntimeFieldsType struct
func (rb *RuntimeFieldsTypeBuilder) Build() RuntimeFieldsType {
	return *rb.v
}

// CharsMax set the CharsMax property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) CharsMax(charsmax int64) *RuntimeFieldsTypeBuilder {
	rb.v.CharsMax = charsmax
	return rb
}

// CharsTotal set the CharsTotal property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) CharsTotal(charstotal int64) *RuntimeFieldsTypeBuilder {
	rb.v.CharsTotal = charstotal
	return rb
}

// Count set the Count property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) Count(count int64) *RuntimeFieldsTypeBuilder {
	rb.v.Count = count
	return rb
}

// DocMax set the DocMax property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) DocMax(docmax int64) *RuntimeFieldsTypeBuilder {
	rb.v.DocMax = docmax
	return rb
}

// DocTotal set the DocTotal property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) DocTotal(doctotal int64) *RuntimeFieldsTypeBuilder {
	rb.v.DocTotal = doctotal
	return rb
}

// IndexCount set the IndexCount property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) IndexCount(indexcount int64) *RuntimeFieldsTypeBuilder {
	rb.v.IndexCount = indexcount
	return rb
}

// Lang set the Lang property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) Lang(lang ...string) *RuntimeFieldsTypeBuilder {
	rb.v.Lang = append(rb.v.Lang, lang...)
	return rb
}

// LinesMax set the LinesMax property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) LinesMax(linesmax int64) *RuntimeFieldsTypeBuilder {
	rb.v.LinesMax = linesmax
	return rb
}

// LinesTotal set the LinesTotal property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) LinesTotal(linestotal int64) *RuntimeFieldsTypeBuilder {
	rb.v.LinesTotal = linestotal
	return rb
}

// Name set the Name property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) Name(name Field) *RuntimeFieldsTypeBuilder {
	rb.v.Name = name
	return rb
}

// ScriptlessCount set the ScriptlessCount property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) ScriptlessCount(scriptlesscount int64) *RuntimeFieldsTypeBuilder {
	rb.v.ScriptlessCount = scriptlesscount
	return rb
}

// ShadowedCount set the ShadowedCount property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) ShadowedCount(shadowedcount int64) *RuntimeFieldsTypeBuilder {
	rb.v.ShadowedCount = shadowedcount
	return rb
}

// SourceMax set the SourceMax property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) SourceMax(sourcemax int64) *RuntimeFieldsTypeBuilder {
	rb.v.SourceMax = sourcemax
	return rb
}

// SourceTotal set the SourceTotal property for RuntimeFieldsTypeBuilder.
func (rb *RuntimeFieldsTypeBuilder) SourceTotal(sourcetotal int64) *RuntimeFieldsTypeBuilder {
	rb.v.SourceTotal = sourcetotal
	return rb
}
