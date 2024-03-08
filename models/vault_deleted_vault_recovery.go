package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type VaultDeletedVaultRecovery struct {
    NonLinkable
    // The prefix property
    prefix *string
    // The privateKey property
    privateKey *string
    // The targetGroup property
    targetGroup GroupGroupPrimerable
}
// NewVaultDeletedVaultRecovery instantiates a new VaultDeletedVaultRecovery and sets the default values.
func NewVaultDeletedVaultRecovery()(*VaultDeletedVaultRecovery) {
    m := &VaultDeletedVaultRecovery{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "vault.DeletedVaultRecovery"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateVaultDeletedVaultRecoveryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateVaultDeletedVaultRecoveryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewVaultDeletedVaultRecovery(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *VaultDeletedVaultRecovery) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["prefix"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrefix(val)
        }
        return nil
    }
    res["privateKey"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrivateKey(val)
        }
        return nil
    }
    res["targetGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    return res
}
// GetPrefix gets the prefix property value. The prefix property
// returns a *string when successful
func (m *VaultDeletedVaultRecovery) GetPrefix()(*string) {
    return m.prefix
}
// GetPrivateKey gets the privateKey property value. The privateKey property
// returns a *string when successful
func (m *VaultDeletedVaultRecovery) GetPrivateKey()(*string) {
    return m.privateKey
}
// GetTargetGroup gets the targetGroup property value. The targetGroup property
// returns a GroupGroupPrimerable when successful
func (m *VaultDeletedVaultRecovery) GetTargetGroup()(GroupGroupPrimerable) {
    return m.targetGroup
}
// Serialize serializes information the current object
func (m *VaultDeletedVaultRecovery) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("prefix", m.GetPrefix())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("privateKey", m.GetPrivateKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetGroup", m.GetTargetGroup())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPrefix sets the prefix property value. The prefix property
func (m *VaultDeletedVaultRecovery) SetPrefix(value *string)() {
    m.prefix = value
}
// SetPrivateKey sets the privateKey property value. The privateKey property
func (m *VaultDeletedVaultRecovery) SetPrivateKey(value *string)() {
    m.privateKey = value
}
// SetTargetGroup sets the targetGroup property value. The targetGroup property
func (m *VaultDeletedVaultRecovery) SetTargetGroup(value GroupGroupPrimerable)() {
    m.targetGroup = value
}
type VaultDeletedVaultRecoveryable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetPrefix()(*string)
    GetPrivateKey()(*string)
    GetTargetGroup()(GroupGroupPrimerable)
    SetPrefix(value *string)()
    SetPrivateKey(value *string)()
    SetTargetGroup(value GroupGroupPrimerable)()
}
