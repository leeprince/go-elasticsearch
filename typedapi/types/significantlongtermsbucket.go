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

// SignificantLongTermsBucket type.
type SignificantLongTermsBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	BgCount      int64                       `json:"bg_count"`
	DocCount     int64                       `json:"doc_count"`
	Key          int64                       `json:"key"`
	KeyAsString  *string                     `json:"key_as_string,omitempty"`
	Score        float64                     `json:"score"`
}

// SignificantLongTermsBucketBuilder holds SignificantLongTermsBucket struct and provides a builder API.
type SignificantLongTermsBucketBuilder struct {
	v *SignificantLongTermsBucket
}

// NewSignificantLongTermsBucket provides a builder for the SignificantLongTermsBucket struct.
func NewSignificantLongTermsBucket() *SignificantLongTermsBucketBuilder {
	r := SignificantLongTermsBucketBuilder{
		&SignificantLongTermsBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SignificantLongTermsBucket struct
func (rb *SignificantLongTermsBucketBuilder) Build() SignificantLongTermsBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for SignificantLongTermsBucketBuilder.
func (rb *SignificantLongTermsBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *SignificantLongTermsBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// BgCount set the BgCount property for SignificantLongTermsBucketBuilder.
func (rb *SignificantLongTermsBucketBuilder) BgCount(bgcount int64) *SignificantLongTermsBucketBuilder {
	rb.v.BgCount = bgcount
	return rb
}

// DocCount set the DocCount property for SignificantLongTermsBucketBuilder.
func (rb *SignificantLongTermsBucketBuilder) DocCount(doccount int64) *SignificantLongTermsBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Key set the Key property for SignificantLongTermsBucketBuilder.
func (rb *SignificantLongTermsBucketBuilder) Key(key int64) *SignificantLongTermsBucketBuilder {
	rb.v.Key = key
	return rb
}

// KeyAsString set the KeyAsString property for SignificantLongTermsBucketBuilder.
func (rb *SignificantLongTermsBucketBuilder) KeyAsString(keyasstring string) *SignificantLongTermsBucketBuilder {
	rb.v.KeyAsString = &keyasstring
	return rb
}

// Score set the Score property for SignificantLongTermsBucketBuilder.
func (rb *SignificantLongTermsBucketBuilder) Score(score float64) *SignificantLongTermsBucketBuilder {
	rb.v.Score = score
	return rb
}
