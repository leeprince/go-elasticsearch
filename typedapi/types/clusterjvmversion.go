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

// ClusterJvmVersion type.
type ClusterJvmVersion struct {
	BundledJdk      bool          `json:"bundled_jdk"`
	Count           int           `json:"count"`
	UsingBundledJdk bool          `json:"using_bundled_jdk"`
	Version         VersionString `json:"version"`
	VmName          string        `json:"vm_name"`
	VmVendor        string        `json:"vm_vendor"`
	VmVersion       VersionString `json:"vm_version"`
}

// ClusterJvmVersionBuilder holds ClusterJvmVersion struct and provides a builder API.
type ClusterJvmVersionBuilder struct {
	v *ClusterJvmVersion
}

// NewClusterJvmVersion provides a builder for the ClusterJvmVersion struct.
func NewClusterJvmVersion() *ClusterJvmVersionBuilder {
	r := ClusterJvmVersionBuilder{
		&ClusterJvmVersion{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterJvmVersion struct
func (rb *ClusterJvmVersionBuilder) Build() ClusterJvmVersion {
	return *rb.v
}

// BundledJdk set the BundledJdk property for ClusterJvmVersionBuilder.
func (rb *ClusterJvmVersionBuilder) BundledJdk(bundledjdk bool) *ClusterJvmVersionBuilder {
	rb.v.BundledJdk = bundledjdk
	return rb
}

// Count set the Count property for ClusterJvmVersionBuilder.
func (rb *ClusterJvmVersionBuilder) Count(count int) *ClusterJvmVersionBuilder {
	rb.v.Count = count
	return rb
}

// UsingBundledJdk set the UsingBundledJdk property for ClusterJvmVersionBuilder.
func (rb *ClusterJvmVersionBuilder) UsingBundledJdk(usingbundledjdk bool) *ClusterJvmVersionBuilder {
	rb.v.UsingBundledJdk = usingbundledjdk
	return rb
}

// Version set the Version property for ClusterJvmVersionBuilder.
func (rb *ClusterJvmVersionBuilder) Version(version VersionString) *ClusterJvmVersionBuilder {
	rb.v.Version = version
	return rb
}

// VmName set the VmName property for ClusterJvmVersionBuilder.
func (rb *ClusterJvmVersionBuilder) VmName(vmname string) *ClusterJvmVersionBuilder {
	rb.v.VmName = vmname
	return rb
}

// VmVendor set the VmVendor property for ClusterJvmVersionBuilder.
func (rb *ClusterJvmVersionBuilder) VmVendor(vmvendor string) *ClusterJvmVersionBuilder {
	rb.v.VmVendor = vmvendor
	return rb
}

// VmVersion set the VmVersion property for ClusterJvmVersionBuilder.
func (rb *ClusterJvmVersionBuilder) VmVersion(vmversion VersionString) *ClusterJvmVersionBuilder {
	rb.v.VmVersion = vmversion
	return rb
}
