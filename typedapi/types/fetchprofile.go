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

// FetchProfile type.
type FetchProfile struct {
	Breakdown   FetchProfileBreakdown `json:"breakdown"`
	Children    []FetchProfile        `json:"children,omitempty"`
	Debug       *FetchProfileDebug    `json:"debug,omitempty"`
	Description string                `json:"description"`
	TimeInNanos int64                 `json:"time_in_nanos"`
	Type        string                `json:"type"`
}

// FetchProfileBuilder holds FetchProfile struct and provides a builder API.
type FetchProfileBuilder struct {
	v *FetchProfile
}

// NewFetchProfile provides a builder for the FetchProfile struct.
func NewFetchProfile() *FetchProfileBuilder {
	r := FetchProfileBuilder{
		&FetchProfile{},
	}

	return &r
}

// Build finalize the chain and returns the FetchProfile struct
func (rb *FetchProfileBuilder) Build() FetchProfile {
	return *rb.v
}

// Breakdown set the Breakdown property for FetchProfileBuilder.
func (rb *FetchProfileBuilder) Breakdown(breakdown FetchProfileBreakdown) *FetchProfileBuilder {
	rb.v.Breakdown = breakdown
	return rb
}

// Children set the Children property for FetchProfileBuilder.
func (rb *FetchProfileBuilder) Children(children ...FetchProfile) *FetchProfileBuilder {
	rb.v.Children = append(rb.v.Children, children...)
	return rb
}

// Debug set the Debug property for FetchProfileBuilder.
func (rb *FetchProfileBuilder) Debug(debug FetchProfileDebug) *FetchProfileBuilder {
	rb.v.Debug = &debug
	return rb
}

// Description set the Description property for FetchProfileBuilder.
func (rb *FetchProfileBuilder) Description(description string) *FetchProfileBuilder {
	rb.v.Description = description
	return rb
}

// TimeInNanos set the TimeInNanos property for FetchProfileBuilder.
func (rb *FetchProfileBuilder) TimeInNanos(timeinnanos int64) *FetchProfileBuilder {
	rb.v.TimeInNanos = timeinnanos
	return rb
}

// Type set the Type property for FetchProfileBuilder.
func (rb *FetchProfileBuilder) Type_(type_ string) *FetchProfileBuilder {
	rb.v.Type = type_
	return rb
}
