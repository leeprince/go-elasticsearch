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

// NodeInfoSettingsIngest type.
type NodeInfoSettingsIngest struct {
	Append          *NodeInfoIngestInfo `json:"append,omitempty"`
	Attachment      *NodeInfoIngestInfo `json:"attachment,omitempty"`
	Bytes           *NodeInfoIngestInfo `json:"bytes,omitempty"`
	Circle          *NodeInfoIngestInfo `json:"circle,omitempty"`
	Convert         *NodeInfoIngestInfo `json:"convert,omitempty"`
	Csv             *NodeInfoIngestInfo `json:"csv,omitempty"`
	Date            *NodeInfoIngestInfo `json:"date,omitempty"`
	DateIndexName   *NodeInfoIngestInfo `json:"date_index_name,omitempty"`
	Dissect         *NodeInfoIngestInfo `json:"dissect,omitempty"`
	DotExpander     *NodeInfoIngestInfo `json:"dot_expander,omitempty"`
	Drop            *NodeInfoIngestInfo `json:"drop,omitempty"`
	Enrich          *NodeInfoIngestInfo `json:"enrich,omitempty"`
	Fail            *NodeInfoIngestInfo `json:"fail,omitempty"`
	Foreach         *NodeInfoIngestInfo `json:"foreach,omitempty"`
	Geoip           *NodeInfoIngestInfo `json:"geoip,omitempty"`
	Grok            *NodeInfoIngestInfo `json:"grok,omitempty"`
	Gsub            *NodeInfoIngestInfo `json:"gsub,omitempty"`
	Inference       *NodeInfoIngestInfo `json:"inference,omitempty"`
	Join            *NodeInfoIngestInfo `json:"join,omitempty"`
	Json            *NodeInfoIngestInfo `json:"json,omitempty"`
	Kv              *NodeInfoIngestInfo `json:"kv,omitempty"`
	Lowercase       *NodeInfoIngestInfo `json:"lowercase,omitempty"`
	Pipeline        *NodeInfoIngestInfo `json:"pipeline,omitempty"`
	Remove          *NodeInfoIngestInfo `json:"remove,omitempty"`
	Rename          *NodeInfoIngestInfo `json:"rename,omitempty"`
	Script          *NodeInfoIngestInfo `json:"script,omitempty"`
	Set             *NodeInfoIngestInfo `json:"set,omitempty"`
	SetSecurityUser *NodeInfoIngestInfo `json:"set_security_user,omitempty"`
	Sort            *NodeInfoIngestInfo `json:"sort,omitempty"`
	Split           *NodeInfoIngestInfo `json:"split,omitempty"`
	Trim            *NodeInfoIngestInfo `json:"trim,omitempty"`
	Uppercase       *NodeInfoIngestInfo `json:"uppercase,omitempty"`
	Urldecode       *NodeInfoIngestInfo `json:"urldecode,omitempty"`
	UserAgent       *NodeInfoIngestInfo `json:"user_agent,omitempty"`
}

// NodeInfoSettingsIngestBuilder holds NodeInfoSettingsIngest struct and provides a builder API.
type NodeInfoSettingsIngestBuilder struct {
	v *NodeInfoSettingsIngest
}

// NewNodeInfoSettingsIngest provides a builder for the NodeInfoSettingsIngest struct.
func NewNodeInfoSettingsIngest() *NodeInfoSettingsIngestBuilder {
	r := NodeInfoSettingsIngestBuilder{
		&NodeInfoSettingsIngest{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsIngest struct
func (rb *NodeInfoSettingsIngestBuilder) Build() NodeInfoSettingsIngest {
	return *rb.v
}

// Append set the Append property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Append(append NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Append = &append
	return rb
}

// Attachment set the Attachment property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Attachment(attachment NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Attachment = &attachment
	return rb
}

// Bytes set the Bytes property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Bytes(bytes NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Bytes = &bytes
	return rb
}

// Circle set the Circle property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Circle(circle NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Circle = &circle
	return rb
}

// Convert set the Convert property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Convert(convert NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Convert = &convert
	return rb
}

// Csv set the Csv property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Csv(csv NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Csv = &csv
	return rb
}

// Date set the Date property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Date(date NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Date = &date
	return rb
}

// DateIndexName set the DateIndexName property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) DateIndexName(dateindexname NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.DateIndexName = &dateindexname
	return rb
}

// Dissect set the Dissect property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Dissect(dissect NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Dissect = &dissect
	return rb
}

// DotExpander set the DotExpander property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) DotExpander(dotexpander NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.DotExpander = &dotexpander
	return rb
}

// Drop set the Drop property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Drop(drop NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Drop = &drop
	return rb
}

// Enrich set the Enrich property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Enrich(enrich NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Enrich = &enrich
	return rb
}

// Fail set the Fail property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Fail(fail NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Fail = &fail
	return rb
}

// Foreach set the Foreach property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Foreach(foreach NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Foreach = &foreach
	return rb
}

// Geoip set the Geoip property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Geoip(geoip NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Geoip = &geoip
	return rb
}

// Grok set the Grok property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Grok(grok NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Grok = &grok
	return rb
}

// Gsub set the Gsub property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Gsub(gsub NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Gsub = &gsub
	return rb
}

// Inference set the Inference property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Inference(inference NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Inference = &inference
	return rb
}

// Join set the Join property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Join(join NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Join = &join
	return rb
}

// Json set the Json property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Json(json NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Json = &json
	return rb
}

// Kv set the Kv property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Kv(kv NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Kv = &kv
	return rb
}

// Lowercase set the Lowercase property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Lowercase(lowercase NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Lowercase = &lowercase
	return rb
}

// Pipeline set the Pipeline property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Pipeline(pipeline NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Pipeline = &pipeline
	return rb
}

// Remove set the Remove property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Remove(remove NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Remove = &remove
	return rb
}

// Rename set the Rename property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Rename(rename NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Rename = &rename
	return rb
}

// Script set the Script property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Script(script NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Script = &script
	return rb
}

// Set set the Set property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Set(set NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Set = &set
	return rb
}

// SetSecurityUser set the SetSecurityUser property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) SetSecurityUser(setsecurityuser NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.SetSecurityUser = &setsecurityuser
	return rb
}

// Sort set the Sort property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Sort(sort NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Sort = &sort
	return rb
}

// Split set the Split property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Split(split NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Split = &split
	return rb
}

// Trim set the Trim property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Trim(trim NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Trim = &trim
	return rb
}

// Uppercase set the Uppercase property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Uppercase(uppercase NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Uppercase = &uppercase
	return rb
}

// Urldecode set the Urldecode property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) Urldecode(urldecode NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.Urldecode = &urldecode
	return rb
}

// UserAgent set the UserAgent property for NodeInfoSettingsIngestBuilder.
func (rb *NodeInfoSettingsIngestBuilder) UserAgent(useragent NodeInfoIngestInfo) *NodeInfoSettingsIngestBuilder {
	rb.v.UserAgent = &useragent
	return rb
}
