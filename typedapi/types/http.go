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

// Http type.
type Http struct {
	Clients     []Client `json:"clients,omitempty"`
	CurrentOpen *int     `json:"current_open,omitempty"`
	TotalOpened *int64   `json:"total_opened,omitempty"`
}

// HttpBuilder holds Http struct and provides a builder API.
type HttpBuilder struct {
	v *Http
}

// NewHttp provides a builder for the Http struct.
func NewHttp() *HttpBuilder {
	r := HttpBuilder{
		&Http{},
	}

	return &r
}

// Build finalize the chain and returns the Http struct
func (rb *HttpBuilder) Build() Http {
	return *rb.v
}

// Clients set the Clients property for HttpBuilder.
func (rb *HttpBuilder) Clients(clients ...Client) *HttpBuilder {
	rb.v.Clients = append(rb.v.Clients, clients...)
	return rb
}

// CurrentOpen set the CurrentOpen property for HttpBuilder.
func (rb *HttpBuilder) CurrentOpen(currentopen int) *HttpBuilder {
	rb.v.CurrentOpen = &currentopen
	return rb
}

// TotalOpened set the TotalOpened property for HttpBuilder.
func (rb *HttpBuilder) TotalOpened(totalopened int64) *HttpBuilder {
	rb.v.TotalOpened = &totalopened
	return rb
}
