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

// RecoveryIndexStatus type.
type RecoveryIndexStatus struct {
	Bytes                      *RecoveryBytes `json:"bytes,omitempty"`
	Files                      RecoveryFiles  `json:"files"`
	Size                       RecoveryBytes  `json:"size"`
	SourceThrottleTime         *Time          `json:"source_throttle_time,omitempty"`
	SourceThrottleTimeInMillis EpochMillis    `json:"source_throttle_time_in_millis"`
	TargetThrottleTime         *Time          `json:"target_throttle_time,omitempty"`
	TargetThrottleTimeInMillis EpochMillis    `json:"target_throttle_time_in_millis"`
	TotalTime                  *Time          `json:"total_time,omitempty"`
	TotalTimeInMillis          EpochMillis    `json:"total_time_in_millis"`
}

// RecoveryIndexStatusBuilder holds RecoveryIndexStatus struct and provides a builder API.
type RecoveryIndexStatusBuilder struct {
	v *RecoveryIndexStatus
}

// NewRecoveryIndexStatus provides a builder for the RecoveryIndexStatus struct.
func NewRecoveryIndexStatus() *RecoveryIndexStatusBuilder {
	r := RecoveryIndexStatusBuilder{
		&RecoveryIndexStatus{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryIndexStatus struct
func (rb *RecoveryIndexStatusBuilder) Build() RecoveryIndexStatus {
	return *rb.v
}

// Bytes set the Bytes property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) Bytes(bytes RecoveryBytes) *RecoveryIndexStatusBuilder {
	rb.v.Bytes = &bytes
	return rb
}

// Files set the Files property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) Files(files RecoveryFiles) *RecoveryIndexStatusBuilder {
	rb.v.Files = files
	return rb
}

// Size set the Size property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) Size(size RecoveryBytes) *RecoveryIndexStatusBuilder {
	rb.v.Size = size
	return rb
}

// SourceThrottleTime set the SourceThrottleTime property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) SourceThrottleTime(sourcethrottletime Time) *RecoveryIndexStatusBuilder {
	rb.v.SourceThrottleTime = &sourcethrottletime
	return rb
}

// SourceThrottleTimeInMillis set the SourceThrottleTimeInMillis property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) SourceThrottleTimeInMillis(sourcethrottletimeinmillis EpochMillis) *RecoveryIndexStatusBuilder {
	rb.v.SourceThrottleTimeInMillis = sourcethrottletimeinmillis
	return rb
}

// TargetThrottleTime set the TargetThrottleTime property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) TargetThrottleTime(targetthrottletime Time) *RecoveryIndexStatusBuilder {
	rb.v.TargetThrottleTime = &targetthrottletime
	return rb
}

// TargetThrottleTimeInMillis set the TargetThrottleTimeInMillis property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) TargetThrottleTimeInMillis(targetthrottletimeinmillis EpochMillis) *RecoveryIndexStatusBuilder {
	rb.v.TargetThrottleTimeInMillis = targetthrottletimeinmillis
	return rb
}

// TotalTime set the TotalTime property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) TotalTime(totaltime Time) *RecoveryIndexStatusBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

// TotalTimeInMillis set the TotalTimeInMillis property for RecoveryIndexStatusBuilder.
func (rb *RecoveryIndexStatusBuilder) TotalTimeInMillis(totaltimeinmillis EpochMillis) *RecoveryIndexStatusBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
