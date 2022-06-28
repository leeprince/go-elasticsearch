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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/pagerdutyeventtype"
)

// PagerDutyEvent type.
type PagerDutyEvent struct {
	Account       *string                                `json:"account,omitempty"`
	AttachPayload bool                                   `json:"attach_payload"`
	Client        *string                                `json:"client,omitempty"`
	ClientUrl     *string                                `json:"client_url,omitempty"`
	Contexts      []PagerDutyContext                     `json:"contexts,omitempty"`
	Description   string                                 `json:"description"`
	EventType     *pagerdutyeventtype.PagerDutyEventType `json:"event_type,omitempty"`
	IncidentKey   string                                 `json:"incident_key"`
	Proxy         *PagerDutyEventProxy                   `json:"proxy,omitempty"`
}

// PagerDutyEventBuilder holds PagerDutyEvent struct and provides a builder API.
type PagerDutyEventBuilder struct {
	v *PagerDutyEvent
}

// NewPagerDutyEvent provides a builder for the PagerDutyEvent struct.
func NewPagerDutyEvent() *PagerDutyEventBuilder {
	r := PagerDutyEventBuilder{
		&PagerDutyEvent{},
	}

	return &r
}

// Build finalize the chain and returns the PagerDutyEvent struct
func (rb *PagerDutyEventBuilder) Build() PagerDutyEvent {
	return *rb.v
}

// Account set the Account property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) Account(account string) *PagerDutyEventBuilder {
	rb.v.Account = &account
	return rb
}

// AttachPayload set the AttachPayload property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) AttachPayload(attachpayload bool) *PagerDutyEventBuilder {
	rb.v.AttachPayload = attachpayload
	return rb
}

// Client set the Client property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) Client(client string) *PagerDutyEventBuilder {
	rb.v.Client = &client
	return rb
}

// ClientUrl set the ClientUrl property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) ClientUrl(clienturl string) *PagerDutyEventBuilder {
	rb.v.ClientUrl = &clienturl
	return rb
}

// Contexts set the Contexts property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) Contexts(contexts ...PagerDutyContext) *PagerDutyEventBuilder {
	rb.v.Contexts = append(rb.v.Contexts, contexts...)
	return rb
}

// Description set the Description property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) Description(description string) *PagerDutyEventBuilder {
	rb.v.Description = description
	return rb
}

// EventType set the EventType property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) EventType(eventtype pagerdutyeventtype.PagerDutyEventType) *PagerDutyEventBuilder {
	rb.v.EventType = &eventtype
	return rb
}

// IncidentKey set the IncidentKey property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) IncidentKey(incidentkey string) *PagerDutyEventBuilder {
	rb.v.IncidentKey = incidentkey
	return rb
}

// Proxy set the Proxy property for PagerDutyEventBuilder.
func (rb *PagerDutyEventBuilder) Proxy(proxy PagerDutyEventProxy) *PagerDutyEventBuilder {
	rb.v.Proxy = &proxy
	return rb
}
