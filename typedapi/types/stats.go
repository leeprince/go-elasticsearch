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

// Stats type.
type Stats struct {
	AdaptiveSelection map[string]AdaptiveSelection `json:"adaptive_selection,omitempty"`
	Attributes        map[Field]string             `json:"attributes,omitempty"`
	Breakers          map[string]Breaker           `json:"breakers,omitempty"`
	Discovery         *Discovery                   `json:"discovery,omitempty"`
	Fs                *FileSystem                  `json:"fs,omitempty"`
	Host              *Host                        `json:"host,omitempty"`
	Http              *Http                        `json:"http,omitempty"`
	IndexingPressure  *IndexingPressure            `json:"indexing_pressure,omitempty"`
	Indices           *ShardStats                  `json:"indices,omitempty"`
	Ingest            *Ingest                      `json:"ingest,omitempty"`
	Ip                []Ip                         `json:"ip,omitempty"`
	Jvm               *Jvm                         `json:"jvm,omitempty"`
	Name              *Name                        `json:"name,omitempty"`
	Os                *OperatingSystem             `json:"os,omitempty"`
	Process           *Process                     `json:"process,omitempty"`
	Roles             *NodeRoles                   `json:"roles,omitempty"`
	Script            *Scripting                   `json:"script,omitempty"`
	ScriptCache       map[string][]ScriptCache     `json:"script_cache,omitempty"`
	ThreadPool        map[string]ThreadCount       `json:"thread_pool,omitempty"`
	Timestamp         *int64                       `json:"timestamp,omitempty"`
	Transport         *Transport                   `json:"transport,omitempty"`
	TransportAddress  *TransportAddress            `json:"transport_address,omitempty"`
}

// StatsBuilder holds Stats struct and provides a builder API.
type StatsBuilder struct {
	v *Stats
}

// NewStats provides a builder for the Stats struct.
func NewStats() *StatsBuilder {
	r := StatsBuilder{
		&Stats{
			AdaptiveSelection: make(map[string]AdaptiveSelection, 0),
			Attributes:        make(map[Field]string, 0),
			Breakers:          make(map[string]Breaker, 0),
			ScriptCache:       make(map[string][]ScriptCache, 0),
			ThreadPool:        make(map[string]ThreadCount, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the Stats struct
func (rb *StatsBuilder) Build() Stats {
	return *rb.v
}

// AdaptiveSelection set the AdaptiveSelection property for StatsBuilder.
func (rb *StatsBuilder) AdaptiveSelection(value map[string]AdaptiveSelection) *StatsBuilder {
	rb.v.AdaptiveSelection = value
	return rb
}

// Attributes set the Attributes property for StatsBuilder.
func (rb *StatsBuilder) Attributes(value map[Field]string) *StatsBuilder {
	rb.v.Attributes = value
	return rb
}

// Breakers set the Breakers property for StatsBuilder.
func (rb *StatsBuilder) Breakers(value map[string]Breaker) *StatsBuilder {
	rb.v.Breakers = value
	return rb
}

// Discovery set the Discovery property for StatsBuilder.
func (rb *StatsBuilder) Discovery(discovery Discovery) *StatsBuilder {
	rb.v.Discovery = &discovery
	return rb
}

// Fs set the Fs property for StatsBuilder.
func (rb *StatsBuilder) Fs(fs FileSystem) *StatsBuilder {
	rb.v.Fs = &fs
	return rb
}

// Host set the Host property for StatsBuilder.
func (rb *StatsBuilder) Host(host Host) *StatsBuilder {
	rb.v.Host = &host
	return rb
}

// Http set the Http property for StatsBuilder.
func (rb *StatsBuilder) Http(http Http) *StatsBuilder {
	rb.v.Http = &http
	return rb
}

// IndexingPressure set the IndexingPressure property for StatsBuilder.
func (rb *StatsBuilder) IndexingPressure(indexingpressure IndexingPressure) *StatsBuilder {
	rb.v.IndexingPressure = &indexingpressure
	return rb
}

// Indices set the Indices property for StatsBuilder.
func (rb *StatsBuilder) Indices(indices ShardStats) *StatsBuilder {
	rb.v.Indices = &indices
	return rb
}

// Ingest set the Ingest property for StatsBuilder.
func (rb *StatsBuilder) Ingest(ingest Ingest) *StatsBuilder {
	rb.v.Ingest = &ingest
	return rb
}

// Ip set the Ip property for StatsBuilder.
func (rb *StatsBuilder) Ip(arg []Ip) *StatsBuilder {
	rb.v.Ip = arg
	return rb
}

// Jvm set the Jvm property for StatsBuilder.
func (rb *StatsBuilder) Jvm(jvm Jvm) *StatsBuilder {
	rb.v.Jvm = &jvm
	return rb
}

// Name set the Name property for StatsBuilder.
func (rb *StatsBuilder) Name(name Name) *StatsBuilder {
	rb.v.Name = &name
	return rb
}

// Os set the Os property for StatsBuilder.
func (rb *StatsBuilder) Os(os OperatingSystem) *StatsBuilder {
	rb.v.Os = &os
	return rb
}

// Process set the Process property for StatsBuilder.
func (rb *StatsBuilder) Process(process Process) *StatsBuilder {
	rb.v.Process = &process
	return rb
}

// Roles set the Roles property for StatsBuilder.
func (rb *StatsBuilder) Roles(roles NodeRoles) *StatsBuilder {
	rb.v.Roles = &roles
	return rb
}

// Script set the Script property for StatsBuilder.
func (rb *StatsBuilder) Script(script Scripting) *StatsBuilder {
	rb.v.Script = &script
	return rb
}

// ScriptCache set the ScriptCache property for StatsBuilder.
//TODO: Union in map Not implemented yet: UnionOf
func (rb *StatsBuilder) ScriptCache(value map[string][]ScriptCache) *StatsBuilder {
	rb.v.ScriptCache = value
	return rb
}

// ThreadPool set the ThreadPool property for StatsBuilder.
func (rb *StatsBuilder) ThreadPool(value map[string]ThreadCount) *StatsBuilder {
	rb.v.ThreadPool = value
	return rb
}

// Timestamp set the Timestamp property for StatsBuilder.
func (rb *StatsBuilder) Timestamp(timestamp int64) *StatsBuilder {
	rb.v.Timestamp = &timestamp
	return rb
}

// Transport set the Transport property for StatsBuilder.
func (rb *StatsBuilder) Transport(transport Transport) *StatsBuilder {
	rb.v.Transport = &transport
	return rb
}

// TransportAddress set the TransportAddress property for StatsBuilder.
func (rb *StatsBuilder) TransportAddress(transportaddress TransportAddress) *StatsBuilder {
	rb.v.TransportAddress = &transportaddress
	return rb
}
