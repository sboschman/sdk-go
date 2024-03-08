package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedSCIM struct {
    ProvisioningProvisionedSystem
    // The authenticationScheme property
    authenticationScheme *HttpAuthenticationScheme
    // The basicAuthPassword property
    basicAuthPassword *string
    // The basicAuthUsername property
    basicAuthUsername *string
    // The bearerToken property
    bearerToken *string
    // The customHeaderName property
    customHeaderName *string
    // The customHeaderValue property
    customHeaderValue *string
    // The url property
    url *string
    // The vendor property
    vendorEscaped *ProvisioningProvisionedSCIMVendor
}
// NewProvisioningProvisionedSCIM instantiates a new ProvisioningProvisionedSCIM and sets the default values.
func NewProvisioningProvisionedSCIM()(*ProvisioningProvisionedSCIM) {
    m := &ProvisioningProvisionedSCIM{
        ProvisioningProvisionedSystem: *NewProvisioningProvisionedSystem(),
    }
    typeEscapedValue := "provisioning.ProvisionedSCIM"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedSCIMFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedSCIMFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedSCIM(), nil
}
// GetAuthenticationScheme gets the authenticationScheme property value. The authenticationScheme property
// returns a *HttpAuthenticationScheme when successful
func (m *ProvisioningProvisionedSCIM) GetAuthenticationScheme()(*HttpAuthenticationScheme) {
    return m.authenticationScheme
}
// GetBasicAuthPassword gets the basicAuthPassword property value. The basicAuthPassword property
// returns a *string when successful
func (m *ProvisioningProvisionedSCIM) GetBasicAuthPassword()(*string) {
    return m.basicAuthPassword
}
// GetBasicAuthUsername gets the basicAuthUsername property value. The basicAuthUsername property
// returns a *string when successful
func (m *ProvisioningProvisionedSCIM) GetBasicAuthUsername()(*string) {
    return m.basicAuthUsername
}
// GetBearerToken gets the bearerToken property value. The bearerToken property
// returns a *string when successful
func (m *ProvisioningProvisionedSCIM) GetBearerToken()(*string) {
    return m.bearerToken
}
// GetCustomHeaderName gets the customHeaderName property value. The customHeaderName property
// returns a *string when successful
func (m *ProvisioningProvisionedSCIM) GetCustomHeaderName()(*string) {
    return m.customHeaderName
}
// GetCustomHeaderValue gets the customHeaderValue property value. The customHeaderValue property
// returns a *string when successful
func (m *ProvisioningProvisionedSCIM) GetCustomHeaderValue()(*string) {
    return m.customHeaderValue
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedSCIM) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystem.GetFieldDeserializers()
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
    res["vendor"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningProvisionedSCIMVendor)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorEscaped(val.(*ProvisioningProvisionedSCIMVendor))
        }
        return nil
    }
    return res
}
// GetUrl gets the url property value. The url property
// returns a *string when successful
func (m *ProvisioningProvisionedSCIM) GetUrl()(*string) {
    return m.url
}
// GetVendorEscaped gets the vendor property value. The vendor property
// returns a *ProvisioningProvisionedSCIMVendor when successful
func (m *ProvisioningProvisionedSCIM) GetVendorEscaped()(*ProvisioningProvisionedSCIMVendor) {
    return m.vendorEscaped
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedSCIM) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystem.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteStringValue("url", m.GetUrl())
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
// SetAuthenticationScheme sets the authenticationScheme property value. The authenticationScheme property
func (m *ProvisioningProvisionedSCIM) SetAuthenticationScheme(value *HttpAuthenticationScheme)() {
    m.authenticationScheme = value
}
// SetBasicAuthPassword sets the basicAuthPassword property value. The basicAuthPassword property
func (m *ProvisioningProvisionedSCIM) SetBasicAuthPassword(value *string)() {
    m.basicAuthPassword = value
}
// SetBasicAuthUsername sets the basicAuthUsername property value. The basicAuthUsername property
func (m *ProvisioningProvisionedSCIM) SetBasicAuthUsername(value *string)() {
    m.basicAuthUsername = value
}
// SetBearerToken sets the bearerToken property value. The bearerToken property
func (m *ProvisioningProvisionedSCIM) SetBearerToken(value *string)() {
    m.bearerToken = value
}
// SetCustomHeaderName sets the customHeaderName property value. The customHeaderName property
func (m *ProvisioningProvisionedSCIM) SetCustomHeaderName(value *string)() {
    m.customHeaderName = value
}
// SetCustomHeaderValue sets the customHeaderValue property value. The customHeaderValue property
func (m *ProvisioningProvisionedSCIM) SetCustomHeaderValue(value *string)() {
    m.customHeaderValue = value
}
// SetUrl sets the url property value. The url property
func (m *ProvisioningProvisionedSCIM) SetUrl(value *string)() {
    m.url = value
}
// SetVendorEscaped sets the vendor property value. The vendor property
func (m *ProvisioningProvisionedSCIM) SetVendorEscaped(value *ProvisioningProvisionedSCIMVendor)() {
    m.vendorEscaped = value
}
type ProvisioningProvisionedSCIMable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetAuthenticationScheme()(*HttpAuthenticationScheme)
    GetBasicAuthPassword()(*string)
    GetBasicAuthUsername()(*string)
    GetBearerToken()(*string)
    GetCustomHeaderName()(*string)
    GetCustomHeaderValue()(*string)
    GetUrl()(*string)
    GetVendorEscaped()(*ProvisioningProvisionedSCIMVendor)
    SetAuthenticationScheme(value *HttpAuthenticationScheme)()
    SetBasicAuthPassword(value *string)()
    SetBasicAuthUsername(value *string)()
    SetBearerToken(value *string)()
    SetCustomHeaderName(value *string)()
    SetCustomHeaderValue(value *string)()
    SetUrl(value *string)()
    SetVendorEscaped(value *ProvisioningProvisionedSCIMVendor)()
}
