package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotificationGroupEditRequiredNotification struct {
    NotificationNotification
    // The group property
    group GroupGroupable
    // The parameters property
    parameters NotificationGroupEditRequiredNotification_parametersable
    // The reason property
    reason *GroupGroupEditRequiredReason
}
// NewNotificationGroupEditRequiredNotification instantiates a new NotificationGroupEditRequiredNotification and sets the default values.
func NewNotificationGroupEditRequiredNotification()(*NotificationGroupEditRequiredNotification) {
    m := &NotificationGroupEditRequiredNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.GroupEditRequiredNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationGroupEditRequiredNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotificationGroupEditRequiredNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationGroupEditRequiredNotification(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotificationGroupEditRequiredNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["group"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroup(val.(GroupGroupable))
        }
        return nil
    }
    res["parameters"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateNotificationGroupEditRequiredNotification_parametersFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParameters(val.(NotificationGroupEditRequiredNotification_parametersable))
        }
        return nil
    }
    res["reason"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupGroupEditRequiredReason)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReason(val.(*GroupGroupEditRequiredReason))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
// returns a GroupGroupable when successful
func (m *NotificationGroupEditRequiredNotification) GetGroup()(GroupGroupable) {
    return m.group
}
// GetParameters gets the parameters property value. The parameters property
// returns a NotificationGroupEditRequiredNotification_parametersable when successful
func (m *NotificationGroupEditRequiredNotification) GetParameters()(NotificationGroupEditRequiredNotification_parametersable) {
    return m.parameters
}
// GetReason gets the reason property value. The reason property
// returns a *GroupGroupEditRequiredReason when successful
func (m *NotificationGroupEditRequiredNotification) GetReason()(*GroupGroupEditRequiredReason) {
    return m.reason
}
// Serialize serializes information the current object
func (m *NotificationGroupEditRequiredNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parameters", m.GetParameters())
        if err != nil {
            return err
        }
    }
    if m.GetReason() != nil {
        cast := (*m.GetReason()).String()
        err = writer.WriteStringValue("reason", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroup sets the group property value. The group property
func (m *NotificationGroupEditRequiredNotification) SetGroup(value GroupGroupable)() {
    m.group = value
}
// SetParameters sets the parameters property value. The parameters property
func (m *NotificationGroupEditRequiredNotification) SetParameters(value NotificationGroupEditRequiredNotification_parametersable)() {
    m.parameters = value
}
// SetReason sets the reason property value. The reason property
func (m *NotificationGroupEditRequiredNotification) SetReason(value *GroupGroupEditRequiredReason)() {
    m.reason = value
}
type NotificationGroupEditRequiredNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetGroup()(GroupGroupable)
    GetParameters()(NotificationGroupEditRequiredNotification_parametersable)
    GetReason()(*GroupGroupEditRequiredReason)
    SetGroup(value GroupGroupable)()
    SetParameters(value NotificationGroupEditRequiredNotification_parametersable)()
    SetReason(value *GroupGroupEditRequiredReason)()
}
