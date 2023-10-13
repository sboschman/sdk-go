package system

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ItemSynclogRequestBuilder builds and executes requests for operations under \system\{systemid}\synclog
type ItemSynclogRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynclogRequestBuilderGetQueryParameters query for sync logs for a provisioned system. The various query parameters can be used to filter the response.
type ItemSynclogRequestBuilderGetQueryParameters struct {
    // Request additional information to be returned for every record.
    Additional []string `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Only return logs for the given systems, specified by id. This parameter supports composition with all parameters from the systems resource.
    System []int64 `uriparametername:"system"`
}
// ItemSynclogRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynclogRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSynclogRequestBuilderGetQueryParameters
}
// BySynclogid gets an item from the github.com/topicuskeyhub/sdk-go.system.item.synclog.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
func (m *ItemSynclogRequestBuilder) BySynclogid(synclogid string)(*ItemSynclogWithSynclogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if synclogid != "" {
        urlTplParams["synclogid"] = synclogid
    }
    return NewItemSynclogWithSynclogItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// BySynclogidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.system.item.synclog.item collection
func (m *ItemSynclogRequestBuilder) BySynclogidInt64(synclogid int64)(*ItemSynclogWithSynclogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["synclogid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(synclogid, 10)
    return NewItemSynclogWithSynclogItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSynclogRequestBuilderInternal instantiates a new SynclogRequestBuilder and sets the default values.
func NewItemSynclogRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynclogRequestBuilder) {
    m := &ItemSynclogRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/system/{systemid}/synclog{?additional*,any*,createdAfter*,createdBefore*,exclude*,id*,modifiedSince*,q*,system*}", pathParameters),
    }
    return m
}
// NewItemSynclogRequestBuilder instantiates a new SynclogRequestBuilder and sets the default values.
func NewItemSynclogRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynclogRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynclogRequestBuilderInternal(urlParams, requestAdapter)
}
// Get query for sync logs for a provisioned system. The various query parameters can be used to filter the response.
func (m *ItemSynclogRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSynclogRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemSyncLogLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
        "5XX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateProvisioningProvisionedSystemSyncLogLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ProvisioningProvisionedSystemSyncLogLinkableWrapperable), nil
}
// ToGetRequestInformation query for sync logs for a provisioned system. The various query parameters can be used to filter the response.
func (m *ItemSynclogRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSynclogRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/vnd.topicus.keyhub+json;version=67")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSynclogRequestBuilder) WithUrl(rawUrl string)(*ItemSynclogRequestBuilder) {
    return NewItemSynclogRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
