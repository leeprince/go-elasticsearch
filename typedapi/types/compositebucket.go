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

// CompositeBucket type.
type CompositeBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Key          map[string]interface{}      `json:"key"`
}

// CompositeBucketBuilder holds CompositeBucket struct and provides a builder API.
type CompositeBucketBuilder struct {
	v *CompositeBucket
}

// NewCompositeBucket provides a builder for the CompositeBucket struct.
func NewCompositeBucket() *CompositeBucketBuilder {
	r := CompositeBucketBuilder{
		&CompositeBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
			Key:          make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the CompositeBucket struct
func (rb *CompositeBucketBuilder) Build() CompositeBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for CompositeBucketBuilder.
func (rb *CompositeBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *CompositeBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for CompositeBucketBuilder.
func (rb *CompositeBucketBuilder) DocCount(doccount int64) *CompositeBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Key set the Key property for CompositeBucketBuilder.
func (rb *CompositeBucketBuilder) Key(value map[string]interface{}) *CompositeBucketBuilder {
	rb.v.Key = value
	return rb
}
