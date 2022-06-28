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


package putwatch

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putwatch
type Request struct {
	Actions map[string]types.Action `json:"actions,omitempty"`

	Condition *types.ConditionContainer `json:"condition,omitempty"`

	Input *types.InputContainer `json:"input,omitempty"`

	Metadata *types.Metadata `json:"metadata,omitempty"`

	ThrottlePeriod *string `json:"throttle_period,omitempty"`

	Transform *types.TransformContainer `json:"transform,omitempty"`

	Trigger *types.TriggerContainer `json:"trigger,omitempty"`
}

// RequestBuilder is the builder API for the putwatch.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequest() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Actions: make(map[string]types.Action, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putwatch request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Actions(value map[string]types.Action) *RequestBuilder {
	rb.v.Actions = value
	return rb
}

func (rb *RequestBuilder) Condition(condition types.ConditionContainer) *RequestBuilder {
	rb.v.Condition = &condition
	return rb
}

func (rb *RequestBuilder) Input(input types.InputContainer) *RequestBuilder {
	rb.v.Input = &input
	return rb
}

func (rb *RequestBuilder) Metadata(metadata types.Metadata) *RequestBuilder {
	rb.v.Metadata = &metadata
	return rb
}

func (rb *RequestBuilder) ThrottlePeriod(throttleperiod string) *RequestBuilder {
	rb.v.ThrottlePeriod = &throttleperiod
	return rb
}

func (rb *RequestBuilder) Transform(transform types.TransformContainer) *RequestBuilder {
	rb.v.Transform = &transform
	return rb
}

func (rb *RequestBuilder) Trigger(trigger types.TriggerContainer) *RequestBuilder {
	rb.v.Trigger = &trigger
	return rb
}
