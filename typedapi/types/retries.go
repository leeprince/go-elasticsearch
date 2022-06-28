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

// Retries type.
type Retries struct {
	Bulk   int64 `json:"bulk"`
	Search int64 `json:"search"`
}

// RetriesBuilder holds Retries struct and provides a builder API.
type RetriesBuilder struct {
	v *Retries
}

// NewRetries provides a builder for the Retries struct.
func NewRetries() *RetriesBuilder {
	r := RetriesBuilder{
		&Retries{},
	}

	return &r
}

// Build finalize the chain and returns the Retries struct
func (rb *RetriesBuilder) Build() Retries {
	return *rb.v
}

// Bulk set the Bulk property for RetriesBuilder.
func (rb *RetriesBuilder) Bulk(bulk int64) *RetriesBuilder {
	rb.v.Bulk = bulk
	return rb
}

// Search set the Search property for RetriesBuilder.
func (rb *RetriesBuilder) Search(search int64) *RetriesBuilder {
	rb.v.Search = search
	return rb
}
