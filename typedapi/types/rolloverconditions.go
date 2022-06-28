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

// RolloverConditions type.
type RolloverConditions struct {
	MaxAge                   *Time        `json:"max_age,omitempty"`
	MaxAgeMillis             *EpochMillis `json:"max_age_millis,omitempty"`
	MaxDocs                  *int64       `json:"max_docs,omitempty"`
	MaxPrimaryShardSize      *ByteSize    `json:"max_primary_shard_size,omitempty"`
	MaxPrimaryShardSizeBytes *ByteSize    `json:"max_primary_shard_size_bytes,omitempty"`
	MaxSize                  *string      `json:"max_size,omitempty"`
	MaxSizeBytes             *ByteSize    `json:"max_size_bytes,omitempty"`
}

// RolloverConditionsBuilder holds RolloverConditions struct and provides a builder API.
type RolloverConditionsBuilder struct {
	v *RolloverConditions
}

// NewRolloverConditions provides a builder for the RolloverConditions struct.
func NewRolloverConditions() *RolloverConditionsBuilder {
	r := RolloverConditionsBuilder{
		&RolloverConditions{},
	}

	return &r
}

// Build finalize the chain and returns the RolloverConditions struct
func (rb *RolloverConditionsBuilder) Build() RolloverConditions {
	return *rb.v
}

// MaxAge set the MaxAge property for RolloverConditionsBuilder.
func (rb *RolloverConditionsBuilder) MaxAge(maxage Time) *RolloverConditionsBuilder {
	rb.v.MaxAge = &maxage
	return rb
}

// MaxAgeMillis set the MaxAgeMillis property for RolloverConditionsBuilder.
func (rb *RolloverConditionsBuilder) MaxAgeMillis(maxagemillis EpochMillis) *RolloverConditionsBuilder {
	rb.v.MaxAgeMillis = &maxagemillis
	return rb
}

// MaxDocs set the MaxDocs property for RolloverConditionsBuilder.
func (rb *RolloverConditionsBuilder) MaxDocs(maxdocs int64) *RolloverConditionsBuilder {
	rb.v.MaxDocs = &maxdocs
	return rb
}

// MaxPrimaryShardSize set the MaxPrimaryShardSize property for RolloverConditionsBuilder.
func (rb *RolloverConditionsBuilder) MaxPrimaryShardSize(maxprimaryshardsize ByteSize) *RolloverConditionsBuilder {
	rb.v.MaxPrimaryShardSize = &maxprimaryshardsize
	return rb
}

// MaxPrimaryShardSizeBytes set the MaxPrimaryShardSizeBytes property for RolloverConditionsBuilder.
func (rb *RolloverConditionsBuilder) MaxPrimaryShardSizeBytes(maxprimaryshardsizebytes ByteSize) *RolloverConditionsBuilder {
	rb.v.MaxPrimaryShardSizeBytes = &maxprimaryshardsizebytes
	return rb
}

// MaxSize set the MaxSize property for RolloverConditionsBuilder.
func (rb *RolloverConditionsBuilder) MaxSize(maxsize string) *RolloverConditionsBuilder {
	rb.v.MaxSize = &maxsize
	return rb
}

// MaxSizeBytes set the MaxSizeBytes property for RolloverConditionsBuilder.
func (rb *RolloverConditionsBuilder) MaxSizeBytes(maxsizebytes ByteSize) *RolloverConditionsBuilder {
	rb.v.MaxSizeBytes = &maxsizebytes
	return rb
}
