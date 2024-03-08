package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type AuditGroupAuditAccount struct {
    Linkable
    // The accountUuid property
    accountUuid *string
    // The accountValid property
    accountValid *bool
    // The action property
    action *AuditAuditAccountAction
    // The comment property
    comment *string
    // The disconnectedNested property
    disconnectedNested *bool
    // The displayName property
    displayName *string
    // The endDate property
    endDate *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The lastActive property
    lastActive *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The lastUsed property
    lastUsed *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly
    // The nested property
    nested *bool
    // The rights property
    rights *GroupGroupRights
    // The username property
    username *string
}
// NewAuditGroupAuditAccount instantiates a new AuditGroupAuditAccount and sets the default values.
func NewAuditGroupAuditAccount()(*AuditGroupAuditAccount) {
    m := &AuditGroupAuditAccount{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "audit.GroupAuditAccount"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateAuditGroupAuditAccountFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateAuditGroupAuditAccountFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuditGroupAuditAccount(), nil
}
// GetAccountUuid gets the accountUuid property value. The accountUuid property
// returns a *string when successful
func (m *AuditGroupAuditAccount) GetAccountUuid()(*string) {
    return m.accountUuid
}
// GetAccountValid gets the accountValid property value. The accountValid property
// returns a *bool when successful
func (m *AuditGroupAuditAccount) GetAccountValid()(*bool) {
    return m.accountValid
}
// GetAction gets the action property value. The action property
// returns a *AuditAuditAccountAction when successful
func (m *AuditGroupAuditAccount) GetAction()(*AuditAuditAccountAction) {
    return m.action
}
// GetComment gets the comment property value. The comment property
// returns a *string when successful
func (m *AuditGroupAuditAccount) GetComment()(*string) {
    return m.comment
}
// GetDisconnectedNested gets the disconnectedNested property value. The disconnectedNested property
// returns a *bool when successful
func (m *AuditGroupAuditAccount) GetDisconnectedNested()(*bool) {
    return m.disconnectedNested
}
// GetDisplayName gets the displayName property value. The displayName property
// returns a *string when successful
func (m *AuditGroupAuditAccount) GetDisplayName()(*string) {
    return m.displayName
}
// GetEndDate gets the endDate property value. The endDate property
// returns a *DateOnly when successful
func (m *AuditGroupAuditAccount) GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.endDate
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *AuditGroupAuditAccount) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["accountUuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountUuid(val)
        }
        return nil
    }
    res["accountValid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountValid(val)
        }
        return nil
    }
    res["action"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuditAuditAccountAction)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAction(val.(*AuditAuditAccountAction))
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    res["disconnectedNested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisconnectedNested(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
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
    res["lastActive"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastActive(val)
        }
        return nil
    }
    res["lastUsed"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetDateOnlyValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastUsed(val)
        }
        return nil
    }
    res["nested"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNested(val)
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
    res["username"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsername(val)
        }
        return nil
    }
    return res
}
// GetLastActive gets the lastActive property value. The lastActive property
// returns a *Time when successful
func (m *AuditGroupAuditAccount) GetLastActive()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastActive
}
// GetLastUsed gets the lastUsed property value. The lastUsed property
// returns a *DateOnly when successful
func (m *AuditGroupAuditAccount) GetLastUsed()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly) {
    return m.lastUsed
}
// GetNested gets the nested property value. The nested property
// returns a *bool when successful
func (m *AuditGroupAuditAccount) GetNested()(*bool) {
    return m.nested
}
// GetRights gets the rights property value. The rights property
// returns a *GroupGroupRights when successful
func (m *AuditGroupAuditAccount) GetRights()(*GroupGroupRights) {
    return m.rights
}
// GetUsername gets the username property value. The username property
// returns a *string when successful
func (m *AuditGroupAuditAccount) GetUsername()(*string) {
    return m.username
}
// Serialize serializes information the current object
func (m *AuditGroupAuditAccount) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("accountUuid", m.GetAccountUuid())
        if err != nil {
            return err
        }
    }
    if m.GetAction() != nil {
        cast := (*m.GetAction()).String()
        err = writer.WriteStringValue("action", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("comment", m.GetComment())
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
    return nil
}
// SetAccountUuid sets the accountUuid property value. The accountUuid property
func (m *AuditGroupAuditAccount) SetAccountUuid(value *string)() {
    m.accountUuid = value
}
// SetAccountValid sets the accountValid property value. The accountValid property
func (m *AuditGroupAuditAccount) SetAccountValid(value *bool)() {
    m.accountValid = value
}
// SetAction sets the action property value. The action property
func (m *AuditGroupAuditAccount) SetAction(value *AuditAuditAccountAction)() {
    m.action = value
}
// SetComment sets the comment property value. The comment property
func (m *AuditGroupAuditAccount) SetComment(value *string)() {
    m.comment = value
}
// SetDisconnectedNested sets the disconnectedNested property value. The disconnectedNested property
func (m *AuditGroupAuditAccount) SetDisconnectedNested(value *bool)() {
    m.disconnectedNested = value
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *AuditGroupAuditAccount) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetEndDate sets the endDate property value. The endDate property
func (m *AuditGroupAuditAccount) SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.endDate = value
}
// SetLastActive sets the lastActive property value. The lastActive property
func (m *AuditGroupAuditAccount) SetLastActive(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastActive = value
}
// SetLastUsed sets the lastUsed property value. The lastUsed property
func (m *AuditGroupAuditAccount) SetLastUsed(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)() {
    m.lastUsed = value
}
// SetNested sets the nested property value. The nested property
func (m *AuditGroupAuditAccount) SetNested(value *bool)() {
    m.nested = value
}
// SetRights sets the rights property value. The rights property
func (m *AuditGroupAuditAccount) SetRights(value *GroupGroupRights)() {
    m.rights = value
}
// SetUsername sets the username property value. The username property
func (m *AuditGroupAuditAccount) SetUsername(value *string)() {
    m.username = value
}
type AuditGroupAuditAccountable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAccountUuid()(*string)
    GetAccountValid()(*bool)
    GetAction()(*AuditAuditAccountAction)
    GetComment()(*string)
    GetDisconnectedNested()(*bool)
    GetDisplayName()(*string)
    GetEndDate()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetLastActive()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetLastUsed()(*i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)
    GetNested()(*bool)
    GetRights()(*GroupGroupRights)
    GetUsername()(*string)
    SetAccountUuid(value *string)()
    SetAccountValid(value *bool)()
    SetAction(value *AuditAuditAccountAction)()
    SetComment(value *string)()
    SetDisconnectedNested(value *bool)()
    SetDisplayName(value *string)()
    SetEndDate(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetLastActive(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetLastUsed(value *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)()
    SetNested(value *bool)()
    SetRights(value *GroupGroupRights)()
    SetUsername(value *string)()
}
