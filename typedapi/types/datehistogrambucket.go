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

// DateHistogramBucket type.
type DateHistogramBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Key          EpochMillis                 `json:"key"`
	KeyAsString  *string                     `json:"key_as_string,omitempty"`
}

// DateHistogramBucketBuilder holds DateHistogramBucket struct and provides a builder API.
type DateHistogramBucketBuilder struct {
	v *DateHistogramBucket
}

// NewDateHistogramBucket provides a builder for the DateHistogramBucket struct.
func NewDateHistogramBucket() *DateHistogramBucketBuilder {
	r := DateHistogramBucketBuilder{
		&DateHistogramBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DateHistogramBucket struct
func (rb *DateHistogramBucketBuilder) Build() DateHistogramBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for DateHistogramBucketBuilder.
func (rb *DateHistogramBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *DateHistogramBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for DateHistogramBucketBuilder.
func (rb *DateHistogramBucketBuilder) DocCount(doccount int64) *DateHistogramBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Key set the Key property for DateHistogramBucketBuilder.
func (rb *DateHistogramBucketBuilder) Key(key EpochMillis) *DateHistogramBucketBuilder {
	rb.v.Key = key
	return rb
}

// KeyAsString set the KeyAsString property for DateHistogramBucketBuilder.
func (rb *DateHistogramBucketBuilder) KeyAsString(keyasstring string) *DateHistogramBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}
