package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type WebhookWebhook struct {
    Linkable
    // The account property
    account AuthAccountPrimerable
    // The active property
    active *bool
    // The additionalObjects property
    additionalObjects WebhookWebhook_additionalObjectsable
    // The allTypes property
    allTypes *bool
    // The authenticationScheme property
    authenticationScheme *HttpAuthenticationScheme
    // The basicAuthPassword property
    basicAuthPassword *string
    // The basicAuthUsername property
    basicAuthUsername *string
    // The bearerToken property
    bearerToken *string
    // The client property
    client ClientClientApplicationPrimerable
    // The clientCertificate property
    clientCertificate CertificateCertificatePrimerable
    // The customHeaderName property
    customHeaderName *string
    // The customHeaderValue property
    customHeaderValue *string
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
    // The group property
    group GroupGroupPrimerable
    // The name property
    name *string
    // The system property
    system ProvisioningProvisionedSystemPrimerable
    // The tls property
    tls *TLSLevel
    // The trustedCertificate property
    trustedCertificate CertificateCertificatePrimerable
    // The types property
    types []AuditAuditRecordType
    // The url property
    url *string
    // The uuid property
    uuid *string
    // The verbosePayloads property
    verbosePayloads *bool
}
// NewWebhookWebhook instantiates a new WebhookWebhook and sets the default values.
func NewWebhookWebhook()(*WebhookWebhook) {
    m := &WebhookWebhook{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "webhook.Webhook"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateWebhookWebhookFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateWebhookWebhookFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWebhookWebhook(), nil
}
// GetAccount gets the account property value. The account property
// returns a AuthAccountPrimerable when successful
func (m *WebhookWebhook) GetAccount()(AuthAccountPrimerable) {
    return m.account
}
// GetActive gets the active property value. The active property
// returns a *bool when successful
func (m *WebhookWebhook) GetActive()(*bool) {
    return m.active
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a WebhookWebhook_additionalObjectsable when successful
func (m *WebhookWebhook) GetAdditionalObjects()(WebhookWebhook_additionalObjectsable) {
    return m.additionalObjects
}
// GetAllTypes gets the allTypes property value. The allTypes property
// returns a *bool when successful
func (m *WebhookWebhook) GetAllTypes()(*bool) {
    return m.allTypes
}
// GetAuthenticationScheme gets the authenticationScheme property value. The authenticationScheme property
// returns a *HttpAuthenticationScheme when successful
func (m *WebhookWebhook) GetAuthenticationScheme()(*HttpAuthenticationScheme) {
    return m.authenticationScheme
}
// GetBasicAuthPassword gets the basicAuthPassword property value. The basicAuthPassword property
// returns a *string when successful
func (m *WebhookWebhook) GetBasicAuthPassword()(*string) {
    return m.basicAuthPassword
}
// GetBasicAuthUsername gets the basicAuthUsername property value. The basicAuthUsername property
// returns a *string when successful
func (m *WebhookWebhook) GetBasicAuthUsername()(*string) {
    return m.basicAuthUsername
}
// GetBearerToken gets the bearerToken property value. The bearerToken property
// returns a *string when successful
func (m *WebhookWebhook) GetBearerToken()(*string) {
    return m.bearerToken
}
// GetClient gets the client property value. The client property
// returns a ClientClientApplicationPrimerable when successful
func (m *WebhookWebhook) GetClient()(ClientClientApplicationPrimerable) {
    return m.client
}
// GetClientCertificate gets the clientCertificate property value. The clientCertificate property
// returns a CertificateCertificatePrimerable when successful
func (m *WebhookWebhook) GetClientCertificate()(CertificateCertificatePrimerable) {
    return m.clientCertificate
}
// GetCustomHeaderName gets the customHeaderName property value. The customHeaderName property
// returns a *string when successful
func (m *WebhookWebhook) GetCustomHeaderName()(*string) {
    return m.customHeaderName
}
// GetCustomHeaderValue gets the customHeaderValue property value. The customHeaderValue property
// returns a *string when successful
func (m *WebhookWebhook) GetCustomHeaderValue()(*string) {
    return m.customHeaderValue
}
// GetDirectory gets the directory property value. The directory property
// returns a DirectoryAccountDirectoryPrimerable when successful
func (m *WebhookWebhook) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *WebhookWebhook) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["account"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccount(val.(AuthAccountPrimerable))
        }
        return nil
    }
    res["active"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActive(val)
        }
        return nil
    }
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateWebhookWebhook_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(WebhookWebhook_additionalObjectsable))
        }
        return nil
    }
    res["allTypes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllTypes(val)
        }
        return nil
    }
    res["authenticationScheme"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseHttpAuthenticationScheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuthenticationScheme(val.(*HttpAuthenticationScheme))
        }
        return nil
    }
    res["basicAuthPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasicAuthPassword(val)
        }
        return nil
    }
    res["basicAuthUsername"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasicAuthUsername(val)
        }
        return nil
    }
    res["bearerToken"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBearerToken(val)
        }
        return nil
    }
    res["client"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientClientApplicationPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClient(val.(ClientClientApplicationPrimerable))
        }
        return nil
    }
    res["clientCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateCertificatePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientCertificate(val.(CertificateCertificatePrimerable))
        }
        return nil
    }
    res["customHeaderName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomHeaderName(val)
        }
        return nil
    }
    res["customHeaderValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCustomHeaderValue(val)
        }
        return nil
    }
    res["directory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryAccountDirectoryPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDirectory(val.(DirectoryAccountDirectoryPrimerable))
        }
        return nil
    }
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(ProvisioningProvisionedSystemPrimerable))
        }
        return nil
    }
    res["tls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTLSLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTls(val.(*TLSLevel))
        }
        return nil
    }
    res["trustedCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateCertificatePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTrustedCertificate(val.(CertificateCertificatePrimerable))
        }
        return nil
    }
    res["types"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseAuditAuditRecordType)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuditAuditRecordType, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*AuditAuditRecordType))
                }
            }
            m.SetTypes(res)
        }
        return nil
    }
    res["url"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUrl(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    res["verbosePayloads"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerbosePayloads(val)
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
// returns a GroupGroupPrimerable when successful
func (m *WebhookWebhook) GetGroup()(GroupGroupPrimerable) {
    return m.group
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *WebhookWebhook) GetName()(*string) {
    return m.name
}
// GetSystem gets the system property value. The system property
// returns a ProvisioningProvisionedSystemPrimerable when successful
func (m *WebhookWebhook) GetSystem()(ProvisioningProvisionedSystemPrimerable) {
    return m.system
}
// GetTls gets the tls property value. The tls property
// returns a *TLSLevel when successful
func (m *WebhookWebhook) GetTls()(*TLSLevel) {
    return m.tls
}
// GetTrustedCertificate gets the trustedCertificate property value. The trustedCertificate property
// returns a CertificateCertificatePrimerable when successful
func (m *WebhookWebhook) GetTrustedCertificate()(CertificateCertificatePrimerable) {
    return m.trustedCertificate
}
// GetTypes gets the types property value. The types property
// returns a []AuditAuditRecordType when successful
func (m *WebhookWebhook) GetTypes()([]AuditAuditRecordType) {
    return m.types
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *WebhookWebhook) GetUrl()(*string) {
    return m.url
}
// GetUuid gets the uuid property value. The uuid property
// returns a *string when successful
func (m *WebhookWebhook) GetUuid()(*string) {
    return m.uuid
}
// GetVerbosePayloads gets the verbosePayloads property value. The verbosePayloads property
// returns a *bool when successful
func (m *WebhookWebhook) GetVerbosePayloads()(*bool) {
    return m.verbosePayloads
}
// Serialize serializes information the current object
func (m *WebhookWebhook) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("account", m.GetAccount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("active", m.GetActive())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allTypes", m.GetAllTypes())
        if err != nil {
            return err
        }
    }
    if m.GetAuthenticationScheme() != nil {
        cast := (*m.GetAuthenticationScheme()).String()
        err = writer.WriteStringValue("authenticationScheme", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("basicAuthPassword", m.GetBasicAuthPassword())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("basicAuthUsername", m.GetBasicAuthUsername())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bearerToken", m.GetBearerToken())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("client", m.GetClient())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("clientCertificate", m.GetClientCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customHeaderName", m.GetCustomHeaderName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("customHeaderValue", m.GetCustomHeaderValue())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    if m.GetTls() != nil {
        cast := (*m.GetTls()).String()
        err = writer.WriteStringValue("tls", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("trustedCertificate", m.GetTrustedCertificate())
        if err != nil {
            return err
        }
    }
    if m.GetTypes() != nil {
        err = writer.WriteCollectionOfStringValues("types", SerializeAuditAuditRecordType(m.GetTypes()))
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("url", m.GetUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("verbosePayloads", m.GetVerbosePayloads())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccount sets the account property value. The account property
func (m *WebhookWebhook) SetAccount(value AuthAccountPrimerable)() {
    m.account = value
}
// SetActive sets the active property value. The active property
func (m *WebhookWebhook) SetActive(value *bool)() {
    m.active = value
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *WebhookWebhook) SetAdditionalObjects(value WebhookWebhook_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetAllTypes sets the allTypes property value. The allTypes property
func (m *WebhookWebhook) SetAllTypes(value *bool)() {
    m.allTypes = value
}
// SetAuthenticationScheme sets the authenticationScheme property value. The authenticationScheme property
func (m *WebhookWebhook) SetAuthenticationScheme(value *HttpAuthenticationScheme)() {
    m.authenticationScheme = value
}
// SetBasicAuthPassword sets the basicAuthPassword property value. The basicAuthPassword property
func (m *WebhookWebhook) SetBasicAuthPassword(value *string)() {
    m.basicAuthPassword = value
}
// SetBasicAuthUsername sets the basicAuthUsername property value. The basicAuthUsername property
func (m *WebhookWebhook) SetBasicAuthUsername(value *string)() {
    m.basicAuthUsername = value
}
// SetBearerToken sets the bearerToken property value. The bearerToken property
func (m *WebhookWebhook) SetBearerToken(value *string)() {
    m.bearerToken = value
}
// SetClient sets the client property value. The client property
func (m *WebhookWebhook) SetClient(value ClientClientApplicationPrimerable)() {
    m.client = value
}
// SetClientCertificate sets the clientCertificate property value. The clientCertificate property
func (m *WebhookWebhook) SetClientCertificate(value CertificateCertificatePrimerable)() {
    m.clientCertificate = value
}
// SetCustomHeaderName sets the customHeaderName property value. The customHeaderName property
func (m *WebhookWebhook) SetCustomHeaderName(value *string)() {
    m.customHeaderName = value
}
// SetCustomHeaderValue sets the customHeaderValue property value. The customHeaderValue property
func (m *WebhookWebhook) SetCustomHeaderValue(value *string)() {
    m.customHeaderValue = value
}
// SetDirectory sets the directory property value. The directory property
func (m *WebhookWebhook) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
// SetGroup sets the group property value. The group property
func (m *WebhookWebhook) SetGroup(value GroupGroupPrimerable)() {
    m.group = value
}
// SetName sets the name property value. The name property
func (m *WebhookWebhook) SetName(value *string)() {
    m.name = value
}
// SetSystem sets the system property value. The system property
func (m *WebhookWebhook) SetSystem(value ProvisioningProvisionedSystemPrimerable)() {
    m.system = value
}
// SetTls sets the tls property value. The tls property
func (m *WebhookWebhook) SetTls(value *TLSLevel)() {
    m.tls = value
}
// SetTrustedCertificate sets the trustedCertificate property value. The trustedCertificate property
func (m *WebhookWebhook) SetTrustedCertificate(value CertificateCertificatePrimerable)() {
    m.trustedCertificate = value
}
// SetTypes sets the types property value. The types property
func (m *WebhookWebhook) SetTypes(value []AuditAuditRecordType)() {
    m.types = value
}
// SetUrl sets the url property value. The url property
func (m *WebhookWebhook) SetUrl(value *string)() {
    m.url = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *WebhookWebhook) SetUuid(value *string)() {
    m.uuid = value
}
// SetVerbosePayloads sets the verbosePayloads property value. The verbosePayloads property
func (m *WebhookWebhook) SetVerbosePayloads(value *bool)() {
    m.verbosePayloads = value
}
type WebhookWebhookable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccount()(AuthAccountPrimerable)
    GetActive()(*bool)
    GetAdditionalObjects()(WebhookWebhook_additionalObjectsable)
    GetAllTypes()(*bool)
    GetAuthenticationScheme()(*HttpAuthenticationScheme)
    GetBasicAuthPassword()(*string)
    GetBasicAuthUsername()(*string)
    GetBearerToken()(*string)
    GetClient()(ClientClientApplicationPrimerable)
    GetClientCertificate()(CertificateCertificatePrimerable)
    GetCustomHeaderName()(*string)
    GetCustomHeaderValue()(*string)
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    GetGroup()(GroupGroupPrimerable)
    GetName()(*string)
    GetSystem()(ProvisioningProvisionedSystemPrimerable)
    GetTls()(*TLSLevel)
    GetTrustedCertificate()(CertificateCertificatePrimerable)
    GetTypes()([]AuditAuditRecordType)
    GetUrl()(*string)
    GetUuid()(*string)
    GetVerbosePayloads()(*bool)
    SetAccount(value AuthAccountPrimerable)()
    SetActive(value *bool)()
    SetAdditionalObjects(value WebhookWebhook_additionalObjectsable)()
    SetAllTypes(value *bool)()
    SetAuthenticationScheme(value *HttpAuthenticationScheme)()
    SetBasicAuthPassword(value *string)()
    SetBasicAuthUsername(value *string)()
    SetBearerToken(value *string)()
    SetClient(value ClientClientApplicationPrimerable)()
    SetClientCertificate(value CertificateCertificatePrimerable)()
    SetCustomHeaderName(value *string)()
    SetCustomHeaderValue(value *string)()
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
    SetGroup(value GroupGroupPrimerable)()
    SetName(value *string)()
    SetSystem(value ProvisioningProvisionedSystemPrimerable)()
    SetTls(value *TLSLevel)()
    SetTrustedCertificate(value CertificateCertificatePrimerable)()
    SetTypes(value []AuditAuditRecordType)()
    SetUrl(value *string)()
    SetUuid(value *string)()
    SetVerbosePayloads(value *bool)()
}
