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

// StatusResponseBase type.
type StatusResponseBase struct {
	CompletionStatus       *int            `json:"completion_status,omitempty"`
	ExpirationTimeInMillis EpochMillis     `json:"expiration_time_in_millis"`
	Id                     *Id             `json:"id,omitempty"`
	IsPartial              bool            `json:"is_partial"`
	IsRunning              bool            `json:"is_running"`
	Shards_                ShardStatistics `json:"_shards"`
	StartTimeInMillis      EpochMillis     `json:"start_time_in_millis"`
}

// StatusResponseBaseBuilder holds StatusResponseBase struct and provides a builder API.
type StatusResponseBaseBuilder struct {
	v *StatusResponseBase
}

// NewStatusResponseBase provides a builder for the StatusResponseBase struct.
func NewStatusResponseBase() *StatusResponseBaseBuilder {
	r := StatusResponseBaseBuilder{
		&StatusResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the StatusResponseBase struct
func (rb *StatusResponseBaseBuilder) Build() StatusResponseBase {
	return *rb.v
}

// CompletionStatus set the CompletionStatus property for StatusResponseBaseBuilder.
func (rb *StatusResponseBaseBuilder) CompletionStatus(completionstatus int) *StatusResponseBaseBuilder {
	rb.v.CompletionStatus = &completionstatus
	return rb
}

// ExpirationTimeInMillis set the ExpirationTimeInMillis property for StatusResponseBaseBuilder.
func (rb *StatusResponseBaseBuilder) ExpirationTimeInMillis(expirationtimeinmillis EpochMillis) *StatusResponseBaseBuilder {
	rb.v.ExpirationTimeInMillis = expirationtimeinmillis
	return rb
}

// Id set the Id property for StatusResponseBaseBuilder.
func (rb *StatusResponseBaseBuilder) Id(id Id) *StatusResponseBaseBuilder {
	rb.v.Id = &id
	return rb
}

// IsPartial set the IsPartial property for StatusResponseBaseBuilder.
func (rb *StatusResponseBaseBuilder) IsPartial(ispartial bool) *StatusResponseBaseBuilder {
	rb.v.IsPartial = ispartial
	return rb
}

// IsRunning set the IsRunning property for StatusResponseBaseBuilder.
func (rb *StatusResponseBaseBuilder) IsRunning(isrunning bool) *StatusResponseBaseBuilder {
	rb.v.IsRunning = isrunning
	return rb
}

// Shards_ set the Shards_ property for StatusResponseBaseBuilder.
func (rb *StatusResponseBaseBuilder) Shards_(shards_ ShardStatistics) *StatusResponseBaseBuilder {
	rb.v.Shards_ = shards_
	return rb
}

// StartTimeInMillis set the StartTimeInMillis property for StatusResponseBaseBuilder.
func (rb *StatusResponseBaseBuilder) StartTimeInMillis(starttimeinmillis EpochMillis) *StatusResponseBaseBuilder {
	rb.v.StartTimeInMillis = starttimeinmillis
	return rb
}
