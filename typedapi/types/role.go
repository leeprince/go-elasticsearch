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

// Role type.
type Role struct {
	Applications      []ApplicationPrivileges                   `json:"applications"`
	Cluster           []string                                  `json:"cluster"`
	Global            map[string]map[string]map[string][]string `json:"global,omitempty"`
	Indices           []IndicesPrivileges                       `json:"indices"`
	Metadata          Metadata                                  `json:"metadata"`
	RoleTemplates     []RoleTemplate                            `json:"role_templates,omitempty"`
	RunAs             []string                                  `json:"run_as"`
	TransientMetadata TransientMetadataConfig                   `json:"transient_metadata"`
}

// RoleBuilder holds Role struct and provides a builder API.
type RoleBuilder struct {
	v *Role
}

// NewRole provides a builder for the Role struct.
func NewRole() *RoleBuilder {
	r := RoleBuilder{
		&Role{
			Global: make(map[string]map[string]map[string][]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Role struct
func (rb *RoleBuilder) Build() Role {
	return *rb.v
}

// Applications set the Applications property for RoleBuilder.
func (rb *RoleBuilder) Applications(applications ...ApplicationPrivileges) *RoleBuilder {
	rb.v.Applications = append(rb.v.Applications, applications...)
	return rb
}

// Cluster set the Cluster property for RoleBuilder.
func (rb *RoleBuilder) Cluster(cluster ...string) *RoleBuilder {
	rb.v.Cluster = append(rb.v.Cluster, cluster...)
	return rb
}

// Global set the Global property for RoleBuilder.
//TODO: Map in map Not implemented yet: Map
func (rb *RoleBuilder) Global(value map[string]map[string]map[string][]string) *RoleBuilder {
	rb.v.Global = value
	return rb
}

// Indices set the Indices property for RoleBuilder.
func (rb *RoleBuilder) Indices(indices ...IndicesPrivileges) *RoleBuilder {
	rb.v.Indices = append(rb.v.Indices, indices...)
	return rb
}

// Metadata set the Metadata property for RoleBuilder.
func (rb *RoleBuilder) Metadata(metadata Metadata) *RoleBuilder {
	rb.v.Metadata = metadata
	return rb
}

// RoleTemplates set the RoleTemplates property for RoleBuilder.
func (rb *RoleBuilder) RoleTemplates(role_templates ...RoleTemplate) *RoleBuilder {
	rb.v.RoleTemplates = append(rb.v.RoleTemplates, role_templates...)
	return rb
}

// RunAs set the RunAs property for RoleBuilder.
func (rb *RoleBuilder) RunAs(run_as ...string) *RoleBuilder {
	rb.v.RunAs = append(rb.v.RunAs, run_as...)
	return rb
}

// TransientMetadata set the TransientMetadata property for RoleBuilder.
func (rb *RoleBuilder) TransientMetadata(transientmetadata TransientMetadataConfig) *RoleBuilder {
	rb.v.TransientMetadata = transientmetadata
	return rb
}
