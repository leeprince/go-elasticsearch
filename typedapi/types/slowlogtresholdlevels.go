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

// SlowlogTresholdLevels type.
type SlowlogTresholdLevels struct {
	Debug *Time `json:"debug,omitempty"`
	Info  *Time `json:"info,omitempty"`
	Trace *Time `json:"trace,omitempty"`
	Warn  *Time `json:"warn,omitempty"`
}

// SlowlogTresholdLevelsBuilder holds SlowlogTresholdLevels struct and provides a builder API.
type SlowlogTresholdLevelsBuilder struct {
	v *SlowlogTresholdLevels
}

// NewSlowlogTresholdLevels provides a builder for the SlowlogTresholdLevels struct.
func NewSlowlogTresholdLevels() *SlowlogTresholdLevelsBuilder {
	r := SlowlogTresholdLevelsBuilder{
		&SlowlogTresholdLevels{},
	}

	return &r
}

// Build finalize the chain and returns the SlowlogTresholdLevels struct
func (rb *SlowlogTresholdLevelsBuilder) Build() SlowlogTresholdLevels {
	return *rb.v
}

// Debug set the Debug property for SlowlogTresholdLevelsBuilder.
func (rb *SlowlogTresholdLevelsBuilder) Debug(debug Time) *SlowlogTresholdLevelsBuilder {
	rb.v.Debug = &debug
	return rb
}

// Info set the Info property for SlowlogTresholdLevelsBuilder.
func (rb *SlowlogTresholdLevelsBuilder) Info(info Time) *SlowlogTresholdLevelsBuilder {
	rb.v.Info = &info
	return rb
}

// Trace set the Trace property for SlowlogTresholdLevelsBuilder.
func (rb *SlowlogTresholdLevelsBuilder) Trace(trace Time) *SlowlogTresholdLevelsBuilder {
	rb.v.Trace = &trace
	return rb
}

// Warn set the Warn property for SlowlogTresholdLevelsBuilder.
func (rb *SlowlogTresholdLevelsBuilder) Warn(warn Time) *SlowlogTresholdLevelsBuilder {
	rb.v.Warn = &warn
	return rb
}
