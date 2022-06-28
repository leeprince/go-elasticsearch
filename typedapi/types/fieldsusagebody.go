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

// FieldsUsageBody type.
type FieldsUsageBody struct {
	FieldsUsageBody map[IndexName]UsageStatsIndex `json:"FieldsUsageBody,omitempty"`
	Shards_         ShardStatistics               `json:"_shards"`
}

// FieldsUsageBodyBuilder holds FieldsUsageBody struct and provides a builder API.
type FieldsUsageBodyBuilder struct {
	v *FieldsUsageBody
}

// NewFieldsUsageBody provides a builder for the FieldsUsageBody struct.
func NewFieldsUsageBody() *FieldsUsageBodyBuilder {
	r := FieldsUsageBodyBuilder{
		&FieldsUsageBody{
			FieldsUsageBody: make(map[IndexName]UsageStatsIndex, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the FieldsUsageBody struct
func (rb *FieldsUsageBodyBuilder) Build() FieldsUsageBody {
	return *rb.v
}

// FieldsUsageBody set the FieldsUsageBody property for FieldsUsageBodyBuilder.
func (rb *FieldsUsageBodyBuilder) FieldsUsageBody(value map[IndexName]UsageStatsIndex) *FieldsUsageBodyBuilder {
	rb.v.FieldsUsageBody = value
	return rb
}

// Shards_ set the Shards_ property for FieldsUsageBodyBuilder.
func (rb *FieldsUsageBodyBuilder) Shards_(shards_ ShardStatistics) *FieldsUsageBodyBuilder {
	rb.v.Shards_ = shards_
	return rb
}
