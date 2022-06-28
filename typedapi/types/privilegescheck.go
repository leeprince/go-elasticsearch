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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/clusterprivilege"
)

// PrivilegesCheck type.
type PrivilegesCheck struct {
	Application []ApplicationPrivilegesCheck `json:"application,omitempty"`
	// Cluster A list of the cluster privileges that you want to check.
	Cluster []clusterprivilege.ClusterPrivilege `json:"cluster,omitempty"`
	Index   []IndexPrivilegesCheck              `json:"index,omitempty"`
}

// PrivilegesCheckBuilder holds PrivilegesCheck struct and provides a builder API.
type PrivilegesCheckBuilder struct {
	v *PrivilegesCheck
}

// NewPrivilegesCheck provides a builder for the PrivilegesCheck struct.
func NewPrivilegesCheck() *PrivilegesCheckBuilder {
	r := PrivilegesCheckBuilder{
		&PrivilegesCheck{},
	}

	return &r
}

// Build finalize the chain and returns the PrivilegesCheck struct
func (rb *PrivilegesCheckBuilder) Build() PrivilegesCheck {
	return *rb.v
}

// Application set the Application property for PrivilegesCheckBuilder.
func (rb *PrivilegesCheckBuilder) Application(application ...ApplicationPrivilegesCheck) *PrivilegesCheckBuilder {
	rb.v.Application = append(rb.v.Application, application...)
	return rb
}

// Cluster A list of the cluster privileges that you want to check.
func (rb *PrivilegesCheckBuilder) Cluster(cluster ...clusterprivilege.ClusterPrivilege) *PrivilegesCheckBuilder {
	rb.v.Cluster = append(rb.v.Cluster, cluster...)
	return rb
}

// Index set the Index property for PrivilegesCheckBuilder.
func (rb *PrivilegesCheckBuilder) Index(index ...IndexPrivilegesCheck) *PrivilegesCheckBuilder {
	rb.v.Index = append(rb.v.Index, index...)
	return rb
}
