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

// SignificantTermsBucketBase type.
type SignificantTermsBucketBase struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	BgCount      int64                       `json:"bg_count"`
	DocCount     int64                       `json:"doc_count"`
	Score        float64                     `json:"score"`
}

// SignificantTermsBucketBaseBuilder holds SignificantTermsBucketBase struct and provides a builder API.
type SignificantTermsBucketBaseBuilder struct {
	v *SignificantTermsBucketBase
}

// NewSignificantTermsBucketBase provides a builder for the SignificantTermsBucketBase struct.
func NewSignificantTermsBucketBase() *SignificantTermsBucketBaseBuilder {
	r := SignificantTermsBucketBaseBuilder{
		&SignificantTermsBucketBase{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SignificantTermsBucketBase struct
func (rb *SignificantTermsBucketBaseBuilder) Build() SignificantTermsBucketBase {
	return *rb.v
}

// Aggregations set the Aggregations property for SignificantTermsBucketBaseBuilder.
func (rb *SignificantTermsBucketBaseBuilder) Aggregations(value map[AggregateName]Aggregate) *SignificantTermsBucketBaseBuilder {
	rb.v.Aggregations = value
	return rb
}

// BgCount set the BgCount property for SignificantTermsBucketBaseBuilder.
func (rb *SignificantTermsBucketBaseBuilder) BgCount(bgcount int64) *SignificantTermsBucketBaseBuilder {
	rb.v.BgCount = bgcount
	return rb
}

// DocCount set the DocCount property for SignificantTermsBucketBaseBuilder.
func (rb *SignificantTermsBucketBaseBuilder) DocCount(doccount int64) *SignificantTermsBucketBaseBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Score set the Score property for SignificantTermsBucketBaseBuilder.
func (rb *SignificantTermsBucketBaseBuilder) Score(score float64) *SignificantTermsBucketBaseBuilder {
	rb.v.Score = score
	return rb
}
