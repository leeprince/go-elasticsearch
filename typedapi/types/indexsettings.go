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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/indexcheckonstartup"
)

// IndexSettings type.
type IndexSettings struct {
	Analysis *IndexSettingsAnalysis `json:"analysis,omitempty"`
	// Analyze Settings to define analyzers, tokenizers, token filters and character
	// filters.
	Analyze            *SettingsAnalyze                         `json:"analyze,omitempty"`
	AutoExpandReplicas *string                                  `json:"auto_expand_replicas,omitempty"`
	Blocks             *IndexSettingBlocks                      `json:"blocks,omitempty"`
	CheckOnStartup     *indexcheckonstartup.IndexCheckOnStartup `json:"check_on_startup,omitempty"`
	Codec              *string                                  `json:"codec,omitempty"`
	CreationDate       *DateString                              `json:"creation_date,omitempty"`
	CreationDateString *DateString                              `json:"creation_date_string,omitempty"`
	DefaultPipeline    *PipelineName                            `json:"default_pipeline,omitempty"`
	FinalPipeline      *PipelineName                            `json:"final_pipeline,omitempty"`
	Format             string                                   `json:"format,omitempty"`
	GcDeletes          *Time                                    `json:"gc_deletes,omitempty"`
	Hidden             string                                   `json:"hidden,omitempty"`
	Highlight          *SettingsHighlight                       `json:"highlight,omitempty"`
	Index              *IndexSettings                           `json:"index,omitempty"`
	// IndexingPressure Configure indexing back pressure limits.
	IndexingPressure              *IndexingPressure       `json:"indexing_pressure,omitempty"`
	IndexingSlowlog               *SlowlogSettings        `json:"indexing.slowlog,omitempty"`
	Lifecycle                     *IndexSettingsLifecycle `json:"lifecycle,omitempty"`
	LoadFixedBitsetFiltersEagerly *bool                   `json:"load_fixed_bitset_filters_eagerly,omitempty"`
	// Mappings Enable or disable dynamic mapping for an index.
	Mappings                *MappingLimitSettings `json:"mappings,omitempty"`
	MaxDocvalueFieldsSearch *int                  `json:"max_docvalue_fields_search,omitempty"`
	MaxInnerResultWindow    *int                  `json:"max_inner_result_window,omitempty"`
	MaxNgramDiff            *int                  `json:"max_ngram_diff,omitempty"`
	MaxRefreshListeners     *int                  `json:"max_refresh_listeners,omitempty"`
	MaxRegexLength          *int                  `json:"max_regex_length,omitempty"`
	MaxRescoreWindow        *int                  `json:"max_rescore_window,omitempty"`
	MaxResultWindow         *int                  `json:"max_result_window,omitempty"`
	MaxScriptFields         *int                  `json:"max_script_fields,omitempty"`
	MaxShingleDiff          *int                  `json:"max_shingle_diff,omitempty"`
	MaxSlicesPerScroll      *int                  `json:"max_slices_per_scroll,omitempty"`
	MaxTermsCount           *int                  `json:"max_terms_count,omitempty"`
	Merge                   *Merge                `json:"merge,omitempty"`
	Mode                    *string               `json:"mode,omitempty"`
	NumberOfReplicas        string                `json:"number_of_replicas,omitempty"`
	NumberOfRoutingShards   *int                  `json:"number_of_routing_shards,omitempty"`
	NumberOfShards          string                `json:"number_of_shards,omitempty"`
	Priority                string                `json:"priority,omitempty"`
	ProvidedName            *Name                 `json:"provided_name,omitempty"`
	Queries                 *Queries              `json:"queries,omitempty"`
	QueryString             *SettingsQueryString  `json:"query_string,omitempty"`
	RefreshInterval         *Time                 `json:"refresh_interval,omitempty"`
	Routing                 *IndexRouting         `json:"routing,omitempty"`
	RoutingPartitionSize    *int                  `json:"routing_partition_size,omitempty"`
	RoutingPath             []string              `json:"routing_path,omitempty"`
	Search                  *SettingsSearch       `json:"search,omitempty"`
	Settings                *IndexSettings        `json:"settings,omitempty"`
	Shards                  *int                  `json:"shards,omitempty"`
	// Similarity Configure custom similarity settings to customize how search results are
	// scored.
	Similarity  *SettingsSimilarity `json:"similarity,omitempty"`
	SoftDeletes *SoftDeletes        `json:"soft_deletes,omitempty"`
	Sort        *IndexSegmentSort   `json:"sort,omitempty"`
	// Store The store module allows you to control how index data is stored and accessed
	// on disk.
	Store               *Storage                 `json:"store,omitempty"`
	TimeSeries          *IndexSettingsTimeSeries `json:"time_series,omitempty"`
	TopMetricsMaxSize   *int                     `json:"top_metrics_max_size,omitempty"`
	Translog            *Translog                `json:"translog,omitempty"`
	Uuid                *Uuid                    `json:"uuid,omitempty"`
	VerifiedBeforeClose string                   `json:"verified_before_close,omitempty"`
	Version             *IndexVersioning         `json:"version,omitempty"`
}

