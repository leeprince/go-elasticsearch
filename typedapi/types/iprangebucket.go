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

// IpRangeBucket type.
type IpRangeBucket struct {
	Aggregations map[AggregateName]Aggregate `json:"aggregations,omitempty"`
	DocCount     int64                       `json:"doc_count"`
	From         *string                     `json:"from,omitempty"`
	To           *string                     `json:"to,omitempty"`
}

// IpRangeBucketBuilder holds IpRangeBucket struct and provides a builder API.
type IpRangeBucketBuilder struct {
	v *IpRangeBucket
}

// NewIpRangeBucket provides a builder for the IpRangeBucket struct.
func NewIpRangeBucket() *IpRangeBucketBuilder {
	r := IpRangeBucketBuilder{
		&IpRangeBucket{
			Aggregations: make(map[AggregateName]Aggregate, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the IpRangeBucket struct
func (rb *IpRangeBucketBuilder) Build() IpRangeBucket {
	return *rb.v
}

// Aggregations set the Aggregations property for IpRangeBucketBuilder.
func (rb *IpRangeBucketBuilder) Aggregations(value map[AggregateName]Aggregate) *IpRangeBucketBuilder {
	rb.v.Aggregations = value
	return rb
}

// DocCount set the DocCount property for IpRangeBucketBuilder.
func (rb *IpRangeBucketBuilder) DocCount(doccount int64) *IpRangeBucketBuilder {
	rb.v.DocCount = doccount
	return rb
}

// From set the From property for IpRangeBucketBuilder.
func (rb *IpRangeBucketBuilder) From(from string) *IpRangeBucketBuilder {
	rb.v.From = &from
	return rb
}

// To set the To property for IpRangeBucketBuilder.
func (rb *IpRangeBucketBuilder) To(to string) *IpRangeBucketBuilder {
	rb.v.To = &to
	return rb
}
