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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/unassignedinformationreason"
)

// UnassignedInformation type.
type UnassignedInformation struct {
	AllocationStatus         *string                                                 `json:"allocation_status,omitempty"`
	At                       DateString                                              `json:"at"`
	Delayed                  *bool                                                   `json:"delayed,omitempty"`
	Details                  *string                                                 `json:"details,omitempty"`
	FailedAllocationAttempts *int                                                    `json:"failed_allocation_attempts,omitempty"`
	LastAllocationStatus     *string                                                 `json:"last_allocation_status,omitempty"`
	Reason                   unassignedinformationreason.UnassignedInformationReason `json:"reason"`
}

// UnassignedInformationBuilder holds UnassignedInformation struct and provides a builder API.
type UnassignedInformationBuilder struct {
	v *UnassignedInformation
}

// NewUnassignedInformation provides a builder for the UnassignedInformation struct.
func NewUnassignedInformation() *UnassignedInformationBuilder {
	r := UnassignedInformationBuilder{
		&UnassignedInformation{},
	}

	return &r
}

// Build finalize the chain and returns the UnassignedInformation struct
func (rb *UnassignedInformationBuilder) Build() UnassignedInformation {
	return *rb.v
}

// AllocationStatus set the AllocationStatus property for UnassignedInformationBuilder.
func (rb *UnassignedInformationBuilder) AllocationStatus(allocationstatus string) *UnassignedInformationBuilder {
	rb.v.AllocationStatus = &allocationstatus
	return rb
}

// At set the At property for UnassignedInformationBuilder.
func (rb *UnassignedInformationBuilder) At(at DateString) *UnassignedInformationBuilder {
	rb.v.At = at
	return rb
}

// Delayed set the Delayed property for UnassignedInformationBuilder.
func (rb *UnassignedInformationBuilder) Delayed(delayed bool) *UnassignedInformationBuilder {
	rb.v.Delayed = &delayed
	return rb
}

// Details set the Details property for UnassignedInformationBuilder.
func (rb *UnassignedInformationBuilder) Details(details string) *UnassignedInformationBuilder {
	rb.v.Details = &details
	return rb
}

// FailedAllocationAttempts set the FailedAllocationAttempts property for UnassignedInformationBuilder.
func (rb *UnassignedInformationBuilder) FailedAllocationAttempts(failedallocationattempts int) *UnassignedInformationBuilder {
	rb.v.FailedAllocationAttempts = &failedallocationattempts
	return rb
}

// LastAllocationStatus set the LastAllocationStatus property for UnassignedInformationBuilder.
func (rb *UnassignedInformationBuilder) LastAllocationStatus(lastallocationstatus string) *UnassignedInformationBuilder {
	rb.v.LastAllocationStatus = &lastallocationstatus
	return rb
}

// Reason set the Reason property for UnassignedInformationBuilder.
func (rb *UnassignedInformationBuilder) Reason(reason unassignedinformationreason.UnassignedInformationReason) *UnassignedInformationBuilder {
	rb.v.Reason = reason
	return rb
}
