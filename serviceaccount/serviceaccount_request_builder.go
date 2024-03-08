package serviceaccount

import (
    "context"
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1 "github.com/topicuskeyhub/sdk-go/models"
)

// ServiceaccountRequestBuilder builds and executes requests for operations under \serviceaccount
type ServiceaccountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ServiceaccountRequestBuilderGetQueryParameters query for all service accounts in Topicus KeyHub. The various query parameters can be used to filter the response.
type ServiceaccountRequestBuilderGetQueryParameters struct {
    // Only return active or inactive service accounts.
    // Deprecated: This property is deprecated, use ActiveAsBooleanEnum instead
    Active []string `uriparametername:"active"`
    // Only return active or inactive service accounts.
    ActiveAsBooleanEnum []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.BooleanEnum `uriparametername:"active"`
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsGetAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsGetAdditionalQueryParameterType []GetAdditionalQueryParameterType `uriparametername:"additional"`
    // Return all or no records. This can be useful when composing parameters.
    Any []bool `uriparametername:"any"`
    // Only return records that have been created after the given instant.
    CreatedAfter []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdAfter"`
    // Only return records that have been created before the given instant.
    CreatedBefore []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"createdBefore"`
    // Filter the results to exclude the given ids.
    Exclude []int64 `uriparametername:"exclude"`
    // Filter the service accounts by groups on systems, specified by id. This parameter supports composition with all parameters from the group on system resource.
    GroupOnSystem []int64 `uriparametername:"groupOnSystem"`
    // Filter the service accounts on group on systems owned by any of the given groups, specified by id.
    GroupOnSystemOwners []int64 `uriparametername:"groupOnSystemOwners"`
    // Filter the results on the given ids.
    Id []int64 `uriparametername:"id"`
    // Only return records that have been modified since the given instant.
    ModifiedSince []i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time `uriparametername:"modifiedSince"`
    // Filter service accounts on the exact name.
    Name []string `uriparametername:"name"`
    // Search service accounts on (part of) the name, username or uuid.
    NameContains []string `uriparametername:"nameContains"`
    // Only return service accounts for which the name does not start with the given prefix.
    NameDoesNotStartWith []string `uriparametername:"nameDoesNotStartWith"`
    // Only return service accounts for which the name starts with the given prefix.
    NameStartsWith []string `uriparametername:"nameStartsWith"`
    // Filter service accounts on their organizational units, specified by id. This parameter supports composition with all parameters from the organizational unit resource.
    OrganizationalUnitForEnforcement []int64 `uriparametername:"organizationalUnitForEnforcement"`
    // Filter the service accounts by the password shared in the vault, specified by id. This parameter supports composition with all parameters from the vault record resource.
    Password []int64 `uriparametername:"password"`
    // Only return service accounts with the given password rotation scheme.
    // Deprecated: This property is deprecated, use PasswordRotationAsServiceaccountPasswordRotationScheme instead
    PasswordRotation []string `uriparametername:"passwordRotation"`
    // Only return service accounts with the given password rotation scheme.
    PasswordRotationAsServiceaccountPasswordRotationScheme []ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountPasswordRotationScheme `uriparametername:"passwordRotation"`
    // Filter records on a complex CQL query.
    Q []string `uriparametername:"q"`
    // Filter the service accounts on active requests for group on systems owned by any of the given groups, specified by id.
    RequestedGroupOnSystemOwners []int64 `uriparametername:"requestedGroupOnSystemOwners"`
    // Sort the items. Use 'asc-<name>' for ascending and 'desc-<name>' for descending order.
    Sort []string `uriparametername:"sort"`
    // Filter the service accounts by provisioned systems, specified by id. This parameter supports composition with all parameters from the provisioned system resource.
    System []int64 `uriparametername:"system"`
    // Filter the service accounts on the content administration group of a provisioned system, specified by id.
    SystemContentAdministrators []int64 `uriparametername:"systemContentAdministrators"`
    // Filter the service accounts on the owning group of a provisioned system, specified by id.
    SystemOwners []int64 `uriparametername:"systemOwners"`
    // Filter the service accounts on groups that perform technical administration for them, specified by id. This parameter supports composition with all parameters from the group resource.
    TechnicalAdministrator []int64 `uriparametername:"technicalAdministrator"`
    // Filter service accounts on the exact username.
    Username []string `uriparametername:"username"`
    // Filter results on one or more UUIDs.
    Uuid []string `uriparametername:"uuid"`
}
// ServiceaccountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceaccountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceaccountRequestBuilderGetQueryParameters
}
// ServiceaccountRequestBuilderPostQueryParameters creates one or more new service accounts and returns the newly created service accounts.
type ServiceaccountRequestBuilderPostQueryParameters struct {
    // Request additional information to be returned for every record.
    // Deprecated: This property is deprecated, use AdditionalAsPostAdditionalQueryParameterType instead
    Additional []string `uriparametername:"additional"`
    // Request additional information to be returned for every record.
    AdditionalAsPostAdditionalQueryParameterType []PostAdditionalQueryParameterType `uriparametername:"additional"`
}
// ServiceaccountRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ServiceaccountRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ServiceaccountRequestBuilderPostQueryParameters
}
// Auditstats the auditstats property
// returns a *AuditstatsRequestBuilder when successful
func (m *ServiceaccountRequestBuilder) Auditstats()(*AuditstatsRequestBuilder) {
    return NewAuditstatsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ByServiceaccountid gets an item from the github.com/topicuskeyhub/sdk-go.serviceaccount.item collection
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *WithServiceaccountItemRequestBuilder when successful
func (m *ServiceaccountRequestBuilder) ByServiceaccountid(serviceaccountid string)(*WithServiceaccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if serviceaccountid != "" {
        urlTplParams["serviceaccountid"] = serviceaccountid
    }
    return NewWithServiceaccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByServiceaccountidInt64 gets an item from the github.com/topicuskeyhub/sdk-go.serviceaccount.item collection
// returns a *WithServiceaccountItemRequestBuilder when successful
func (m *ServiceaccountRequestBuilder) ByServiceaccountidInt64(serviceaccountid int64)(*WithServiceaccountItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["serviceaccountid"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(serviceaccountid, 10)
    return NewWithServiceaccountItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewServiceaccountRequestBuilderInternal instantiates a new ServiceaccountRequestBuilder and sets the default values.
func NewServiceaccountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceaccountRequestBuilder) {
    m := &ServiceaccountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/serviceaccount{?active*,additional*,any*,createdAfter*,createdBefore*,exclude*,groupOnSystem*,groupOnSystemOwners*,id*,modifiedSince*,name*,nameContains*,nameDoesNotStartWith*,nameStartsWith*,organizationalUnitForEnforcement*,password*,passwordRotation*,q*,requestedGroupOnSystemOwners*,sort*,system*,systemContentAdministrators*,systemOwners*,technicalAdministrator*,username*,uuid*}", pathParameters),
    }
    return m
}
// NewServiceaccountRequestBuilder instantiates a new ServiceaccountRequestBuilder and sets the default values.
func NewServiceaccountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServiceaccountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServiceaccountRequestBuilderInternal(urlParams, requestAdapter)
}
// Export the export property
// returns a *ExportRequestBuilder when successful
func (m *ServiceaccountRequestBuilder) Export()(*ExportRequestBuilder) {
    return NewExportRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get query for all service accounts in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a ServiceaccountServiceAccountLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ServiceaccountRequestBuilder) Get(ctx context.Context, requestConfiguration *ServiceaccountRequestBuilderGetRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateServiceaccountServiceAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable), nil
}
// Group the group property
// returns a *GroupRequestBuilder when successful
func (m *ServiceaccountRequestBuilder) Group()(*GroupRequestBuilder) {
    return NewGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post creates one or more new service accounts and returns the newly created service accounts.
// returns a ServiceaccountServiceAccountLinkableWrapperable when successful
// returns a ErrorReport error when the service returns a 4XX or 5XX status code
func (m *ServiceaccountRequestBuilder) Post(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable, requestConfiguration *ServiceaccountRequestBuilderPostRequestConfiguration)(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateErrorReportFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.CreateServiceaccountServiceAccountLinkableWrapperFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable), nil
}
// ToGetRequestInformation query for all service accounts in Topicus KeyHub. The various query parameters can be used to filter the response.
// returns a *RequestInformation when successful
func (m *ServiceaccountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ServiceaccountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=69")
    return requestInfo, nil
}
// ToPostRequestInformation creates one or more new service accounts and returns the newly created service accounts.
// returns a *RequestInformation when successful
func (m *ServiceaccountRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie2969523f41a2fae7d38164656da4464a9222947e5ea7fbe5cbfbbf94304e5c1.ServiceaccountServiceAccountLinkableWrapperable, requestConfiguration *ServiceaccountRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/serviceaccount{?additional*}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/vnd.topicus.keyhub+json;version=69")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/vnd.topicus.keyhub+json;version=69", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ServiceaccountRequestBuilder when successful
func (m *ServiceaccountRequestBuilder) WithUrl(rawUrl string)(*ServiceaccountRequestBuilder) {
    return NewServiceaccountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
