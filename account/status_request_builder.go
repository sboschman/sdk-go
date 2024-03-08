package account

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// StatusRequestBuilder builds and executes requests for operations under \account\status
type StatusRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Bulk the bulk property
// returns a *StatusBulkRequestBuilder when successful
func (m *StatusRequestBuilder) Bulk()(*StatusBulkRequestBuilder) {
    return NewStatusBulkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewStatusRequestBuilderInternal instantiates a new StatusRequestBuilder and sets the default values.
func NewStatusRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatusRequestBuilder) {
    m := &StatusRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/account/status", pathParameters),
    }
    return m
}
// NewStatusRequestBuilder instantiates a new StatusRequestBuilder and sets the default values.
func NewStatusRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*StatusRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewStatusRequestBuilderInternal(urlParams, requestAdapter)
}
