package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationOrganizationalUnitAccount struct {
    AuthAccountPrimer
    // The additionalObjects property
    additionalObjects OrganizationOrganizationalUnitAccount_additionalObjectsable
    // The directory property
    directory DirectoryAccountDirectoryPrimerable
}
// NewOrganizationOrganizationalUnitAccount instantiates a new OrganizationOrganizationalUnitAccount and sets the default values.
func NewOrganizationOrganizationalUnitAccount()(*OrganizationOrganizationalUnitAccount) {
    m := &OrganizationOrganizationalUnitAccount{
        AuthAccountPrimer: *NewAuthAccountPrimer(),
    }
    typeEscapedValue := "organization.OrganizationalUnitAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateOrganizationOrganizationalUnitAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationOrganizationalUnitAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOrganizationOrganizationalUnitAccount(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a OrganizationOrganizationalUnitAccount_additionalObjectsable when successful
func (m *OrganizationOrganizationalUnitAccount) GetAdditionalObjects()(OrganizationOrganizationalUnitAccount_additionalObjectsable) {
    return m.additionalObjects
}
// GetDirectory gets the directory property value. The directory property
// returns a DirectoryAccountDirectoryPrimerable when successful
func (m *OrganizationOrganizationalUnitAccount) GetDirectory()(DirectoryAccountDirectoryPrimerable) {
    return m.directory
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationOrganizationalUnitAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AuthAccountPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitAccount_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(OrganizationOrganizationalUnitAccount_additionalObjectsable))
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
    return res
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnitAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AuthAccountPrimer.Serialize(writer)
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
        err = writer.WriteObjectValue("directory", m.GetDirectory())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *OrganizationOrganizationalUnitAccount) SetAdditionalObjects(value OrganizationOrganizationalUnitAccount_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetDirectory sets the directory property value. The directory property
func (m *OrganizationOrganizationalUnitAccount) SetDirectory(value DirectoryAccountDirectoryPrimerable)() {
    m.directory = value
}
type OrganizationOrganizationalUnitAccountable interface {
    AuthAccountPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(OrganizationOrganizationalUnitAccount_additionalObjectsable)
    GetDirectory()(DirectoryAccountDirectoryPrimerable)
    SetAdditionalObjects(value OrganizationOrganizationalUnitAccount_additionalObjectsable)()
    SetDirectory(value DirectoryAccountDirectoryPrimerable)()
}
