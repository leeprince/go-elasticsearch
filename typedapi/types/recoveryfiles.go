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

// RecoveryFiles type.
type RecoveryFiles struct {
	Details   []FileDetails `json:"details,omitempty"`
	Percent   Percentage    `json:"percent"`
	Recovered int64         `json:"recovered"`
	Reused    int64         `json:"reused"`
	Total     int64         `json:"total"`
}

// RecoveryFilesBuilder holds RecoveryFiles struct and provides a builder API.
type RecoveryFilesBuilder struct {
	v *RecoveryFiles
}

// NewRecoveryFiles provides a builder for the RecoveryFiles struct.
func NewRecoveryFiles() *RecoveryFilesBuilder {
	r := RecoveryFilesBuilder{
		&RecoveryFiles{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryFiles struct
func (rb *RecoveryFilesBuilder) Build() RecoveryFiles {
	return *rb.v
}

// Details set the Details property for RecoveryFilesBuilder.
func (rb *RecoveryFilesBuilder) Details(details ...FileDetails) *RecoveryFilesBuilder {
	rb.v.Details = append(rb.v.Details, details...)
	return rb
}

// Percent set the Percent property for RecoveryFilesBuilder.
func (rb *RecoveryFilesBuilder) Percent(percent Percentage) *RecoveryFilesBuilder {
	rb.v.Percent = percent
	return rb
}

// Recovered set the Recovered property for RecoveryFilesBuilder.
func (rb *RecoveryFilesBuilder) Recovered(recovered int64) *RecoveryFilesBuilder {
	rb.v.Recovered = recovered
	return rb
}

// Reused set the Reused property for RecoveryFilesBuilder.
func (rb *RecoveryFilesBuilder) Reused(reused int64) *RecoveryFilesBuilder {
	rb.v.Reused = reused
	return rb
}

// Total set the Total property for RecoveryFilesBuilder.
func (rb *RecoveryFilesBuilder) Total(total int64) *RecoveryFilesBuilder {
	rb.v.Total = total
	return rb
}
