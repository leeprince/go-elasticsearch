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

// LifecycleExplainPhaseExecution type.
type LifecycleExplainPhaseExecution struct {
	ModifiedDateInMillis EpochMillis   `json:"modified_date_in_millis"`
	Policy               Name          `json:"policy"`
	Version              VersionNumber `json:"version"`
}

// LifecycleExplainPhaseExecutionBuilder holds LifecycleExplainPhaseExecution struct and provides a builder API.
type LifecycleExplainPhaseExecutionBuilder struct {
	v *LifecycleExplainPhaseExecution
}

// NewLifecycleExplainPhaseExecution provides a builder for the LifecycleExplainPhaseExecution struct.
func NewLifecycleExplainPhaseExecution() *LifecycleExplainPhaseExecutionBuilder {
	r := LifecycleExplainPhaseExecutionBuilder{
		&LifecycleExplainPhaseExecution{},
	}

	return &r
}

// Build finalize the chain and returns the LifecycleExplainPhaseExecution struct
func (rb *LifecycleExplainPhaseExecutionBuilder) Build() LifecycleExplainPhaseExecution {
	return *rb.v
}

// ModifiedDateInMillis set the ModifiedDateInMillis property for LifecycleExplainPhaseExecutionBuilder.
func (rb *LifecycleExplainPhaseExecutionBuilder) ModifiedDateInMillis(modifieddateinmillis EpochMillis) *LifecycleExplainPhaseExecutionBuilder {
	rb.v.ModifiedDateInMillis = modifieddateinmillis
	return rb
}

// Policy set the Policy property for LifecycleExplainPhaseExecutionBuilder.
func (rb *LifecycleExplainPhaseExecutionBuilder) Policy(policy Name) *LifecycleExplainPhaseExecutionBuilder {
	rb.v.Policy = policy
	return rb
}

// Version set the Version property for LifecycleExplainPhaseExecutionBuilder.
func (rb *LifecycleExplainPhaseExecutionBuilder) Version(version VersionNumber) *LifecycleExplainPhaseExecutionBuilder {
	rb.v.Version = version
	return rb
}
