package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestMoveGroupsRequest struct {
    RequestAbstractOrganizationalUnitModificationRequest
    // The groups property
    groups []GroupGroupPrimerable
}
// NewRequestMoveGroupsRequest instantiates a new RequestMoveGroupsRequest and sets the default values.
func NewRequestMoveGroupsRequest()(*RequestMoveGroupsRequest) {
    m := &RequestMoveGroupsRequest{
        RequestAbstractOrganizationalUnitModificationRequest: *NewRequestAbstractOrganizationalUnitModificationRequest(),
    }
    typeEscapedValue := "request.MoveGroupsRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestMoveGroupsRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestMoveGroupsRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestMoveGroupsRequest(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestMoveGroupsRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestAbstractOrganizationalUnitModificationRequest.GetFieldDeserializers()
    res["groups"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]GroupGroupPrimerable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(GroupGroupPrimerable)
                }
            }
            m.SetGroups(res)
        }
        return nil
    }
    return res
}
// GetGroups gets the groups property value. The groups property
// returns a []GroupGroupPrimerable when successful
func (m *RequestMoveGroupsRequest) GetGroups()([]GroupGroupPrimerable) {
    return m.groups
}
// Serialize serializes information the current object
func (m *RequestMoveGroupsRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestAbstractOrganizationalUnitModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetGroups() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetGroups()))
        for i, v := range m.GetGroups() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("groups", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetGroups sets the groups property value. The groups property
func (m *RequestMoveGroupsRequest) SetGroups(value []GroupGroupPrimerable)() {
    m.groups = value
}
type RequestMoveGroupsRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestAbstractOrganizationalUnitModificationRequestable
    GetGroups()([]GroupGroupPrimerable)
    SetGroups(value []GroupGroupPrimerable)()
}
