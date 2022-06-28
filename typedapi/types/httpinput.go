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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/responsecontenttype"
)

// HttpInput type.
type HttpInput struct {
	Extract             []string                                 `json:"extract,omitempty"`
	Http                *HttpInput                               `json:"http,omitempty"`
	Request             *HttpInputRequestDefinition              `json:"request,omitempty"`
	ResponseContentType *responsecontenttype.ResponseContentType `json:"response_content_type,omitempty"`
}

// HttpInputBuilder holds HttpInput struct and provides a builder API.
type HttpInputBuilder struct {
	v *HttpInput
}

// NewHttpInput provides a builder for the HttpInput struct.
func NewHttpInput() *HttpInputBuilder {
	r := HttpInputBuilder{
		&HttpInput{},
	}

	return &r
}

// Build finalize the chain and returns the HttpInput struct
func (rb *HttpInputBuilder) Build() HttpInput {
	return *rb.v
}

// Extract set the Extract property for HttpInputBuilder.
func (rb *HttpInputBuilder) Extract(extract ...string) *HttpInputBuilder {
	rb.v.Extract = append(rb.v.Extract, extract...)
	return rb
}

// Http set the Http property for HttpInputBuilder.
func (rb *HttpInputBuilder) Http(http HttpInput) *HttpInputBuilder {
	rb.v.Http = &http
	return rb
}

// Request set the Request property for HttpInputBuilder.
func (rb *HttpInputBuilder) Request(request HttpInputRequestDefinition) *HttpInputBuilder {
	rb.v.Request = &request
	return rb
}

// ResponseContentType set the ResponseContentType property for HttpInputBuilder.
func (rb *HttpInputBuilder) ResponseContentType(responsecontenttype responsecontenttype.ResponseContentType) *HttpInputBuilder {
	rb.v.ResponseContentType = &responsecontenttype
	return rb
}
