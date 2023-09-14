package models
import (
    "errors"
)
// 
type AuthSecurityLevel int

const (
    LOW_AUTHSECURITYLEVEL AuthSecurityLevel = iota
    MEDIUM_AUTHSECURITYLEVEL
    HIGH_AUTHSECURITYLEVEL
)

func (i AuthSecurityLevel) String() string {
    return []string{"LOW", "MEDIUM", "HIGH"}[i]
}
func ParseAuthSecurityLevel(v string) (any, error) {
    result := LOW_AUTHSECURITYLEVEL
    switch v {
        case "LOW":
            result = LOW_AUTHSECURITYLEVEL
        case "MEDIUM":
            result = MEDIUM_AUTHSECURITYLEVEL
        case "HIGH":
            result = HIGH_AUTHSECURITYLEVEL
        default:
            return 0, errors.New("Unknown AuthSecurityLevel value: " + v)
    }
    return &result, nil
}
func SerializeAuthSecurityLevel(values []AuthSecurityLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthSecurityLevel) isMultiValue() bool {
    return false
}
