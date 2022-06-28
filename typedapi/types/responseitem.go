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

// ResponseItem holds the union for the following types:
//     GetResult
//     MultiGetError
type ResponseItem interface{}

// ResponseItemBuilder holds ResponseItem struct and provides a builder API.
type ResponseItemBuilder struct {
	v ResponseItem
}

// NewResponseItem provides a builder for the ResponseItem struct.
func NewResponseItem() *ResponseItemBuilder {
	return &ResponseItemBuilder{}
}

// Build finalize the chain and returns the ResponseItem struct
func (u *ResponseItemBuilder) Build() ResponseItem {
	return u.v
}

// GetResult set the GetResult property for ResponseItemBuilder.
func (u *ResponseItemBuilder) GetResult(v GetResult) *ResponseItemBuilder {
	u.v = v
	return u
}

// MultiGetError set the MultiGetError property for ResponseItemBuilder.
func (u *ResponseItemBuilder) MultiGetError(v MultiGetError) *ResponseItemBuilder {
	u.v = v
	return u
}
