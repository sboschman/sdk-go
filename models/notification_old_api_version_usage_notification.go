package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type NotificationOldApiVersionUsageNotification struct {
    NotificationNotification
    // The apiVersion property
    apiVersion *int32
    // The client property
    client ClientOAuth2Clientable
}
// NewNotificationOldApiVersionUsageNotification instantiates a new NotificationOldApiVersionUsageNotification and sets the default values.
func NewNotificationOldApiVersionUsageNotification()(*NotificationOldApiVersionUsageNotification) {
    m := &NotificationOldApiVersionUsageNotification{
        NotificationNotification: *NewNotificationNotification(),
    }
    typeEscapedValue := "notification.OldApiVersionUsageNotification"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateNotificationOldApiVersionUsageNotificationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateNotificationOldApiVersionUsageNotificationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewNotificationOldApiVersionUsageNotification(), nil
}
// GetApiVersion gets the apiVersion property value. The apiVersion property
// returns a *int32 when successful
func (m *NotificationOldApiVersionUsageNotification) GetApiVersion()(*int32) {
    return m.apiVersion
}
// GetClient gets the client property value. The client property
// returns a ClientOAuth2Clientable when successful
func (m *NotificationOldApiVersionUsageNotification) GetClient()(ClientOAuth2Clientable) {
    return m.client
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *NotificationOldApiVersionUsageNotification) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NotificationNotification.GetFieldDeserializers()
    res["apiVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApiVersion(val)
        }
        return nil
    }
    res["client"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientOAuth2ClientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClient(val.(ClientOAuth2Clientable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *NotificationOldApiVersionUsageNotification) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NotificationNotification.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt32Value("apiVersion", m.GetApiVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("client", m.GetClient())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApiVersion sets the apiVersion property value. The apiVersion property
func (m *NotificationOldApiVersionUsageNotification) SetApiVersion(value *int32)() {
    m.apiVersion = value
}
// SetClient sets the client property value. The client property
func (m *NotificationOldApiVersionUsageNotification) SetClient(value ClientOAuth2Clientable)() {
    m.client = value
}
type NotificationOldApiVersionUsageNotificationable interface {
    NotificationNotificationable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetApiVersion()(*int32)
    GetClient()(ClientOAuth2Clientable)
    SetApiVersion(value *int32)()
    SetClient(value ClientOAuth2Clientable)()
}
