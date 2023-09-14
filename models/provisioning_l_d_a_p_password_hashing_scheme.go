package models
import (
    "errors"
)
// 
type ProvisioningLDAPPasswordHashingScheme int

const (
    SSHA_PROVISIONINGLDAPPASSWORDHASHINGSCHEME ProvisioningLDAPPasswordHashingScheme = iota
    PBKDF2_PROVISIONINGLDAPPASSWORDHASHINGSCHEME
)

func (i ProvisioningLDAPPasswordHashingScheme) String() string {
    return []string{"SSHA", "PBKDF2"}[i]
}
func ParseProvisioningLDAPPasswordHashingScheme(v string) (any, error) {
    result := SSHA_PROVISIONINGLDAPPASSWORDHASHINGSCHEME
    switch v {
        case "SSHA":
            result = SSHA_PROVISIONINGLDAPPASSWORDHASHINGSCHEME
        case "PBKDF2":
            result = PBKDF2_PROVISIONINGLDAPPASSWORDHASHINGSCHEME
        default:
            return 0, errors.New("Unknown ProvisioningLDAPPasswordHashingScheme value: " + v)
    }
    return &result, nil
}
func SerializeProvisioningLDAPPasswordHashingScheme(values []ProvisioningLDAPPasswordHashingScheme) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i ProvisioningLDAPPasswordHashingScheme) isMultiValue() bool {
    return false
}
