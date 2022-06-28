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

// SecurityRolesDlsBitSetCache type.
type SecurityRolesDlsBitSetCache struct {
	Count         int       `json:"count"`
	Memory        *ByteSize `json:"memory,omitempty"`
	MemoryInBytes uint64    `json:"memory_in_bytes"`
}

// SecurityRolesDlsBitSetCacheBuilder holds SecurityRolesDlsBitSetCache struct and provides a builder API.
type SecurityRolesDlsBitSetCacheBuilder struct {
	v *SecurityRolesDlsBitSetCache
}

// NewSecurityRolesDlsBitSetCache provides a builder for the SecurityRolesDlsBitSetCache struct.
func NewSecurityRolesDlsBitSetCache() *SecurityRolesDlsBitSetCacheBuilder {
	r := SecurityRolesDlsBitSetCacheBuilder{
		&SecurityRolesDlsBitSetCache{},
	}

	return &r
}

// Build finalize the chain and returns the SecurityRolesDlsBitSetCache struct
func (rb *SecurityRolesDlsBitSetCacheBuilder) Build() SecurityRolesDlsBitSetCache {
	return *rb.v
}

// Count set the Count property for SecurityRolesDlsBitSetCacheBuilder.
func (rb *SecurityRolesDlsBitSetCacheBuilder) Count(count int) *SecurityRolesDlsBitSetCacheBuilder {
	rb.v.Count = count
	return rb
}

// Memory set the Memory property for SecurityRolesDlsBitSetCacheBuilder.
func (rb *SecurityRolesDlsBitSetCacheBuilder) Memory(memory ByteSize) *SecurityRolesDlsBitSetCacheBuilder {
	rb.v.Memory = &memory
	return rb
}

// MemoryInBytes set the MemoryInBytes property for SecurityRolesDlsBitSetCacheBuilder.
func (rb *SecurityRolesDlsBitSetCacheBuilder) MemoryInBytes(memoryinbytes uint64) *SecurityRolesDlsBitSetCacheBuilder {
	rb.v.MemoryInBytes = memoryinbytes
	return rb
}
