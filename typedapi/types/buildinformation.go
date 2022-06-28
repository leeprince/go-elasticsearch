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

// BuildInformation type.
type BuildInformation struct {
	Date DateString `json:"date"`
	Hash string     `json:"hash"`
}

// BuildInformationBuilder holds BuildInformation struct and provides a builder API.
type BuildInformationBuilder struct {
	v *BuildInformation
}

// NewBuildInformation provides a builder for the BuildInformation struct.
func NewBuildInformation() *BuildInformationBuilder {
	r := BuildInformationBuilder{
		&BuildInformation{},
	}

	return &r
}

// Build finalize the chain and returns the BuildInformation struct
func (rb *BuildInformationBuilder) Build() BuildInformation {
	return *rb.v
}

// Date set the Date property for BuildInformationBuilder.
func (rb *BuildInformationBuilder) Date(date DateString) *BuildInformationBuilder {
	rb.v.Date = date
	return rb
}

// Hash set the Hash property for BuildInformationBuilder.
func (rb *BuildInformationBuilder) Hash(hash string) *BuildInformationBuilder {
	rb.v.Hash = hash
	return rb
}
