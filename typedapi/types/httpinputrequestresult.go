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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/connectionscheme"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/httpinputmethod"
)

// HttpInputRequestResult type.
type HttpInputRequestResult struct {
	Auth              *HttpInputAuthentication           `json:"auth,omitempty"`
	Body              *string                            `json:"body,omitempty"`
	ConnectionTimeout *Time                              `json:"connection_timeout,omitempty"`
	Headers           map[string]string                  `json:"headers,omitempty"`
	Host              *Host                              `json:"host,omitempty"`
	Method            *httpinputmethod.HttpInputMethod   `json:"method,omitempty"`
	Params            map[string]string                  `json:"params,omitempty"`
	Path              *string                            `json:"path,omitempty"`
	Port              *uint                              `json:"port,omitempty"`
	Proxy             *HttpInputProxy                    `json:"proxy,omitempty"`
	ReadTimeout       *Time                              `json:"read_timeout,omitempty"`
	Scheme            *connectionscheme.ConnectionScheme `json:"scheme,omitempty"`
	Url               *string                            `json:"url,omitempty"`
}

// HttpInputRequestResultBuilder holds HttpInputRequestResult struct and provides a builder API.
type HttpInputRequestResultBuilder struct {
	v *HttpInputRequestResult
}

// NewHttpInputRequestResult provides a builder for the HttpInputRequestResult struct.
func NewHttpInputRequestResult() *HttpInputRequestResultBuilder {
	r := HttpInputRequestResultBuilder{
		&HttpInputRequestResult{
			Headers: make(map[string]string, 0),
			Params:  make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the HttpInputRequestResult struct
func (rb *HttpInputRequestResultBuilder) Build() HttpInputRequestResult {
	return *rb.v
}

// Auth set the Auth property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Auth(auth HttpInputAuthentication) *HttpInputRequestResultBuilder {
	rb.v.Auth = &auth
	return rb
}

// Body set the Body property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Body(body string) *HttpInputRequestResultBuilder {
	rb.v.Body = &body
	return rb
}

// ConnectionTimeout set the ConnectionTimeout property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) ConnectionTimeout(connectiontimeout Time) *HttpInputRequestResultBuilder {
	rb.v.ConnectionTimeout = &connectiontimeout
	return rb
}

// Headers set the Headers property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Headers(value map[string]string) *HttpInputRequestResultBuilder {
	rb.v.Headers = value
	return rb
}

// Host set the Host property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Host(host Host) *HttpInputRequestResultBuilder {
	rb.v.Host = &host
	return rb
}

// Method set the Method property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Method(method httpinputmethod.HttpInputMethod) *HttpInputRequestResultBuilder {
	rb.v.Method = &method
	return rb
}

// Params set the Params property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Params(value map[string]string) *HttpInputRequestResultBuilder {
	rb.v.Params = value
	return rb
}

// Path set the Path property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Path(path string) *HttpInputRequestResultBuilder {
	rb.v.Path = &path
	return rb
}

// Port set the Port property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Port(port uint) *HttpInputRequestResultBuilder {
	rb.v.Port = &port
	return rb
}

// Proxy set the Proxy property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Proxy(proxy HttpInputProxy) *HttpInputRequestResultBuilder {
	rb.v.Proxy = &proxy
	return rb
}

// ReadTimeout set the ReadTimeout property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) ReadTimeout(readtimeout Time) *HttpInputRequestResultBuilder {
	rb.v.ReadTimeout = &readtimeout
	return rb
}

// Scheme set the Scheme property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Scheme(scheme connectionscheme.ConnectionScheme) *HttpInputRequestResultBuilder {
	rb.v.Scheme = &scheme
	return rb
}

// Url set the Url property for HttpInputRequestResultBuilder.
func (rb *HttpInputRequestResultBuilder) Url(url string) *HttpInputRequestResultBuilder {
	rb.v.Url = &url
	return rb
}
