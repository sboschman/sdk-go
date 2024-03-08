package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultVaultUnlock struct {
    NonLinkable
    // The password property
    password *string
}
// NewVaultVaultUnlock instantiates a new VaultVaultUnlock and sets the default values.
func NewVaultVaultUnlock()(*VaultVaultUnlock) {
    m := &VaultVaultUnlock{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.VaultUnlock"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultVaultUnlockFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultVaultUnlockFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultVaultUnlock(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultVaultUnlock) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["password"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPassword(val)
        }
        return nil
    }
    return res
}
// GetPassword gets the password property value. The password property
// returns a *string when successful
func (m *VaultVaultUnlock) GetPassword()(*string) {
    return m.password
}
// Serialize serializes information the current object
func (m *VaultVaultUnlock) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("password", m.GetPassword())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPassword sets the password property value. The password property
func (m *VaultVaultUnlock) SetPassword(value *string)() {
    m.password = value
}
type VaultVaultUnlockable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPassword()(*string)
    SetPassword(value *string)()
}
