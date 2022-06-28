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

// ReindexStatus type.
type ReindexStatus struct {
	Batches              int64   `json:"batches"`
	Created              int64   `json:"created"`
	Deleted              int64   `json:"deleted"`
	Noops                int64   `json:"noops"`
	RequestsPerSecond    float32 `json:"requests_per_second"`
	Retries              Retries `json:"retries"`
	ThrottledMillis      int64   `json:"throttled_millis"`
	ThrottledUntilMillis int64   `json:"throttled_until_millis"`
	Total                int64   `json:"total"`
	Updated              int64   `json:"updated"`
	VersionConflicts     int64   `json:"version_conflicts"`
}

// ReindexStatusBuilder holds ReindexStatus struct and provides a builder API.
type ReindexStatusBuilder struct {
	v *ReindexStatus
}

// NewReindexStatus provides a builder for the ReindexStatus struct.
func NewReindexStatus() *ReindexStatusBuilder {
	r := ReindexStatusBuilder{
		&ReindexStatus{},
	}

	return &r
}

// Build finalize the chain and returns the ReindexStatus struct
func (rb *ReindexStatusBuilder) Build() ReindexStatus {
	return *rb.v
}

// Batches set the Batches property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) Batches(batches int64) *ReindexStatusBuilder {
	rb.v.Batches = batches
	return rb
}

// Created set the Created property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) Created(created int64) *ReindexStatusBuilder {
	rb.v.Created = created
	return rb
}

// Deleted set the Deleted property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) Deleted(deleted int64) *ReindexStatusBuilder {
	rb.v.Deleted = deleted
	return rb
}

// Noops set the Noops property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) Noops(noops int64) *ReindexStatusBuilder {
	rb.v.Noops = noops
	return rb
}

// RequestsPerSecond set the RequestsPerSecond property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) RequestsPerSecond(requestspersecond float32) *ReindexStatusBuilder {
	rb.v.RequestsPerSecond = requestspersecond
	return rb
}

// Retries set the Retries property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) Retries(retries Retries) *ReindexStatusBuilder {
	rb.v.Retries = retries
	return rb
}

// ThrottledMillis set the ThrottledMillis property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) ThrottledMillis(throttledmillis int64) *ReindexStatusBuilder {
	rb.v.ThrottledMillis = throttledmillis
	return rb
}

// ThrottledUntilMillis set the ThrottledUntilMillis property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) ThrottledUntilMillis(throttleduntilmillis int64) *ReindexStatusBuilder {
	rb.v.ThrottledUntilMillis = throttleduntilmillis
	return rb
}

// Total set the Total property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) Total(total int64) *ReindexStatusBuilder {
	rb.v.Total = total
	return rb
}

// Updated set the Updated property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) Updated(updated int64) *ReindexStatusBuilder {
	rb.v.Updated = updated
	return rb
}

// VersionConflicts set the VersionConflicts property for ReindexStatusBuilder.
func (rb *ReindexStatusBuilder) VersionConflicts(versionconflicts int64) *ReindexStatusBuilder {
	rb.v.VersionConflicts = versionconflicts
	return rb
}
