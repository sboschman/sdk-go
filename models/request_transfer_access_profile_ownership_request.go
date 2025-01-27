package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestTransferAccessProfileOwnershipRequest struct {
    RequestModificationRequest
    // The accessProfile property
    accessProfile ProfileAccessProfilePrimerable
}
// NewRequestTransferAccessProfileOwnershipRequest instantiates a new RequestTransferAccessProfileOwnershipRequest and sets the default values.
func NewRequestTransferAccessProfileOwnershipRequest()(*RequestTransferAccessProfileOwnershipRequest) {
    m := &RequestTransferAccessProfileOwnershipRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.TransferAccessProfileOwnershipRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestTransferAccessProfileOwnershipRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestTransferAccessProfileOwnershipRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestTransferAccessProfileOwnershipRequest(), nil
}
// GetAccessProfile gets the accessProfile property value. The accessProfile property
// returns a ProfileAccessProfilePrimerable when successful
func (m *RequestTransferAccessProfileOwnershipRequest) GetAccessProfile()(ProfileAccessProfilePrimerable) {
    return m.accessProfile
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestTransferAccessProfileOwnershipRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["accessProfile"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProfileAccessProfilePrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccessProfile(val.(ProfileAccessProfilePrimerable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestTransferAccessProfileOwnershipRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accessProfile", m.GetAccessProfile())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessProfile sets the accessProfile property value. The accessProfile property
func (m *RequestTransferAccessProfileOwnershipRequest) SetAccessProfile(value ProfileAccessProfilePrimerable)() {
    m.accessProfile = value
}
type RequestTransferAccessProfileOwnershipRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetAccessProfile()(ProfileAccessProfilePrimerable)
    SetAccessProfile(value ProfileAccessProfilePrimerable)()
}