// IndexSettingsBuilder holds IndexSettings struct and provides a builder API.
type IndexSettingsBuilder struct {
	v *IndexSettings
}

// NewIndexSettings provides a builder for the IndexSettings struct.
func NewIndexSettings() *IndexSettingsBuilder {
	r := IndexSettingsBuilder{
		&IndexSettings{},
	}

	return &r
}

// Build finalize the chain and returns the IndexSettings struct
func (rb *IndexSettingsBuilder) Build() IndexSettings {
	return *rb.v
}

// Analysis set the Analysis property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Analysis(analysis IndexSettingsAnalysis) *IndexSettingsBuilder {
	rb.v.Analysis = &analysis
	return rb
}

// Analyze Settings to define analyzers, tokenizers, token filters and character
// filters.
func (rb *IndexSettingsBuilder) Analyze(analyze SettingsAnalyze) *IndexSettingsBuilder {
	rb.v.Analyze = &analyze
	return rb
}

// AutoExpandReplicas set the AutoExpandReplicas property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) AutoExpandReplicas(autoexpandreplicas string) *IndexSettingsBuilder {
	rb.v.AutoExpandReplicas = &autoexpandreplicas
	return rb
}

// Blocks set the Blocks property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Blocks(blocks IndexSettingBlocks) *IndexSettingsBuilder {
	rb.v.Blocks = &blocks
	return rb
}

// CheckOnStartup set the CheckOnStartup property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) CheckOnStartup(checkonstartup indexcheckonstartup.IndexCheckOnStartup) *IndexSettingsBuilder {
	rb.v.CheckOnStartup = &checkonstartup
	return rb
}

// Codec set the Codec property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Codec(codec string) *IndexSettingsBuilder {
	rb.v.Codec = &codec
	return rb
}

// CreationDate set the CreationDate property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) CreationDate(creationdate DateString) *IndexSettingsBuilder {
	rb.v.CreationDate = &creationdate
	return rb
}

// CreationDateString set the CreationDateString property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) CreationDateString(creationdatestring DateString) *IndexSettingsBuilder {
	rb.v.CreationDateString = &creationdatestring
	return rb
}

// DefaultPipeline set the DefaultPipeline property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) DefaultPipeline(defaultpipeline PipelineName) *IndexSettingsBuilder {
	rb.v.DefaultPipeline = &defaultpipeline
	return rb
}

// FinalPipeline set the FinalPipeline property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) FinalPipeline(finalpipeline PipelineName) *IndexSettingsBuilder {
	rb.v.FinalPipeline = &finalpipeline
	return rb
}

// Format set the Format property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Format(arg string) *IndexSettingsBuilder {
	rb.v.Format = arg
	return rb
}

// GcDeletes set the GcDeletes property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) GcDeletes(gcdeletes Time) *IndexSettingsBuilder {
	rb.v.GcDeletes = &gcdeletes
	return rb
}

// Hidden set the Hidden property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Hidden(arg string) *IndexSettingsBuilder {
	rb.v.Hidden = arg
	return rb
}

// Highlight set the Highlight property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Highlight(highlight SettingsHighlight) *IndexSettingsBuilder {
	rb.v.Highlight = &highlight
	return rb
}

// Index set the Index property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Index(index IndexSettings) *IndexSettingsBuilder {
	rb.v.Index = &index
	return rb
}

