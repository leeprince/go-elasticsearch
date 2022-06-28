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

// FiltersBucket type.
type FiltersBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
}

// FiltersBucketBuilder holds FiltersBucket struct and provides a builder API.
type FiltersBucketBuilder struct {
	v *FiltersBucket
}

// NewFiltersBucket provides a builder for the FiltersBucket struct.
func NewFiltersBucket() *FiltersBucketBuilder {
	r := FiltersBucketBuilder{
		&FiltersBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the FiltersBucket struct
func (rb *FiltersBucketBuilder) Build() FiltersBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for FiltersBucketBuilder.
func (rb *FiltersBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *FiltersBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for FiltersBucketBuilder.
func (rb *FiltersBucketBuilder) DocCount(doccount int64) *FiltersBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}
