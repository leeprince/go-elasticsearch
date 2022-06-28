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

// AsyncSearchDocumentResponseBase type.
type AsyncSearchDocumentResponseBase struct {
	ExpirationTimeInMillis EpochMillis `json:"expiration_time_in_millis"`
	Id                     *Id         `json:"id,omitempty"`
	IsPartial              bool        `json:"is_partial"`
	IsRunning              bool        `json:"is_running"`
	Response               AsyncSearch `json:"response"`
	StartTimeInMillis      EpochMillis `json:"start_time_in_millis"`
}

// AsyncSearchDocumentResponseBaseBuilder holds AsyncSearchDocumentResponseBase struct and provides a builder API.
type AsyncSearchDocumentResponseBaseBuilder struct {
	v *AsyncSearchDocumentResponseBase
}

// NewAsyncSearchDocumentResponseBase provides a builder for the AsyncSearchDocumentResponseBase struct.
func NewAsyncSearchDocumentResponseBase() *AsyncSearchDocumentResponseBaseBuilder {
	r := AsyncSearchDocumentResponseBaseBuilder{
		&AsyncSearchDocumentResponseBase{},
	}

	return &r
}

// Build finalize the chain and returns the AsyncSearchDocumentResponseBase struct
func (rb *AsyncSearchDocumentResponseBaseBuilder) Build() AsyncSearchDocumentResponseBase {
	return *rb.v
}

// ExpirationTimeInMillis set the ExpirationTimeInMillis property for AsyncSearchDocumentResponseBaseBuilder.
func (rb *AsyncSearchDocumentResponseBaseBuilder) ExpirationTimeInMillis(expirationtimeinmillis EpochMillis) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.ExpirationTimeInMillis = expirationtimeinmillis
	return rb
}

// Id set the Id property for AsyncSearchDocumentResponseBaseBuilder.
func (rb *AsyncSearchDocumentResponseBaseBuilder) Id(id Id) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.Id = &id
	return rb
}

// IsPartial set the IsPartial property for AsyncSearchDocumentResponseBaseBuilder.
func (rb *AsyncSearchDocumentResponseBaseBuilder) IsPartial(ispartial bool) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.IsPartial = ispartial
	return rb
}

// IsRunning set the IsRunning property for AsyncSearchDocumentResponseBaseBuilder.
func (rb *AsyncSearchDocumentResponseBaseBuilder) IsRunning(isrunning bool) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.IsRunning = isrunning
	return rb
}

// Response set the Response property for AsyncSearchDocumentResponseBaseBuilder.
func (rb *AsyncSearchDocumentResponseBaseBuilder) Response(response AsyncSearch) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.Response = response
	return rb
}

// StartTimeInMillis set the StartTimeInMillis property for AsyncSearchDocumentResponseBaseBuilder.
func (rb *AsyncSearchDocumentResponseBaseBuilder) StartTimeInMillis(starttimeinmillis EpochMillis) *AsyncSearchDocumentResponseBaseBuilder {
	rb.v.StartTimeInMillis = starttimeinmillis
	return rb
}
