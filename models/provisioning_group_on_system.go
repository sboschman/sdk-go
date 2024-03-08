package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ProvisioningGroupOnSystem struct {
    ProvisioningGroupOnSystemPrimer
    // The additionalObjects property
    additionalObjects ProvisioningGroupOnSystem_additionalObjectsable
    // The owner property
    owner GroupGroupPrimerable
    // The system property
    system ProvisioningProvisionedSystemPrimerable
}
// NewProvisioningGroupOnSystem instantiates a new ProvisioningGroupOnSystem and sets the default values.
func NewProvisioningGroupOnSystem()(*ProvisioningGroupOnSystem) {
    m := &ProvisioningGroupOnSystem{
        ProvisioningGroupOnSystemPrimer: *NewProvisioningGroupOnSystemPrimer(),
    }
    typeEscapedValue := "provisioning.GroupOnSystem"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningGroupOnSystemFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateProvisioningGroupOnSystemFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningGroupOnSystem(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ProvisioningGroupOnSystem_additionalObjectsable when successful
func (m *ProvisioningGroupOnSystem) GetAdditionalObjects()(ProvisioningGroupOnSystem_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ProvisioningGroupOnSystem) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ProvisioningGroupOnSystemPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningGroupOnSystem_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ProvisioningGroupOnSystem_additionalObjectsable))
        }
        return nil
    }
    res["owner"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOwner(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["system"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningProvisionedSystemPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSystem(val.(ProvisioningProvisionedSystemPrimerable))
        }
        return nil
    }
    return res
}
// GetOwner gets the owner property value. The owner property
// returns a GroupGroupPrimerable when successful
func (m *ProvisioningGroupOnSystem) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// GetSystem gets the system property value. The system property
// returns a ProvisioningProvisionedSystemPrimerable when successful
func (m *ProvisioningGroupOnSystem) GetSystem()(ProvisioningProvisionedSystemPrimerable) {
    return m.system
}
// Serialize serializes information the current object
func (m *ProvisioningGroupOnSystem) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ProvisioningGroupOnSystemPrimer.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("additionalObjects", m.GetAdditionalObjects())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("owner", m.GetOwner())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("system", m.GetSystem())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ProvisioningGroupOnSystem) SetAdditionalObjects(value ProvisioningGroupOnSystem_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetOwner sets the owner property value. The owner property
func (m *ProvisioningGroupOnSystem) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
// SetSystem sets the system property value. The system property
func (m *ProvisioningGroupOnSystem) SetSystem(value ProvisioningProvisionedSystemPrimerable)() {
    m.system = value
}
type ProvisioningGroupOnSystemable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    ProvisioningGroupOnSystemPrimerable
    GetAdditionalObjects()(ProvisioningGroupOnSystem_additionalObjectsable)
    GetOwner()(GroupGroupPrimerable)
    GetSystem()(ProvisioningProvisionedSystemPrimerable)
    SetAdditionalObjects(value ProvisioningGroupOnSystem_additionalObjectsable)()
    SetOwner(value GroupGroupPrimerable)()
    SetSystem(value ProvisioningProvisionedSystemPrimerable)()
}
