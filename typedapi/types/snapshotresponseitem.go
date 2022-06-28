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

// SnapshotResponseItem type.
type SnapshotResponseItem struct {
	Error      *ErrorCause    `json:"error,omitempty"`
	Repository Name           `json:"repository"`
	Snapshots  []SnapshotInfo `json:"snapshots,omitempty"`
}

// SnapshotResponseItemBuilder holds SnapshotResponseItem struct and provides a builder API.
type SnapshotResponseItemBuilder struct {
	v *SnapshotResponseItem
}

// NewSnapshotResponseItem provides a builder for the SnapshotResponseItem struct.
func NewSnapshotResponseItem() *SnapshotResponseItemBuilder {
	r := SnapshotResponseItemBuilder{
		&SnapshotResponseItem{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotResponseItem struct
func (rb *SnapshotResponseItemBuilder) Build() SnapshotResponseItem {
	return *rb.v
}

// Error set the Error property for SnapshotResponseItemBuilder.
func (rb *SnapshotResponseItemBuilder) Error(error ErrorCause) *SnapshotResponseItemBuilder {
	rb.v.Error = &error
	return rb
}

// Repository set the Repository property for SnapshotResponseItemBuilder.
func (rb *SnapshotResponseItemBuilder) Repository(repository Name) *SnapshotResponseItemBuilder {
	rb.v.Repository = repository
	return rb
}

// Snapshots set the Snapshots property for SnapshotResponseItemBuilder.
func (rb *SnapshotResponseItemBuilder) Snapshots(snapshots ...SnapshotInfo) *SnapshotResponseItemBuilder {
	rb.v.Snapshots = append(rb.v.Snapshots, snapshots...)
	return rb
}
