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

// Datafeed type.
type Datafeed struct {
	Count int64 `json:"count"`
}

// DatafeedBuilder holds Datafeed struct and provides a builder API.
type DatafeedBuilder struct {
	v *Datafeed
}

// NewDatafeed provides a builder for the Datafeed struct.
func NewDatafeed() *DatafeedBuilder {
	r := DatafeedBuilder{
		&Datafeed{},
	}

	return &r
}

// Build finalize the chain and returns the Datafeed struct
func (rb *DatafeedBuilder) Build() Datafeed {
	return *rb.v
}

// Count set the Count property for DatafeedBuilder.
func (rb *DatafeedBuilder) Count(count int64) *DatafeedBuilder {
	rb.v.Count = count
	return rb
}
