package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningProvisionedInternalLDAP struct {
    ProvisioningProvisionedSystem
    // The client property
    client ClientLdapClientable
}
// NewProvisioningProvisionedInternalLDAP instantiates a new ProvisioningProvisionedInternalLDAP and sets the default values.
func NewProvisioningProvisionedInternalLDAP()(*ProvisioningProvisionedInternalLDAP) {
    m := &ProvisioningProvisionedInternalLDAP{
        ProvisioningProvisionedSystem: *NewProvisioningProvisionedSystem(),
    }
    typeEscapedValue := "provisioning.ProvisionedInternalLDAP"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningProvisionedInternalLDAPFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningProvisionedInternalLDAPFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningProvisionedInternalLDAP(), nil
}
// GetClient gets the client property value. The client property
// returns a ClientLdapClientable when successful
func (m *ProvisioningProvisionedInternalLDAP) GetClient()(ClientLdapClientable) {
    return m.client
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningProvisionedInternalLDAP) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningProvisionedSystem.GetFieldDeserializers()
    res["client"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientLdapClientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClient(val.(ClientLdapClientable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *ProvisioningProvisionedInternalLDAP) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningProvisionedSystem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("client", m.GetClient())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClient sets the client property value. The client property
func (m *ProvisioningProvisionedInternalLDAP) SetClient(value ClientLdapClientable)() {
    m.client = value
}
type ProvisioningProvisionedInternalLDAPable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningProvisionedSystemable
    GetClient()(ClientLdapClientable)
    SetClient(value ClientLdapClientable)()
}
