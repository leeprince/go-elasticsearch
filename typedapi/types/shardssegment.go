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

// ShardsSegment type.
type ShardsSegment struct {
	NumCommittedSegments int                 `json:"num_committed_segments"`
	NumSearchSegments    int                 `json:"num_search_segments"`
	Routing              ShardSegmentRouting `json:"routing"`
	Segments             map[string]Segment  `json:"segments"`
}

// ShardsSegmentBuilder holds ShardsSegment struct and provides a builder API.
type ShardsSegmentBuilder struct {
	v *ShardsSegment
}

// NewShardsSegment provides a builder for the ShardsSegment struct.
func NewShardsSegment() *ShardsSegmentBuilder {
	r := ShardsSegmentBuilder{
		&ShardsSegment{
			Segments: make(map[string]Segment, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ShardsSegment struct
func (rb *ShardsSegmentBuilder) Build() ShardsSegment {
	return *rb.v
}

// NumCommittedSegments set the NumCommittedSegments property for ShardsSegmentBuilder.
func (rb *ShardsSegmentBuilder) NumCommittedSegments(numcommittedsegments int) *ShardsSegmentBuilder {
	rb.v.NumCommittedSegments = numcommittedsegments
	return rb
}

// NumSearchSegments set the NumSearchSegments property for ShardsSegmentBuilder.
func (rb *ShardsSegmentBuilder) NumSearchSegments(numsearchsegments int) *ShardsSegmentBuilder {
	rb.v.NumSearchSegments = numsearchsegments
	return rb
}

// Routing set the Routing property for ShardsSegmentBuilder.
func (rb *ShardsSegmentBuilder) Routing(routing ShardSegmentRouting) *ShardsSegmentBuilder {
	rb.v.Routing = routing
	return rb
}

// Segments set the Segments property for ShardsSegmentBuilder.
func (rb *ShardsSegmentBuilder) Segments(value map[string]Segment) *ShardsSegmentBuilder {
	rb.v.Segments = value
	return rb
}
