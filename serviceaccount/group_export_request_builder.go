package serviceaccount

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// GroupExportRequestBuilder builds and executes requests for operations under \serviceaccount\group\export
type GroupExportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GroupExportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupExportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupExportRequestBuilderInternal instantiates a new GroupExportRequestBuilder and sets the default values.
func NewGroupExportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupExportRequestBuilder) {
    m := &GroupExportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/serviceaccount/group/export", pathParameters),
    }
    return m
}
// NewGroupExportRequestBuilder instantiates a new GroupExportRequestBuilder and sets the default values.
func NewGroupExportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupExportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupExportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post prepares an export of service account group on systems using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *GroupExportRequestBuilder) Post(ctx context.Context, requestConfiguration *GroupExportRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation prepares an export of service account group on systems using the filtering specified in the query parameters. The URI of the export can be found in the returned Location header. This URI is valid for 2 minutes after being generated.
// returns a *RequestInformation when successful
func (m *GroupExportRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *GroupExportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=70")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GroupExportRequestBuilder when successful
func (m *GroupExportRequestBuilder) WithUrl(rawUrl string)(*GroupExportRequestBuilder) {
    return NewGroupExportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
