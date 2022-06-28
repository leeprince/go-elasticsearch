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

// Configurations type.
type Configurations struct {
	Forcemerge *ForceMergeConfiguration `json:"forcemerge,omitempty"`
	Rollover   *RolloverConditions      `json:"rollover,omitempty"`
	Shrink     *ShrinkConfiguration     `json:"shrink,omitempty"`
}

// ConfigurationsBuilder holds Configurations struct and provides a builder API.
type ConfigurationsBuilder struct {
	v *Configurations
}

// NewConfigurations provides a builder for the Configurations struct.
func NewConfigurations() *ConfigurationsBuilder {
	r := ConfigurationsBuilder{
		&Configurations{},
	}

	return &r
}

// Build finalize the chain and returns the Configurations struct
func (rb *ConfigurationsBuilder) Build() Configurations {
	return *rb.v
}

// Forcemerge set the Forcemerge property for ConfigurationsBuilder.
func (rb *ConfigurationsBuilder) Forcemerge(forcemerge ForceMergeConfiguration) *ConfigurationsBuilder {
	rb.v.Forcemerge = &forcemerge
	return rb
}

// Rollover set the Rollover property for ConfigurationsBuilder.
func (rb *ConfigurationsBuilder) Rollover(rollover RolloverConditions) *ConfigurationsBuilder {
	rb.v.Rollover = &rollover
	return rb
}

// Shrink set the Shrink property for ConfigurationsBuilder.
func (rb *ConfigurationsBuilder) Shrink(shrink ShrinkConfiguration) *ConfigurationsBuilder {
	rb.v.Shrink = &shrink
	return rb
}
