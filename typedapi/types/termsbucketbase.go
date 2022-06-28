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

// TermsBucketBase type.
type TermsBucketBase struct {
	Aggregations  map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount      int64                       `json:"doc_count"`
	DocCountError *int64                      `json:"doc_count_error,omitempty"`
}

// TermsBucketBaseBuilder holds TermsBucketBase struct and provides a builder API.
type TermsBucketBaseBuilder struct {
	v *TermsBucketBase
}

// NewTermsBucketBase provides a builder for the TermsBucketBase struct.
func NewTermsBucketBase() *TermsBucketBaseBuilder {
	r := TermsBucketBaseBuilder{
		&TermsBucketBase{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TermsBucketBase struct
func (rb *TermsBucketBaseBuilder) Build() TermsBucketBase {
	return *rb.v
}

// Aggregations set the Aggregations property for TermsBucketBaseBuilder.
func (rb *TermsBucketBaseBuilder) Aggregations(value map[AggregateName]Aggregate) *TermsBucketBaseBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for TermsBucketBaseBuilder.
func (rb *TermsBucketBaseBuilder) DocCount(doccount int64) *TermsBucketBaseBuilder {
	rb.v.DocCount = doccount
	return rb
}

// DocCountError set the DocCountError property for TermsBucketBaseBuilder.
func (rb *TermsBucketBaseBuilder) DocCountError(doccounterror int64) *TermsBucketBaseBuilder {
	rb.v.DocCountError = &doccounterror
	return rb
}
