package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningCircuitBreakerStatistics 
type ProvisioningCircuitBreakerStatistics struct {
    NonLinkable
    // The numberOfFailedCalls property
    numberOfFailedCalls *int64
    // The numberOfNotPermittedCalls property
    numberOfNotPermittedCalls *int64
    // The numberOfSuccessfulCalls property
    numberOfSuccessfulCalls *int64
    // The state property
    state *ProvisioningCircuitBreakerState
}
// NewProvisioningCircuitBreakerStatistics instantiates a new provisioningCircuitBreakerStatistics and sets the default values.
func NewProvisioningCircuitBreakerStatistics()(*ProvisioningCircuitBreakerStatistics) {
    m := &ProvisioningCircuitBreakerStatistics{
        NonLinkable: *NewNonLinkable(),
    }
    typeEscapedValue := "provisioning.CircuitBreakerStatistics"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateProvisioningCircuitBreakerStatisticsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningCircuitBreakerStatisticsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningCircuitBreakerStatistics(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningCircuitBreakerStatistics) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.NonLinkable.GetFieldDeserializers()
    res["numberOfFailedCalls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfFailedCalls(val)
        }
        return nil
    }
    res["numberOfNotPermittedCalls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfNotPermittedCalls(val)
        }
        return nil
    }
    res["numberOfSuccessfulCalls"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNumberOfSuccessfulCalls(val)
        }
        return nil
    }
    res["state"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningCircuitBreakerState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ProvisioningCircuitBreakerState))
        }
        return nil
    }
    return res
}
// GetNumberOfFailedCalls gets the numberOfFailedCalls property value. The numberOfFailedCalls property
func (m *ProvisioningCircuitBreakerStatistics) GetNumberOfFailedCalls()(*int64) {
    return m.numberOfFailedCalls
}
// GetNumberOfNotPermittedCalls gets the numberOfNotPermittedCalls property value. The numberOfNotPermittedCalls property
func (m *ProvisioningCircuitBreakerStatistics) GetNumberOfNotPermittedCalls()(*int64) {
    return m.numberOfNotPermittedCalls
}
// GetNumberOfSuccessfulCalls gets the numberOfSuccessfulCalls property value. The numberOfSuccessfulCalls property
func (m *ProvisioningCircuitBreakerStatistics) GetNumberOfSuccessfulCalls()(*int64) {
    return m.numberOfSuccessfulCalls
}
// GetState gets the state property value. The state property
func (m *ProvisioningCircuitBreakerStatistics) GetState()(*ProvisioningCircuitBreakerState) {
    return m.state
}
// Serialize serializes information the current object
func (m *ProvisioningCircuitBreakerStatistics) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.NonLinkable.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetNumberOfFailedCalls sets the numberOfFailedCalls property value. The numberOfFailedCalls property
func (m *ProvisioningCircuitBreakerStatistics) SetNumberOfFailedCalls(value *int64)() {
    m.numberOfFailedCalls = value
}
// SetNumberOfNotPermittedCalls sets the numberOfNotPermittedCalls property value. The numberOfNotPermittedCalls property
func (m *ProvisioningCircuitBreakerStatistics) SetNumberOfNotPermittedCalls(value *int64)() {
    m.numberOfNotPermittedCalls = value
}
// SetNumberOfSuccessfulCalls sets the numberOfSuccessfulCalls property value. The numberOfSuccessfulCalls property
func (m *ProvisioningCircuitBreakerStatistics) SetNumberOfSuccessfulCalls(value *int64)() {
    m.numberOfSuccessfulCalls = value
}
// SetState sets the state property value. The state property
func (m *ProvisioningCircuitBreakerStatistics) SetState(value *ProvisioningCircuitBreakerState)() {
    m.state = value
}
// ProvisioningCircuitBreakerStatisticsable 
type ProvisioningCircuitBreakerStatisticsable interface {
    NonLinkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetNumberOfFailedCalls()(*int64)
    GetNumberOfNotPermittedCalls()(*int64)
    GetNumberOfSuccessfulCalls()(*int64)
    GetState()(*ProvisioningCircuitBreakerState)
    SetNumberOfFailedCalls(value *int64)()
    SetNumberOfNotPermittedCalls(value *int64)()
    SetNumberOfSuccessfulCalls(value *int64)()
    SetState(value *ProvisioningCircuitBreakerState)()
}
