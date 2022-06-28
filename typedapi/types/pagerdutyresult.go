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

// PagerDutyResult type.
type PagerDutyResult struct {
	Event    PagerDutyEvent           `json:"event"`
	Reason   *string                  `json:"reason,omitempty"`
	Request  *HttpInputRequestResult  `json:"request,omitempty"`
	Response *HttpInputResponseResult `json:"response,omitempty"`
}

// PagerDutyResultBuilder holds PagerDutyResult struct and provides a builder API.
type PagerDutyResultBuilder struct {
	v *PagerDutyResult
}

// NewPagerDutyResult provides a builder for the PagerDutyResult struct.
func NewPagerDutyResult() *PagerDutyResultBuilder {
	r := PagerDutyResultBuilder{
		&PagerDutyResult{},
	}

	return &r
}

// Build finalize the chain and returns the PagerDutyResult struct
func (rb *PagerDutyResultBuilder) Build() PagerDutyResult {
	return *rb.v
}

// Event set the Event property for PagerDutyResultBuilder.
func (rb *PagerDutyResultBuilder) Event(event PagerDutyEvent) *PagerDutyResultBuilder {
	rb.v.Event = event
	return rb
}

// Reason set the Reason property for PagerDutyResultBuilder.
func (rb *PagerDutyResultBuilder) Reason(reason string) *PagerDutyResultBuilder {
	rb.v.Reason = &reason
	return rb
}

// Request set the Request property for PagerDutyResultBuilder.
func (rb *PagerDutyResultBuilder) Request(request HttpInputRequestResult) *PagerDutyResultBuilder {
	rb.v.Request = &request
	return rb
}

// Response set the Response property for PagerDutyResultBuilder.
func (rb *PagerDutyResultBuilder) Response(response HttpInputResponseResult) *PagerDutyResultBuilder {
	rb.v.Response = &response
	return rb
}
