package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LaunchpadDisplayedLaunchpadTiles struct {
    NonLinkable
    // The items property
    items []LaunchpadDisplayedLaunchpadTileable
}
// NewLaunchpadDisplayedLaunchpadTiles instantiates a new LaunchpadDisplayedLaunchpadTiles and sets the default values.
func NewLaunchpadDisplayedLaunchpadTiles()(*LaunchpadDisplayedLaunchpadTiles) {
    m := &LaunchpadDisplayedLaunchpadTiles{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "launchpad.DisplayedLaunchpadTiles"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateLaunchpadDisplayedLaunchpadTilesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLaunchpadDisplayedLaunchpadTilesFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLaunchpadDisplayedLaunchpadTiles(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LaunchpadDisplayedLaunchpadTiles) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["items"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLaunchpadDisplayedLaunchpadTileFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LaunchpadDisplayedLaunchpadTileable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(LaunchpadDisplayedLaunchpadTileable)
                }
            }
            m.SetItems(res)
        }
        return nil
    }
    return res
}
// GetItems gets the items property value. The items property
// returns a []LaunchpadDisplayedLaunchpadTileable when successful
func (m *LaunchpadDisplayedLaunchpadTiles) GetItems()([]LaunchpadDisplayedLaunchpadTileable) {
    return m.items
}
// Serialize serializes information the current object
func (m *LaunchpadDisplayedLaunchpadTiles) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetItems()))
        for i, v := range m.GetItems() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("items", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetItems sets the items property value. The items property
func (m *LaunchpadDisplayedLaunchpadTiles) SetItems(value []LaunchpadDisplayedLaunchpadTileable)() {
    m.items = value
}
type LaunchpadDisplayedLaunchpadTilesable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetItems()([]LaunchpadDisplayedLaunchpadTileable)
    SetItems(value []LaunchpadDisplayedLaunchpadTileable)()
}
