package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type RequestUpdateGroupMembershipRequest struct {
    RequestModificationRequest
    // The accountToUpdate property
    accountToUpdate AuthAccountPrimerable
    // The currentEndDate property
    currentEndDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The currentRights property
    currentRights *GroupGroupRights
    // The endDate property
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The rights property
    rights *GroupGroupRights
    // The updateGroupMembershipType property
    updateGroupMembershipType *RequestUpdateGroupMembershipType
}
// NewRequestUpdateGroupMembershipRequest instantiates a new RequestUpdateGroupMembershipRequest and sets the default values.
func NewRequestUpdateGroupMembershipRequest()(*RequestUpdateGroupMembershipRequest) {
    m := &RequestUpdateGroupMembershipRequest{
        RequestModificationRequest: *NewRequestModificationRequest(),
    }
    typeEscapedValue := "request.UpdateGroupMembershipRequest"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateRequestUpdateGroupMembershipRequestFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateRequestUpdateGroupMembershipRequestFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewRequestUpdateGroupMembershipRequest(), nil
}
// GetAccountToUpdate gets the accountToUpdate property value. The accountToUpdate property
// returns a AuthAccountPrimerable when successful
func (m *RequestUpdateGroupMembershipRequest) GetAccountToUpdate()(AuthAccountPrimerable) {
    return m.accountToUpdate
}
// GetCurrentEndDate gets the currentEndDate property value. The currentEndDate property
// returns a *DateOnly when successful
func (m *RequestUpdateGroupMembershipRequest) GetCurrentEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.currentEndDate
}
// GetCurrentRights gets the currentRights property value. The currentRights property
// returns a *GroupGroupRights when successful
func (m *RequestUpdateGroupMembershipRequest) GetCurrentRights()(*GroupGroupRights) {
    return m.currentRights
}
// GetEndDate gets the endDate property value. The endDate property
// returns a *DateOnly when successful
func (m *RequestUpdateGroupMembershipRequest) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *RequestUpdateGroupMembershipRequest) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.RequestModificationRequest.GetFieldDeserializers()
    res["accountToUpdate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateAuthAccountPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountToUpdate(val.(AuthAccountPrimerable))
        }
        return nil
    }
    res["currentEndDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentEndDate(val)
        }
        return nil
    }
    res["currentRights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupGroupRights)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCurrentRights(val.(*GroupGroupRights))
        }
        return nil
    }
    res["endDate"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["rights"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseGroupGroupRights)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRights(val.(*GroupGroupRights))
        }
        return nil
    }
    res["updateGroupMembershipType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseRequestUpdateGroupMembershipType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUpdateGroupMembershipType(val.(*RequestUpdateGroupMembershipType))
        }
        return nil
    }
    return res
}
// GetRights gets the rights property value. The rights property
// returns a *GroupGroupRights when successful
func (m *RequestUpdateGroupMembershipRequest) GetRights()(*GroupGroupRights) {
    return m.rights
}
// GetUpdateGroupMembershipType gets the updateGroupMembershipType property value. The updateGroupMembershipType property
// returns a *RequestUpdateGroupMembershipType when successful
func (m *RequestUpdateGroupMembershipRequest) GetUpdateGroupMembershipType()(*RequestUpdateGroupMembershipType) {
    return m.updateGroupMembershipType
}
// Serialize serializes information the current object
func (m *RequestUpdateGroupMembershipRequest) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.RequestModificationRequest.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("accountToUpdate", m.GetAccountToUpdate())
        if err != nil {
            return err
        }
    }
    if m.GetCurrentRights() != nil {
        cast := (*m.GetCurrentRights()).String()
        err = writer.WriteStringValue("currentRights", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteDateOnlyValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    if m.GetRights() != nil {
        cast := (*m.GetRights()).String()
        err = writer.WriteStringValue("rights", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetUpdateGroupMembershipType() != nil {
        cast := (*m.GetUpdateGroupMembershipType()).String()
        err = writer.WriteStringValue("updateGroupMembershipType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccountToUpdate sets the accountToUpdate property value. The accountToUpdate property
func (m *RequestUpdateGroupMembershipRequest) SetAccountToUpdate(value AuthAccountPrimerable)() {
    m.accountToUpdate = value
}
// SetCurrentEndDate sets the currentEndDate property value. The currentEndDate property
func (m *RequestUpdateGroupMembershipRequest) SetCurrentEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.currentEndDate = value
}
// SetCurrentRights sets the currentRights property value. The currentRights property
func (m *RequestUpdateGroupMembershipRequest) SetCurrentRights(value *GroupGroupRights)() {
    m.currentRights = value
}
// SetEndDate sets the endDate property value. The endDate property
func (m *RequestUpdateGroupMembershipRequest) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetRights sets the rights property value. The rights property
func (m *RequestUpdateGroupMembershipRequest) SetRights(value *GroupGroupRights)() {
    m.rights = value
}
// SetUpdateGroupMembershipType sets the updateGroupMembershipType property value. The updateGroupMembershipType property
func (m *RequestUpdateGroupMembershipRequest) SetUpdateGroupMembershipType(value *RequestUpdateGroupMembershipType)() {
    m.updateGroupMembershipType = value
}
type RequestUpdateGroupMembershipRequestable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    RequestModificationRequestable
    GetAccountToUpdate()(AuthAccountPrimerable)
    GetCurrentEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetCurrentRights()(*GroupGroupRights)
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetRights()(*GroupGroupRights)
    GetUpdateGroupMembershipType()(*RequestUpdateGroupMembershipType)
    SetAccountToUpdate(value AuthAccountPrimerable)()
    SetCurrentEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetCurrentRights(value *GroupGroupRights)()
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetRights(value *GroupGroupRights)()
    SetUpdateGroupMembershipType(value *RequestUpdateGroupMembershipType)()
}
