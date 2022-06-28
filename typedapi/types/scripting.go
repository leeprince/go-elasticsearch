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

// Scripting type.
type Scripting struct {
	CacheEvictions            *int64    `json:"cache_evictions,omitempty"`
	CompilationLimitTriggered *int64    `json:"compilation_limit_triggered,omitempty"`
	Compilations              *int64    `json:"compilations,omitempty"`
	Contexts                  []Context `json:"contexts,omitempty"`
}

// ScriptingBuilder holds Scripting struct and provides a builder API.
type ScriptingBuilder struct {
	v *Scripting
}

// NewScripting provides a builder for the Scripting struct.
func NewScripting() *ScriptingBuilder {
	r := ScriptingBuilder{
		&Scripting{},
	}

	return &r
}

// Build finalize the chain and returns the Scripting struct
func (rb *ScriptingBuilder) Build() Scripting {
	return *rb.v
}

// CacheEvictions set the CacheEvictions property for ScriptingBuilder.
func (rb *ScriptingBuilder) CacheEvictions(cacheevictions int64) *ScriptingBuilder {
	rb.v.CacheEvictions = &cacheevictions
	return rb
}

// CompilationLimitTriggered set the CompilationLimitTriggered property for ScriptingBuilder.
func (rb *ScriptingBuilder) CompilationLimitTriggered(compilationlimittriggered int64) *ScriptingBuilder {
	rb.v.CompilationLimitTriggered = &compilationlimittriggered
	return rb
}

// Compilations set the Compilations property for ScriptingBuilder.
func (rb *ScriptingBuilder) Compilations(compilations int64) *ScriptingBuilder {
	rb.v.Compilations = &compilations
	return rb
}

// Contexts set the Contexts property for ScriptingBuilder.
func (rb *ScriptingBuilder) Contexts(contexts ...Context) *ScriptingBuilder {
	rb.v.Contexts = append(rb.v.Contexts, contexts...)
	return rb
}
