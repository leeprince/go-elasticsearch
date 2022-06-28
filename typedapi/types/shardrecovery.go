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

// ShardRecovery type.
type ShardRecovery struct {
	Id                int64                `json:"id"`
	Index             RecoveryIndexStatus  `json:"index"`
	Primary           bool                 `json:"primary"`
	Source            RecoveryOrigin       `json:"source"`
	Stage             string               `json:"stage"`
	Start             *RecoveryStartStatus `json:"start,omitempty"`
	StartTime         *DateString          `json:"start_time,omitempty"`
	StartTimeInMillis EpochMillis          `json:"start_time_in_millis"`
	StopTime          *DateString          `json:"stop_time,omitempty"`
	StopTimeInMillis  EpochMillis          `json:"stop_time_in_millis"`
	Target            RecoveryOrigin       `json:"target"`
	TotalTime         *DateString          `json:"total_time,omitempty"`
	TotalTimeInMillis EpochMillis          `json:"total_time_in_millis"`
	Translog          TranslogStatus       `json:"translog"`
	Type              string               `json:"type"`
	VerifyIndex       VerifyIndex          `json:"verify_index"`
}

// ShardRecoveryBuilder holds ShardRecovery struct and provides a builder API.
type ShardRecoveryBuilder struct {
	v *ShardRecovery
}

// NewShardRecovery provides a builder for the ShardRecovery struct.
func NewShardRecovery() *ShardRecoveryBuilder {
	r := ShardRecoveryBuilder{
		&ShardRecovery{},
	}

	return &r
}

// Build finalize the chain and returns the ShardRecovery struct
func (rb *ShardRecoveryBuilder) Build() ShardRecovery {
	return *rb.v
}

// Id set the Id property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Id(id int64) *ShardRecoveryBuilder {
	rb.v.Id = id
	return rb
}

// Index set the Index property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Index(index RecoveryIndexStatus) *ShardRecoveryBuilder {
	rb.v.Index = index
	return rb
}

// Primary set the Primary property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Primary(primary bool) *ShardRecoveryBuilder {
	rb.v.Primary = primary
	return rb
}

// Source set the Source property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Source(source RecoveryOrigin) *ShardRecoveryBuilder {
	rb.v.Source = source
	return rb
}

// Stage set the Stage property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Stage(stage string) *ShardRecoveryBuilder {
	rb.v.Stage = stage
	return rb
}

// Start set the Start property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Start(start RecoveryStartStatus) *ShardRecoveryBuilder {
	rb.v.Start = &start
	return rb
}

// StartTime set the StartTime property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) StartTime(starttime DateString) *ShardRecoveryBuilder {
	rb.v.StartTime = &starttime
	return rb
}

// StartTimeInMillis set the StartTimeInMillis property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) StartTimeInMillis(starttimeinmillis EpochMillis) *ShardRecoveryBuilder {
	rb.v.StartTimeInMillis = starttimeinmillis
	return rb
}

// StopTime set the StopTime property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) StopTime(stoptime DateString) *ShardRecoveryBuilder {
	rb.v.StopTime = &stoptime
	return rb
}

// StopTimeInMillis set the StopTimeInMillis property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) StopTimeInMillis(stoptimeinmillis EpochMillis) *ShardRecoveryBuilder {
	rb.v.StopTimeInMillis = stoptimeinmillis
	return rb
}

// Target set the Target property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Target(target RecoveryOrigin) *ShardRecoveryBuilder {
	rb.v.Target = target
	return rb
}

// TotalTime set the TotalTime property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) TotalTime(totaltime DateString) *ShardRecoveryBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

// TotalTimeInMillis set the TotalTimeInMillis property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) TotalTimeInMillis(totaltimeinmillis EpochMillis) *ShardRecoveryBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}

// Translog set the Translog property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Translog(translog TranslogStatus) *ShardRecoveryBuilder {
	rb.v.Translog = translog
	return rb
}

// Type set the Type property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) Type_(type_ string) *ShardRecoveryBuilder {
	rb.v.Type = type_
	return rb
}

// VerifyIndex set the VerifyIndex property for ShardRecoveryBuilder.
func (rb *ShardRecoveryBuilder) VerifyIndex(verifyindex VerifyIndex) *ShardRecoveryBuilder {
	rb.v.VerifyIndex = verifyindex
	return rb
}