// IndexingPressure Configure indexing back pressure limits.
func (rb *IndexSettingsBuilder) IndexingPressure(indexingpressure IndexingPressure) *IndexSettingsBuilder {
	rb.v.IndexingPressure = &indexingpressure
	return rb
}

// IndexingSlowlog set the IndexingSlowlog property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) IndexingSlowlog(indexingslowlog SlowlogSettings) *IndexSettingsBuilder {
	rb.v.IndexingSlowlog = &indexingslowlog
	return rb
}

// Lifecycle set the Lifecycle property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Lifecycle(lifecycle IndexSettingsLifecycle) *IndexSettingsBuilder {
	rb.v.Lifecycle = &lifecycle
	return rb
}

// LoadFixedBitsetFiltersEagerly set the LoadFixedBitsetFiltersEagerly property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) LoadFixedBitsetFiltersEagerly(loadfixedbitsetfilterseagerly bool) *IndexSettingsBuilder {
	rb.v.LoadFixedBitsetFiltersEagerly = &loadfixedbitsetfilterseagerly
	return rb
}

// Mappings Enable or disable dynamic mapping for an index.
func (rb *IndexSettingsBuilder) Mappings(mappings MappingLimitSettings) *IndexSettingsBuilder {
	rb.v.Mappings = &mappings
	return rb
}

// MaxDocvalueFieldsSearch set the MaxDocvalueFieldsSearch property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxDocvalueFieldsSearch(maxdocvaluefieldssearch int) *IndexSettingsBuilder {
	rb.v.MaxDocvalueFieldsSearch = &maxdocvaluefieldssearch
	return rb
}

// MaxInnerResultWindow set the MaxInnerResultWindow property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxInnerResultWindow(maxinnerresultwindow int) *IndexSettingsBuilder {
	rb.v.MaxInnerResultWindow = &maxinnerresultwindow
	return rb
}

// MaxNgramDiff set the MaxNgramDiff property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxNgramDiff(maxngramdiff int) *IndexSettingsBuilder {
	rb.v.MaxNgramDiff = &maxngramdiff
	return rb
}

// MaxRefreshListeners set the MaxRefreshListeners property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxRefreshListeners(maxrefreshlisteners int) *IndexSettingsBuilder {
	rb.v.MaxRefreshListeners = &maxrefreshlisteners
	return rb
}

// MaxRegexLength set the MaxRegexLength property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxRegexLength(maxregexlength int) *IndexSettingsBuilder {
	rb.v.MaxRegexLength = &maxregexlength
	return rb
}

// MaxRescoreWindow set the MaxRescoreWindow property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxRescoreWindow(maxrescorewindow int) *IndexSettingsBuilder {
	rb.v.MaxRescoreWindow = &maxrescorewindow
	return rb
}

// MaxResultWindow set the MaxResultWindow property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxResultWindow(maxresultwindow int) *IndexSettingsBuilder {
	rb.v.MaxResultWindow = &maxresultwindow
	return rb
}

// MaxScriptFields set the MaxScriptFields property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxScriptFields(maxscriptfields int) *IndexSettingsBuilder {
	rb.v.MaxScriptFields = &maxscriptfields
	return rb
}

// MaxShingleDiff set the MaxShingleDiff property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxShingleDiff(maxshinglediff int) *IndexSettingsBuilder {
	rb.v.MaxShingleDiff = &maxshinglediff
	return rb
}

// MaxSlicesPerScroll set the MaxSlicesPerScroll property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxSlicesPerScroll(maxslicesperscroll int) *IndexSettingsBuilder {
	rb.v.MaxSlicesPerScroll = &maxslicesperscroll
	return rb
}

// MaxTermsCount set the MaxTermsCount property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) MaxTermsCount(maxtermscount int) *IndexSettingsBuilder {
	rb.v.MaxTermsCount = &maxtermscount
	return rb
}

// Merge set the Merge property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Merge(merge Merge) *IndexSettingsBuilder {
	rb.v.Merge = &merge
	return rb
}

// Mode set the Mode property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Mode(mode string) *IndexSettingsBuilder {
	rb.v.Mode = &mode
	return rb
}

// NumberOfReplicas set the NumberOfReplicas property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) NumberOfReplicas(arg string) *IndexSettingsBuilder {
	rb.v.NumberOfReplicas = arg
	return rb
}

