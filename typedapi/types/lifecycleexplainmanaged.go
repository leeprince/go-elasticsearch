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

// LifecycleExplainManaged type.
type LifecycleExplainManaged struct {
	Action                  Name                           `json:"action"`
	ActionTimeMillis        EpochMillis                    `json:"action_time_millis"`
	Age                     Time                           `json:"age"`
	FailedStep              *Name                          `json:"failed_step,omitempty"`
	FailedStepRetryCount    *int                           `json:"failed_step_retry_count,omitempty"`
	Index                   IndexName                      `json:"index"`
	IndexCreationDateMillis *EpochMillis                   `json:"index_creation_date_millis,omitempty"`
	IsAutoRetryableError    *bool                          `json:"is_auto_retryable_error,omitempty"`
	LifecycleDateMillis     EpochMillis                    `json:"lifecycle_date_millis"`
	Managed                 string                         `json:"managed,omitempty"`
	Phase                   Name                           `json:"phase"`
	PhaseExecution          LifecycleExplainPhaseExecution `json:"phase_execution"`
	PhaseTimeMillis         EpochMillis                    `json:"phase_time_millis"`
	Policy                  Name                           `json:"policy"`
	Step                    Name                           `json:"step"`
	StepInfo                map[string]interface{}         `json:"step_info,omitempty"`
	StepTimeMillis          EpochMillis                    `json:"step_time_millis"`
	TimeSinceIndexCreation  *Time                          `json:"time_since_index_creation,omitempty"`
}

// LifecycleExplainManagedBuilder holds LifecycleExplainManaged struct and provides a builder API.
type LifecycleExplainManagedBuilder struct {
	v *LifecycleExplainManaged
}

// NewLifecycleExplainManaged provides a builder for the LifecycleExplainManaged struct.
func NewLifecycleExplainManaged() *LifecycleExplainManagedBuilder {
	r := LifecycleExplainManagedBuilder{
		&LifecycleExplainManaged{
			StepInfo: make(map[string]interface{}, 0),
		},
	}

	r.v.Managed = "true"

	return &r
}

// Build finalize the chain and returns the LifecycleExplainManaged struct
func (rb *LifecycleExplainManagedBuilder) Build() LifecycleExplainManaged {
	return *rb.v
}

// Action set the Action property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) Action(action Name) *LifecycleExplainManagedBuilder {
	rb.v.Action = action
	return rb
}

// ActionTimeMillis set the ActionTimeMillis property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) ActionTimeMillis(actiontimemillis EpochMillis) *LifecycleExplainManagedBuilder {
	rb.v.ActionTimeMillis = actiontimemillis
	return rb
}

// Age set the Age property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) Age(age Time) *LifecycleExplainManagedBuilder {
	rb.v.Age = age
	return rb
}

// FailedStep set the FailedStep property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) FailedStep(failedstep Name) *LifecycleExplainManagedBuilder {
	rb.v.FailedStep = &failedstep
	return rb
}

// FailedStepRetryCount set the FailedStepRetryCount property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) FailedStepRetryCount(failedstepretrycount int) *LifecycleExplainManagedBuilder {
	rb.v.FailedStepRetryCount = &failedstepretrycount
	return rb
}

// Index set the Index property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) Index(index IndexName) *LifecycleExplainManagedBuilder {
	rb.v.Index = index
	return rb
}

// IndexCreationDateMillis set the IndexCreationDateMillis property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) IndexCreationDateMillis(indexcreationdatemillis EpochMillis) *LifecycleExplainManagedBuilder {
	rb.v.IndexCreationDateMillis = &indexcreationdatemillis
	return rb
}

// IsAutoRetryableError set the IsAutoRetryableError property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) IsAutoRetryableError(isautoretryableerror bool) *LifecycleExplainManagedBuilder {
	rb.v.IsAutoRetryableError = &isautoretryableerror
	return rb
}

// LifecycleDateMillis set the LifecycleDateMillis property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) LifecycleDateMillis(lifecycledatemillis EpochMillis) *LifecycleExplainManagedBuilder {
	rb.v.LifecycleDateMillis = lifecycledatemillis
	return rb
}

// Managed set the Managed property for LifecycleExplainManagedBuilder.

// Phase set the Phase property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) Phase(phase Name) *LifecycleExplainManagedBuilder {
	rb.v.Phase = phase
	return rb
}

// PhaseExecution set the PhaseExecution property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) PhaseExecution(phaseexecution LifecycleExplainPhaseExecution) *LifecycleExplainManagedBuilder {
	rb.v.PhaseExecution = phaseexecution
	return rb
}

// PhaseTimeMillis set the PhaseTimeMillis property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) PhaseTimeMillis(phasetimemillis EpochMillis) *LifecycleExplainManagedBuilder {
	rb.v.PhaseTimeMillis = phasetimemillis
	return rb
}

// Policy set the Policy property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) Policy(policy Name) *LifecycleExplainManagedBuilder {
	rb.v.Policy = policy
	return rb
}

// Step set the Step property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) Step(step Name) *LifecycleExplainManagedBuilder {
	rb.v.Step = step
	return rb
}

// StepInfo set the StepInfo property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) StepInfo(value map[string]interface{}) *LifecycleExplainManagedBuilder {
	rb.v.StepInfo = value
	return rb
}

// StepTimeMillis set the StepTimeMillis property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) StepTimeMillis(steptimemillis EpochMillis) *LifecycleExplainManagedBuilder {
	rb.v.StepTimeMillis = steptimemillis
	return rb
}

// TimeSinceIndexCreation set the TimeSinceIndexCreation property for LifecycleExplainManagedBuilder.
func (rb *LifecycleExplainManagedBuilder) TimeSinceIndexCreation(timesinceindexcreation Time) *LifecycleExplainManagedBuilder {
	rb.v.TimeSinceIndexCreation = &timesinceindexcreation
	return rb
}
