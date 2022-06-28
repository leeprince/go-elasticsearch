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

// MigrationFeatureIndexInfo type.
type MigrationFeatureIndexInfo struct {
	FailureCause *ErrorCause   `json:"failure_cause,omitempty"`
	Index        IndexName     `json:"index"`
	Version      VersionString `json:"version"`
}

// MigrationFeatureIndexInfoBuilder holds MigrationFeatureIndexInfo struct and provides a builder API.
type MigrationFeatureIndexInfoBuilder struct {
	v *MigrationFeatureIndexInfo
}

// NewMigrationFeatureIndexInfo provides a builder for the MigrationFeatureIndexInfo struct.
func NewMigrationFeatureIndexInfo() *MigrationFeatureIndexInfoBuilder {
	r := MigrationFeatureIndexInfoBuilder{
		&MigrationFeatureIndexInfo{},
	}

	return &r
}

// Build finalize the chain and returns the MigrationFeatureIndexInfo struct
func (rb *MigrationFeatureIndexInfoBuilder) Build() MigrationFeatureIndexInfo {
	return *rb.v
}

// FailureCause set the FailureCause property for MigrationFeatureIndexInfoBuilder.
func (rb *MigrationFeatureIndexInfoBuilder) FailureCause(failurecause ErrorCause) *MigrationFeatureIndexInfoBuilder {
	rb.v.FailureCause = &failurecause
	return rb
}

// Index set the Index property for MigrationFeatureIndexInfoBuilder.
func (rb *MigrationFeatureIndexInfoBuilder) Index(index IndexName) *MigrationFeatureIndexInfoBuilder {
	rb.v.Index = index
	return rb
}

// Version set the Version property for MigrationFeatureIndexInfoBuilder.
func (rb *MigrationFeatureIndexInfoBuilder) Version(version VersionString) *MigrationFeatureIndexInfoBuilder {
	rb.v.Version = version
	return rb
}
