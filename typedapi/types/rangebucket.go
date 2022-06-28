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

// RangeBucket type.
type RangeBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	From         *float64                    `json:"from,omitempty"`
	FromAsString *string                     `json:"from_as_string,omitempty"`
	// Key The bucket key. Present if the aggregation is _not_ keyed
	Key        *string  `json:"key,omitempty"`
	To         *float64 `json:"to,omitempty"`
	ToAsString *string  `json:"to_as_string,omitempty"`
}

// RangeBucketBuilder holds RangeBucket struct and provides a builder API.
type RangeBucketBuilder struct {
	v *RangeBucket
}

// NewRangeBucket provides a builder for the RangeBucket struct.
func NewRangeBucket() *RangeBucketBuilder {
	r := RangeBucketBuilder{
		&RangeBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RangeBucket struct
func (rb *RangeBucketBuilder) Build() RangeBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for RangeBucketBuilder.
func (rb *RangeBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *RangeBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for RangeBucketBuilder.
func (rb *RangeBucketBuilder) DocCount(doccount int64) *RangeBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

// From set the From property for RangeBucketBuilder.
func (rb *RangeBucketBuilder) From(from float64) *RangeBucketBuilder {
	rb.v.From = &from
	return rb
}

// FromAsString set the FromAsString property for RangeBucketBuilder.
func (rb *RangeBucketBuilder) FromAsString(fromasstring string) *RangeBucketBuilder {
	rb.v.FromAsString = &fromasstring
	return rb
}

// Key The bucket key. Present if the aggregation is _not_ keyed
func (rb *RangeBucketBuilder) Key(key string) *RangeBucketBuilder {
	rb.v.Key = &key
	return rb
}

// To set the To property for RangeBucketBuilder.
func (rb *RangeBucketBuilder) To(to float64) *RangeBucketBuilder {
	rb.v.To = &to
	return rb
}

// ToAsString set the ToAsString property for RangeBucketBuilder.
func (rb *RangeBucketBuilder) ToAsString(toasstring string) *RangeBucketBuilder {
	rb.v.ToAsString = &toasstring
	return rb
}
