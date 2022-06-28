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

// Defaults type.
type Defaults struct {
	AnomalyDetectors AnomalyDetectors `json:"anomaly_detectors"`
	Datafeeds        Datafeeds        `json:"datafeeds"`
}

// DefaultsBuilder holds Defaults struct and provides a builder API.
type DefaultsBuilder struct {
	v *Defaults
}

// NewDefaults provides a builder for the Defaults struct.
func NewDefaults() *DefaultsBuilder {
	r := DefaultsBuilder{
		&Defaults{},
	}

	return &r
}

// Build finalize the chain and returns the Defaults struct
func (rb *DefaultsBuilder) Build() Defaults {
	return *rb.v
}

// AnomalyDetectors set the AnomalyDetectors property for DefaultsBuilder.
func (rb *DefaultsBuilder) AnomalyDetectors(anomalydetectors AnomalyDetectors) *DefaultsBuilder {
	rb.v.AnomalyDetectors = anomalydetectors
	return rb
}

// Datafeeds set the Datafeeds property for DefaultsBuilder.
func (rb *DefaultsBuilder) Datafeeds(datafeeds Datafeeds) *DefaultsBuilder {
	rb.v.Datafeeds = datafeeds
	return rb
}
