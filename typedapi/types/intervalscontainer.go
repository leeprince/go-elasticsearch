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

// IntervalsContainer type.
type IntervalsContainer struct {
	AllOf    *IntervalsAllOf    `json:"all_of,omitempty"`
	AnyOf    *IntervalsAnyOf    `json:"any_of,omitempty"`
	Fuzzy    *IntervalsFuzzy    `json:"fuzzy,omitempty"`
	Match    *IntervalsMatch    `json:"match,omitempty"`
	Prefix   *IntervalsPrefix   `json:"prefix,omitempty"`
	Wildcard *IntervalsWildcard `json:"wildcard,omitempty"`
}

// IntervalsContainerBuilder holds IntervalsContainer struct and provides a builder API.
type IntervalsContainerBuilder struct {
	v *IntervalsContainer
}

// NewIntervalsContainer provides a builder for the IntervalsContainer struct.
func NewIntervalsContainer() *IntervalsContainerBuilder {
	r := IntervalsContainerBuilder{
		&IntervalsContainer{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsContainer struct
func (rb *IntervalsContainerBuilder) Build() IntervalsContainer {
	return *rb.v
}

// AllOf set the AllOf property for IntervalsContainerBuilder.
func (rb *IntervalsContainerBuilder) AllOf(allof IntervalsAllOf) *IntervalsContainerBuilder {
	rb.v.AllOf = &allof
	return rb
}

// AnyOf set the AnyOf property for IntervalsContainerBuilder.
func (rb *IntervalsContainerBuilder) AnyOf(anyof IntervalsAnyOf) *IntervalsContainerBuilder {
	rb.v.AnyOf = &anyof
	return rb
}

// Fuzzy set the Fuzzy property for IntervalsContainerBuilder.
func (rb *IntervalsContainerBuilder) Fuzzy(fuzzy IntervalsFuzzy) *IntervalsContainerBuilder {
	rb.v.Fuzzy = &fuzzy
	return rb
}

// Match set the Match property for IntervalsContainerBuilder.
func (rb *IntervalsContainerBuilder) Match(match IntervalsMatch) *IntervalsContainerBuilder {
	rb.v.Match = &match
	return rb
}

// Prefix set the Prefix property for IntervalsContainerBuilder.
func (rb *IntervalsContainerBuilder) Prefix(prefix IntervalsPrefix) *IntervalsContainerBuilder {
	rb.v.Prefix = &prefix
	return rb
}

// Wildcard set the Wildcard property for IntervalsContainerBuilder.
func (rb *IntervalsContainerBuilder) Wildcard(wildcard IntervalsWildcard) *IntervalsContainerBuilder {
	rb.v.Wildcard = &wildcard
	return rb
}
