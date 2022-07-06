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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deprecationlevel"
)

// Deprecation type.
type Deprecation struct {
	Details string `json:"details"`
	// Level The level property describes the significance of the issue.
	Level   deprecationlevel.DeprecationLevel `json:"level"`
	Message string                            `json:"message"`
	Url     string                            `json:"url"`
}

// DeprecationBuilder holds Deprecation struct and provides a builder API.
type DeprecationBuilder struct {
	v *Deprecation
}

// NewDeprecation provides a builder for the Deprecation struct.
func NewDeprecation() *DeprecationBuilder {
	r := DeprecationBuilder{
		&Deprecation{},
	}

	return &r
}

// Build finalize the chain and returns the Deprecation struct
func (rb *DeprecationBuilder) Build() Deprecation {
	return *rb.v
}

// Details set the Details property for DeprecationBuilder.
func (rb *DeprecationBuilder) Details(details string) *DeprecationBuilder {
	rb.v.Details = details
	return rb
}

// Level The level property describes the significance of the issue.
func (rb *DeprecationBuilder) Level(level deprecationlevel.DeprecationLevel) *DeprecationBuilder {
	rb.v.Level = level
	return rb
}

// Message set the Message property for DeprecationBuilder.
func (rb *DeprecationBuilder) Message(message string) *DeprecationBuilder {
	rb.v.Message = message
	return rb
}

// Url set the Url property for DeprecationBuilder.
func (rb *DeprecationBuilder) Url(url string) *DeprecationBuilder {
	rb.v.Url = url
	return rb
}
