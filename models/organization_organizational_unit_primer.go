package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type OrganizationOrganizationalUnitPrimer struct {
    Linkable
    // The name property
    name *string
    // The uuid property
    uuid *string
}
// NewOrganizationOrganizationalUnitPrimer instantiates a new OrganizationOrganizationalUnitPrimer and sets the default values.
func NewOrganizationOrganizationalUnitPrimer()(*OrganizationOrganizationalUnitPrimer) {
    m := &OrganizationOrganizationalUnitPrimer{
        Linkable: *NewLinkable(),
    }
    typeEscapedValue := "organization.OrganizationalUnitPrimer"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateOrganizationOrganizationalUnitPrimerFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "organization.AccountOrganizationalUnit":
                        return NewOrganizationAccountOrganizationalUnit(), nil
                    case "organization.OrganizationalUnit":
                        return NewOrganizationOrganizationalUnit(), nil
                }
            }
        }
    }
    return NewOrganizationOrganizationalUnitPrimer(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *OrganizationOrganizationalUnitPrimer) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Linkable.GetFieldDeserializers()
    res["name"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["uuid"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUuid(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name property
// returns a *string when successful
func (m *OrganizationOrganizationalUnitPrimer) GetName()(*string) {
    return m.name
}
// GetUuid gets the uuid property value. The uuid property
// returns a *string when successful
func (m *OrganizationOrganizationalUnitPrimer) GetUuid()(*string) {
    return m.uuid
}
// Serialize serializes information the current object
func (m *OrganizationOrganizationalUnitPrimer) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Linkable.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name property
func (m *OrganizationOrganizationalUnitPrimer) SetName(value *string)() {
    m.name = value
}
// SetUuid sets the uuid property value. The uuid property
func (m *OrganizationOrganizationalUnitPrimer) SetUuid(value *string)() {
    m.uuid = value
}
type OrganizationOrganizationalUnitPrimerable interface {
    Linkableable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetName()(*string)
    GetUuid()(*string)
    SetName(value *string)()
    SetUuid(value *string)()
}
