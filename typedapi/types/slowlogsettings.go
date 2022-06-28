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

// SlowlogSettings type.
type SlowlogSettings struct {
	Level     *string           `json:"level,omitempty"`
	Reformat  *bool             `json:"reformat,omitempty"`
	Source    *int              `json:"source,omitempty"`
	Threshold *SlowlogTresholds `json:"threshold,omitempty"`
}

// SlowlogSettingsBuilder holds SlowlogSettings struct and provides a builder API.
type SlowlogSettingsBuilder struct {
	v *SlowlogSettings
}

// NewSlowlogSettings provides a builder for the SlowlogSettings struct.
func NewSlowlogSettings() *SlowlogSettingsBuilder {
	r := SlowlogSettingsBuilder{
		&SlowlogSettings{},
	}

	return &r
}

// Build finalize the chain and returns the SlowlogSettings struct
func (rb *SlowlogSettingsBuilder) Build() SlowlogSettings {
	return *rb.v
}

// Level set the Level property for SlowlogSettingsBuilder.
func (rb *SlowlogSettingsBuilder) Level(level string) *SlowlogSettingsBuilder {
	rb.v.Level = &level
	return rb
}

// Reformat set the Reformat property for SlowlogSettingsBuilder.
func (rb *SlowlogSettingsBuilder) Reformat(reformat bool) *SlowlogSettingsBuilder {
	rb.v.Reformat = &reformat
	return rb
}

// Source set the Source property for SlowlogSettingsBuilder.
func (rb *SlowlogSettingsBuilder) Source(source int) *SlowlogSettingsBuilder {
	rb.v.Source = &source
	return rb
}

// Threshold set the Threshold property for SlowlogSettingsBuilder.
func (rb *SlowlogSettingsBuilder) Threshold(threshold SlowlogTresholds) *SlowlogSettingsBuilder {
	rb.v.Threshold = &threshold
	return rb
}
