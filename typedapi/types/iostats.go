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

// IoStats type.
type IoStats struct {
	Devices []IoStatDevice `json:"devices,omitempty"`
	Total   *IoStatDevice  `json:"total,omitempty"`
}

// IoStatsBuilder holds IoStats struct and provides a builder API.
type IoStatsBuilder struct {
	v *IoStats
}

// NewIoStats provides a builder for the IoStats struct.
func NewIoStats() *IoStatsBuilder {
	r := IoStatsBuilder{
		&IoStats{},
	}

	return &r
}

// Build finalize the chain and returns the IoStats struct
func (rb *IoStatsBuilder) Build() IoStats {
	return *rb.v
}

// Devices set the Devices property for IoStatsBuilder.
func (rb *IoStatsBuilder) Devices(devices ...IoStatDevice) *IoStatsBuilder {
	rb.v.Devices = append(rb.v.Devices, devices...)
	return rb
}

// Total set the Total property for IoStatsBuilder.
func (rb *IoStatsBuilder) Total(total IoStatDevice) *IoStatsBuilder {
	rb.v.Total = &total
	return rb
}
