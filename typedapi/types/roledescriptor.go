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

// RoleDescriptor type.
type RoleDescriptor struct {
	Applications      []ApplicationPrivileges  `json:"applications,omitempty"`
	Cluster           []string                 `json:"cluster"`
	Global            []GlobalPrivilege        `json:"global,omitempty"`
	Indices           []IndicesPrivileges      `json:"indices"`
	Metadata          *Metadata                `json:"metadata,omitempty"`
	RunAs             []string                 `json:"run_as,omitempty"`
	TransientMetadata *TransientMetadataConfig `json:"transient_metadata,omitempty"`
}

// RoleDescriptorBuilder holds RoleDescriptor struct and provides a builder API.
type RoleDescriptorBuilder struct {
	v *RoleDescriptor
}

// NewRoleDescriptor provides a builder for the RoleDescriptor struct.
func NewRoleDescriptor() *RoleDescriptorBuilder {
	r := RoleDescriptorBuilder{
		&RoleDescriptor{},
	}

	return &r
}

// Build finalize the chain and returns the RoleDescriptor struct
func (rb *RoleDescriptorBuilder) Build() RoleDescriptor {
	return *rb.v
}

// Applications set the Applications property for RoleDescriptorBuilder.
func (rb *RoleDescriptorBuilder) Applications(applications ...ApplicationPrivileges) *RoleDescriptorBuilder {
	rb.v.Applications = append(rb.v.Applications, applications...)
	return rb
}

// Cluster set the Cluster property for RoleDescriptorBuilder.
func (rb *RoleDescriptorBuilder) Cluster(cluster ...string) *RoleDescriptorBuilder {
	rb.v.Cluster = append(rb.v.Cluster, cluster...)
	return rb
}

// Global set the Global property for RoleDescriptorBuilder.
func (rb *RoleDescriptorBuilder) Global(arg []GlobalPrivilege) *RoleDescriptorBuilder {
	rb.v.Global = arg
	return rb
}

// Indices set the Indices property for RoleDescriptorBuilder.
func (rb *RoleDescriptorBuilder) Indices(indices ...IndicesPrivileges) *RoleDescriptorBuilder {
	rb.v.Indices = append(rb.v.Indices, indices...)
	return rb
}

// Metadata set the Metadata property for RoleDescriptorBuilder.
func (rb *RoleDescriptorBuilder) Metadata(metadata Metadata) *RoleDescriptorBuilder {
	rb.v.Metadata = &metadata
	return rb
}

// RunAs set the RunAs property for RoleDescriptorBuilder.
func (rb *RoleDescriptorBuilder) RunAs(run_as ...string) *RoleDescriptorBuilder {
	rb.v.RunAs = append(rb.v.RunAs, run_as...)
	return rb
}

// TransientMetadata set the TransientMetadata property for RoleDescriptorBuilder.
func (rb *RoleDescriptorBuilder) TransientMetadata(transientmetadata TransientMetadataConfig) *RoleDescriptorBuilder {
	rb.v.TransientMetadata = &transientmetadata
	return rb
}
