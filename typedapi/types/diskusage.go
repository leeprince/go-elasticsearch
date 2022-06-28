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

// DiskUsage type.
type DiskUsage struct {
	FreeBytes       int64   `json:"free_bytes"`
	FreeDiskPercent float64 `json:"free_disk_percent"`
	Path            string  `json:"path"`
	TotalBytes      int64   `json:"total_bytes"`
	UsedBytes       int64   `json:"used_bytes"`
	UsedDiskPercent float64 `json:"used_disk_percent"`
}

// DiskUsageBuilder holds DiskUsage struct and provides a builder API.
type DiskUsageBuilder struct {
	v *DiskUsage
}

// NewDiskUsage provides a builder for the DiskUsage struct.
func NewDiskUsage() *DiskUsageBuilder {
	r := DiskUsageBuilder{
		&DiskUsage{},
	}

	return &r
}

// Build finalize the chain and returns the DiskUsage struct
func (rb *DiskUsageBuilder) Build() DiskUsage {
	return *rb.v
}

// FreeBytes set the FreeBytes property for DiskUsageBuilder.
func (rb *DiskUsageBuilder) FreeBytes(freebytes int64) *DiskUsageBuilder {
	rb.v.FreeBytes = freebytes
	return rb
}

// FreeDiskPercent set the FreeDiskPercent property for DiskUsageBuilder.
func (rb *DiskUsageBuilder) FreeDiskPercent(freediskpercent float64) *DiskUsageBuilder {
	rb.v.FreeDiskPercent = freediskpercent
	return rb
}

// Path set the Path property for DiskUsageBuilder.
func (rb *DiskUsageBuilder) Path(path string) *DiskUsageBuilder {
	rb.v.Path = path
	return rb
}

// TotalBytes set the TotalBytes property for DiskUsageBuilder.
func (rb *DiskUsageBuilder) TotalBytes(totalbytes int64) *DiskUsageBuilder {
	rb.v.TotalBytes = totalbytes
	return rb
}

// UsedBytes set the UsedBytes property for DiskUsageBuilder.
func (rb *DiskUsageBuilder) UsedBytes(usedbytes int64) *DiskUsageBuilder {
	rb.v.UsedBytes = usedbytes
	return rb
}

// UsedDiskPercent set the UsedDiskPercent property for DiskUsageBuilder.
func (rb *DiskUsageBuilder) UsedDiskPercent(useddiskpercent float64) *DiskUsageBuilder {
	rb.v.UsedDiskPercent = useddiskpercent
	return rb
}
