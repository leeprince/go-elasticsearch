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

// PagerDutyEventProxy type.
type PagerDutyEventProxy struct {
	Host *Host `json:"host,omitempty"`
	Port *int  `json:"port,omitempty"`
}

// PagerDutyEventProxyBuilder holds PagerDutyEventProxy struct and provides a builder API.
type PagerDutyEventProxyBuilder struct {
	v *PagerDutyEventProxy
}

// NewPagerDutyEventProxy provides a builder for the PagerDutyEventProxy struct.
func NewPagerDutyEventProxy() *PagerDutyEventProxyBuilder {
	r := PagerDutyEventProxyBuilder{
		&PagerDutyEventProxy{},
	}

	return &r
}

// Build finalize the chain and returns the PagerDutyEventProxy struct
func (rb *PagerDutyEventProxyBuilder) Build() PagerDutyEventProxy {
	return *rb.v
}

// Host set the Host property for PagerDutyEventProxyBuilder.
func (rb *PagerDutyEventProxyBuilder) Host(host Host) *PagerDutyEventProxyBuilder {
	rb.v.Host = &host
	return rb
}

// Port set the Port property for PagerDutyEventProxyBuilder.
func (rb *PagerDutyEventProxyBuilder) Port(port int) *PagerDutyEventProxyBuilder {
	rb.v.Port = &port
	return rb
}
