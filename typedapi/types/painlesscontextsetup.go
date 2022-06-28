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

// PainlessContextSetup type.
type PainlessContextSetup struct {
	Document interface{}    `json:"document,omitempty"`
	Index    IndexName      `json:"index"`
	Query    QueryContainer `json:"query"`
}

// PainlessContextSetupBuilder holds PainlessContextSetup struct and provides a builder API.
type PainlessContextSetupBuilder struct {
	v *PainlessContextSetup
}

// NewPainlessContextSetup provides a builder for the PainlessContextSetup struct.
func NewPainlessContextSetup() *PainlessContextSetupBuilder {
	r := PainlessContextSetupBuilder{
		&PainlessContextSetup{},
	}

	return &r
}

// Build finalize the chain and returns the PainlessContextSetup struct
func (rb *PainlessContextSetupBuilder) Build() PainlessContextSetup {
	return *rb.v
}

// Document set the Document property for PainlessContextSetupBuilder.
func (rb *PainlessContextSetupBuilder) Document(document interface{}) *PainlessContextSetupBuilder {
	rb.v.Document = document
	return rb
}

// Index set the Index property for PainlessContextSetupBuilder.
func (rb *PainlessContextSetupBuilder) Index(index IndexName) *PainlessContextSetupBuilder {
	rb.v.Index = index
	return rb
}

// Query set the Query property for PainlessContextSetupBuilder.
func (rb *PainlessContextSetupBuilder) Query(query QueryContainer) *PainlessContextSetupBuilder {
	rb.v.Query = query
	return rb
}
