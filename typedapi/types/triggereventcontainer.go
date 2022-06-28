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

// TriggerEventContainer type.
type TriggerEventContainer struct {
	Schedule *ScheduleTriggerEvent `json:"schedule,omitempty"`
}

// TriggerEventContainerBuilder holds TriggerEventContainer struct and provides a builder API.
type TriggerEventContainerBuilder struct {
	v *TriggerEventContainer
}

// NewTriggerEventContainer provides a builder for the TriggerEventContainer struct.
func NewTriggerEventContainer() *TriggerEventContainerBuilder {
	r := TriggerEventContainerBuilder{
		&TriggerEventContainer{},
	}

	return &r
}

// Build finalize the chain and returns the TriggerEventContainer struct
func (rb *TriggerEventContainerBuilder) Build() TriggerEventContainer {
	return *rb.v
}

// Schedule set the Schedule property for TriggerEventContainerBuilder.
func (rb *TriggerEventContainerBuilder) Schedule(schedule ScheduleTriggerEvent) *TriggerEventContainerBuilder {
	rb.v.Schedule = &schedule
	return rb
}
