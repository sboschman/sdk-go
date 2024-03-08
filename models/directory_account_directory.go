package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type DirectoryAccountDirectory struct {
    DirectoryAccountDirectoryPrimer
    // The additionalObjects property
    additionalObjects DirectoryAccountDirectory_additionalObjectsable
    // The baseOrganizationalUnit property
    baseOrganizationalUnit OrganizationOrganizationalUnitPrimerable
    // The defaultDirectory property
    defaultDirectory *bool
    // The helpdeskGroup property
    helpdeskGroup GroupGroupPrimerable
    // The restrict2fa property
    restrict2fa *bool
    // The rotatingPassword property
    rotatingPassword *DirectoryDirectoryRotatingPassword
    // The usernameCustomizable property
    usernameCustomizable *bool
}
// NewDirectoryAccountDirectory instantiates a new DirectoryAccountDirectory and sets the default values.
func NewDirectoryAccountDirectory()(*DirectoryAccountDirectory) {
    m := &DirectoryAccountDirectory{
        DirectoryAccountDirectoryPrimer: *NewDirectoryAccountDirectoryPrimer(),
    }
    typeEscapedValue := "directory.AccountDirectory"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateDirectoryAccountDirectoryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryAccountDirectoryFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "directory.InternalDirectory":
                        return NewDirectoryInternalDirectory(), nil
                    case "directory.LDAPDirectory":
                        return NewDirectoryLDAPDirectory(), nil
                    case "directory.MaintenanceDirectory":
                        return NewDirectoryMaintenanceDirectory(), nil
                    case "directory.OIDCDirectory":
                        return NewDirectoryOIDCDirectory(), nil
                }
            }
        }
    }
    return NewDirectoryAccountDirectory(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a DirectoryAccountDirectory_additionalObjectsable when successful
func (m *DirectoryAccountDirectory) GetAdditionalObjects()(DirectoryAccountDirectory_additionalObjectsable) {
    return m.additionalObjects
}
// GetBaseOrganizationalUnit gets the baseOrganizationalUnit property value. The baseOrganizationalUnit property
// returns a OrganizationOrganizationalUnitPrimerable when successful
func (m *DirectoryAccountDirectory) GetBaseOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable) {
    return m.baseOrganizationalUnit
}
// GetDefaultDirectory gets the defaultDirectory property value. The defaultDirectory property
// returns a *bool when successful
func (m *DirectoryAccountDirectory) GetDefaultDirectory()(*bool) {
    return m.defaultDirectory
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectoryAccountDirectory) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.DirectoryAccountDirectoryPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDirectoryAccountDirectory_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(DirectoryAccountDirectory_additionalObjectsable))
        }
        return nil
    }
    res["baseOrganizationalUnit"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseOrganizationalUnit(val.(OrganizationOrganizationalUnitPrimerable))
        }
        return nil
    }
    res["defaultDirectory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefaultDirectory(val)
        }
        return nil
    }
    res["helpdeskGroup"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHelpdeskGroup(val.(GroupGroupPrimerable))
        }
        return nil
    }
    res["restrict2fa"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRestrict2fa(val)
        }
        return nil
    }
    res["rotatingPassword"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDirectoryDirectoryRotatingPassword)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRotatingPassword(val.(*DirectoryDirectoryRotatingPassword))
        }
        return nil
    }
    res["usernameCustomizable"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUsernameCustomizable(val)
        }
        return nil
    }
    return res
}
// GetHelpdeskGroup gets the helpdeskGroup property value. The helpdeskGroup property
// returns a GroupGroupPrimerable when successful
func (m *DirectoryAccountDirectory) GetHelpdeskGroup()(GroupGroupPrimerable) {
    return m.helpdeskGroup
}
// GetRestrict2fa gets the restrict2fa property value. The restrict2fa property
// returns a *bool when successful
func (m *DirectoryAccountDirectory) GetRestrict2fa()(*bool) {
    return m.restrict2fa
}
// GetRotatingPassword gets the rotatingPassword property value. The rotatingPassword property
// returns a *DirectoryDirectoryRotatingPassword when successful
func (m *DirectoryAccountDirectory) GetRotatingPassword()(*DirectoryDirectoryRotatingPassword) {
    return m.rotatingPassword
}
// GetUsernameCustomizable gets the usernameCustomizable property value. The usernameCustomizable property
// returns a *bool when successful
func (m *DirectoryAccountDirectory) GetUsernameCustomizable()(*bool) {
    return m.usernameCustomizable
}
// Serialize serializes information the current object
func (m *DirectoryAccountDirectory) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.DirectoryAccountDirectoryPrimer.Serialize(writer)
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
        err = writer.WriteObjectValue("baseOrganizationalUnit", m.GetBaseOrganizationalUnit())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("defaultDirectory", m.GetDefaultDirectory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("helpdeskGroup", m.GetHelpdeskGroup())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("restrict2fa", m.GetRestrict2fa())
        if err != nil {
            return err
        }
    }
    if m.GetRotatingPassword() != nil {
        cast := (*m.GetRotatingPassword()).String()
        err = writer.WriteStringValue("rotatingPassword", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("usernameCustomizable", m.GetUsernameCustomizable())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *DirectoryAccountDirectory) SetAdditionalObjects(value DirectoryAccountDirectory_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetBaseOrganizationalUnit sets the baseOrganizationalUnit property value. The baseOrganizationalUnit property
func (m *DirectoryAccountDirectory) SetBaseOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)() {
    m.baseOrganizationalUnit = value
}
// SetDefaultDirectory sets the defaultDirectory property value. The defaultDirectory property
func (m *DirectoryAccountDirectory) SetDefaultDirectory(value *bool)() {
    m.defaultDirectory = value
}
// SetHelpdeskGroup sets the helpdeskGroup property value. The helpdeskGroup property
func (m *DirectoryAccountDirectory) SetHelpdeskGroup(value GroupGroupPrimerable)() {
    m.helpdeskGroup = value
}
// SetRestrict2fa sets the restrict2fa property value. The restrict2fa property
func (m *DirectoryAccountDirectory) SetRestrict2fa(value *bool)() {
    m.restrict2fa = value
}
// SetRotatingPassword sets the rotatingPassword property value. The rotatingPassword property
func (m *DirectoryAccountDirectory) SetRotatingPassword(value *DirectoryDirectoryRotatingPassword)() {
    m.rotatingPassword = value
}
// SetUsernameCustomizable sets the usernameCustomizable property value. The usernameCustomizable property
func (m *DirectoryAccountDirectory) SetUsernameCustomizable(value *bool)() {
    m.usernameCustomizable = value
}
type DirectoryAccountDirectoryable interface {
    DirectoryAccountDirectoryPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(DirectoryAccountDirectory_additionalObjectsable)
    GetBaseOrganizationalUnit()(OrganizationOrganizationalUnitPrimerable)
    GetDefaultDirectory()(*bool)
    GetHelpdeskGroup()(GroupGroupPrimerable)
    GetRestrict2fa()(*bool)
    GetRotatingPassword()(*DirectoryDirectoryRotatingPassword)
    GetUsernameCustomizable()(*bool)
    SetAdditionalObjects(value DirectoryAccountDirectory_additionalObjectsable)()
    SetBaseOrganizationalUnit(value OrganizationOrganizationalUnitPrimerable)()
    SetDefaultDirectory(value *bool)()
    SetHelpdeskGroup(value GroupGroupPrimerable)()
    SetRestrict2fa(value *bool)()
    SetRotatingPassword(value *DirectoryDirectoryRotatingPassword)()
    SetUsernameCustomizable(value *bool)()
}
