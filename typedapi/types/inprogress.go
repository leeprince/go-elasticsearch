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

// InProgress type.
type InProgress struct {
	Name            Name       `json:"name"`
	StartTimeMillis DateString `json:"start_time_millis"`
	State           string     `json:"state"`
	Uuid            Uuid       `json:"uuid"`
}

// InProgressBuilder holds InProgress struct and provides a builder API.
type InProgressBuilder struct {
	v *InProgress
}

// NewInProgress provides a builder for the InProgress struct.
func NewInProgress() *InProgressBuilder {
	r := InProgressBuilder{
		&InProgress{},
	}

	return &r
}

// Build finalize the chain and returns the InProgress struct
func (rb *InProgressBuilder) Build() InProgress {
	return *rb.v
}

// Name set the Name property for InProgressBuilder.
func (rb *InProgressBuilder) Name(name Name) *InProgressBuilder {
	rb.v.Name = name
	return rb
}

// StartTimeMillis set the StartTimeMillis property for InProgressBuilder.
func (rb *InProgressBuilder) StartTimeMillis(starttimemillis DateString) *InProgressBuilder {
	rb.v.StartTimeMillis = starttimemillis
	return rb
}

// State set the State property for InProgressBuilder.
func (rb *InProgressBuilder) State(state string) *InProgressBuilder {
	rb.v.State = state
	return rb
}

// Uuid set the Uuid property for InProgressBuilder.
func (rb *InProgressBuilder) Uuid(uuid Uuid) *InProgressBuilder {
	rb.v.Uuid = uuid
	return rb
}
