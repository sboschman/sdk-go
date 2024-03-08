package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupGroupAccountStatus struct {
    NonLinkable
    // The folder property
    folder GroupGroupFolderable
    // The visibleForProvisioning property
    visibleForProvisioning *bool
}
// NewGroupGroupAccountStatus instantiates a new GroupGroupAccountStatus and sets the default values.
func NewGroupGroupAccountStatus()(*GroupGroupAccountStatus) {
    m := &GroupGroupAccountStatus{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupAccountStatus"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAccountStatusFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupGroupAccountStatusFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAccountStatus(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupGroupAccountStatus) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["folder"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupFolderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFolder(val.(GroupGroupFolderable))
        }
        return nil
    }
    res["visibleForProvisioning"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVisibleForProvisioning(val)
        }
        return nil
    }
    return res
}
// GetFolder gets the folder property value. The folder property
// returns a GroupGroupFolderable when successful
func (m *GroupGroupAccountStatus) GetFolder()(GroupGroupFolderable) {
    return m.folder
}
// GetVisibleForProvisioning gets the visibleForProvisioning property value. The visibleForProvisioning property
// returns a *bool when successful
func (m *GroupGroupAccountStatus) GetVisibleForProvisioning()(*bool) {
    return m.visibleForProvisioning
}
// Serialize serializes information the current object
func (m *GroupGroupAccountStatus) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("folder", m.GetFolder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("visibleForProvisioning", m.GetVisibleForProvisioning())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFolder sets the folder property value. The folder property
func (m *GroupGroupAccountStatus) SetFolder(value GroupGroupFolderable)() {
    m.folder = value
}
// SetVisibleForProvisioning sets the visibleForProvisioning property value. The visibleForProvisioning property
func (m *GroupGroupAccountStatus) SetVisibleForProvisioning(value *bool)() {
    m.visibleForProvisioning = value
}
type GroupGroupAccountStatusable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetFolder()(GroupGroupFolderable)
    GetVisibleForProvisioning()(*bool)
    SetFolder(value GroupGroupFolderable)()
    SetVisibleForProvisioning(value *bool)()
}
