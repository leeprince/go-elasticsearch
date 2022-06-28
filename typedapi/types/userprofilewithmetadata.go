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

// UserProfileWithMetadata type.
type UserProfileWithMetadata struct {
	Data             map[string]interface{} `json:"data"`
	Doc_             UserProfileHitMetadata `json:"_doc"`
	Enabled          *bool                  `json:"enabled,omitempty"`
	Labels           map[string]interface{} `json:"labels"`
	LastSynchronized int64                  `json:"last_synchronized"`
	Uid              UserProfileId          `json:"uid"`
	User             UserProfileUser        `json:"user"`
}

// UserProfileWithMetadataBuilder holds UserProfileWithMetadata struct and provides a builder API.
type UserProfileWithMetadataBuilder struct {
	v *UserProfileWithMetadata
}

// NewUserProfileWithMetadata provides a builder for the UserProfileWithMetadata struct.
func NewUserProfileWithMetadata() *UserProfileWithMetadataBuilder {
	r := UserProfileWithMetadataBuilder{
		&UserProfileWithMetadata{
			Data:   make(map[string]interface{}, 0),
			Labels: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the UserProfileWithMetadata struct
func (rb *UserProfileWithMetadataBuilder) Build() UserProfileWithMetadata {
	return *rb.v
}

// Data set the Data property for UserProfileWithMetadataBuilder.
func (rb *UserProfileWithMetadataBuilder) Data(value map[string]interface{}) *UserProfileWithMetadataBuilder {
	rb.v.Data = value
	return rb
}

// Doc_ set the Doc_ property for UserProfileWithMetadataBuilder.
func (rb *UserProfileWithMetadataBuilder) Doc_(doc_ UserProfileHitMetadata) *UserProfileWithMetadataBuilder {
	rb.v.Doc_ = doc_
	return rb
}

// Enabled set the Enabled property for UserProfileWithMetadataBuilder.
func (rb *UserProfileWithMetadataBuilder) Enabled(enabled bool) *UserProfileWithMetadataBuilder {
	rb.v.Enabled = &enabled
	return rb
}

// Labels set the Labels property for UserProfileWithMetadataBuilder.
func (rb *UserProfileWithMetadataBuilder) Labels(value map[string]interface{}) *UserProfileWithMetadataBuilder {
	rb.v.Labels = value
	return rb
}

// LastSynchronized set the LastSynchronized property for UserProfileWithMetadataBuilder.
func (rb *UserProfileWithMetadataBuilder) LastSynchronized(lastsynchronized int64) *UserProfileWithMetadataBuilder {
	rb.v.LastSynchronized = lastsynchronized
	return rb
}

// Uid set the Uid property for UserProfileWithMetadataBuilder.
func (rb *UserProfileWithMetadataBuilder) Uid(uid UserProfileId) *UserProfileWithMetadataBuilder {
	rb.v.Uid = uid
	return rb
}

// User set the User property for UserProfileWithMetadataBuilder.
func (rb *UserProfileWithMetadataBuilder) User(user UserProfileUser) *UserProfileWithMetadataBuilder {
	rb.v.User = user
	return rb
}
