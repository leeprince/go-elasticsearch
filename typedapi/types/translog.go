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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/translogdurability"
)

// Translog type.
type Translog struct {
	// Durability Whether or not to `fsync` and commit the translog after every index, delete,
	// update, or bulk request.
	Durability *translogdurability.TranslogDurability `json:"durability,omitempty"`
	// FlushThresholdSize The translog stores all operations that are not yet safely persisted in
	// Lucene (i.e., are not
	// part of a Lucene commit point). Although these operations are available for
	// reads, they will need
	// to be replayed if the shard was stopped and had to be recovered. This setting
	// controls the
	// maximum total size of these operations, to prevent recoveries from taking too
	// long. Once the
	// maximum size has been reached a flush will happen, generating a new Lucene
	// commit point.
	FlushThresholdSize *ByteSize          `json:"flush_threshold_size,omitempty"`
	Retention          *TranslogRetention `json:"retention,omitempty"`
	// SyncInterval How often the translog is fsynced to disk and committed, regardless of write
	// operations.
	// Values less than 100ms are not allowed.
	SyncInterval *Time `json:"sync_interval,omitempty"`
}

// TranslogBuilder holds Translog struct and provides a builder API.
type TranslogBuilder struct {
	v *Translog
}

// NewTranslog provides a builder for the Translog struct.
func NewTranslog() *TranslogBuilder {
	r := TranslogBuilder{
		&Translog{},
	}

	return &r
}

// Build finalize the chain and returns the Translog struct
func (rb *TranslogBuilder) Build() Translog {
	return *rb.v
}

// Durability Whether or not to `fsync` and commit the translog after every index, delete,
// update, or bulk request.
func (rb *TranslogBuilder) Durability(durability translogdurability.TranslogDurability) *TranslogBuilder {
	rb.v.Durability = &durability
	return rb
}

// FlushThresholdSize The translog stores all operations that are not yet safely persisted in
// Lucene (i.e., are not
// part of a Lucene commit point). Although these operations are available for
// reads, they will need
// to be replayed if the shard was stopped and had to be recovered. This setting
// controls the
// maximum total size of these operations, to prevent recoveries from taking too
// long. Once the
// maximum size has been reached a flush will happen, generating a new Lucene
// commit point.
func (rb *TranslogBuilder) FlushThresholdSize(flushthresholdsize ByteSize) *TranslogBuilder {
	rb.v.FlushThresholdSize = &flushthresholdsize
	return rb
}

// Retention set the Retention property for TranslogBuilder.
func (rb *TranslogBuilder) Retention(retention TranslogRetention) *TranslogBuilder {
	rb.v.Retention = &retention
	return rb
}

// SyncInterval How often the translog is fsynced to disk and committed, regardless of write
// operations.
// Values less than 100ms are not allowed.
func (rb *TranslogBuilder) SyncInterval(syncinterval Time) *TranslogBuilder {
	rb.v.SyncInterval = &syncinterval
	return rb
}
