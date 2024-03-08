package item
import (
    "errors"
)
type PutAdditionalQueryParameterType int

const (
    ACCEPTPARAMS_PUTADDITIONALQUERYPARAMETERTYPE PutAdditionalQueryParameterType = iota
    AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
    RESETSTATUS_PUTADDITIONALQUERYPARAMETERTYPE
)

func (i PutAdditionalQueryParameterType) String() string {
    return []string{"acceptParams", "audit", "resetStatus"}[i]
}
func ParsePutAdditionalQueryParameterType(v string) (any, error) {
    result := ACCEPTPARAMS_PUTADDITIONALQUERYPARAMETERTYPE
    switch v {
        case "acceptParams":
            result = ACCEPTPARAMS_PUTADDITIONALQUERYPARAMETERTYPE
        case "audit":
            result = AUDIT_PUTADDITIONALQUERYPARAMETERTYPE
        case "resetStatus":
            result = RESETSTATUS_PUTADDITIONALQUERYPARAMETERTYPE
        default:
            return 0, errors.New("Unknown PutAdditionalQueryParameterType value: " + v)
    }
    return &result, nil
}
func SerializePutAdditionalQueryParameterType(values []PutAdditionalQueryParameterType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i PutAdditionalQueryParameterType) isMultiValue() bool {
    return false
}
