package models
import (
    "errors"
)
type GroupGroupEditRequiredReason int

const (
    ONE_MANAGER_GROUPGROUPEDITREQUIREDREASON GroupGroupEditRequiredReason = iota
    CRITERIA_UNSATISFIED_AUDIT_INTERVAL_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_AUDIT_MONTHS_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_PROVISIONING_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_MEMBERSHIP_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_DELEGATION_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_AUDITING_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_RECORD_TRAIL_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_ROTATING_PASSWORD_REQUIRED_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_VAULT_REQUIRES_ACTIVATION_GROUPGROUPEDITREQUIREDREASON
    CRITERIA_UNSATISFIED_MINIMUM_NR_MANAGERS_GROUPGROUPEDITREQUIREDREASON
)

func (i GroupGroupEditRequiredReason) String() string {
    return []string{"ONE_MANAGER", "CRITERIA_UNSATISFIED_AUDIT_INTERVAL", "CRITERIA_UNSATISFIED_AUDIT_MONTHS", "CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_PROVISIONING", "CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_MEMBERSHIP", "CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_DELEGATION", "CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_AUDITING", "CRITERIA_UNSATISFIED_RECORD_TRAIL", "CRITERIA_UNSATISFIED_ROTATING_PASSWORD_REQUIRED", "CRITERIA_UNSATISFIED_VAULT_REQUIRES_ACTIVATION", "CRITERIA_UNSATISFIED_MINIMUM_NR_MANAGERS"}[i]
}
func ParseGroupGroupEditRequiredReason(v string) (any, error) {
    result := ONE_MANAGER_GROUPGROUPEDITREQUIREDREASON
    switch v {
        case "ONE_MANAGER":
            result = ONE_MANAGER_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_AUDIT_INTERVAL":
            result = CRITERIA_UNSATISFIED_AUDIT_INTERVAL_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_AUDIT_MONTHS":
            result = CRITERIA_UNSATISFIED_AUDIT_MONTHS_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_PROVISIONING":
            result = CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_PROVISIONING_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_MEMBERSHIP":
            result = CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_MEMBERSHIP_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_DELEGATION":
            result = CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_DELEGATION_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_AUDITING":
            result = CRITERIA_UNSATISFIED_AUTHORIZING_GROUP_AUDITING_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_RECORD_TRAIL":
            result = CRITERIA_UNSATISFIED_RECORD_TRAIL_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_ROTATING_PASSWORD_REQUIRED":
            result = CRITERIA_UNSATISFIED_ROTATING_PASSWORD_REQUIRED_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_VAULT_REQUIRES_ACTIVATION":
            result = CRITERIA_UNSATISFIED_VAULT_REQUIRES_ACTIVATION_GROUPGROUPEDITREQUIREDREASON
        case "CRITERIA_UNSATISFIED_MINIMUM_NR_MANAGERS":
            result = CRITERIA_UNSATISFIED_MINIMUM_NR_MANAGERS_GROUPGROUPEDITREQUIREDREASON
        default:
            return 0, errors.New("Unknown GroupGroupEditRequiredReason value: " + v)
    }
    return &result, nil
}
func SerializeGroupGroupEditRequiredReason(values []GroupGroupEditRequiredReason) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupGroupEditRequiredReason) isMultiValue() bool {
    return false
}
