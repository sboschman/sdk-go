package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupGroupAuditConfig struct {
    Linkable
    // The months property
    months []Month
}
// NewGroupGroupAuditConfig instantiates a new GroupGroupAuditConfig and sets the default values.
func NewGroupGroupAuditConfig()(*GroupGroupAuditConfig) {
    m := &GroupGroupAuditConfig{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "group.GroupAuditConfig"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAuditConfigFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupGroupAuditConfigFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAuditConfig(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupGroupAuditConfig) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["months"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseMonth)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Month, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*Month))
                }
            }
            m.SetMonths(res)
        }
        return nil
    }
    return res
}
// GetMonths gets the months property value. The months property
// returns a []Month when successful
func (m *GroupGroupAuditConfig) GetMonths()([]Month) {
    return m.months
}
// Serialize serializes information the current object
func (m *GroupGroupAuditConfig) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetMonths() != nil {
        err = writer.WriteCollectionOfStringValues("months", SerializeMonth(m.GetMonths()))
        if err != nil {
            return err
        }
    }
    return nil
}
// SetMonths sets the months property value. The months property
func (m *GroupGroupAuditConfig) SetMonths(value []Month)() {
    m.months = value
}
type GroupGroupAuditConfigable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetMonths()([]Month)
    SetMonths(value []Month)()
}
