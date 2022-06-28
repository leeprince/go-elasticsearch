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

// IntervalsFilter type.
type IntervalsFilter struct {
	After          *IntervalsContainer `json:"after,omitempty"`
	Before         *IntervalsContainer `json:"before,omitempty"`
	ContainedBy    *IntervalsContainer `json:"contained_by,omitempty"`
	Containing     *IntervalsContainer `json:"containing,omitempty"`
	NotContainedBy *IntervalsContainer `json:"not_contained_by,omitempty"`
	NotContaining  *IntervalsContainer `json:"not_containing,omitempty"`
	NotOverlapping *IntervalsContainer `json:"not_overlapping,omitempty"`
	Overlapping    *IntervalsContainer `json:"overlapping,omitempty"`
	Script         *Script             `json:"script,omitempty"`
}

// IntervalsFilterBuilder holds IntervalsFilter struct and provides a builder API.
type IntervalsFilterBuilder struct {
	v *IntervalsFilter
}

// NewIntervalsFilter provides a builder for the IntervalsFilter struct.
func NewIntervalsFilter() *IntervalsFilterBuilder {
	r := IntervalsFilterBuilder{
		&IntervalsFilter{},
	}

	return &r
}

// Build finalize the chain and returns the IntervalsFilter struct
func (rb *IntervalsFilterBuilder) Build() IntervalsFilter {
	return *rb.v
}

// After set the After property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) After(after IntervalsContainer) *IntervalsFilterBuilder {
	rb.v.After = &after
	return rb
}

// Before set the Before property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) Before(before IntervalsContainer) *IntervalsFilterBuilder {
	rb.v.Before = &before
	return rb
}

// ContainedBy set the ContainedBy property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) ContainedBy(containedby IntervalsContainer) *IntervalsFilterBuilder {
	rb.v.ContainedBy = &containedby
	return rb
}

// Containing set the Containing property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) Containing(containing IntervalsContainer) *IntervalsFilterBuilder {
	rb.v.Containing = &containing
	return rb
}

// NotContainedBy set the NotContainedBy property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) NotContainedBy(notcontainedby IntervalsContainer) *IntervalsFilterBuilder {
	rb.v.NotContainedBy = &notcontainedby
	return rb
}

// NotContaining set the NotContaining property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) NotContaining(notcontaining IntervalsContainer) *IntervalsFilterBuilder {
	rb.v.NotContaining = &notcontaining
	return rb
}

// NotOverlapping set the NotOverlapping property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) NotOverlapping(notoverlapping IntervalsContainer) *IntervalsFilterBuilder {
	rb.v.NotOverlapping = &notoverlapping
	return rb
}

// Overlapping set the Overlapping property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) Overlapping(overlapping IntervalsContainer) *IntervalsFilterBuilder {
	rb.v.Overlapping = &overlapping
	return rb
}

// Script set the Script property for IntervalsFilterBuilder.
func (rb *IntervalsFilterBuilder) Script(script Script) *IntervalsFilterBuilder {
	rb.v.Script = &script
	return rb
}
