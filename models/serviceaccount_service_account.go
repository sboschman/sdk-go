package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ServiceaccountServiceAccount struct {
    ServiceaccountServiceAccountPrimer
    // The additionalObjects property
    additionalObjects ServiceaccountServiceAccount_additionalObjectsable
    // The description property
    description *string
    // The password property
    password VaultVaultRecordPrimerable
    // The passwordRotation property
    passwordRotation *ServiceaccountPasswordRotationScheme
    // The technicalAdministrator property
    technicalAdministrator GroupGroupPrimerable
}
// NewServiceaccountServiceAccount instantiates a new ServiceaccountServiceAccount and sets the default values.
func NewServiceaccountServiceAccount()(*ServiceaccountServiceAccount) {
    m := &ServiceaccountServiceAccount{
        ServiceaccountServiceAccountPrimer: *NewServiceaccountServiceAccountPrimer(),
    }
    typeEscapedValue := "serviceaccount.ServiceAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateServiceaccountServiceAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateServiceaccountServiceAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewServiceaccountServiceAccount(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ServiceaccountServiceAccount_additionalObjectsable when successful
func (m *ServiceaccountServiceAccount) GetAdditionalObjects()(ServiceaccountServiceAccount_additionalObjectsable) {
    return m.additionalObjects
}
// GetDescription gets the description property value. The description property
// returns a *string when successful
func (m *ServiceaccountServiceAccount) GetDescription()(*string) {
    return m.description
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ServiceaccountServiceAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ServiceaccountServiceAccountPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateServiceaccountServiceAccount_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ServiceaccountServiceAccount_additionalObjectsable))
        }
        return nil
    }
    res["description"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVaultVaultRecordPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val.(VaultVaultRecordPrimerable))
        }
        return nil
    }
    res["passwordRotation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseServiceaccountPasswordRotationScheme)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRotation(val.(*ServiceaccountPasswordRotationScheme))
        }
        return nil
    }
    res["technicalAdministrator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnicalAdministrator(val.(GroupGroupPrimerable))
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
// returns a VaultVaultRecordPrimerable when successful
func (m *ServiceaccountServiceAccount) GetPassword()(VaultVaultRecordPrimerable) {
    return m.password
}
// GetPasswordRotation gets the passwordRotation property value. The passwordRotation property
// returns a *ServiceaccountPasswordRotationScheme when successful
func (m *ServiceaccountServiceAccount) GetPasswordRotation()(*ServiceaccountPasswordRotationScheme) {
    return m.passwordRotation
}
// GetTechnicalAdministrator gets the technicalAdministrator property value. The technicalAdministrator property
// returns a GroupGroupPrimerable when successful
func (m *ServiceaccountServiceAccount) GetTechnicalAdministrator()(GroupGroupPrimerable) {
    return m.technicalAdministrator
}
// Serialize serializes information the current object
func (m *ServiceaccountServiceAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ServiceaccountServiceAccountPrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    if m.GetPasswordRotation() != nil {
        cast := (*m.GetPasswordRotation()).String()
        err = writer.WriteStringValue("passwordRotation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("technicalAdministrator", m.GetTechnicalAdministrator())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ServiceaccountServiceAccount) SetAdditionalObjects(value ServiceaccountServiceAccount_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetDescription sets the description property value. The description property
func (m *ServiceaccountServiceAccount) SetDescription(value *string)() {
    m.description = value
}
// SetPassword sets the password property value. The password property
func (m *ServiceaccountServiceAccount) SetPassword(value VaultVaultRecordPrimerable)() {
    m.password = value
}
// SetPasswordRotation sets the passwordRotation property value. The passwordRotation property
func (m *ServiceaccountServiceAccount) SetPasswordRotation(value *ServiceaccountPasswordRotationScheme)() {
    m.passwordRotation = value
}
// SetTechnicalAdministrator sets the technicalAdministrator property value. The technicalAdministrator property
func (m *ServiceaccountServiceAccount) SetTechnicalAdministrator(value GroupGroupPrimerable)() {
    m.technicalAdministrator = value
}
type ServiceaccountServiceAccountable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ServiceaccountServiceAccountPrimerable
    GetAdditionalObjects()(ServiceaccountServiceAccount_additionalObjectsable)
    GetDescription()(*string)
    GetPassword()(VaultVaultRecordPrimerable)
    GetPasswordRotation()(*ServiceaccountPasswordRotationScheme)
    GetTechnicalAdministrator()(GroupGroupPrimerable)
    SetAdditionalObjects(value ServiceaccountServiceAccount_additionalObjectsable)()
    SetDescription(value *string)()
    SetPassword(value VaultVaultRecordPrimerable)()
    SetPasswordRotation(value *ServiceaccountPasswordRotationScheme)()
    SetTechnicalAdministrator(value GroupGroupPrimerable)()
}
