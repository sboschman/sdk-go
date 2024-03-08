package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningAbstractProvisionedLDAP struct {
    ProvisioningProvisionedSystem
    // The attributes property
    attributes ProvisioningAbstractProvisionedLDAP_attributesable
    // The baseDN property
    baseDN *string
    // The bindDN property
    bindDN *string
    // The bindPassword property
    bindPassword *string
    // The clientCertificate property
    clientCertificate CertificateCertificatePrimerable
    // The failoverHost property
    failoverHost *string
    // The failoverTrustedCertificate property
    failoverTrustedCertificate CertificateCertificatePrimerable
    // The groupDN property
    groupDN *string
    // The host property
    host *string
    // The objectClasses property
    objectClasses *string
    // The port property
    port *int32
    // The serviceAccountDN property
    serviceAccountDN *string
    // The sshPublicKeySupported property
    sshPublicKeySupported *bool
    // The tls property
    tls *TLSLevel
    // The trustedCertificate property
    trustedCertificate CertificateCertificatePrimerable
    // The userDN property
    userDN *string
}
// NewProvisioningAbstractProvisionedLDAP instantiates a new ProvisioningAbstractProvisionedLDAP and sets the default values.
func NewProvisioningAbstractProvisionedLDAP()(*ProvisioningAbstractProvisionedLDAP) {
    m := &ProvisioningAbstractProvisionedLDAP{
        ProvisioningProvisionedSystem: *NewProvisioningProvisionedSystem(),
    }
    typeEscapedValue := "provisioning.AbstractProvisionedLDAP"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningAbstractProvisionedLDAPFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningAbstractProvisionedLDAPFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "provisioning.ProvisionedAD":
                        return NewProvisioningProvisionedAD(), nil
                    case "provisioning.ProvisionedLDAP":
                        return NewProvisioningProvisionedLDAP(), nil
                }
            }
        }
    }
    return NewProvisioningAbstractProvisionedLDAP(), nil
}
// GetAttributes gets the attributes property value. The attributes property
// returns a ProvisioningAbstractProvisionedLDAP_attributesable when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetAttributes()(ProvisioningAbstractProvisionedLDAP_attributesable) {
    return m.attributes
}
// GetBaseDN gets the baseDN property value. The baseDN property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetBaseDN()(*string) {
    return m.baseDN
}
// GetBindDN gets the bindDN property value. The bindDN property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetBindDN()(*string) {
    return m.bindDN
}
// GetBindPassword gets the bindPassword property value. The bindPassword property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetBindPassword()(*string) {
    return m.bindPassword
}
// GetClientCertificate gets the clientCertificate property value. The clientCertificate property
// returns a CertificateCertificatePrimerable when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetClientCertificate()(CertificateCertificatePrimerable) {
    return m.clientCertificate
}
// GetFailoverHost gets the failoverHost property value. The failoverHost property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetFailoverHost()(*string) {
    return m.failoverHost
}
// GetFailoverTrustedCertificate gets the failoverTrustedCertificate property value. The failoverTrustedCertificate property
// returns a CertificateCertificatePrimerable when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetFailoverTrustedCertificate()(CertificateCertificatePrimerable) {
    return m.failoverTrustedCertificate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystem.GetFieldDeserializers()
    res["attributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningAbstractProvisionedLDAP_attributesFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttributes(val.(ProvisioningAbstractProvisionedLDAP_attributesable))
        }
        return nil
    }
    res["baseDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseDN(val)
        }
        return nil
    }
    res["bindDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindDN(val)
        }
        return nil
    }
    res["bindPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBindPassword(val)
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
    res["failoverHost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailoverHost(val)
        }
        return nil
    }
    res["failoverTrustedCertificate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateCertificateCertificatePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFailoverTrustedCertificate(val.(CertificateCertificatePrimerable))
        }
        return nil
    }
    res["groupDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupDN(val)
        }
        return nil
    }
    res["host"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHost(val)
        }
        return nil
    }
    res["objectClasses"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetObjectClasses(val)
        }
        return nil
    }
    res["port"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["serviceAccountDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAccountDN(val)
        }
        return nil
    }
    res["sshPublicKeySupported"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSshPublicKeySupported(val)
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
    res["userDN"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserDN(val)
        }
        return nil
    }
    return res
}
// GetGroupDN gets the groupDN property value. The groupDN property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetGroupDN()(*string) {
    return m.groupDN
}
// GetHost gets the host property value. The host property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetHost()(*string) {
    return m.host
}
// GetObjectClasses gets the objectClasses property value. The objectClasses property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetObjectClasses()(*string) {
    return m.objectClasses
}
// GetPort gets the port property value. The port property
// returns a *int32 when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetPort()(*int32) {
    return m.port
}
// GetServiceAccountDN gets the serviceAccountDN property value. The serviceAccountDN property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetServiceAccountDN()(*string) {
    return m.serviceAccountDN
}
// GetSshPublicKeySupported gets the sshPublicKeySupported property value. The sshPublicKeySupported property
// returns a *bool when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetSshPublicKeySupported()(*bool) {
    return m.sshPublicKeySupported
}
// GetTls gets the tls property value. The tls property
// returns a *TLSLevel when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetTls()(*TLSLevel) {
    return m.tls
}
// GetTrustedCertificate gets the trustedCertificate property value. The trustedCertificate property
// returns a CertificateCertificatePrimerable when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetTrustedCertificate()(CertificateCertificatePrimerable) {
    return m.trustedCertificate
}
// GetUserDN gets the userDN property value. The userDN property
// returns a *string when successful
func (m *ProvisioningAbstractProvisionedLDAP) GetUserDN()(*string) {
    return m.userDN
}
// Serialize serializes information the current object
func (m *ProvisioningAbstractProvisionedLDAP) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("attributes", m.GetAttributes())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("baseDN", m.GetBaseDN())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bindDN", m.GetBindDN())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bindPassword", m.GetBindPassword())
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
        err = writer.WriteStringValue("failoverHost", m.GetFailoverHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("failoverTrustedCertificate", m.GetFailoverTrustedCertificate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupDN", m.GetGroupDN())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("host", m.GetHost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("objectClasses", m.GetObjectClasses())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serviceAccountDN", m.GetServiceAccountDN())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("sshPublicKeySupported", m.GetSshPublicKeySupported())
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
    {
        err = writer.WriteStringValue("userDN", m.GetUserDN())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttributes sets the attributes property value. The attributes property
func (m *ProvisioningAbstractProvisionedLDAP) SetAttributes(value ProvisioningAbstractProvisionedLDAP_attributesable)() {
    m.attributes = value
}
// SetBaseDN sets the baseDN property value. The baseDN property
func (m *ProvisioningAbstractProvisionedLDAP) SetBaseDN(value *string)() {
    m.baseDN = value
}
// SetBindDN sets the bindDN property value. The bindDN property
func (m *ProvisioningAbstractProvisionedLDAP) SetBindDN(value *string)() {
    m.bindDN = value
}
// SetBindPassword sets the bindPassword property value. The bindPassword property
func (m *ProvisioningAbstractProvisionedLDAP) SetBindPassword(value *string)() {
    m.bindPassword = value
}
// SetClientCertificate sets the clientCertificate property value. The clientCertificate property
func (m *ProvisioningAbstractProvisionedLDAP) SetClientCertificate(value CertificateCertificatePrimerable)() {
    m.clientCertificate = value
}
// SetFailoverHost sets the failoverHost property value. The failoverHost property
func (m *ProvisioningAbstractProvisionedLDAP) SetFailoverHost(value *string)() {
    m.failoverHost = value
}
// SetFailoverTrustedCertificate sets the failoverTrustedCertificate property value. The failoverTrustedCertificate property
func (m *ProvisioningAbstractProvisionedLDAP) SetFailoverTrustedCertificate(value CertificateCertificatePrimerable)() {
    m.failoverTrustedCertificate = value
}
// SetGroupDN sets the groupDN property value. The groupDN property
func (m *ProvisioningAbstractProvisionedLDAP) SetGroupDN(value *string)() {
    m.groupDN = value
}
// SetHost sets the host property value. The host property
func (m *ProvisioningAbstractProvisionedLDAP) SetHost(value *string)() {
    m.host = value
}
// SetObjectClasses sets the objectClasses property value. The objectClasses property
func (m *ProvisioningAbstractProvisionedLDAP) SetObjectClasses(value *string)() {
    m.objectClasses = value
}
// SetPort sets the port property value. The port property
func (m *ProvisioningAbstractProvisionedLDAP) SetPort(value *int32)() {
    m.port = value
}
// SetServiceAccountDN sets the serviceAccountDN property value. The serviceAccountDN property
func (m *ProvisioningAbstractProvisionedLDAP) SetServiceAccountDN(value *string)() {
    m.serviceAccountDN = value
}
// SetSshPublicKeySupported sets the sshPublicKeySupported property value. The sshPublicKeySupported property
func (m *ProvisioningAbstractProvisionedLDAP) SetSshPublicKeySupported(value *bool)() {
    m.sshPublicKeySupported = value
}
// SetTls sets the tls property value. The tls property
func (m *ProvisioningAbstractProvisionedLDAP) SetTls(value *TLSLevel)() {
    m.tls = value
}
// SetTrustedCertificate sets the trustedCertificate property value. The trustedCertificate property
func (m *ProvisioningAbstractProvisionedLDAP) SetTrustedCertificate(value CertificateCertificatePrimerable)() {
    m.trustedCertificate = value
}
// SetUserDN sets the userDN property value. The userDN property
func (m *ProvisioningAbstractProvisionedLDAP) SetUserDN(value *string)() {
    m.userDN = value
}
type ProvisioningAbstractProvisionedLDAPable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetAttributes()(ProvisioningAbstractProvisionedLDAP_attributesable)
    GetBaseDN()(*string)
    GetBindDN()(*string)
    GetBindPassword()(*string)
    GetClientCertificate()(CertificateCertificatePrimerable)
    GetFailoverHost()(*string)
    GetFailoverTrustedCertificate()(CertificateCertificatePrimerable)
    GetGroupDN()(*string)
    GetHost()(*string)
    GetObjectClasses()(*string)
    GetPort()(*int32)
    GetServiceAccountDN()(*string)
    GetSshPublicKeySupported()(*bool)
    GetTls()(*TLSLevel)
    GetTrustedCertificate()(CertificateCertificatePrimerable)
    GetUserDN()(*string)
    SetAttributes(value ProvisioningAbstractProvisionedLDAP_attributesable)()
    SetBaseDN(value *string)()
    SetBindDN(value *string)()
    SetBindPassword(value *string)()
    SetClientCertificate(value CertificateCertificatePrimerable)()
    SetFailoverHost(value *string)()
    SetFailoverTrustedCertificate(value CertificateCertificatePrimerable)()
    SetGroupDN(value *string)()
    SetHost(value *string)()
    SetObjectClasses(value *string)()
    SetPort(value *int32)()
    SetServiceAccountDN(value *string)()
    SetSshPublicKeySupported(value *bool)()
    SetTls(value *TLSLevel)()
    SetTrustedCertificate(value CertificateCertificatePrimerable)()
    SetUserDN(value *string)()
}
