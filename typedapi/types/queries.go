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

// Queries type.
type Queries struct {
	Cache *CacheQueries `json:"cache,omitempty"`
}

// QueriesBuilder holds Queries struct and provides a builder API.
type QueriesBuilder struct {
	v *Queries
}

// NewQueries provides a builder for the Queries struct.
func NewQueries() *QueriesBuilder {
	r := QueriesBuilder{
		&Queries{},
	}

	return &r
}

// Build finalize the chain and returns the Queries struct
func (rb *QueriesBuilder) Build() Queries {
	return *rb.v
}

// Cache set the Cache property for QueriesBuilder.
func (rb *QueriesBuilder) Cache(cache CacheQueries) *QueriesBuilder {
	rb.v.Cache = &cache
	return rb
}
