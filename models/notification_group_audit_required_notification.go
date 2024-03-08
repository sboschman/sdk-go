package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotificationGroupAuditRequiredNotification struct {
    NotificationNotification
    // The dueDate property
    dueDate *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The group property
    group GroupGroupable
    // The type property
    notificationGroupAuditRequiredNotificationType *AuditAuditRequiredSourceType
}
// NewNotificationGroupAuditRequiredNotification instantiates a new NotificationGroupAuditRequiredNotification and sets the default values.
func NewNotificationGroupAuditRequiredNotification()(*NotificationGroupAuditRequiredNotification) {
    m := &NotificationGroupAuditRequiredNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.GroupAuditRequiredNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationGroupAuditRequiredNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotificationGroupAuditRequiredNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationGroupAuditRequiredNotification(), nil
}
// GetDueDate gets the dueDate property value. The dueDate property
// returns a *Time when successful
func (m *NotificationGroupAuditRequiredNotification) GetDueDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.dueDate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotificationGroupAuditRequiredNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["dueDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDate(val)
        }
        return nil
    }
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
    res["type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditAuditRequiredSourceType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationGroupAuditRequiredNotificationType(val.(*AuditAuditRequiredSourceType))
        }
        return nil
    }
    return res
}
// GetGroup gets the group property value. The group property
// returns a GroupGroupable when successful
func (m *NotificationGroupAuditRequiredNotification) GetGroup()(GroupGroupable) {
    return m.group
}
// GetNotificationGroupAuditRequiredNotificationType gets the type property value. The type property
// returns a *AuditAuditRequiredSourceType when successful
func (m *NotificationGroupAuditRequiredNotification) GetNotificationGroupAuditRequiredNotificationType()(*AuditAuditRequiredSourceType) {
    return m.notificationGroupAuditRequiredNotificationType
}
// Serialize serializes information the current object
func (m *NotificationGroupAuditRequiredNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("dueDate", m.GetDueDate())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("group", m.GetGroup())
        if err != nil {
            return err
        }
    }
    if m.GetNotificationGroupAuditRequiredNotificationType() != nil {
        cast := (*m.GetNotificationGroupAuditRequiredNotificationType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDueDate sets the dueDate property value. The dueDate property
func (m *NotificationGroupAuditRequiredNotification) SetDueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.dueDate = value
}
// SetGroup sets the group property value. The group property
func (m *NotificationGroupAuditRequiredNotification) SetGroup(value GroupGroupable)() {
    m.group = value
}
// SetNotificationGroupAuditRequiredNotificationType sets the type property value. The type property
func (m *NotificationGroupAuditRequiredNotification) SetNotificationGroupAuditRequiredNotificationType(value *AuditAuditRequiredSourceType)() {
    m.notificationGroupAuditRequiredNotificationType = value
}
type NotificationGroupAuditRequiredNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetDueDate()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetGroup()(GroupGroupable)
    GetNotificationGroupAuditRequiredNotificationType()(*AuditAuditRequiredSourceType)
    SetDueDate(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetGroup(value GroupGroupable)()
    SetNotificationGroupAuditRequiredNotificationType(value *AuditAuditRequiredSourceType)()
}
