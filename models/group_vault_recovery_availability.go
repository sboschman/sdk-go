package models
import (
    "errors"
)
type GroupVaultRecoveryAvailability int

const (
    NONE_GROUPVAULTRECOVERYAVAILABILITY GroupVaultRecoveryAvailability = iota
    RECOVERY_KEY_ONLY_GROUPVAULTRECOVERYAVAILABILITY
    FULL_GROUPVAULTRECOVERYAVAILABILITY
)

func (i GroupVaultRecoveryAvailability) String() string {
    return []string{"NONE", "RECOVERY_KEY_ONLY", "FULL"}[i]
}
func ParseGroupVaultRecoveryAvailability(v string) (any, error) {
    result := NONE_GROUPVAULTRECOVERYAVAILABILITY
    switch v {
        case "NONE":
            result = NONE_GROUPVAULTRECOVERYAVAILABILITY
        case "RECOVERY_KEY_ONLY":
            result = RECOVERY_KEY_ONLY_GROUPVAULTRECOVERYAVAILABILITY
        case "FULL":
            result = FULL_GROUPVAULTRECOVERYAVAILABILITY
        default:
            return 0, errors.New("Unknown GroupVaultRecoveryAvailability value: " + v)
    }
    return &result, nil
}
func SerializeGroupVaultRecoveryAvailability(values []GroupVaultRecoveryAvailability) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i GroupVaultRecoveryAvailability) isMultiValue() bool {
    return false
}
