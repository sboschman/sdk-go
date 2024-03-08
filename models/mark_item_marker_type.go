package models
import (
    "errors"
)
type MarkItemMarkerType int

const (
    TLS_DISABLED_MARKITEMMARKERTYPE MarkItemMarkerType = iota
    TLS_UNVERIFIED_MARKITEMMARKERTYPE
    WEAK_PASSWORD_HASHING_MARKITEMMARKERTYPE
    GROUP_NO_MANAGER_MARKITEMMARKERTYPE
    GROUP_ONE_MANAGER_MARKITEMMARKERTYPE
    GROUP_UNREADABLE_VAULT_MARKITEMMARKERTYPE
    GROUP_UNRECOVERABLE_VAULT_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_AUDIT_INTERVAL_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_AUDIT_MONTHS_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_PROVISIONING_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_MEMBERSHIP_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_DELEGATION_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_AUDITING_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_RECORD_TRAIL_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_ROTATING_PASSWORD_REQUIRED_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_VAULT_REQUIRES_ACTIVATION_MARKITEMMARKERTYPE
    GROUP_CRITERIA_UNSATISFIED_MINIMUM_NR_MANAGERS_MARKITEMMARKERTYPE
)

func (i MarkItemMarkerType) String() string {
    return []string{"TLS_DISABLED", "TLS_UNVERIFIED", "WEAK_PASSWORD_HASHING", "GROUP_NO_MANAGER", "GROUP_ONE_MANAGER", "GROUP_UNREADABLE_VAULT", "GROUP_UNRECOVERABLE_VAULT", "GROUP_CRITERIA_UNSATISFIED_AUDIT_INTERVAL", "GROUP_CRITERIA_UNSATISFIED_AUDIT_MONTHS", "GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_PROVISIONING", "GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_MEMBERSHIP", "GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_DELEGATION", "GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_AUDITING", "GROUP_CRITERIA_UNSATISFIED_RECORD_TRAIL", "GROUP_CRITERIA_UNSATISFIED_ROTATING_PASSWORD_REQUIRED", "GROUP_CRITERIA_UNSATISFIED_VAULT_REQUIRES_ACTIVATION", "GROUP_CRITERIA_UNSATISFIED_MINIMUM_NR_MANAGERS"}[i]
}
func ParseMarkItemMarkerType(v string) (any, error) {
    result := TLS_DISABLED_MARKITEMMARKERTYPE
    switch v {
        case "TLS_DISABLED":
            result = TLS_DISABLED_MARKITEMMARKERTYPE
        case "TLS_UNVERIFIED":
            result = TLS_UNVERIFIED_MARKITEMMARKERTYPE
        case "WEAK_PASSWORD_HASHING":
            result = WEAK_PASSWORD_HASHING_MARKITEMMARKERTYPE
        case "GROUP_NO_MANAGER":
            result = GROUP_NO_MANAGER_MARKITEMMARKERTYPE
        case "GROUP_ONE_MANAGER":
            result = GROUP_ONE_MANAGER_MARKITEMMARKERTYPE
        case "GROUP_UNREADABLE_VAULT":
            result = GROUP_UNREADABLE_VAULT_MARKITEMMARKERTYPE
        case "GROUP_UNRECOVERABLE_VAULT":
            result = GROUP_UNRECOVERABLE_VAULT_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_AUDIT_INTERVAL":
            result = GROUP_CRITERIA_UNSATISFIED_AUDIT_INTERVAL_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_AUDIT_MONTHS":
            result = GROUP_CRITERIA_UNSATISFIED_AUDIT_MONTHS_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_PROVISIONING":
            result = GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_PROVISIONING_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_MEMBERSHIP":
            result = GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_MEMBERSHIP_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_DELEGATION":
            result = GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_DELEGATION_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_AUDITING":
            result = GROUP_CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_AUDITING_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_RECORD_TRAIL":
            result = GROUP_CRITERIA_UNSATISFIED_RECORD_TRAIL_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_ROTATING_PASSWORD_REQUIRED":
            result = GROUP_CRITERIA_UNSATISFIED_ROTATING_PASSWORD_REQUIRED_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_VAULT_REQUIRES_ACTIVATION":
            result = GROUP_CRITERIA_UNSATISFIED_VAULT_REQUIRES_ACTIVATION_MARKITEMMARKERTYPE
        case "GROUP_CRITERIA_UNSATISFIED_MINIMUM_NR_MANAGERS":
            result = GROUP_CRITERIA_UNSATISFIED_MINIMUM_NR_MANAGERS_MARKITEMMARKERTYPE
        default:
            return 0, errors.New("Unknown MarkItemMarkerType value: " + v)
    }
    return &result, nil
}
func SerializeMarkItemMarkerType(values []MarkItemMarkerType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MarkItemMarkerType) isMultiValue() bool {
    return false
}
