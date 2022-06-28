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

// ProcessorContainer type.
type ProcessorContainer struct {
	Append          *AppendProcessor          `json:"append,omitempty"`
	Attachment      *AttachmentProcessor      `json:"attachment,omitempty"`
	Bytes           *BytesProcessor           `json:"bytes,omitempty"`
	Circle          *CircleProcessor          `json:"circle,omitempty"`
	Convert         *ConvertProcessor         `json:"convert,omitempty"`
	Csv             *CsvProcessor             `json:"csv,omitempty"`
	Date            *DateProcessor            `json:"date,omitempty"`
	DateIndexName   *DateIndexNameProcessor   `json:"date_index_name,omitempty"`
	Dissect         *DissectProcessor         `json:"dissect,omitempty"`
	DotExpander     *DotExpanderProcessor     `json:"dot_expander,omitempty"`
	Drop            *DropProcessor            `json:"drop,omitempty"`
	Enrich          *EnrichProcessor          `json:"enrich,omitempty"`
	Fail            *FailProcessor            `json:"fail,omitempty"`
	Foreach         *ForeachProcessor         `json:"foreach,omitempty"`
	Geoip           *GeoIpProcessor           `json:"geoip,omitempty"`
	Grok            *GrokProcessor            `json:"grok,omitempty"`
	Gsub            *GsubProcessor            `json:"gsub,omitempty"`
	Inference       *InferenceProcessor       `json:"inference,omitempty"`
	Join            *JoinProcessor            `json:"join,omitempty"`
	Json            *JsonProcessor            `json:"json,omitempty"`
	Kv              *KeyValueProcessor        `json:"kv,omitempty"`
	Lowercase       *LowercaseProcessor       `json:"lowercase,omitempty"`
	Pipeline        *PipelineProcessor        `json:"pipeline,omitempty"`
	Remove          *RemoveProcessor          `json:"remove,omitempty"`
	Rename          *RenameProcessor          `json:"rename,omitempty"`
	Script          *Script                   `json:"script,omitempty"`
	Set             *SetProcessor             `json:"set,omitempty"`
	SetSecurityUser *SetSecurityUserProcessor `json:"set_security_user,omitempty"`
	Sort            *SortProcessor            `json:"sort,omitempty"`
	Split           *SplitProcessor           `json:"split,omitempty"`
	Trim            *TrimProcessor            `json:"trim,omitempty"`
	Uppercase       *UppercaseProcessor       `json:"uppercase,omitempty"`
	Urldecode       *UrlDecodeProcessor       `json:"urldecode,omitempty"`
	UserAgent       *UserAgentProcessor       `json:"user_agent,omitempty"`
}

// ProcessorContainerBuilder holds ProcessorContainer struct and provides a builder API.
type ProcessorContainerBuilder struct {
	v *ProcessorContainer
}

// NewProcessorContainer provides a builder for the ProcessorContainer struct.
func NewProcessorContainer() *ProcessorContainerBuilder {
	r := ProcessorContainerBuilder{
		&ProcessorContainer{},
	}

	return &r
}

// Build finalize the chain and returns the ProcessorContainer struct
func (rb *ProcessorContainerBuilder) Build() ProcessorContainer {
	return *rb.v
}

// Append set the Append property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Append(append AppendProcessor) *ProcessorContainerBuilder {
	rb.v.Append = &append
	return rb
}

// Attachment set the Attachment property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Attachment(attachment AttachmentProcessor) *ProcessorContainerBuilder {
	rb.v.Attachment = &attachment
	return rb
}

// Bytes set the Bytes property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Bytes(bytes BytesProcessor) *ProcessorContainerBuilder {
	rb.v.Bytes = &bytes
	return rb
}

// Circle set the Circle property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Circle(circle CircleProcessor) *ProcessorContainerBuilder {
	rb.v.Circle = &circle
	return rb
}

// Convert set the Convert property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Convert(convert ConvertProcessor) *ProcessorContainerBuilder {
	rb.v.Convert = &convert
	return rb
}

// Csv set the Csv property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Csv(csv CsvProcessor) *ProcessorContainerBuilder {
	rb.v.Csv = &csv
	return rb
}

// Date set the Date property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Date(date DateProcessor) *ProcessorContainerBuilder {
	rb.v.Date = &date
	return rb
}

// DateIndexName set the DateIndexName property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) DateIndexName(dateindexname DateIndexNameProcessor) *ProcessorContainerBuilder {
	rb.v.DateIndexName = &dateindexname
	return rb
}

