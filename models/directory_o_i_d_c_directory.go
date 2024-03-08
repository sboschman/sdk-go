package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DirectoryOIDCDirectory struct {
    DirectoryAccountDirectory
    // The acrValues property
    acrValues *string
    // The attributesToStore property
    attributesToStore *string
    // The clientId property
    clientId *string
    // The clientSecret property
    clientSecret *string
    // The domainRestriction property
    domainRestriction *string
    // The enforces2fa property
    enforces2fa *bool
    // The fullyResolvedIssuer property
    fullyResolvedIssuer *string
    // The issuer property
    issuer *string
    // The logoutUrl property
    logoutUrl *string
    // The sendLoginHint property
    sendLoginHint *bool
    // The vendor property
    vendorEscaped *DirectoryOIDCVendor
}
// NewDirectoryOIDCDirectory instantiates a new DirectoryOIDCDirectory and sets the default values.
func NewDirectoryOIDCDirectory()(*DirectoryOIDCDirectory) {
    m := &DirectoryOIDCDirectory{
        DirectoryAccountDirectory: *NewDirectoryAccountDirectory(),
    }
    typeEscapedValue := "directory.OIDCDirectory"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateDirectoryOIDCDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryOIDCDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryOIDCDirectory(), nil
}
// GetAcrValues gets the acrValues property value. The acrValues property
// returns a *string when successful
func (m *DirectoryOIDCDirectory) GetAcrValues()(*string) {
    return m.acrValues
}
// GetAttributesToStore gets the attributesToStore property value. The attributesToStore property
// returns a *string when successful
func (m *DirectoryOIDCDirectory) GetAttributesToStore()(*string) {
    return m.attributesToStore
}
// GetClientId gets the clientId property value. The clientId property
// returns a *string when successful
func (m *DirectoryOIDCDirectory) GetClientId()(*string) {
    return m.clientId
}
// GetClientSecret gets the clientSecret property value. The clientSecret property
// returns a *string when successful
func (m *DirectoryOIDCDirectory) GetClientSecret()(*string) {
    return m.clientSecret
}
// GetDomainRestriction gets the domainRestriction property value. The domainRestriction property
// returns a *string when successful
func (m *DirectoryOIDCDirectory) GetDomainRestriction()(*string) {
    return m.domainRestriction
}
// GetEnforces2fa gets the enforces2fa property value. The enforces2fa property
// returns a *bool when successful
func (m *DirectoryOIDCDirectory) GetEnforces2fa()(*bool) {
    return m.enforces2fa
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectoryOIDCDirectory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryAccountDirectory.GetFieldDeserializers()
    res["acrValues"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAcrValues(val)
        }
        return nil
    }
    res["attributesToStore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributesToStore(val)
        }
        return nil
    }
    res["clientId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientId(val)
        }
        return nil
    }
    res["clientSecret"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClientSecret(val)
        }
        return nil
    }
    res["domainRestriction"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDomainRestriction(val)
        }
        return nil
    }
    res["enforces2fa"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnforces2fa(val)
        }
        return nil
    }
    res["fullyResolvedIssuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFullyResolvedIssuer(val)
        }
        return nil
    }
    res["issuer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIssuer(val)
        }
        return nil
    }
    res["logoutUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLogoutUrl(val)
        }
        return nil
    }
    res["sendLoginHint"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSendLoginHint(val)
        }
        return nil
    }
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryOIDCVendor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorEscaped(val.(*DirectoryOIDCVendor))
        }
        return nil
    }
    return res
}
// GetFullyResolvedIssuer gets the fullyResolvedIssuer property value. The fullyResolvedIssuer property
// returns a *string when successful
func (m *DirectoryOIDCDirectory) GetFullyResolvedIssuer()(*string) {
    return m.fullyResolvedIssuer
}
// GetIssuer gets the issuer property value. The issuer property
// returns a *string when successful
func (m *DirectoryOIDCDirectory) GetIssuer()(*string) {
    return m.issuer
}
// GetLogoutUrl gets the logoutUrl property value. The logoutUrl property
// returns a *string when successful
func (m *DirectoryOIDCDirectory) GetLogoutUrl()(*string) {
    return m.logoutUrl
}
// GetSendLoginHint gets the sendLoginHint property value. The sendLoginHint property
// returns a *bool when successful
func (m *DirectoryOIDCDirectory) GetSendLoginHint()(*bool) {
    return m.sendLoginHint
}
// GetVendorEscaped gets the vendor property value. The vendor property
// returns a *DirectoryOIDCVendor when successful
func (m *DirectoryOIDCDirectory) GetVendorEscaped()(*DirectoryOIDCVendor) {
    return m.vendorEscaped
}
// Serialize serializes information the current object
func (m *DirectoryOIDCDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryAccountDirectory.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("acrValues", m.GetAcrValues())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("attributesToStore", m.GetAttributesToStore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientId", m.GetClientId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("clientSecret", m.GetClientSecret())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("domainRestriction", m.GetDomainRestriction())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("enforces2fa", m.GetEnforces2fa())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("issuer", m.GetIssuer())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("logoutUrl", m.GetLogoutUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sendLoginHint", m.GetSendLoginHint())
        if err != nil {
            return err
        }
    }
    if m.GetVendorEscaped() != nil {
        cast := (*m.GetVendorEscaped()).String()
        err = writer.WriteStringValue("vendor", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAcrValues sets the acrValues property value. The acrValues property
func (m *DirectoryOIDCDirectory) SetAcrValues(value *string)() {
    m.acrValues = value
}
// SetAttributesToStore sets the attributesToStore property value. The attributesToStore property
func (m *DirectoryOIDCDirectory) SetAttributesToStore(value *string)() {
    m.attributesToStore = value
}
// SetClientId sets the clientId property value. The clientId property
func (m *DirectoryOIDCDirectory) SetClientId(value *string)() {
    m.clientId = value
}
// SetClientSecret sets the clientSecret property value. The clientSecret property
func (m *DirectoryOIDCDirectory) SetClientSecret(value *string)() {
    m.clientSecret = value
}
// SetDomainRestriction sets the domainRestriction property value. The domainRestriction property
func (m *DirectoryOIDCDirectory) SetDomainRestriction(value *string)() {
    m.domainRestriction = value
}
// SetEnforces2fa sets the enforces2fa property value. The enforces2fa property
func (m *DirectoryOIDCDirectory) SetEnforces2fa(value *bool)() {
    m.enforces2fa = value
}
// SetFullyResolvedIssuer sets the fullyResolvedIssuer property value. The fullyResolvedIssuer property
func (m *DirectoryOIDCDirectory) SetFullyResolvedIssuer(value *string)() {
    m.fullyResolvedIssuer = value
}
// SetIssuer sets the issuer property value. The issuer property
func (m *DirectoryOIDCDirectory) SetIssuer(value *string)() {
    m.issuer = value
}
// SetLogoutUrl sets the logoutUrl property value. The logoutUrl property
func (m *DirectoryOIDCDirectory) SetLogoutUrl(value *string)() {
    m.logoutUrl = value
}
// SetSendLoginHint sets the sendLoginHint property value. The sendLoginHint property
func (m *DirectoryOIDCDirectory) SetSendLoginHint(value *bool)() {
    m.sendLoginHint = value
}
// SetVendorEscaped sets the vendor property value. The vendor property
func (m *DirectoryOIDCDirectory) SetVendorEscaped(value *DirectoryOIDCVendor)() {
    m.vendorEscaped = value
}
type DirectoryOIDCDirectoryable interface {
    DirectoryAccountDirectoryable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAcrValues()(*string)
    GetAttributesToStore()(*string)
    GetClientId()(*string)
    GetClientSecret()(*string)
    GetDomainRestriction()(*string)
    GetEnforces2fa()(*bool)
    GetFullyResolvedIssuer()(*string)
    GetIssuer()(*string)
    GetLogoutUrl()(*string)
    GetSendLoginHint()(*bool)
    GetVendorEscaped()(*DirectoryOIDCVendor)
    SetAcrValues(value *string)()
    SetAttributesToStore(value *string)()
    SetClientId(value *string)()
    SetClientSecret(value *string)()
    SetDomainRestriction(value *string)()
    SetEnforces2fa(value *bool)()
    SetFullyResolvedIssuer(value *string)()
    SetIssuer(value *string)()
    SetLogoutUrl(value *string)()
    SetSendLoginHint(value *bool)()
    SetVendorEscaped(value *DirectoryOIDCVendor)()
}
