package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

type ClientClientApplication struct {
    ClientClientApplicationPrimer
    // The additionalObjects property
    additionalObjects ClientClientApplication_additionalObjectsable
    // The lastModifiedAt property
    lastModifiedAt *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The owner property
    owner GroupGroupPrimerable
    // The technicalAdministrator property
    technicalAdministrator GroupGroupPrimerable
}
// NewClientClientApplication instantiates a new ClientClientApplication and sets the default values.
func NewClientClientApplication()(*ClientClientApplication) {
    m := &ClientClientApplication{
        ClientClientApplicationPrimer: *NewClientClientApplicationPrimer(),
    }
    typeEscapedValue := "client.ClientApplication"
    m.SetTypeEscaped(&typeEscapedValue)
    return m
}
// CreateClientClientApplicationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateClientClientApplicationFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
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
                    case "client.LdapClient":
                        return NewClientLdapClient(), nil
                    case "client.OAuth2Client":
                        return NewClientOAuth2Client(), nil
                    case "client.Saml2Client":
                        return NewClientSaml2Client(), nil
                }
            }
        }
    }
    return NewClientClientApplication(), nil
}
// GetAdditionalObjects gets the additionalObjects property value. The additionalObjects property
// returns a ClientClientApplication_additionalObjectsable when successful
func (m *ClientClientApplication) GetAdditionalObjects()(ClientClientApplication_additionalObjectsable) {
    return m.additionalObjects
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *ClientClientApplication) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.ClientClientApplicationPrimer.GetFieldDeserializers()
    res["additionalObjects"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateClientClientApplication_additionalObjectsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAdditionalObjects(val.(ClientClientApplication_additionalObjectsable))
        }
        return nil
    }
    res["lastModifiedAt"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedAt(val)
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
    res["technicalAdministrator"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateGroupGroupPrimerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTechnicalAdministrator(val.(GroupGroupPrimerable))
        }
        return nil
    }
    return res
}
// GetLastModifiedAt gets the lastModifiedAt property value. The lastModifiedAt property
// returns a *Time when successful
func (m *ClientClientApplication) GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.lastModifiedAt
}
// GetOwner gets the owner property value. The owner property
// returns a GroupGroupPrimerable when successful
func (m *ClientClientApplication) GetOwner()(GroupGroupPrimerable) {
    return m.owner
}
// GetTechnicalAdministrator gets the technicalAdministrator property value. The technicalAdministrator property
// returns a GroupGroupPrimerable when successful
func (m *ClientClientApplication) GetTechnicalAdministrator()(GroupGroupPrimerable) {
    return m.technicalAdministrator
}
// Serialize serializes information the current object
func (m *ClientClientApplication) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.ClientClientApplicationPrimer.Serialize(writer)
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
        err = writer.WriteObjectValue("technicalAdministrator", m.GetTechnicalAdministrator())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalObjects sets the additionalObjects property value. The additionalObjects property
func (m *ClientClientApplication) SetAdditionalObjects(value ClientClientApplication_additionalObjectsable)() {
    m.additionalObjects = value
}
// SetLastModifiedAt sets the lastModifiedAt property value. The lastModifiedAt property
func (m *ClientClientApplication) SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedAt = value
}
// SetOwner sets the owner property value. The owner property
func (m *ClientClientApplication) SetOwner(value GroupGroupPrimerable)() {
    m.owner = value
}
// SetTechnicalAdministrator sets the technicalAdministrator property value. The technicalAdministrator property
func (m *ClientClientApplication) SetTechnicalAdministrator(value GroupGroupPrimerable)() {
    m.technicalAdministrator = value
}
type ClientClientApplicationable interface {
    ClientClientApplicationPrimerable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAdditionalObjects()(ClientClientApplication_additionalObjectsable)
    GetLastModifiedAt()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOwner()(GroupGroupPrimerable)
    GetTechnicalAdministrator()(GroupGroupPrimerable)
    SetAdditionalObjects(value ClientClientApplication_additionalObjectsable)()
    SetLastModifiedAt(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOwner(value GroupGroupPrimerable)()
    SetTechnicalAdministrator(value GroupGroupPrimerable)()
}
