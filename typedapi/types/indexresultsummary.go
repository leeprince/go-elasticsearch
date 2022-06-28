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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/result"
)

// IndexResultSummary type.
type IndexResultSummary struct {
	Created bool          `json:"created"`
	Id      Id            `json:"id"`
	Index   IndexName     `json:"index"`
	Result  result.Result `json:"result"`
	Version VersionNumber `json:"version"`
}

// IndexResultSummaryBuilder holds IndexResultSummary struct and provides a builder API.
type IndexResultSummaryBuilder struct {
	v *IndexResultSummary
}

// NewIndexResultSummary provides a builder for the IndexResultSummary struct.
func NewIndexResultSummary() *IndexResultSummaryBuilder {
	r := IndexResultSummaryBuilder{
		&IndexResultSummary{},
	}

	return &r
}

// Build finalize the chain and returns the IndexResultSummary struct
func (rb *IndexResultSummaryBuilder) Build() IndexResultSummary {
	return *rb.v
}

// Created set the Created property for IndexResultSummaryBuilder.
func (rb *IndexResultSummaryBuilder) Created(created bool) *IndexResultSummaryBuilder {
	rb.v.Created = created
	return rb
}

// Id set the Id property for IndexResultSummaryBuilder.
func (rb *IndexResultSummaryBuilder) Id(id Id) *IndexResultSummaryBuilder {
	rb.v.Id = id
	return rb
}

// Index set the Index property for IndexResultSummaryBuilder.
func (rb *IndexResultSummaryBuilder) Index(index IndexName) *IndexResultSummaryBuilder {
	rb.v.Index = index
	return rb
}

// Result set the Result property for IndexResultSummaryBuilder.
func (rb *IndexResultSummaryBuilder) Result(result result.Result) *IndexResultSummaryBuilder {
	rb.v.Result = result
	return rb
}

// Version set the Version property for IndexResultSummaryBuilder.
func (rb *IndexResultSummaryBuilder) Version(version VersionNumber) *IndexResultSummaryBuilder {
	rb.v.Version = version
	return rb
}
