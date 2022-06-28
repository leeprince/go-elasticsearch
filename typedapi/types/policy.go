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

// Policy type.
type Policy struct {
	Config     Configuration  `json:"config"`
	Name       Name           `json:"name"`
	Repository string         `json:"repository"`
	Retention  Retention      `json:"retention"`
	Schedule   CronExpression `json:"schedule"`
}

// PolicyBuilder holds Policy struct and provides a builder API.
type PolicyBuilder struct {
	v *Policy
}

// NewPolicy provides a builder for the Policy struct.
func NewPolicy() *PolicyBuilder {
	r := PolicyBuilder{
		&Policy{},
	}

	return &r
}

// Build finalize the chain and returns the Policy struct
func (rb *PolicyBuilder) Build() Policy {
	return *rb.v
}

// Config set the Config property for PolicyBuilder.
func (rb *PolicyBuilder) Config(config Configuration) *PolicyBuilder {
	rb.v.Config = config
	return rb
}

// Name set the Name property for PolicyBuilder.
func (rb *PolicyBuilder) Name(name Name) *PolicyBuilder {
	rb.v.Name = name
	return rb
}

// Repository set the Repository property for PolicyBuilder.
func (rb *PolicyBuilder) Repository(repository string) *PolicyBuilder {
	rb.v.Repository = repository
	return rb
}

// Retention set the Retention property for PolicyBuilder.
func (rb *PolicyBuilder) Retention(retention Retention) *PolicyBuilder {
	rb.v.Retention = retention
	return rb
}

// Schedule set the Schedule property for PolicyBuilder.
func (rb *PolicyBuilder) Schedule(schedule CronExpression) *PolicyBuilder {
	rb.v.Schedule = schedule
	return rb
}
