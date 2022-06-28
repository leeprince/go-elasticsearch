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

// SingleBucketAggregateBase type.
type SingleBucketAggregateBase struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	Meta         *Metadata                   `json:"meta,omitempty"`
}

// SingleBucketAggregateBaseBuilder holds SingleBucketAggregateBase struct and provides a builder API.
type SingleBucketAggregateBaseBuilder struct {
	v *SingleBucketAggregateBase
}

// NewSingleBucketAggregateBase provides a builder for the SingleBucketAggregateBase struct.
func NewSingleBucketAggregateBase() *SingleBucketAggregateBaseBuilder {
	r := SingleBucketAggregateBaseBuilder{
		&SingleBucketAggregateBase{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SingleBucketAggregateBase struct
func (rb *SingleBucketAggregateBaseBuilder) Build() SingleBucketAggregateBase {
	return *rb.v
}

// Aggregations set the Aggregations property for SingleBucketAggregateBaseBuilder.
func (rb *SingleBucketAggregateBaseBuilder) Aggregations(value map[AggregateName]Aggregate) *SingleBucketAggregateBaseBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for SingleBucketAggregateBaseBuilder.
func (rb *SingleBucketAggregateBaseBuilder) DocCount(doccount int64) *SingleBucketAggregateBaseBuilder {
	rb.v.DocCount = doccount
	return rb
}

// Meta set the Meta property for SingleBucketAggregateBaseBuilder.
func (rb *SingleBucketAggregateBaseBuilder) Meta(meta Metadata) *SingleBucketAggregateBaseBuilder {
	rb.v.Meta = &meta
	return rb
}
