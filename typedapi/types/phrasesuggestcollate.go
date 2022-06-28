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

// PhraseSuggestCollate type.
type PhraseSuggestCollate struct {
	Params map[string]interface{}    `json:"params,omitempty"`
	Prune  *bool                     `json:"prune,omitempty"`
	Query  PhraseSuggestCollateQuery `json:"query"`
}

// PhraseSuggestCollateBuilder holds PhraseSuggestCollate struct and provides a builder API.
type PhraseSuggestCollateBuilder struct {
	v *PhraseSuggestCollate
}

// NewPhraseSuggestCollate provides a builder for the PhraseSuggestCollate struct.
func NewPhraseSuggestCollate() *PhraseSuggestCollateBuilder {
	r := PhraseSuggestCollateBuilder{
		&PhraseSuggestCollate{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the PhraseSuggestCollate struct
func (rb *PhraseSuggestCollateBuilder) Build() PhraseSuggestCollate {
	return *rb.v
}

// Params set the Params property for PhraseSuggestCollateBuilder.
func (rb *PhraseSuggestCollateBuilder) Params(value map[string]interface{}) *PhraseSuggestCollateBuilder {
	rb.v.Params = value
	return rb
}

// Prune set the Prune property for PhraseSuggestCollateBuilder.
func (rb *PhraseSuggestCollateBuilder) Prune(prune bool) *PhraseSuggestCollateBuilder {
	rb.v.Prune = &prune
	return rb
}

// Query set the Query property for PhraseSuggestCollateBuilder.
func (rb *PhraseSuggestCollateBuilder) Query(query PhraseSuggestCollateQuery) *PhraseSuggestCollateBuilder {
	rb.v.Query = query
	return rb
}
