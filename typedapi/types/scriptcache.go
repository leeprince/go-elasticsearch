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

// ScriptCache type.
type ScriptCache struct {
	CacheEvictions            *int64  `json:"cache_evictions,omitempty"`
	CompilationLimitTriggered *int64  `json:"compilation_limit_triggered,omitempty"`
	Compilations              *int64  `json:"compilations,omitempty"`
	Context                   *string `json:"context,omitempty"`
}

// ScriptCacheBuilder holds ScriptCache struct and provides a builder API.
type ScriptCacheBuilder struct {
	v *ScriptCache
}

// NewScriptCache provides a builder for the ScriptCache struct.
func NewScriptCache() *ScriptCacheBuilder {
	r := ScriptCacheBuilder{
		&ScriptCache{},
	}

	return &r
}

// Build finalize the chain and returns the ScriptCache struct
func (rb *ScriptCacheBuilder) Build() ScriptCache {
	return *rb.v
}

// CacheEvictions set the CacheEvictions property for ScriptCacheBuilder.
func (rb *ScriptCacheBuilder) CacheEvictions(cacheevictions int64) *ScriptCacheBuilder {
	rb.v.CacheEvictions = &cacheevictions
	return rb
}

// CompilationLimitTriggered set the CompilationLimitTriggered property for ScriptCacheBuilder.
func (rb *ScriptCacheBuilder) CompilationLimitTriggered(compilationlimittriggered int64) *ScriptCacheBuilder {
	rb.v.CompilationLimitTriggered = &compilationlimittriggered
	return rb
}

// Compilations set the Compilations property for ScriptCacheBuilder.
func (rb *ScriptCacheBuilder) Compilations(compilations int64) *ScriptCacheBuilder {
	rb.v.Compilations = &compilations
	return rb
}

// Context set the Context property for ScriptCacheBuilder.
func (rb *ScriptCacheBuilder) Context(context string) *ScriptCacheBuilder {
	rb.v.Context = &context
	return rb
}
