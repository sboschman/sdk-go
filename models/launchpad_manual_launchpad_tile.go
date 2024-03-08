package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type LaunchpadManualLaunchpadTile struct {
    LaunchpadLaunchpadTile
    // The title property
    title *string
    // The uri property
    uri *string
}
// NewLaunchpadManualLaunchpadTile instantiates a new LaunchpadManualLaunchpadTile and sets the default values.
func NewLaunchpadManualLaunchpadTile()(*LaunchpadManualLaunchpadTile) {
    m := &LaunchpadManualLaunchpadTile{
        LaunchpadLaunchpadTile: *NewLaunchpadLaunchpadTile(),
    }
    typeEscapedValue := "launchpad.ManualLaunchpadTile"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateLaunchpadManualLaunchpadTileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateLaunchpadManualLaunchpadTileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewLaunchpadManualLaunchpadTile(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *LaunchpadManualLaunchpadTile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.LaunchpadLaunchpadTile.GetFieldDeserializers()
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["uri"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUri(val)
        }
        return nil
    }
    return res
}
// GetTitle gets the title property value. The title property
// returns a *string when successful
func (m *LaunchpadManualLaunchpadTile) GetTitle()(*string) {
    return m.title
}
// GetUri gets the uri property value. The uri property
// returns a *string when successful
func (m *LaunchpadManualLaunchpadTile) GetUri()(*string) {
    return m.uri
}
// Serialize serializes information the current object
func (m *LaunchpadManualLaunchpadTile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.LaunchpadLaunchpadTile.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("uri", m.GetUri())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTitle sets the title property value. The title property
func (m *LaunchpadManualLaunchpadTile) SetTitle(value *string)() {
    m.title = value
}
// SetUri sets the uri property value. The uri property
func (m *LaunchpadManualLaunchpadTile) SetUri(value *string)() {
    m.uri = value
}
type LaunchpadManualLaunchpadTileable interface {
    LaunchpadLaunchpadTileable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetTitle()(*string)
    GetUri()(*string)
    SetTitle(value *string)()
    SetUri(value *string)()
}
