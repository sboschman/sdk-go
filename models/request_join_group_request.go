package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestJoinGroupRequest struct {
    RequestModificationRequest
}
// NewRequestJoinGroupRequest instantiates a new RequestJoinGroupRequest and sets the default values.
func NewRequestJoinGroupRequest()(*RequestJoinGroupRequest) {
    m := &RequestJoinGroupRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.JoinGroupRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestJoinGroupRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestJoinGroupRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestJoinGroupRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestJoinGroupRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *RequestJoinGroupRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
type RequestJoinGroupRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
}
