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

// QueryProfile type.
type QueryProfile struct {
	Breakdown   QueryBreakdown `json:"breakdown"`
	Children    []QueryProfile `json:"children,omitempty"`
	Description string         `json:"description"`
	TimeInNanos int64          `json:"time_in_nanos"`
	Type        string         `json:"type"`
}

// QueryProfileBuilder holds QueryProfile struct and provides a builder API.
type QueryProfileBuilder struct {
	v *QueryProfile
}

// NewQueryProfile provides a builder for the QueryProfile struct.
func NewQueryProfile() *QueryProfileBuilder {
	r := QueryProfileBuilder{
		&QueryProfile{},
	}

	return &r
}

// Build finalize the chain and returns the QueryProfile struct
func (rb *QueryProfileBuilder) Build() QueryProfile {
	return *rb.v
}

// Breakdown set the Breakdown property for QueryProfileBuilder.
func (rb *QueryProfileBuilder) Breakdown(breakdown QueryBreakdown) *QueryProfileBuilder {
	rb.v.Breakdown = breakdown
	return rb
}

// Children set the Children property for QueryProfileBuilder.
func (rb *QueryProfileBuilder) Children(children ...QueryProfile) *QueryProfileBuilder {
	rb.v.Children = append(rb.v.Children, children...)
	return rb
}

// Description set the Description property for QueryProfileBuilder.
func (rb *QueryProfileBuilder) Description(description string) *QueryProfileBuilder {
	rb.v.Description = description
	return rb
}

// TimeInNanos set the TimeInNanos property for QueryProfileBuilder.
func (rb *QueryProfileBuilder) TimeInNanos(timeinnanos int64) *QueryProfileBuilder {
	rb.v.TimeInNanos = timeinnanos
	return rb
}

// Type set the Type property for QueryProfileBuilder.
func (rb *QueryProfileBuilder) Type_(type_ string) *QueryProfileBuilder {
	rb.v.Type = type_
	return rb
}
