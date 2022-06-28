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

// QueryWatch type.
type QueryWatch struct {
	Id_          Id              `json:"_id"`
	PrimaryTerm_ *int            `json:"_primary_term,omitempty"`
	SeqNo_       *SequenceNumber `json:"_seq_no,omitempty"`
	Status       *WatchStatus    `json:"status,omitempty"`
	Watch        *Watch          `json:"watch,omitempty"`
}

// QueryWatchBuilder holds QueryWatch struct and provides a builder API.
type QueryWatchBuilder struct {
	v *QueryWatch
}

// NewQueryWatch provides a builder for the QueryWatch struct.
func NewQueryWatch() *QueryWatchBuilder {
	r := QueryWatchBuilder{
		&QueryWatch{},
	}

	return &r
}

// Build finalize the chain and returns the QueryWatch struct
func (rb *QueryWatchBuilder) Build() QueryWatch {
	return *rb.v
}

// Id_ set the Id_ property for QueryWatchBuilder.
func (rb *QueryWatchBuilder) Id_(id_ Id) *QueryWatchBuilder {
	rb.v.Id_ = id_
	return rb
}

// PrimaryTerm_ set the PrimaryTerm_ property for QueryWatchBuilder.
func (rb *QueryWatchBuilder) PrimaryTerm_(primaryterm_ int) *QueryWatchBuilder {
	rb.v.PrimaryTerm_ = &primaryterm_
	return rb
}

// SeqNo_ set the SeqNo_ property for QueryWatchBuilder.
func (rb *QueryWatchBuilder) SeqNo_(seqno_ SequenceNumber) *QueryWatchBuilder {
	rb.v.SeqNo_ = &seqno_
	return rb
}

// Status set the Status property for QueryWatchBuilder.
func (rb *QueryWatchBuilder) Status(status WatchStatus) *QueryWatchBuilder {
	rb.v.Status = &status
	return rb
}

// Watch set the Watch property for QueryWatchBuilder.
func (rb *QueryWatchBuilder) Watch(watch Watch) *QueryWatchBuilder {
	rb.v.Watch = &watch
	return rb
}
