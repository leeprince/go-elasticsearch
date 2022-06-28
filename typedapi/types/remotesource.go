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

// RemoteSource type.
type RemoteSource struct {
	ConnectTimeout *Time             `json:"connect_timeout,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	Host           Host              `json:"host"`
	Password       *Password         `json:"password,omitempty"`
	SocketTimeout  *Time             `json:"socket_timeout,omitempty"`
	Username       *Username         `json:"username,omitempty"`
}

// RemoteSourceBuilder holds RemoteSource struct and provides a builder API.
type RemoteSourceBuilder struct {
	v *RemoteSource
}

// NewRemoteSource provides a builder for the RemoteSource struct.
func NewRemoteSource() *RemoteSourceBuilder {
	r := RemoteSourceBuilder{
		&RemoteSource{
			Headers: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the RemoteSource struct
func (rb *RemoteSourceBuilder) Build() RemoteSource {
	return *rb.v
}

// ConnectTimeout set the ConnectTimeout property for RemoteSourceBuilder.
func (rb *RemoteSourceBuilder) ConnectTimeout(connecttimeout Time) *RemoteSourceBuilder {
	rb.v.ConnectTimeout = &connecttimeout
	return rb
}

// Headers set the Headers property for RemoteSourceBuilder.
func (rb *RemoteSourceBuilder) Headers(value map[string]string) *RemoteSourceBuilder {
	rb.v.Headers = value
	return rb
}

// Host set the Host property for RemoteSourceBuilder.
func (rb *RemoteSourceBuilder) Host(host Host) *RemoteSourceBuilder {
	rb.v.Host = host
	return rb
}

// Password set the Password property for RemoteSourceBuilder.
func (rb *RemoteSourceBuilder) Password(password Password) *RemoteSourceBuilder {
	rb.v.Password = &password
	return rb
}

// SocketTimeout set the SocketTimeout property for RemoteSourceBuilder.
func (rb *RemoteSourceBuilder) SocketTimeout(sockettimeout Time) *RemoteSourceBuilder {
	rb.v.SocketTimeout = &sockettimeout
	return rb
}

// Username set the Username property for RemoteSourceBuilder.
func (rb *RemoteSourceBuilder) Username(username Username) *RemoteSourceBuilder {
	rb.v.Username = &username
	return rb
}
