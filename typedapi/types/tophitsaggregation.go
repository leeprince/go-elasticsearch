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

// TopHitsAggregation type.
type TopHitsAggregation struct {
	DocvalueFields   *Fields                `json:"docvalue_fields,omitempty"`
	Explain          *bool                  `json:"explain,omitempty"`
	Field            *Field                 `json:"field,omitempty"`
	From             *int                   `json:"from,omitempty"`
	Highlight        *Highlight             `json:"highlight,omitempty"`
	Missing          *Missing               `json:"missing,omitempty"`
	Script           *Script                `json:"script,omitempty"`
	ScriptFields     map[string]ScriptField `json:"script_fields,omitempty"`
	SeqNoPrimaryTerm *bool                  `json:"seq_no_primary_term,omitempty"`
	Size             *int                   `json:"size,omitempty"`
	Sort             *Sort                  `json:"sort,omitempty"`
	Source_          *SourceConfig          `json:"_source,omitempty"`
	StoredFields     *Fields                `json:"stored_fields,omitempty"`
	TrackScores      *bool                  `json:"track_scores,omitempty"`
	Version          *bool                  `json:"version,omitempty"`
}

// TopHitsAggregationBuilder holds TopHitsAggregation struct and provides a builder API.
type TopHitsAggregationBuilder struct {
	v *TopHitsAggregation
}

// NewTopHitsAggregation provides a builder for the TopHitsAggregation struct.
func NewTopHitsAggregation() *TopHitsAggregationBuilder {
	r := TopHitsAggregationBuilder{
		&TopHitsAggregation{
			ScriptFields: make(map[string]ScriptField, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TopHitsAggregation struct
func (rb *TopHitsAggregationBuilder) Build() TopHitsAggregation {
	return *rb.v
}

// DocvalueFields set the DocvalueFields property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) DocvalueFields(docvaluefields Fields) *TopHitsAggregationBuilder {
	rb.v.DocvalueFields = &docvaluefields
	return rb
}

// Explain set the Explain property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Explain(explain bool) *TopHitsAggregationBuilder {
	rb.v.Explain = &explain
	return rb
}

// Field set the Field property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Field(field Field) *TopHitsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

// From set the From property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) From(from int) *TopHitsAggregationBuilder {
	rb.v.From = &from
	return rb
}

// Highlight set the Highlight property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Highlight(highlight Highlight) *TopHitsAggregationBuilder {
	rb.v.Highlight = &highlight
	return rb
}

// Missing set the Missing property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Missing(missing Missing) *TopHitsAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

// Script set the Script property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Script(script Script) *TopHitsAggregationBuilder {
	rb.v.Script = &script
	return rb
}

// ScriptFields set the ScriptFields property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) ScriptFields(value map[string]ScriptField) *TopHitsAggregationBuilder {
	rb.v.ScriptFields = value
	return rb
}

// SeqNoPrimaryTerm set the SeqNoPrimaryTerm property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) SeqNoPrimaryTerm(seqnoprimaryterm bool) *TopHitsAggregationBuilder {
	rb.v.SeqNoPrimaryTerm = &seqnoprimaryterm
	return rb
}

// Size set the Size property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Size(size int) *TopHitsAggregationBuilder {
	rb.v.Size = &size
	return rb
}

// Sort set the Sort property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Sort(sort Sort) *TopHitsAggregationBuilder {
	rb.v.Sort = &sort
	return rb
}

// Source_ set the Source_ property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Source_(source_ SourceConfig) *TopHitsAggregationBuilder {
	rb.v.Source_ = &source_
	return rb
}

// StoredFields set the StoredFields property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) StoredFields(storedfields Fields) *TopHitsAggregationBuilder {
	rb.v.StoredFields = &storedfields
	return rb
}

// TrackScores set the TrackScores property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) TrackScores(trackscores bool) *TopHitsAggregationBuilder {
	rb.v.TrackScores = &trackscores
	return rb
}

// Version set the Version property for TopHitsAggregationBuilder.
func (rb *TopHitsAggregationBuilder) Version(version bool) *TopHitsAggregationBuilder {
	rb.v.Version = &version
	return rb
}
