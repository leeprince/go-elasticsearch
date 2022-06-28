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

// AllJobs type.
type AllJobs struct {
	Count     int               `json:"count"`
	CreatedBy map[string]string `json:"created_by"`
	Detectors map[string]int    `json:"detectors"`
	Forecasts map[string]int    `json:"forecasts"`
	ModelSize map[string]int    `json:"model_size"`
}

// AllJobsBuilder holds AllJobs struct and provides a builder API.
type AllJobsBuilder struct {
	v *AllJobs
}

// NewAllJobs provides a builder for the AllJobs struct.
func NewAllJobs() *AllJobsBuilder {
	r := AllJobsBuilder{
		&AllJobs{
			CreatedBy: make(map[string]string, 0),
			Detectors: make(map[string]int, 0),
			Forecasts: make(map[string]int, 0),
			ModelSize: make(map[string]int, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the AllJobs struct
func (rb *AllJobsBuilder) Build() AllJobs {
	return *rb.v
}

// Count set the Count property for AllJobsBuilder.
func (rb *AllJobsBuilder) Count(count int) *AllJobsBuilder {
	rb.v.Count = count
	return rb
}

// CreatedBy set the CreatedBy property for AllJobsBuilder.
//TODO: Union in map Not implemented yet: UnionOf
func (rb *AllJobsBuilder) CreatedBy(value map[string]string) *AllJobsBuilder {
	rb.v.CreatedBy = value
	return rb
}

// Detectors set the Detectors property for AllJobsBuilder.
func (rb *AllJobsBuilder) Detectors(value map[string]int) *AllJobsBuilder {
	rb.v.Detectors = value
	return rb
}

// Forecasts set the Forecasts property for AllJobsBuilder.
func (rb *AllJobsBuilder) Forecasts(value map[string]int) *AllJobsBuilder {
	rb.v.Forecasts = value
	return rb
}

// ModelSize set the ModelSize property for AllJobsBuilder.
func (rb *AllJobsBuilder) ModelSize(value map[string]int) *AllJobsBuilder {
	rb.v.ModelSize = value
	return rb
}
