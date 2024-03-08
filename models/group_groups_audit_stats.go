package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupGroupsAuditStats struct {
    NonLinkable
    // The auditedGroups property
    auditedGroups *int64
    // The classifications property
    classifications []GroupGroupClassificationable
    // The overdueAuditGroups property
    overdueAuditGroups *int64
    // The overdueVaultPolicies property
    overdueVaultPolicies *int64
    // The unauditedGroups property
    unauditedGroups *int64
    // The validVaultPolicies property
    validVaultPolicies *int64
    // The vaultsWithoutPolicies property
    vaultsWithoutPolicies *int64
}
// NewGroupGroupsAuditStats instantiates a new GroupGroupsAuditStats and sets the default values.
func NewGroupGroupsAuditStats()(*GroupGroupsAuditStats) {
    m := &GroupGroupsAuditStats{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupsAuditStats"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupsAuditStatsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupGroupsAuditStatsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupsAuditStats(), nil
}
// GetAuditedGroups gets the auditedGroups property value. The auditedGroups property
// returns a *int64 when successful
func (m *GroupGroupsAuditStats) GetAuditedGroups()(*int64) {
    return m.auditedGroups
}
// GetClassifications gets the classifications property value. The classifications property
// returns a []GroupGroupClassificationable when successful
func (m *GroupGroupsAuditStats) GetClassifications()([]GroupGroupClassificationable) {
    return m.classifications
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupGroupsAuditStats) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["auditedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAuditedGroups(val)
        }
        return nil
    }
    res["classifications"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupGroupClassificationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupGroupClassificationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupGroupClassificationable)
                }
            }
            m.SetClassifications(res)
        }
        return nil
    }
    res["overdueAuditGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverdueAuditGroups(val)
        }
        return nil
    }
    res["overdueVaultPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOverdueVaultPolicies(val)
        }
        return nil
    }
    res["unauditedGroups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUnauditedGroups(val)
        }
        return nil
    }
    res["validVaultPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValidVaultPolicies(val)
        }
        return nil
    }
    res["vaultsWithoutPolicies"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVaultsWithoutPolicies(val)
        }
        return nil
    }
    return res
}
// GetOverdueAuditGroups gets the overdueAuditGroups property value. The overdueAuditGroups property
// returns a *int64 when successful
func (m *GroupGroupsAuditStats) GetOverdueAuditGroups()(*int64) {
    return m.overdueAuditGroups
}
// GetOverdueVaultPolicies gets the overdueVaultPolicies property value. The overdueVaultPolicies property
// returns a *int64 when successful
func (m *GroupGroupsAuditStats) GetOverdueVaultPolicies()(*int64) {
    return m.overdueVaultPolicies
}
// GetUnauditedGroups gets the unauditedGroups property value. The unauditedGroups property
// returns a *int64 when successful
func (m *GroupGroupsAuditStats) GetUnauditedGroups()(*int64) {
    return m.unauditedGroups
}
// GetValidVaultPolicies gets the validVaultPolicies property value. The validVaultPolicies property
// returns a *int64 when successful
func (m *GroupGroupsAuditStats) GetValidVaultPolicies()(*int64) {
    return m.validVaultPolicies
}
// GetVaultsWithoutPolicies gets the vaultsWithoutPolicies property value. The vaultsWithoutPolicies property
// returns a *int64 when successful
func (m *GroupGroupsAuditStats) GetVaultsWithoutPolicies()(*int64) {
    return m.vaultsWithoutPolicies
}
// Serialize serializes information the current object
func (m *GroupGroupsAuditStats) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("auditedGroups", m.GetAuditedGroups())
        if err != nil {
            return err
        }
    }
    if m.GetClassifications() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetClassifications()))
        for i, v := range m.GetClassifications() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("classifications", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("overdueAuditGroups", m.GetOverdueAuditGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("overdueVaultPolicies", m.GetOverdueVaultPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("unauditedGroups", m.GetUnauditedGroups())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("validVaultPolicies", m.GetValidVaultPolicies())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("vaultsWithoutPolicies", m.GetVaultsWithoutPolicies())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAuditedGroups sets the auditedGroups property value. The auditedGroups property
func (m *GroupGroupsAuditStats) SetAuditedGroups(value *int64)() {
    m.auditedGroups = value
}
// SetClassifications sets the classifications property value. The classifications property
func (m *GroupGroupsAuditStats) SetClassifications(value []GroupGroupClassificationable)() {
    m.classifications = value
}
// SetOverdueAuditGroups sets the overdueAuditGroups property value. The overdueAuditGroups property
func (m *GroupGroupsAuditStats) SetOverdueAuditGroups(value *int64)() {
    m.overdueAuditGroups = value
}
// SetOverdueVaultPolicies sets the overdueVaultPolicies property value. The overdueVaultPolicies property
func (m *GroupGroupsAuditStats) SetOverdueVaultPolicies(value *int64)() {
    m.overdueVaultPolicies = value
}
// SetUnauditedGroups sets the unauditedGroups property value. The unauditedGroups property
func (m *GroupGroupsAuditStats) SetUnauditedGroups(value *int64)() {
    m.unauditedGroups = value
}
// SetValidVaultPolicies sets the validVaultPolicies property value. The validVaultPolicies property
func (m *GroupGroupsAuditStats) SetValidVaultPolicies(value *int64)() {
    m.validVaultPolicies = value
}
// SetVaultsWithoutPolicies sets the vaultsWithoutPolicies property value. The vaultsWithoutPolicies property
func (m *GroupGroupsAuditStats) SetVaultsWithoutPolicies(value *int64)() {
    m.vaultsWithoutPolicies = value
}
type GroupGroupsAuditStatsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAuditedGroups()(*int64)
    GetClassifications()([]GroupGroupClassificationable)
    GetOverdueAuditGroups()(*int64)
    GetOverdueVaultPolicies()(*int64)
    GetUnauditedGroups()(*int64)
    GetValidVaultPolicies()(*int64)
    GetVaultsWithoutPolicies()(*int64)
    SetAuditedGroups(value *int64)()
    SetClassifications(value []GroupGroupClassificationable)()
    SetOverdueAuditGroups(value *int64)()
    SetOverdueVaultPolicies(value *int64)()
    SetUnauditedGroups(value *int64)()
    SetValidVaultPolicies(value *int64)()
    SetVaultsWithoutPolicies(value *int64)()
}
