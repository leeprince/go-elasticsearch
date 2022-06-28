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

// SlicedScroll type.
type SlicedScroll struct {
	Field *Field `json:"field,omitempty"`
	Id    int    `json:"id"`
	Max   int    `json:"max"`
}

// SlicedScrollBuilder holds SlicedScroll struct and provides a builder API.
type SlicedScrollBuilder struct {
	v *SlicedScroll
}

// NewSlicedScroll provides a builder for the SlicedScroll struct.
func NewSlicedScroll() *SlicedScrollBuilder {
	r := SlicedScrollBuilder{
		&SlicedScroll{},
	}

	return &r
}

// Build finalize the chain and returns the SlicedScroll struct
func (rb *SlicedScrollBuilder) Build() SlicedScroll {
	return *rb.v
}

// Field set the Field property for SlicedScrollBuilder.
func (rb *SlicedScrollBuilder) Field(field Field) *SlicedScrollBuilder {
	rb.v.Field = &field
	return rb
}

// Id set the Id property for SlicedScrollBuilder.
func (rb *SlicedScrollBuilder) Id(id int) *SlicedScrollBuilder {
	rb.v.Id = id
	return rb
}

// Max set the Max property for SlicedScrollBuilder.
func (rb *SlicedScrollBuilder) Max(max int) *SlicedScrollBuilder {
	rb.v.Max = max
	return rb
}
