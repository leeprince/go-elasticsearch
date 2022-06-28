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

// SnapshotLifecycle type.
type SnapshotLifecycle struct {
	InProgress          *InProgress   `json:"in_progress,omitempty"`
	LastFailure         *Invocation   `json:"last_failure,omitempty"`
	LastSuccess         *Invocation   `json:"last_success,omitempty"`
	ModifiedDate        *DateString   `json:"modified_date,omitempty"`
	ModifiedDateMillis  EpochMillis   `json:"modified_date_millis"`
	NextExecution       *DateString   `json:"next_execution,omitempty"`
	NextExecutionMillis EpochMillis   `json:"next_execution_millis"`
	Policy              Policy        `json:"policy"`
	Stats               Statistics    `json:"stats"`
	Version             VersionNumber `json:"version"`
}

// SnapshotLifecycleBuilder holds SnapshotLifecycle struct and provides a builder API.
type SnapshotLifecycleBuilder struct {
	v *SnapshotLifecycle
}

// NewSnapshotLifecycle provides a builder for the SnapshotLifecycle struct.
func NewSnapshotLifecycle() *SnapshotLifecycleBuilder {
	r := SnapshotLifecycleBuilder{
		&SnapshotLifecycle{},
	}

	return &r
}

// Build finalize the chain and returns the SnapshotLifecycle struct
func (rb *SnapshotLifecycleBuilder) Build() SnapshotLifecycle {
	return *rb.v
}

// InProgress set the InProgress property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) InProgress(inprogress InProgress) *SnapshotLifecycleBuilder {
	rb.v.InProgress = &inprogress
	return rb
}

// LastFailure set the LastFailure property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) LastFailure(lastfailure Invocation) *SnapshotLifecycleBuilder {
	rb.v.LastFailure = &lastfailure
	return rb
}

// LastSuccess set the LastSuccess property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) LastSuccess(lastsuccess Invocation) *SnapshotLifecycleBuilder {
	rb.v.LastSuccess = &lastsuccess
	return rb
}

// ModifiedDate set the ModifiedDate property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) ModifiedDate(modifieddate DateString) *SnapshotLifecycleBuilder {
	rb.v.ModifiedDate = &modifieddate
	return rb
}

// ModifiedDateMillis set the ModifiedDateMillis property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) ModifiedDateMillis(modifieddatemillis EpochMillis) *SnapshotLifecycleBuilder {
	rb.v.ModifiedDateMillis = modifieddatemillis
	return rb
}

// NextExecution set the NextExecution property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) NextExecution(nextexecution DateString) *SnapshotLifecycleBuilder {
	rb.v.NextExecution = &nextexecution
	return rb
}

// NextExecutionMillis set the NextExecutionMillis property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) NextExecutionMillis(nextexecutionmillis EpochMillis) *SnapshotLifecycleBuilder {
	rb.v.NextExecutionMillis = nextexecutionmillis
	return rb
}

// Policy set the Policy property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) Policy(policy Policy) *SnapshotLifecycleBuilder {
	rb.v.Policy = policy
	return rb
}

// Stats set the Stats property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) Stats(stats Statistics) *SnapshotLifecycleBuilder {
	rb.v.Stats = stats
	return rb
}

// Version set the Version property for SnapshotLifecycleBuilder.
func (rb *SnapshotLifecycleBuilder) Version(version VersionNumber) *SnapshotLifecycleBuilder {
	rb.v.Version = version
	return rb
}
