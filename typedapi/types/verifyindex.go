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

// VerifyIndex type.
type VerifyIndex struct {
	CheckIndexTime         *Time       `json:"check_index_time,omitempty"`
	CheckIndexTimeInMillis EpochMillis `json:"check_index_time_in_millis"`
	TotalTime              *Time       `json:"total_time,omitempty"`
	TotalTimeInMillis      EpochMillis `json:"total_time_in_millis"`
}

// VerifyIndexBuilder holds VerifyIndex struct and provides a builder API.
type VerifyIndexBuilder struct {
	v *VerifyIndex
}

// NewVerifyIndex provides a builder for the VerifyIndex struct.
func NewVerifyIndex() *VerifyIndexBuilder {
	r := VerifyIndexBuilder{
		&VerifyIndex{},
	}

	return &r
}

// Build finalize the chain and returns the VerifyIndex struct
func (rb *VerifyIndexBuilder) Build() VerifyIndex {
	return *rb.v
}

// CheckIndexTime set the CheckIndexTime property for VerifyIndexBuilder.
func (rb *VerifyIndexBuilder) CheckIndexTime(checkindextime Time) *VerifyIndexBuilder {
	rb.v.CheckIndexTime = &checkindextime
	return rb
}

// CheckIndexTimeInMillis set the CheckIndexTimeInMillis property for VerifyIndexBuilder.
func (rb *VerifyIndexBuilder) CheckIndexTimeInMillis(checkindextimeinmillis EpochMillis) *VerifyIndexBuilder {
	rb.v.CheckIndexTimeInMillis = checkindextimeinmillis
	return rb
}

// TotalTime set the TotalTime property for VerifyIndexBuilder.
func (rb *VerifyIndexBuilder) TotalTime(totaltime Time) *VerifyIndexBuilder {
	rb.v.TotalTime = &totaltime
	return rb
}

// TotalTimeInMillis set the TotalTimeInMillis property for VerifyIndexBuilder.
func (rb *VerifyIndexBuilder) TotalTimeInMillis(totaltimeinmillis EpochMillis) *VerifyIndexBuilder {
	rb.v.TotalTimeInMillis = totaltimeinmillis
	return rb
}
