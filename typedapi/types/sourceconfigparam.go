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

// SourceConfigParam holds the union for the following types:
//     bool
//     Fields
type SourceConfigParam interface{}

// SourceConfigParamBuilder holds SourceConfigParam struct and provides a builder API.
type SourceConfigParamBuilder struct {
	v SourceConfigParam
}

// NewSourceConfigParam provides a builder for the SourceConfigParam struct.
func NewSourceConfigParam() *SourceConfigParamBuilder {
	return &SourceConfigParamBuilder{}
}

// Build finalize the chain and returns the SourceConfigParam struct
func (u *SourceConfigParamBuilder) Build() SourceConfigParam {
	return u.v
}

// Bool set the Bool property for SourceConfigParamBuilder.
func (u *SourceConfigParamBuilder) Bool(v bool) *SourceConfigParamBuilder {
	u.v = v
	return u
}

// Fields set the Fields property for SourceConfigParamBuilder.
func (u *SourceConfigParamBuilder) Fields(v Fields) *SourceConfigParamBuilder {
	u.v = v
	return u
}
