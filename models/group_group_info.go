package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GroupGroupInfo 
type GroupGroupInfo struct {
    NonLinkable
    // The nrAccounts property
    nrAccounts *int32
    // The nrAccountsWithVault property
    nrAccountsWithVault *int32
    // The nrAudits property
    nrAudits *int32
    // The nrClients property
    nrClients *int32
    // The nrProvisionedSystems property
    nrProvisionedSystems *int32
    // The nrVaultRecords property
    nrVaultRecords *int32
}
// NewGroupGroupInfo instantiates a new groupGroupInfo and sets the default values.
func NewGroupGroupInfo()(*GroupGroupInfo) {
    m := &GroupGroupInfo{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupInfo"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGroupGroupInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupInfo(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GroupGroupInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["nrAccounts"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrAccounts(val)
        }
        return nil
    }
    res["nrAccountsWithVault"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrAccountsWithVault(val)
        }
        return nil
    }
    res["nrAudits"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrAudits(val)
        }
        return nil
    }
    res["nrClients"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrClients(val)
        }
        return nil
    }
    res["nrProvisionedSystems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrProvisionedSystems(val)
        }
        return nil
    }
    res["nrVaultRecords"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNrVaultRecords(val)
        }
        return nil
    }
    return res
}
// GetNrAccounts gets the nrAccounts property value. The nrAccounts property
func (m *GroupGroupInfo) GetNrAccounts()(*int32) {
    return m.nrAccounts
}
// GetNrAccountsWithVault gets the nrAccountsWithVault property value. The nrAccountsWithVault property
func (m *GroupGroupInfo) GetNrAccountsWithVault()(*int32) {
    return m.nrAccountsWithVault
}
// GetNrAudits gets the nrAudits property value. The nrAudits property
func (m *GroupGroupInfo) GetNrAudits()(*int32) {
    return m.nrAudits
}
// GetNrClients gets the nrClients property value. The nrClients property
func (m *GroupGroupInfo) GetNrClients()(*int32) {
    return m.nrClients
}
// GetNrProvisionedSystems gets the nrProvisionedSystems property value. The nrProvisionedSystems property
func (m *GroupGroupInfo) GetNrProvisionedSystems()(*int32) {
    return m.nrProvisionedSystems
}
// GetNrVaultRecords gets the nrVaultRecords property value. The nrVaultRecords property
func (m *GroupGroupInfo) GetNrVaultRecords()(*int32) {
    return m.nrVaultRecords
}
// Serialize serializes information the current object
func (m *GroupGroupInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
// SetNrAccounts sets the nrAccounts property value. The nrAccounts property
func (m *GroupGroupInfo) SetNrAccounts(value *int32)() {
    m.nrAccounts = value
}
// SetNrAccountsWithVault sets the nrAccountsWithVault property value. The nrAccountsWithVault property
func (m *GroupGroupInfo) SetNrAccountsWithVault(value *int32)() {
    m.nrAccountsWithVault = value
}
// SetNrAudits sets the nrAudits property value. The nrAudits property
func (m *GroupGroupInfo) SetNrAudits(value *int32)() {
    m.nrAudits = value
}
// SetNrClients sets the nrClients property value. The nrClients property
func (m *GroupGroupInfo) SetNrClients(value *int32)() {
    m.nrClients = value
}
// SetNrProvisionedSystems sets the nrProvisionedSystems property value. The nrProvisionedSystems property
func (m *GroupGroupInfo) SetNrProvisionedSystems(value *int32)() {
    m.nrProvisionedSystems = value
}
// SetNrVaultRecords sets the nrVaultRecords property value. The nrVaultRecords property
func (m *GroupGroupInfo) SetNrVaultRecords(value *int32)() {
    m.nrVaultRecords = value
}
// GroupGroupInfoable 
type GroupGroupInfoable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNrAccounts()(*int32)
    GetNrAccountsWithVault()(*int32)
    GetNrAudits()(*int32)
    GetNrClients()(*int32)
    GetNrProvisionedSystems()(*int32)
    GetNrVaultRecords()(*int32)
    SetNrAccounts(value *int32)()
    SetNrAccountsWithVault(value *int32)()
    SetNrAudits(value *int32)()
    SetNrClients(value *int32)()
    SetNrProvisionedSystems(value *int32)()
    SetNrVaultRecords(value *int32)()
}
