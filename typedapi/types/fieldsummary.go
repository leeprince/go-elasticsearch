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

// FieldSummary type.
type FieldSummary struct {
	Any           uint          `json:"any"`
	DocValues     uint          `json:"doc_values"`
	InvertedIndex InvertedIndex `json:"inverted_index"`
	KnnVectors    uint          `json:"knn_vectors"`
	Norms         uint          `json:"norms"`
	Points        uint          `json:"points"`
	StoredFields  uint          `json:"stored_fields"`
	TermVectors   uint          `json:"term_vectors"`
}

// FieldSummaryBuilder holds FieldSummary struct and provides a builder API.
type FieldSummaryBuilder struct {
	v *FieldSummary
}

// NewFieldSummary provides a builder for the FieldSummary struct.
func NewFieldSummary() *FieldSummaryBuilder {
	r := FieldSummaryBuilder{
		&FieldSummary{},
	}

	return &r
}

// Build finalize the chain and returns the FieldSummary struct
func (rb *FieldSummaryBuilder) Build() FieldSummary {
	return *rb.v
}

// Any set the Any property for FieldSummaryBuilder.
func (rb *FieldSummaryBuilder) Any(any uint) *FieldSummaryBuilder {
	rb.v.Any = any
	return rb
}

// DocValues set the DocValues property for FieldSummaryBuilder.
func (rb *FieldSummaryBuilder) DocValues(docvalues uint) *FieldSummaryBuilder {
	rb.v.DocValues = docvalues
	return rb
}

// InvertedIndex set the InvertedIndex property for FieldSummaryBuilder.
func (rb *FieldSummaryBuilder) InvertedIndex(invertedindex InvertedIndex) *FieldSummaryBuilder {
	rb.v.InvertedIndex = invertedindex
	return rb
}

// KnnVectors set the KnnVectors property for FieldSummaryBuilder.
func (rb *FieldSummaryBuilder) KnnVectors(knnvectors uint) *FieldSummaryBuilder {
	rb.v.KnnVectors = knnvectors
	return rb
}

// Norms set the Norms property for FieldSummaryBuilder.
func (rb *FieldSummaryBuilder) Norms(norms uint) *FieldSummaryBuilder {
	rb.v.Norms = norms
	return rb
}

// Points set the Points property for FieldSummaryBuilder.
func (rb *FieldSummaryBuilder) Points(points uint) *FieldSummaryBuilder {
	rb.v.Points = points
	return rb
}

// StoredFields set the StoredFields property for FieldSummaryBuilder.
func (rb *FieldSummaryBuilder) StoredFields(storedfields uint) *FieldSummaryBuilder {
	rb.v.StoredFields = storedfields
	return rb
}

// TermVectors set the TermVectors property for FieldSummaryBuilder.
func (rb *FieldSummaryBuilder) TermVectors(termvectors uint) *FieldSummaryBuilder {
	rb.v.TermVectors = termvectors
	return rb
}
