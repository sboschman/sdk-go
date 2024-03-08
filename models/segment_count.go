package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type SegmentCount struct {
    NonLinkable
    // The count property
    count *int64
    // The name property
    name *string
}
// NewSegmentCount instantiates a new SegmentCount and sets the default values.
func NewSegmentCount()(*SegmentCount) {
    m := &SegmentCount{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "SegmentCount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateSegmentCountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSegmentCountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSegmentCount(), nil
}
// GetCount gets the count property value. The count property
// returns a *int64 when successful
func (m *SegmentCount) GetCount()(*int64) {
    return m.count
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SegmentCount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["count"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCount(val)
        }
        return nil
    }
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *SegmentCount) GetName()(*string) {
    return m.name
}
// Serialize serializes information the current object
func (m *SegmentCount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteInt64Value("count", m.GetCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCount sets the count property value. The count property
func (m *SegmentCount) SetCount(value *int64)() {
    m.count = value
}
// SetName sets the name property value. The name property
func (m *SegmentCount) SetName(value *string)() {
    m.name = value
}
type SegmentCountable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetCount()(*int64)
    GetName()(*string)
    SetCount(value *int64)()
    SetName(value *string)()
}