// NumberOfRoutingShards set the NumberOfRoutingShards property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) NumberOfRoutingShards(numberofroutingshards int) *IndexSettingsBuilder {
	rb.v.NumberOfRoutingShards = &numberofroutingshards
	return rb
}

// NumberOfShards set the NumberOfShards property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) NumberOfShards(arg string) *IndexSettingsBuilder {
	rb.v.NumberOfShards = arg
	return rb
}

// Priority set the Priority property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Priority(arg string) *IndexSettingsBuilder {
	rb.v.Priority = arg
	return rb
}

// ProvidedName set the ProvidedName property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) ProvidedName(providedname Name) *IndexSettingsBuilder {
	rb.v.ProvidedName = &providedname
	return rb
}

// Queries set the Queries property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Queries(queries Queries) *IndexSettingsBuilder {
	rb.v.Queries = &queries
	return rb
}

// QueryString set the QueryString property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) QueryString(querystring SettingsQueryString) *IndexSettingsBuilder {
	rb.v.QueryString = &querystring
	return rb
}

// RefreshInterval set the RefreshInterval property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) RefreshInterval(refreshinterval Time) *IndexSettingsBuilder {
	rb.v.RefreshInterval = &refreshinterval
	return rb
}

// Routing set the Routing property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Routing(routing IndexRouting) *IndexSettingsBuilder {
	rb.v.Routing = &routing
	return rb
}

// RoutingPartitionSize set the RoutingPartitionSize property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) RoutingPartitionSize(routingpartitionsize int) *IndexSettingsBuilder {
	rb.v.RoutingPartitionSize = &routingpartitionsize
	return rb
}

// RoutingPath set the RoutingPath property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) RoutingPath(arg []string) *IndexSettingsBuilder {
	rb.v.RoutingPath = arg
	return rb
}

// Search set the Search property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Search(search SettingsSearch) *IndexSettingsBuilder {
	rb.v.Search = &search
	return rb
}

// Settings set the Settings property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Settings(settings IndexSettings) *IndexSettingsBuilder {
	rb.v.Settings = &settings
	return rb
}

// Shards set the Shards property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Shards(shards int) *IndexSettingsBuilder {
	rb.v.Shards = &shards
	return rb
}

// Similarity Configure custom similarity settings to customize how search results are
// scored.
func (rb *IndexSettingsBuilder) Similarity(similarity SettingsSimilarity) *IndexSettingsBuilder {
	rb.v.Similarity = &similarity
	return rb
}

// SoftDeletes set the SoftDeletes property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) SoftDeletes(softdeletes SoftDeletes) *IndexSettingsBuilder {
	rb.v.SoftDeletes = &softdeletes
	return rb
}

// Sort set the Sort property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Sort(sort IndexSegmentSort) *IndexSettingsBuilder {
	rb.v.Sort = &sort
	return rb
}

// Store The store module allows you to control how index data is stored and accessed
// on disk.
func (rb *IndexSettingsBuilder) Store(store Storage) *IndexSettingsBuilder {
	rb.v.Store = &store
	return rb
}

// TimeSeries set the TimeSeries property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) TimeSeries(timeseries IndexSettingsTimeSeries) *IndexSettingsBuilder {
	rb.v.TimeSeries = &timeseries
	return rb
}

// TopMetricsMaxSize set the TopMetricsMaxSize property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) TopMetricsMaxSize(topmetricsmaxsize int) *IndexSettingsBuilder {
	rb.v.TopMetricsMaxSize = &topmetricsmaxsize
	return rb
}

// Translog set the Translog property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Translog(translog Translog) *IndexSettingsBuilder {
	rb.v.Translog = &translog
	return rb
}

// Uuid set the Uuid property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Uuid(uuid Uuid) *IndexSettingsBuilder {
	rb.v.Uuid = &uuid
	return rb
}

// VerifiedBeforeClose set the VerifiedBeforeClose property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) VerifiedBeforeClose(arg string) *IndexSettingsBuilder {
	rb.v.VerifiedBeforeClose = arg
	return rb
}

// Version set the Version property for IndexSettingsBuilder.
func (rb *IndexSettingsBuilder) Version(version IndexVersioning) *IndexSettingsBuilder {
	rb.v.Version = &version
	return rb
}
