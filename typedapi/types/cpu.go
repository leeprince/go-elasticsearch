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

// Cpu type.
type Cpu struct {
	LoadAverage   map[string]float64 `json:"load_average,omitempty"`
	Percent       *int               `json:"percent,omitempty"`
	Sys           *string            `json:"sys,omitempty"`
	SysInMillis   *int64             `json:"sys_in_millis,omitempty"`
	Total         *string            `json:"total,omitempty"`
	TotalInMillis *int64             `json:"total_in_millis,omitempty"`
	User          *string            `json:"user,omitempty"`
	UserInMillis  *int64             `json:"user_in_millis,omitempty"`
}

// CpuBuilder holds Cpu struct and provides a builder API.
type CpuBuilder struct {
	v *Cpu
}

// NewCpu provides a builder for the Cpu struct.
func NewCpu() *CpuBuilder {
	r := CpuBuilder{
		&Cpu{
			LoadAverage: make(map[string]float64, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Cpu struct
func (rb *CpuBuilder) Build() Cpu {
	return *rb.v
}

// LoadAverage set the LoadAverage property for CpuBuilder.
func (rb *CpuBuilder) LoadAverage(value map[string]float64) *CpuBuilder {
	rb.v.LoadAverage = value
	return rb
}

// Percent set the Percent property for CpuBuilder.
func (rb *CpuBuilder) Percent(percent int) *CpuBuilder {
	rb.v.Percent = &percent
	return rb
}

// Sys set the Sys property for CpuBuilder.
func (rb *CpuBuilder) Sys(sys string) *CpuBuilder {
	rb.v.Sys = &sys
	return rb
}

// SysInMillis set the SysInMillis property for CpuBuilder.
func (rb *CpuBuilder) SysInMillis(sysinmillis int64) *CpuBuilder {
	rb.v.SysInMillis = &sysinmillis
	return rb
}

// Total set the Total property for CpuBuilder.
func (rb *CpuBuilder) Total(total string) *CpuBuilder {
	rb.v.Total = &total
	return rb
}

// TotalInMillis set the TotalInMillis property for CpuBuilder.
func (rb *CpuBuilder) TotalInMillis(totalinmillis int64) *CpuBuilder {
	rb.v.TotalInMillis = &totalinmillis
	return rb
}

// User set the User property for CpuBuilder.
func (rb *CpuBuilder) User(user string) *CpuBuilder {
	rb.v.User = &user
	return rb
}

// UserInMillis set the UserInMillis property for CpuBuilder.
func (rb *CpuBuilder) UserInMillis(userinmillis int64) *CpuBuilder {
	rb.v.UserInMillis = &userinmillis
	return rb
}
