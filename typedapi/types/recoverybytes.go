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

// RecoveryBytes type.
type RecoveryBytes struct {
	Percent                      Percentage `json:"percent"`
	Recovered                    *ByteSize  `json:"recovered,omitempty"`
	RecoveredFromSnapshot        *ByteSize  `json:"recovered_from_snapshot,omitempty"`
	RecoveredFromSnapshotInBytes *ByteSize  `json:"recovered_from_snapshot_in_bytes,omitempty"`
	RecoveredInBytes             ByteSize   `json:"recovered_in_bytes"`
	Reused                       *ByteSize  `json:"reused,omitempty"`
	ReusedInBytes                ByteSize   `json:"reused_in_bytes"`
	Total                        *ByteSize  `json:"total,omitempty"`
	TotalInBytes                 ByteSize   `json:"total_in_bytes"`
}

// RecoveryBytesBuilder holds RecoveryBytes struct and provides a builder API.
type RecoveryBytesBuilder struct {
	v *RecoveryBytes
}

// NewRecoveryBytes provides a builder for the RecoveryBytes struct.
func NewRecoveryBytes() *RecoveryBytesBuilder {
	r := RecoveryBytesBuilder{
		&RecoveryBytes{},
	}

	return &r
}

// Build finalize the chain and returns the RecoveryBytes struct
func (rb *RecoveryBytesBuilder) Build() RecoveryBytes {
	return *rb.v
}

// Percent set the Percent property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) Percent(percent Percentage) *RecoveryBytesBuilder {
	rb.v.Percent = percent
	return rb
}

// Recovered set the Recovered property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) Recovered(recovered ByteSize) *RecoveryBytesBuilder {
	rb.v.Recovered = &recovered
	return rb
}

// RecoveredFromSnapshot set the RecoveredFromSnapshot property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) RecoveredFromSnapshot(recoveredfromsnapshot ByteSize) *RecoveryBytesBuilder {
	rb.v.RecoveredFromSnapshot = &recoveredfromsnapshot
	return rb
}

// RecoveredFromSnapshotInBytes set the RecoveredFromSnapshotInBytes property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) RecoveredFromSnapshotInBytes(recoveredfromsnapshotinbytes ByteSize) *RecoveryBytesBuilder {
	rb.v.RecoveredFromSnapshotInBytes = &recoveredfromsnapshotinbytes
	return rb
}

// RecoveredInBytes set the RecoveredInBytes property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) RecoveredInBytes(recoveredinbytes ByteSize) *RecoveryBytesBuilder {
	rb.v.RecoveredInBytes = recoveredinbytes
	return rb
}

// Reused set the Reused property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) Reused(reused ByteSize) *RecoveryBytesBuilder {
	rb.v.Reused = &reused
	return rb
}

// ReusedInBytes set the ReusedInBytes property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) ReusedInBytes(reusedinbytes ByteSize) *RecoveryBytesBuilder {
	rb.v.ReusedInBytes = reusedinbytes
	return rb
}

// Total set the Total property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) Total(total ByteSize) *RecoveryBytesBuilder {
	rb.v.Total = &total
	return rb
}

// TotalInBytes set the TotalInBytes property for RecoveryBytesBuilder.
func (rb *RecoveryBytesBuilder) TotalInBytes(totalinbytes ByteSize) *RecoveryBytesBuilder {
	rb.v.TotalInBytes = totalinbytes
	return rb
}
