package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type GroupGroupAdminMail struct {
    NonLinkable
    // The body property
    body *string
}
// NewGroupGroupAdminMail instantiates a new GroupGroupAdminMail and sets the default values.
func NewGroupGroupAdminMail()(*GroupGroupAdminMail) {
    m := &GroupGroupAdminMail{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "group.GroupAdminMail"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateGroupGroupAdminMailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateGroupGroupAdminMailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGroupGroupAdminMail(), nil
}
// GetBody gets the body property value. The body property
// returns a *string when successful
func (m *GroupGroupAdminMail) GetBody()(*string) {
    return m.body
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *GroupGroupAdminMail) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *GroupGroupAdminMail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBody sets the body property value. The body property
func (m *GroupGroupAdminMail) SetBody(value *string)() {
    m.body = value
}
type GroupGroupAdminMailable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBody()(*string)
    SetBody(value *string)()
}
