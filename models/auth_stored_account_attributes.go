package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuthStoredAccountAttributes struct {
    NonLinkable
    // The attributes property
    attributes []AuthStoredAccountAttributeable
}
// NewAuthStoredAccountAttributes instantiates a new AuthStoredAccountAttributes and sets the default values.
func NewAuthStoredAccountAttributes()(*AuthStoredAccountAttributes) {
    m := &AuthStoredAccountAttributes{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "auth.StoredAccountAttributes"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuthStoredAccountAttributesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuthStoredAccountAttributesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthStoredAccountAttributes(), nil
}
// GetAttributes gets the attributes property value. The attributes property
// returns a []AuthStoredAccountAttributeable when successful
func (m *AuthStoredAccountAttributes) GetAttributes()([]AuthStoredAccountAttributeable) {
    return m.attributes
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuthStoredAccountAttributes) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["attributes"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAuthStoredAccountAttributeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AuthStoredAccountAttributeable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(AuthStoredAccountAttributeable)
                }
            }
            m.SetAttributes(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *AuthStoredAccountAttributes) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttributes() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttributes()))
        for i, v := range m.GetAttributes() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("attributes", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttributes sets the attributes property value. The attributes property
func (m *AuthStoredAccountAttributes) SetAttributes(value []AuthStoredAccountAttributeable)() {
    m.attributes = value
}
type AuthStoredAccountAttributesable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAttributes()([]AuthStoredAccountAttributeable)
    SetAttributes(value []AuthStoredAccountAttributeable)()
}
