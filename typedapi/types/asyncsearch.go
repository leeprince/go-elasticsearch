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

// AsyncSearch type.
type AsyncSearch struct {
	Aggregations    map[AggregateName]Aggregate  `json:"aggregations,omitempty"`
	Clusters_       *ClusterStatistics           `json:"_clusters,omitempty"`
	Fields          map[string]interface{}       `json:"fields,omitempty"`
	Hits            HitsMetadata                 `json:"hits"`
	MaxScore        *float64                     `json:"max_score,omitempty"`
	NumReducePhases *int64                       `json:"num_reduce_phases,omitempty"`
	PitId           *Id                          `json:"pit_id,omitempty"`
	Profile         *Profile                     `json:"profile,omitempty"`
	ScrollId_       *Id                          `json:"_scroll_id,omitempty"`
	Shards_         ShardStatistics              `json:"_shards"`
	Suggest         map[SuggestionName][]Suggest `json:"suggest,omitempty"`
	TerminatedEarly *bool                        `json:"terminated_early,omitempty"`
	TimedOut        bool                         `json:"timed_out"`
	Took            int64                        `json:"took"`
}

// AsyncSearchBuilder holds AsyncSearch struct and provides a builder API.
type AsyncSearchBuilder struct {
	v *AsyncSearch
}

// NewAsyncSearch provides a builder for the AsyncSearch struct.
func NewAsyncSearch() *AsyncSearchBuilder {
	r := AsyncSearchBuilder{
		&AsyncSearch{
			Aggregations: make(map[AggregateName]Aggregate, 0),
			Fields:       make(map[string]interface{}, 0),
			Suggest:      make(map[SuggestionName][]Suggest, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AsyncSearch struct
func (rb *AsyncSearchBuilder) Build() AsyncSearch {
	return *rb.v
}

// Aggregations set the Aggregations property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) Aggregations(value map[AggregateName]Aggregate) *AsyncSearchBuilder {
	rb.v.Aggregations = value
	return rb
}

// Clusters_ set the Clusters_ property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) Clusters_(clusters_ ClusterStatistics) *AsyncSearchBuilder {
	rb.v.Clusters_ = &clusters_
	return rb
}

// Fields set the Fields property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) Fields(value map[string]interface{}) *AsyncSearchBuilder {
	rb.v.Fields = value
	return rb
}

// Hits set the Hits property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) Hits(hits HitsMetadata) *AsyncSearchBuilder {
	rb.v.Hits = hits
	return rb
}

// MaxScore set the MaxScore property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) MaxScore(maxscore float64) *AsyncSearchBuilder {
	rb.v.MaxScore = &maxscore
	return rb
}

// NumReducePhases set the NumReducePhases property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) NumReducePhases(numreducephases int64) *AsyncSearchBuilder {
	rb.v.NumReducePhases = &numreducephases
	return rb
}

// PitId set the PitId property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) PitId(pitid Id) *AsyncSearchBuilder {
	rb.v.PitId = &pitid
	return rb
}

// Profile set the Profile property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) Profile(profile Profile) *AsyncSearchBuilder {
	rb.v.Profile = &profile
	return rb
}

// ScrollId_ set the ScrollId_ property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) ScrollId_(scrollid_ Id) *AsyncSearchBuilder {
	rb.v.ScrollId_ = &scrollid_
	return rb
}

// Shards_ set the Shards_ property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) Shards_(shards_ ShardStatistics) *AsyncSearchBuilder {
	rb.v.Shards_ = shards_
	return rb
}

// Suggest set the Suggest property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) Suggest(value map[SuggestionName][]Suggest) *AsyncSearchBuilder {
	rb.v.Suggest = value
	return rb
}

// TerminatedEarly set the TerminatedEarly property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) TerminatedEarly(terminatedearly bool) *AsyncSearchBuilder {
	rb.v.TerminatedEarly = &terminatedearly
	return rb
}

// TimedOut set the TimedOut property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) TimedOut(timedout bool) *AsyncSearchBuilder {
	rb.v.TimedOut = timedout
	return rb
}

// Took set the Took property for AsyncSearchBuilder.
func (rb *AsyncSearchBuilder) Took(took int64) *AsyncSearchBuilder {
	rb.v.Took = took
	return rb
}
