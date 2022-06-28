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

// TransformSummary type.
type TransformSummary struct {
	CreateTime *EpochMillis `json:"create_time,omitempty"`
	// Description Free text description of the transform.
	Description *string `json:"description,omitempty"`
	// Dest The destination for the transform.
	Dest      Destination `json:"dest"`
	Frequency *Time       `json:"frequency,omitempty"`
	Id        Id          `json:"id"`
	Latest    *Latest     `json:"latest,omitempty"`
	Meta_     *Metadata   `json:"_meta,omitempty"`
	// Pivot The pivot method transforms the data by aggregating and grouping it.
	Pivot *Pivot `json:"pivot,omitempty"`
	// Settings Defines optional transform settings.
	Settings *Settings `json:"settings,omitempty"`
	// Source The source of the data for the transform.
	Source Source `json:"source"`
	// Sync Defines the properties transforms require to run continuously.
	Sync    *SyncContainer `json:"sync,omitempty"`
	Version *VersionString `json:"version,omitempty"`
}

// TransformSummaryBuilder holds TransformSummary struct and provides a builder API.
type TransformSummaryBuilder struct {
	v *TransformSummary
}

// NewTransformSummary provides a builder for the TransformSummary struct.
func NewTransformSummary() *TransformSummaryBuilder {
	r := TransformSummaryBuilder{
		&TransformSummary{},
	}

	return &r
}

// Build finalize the chain and returns the TransformSummary struct
func (rb *TransformSummaryBuilder) Build() TransformSummary {
	return *rb.v
}

// CreateTime set the CreateTime property for TransformSummaryBuilder.
func (rb *TransformSummaryBuilder) CreateTime(createtime EpochMillis) *TransformSummaryBuilder {
	rb.v.CreateTime = &createtime
	return rb
}

// Description Free text description of the transform.
func (rb *TransformSummaryBuilder) Description(description string) *TransformSummaryBuilder {
	rb.v.Description = &description
	return rb
}

// Dest The destination for the transform.
func (rb *TransformSummaryBuilder) Dest(dest Destination) *TransformSummaryBuilder {
	rb.v.Dest = dest
	return rb
}

// Frequency set the Frequency property for TransformSummaryBuilder.
func (rb *TransformSummaryBuilder) Frequency(frequency Time) *TransformSummaryBuilder {
	rb.v.Frequency = &frequency
	return rb
}

// Id set the Id property for TransformSummaryBuilder.
func (rb *TransformSummaryBuilder) Id(id Id) *TransformSummaryBuilder {
	rb.v.Id = id
	return rb
}

// Latest set the Latest property for TransformSummaryBuilder.
func (rb *TransformSummaryBuilder) Latest(latest Latest) *TransformSummaryBuilder {
	rb.v.Latest = &latest
	return rb
}

// Meta_ set the Meta_ property for TransformSummaryBuilder.
func (rb *TransformSummaryBuilder) Meta_(meta_ Metadata) *TransformSummaryBuilder {
	rb.v.Meta_ = &meta_
	return rb
}

// Pivot The pivot method transforms the data by aggregating and grouping it.
func (rb *TransformSummaryBuilder) Pivot(pivot Pivot) *TransformSummaryBuilder {
	rb.v.Pivot = &pivot
	return rb
}

// Settings Defines optional transform settings.
func (rb *TransformSummaryBuilder) Settings(settings Settings) *TransformSummaryBuilder {
	rb.v.Settings = &settings
	return rb
}

// Source The source of the data for the transform.
func (rb *TransformSummaryBuilder) Source(source Source) *TransformSummaryBuilder {
	rb.v.Source = source
	return rb
}

// Sync Defines the properties transforms require to run continuously.
func (rb *TransformSummaryBuilder) Sync(sync SyncContainer) *TransformSummaryBuilder {
	rb.v.Sync = &sync
	return rb
}

// Version set the Version property for TransformSummaryBuilder.
func (rb *TransformSummaryBuilder) Version(version VersionString) *TransformSummaryBuilder {
	rb.v.Version = &version
	return rb
}
