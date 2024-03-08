package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestAbstractApplicationModificationRequest struct {
    RequestModificationRequest
    // The application property
    application ClientClientApplicationPrimerable
}
// NewRequestAbstractApplicationModificationRequest instantiates a new RequestAbstractApplicationModificationRequest and sets the default values.
func NewRequestAbstractApplicationModificationRequest()(*RequestAbstractApplicationModificationRequest) {
    m := &RequestAbstractApplicationModificationRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.AbstractApplicationModificationRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestAbstractApplicationModificationRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestAbstractApplicationModificationRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("$type")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
                switch *mappingValue {
                    case "request.GrantApplicationRequest":
                        return NewRequestGrantApplicationRequest(), nil
                    case "request.GrantClientPermissionRequest":
                        return NewRequestGrantClientPermissionRequest(), nil
                    case "request.TransferApplicationAdministrationRequest":
                        return NewRequestTransferApplicationAdministrationRequest(), nil
                    case "request.TransferApplicationOwnershipRequest":
                        return NewRequestTransferApplicationOwnershipRequest(), nil
                }
            }
        }
    }
    return NewRequestAbstractApplicationModificationRequest(), nil
}
// GetApplication gets the application property value. The application property
// returns a ClientClientApplicationPrimerable when successful
func (m *RequestAbstractApplicationModificationRequest) GetApplication()(ClientClientApplicationPrimerable) {
    return m.application
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestAbstractApplicationModificationRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["application"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientClientApplicationPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetApplication(val.(ClientClientApplicationPrimerable))
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *RequestAbstractApplicationModificationRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("application", m.GetApplication())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetApplication sets the application property value. The application property
func (m *RequestAbstractApplicationModificationRequest) SetApplication(value ClientClientApplicationPrimerable)() {
    m.application = value
}
type RequestAbstractApplicationModificationRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetApplication()(ClientClientApplicationPrimerable)
    SetApplication(value ClientClientApplicationPrimerable)()
}
