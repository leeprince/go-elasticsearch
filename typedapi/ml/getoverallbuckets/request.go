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


package getoverallbuckets

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package getoverallbuckets
type Request struct {

	// AllowNoMatch Refer to the description for the `allow_no_match` query parameter.
	AllowNoMatch *bool `json:"allow_no_match,omitempty"`

	// BucketSpan Refer to the description for the `bucket_span` query parameter.
	BucketSpan *types.Time `json:"bucket_span,omitempty"`

	// End Refer to the description for the `end` query parameter.
	End *types.Time `json:"end,omitempty"`

	// ExcludeInterim Refer to the description for the `exclude_interim` query parameter.
	ExcludeInterim *bool `json:"exclude_interim,omitempty"`

	// OverallScore Refer to the description for the `overall_score` query parameter.
	OverallScore string `json:"overall_score,omitempty"`

	// Start Refer to the description for the `start` query parameter.
	Start *types.Time `json:"start,omitempty"`

	// TopN Refer to the description for the `top_n` query parameter.
	TopN *int `json:"top_n,omitempty"`
}

// RequestBuilder is the builder API for the getoverallbuckets.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequest() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Getoverallbuckets request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) AllowNoMatch(allownomatch bool) *RequestBuilder {
	rb.v.AllowNoMatch = &allownomatch
	return rb
}

func (rb *RequestBuilder) BucketSpan(bucketspan types.Time) *RequestBuilder {
	rb.v.BucketSpan = &bucketspan
	return rb
}

func (rb *RequestBuilder) End(end types.Time) *RequestBuilder {
	rb.v.End = &end
	return rb
}

func (rb *RequestBuilder) ExcludeInterim(excludeinterim bool) *RequestBuilder {
	rb.v.ExcludeInterim = &excludeinterim
	return rb
}

func (rb *RequestBuilder) OverallScore(arg string) *RequestBuilder {
	rb.v.OverallScore = arg
	return rb
}

func (rb *RequestBuilder) Start(start types.Time) *RequestBuilder {
	rb.v.Start = &start
	return rb
}

func (rb *RequestBuilder) TopN(topn int) *RequestBuilder {
	rb.v.TopN = &topn
	return rb
}
