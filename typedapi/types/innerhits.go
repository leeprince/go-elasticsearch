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

// InnerHits type.
type InnerHits struct {
	Collapse         *FieldCollapse        `json:"collapse,omitempty"`
	DocvalueFields   []FieldAndFormat      `json:"docvalue_fields,omitempty"`
	Explain          *bool                 `json:"explain,omitempty"`
	Fields           *Fields               `json:"fields,omitempty"`
	From             *int                  `json:"from,omitempty"`
	Highlight        *Highlight            `json:"highlight,omitempty"`
	IgnoreUnmapped   *bool                 `json:"ignore_unmapped,omitempty"`
	Name             *Name                 `json:"name,omitempty"`
	ScriptFields     map[Field]ScriptField `json:"script_fields,omitempty"`
	SeqNoPrimaryTerm *bool                 `json:"seq_no_primary_term,omitempty"`
	Size             *int                  `json:"size,omitempty"`
	Sort             *Sort                 `json:"sort,omitempty"`
	Source_          *SourceConfig         `json:"_source,omitempty"`
	StoredField      *Fields               `json:"stored_field,omitempty"`
	TrackScores      *bool                 `json:"track_scores,omitempty"`
	Version          *bool                 `json:"version,omitempty"`
}

// InnerHitsBuilder holds InnerHits struct and provides a builder API.
type InnerHitsBuilder struct {
	v *InnerHits
}

// NewInnerHits provides a builder for the InnerHits struct.
func NewInnerHits() *InnerHitsBuilder {
	r := InnerHitsBuilder{
		&InnerHits{
			ScriptFields: make(map[Field]ScriptField, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the InnerHits struct
func (rb *InnerHitsBuilder) Build() InnerHits {
	return *rb.v
}

// Collapse set the Collapse property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Collapse(collapse FieldCollapse) *InnerHitsBuilder {
	rb.v.Collapse = &collapse
	return rb
}

// DocvalueFields set the DocvalueFields property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) DocvalueFields(docvalue_fields ...FieldAndFormat) *InnerHitsBuilder {
	rb.v.DocvalueFields = append(rb.v.DocvalueFields, docvalue_fields...)
	return rb
}

// Explain set the Explain property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Explain(explain bool) *InnerHitsBuilder {
	rb.v.Explain = &explain
	return rb
}

// Fields set the Fields property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Fields(fields Fields) *InnerHitsBuilder {
	rb.v.Fields = &fields
	return rb
}

// From set the From property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) From(from int) *InnerHitsBuilder {
	rb.v.From = &from
	return rb
}

// Highlight set the Highlight property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Highlight(highlight Highlight) *InnerHitsBuilder {
	rb.v.Highlight = &highlight
	return rb
}

// IgnoreUnmapped set the IgnoreUnmapped property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) IgnoreUnmapped(ignoreunmapped bool) *InnerHitsBuilder {
	rb.v.IgnoreUnmapped = &ignoreunmapped
	return rb
}

// Name set the Name property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Name(name Name) *InnerHitsBuilder {
	rb.v.Name = &name
	return rb
}

// ScriptFields set the ScriptFields property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) ScriptFields(value map[Field]ScriptField) *InnerHitsBuilder {
	rb.v.ScriptFields = value
	return rb
}

// SeqNoPrimaryTerm set the SeqNoPrimaryTerm property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) SeqNoPrimaryTerm(seqnoprimaryterm bool) *InnerHitsBuilder {
	rb.v.SeqNoPrimaryTerm = &seqnoprimaryterm
	return rb
}

// Size set the Size property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Size(size int) *InnerHitsBuilder {
	rb.v.Size = &size
	return rb
}

// Sort set the Sort property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Sort(sort Sort) *InnerHitsBuilder {
	rb.v.Sort = &sort
	return rb
}

// Source_ set the Source_ property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Source_(source_ SourceConfig) *InnerHitsBuilder {
	rb.v.Source_ = &source_
	return rb
}

// StoredField set the StoredField property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) StoredField(storedfield Fields) *InnerHitsBuilder {
	rb.v.StoredField = &storedfield
	return rb
}

// TrackScores set the TrackScores property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) TrackScores(trackscores bool) *InnerHitsBuilder {
	rb.v.TrackScores = &trackscores
	return rb
}

// Version set the Version property for InnerHitsBuilder.
func (rb *InnerHitsBuilder) Version(version bool) *InnerHitsBuilder {
	rb.v.Version = &version
	return rb
}
