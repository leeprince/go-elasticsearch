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

// ResponseBody type.
type ResponseBody struct {
	Aggregations    map[AggregateName]Aggregate  `json:"aggregations,omitempty"`
	Clusters_       *ClusterStatistics           `json:"_clusters,omitempty"`
	Fields          map[string]interface{}       `json:"fields,omitempty"`
	Hits            HitsMetadata                 `json:"hits"`
	MaxScore        *float64                     `json:"max_score,omitempty"`
	NumReducePhases *int64                       `json:"num_reduce_phases,omitempty"`
	PitId           *Id                          `json:"pit_id,omitempty"`
	Profile         *Profile                     `json:"profile,omitempty"`
	ScrollId_       *ScrollId                    `json:"_scroll_id,omitempty"`
	Shards_         ShardStatistics              `json:"_shards"`
	Suggest         map[SuggestionName][]Suggest `json:"suggest,omitempty"`
	TerminatedEarly *bool                        `json:"terminated_early,omitempty"`
	TimedOut        bool                         `json:"timed_out"`
	Took            int64                        `json:"took"`
}

// ResponseBodyBuilder holds ResponseBody struct and provides a builder API.
type ResponseBodyBuilder struct {
	v *ResponseBody
}

// NewResponseBody provides a builder for the ResponseBody struct.
func NewResponseBody() *ResponseBodyBuilder {
	r := ResponseBodyBuilder{
		&ResponseBody{
			Aggregations: make(map[AggregateName]Aggregate, 0),
			Fields:       make(map[string]interface{}, 0),
			Suggest:      make(map[SuggestionName][]Suggest, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ResponseBody struct
func (rb *ResponseBodyBuilder) Build() ResponseBody {
	return *rb.v
}

// Aggregations set the Aggregations property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) Aggregations(value map[AggregateName]Aggregate) *ResponseBodyBuilder {
	rb.v.Aggregations = value
	return rb
}

// Clusters_ set the Clusters_ property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) Clusters_(clusters_ ClusterStatistics) *ResponseBodyBuilder {
	rb.v.Clusters_ = &clusters_
	return rb
}

// Fields set the Fields property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) Fields(value map[string]interface{}) *ResponseBodyBuilder {
	rb.v.Fields = value
	return rb
}

// Hits set the Hits property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) Hits(hits HitsMetadata) *ResponseBodyBuilder {
	rb.v.Hits = hits
	return rb
}

// MaxScore set the MaxScore property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) MaxScore(maxscore float64) *ResponseBodyBuilder {
	rb.v.MaxScore = &maxscore
	return rb
}

// NumReducePhases set the NumReducePhases property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) NumReducePhases(numreducephases int64) *ResponseBodyBuilder {
	rb.v.NumReducePhases = &numreducephases
	return rb
}

// PitId set the PitId property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) PitId(pitid Id) *ResponseBodyBuilder {
	rb.v.PitId = &pitid
	return rb
}

// Profile set the Profile property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) Profile(profile Profile) *ResponseBodyBuilder {
	rb.v.Profile = &profile
	return rb
}

// ScrollId_ set the ScrollId_ property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) ScrollId_(scrollid_ ScrollId) *ResponseBodyBuilder {
	rb.v.ScrollId_ = &scrollid_
	return rb
}

// Shards_ set the Shards_ property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) Shards_(shards_ ShardStatistics) *ResponseBodyBuilder {
	rb.v.Shards_ = shards_
	return rb
}

// Suggest set the Suggest property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) Suggest(value map[SuggestionName][]Suggest) *ResponseBodyBuilder {
	rb.v.Suggest = value
	return rb
}

// TerminatedEarly set the TerminatedEarly property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) TerminatedEarly(terminatedearly bool) *ResponseBodyBuilder {
	rb.v.TerminatedEarly = &terminatedearly
	return rb
}

// TimedOut set the TimedOut property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) TimedOut(timedout bool) *ResponseBodyBuilder {
	rb.v.TimedOut = timedout
	return rb
}

// Took set the Took property for ResponseBodyBuilder.
func (rb *ResponseBodyBuilder) Took(took int64) *ResponseBodyBuilder {
	rb.v.Took = took
	return rb
}
