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

// TopLeftBottomRightGeoBounds type.
type TopLeftBottomRightGeoBounds struct {
	BottomRight GeoLocation `json:"bottom_right"`
	TopLeft     GeoLocation `json:"top_left"`
}

// TopLeftBottomRightGeoBoundsBuilder holds TopLeftBottomRightGeoBounds struct and provides a builder API.
type TopLeftBottomRightGeoBoundsBuilder struct {
	v *TopLeftBottomRightGeoBounds
}

// NewTopLeftBottomRightGeoBounds provides a builder for the TopLeftBottomRightGeoBounds struct.
func NewTopLeftBottomRightGeoBounds() *TopLeftBottomRightGeoBoundsBuilder {
	r := TopLeftBottomRightGeoBoundsBuilder{
		&TopLeftBottomRightGeoBounds{},
	}

	return &r
}

// Build finalize the chain and returns the TopLeftBottomRightGeoBounds struct
func (rb *TopLeftBottomRightGeoBoundsBuilder) Build() TopLeftBottomRightGeoBounds {
	return *rb.v
}

// BottomRight set the BottomRight property for TopLeftBottomRightGeoBoundsBuilder.
func (rb *TopLeftBottomRightGeoBoundsBuilder) BottomRight(bottomright GeoLocation) *TopLeftBottomRightGeoBoundsBuilder {
	rb.v.BottomRight = bottomright
	return rb
}

// TopLeft set the TopLeft property for TopLeftBottomRightGeoBoundsBuilder.
func (rb *TopLeftBottomRightGeoBoundsBuilder) TopLeft(topleft GeoLocation) *TopLeftBottomRightGeoBoundsBuilder {
	rb.v.TopLeft = topleft
	return rb
}