// Dissect set the Dissect property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Dissect(dissect DissectProcessor) *ProcessorContainerBuilder {
	rb.v.Dissect = &dissect
	return rb
}

// DotExpander set the DotExpander property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) DotExpander(dotexpander DotExpanderProcessor) *ProcessorContainerBuilder {
	rb.v.DotExpander = &dotexpander
	return rb
}

// Drop set the Drop property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Drop(drop DropProcessor) *ProcessorContainerBuilder {
	rb.v.Drop = &drop
	return rb
}

// Enrich set the Enrich property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Enrich(enrich EnrichProcessor) *ProcessorContainerBuilder {
	rb.v.Enrich = &enrich
	return rb
}

// Fail set the Fail property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Fail(fail FailProcessor) *ProcessorContainerBuilder {
	rb.v.Fail = &fail
	return rb
}

// Foreach set the Foreach property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Foreach(foreach ForeachProcessor) *ProcessorContainerBuilder {
	rb.v.Foreach = &foreach
	return rb
}

// Geoip set the Geoip property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Geoip(geoip GeoIpProcessor) *ProcessorContainerBuilder {
	rb.v.Geoip = &geoip
	return rb
}

// Grok set the Grok property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Grok(grok GrokProcessor) *ProcessorContainerBuilder {
	rb.v.Grok = &grok
	return rb
}

// Gsub set the Gsub property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Gsub(gsub GsubProcessor) *ProcessorContainerBuilder {
	rb.v.Gsub = &gsub
	return rb
}

// Inference set the Inference property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Inference(inference InferenceProcessor) *ProcessorContainerBuilder {
	rb.v.Inference = &inference
	return rb
}

// Join set the Join property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Join(join JoinProcessor) *ProcessorContainerBuilder {
	rb.v.Join = &join
	return rb
}

// Json set the Json property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Json(json JsonProcessor) *ProcessorContainerBuilder {
	rb.v.Json = &json
	return rb
}

// Kv set the Kv property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Kv(kv KeyValueProcessor) *ProcessorContainerBuilder {
	rb.v.Kv = &kv
	return rb
}

// Lowercase set the Lowercase property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Lowercase(lowercase LowercaseProcessor) *ProcessorContainerBuilder {
	rb.v.Lowercase = &lowercase
	return rb
}

// Pipeline set the Pipeline property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Pipeline(pipeline PipelineProcessor) *ProcessorContainerBuilder {
	rb.v.Pipeline = &pipeline
	return rb
}

// Remove set the Remove property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Remove(remove RemoveProcessor) *ProcessorContainerBuilder {
	rb.v.Remove = &remove
	return rb
}

// Rename set the Rename property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Rename(rename RenameProcessor) *ProcessorContainerBuilder {
	rb.v.Rename = &rename
	return rb
}

// Script set the Script property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Script(script Script) *ProcessorContainerBuilder {
	rb.v.Script = &script
	return rb
}

// Set set the Set property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Set(set SetProcessor) *ProcessorContainerBuilder {
	rb.v.Set = &set
	return rb
}

// SetSecurityUser set the SetSecurityUser property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) SetSecurityUser(setsecurityuser SetSecurityUserProcessor) *ProcessorContainerBuilder {
	rb.v.SetSecurityUser = &setsecurityuser
	return rb
}

// Sort set the Sort property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Sort(sort SortProcessor) *ProcessorContainerBuilder {
	rb.v.Sort = &sort
	return rb
}

// Split set the Split property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Split(split SplitProcessor) *ProcessorContainerBuilder {
	rb.v.Split = &split
	return rb
}

// Trim set the Trim property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Trim(trim TrimProcessor) *ProcessorContainerBuilder {
	rb.v.Trim = &trim
	return rb
}

// Uppercase set the Uppercase property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Uppercase(uppercase UppercaseProcessor) *ProcessorContainerBuilder {
	rb.v.Uppercase = &uppercase
	return rb
}

// Urldecode set the Urldecode property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) Urldecode(urldecode UrlDecodeProcessor) *ProcessorContainerBuilder {
	rb.v.Urldecode = &urldecode
	return rb
}

// UserAgent set the UserAgent property for ProcessorContainerBuilder.
func (rb *ProcessorContainerBuilder) UserAgent(useragent UserAgentProcessor) *ProcessorContainerBuilder {
	rb.v.UserAgent = &useragent
	return rb
}
