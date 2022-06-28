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

// Segment type.
type Segment struct {
	Attributes    map[string]string `json:"attributes"`
	Committed     bool              `json:"committed"`
	Compound      bool              `json:"compound"`
	DeletedDocs   int64             `json:"deleted_docs"`
	Generation    int               `json:"generation"`
	MemoryInBytes float64           `json:"memory_in_bytes"`
	NumDocs       int64             `json:"num_docs"`
	Search        bool              `json:"search"`
	SizeInBytes   float64           `json:"size_in_bytes"`
	Version       VersionString     `json:"version"`
}

// SegmentBuilder holds Segment struct and provides a builder API.
type SegmentBuilder struct {
	v *Segment
}

// NewSegment provides a builder for the Segment struct.
func NewSegment() *SegmentBuilder {
	r := SegmentBuilder{
		&Segment{
			Attributes: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Segment struct
func (rb *SegmentBuilder) Build() Segment {
	return *rb.v
}

// Attributes set the Attributes property for SegmentBuilder.
func (rb *SegmentBuilder) Attributes(value map[string]string) *SegmentBuilder {
	rb.v.Attributes = value
	return rb
}

// Committed set the Committed property for SegmentBuilder.
func (rb *SegmentBuilder) Committed(committed bool) *SegmentBuilder {
	rb.v.Committed = committed
	return rb
}

// Compound set the Compound property for SegmentBuilder.
func (rb *SegmentBuilder) Compound(compound bool) *SegmentBuilder {
	rb.v.Compound = compound
	return rb
}

// DeletedDocs set the DeletedDocs property for SegmentBuilder.
func (rb *SegmentBuilder) DeletedDocs(deleteddocs int64) *SegmentBuilder {
	rb.v.DeletedDocs = deleteddocs
	return rb
}

// Generation set the Generation property for SegmentBuilder.
func (rb *SegmentBuilder) Generation(generation int) *SegmentBuilder {
	rb.v.Generation = generation
	return rb
}

// MemoryInBytes set the MemoryInBytes property for SegmentBuilder.
func (rb *SegmentBuilder) MemoryInBytes(memoryinbytes float64) *SegmentBuilder {
	rb.v.MemoryInBytes = memoryinbytes
	return rb
}

// NumDocs set the NumDocs property for SegmentBuilder.
func (rb *SegmentBuilder) NumDocs(numdocs int64) *SegmentBuilder {
	rb.v.NumDocs = numdocs
	return rb
}

// Search set the Search property for SegmentBuilder.
func (rb *SegmentBuilder) Search(search bool) *SegmentBuilder {
	rb.v.Search = search
	return rb
}

// SizeInBytes set the SizeInBytes property for SegmentBuilder.
func (rb *SegmentBuilder) SizeInBytes(sizeinbytes float64) *SegmentBuilder {
	rb.v.SizeInBytes = sizeinbytes
	return rb
}

// Version set the Version property for SegmentBuilder.
func (rb *SegmentBuilder) Version(version VersionString) *SegmentBuilder {
	rb.v.Version = version
	return rb
}
