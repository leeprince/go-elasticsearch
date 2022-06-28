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

// StringTermsBucket type.
type StringTermsBucket struct {
	Aggregations  map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount      int64                       `json:"doc_count"`
	DocCountError *int64                      `json:"doc_count_error,omitempty"`
	Key           string                      `json:"key"`
}

// StringTermsBucketBuilder holds StringTermsBucket struct and provides a builder API.
type StringTermsBucketBuilder struct {
	v *StringTermsBucket
}

// NewStringTermsBucket provides a builder for the StringTermsBucket struct.
func NewStringTermsBucket() *StringTermsBucketBuilder {
	r := StringTermsBucketBuilder{
		&StringTermsBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the StringTermsBucket struct
func (rb *StringTermsBucketBuilder) Build() StringTermsBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for StringTermsBucketBuilder.
func (rb *StringTermsBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *StringTermsBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for StringTermsBucketBuilder.
func (rb *StringTermsBucketBuilder) DocCount(doccount int64) *StringTermsBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

// DocCountError set the DocCountError property for StringTermsBucketBuilder.
func (rb *StringTermsBucketBuilder) DocCountError(doccounterror int64) *StringTermsBucketBuilder {
	rb.v.DocCountError = &doccounterror
	return rb
}

// Key set the Key property for StringTermsBucketBuilder.
func (rb *StringTermsBucketBuilder) Key(key string) *StringTermsBucketBuilder {
	rb.v.Key = key
	return rb
}
