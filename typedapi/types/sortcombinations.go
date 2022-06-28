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

// SortCombinations holds the union for the following types:
//     Field
//     SortOptions
type SortCombinations interface{}

// SortCombinationsBuilder holds SortCombinations struct and provides a builder API.
type SortCombinationsBuilder struct {
	v SortCombinations
}

// NewSortCombinations provides a builder for the SortCombinations struct.
func NewSortCombinations() *SortCombinationsBuilder {
	return &SortCombinationsBuilder{}
}

// Build finalize the chain and returns the SortCombinations struct
func (u *SortCombinationsBuilder) Build() SortCombinations {
	return u.v
}

// Field set the Field property for SortCombinationsBuilder.
func (u *SortCombinationsBuilder) Field(v Field) *SortCombinationsBuilder {
	u.v = v
	return u
}

// SortOptions set the SortOptions property for SortCombinationsBuilder.
func (u *SortCombinationsBuilder) SortOptions(v SortOptions) *SortCombinationsBuilder {
	u.v = v
	return u
}
