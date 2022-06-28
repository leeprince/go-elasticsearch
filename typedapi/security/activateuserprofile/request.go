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


package activateuserprofile

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/granttype"
)

// Request holds the request body struct for the package activateuserprofile
type Request struct {
	AccessToken *string `json:"access_token,omitempty"`

	GrantType granttype.GrantType `json:"grant_type"`

	Password *string `json:"password,omitempty"`

	Username *string `json:"username,omitempty"`
}

// RequestBuilder is the builder API for the activateuserprofile.Request
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
		return nil, fmt.Errorf("could not deserialise json into Activateuserprofile request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) AccessToken(accesstoken string) *RequestBuilder {
	rb.v.AccessToken = &accesstoken
	return rb
}

func (rb *RequestBuilder) GrantType(granttype granttype.GrantType) *RequestBuilder {
	rb.v.GrantType = granttype
	return rb
}

func (rb *RequestBuilder) Password(password string) *RequestBuilder {
	rb.v.Password = &password
	return rb
}

func (rb *RequestBuilder) Username(username string) *RequestBuilder {
	rb.v.Username = &username
	return rb
}
